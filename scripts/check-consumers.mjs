import { mkdtemp, rm } from "node:fs/promises";
import { existsSync, mkdirSync, readFileSync, writeFileSync } from "node:fs";
import { tmpdir } from "node:os";
import { basename, dirname, join, resolve } from "node:path";
import { fileURLToPath } from "node:url";
import { spawnSync } from "node:child_process";

const contractsRoot = resolve(fileURLToPath(new URL("..", import.meta.url)));
const consumerRoot = resolve(process.env.SOHA_CONSUMER_ROOT ?? dirname(contractsRoot));

const consumers = [
  {
    name: "soha",
    path: "soha",
    kind: "go",
    commands: [["go", ["test", "./..."]]],
  },
  {
    name: "soha-web",
    path: "soha-web",
    kind: "npm",
    commands: [
      ["npm", ["ci"]],
      ["npm", ["run", "build"]],
    ],
  },
  {
    name: "soha-cli",
    path: "soha-cli",
    kind: "go",
    commands: [["go", ["test", "./..."]]],
  },
  {
    name: "soha-agent",
    path: "soha-agent",
    kind: "go",
    commands: [["go", ["test", "./..."]]],
  },
];

const args = process.argv.slice(2);
const requireAll = args.includes("--require-all") || process.env.SOHA_CONSUMER_REQUIRE_ALL === "1";
const dryRun = args.includes("--dry-run");
const listOnly = args.includes("--list");
const outputPath = valueArg("--out");
const selected = selectedConsumers(args);

if (listOnly) {
  console.log(consumers.map((consumer) => consumer.name).join("\n"));
  process.exit(0);
}

const targets = selected.length > 0 ? consumers.filter((consumer) => selected.includes(consumer.name)) : consumers;
const unknown = selected.filter((name) => !consumers.some((consumer) => consumer.name === name));
if (unknown.length > 0) {
  throw new Error(`unknown consumer(s): ${unknown.join(", ")}`);
}

let checked = 0;
let skipped = 0;
const results = [];
for (const consumer of targets) {
  const dir = resolve(consumerRoot, consumer.path);
  if (!existsSync(dir)) {
    const message = `consumer ${consumer.name} not found at ${dir}`;
    if (requireAll || selected.includes(consumer.name)) {
      throw new Error(message);
    }
    console.warn(`warning: ${message}; skipping`);
    results.push({
      consumer: consumer.name,
      commands: consumer.commands,
      path: dir,
      reason: "missing checkout",
      status: "skipped",
    });
    skipped += 1;
    continue;
  }

  if (dryRun) {
    console.log(`[dry-run] ${consumer.name}: ${dir}`);
    for (const [command, commandArgs] of consumer.commands) {
      console.log(`[dry-run]   ${command} ${commandArgs.join(" ")}`);
    }
    results.push({
      consumer: consumer.name,
      commands: consumer.commands,
      path: dir,
      status: "dry-run",
    });
    checked += 1;
    continue;
  }

  console.log(`checking consumer ${consumer.name} in ${dir}`);
  const env = consumer.kind === "go" ? await goWorkspaceEnv(dir) : process.env;
  try {
    for (const [command, commandArgs] of consumer.commands) {
      run(command, commandArgs, dir, env);
    }
  } finally {
    if (env !== process.env && env.GOWORK) {
      await rm(dirname(env.GOWORK), { recursive: true, force: true });
    }
  }
  results.push({
    consumer: consumer.name,
    commands: consumer.commands,
    path: dir,
    status: "passed",
  });
  checked += 1;
}

if (requireAll && checked !== targets.length) {
  throw new Error(`consumer matrix checked ${checked}/${targets.length} required consumer(s)`);
}

console.log(`consumer matrix complete: checked ${checked}, skipped ${skipped}, root ${consumerRoot}`);

if (outputPath) {
  const resolvedOutputPath = resolve(outputPath);
  mkdirSync(dirname(resolvedOutputPath), { recursive: true });
  writeFileSync(
    resolvedOutputPath,
    `${JSON.stringify(
      {
        schemaVersion: 1,
        contractsVersion: packageVersion(),
        contractsRoot,
        consumerRoot,
        requireAll,
        dryRun,
        checked,
        skipped,
        results,
      },
      null,
      2,
    )}\n`,
  );
}

function selectedConsumers(argv) {
  const names = [];
  for (let i = 0; i < argv.length; i += 1) {
    if (argv[i] === "--consumer" && argv[i + 1]) {
      names.push(argv[i + 1]);
      i += 1;
    }
  }
  if (process.env.SOHA_CONSUMER) {
    names.push(...process.env.SOHA_CONSUMER.split(",").map((name) => name.trim()).filter(Boolean));
  }
  return [...new Set(names)];
}

function valueArg(name) {
  const prefix = `${name}=`;
  const value = args.find((arg) => arg.startsWith(prefix));
  if (value) {
    return value.slice(prefix.length);
  }
  const index = args.indexOf(name);
  return index >= 0 ? args[index + 1] : undefined;
}

function packageVersion() {
  const packageJson = JSON.parse(readFileSync(join(contractsRoot, "package.json"), "utf8"));
  return packageJson.version;
}

async function goWorkspaceEnv(consumerDir) {
  const workDir = await mkdtemp(join(tmpdir(), "opensoha-consumer-work-"));
  const goWork = join(workDir, "go.work");
  run("go", ["work", "init", consumerDir, contractsRoot], workDir, process.env);
  return {
    ...process.env,
    GOWORK: goWork,
  };
}

function run(command, args, cwd, env) {
  const label = `${basename(cwd)}$ ${command} ${args.join(" ")}`;
  console.log(label);
  const result = spawnSync(command, args, {
    cwd,
    env,
    encoding: "utf8",
    stdio: "inherit",
  });
  if (result.status !== 0) {
    throw new Error(`${label} failed with status ${result.status}`);
  }
}
