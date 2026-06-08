import { mkdir, readFile, writeFile } from "node:fs/promises";
import { dirname } from "node:path";
import { fileURLToPath } from "node:url";
import { parse } from "yaml";

const root = new URL("..", import.meta.url);
const checkOnly = process.argv.includes("--check");

const openapi = parse(await readFile(new URL("openapi/soha-api.yaml", root), "utf8"));
const schemas = openapi?.components?.schemas ?? {};
const orderedNames = Object.keys(schemas);

const generated = new Map([
  ["gen/go/sohaapi/types.go", generateGoTypes(orderedNames, schemas)],
  ["gen/ts/sohaapi/index.ts", generateTSTypes(orderedNames, schemas)],
]);

for (const [path, content] of generated) {
  const url = new URL(path, root);
  if (checkOnly) {
    const current = await readFile(url, "utf8").catch(() => "");
    if (current !== content) {
      throw new Error(`${path} is out of date. Run npm run generate.`);
    }
    continue;
  }
  await mkdir(dirname(fileURLToPath(url)), { recursive: true });
  await writeFile(url, content);
}

function generateTSTypes(names, schemaMap) {
  const lines = [
    "// Code generated from OpenSoha contracts. DO NOT EDIT.",
    "",
    "export type AnyObject = Record<string, unknown>",
    "export type ApiResponse<T = unknown> = { data: T }",
    "export type ApiItemsResponse<T = unknown> = { items: T[] }",
    "export type AuthTokens = TokenSet",
    "",
  ];

  for (const name of names) {
    const schema = schemaMap[name] ?? {};
    if (name === "AnyValue") {
      lines.push("export type AnyValue = unknown", "");
      continue;
    }
    if (name === "GenericObject" || name === "JSONSchema") {
      lines.push(`export type ${name} = AnyObject`, "");
      continue;
    }
    if (schema.enum?.length) {
      lines.push(`export type ${name} = ${schema.enum.map((item) => JSON.stringify(item)).join(" | ")}`, "");
      continue;
    }
    if (schema.type === "object" || schema.properties) {
      lines.push(`export interface ${name} ${tsObjectBody(schema, schema.required ?? [], schemaMap)}`, "");
      continue;
    }
    lines.push(`export type ${name} = ${tsType(schema, schemaMap)}`, "");
  }

  return `${lines.join("\n").trimEnd()}\n`;
}

function tsObjectBody(schema, required, schemaMap) {
  if (!schema.properties || Object.keys(schema.properties).length === 0) {
    return "Record<string, unknown>";
  }
  const requiredSet = new Set(required);
  const lines = ["{"];
  for (const [propName, propSchema] of Object.entries(schema.properties)) {
    const optional = requiredSet.has(propName) ? "" : "?";
    lines.push(`  ${quoteTSProperty(propName)}${optional}: ${tsType(propSchema, schemaMap)}`);
  }
  if (schema.additionalProperties === true) {
    lines.push("  [key: string]: unknown");
  }
  lines.push("}");
  return lines.join("\n");
}

function tsType(schema, schemaMap) {
  if (!schema || Object.keys(schema).length === 0) {
    return "unknown";
  }
  if (schema.$ref) {
    const name = refName(schema.$ref);
    return name === "AnyValue" ? "unknown" : name;
  }
  if (schema.enum?.length) {
    return schema.enum.map((item) => JSON.stringify(item)).join(" | ");
  }
  if (Array.isArray(schema.type)) {
    return schema.type.filter((item) => item !== "null").map((item) => tsType({ ...schema, type: item }, schemaMap)).join(" | ");
  }
  if (schema.anyOf || schema.oneOf) {
    return (schema.anyOf ?? schema.oneOf)
      .filter((item) => item.type !== "null")
      .map((item) => tsType(item, schemaMap))
      .join(" | ") || "unknown";
  }
  switch (schema.type) {
    case "array":
      return `${tsType(schema.items ?? {}, schemaMap)}[]`;
    case "boolean":
      return "boolean";
    case "integer":
    case "number":
      return "number";
    case "string":
      return "string";
    case "object":
      if (schema.properties) {
        return tsObjectBody(schema, schema.required ?? [], schemaMap);
      }
      if (schema.additionalProperties && schema.additionalProperties !== true) {
        return `Record<string, ${tsType(schema.additionalProperties, schemaMap)}>`;
      }
      return "AnyObject";
    default:
      return "unknown";
  }
}

