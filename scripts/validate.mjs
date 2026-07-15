import { readdir, readFile } from "node:fs/promises";
import { join } from "node:path";
import { fileURLToPath } from "node:url";
import Ajv2020 from "ajv/dist/2020.js";
import addFormats from "ajv-formats";
import { parse } from "yaml";

const root = new URL("..", import.meta.url);
const rootPath = fileURLToPath(root);
const openapiPath = "openapi/soha-api.yaml";

const requiredJsonSchemas = [
  {
    file: "capabilities/cluster-capability-matrix.schema.json",
    fixtureDir: "fixtures/json-schema/cluster-capability-matrix",
  },
  { file: "profiles/agent-profile.schema.json", fixtureDir: "fixtures/json-schema/agent-profile" },
  { file: "presets/mcp-preset.schema.json", fixtureDir: "fixtures/json-schema/mcp-preset" },
  { file: "governance/approval-request.schema.json", fixtureDir: "fixtures/json-schema/approval-request" },
  { file: "governance/audit-log.schema.json", fixtureDir: "fixtures/json-schema/audit-log" },
  { file: "agent/agent-protocol.schema.json", fixtureDir: "fixtures/json-schema/agent-protocol" },
  {
    file: "runner/operation-lifecycle.schema.json",
    fixtureDir: "fixtures/json-schema/runner-operation-lifecycle",
  },
  { file: "mcp/gateway-manifest.schema.json", fixtureDir: "fixtures/json-schema/gateway-manifest" },
  {
    file: "mcp/gateway-capability-baseline.schema.json",
    fixtureDir: "fixtures/json-schema/gateway-capability-baseline",
  },
  { file: "skills/skill-manifest.schema.json", fixtureDir: "fixtures/json-schema/skill-manifest" },
  { file: "plugins/plugin-manifest.schema.json", fixtureDir: "fixtures/json-schema/plugin-manifest" },
  {
    file: "plugins/marketplace-catalog.schema.json",
    fixtureDir: "fixtures/json-schema/marketplace-catalog",
  },
  { file: "events/event-envelope.schema.json", fixtureDir: "fixtures/json-schema/event-envelope" },
  { file: "auth/token-claims.schema.json", fixtureDir: "fixtures/json-schema/token-claims" },
  {
    file: "connectors/connector-event-envelope.schema.json",
    fixtureDir: "fixtures/json-schema/connector-event-envelope",
  },
];

const requiredOpenapiFixtureSchemas = [
  "PasswordLoginRequest",
  "AuthResultEnvelope",
  "ExecutionCallbackRequest",
  "DockerOperationCallbackRequest",
  "AgentRunCallbackRequest",
  "AIGatewayManifestEnvelope",
  "PluginManifest",
  "PersonalAccessTokenInput",
  "ServiceAccountInput",
  "GovernanceStatusEnvelope",
  "MCPCapabilityListEnvelope",
  "CloudExtensionPointListEnvelope",
  "WorkbenchStreamEvent",
  "ComputeOverviewEnvelope",
  "ComputeAccessSourceListEnvelope",
  "ComputeProviderListEnvelope",
  "ComputeResourceRelationListEnvelope",
  "ComputeTaskListEnvelope",
  "ErrorEnvelope",
];

const requiredJsonExamples = [
  {
    file: "examples/cluster-capability-matrix/runtime-parity.json",
    schemaFile: "capabilities/cluster-capability-matrix.schema.json",
  },
  { file: "examples/agent-profile/k8s-sre-readonly.json", schemaFile: "profiles/agent-profile.schema.json" },
  { file: "examples/mcp-preset/k8s-readonly.json", schemaFile: "presets/mcp-preset.schema.json" },
  { file: "examples/governance/approval-request.json", schemaFile: "governance/approval-request.schema.json" },
  { file: "examples/governance/audit-log.json", schemaFile: "governance/audit-log.schema.json" },
  { file: "examples/gateway-manifest/minimal.json", schemaFile: "mcp/gateway-manifest.schema.json" },
  {
    file: "examples/gateway-capability-baseline/minimal.json",
    schemaFile: "mcp/gateway-capability-baseline.schema.json",
  },
  { file: "examples/token-claims/service-account.json", schemaFile: "auth/token-claims.schema.json" },
  { file: "examples/plugin-manifest/skill-pack.json", schemaFile: "plugins/plugin-manifest.schema.json" },
  { file: "examples/plugin-manifest/connector.json", schemaFile: "plugins/plugin-manifest.schema.json" },
  { file: "examples/plugin-manifest/compute-provider.json", schemaFile: "plugins/plugin-manifest.schema.json" },
  { file: "examples/marketplace-catalog/default.json", schemaFile: "plugins/marketplace-catalog.schema.json" },
  { file: "examples/skill-manifest/k8s-sre.json", schemaFile: "skills/skill-manifest.schema.json" },
  {
    file: "examples/connector-event-envelope/feishu-message.json",
    schemaFile: "connectors/connector-event-envelope.schema.json",
  },
  {
    file: "examples/runner-operation-lifecycle/late-callback.json",
    schemaFile: "runner/operation-lifecycle.schema.json",
  },
];

