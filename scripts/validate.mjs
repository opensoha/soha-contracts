import { readFile } from "node:fs/promises";
import { parse } from "yaml";

const root = new URL("..", import.meta.url);

const requiredJsonSchemas = [
  "agent/agent-protocol.schema.json",
  "mcp/gateway-manifest.schema.json",
  "skills/skill-manifest.schema.json",
  "plugins/plugin-manifest.schema.json",
  "events/event-envelope.schema.json",
  "auth/token-claims.schema.json",
];

const requiredSdkEntrypoints = [
  "gen/go/sohaapi/types.go",
  "gen/ts/sohaapi/index.ts",
];

for (const file of requiredJsonSchemas) {
  const text = await readFile(new URL(file, root), "utf8");
  const parsed = JSON.parse(text);
  if (!parsed.$schema || !parsed.$id) {
    throw new Error(`${file} must declare $schema and $id`);
  }
}

for (const file of requiredSdkEntrypoints) {
  const text = await readFile(new URL(file, root), "utf8");
  if (!text.includes("Code generated from OpenSoha contracts")) {
    throw new Error(`${file} must be a generated SDK entrypoint`);
  }
}

const openapiText = await readFile(new URL("openapi/soha-api.yaml", root), "utf8");
const openapi = parse(openapiText);
if (openapi?.openapi !== "3.1.0") {
  throw new Error("openapi/soha-api.yaml must use OpenAPI 3.1.0");
}
if (!openapi.info?.title || !openapi.info?.version) {
  throw new Error("openapi/soha-api.yaml must declare info.title and info.version");
}
if (!openapi.paths || Object.keys(openapi.paths).length === 0) {
  throw new Error("openapi/soha-api.yaml must declare paths");
}
if (!openapi.components?.schemas || Object.keys(openapi.components.schemas).length === 0) {
  throw new Error("openapi/soha-api.yaml must declare components.schemas");
}

const operationIds = [];
for (const pathItem of Object.values(openapi.paths)) {
  for (const operation of Object.values(pathItem ?? {})) {
    if (operation && typeof operation === "object" && operation.operationId) {
      operationIds.push(operation.operationId);
    }
  }
}
if (operationIds.length === 0) {
  throw new Error("openapi/soha-api.yaml must declare operationId values");
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
    if (!["get", "post", "put", "patch", "delete"].includes(method)) {
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

console.log(`validated ${requiredJsonSchemas.length} JSON schemas, ${requiredSdkEntrypoints.length} SDK entrypoints, ${Object.keys(openapi.components.schemas).length} schemas, and ${operationIds.length} OpenAPI operations`);