function generateGoTypes(names, schemaMap) {
  const body = [];
  let usesTime = false;

  for (const name of names) {
    const schema = schemaMap[name] ?? {};
    if (name === "AnyValue") {
      body.push("type AnyValue = any", "");
      continue;
    }
    if (name === "GenericObject" || name === "JSONSchema") {
      body.push(`type ${name} map[string]any`, "");
      continue;
    }
    if (schema.enum?.length) {
      body.push(`type ${name} string`, "");
      continue;
    }
    if (schema.type === "object" || schema.properties) {
      const rendered = goObject(name, schema, schema.required ?? [], schemaMap, false);
      usesTime ||= rendered.usesTime;
      body.push(rendered.text, "");
      continue;
    }
    const renderedType = goType(schema, schemaMap, true);
    usesTime ||= renderedType.usesTime;
    body.push(`type ${name} ${renderedType.text}`, "");
  }

  const imports = usesTime ? "\nimport \"time\"\n" : "";
  return [
    "// Code generated from OpenSoha contracts. DO NOT EDIT.",
    "package sohaapi",
    imports,
    body.join("\n").trimEnd(),
    "",
  ].join("\n");
}

function goObject(name, schema, required, schemaMap, inline) {
  if (!schema.properties || Object.keys(schema.properties).length === 0) {
    return { text: inline ? "map[string]any" : `type ${name} map[string]any`, usesTime: false };
  }
  const requiredSet = new Set(required);
  const lines = inline ? ["struct {"] : [`type ${name} struct {`];
  let usesTime = false;
  for (const [propName, propSchema] of Object.entries(schema.properties)) {
    const optional = !requiredSet.has(propName);
    const rendered = goType(propSchema, schemaMap, optional);
    usesTime ||= rendered.usesTime;
    const tag = optional ? `${propName},omitempty` : propName;
    lines.push(`\t${goFieldName(propName)} ${rendered.text} \`json:"${tag}"\``);
  }
  lines.push("}");
  return { text: lines.join("\n"), usesTime };
}

function goType(schema, schemaMap, optional) {
  if (!schema || Object.keys(schema).length === 0) {
    return { text: "any", usesTime: false };
  }
  const forcePointer = optional && schema["x-go-pointer"] === true;
  if (schema.$ref) {
    const name = refName(schema.$ref);
    if (name === "AnyValue") {
      return { text: "any", usesTime: false };
    }
    if (name === "RiskLevel") {
      return { text: "string", usesTime: false };
    }
    const refSchema = schemaMap[name];
    const shouldPointer = optional && refSchema?.type === "object" && !["GenericObject", "JSONSchema"].includes(name);
    return { text: shouldPointer ? `*${name}` : name, usesTime: false };
  }
  if (Array.isArray(schema.type)) {
    const nonNull = schema.type.find((item) => item !== "null");
    return goType({ ...schema, type: nonNull }, schemaMap, optional);
  }
  if (schema.anyOf || schema.oneOf) {
    const first = (schema.anyOf ?? schema.oneOf).find((item) => item.type !== "null") ?? {};
    return goType(first, schemaMap, optional);
  }
  switch (schema.type) {
    case "array": {
      const item = goType(schema.items ?? {}, schemaMap, false);
      return { text: `[]${item.text}`, usesTime: item.usesTime };
    }
    case "boolean":
      return { text: forcePointer ? "*bool" : "bool", usesTime: false };
    case "integer":
      return { text: schema.format === "int64" ? "int64" : "int", usesTime: false };
    case "number":
      return { text: "float64", usesTime: false };
    case "string":
      if (schema.format === "date-time") {
        return { text: optional ? "*time.Time" : "time.Time", usesTime: true };
      }
      return { text: "string", usesTime: false };
    case "object":
      if (schema.properties) {
        const rendered = goObject("", schema, schema.required ?? [], schemaMap, true);
        return { text: optional ? `*${rendered.text}` : rendered.text, usesTime: rendered.usesTime };
      }
      if (schema.additionalProperties && schema.additionalProperties !== true) {
        const value = goType(schema.additionalProperties, schemaMap, false);
        return { text: `map[string]${value.text}`, usesTime: value.usesTime };
      }
      return { text: "map[string]any", usesTime: false };
    default:
      return { text: "any", usesTime: false };
  }
}

function refName(ref) {
  return ref.split("/").at(-1);
}

function quoteTSProperty(name) {
  return /^[A-Za-z_$][A-Za-z0-9_$]*$/.test(name) ? name : JSON.stringify(name);
}

function goFieldName(name) {
  const words = name
    .replace(/([a-z0-9])([A-Z])/g, "$1 $2")
    .replace(/[_-]+/g, " ")
    .split(/\s+/)
    .filter(Boolean);
  return words.map(goTitleWord).join("");
}

function goTitleWord(word) {
  const lower = word.toLowerCase();
  const acronyms = new Map([
    ["ai", "AI"],
    ["api", "API"],
    ["id", "ID"],
    ["ids", "IDs"],
    ["ip", "IP"],
    ["jwt", "JWT"],
    ["mime", "MIME"],
    ["mcp", "MCP"],
    ["oidc", "OIDC"],
    ["rbac", "RBAC"],
    ["sso", "SSO"],
    ["ttl", "TTL"],
    ["uri", "URI"],
    ["url", "URL"],
  ]);
  if (acronyms.has(lower)) {
    return acronyms.get(lower);
  }
  return lower.slice(0, 1).toUpperCase() + lower.slice(1);
}
