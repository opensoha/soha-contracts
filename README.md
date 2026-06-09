# soha-contracts

Shared public contracts for OpenSoha.

This repository is the source of truth for schemas consumed across Soha core,
Soha CLI, Soha Agent, Soha Web, official skills, connectors, and Soha
Cloud services. Product repositories should depend on generated SDKs, HTTP
APIs, JSON Schemas, or released artifacts from this repository instead of
copying business logic between repos.

## Layout

```text
openapi/soha-api.yaml                 Public Soha HTTP API surface
agent/agent-protocol.schema.json      Runner and agent control-plane protocol
mcp/gateway-manifest.schema.json      AI Gateway / MCP capability manifest
skills/skill-manifest.schema.json     Official skill front matter contract
plugins/plugin-manifest.schema.json   Plugin and marketplace package manifest
events/event-envelope.schema.json     Shared event envelope contract
auth/token-claims.schema.json         Token/JWT claims contract
gen/go/sohaapi                        Go SDK types for stable cross-repo contracts
gen/ts/sohaapi                        TypeScript SDK types for stable Web contracts
```

## Boundary Rules

- `soha` remains the open-source core and must run standalone.
- Cloud-only tenant lifecycle, billing, quota, SaaS IAM, and operations
  contracts do not belong here unless they are explicitly public extension
  points.
- Open-source Soha may keep tenant-aware fields, but the default distribution
  uses `tenantId = default` and `workspaceId = default`.
- Public behavior changes should update this repository first, then update
  generated SDKs and consumers.

## SDK Consumers

Current SDK entry points:

- Go: `github.com/opensoha/soha-contracts/gen/go/sohaapi`
- TypeScript: `@opensoha/contracts/gen/ts/sohaapi`

The current SDK pass covers auth, AI Gateway capability/tool/result types,
runner-facing protocol DTOs, plugin marketplace/install DTOs, and AI Gateway
token management, audit, approval, and governance DTOs used by `soha-cli`,
`soha-agent`, and `soha-web`.

SDK DTOs are generated from `openapi/soha-api.yaml` with
`openapi-typescript` and `oapi-codegen`:

```sh
npm run generate
npm run check:generated
```

The runner-facing Go client in `gen/go/sohaapi/client.go` intentionally keeps
the stable hand-written method surface while DTOs are generated from OpenAPI.
The Go generator currently runs with Go 1.24.x in CI because the pinned
`oapi-codegen` release requires that toolchain; generated DTO consumers can
continue compiling with their own repository Go versions.

## Validation

Run:

```sh
npm test
```

The validation gate checks generated SDK freshness, JSON schema syntax,
OpenAPI structure and operation ID uniqueness, TypeScript type declarations,
and Go compilation for the generated SDK package.

## Release Shape

The stable TypeScript package name is `@opensoha/contracts`, and the stable Go
module path is `github.com/opensoha/soha-contracts/gen/go/sohaapi`.

During private staging, consumers intentionally use sibling repository
checkouts, `go.work` / `replace` entries, and `file:../soha-contracts` package
dependencies instead of resolving this repository through anonymous GitHub
module fetches or npm. After the repository is publicly readable and the npm
package is published, released consumer modules can switch to a real
`soha-contracts` version tag and the published npm package without changing SDK
import paths.

## License

This repository is licensed under the Apache License 2.0. See
[LICENSE](./LICENSE) for the full license text.