const requiredOpenapiExamples = [
  { file: "examples/openapi/gateway-manifest-envelope.json", schemaName: "AIGatewayManifestEnvelope" },
  { file: "examples/openapi/tool-invocation-request.json", schemaName: "ToolInvocationRequest" },
  { file: "examples/openapi/tool-invocation-result-envelope.json", schemaName: "ToolInvocationResultEnvelope" },
  { file: "examples/openapi/agent-run-tool-call-request.json", schemaName: "AgentRunToolCallRequest" },
  { file: "examples/openapi/agent-tool-call-result-envelope.json", schemaName: "AgentToolCallResultEnvelope" },
  { file: "examples/openapi/cloud-extension-point-list-envelope.json", schemaName: "CloudExtensionPointListEnvelope" },
];

const requiredSdkEntrypoints = ["gen/go/sohaapi/types.go", "gen/ts/sohaapi/index.ts"];

const ajv = new Ajv2020({
  allErrors: true,
  strict: false,
  validateFormats: true,
});
addFormats(ajv);
ajv.addFormat("int64", true);

const jsonSchemaValidators = new Map();
let jsonSchemaFixtureCount = 0;
for (const contract of requiredJsonSchemas) {
  const text = await readFile(new URL(contract.file, root), "utf8");
  const parsed = JSON.parse(text);
  if (!parsed.$schema || !parsed.$id) {
    throw new Error(`${contract.file} must declare $schema and $id`);
  }

  const validate = ajv.compile(parsed);
  jsonSchemaValidators.set(contract.file, validate);
  jsonSchemaFixtureCount += await validateFixtureDirectory(contract.fixtureDir, validate);
}

for (const file of requiredSdkEntrypoints) {
  const text = await readFile(new URL(file, root), "utf8");
  if (!text.includes("Code generated from OpenSoha contracts")) {
    throw new Error(`${file} must be a generated SDK entrypoint`);
  }
}

const openapiText = await readFile(new URL(openapiPath, root), "utf8");
const openapi = parse(openapiText);
validateOpenapiStructure(openapi);
const openapiFixtureCount = await validateOpenapiFixtures(openapi);
const jsonExampleCount = await validateJsonExamples();
const openapiExampleCount = await validateOpenapiExamples(openapi);

console.log(
  `validated ${requiredJsonSchemas.length} JSON schemas, ${jsonSchemaFixtureCount} JSON schema fixtures, ${jsonExampleCount} JSON examples, ${requiredSdkEntrypoints.length} SDK entrypoints, ${Object.keys(openapi.components.schemas).length} OpenAPI schemas, ${openapiFixtureCount} OpenAPI fixtures, ${openapiExampleCount} OpenAPI examples, and ${collectOperationIds(openapi).length} OpenAPI operations`,
);

async function validateOpenapiFixtures(openapi) {
  const openapiAjv = buildOpenapiAjv(openapi);
  let fixtureCount = 0;
  for (const schemaName of requiredOpenapiFixtureSchemas) {
    const validate = getOpenapiSchemaValidator(openapiAjv, schemaName);
    fixtureCount += await validateFixtureDirectory(`fixtures/openapi/${schemaName}`, validate);
  }
  return fixtureCount;
}

async function validateJsonExamples() {
  for (const example of requiredJsonExamples) {
    const validate = jsonSchemaValidators.get(example.schemaFile);
    if (!validate) {
      throw new Error(`${example.file} references unregistered JSON schema ${example.schemaFile}`);
    }
    await validateExampleFile(example.file, validate);
  }
  return requiredJsonExamples.length;
}

async function validateOpenapiExamples(openapi) {
  const openapiAjv = buildOpenapiAjv(openapi);
  for (const example of requiredOpenapiExamples) {
    const validate = getOpenapiSchemaValidator(openapiAjv, example.schemaName);
    await validateExampleFile(example.file, validate);
  }
  return requiredOpenapiExamples.length;
}

function buildOpenapiAjv(openapi) {
  const documentId = "https://contracts.opensoha.dev/openapi/soha-api.yaml";
  const openapiSchemaDocument = {
    $id: documentId,
    $schema: "https://json-schema.org/draft/2020-12/schema",
    components: {
      schemas: openapi.components.schemas,
    },
  };
  const openapiAjv = new Ajv2020({
    allErrors: true,
    strict: false,
    validateFormats: true,
  });
  addFormats(openapiAjv);
  openapiAjv.addFormat("int64", true);
  openapiAjv.addSchema(openapiSchemaDocument);
  openapiAjv.contractDocumentId = documentId;
  return openapiAjv;
}

function getOpenapiSchemaValidator(openapiAjv, schemaName) {
  const validate = openapiAjv.getSchema(`${openapiAjv.contractDocumentId}#/components/schemas/${schemaName}`);
  if (!validate) {
    throw new Error(`${openapiPath} is missing fixture/example-covered schema ${schemaName}`);
  }
  return validate;
}

async function validateFixtureDirectory(fixtureDir, validate) {
  const validDir = join(rootPath, fixtureDir, "valid");
  const invalidDir = join(rootPath, fixtureDir, "invalid");
  const validCount = await validateFixtures(validDir, true, validate);
  const invalidCount = await validateFixtures(invalidDir, false, validate);
  if (validCount === 0 || invalidCount === 0) {
    throw new Error(`${fixtureDir} must contain at least one valid and one invalid fixture`);
  }
  return validCount + invalidCount;
}

async function validateFixtures(dir, shouldPass, validate) {
  const files = (await readdir(dir).catch(() => [])).filter((file) => file.endsWith(".json")).sort();
  for (const file of files) {
    const fixturePath = join(dir, file);
    const payload = JSON.parse(await readFile(fixturePath, "utf8"));
    const passed = validate(payload);
    if (shouldPass && !passed) {
      throw new Error(`${fixturePath} should be valid:\n${formatAjvErrors(validate.errors)}`);
    }
    if (!shouldPass && passed) {
      throw new Error(`${fixturePath} should be invalid`);
    }
  }
  return files.length;
}

async function validateExampleFile(file, validate) {
  const payload = JSON.parse(await readFile(new URL(file, root), "utf8"));
  const passed = validate(payload);
  if (!passed) {
    throw new Error(`${file} should be valid:\n${formatAjvErrors(validate.errors)}`);
  }
}

function validateOpenapiStructure(openapi) {
  if (openapi?.openapi !== "3.1.0") {
    throw new Error(`${openapiPath} must use OpenAPI 3.1.0`);
  }
  if (!openapi.info?.title || !openapi.info?.version) {
    throw new Error(`${openapiPath} must declare info.title and info.version`);
  }
  if (!openapi.paths || Object.keys(openapi.paths).length === 0) {
    throw new Error(`${openapiPath} must declare paths`);
  }
  if (!openapi.components?.schemas || Object.keys(openapi.components.schemas).length === 0) {
    throw new Error(`${openapiPath} must declare components.schemas`);
  }

  const operationIds = collectOperationIds(openapi);
  if (operationIds.length === 0) {
    throw new Error(`${openapiPath} must declare operationId values`);
  }
  const duplicateOperationIds = operationIds.filter((id, index) => operationIds.indexOf(id) !== index);
  if (duplicateOperationIds.length > 0) {
    throw new Error(`duplicate operationId values: ${[...new Set(duplicateOperationIds)].join(", ")}`);
  }

  const invalidColonParams = Object.keys(openapi.paths).filter((path) => /\/:[A-Za-z]/.test(path));
  if (invalidColonParams.length > 0) {
    throw new Error(`OpenAPI paths must use {param} syntax: ${invalidColonParams.join(", ")}`);
  }

  for (const [path, pathItem] of Object.entries(openapi.paths)) {
    for (const [method, operation] of Object.entries(pathItem ?? {})) {
      if (!isHttpMethod(method)) {
        continue;
      }
      if (!operation.operationId) {
        throw new Error(`${method.toUpperCase()} ${path} is missing operationId`);
      }
      if (!operation.responses || Object.keys(operation.responses).length === 0) {
        throw new Error(`${method.toUpperCase()} ${path} is missing responses`);
      }
    }
  }
}

function collectOperationIds(openapi) {
  const operationIds = [];
  for (const pathItem of Object.values(openapi.paths ?? {})) {
    for (const [method, operation] of Object.entries(pathItem ?? {})) {
      if (isHttpMethod(method) && operation?.operationId) {
        operationIds.push(operation.operationId);
      }
    }
  }
  return operationIds;
}

function isHttpMethod(method) {
  return ["get", "post", "put", "patch", "delete"].includes(method);
}

function formatAjvErrors(errors) {
  return (errors ?? [])
    .map((error) => `- ${error.instancePath || "/"} ${error.message}`)
    .join("\n");
}
