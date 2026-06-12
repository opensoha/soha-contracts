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
compat/openapi-0.1.0.snapshot.json    OpenAPI compatibility baseline
examples                              Valid request/response and manifest examples
fixtures                              Valid and invalid contract payloads
capabilities/cluster-capability-matrix.schema.json  Direct/Agent runtime capability matrix
profiles/agent-profile.schema.json    Agent runtime profile manifest
presets/mcp-preset.schema.json        MCP preset manifest
governance/approval-request.schema.json  Approval request and trace contract
governance/audit-log.schema.json      Audit log contract
agent/agent-protocol.schema.json      Runner and agent control-plane protocol
runner/operation-lifecycle.schema.json  Long-running operation lifecycle events
mcp/gateway-manifest.schema.json      AI Gateway / MCP capability manifest
skills/skill-manifest.schema.json     Official skill front matter contract
plugins/plugin-manifest.schema.json   Plugin and marketplace package manifest
events/event-envelope.schema.json     Shared event envelope contract
auth/token-claims.schema.json         Token/JWT claims contract
connectors/connector-event-envelope.schema.json  Connector runtime event batch envelope
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
runner-facing protocol DTOs, plugin marketplace/install DTOs, public Cloud
extension point DTOs, and AI Gateway token management, audit, approval, and
governance DTOs used by `soha-cli`, `soha-agent`, and `soha-web`.

SDK DTOs are generated from `openapi/soha-api.yaml` with
`openapi-typescript` and `oapi-codegen`:

```sh
npm run generate
npm run check:generated
```

`npm run check:generated` generates SDK output in a temporary directory and
compares it to committed files. It must not rewrite `gen/` during a freshness
check.

The Go client in `gen/go/sohaapi/client.go` intentionally keeps a stable
hand-written method surface while DTOs are generated from OpenAPI. For
`0.1.x`, the supported Go client method surface covers:

- system health and readiness checks;
- public auth discovery, password login, refresh, OIDC code exchange, current
  principal/profile/bootstrap, logout, and stream tickets;
- delivery, Docker, and agent-run claim/status/callback runner APIs;
- AI Gateway audit, approval, governance, token, service-account, and MCP
  capability APIs;
- plugin marketplace, install, enable/disable, upgrade, configuration, and
  manifest APIs.

Other Go consumers should use the generated DTOs directly and keep HTTP
calling code in the consumer until those APIs are promoted into the supported
client surface. Redirect-oriented flows such as OIDC provider redirects should
not use the JSON-only `doJSON` helper until a non-JSON client mode is added.

The Go generator currently runs with Go 1.24.x in CI because the pinned
`oapi-codegen` release requires that toolchain; generated DTO consumers can
continue compiling with their own repository Go versions.

## Validation

Run:

```sh
npm test
```

The validation gate checks generated SDK freshness, JSON schema syntax,
valid and invalid JSON Schema fixtures, valid and invalid OpenAPI DTO
fixtures, request/response and manifest examples, OpenAPI structure and
operation ID uniqueness, version synchronization, OpenAPI lint rules,
breaking-change compatibility against the checked-in baseline, TypeScript
type declarations, and Go compilation for the generated SDK package.

Useful focused checks:

```sh
npm run check:version-sync
npm run validate
npm run lint:openapi
npm run check:openapi-breaking
npm run check:openapi-loose-boundaries
npm run check:json-schema-compat
npm run release:artifacts
npm run release:verify -- dist/release/opensoha-contracts-0.1.0.tgz
```

When intentionally establishing a new OpenAPI compatibility baseline, update
the snapshot with:

```sh
npm run check:openapi-breaking:update
```

Do not update the baseline just to silence an unexpected breaking-change
failure.

## Consumer Matrix

The cross-repository consumer gate is:

```sh
npm run check:consumers
```

By default, the script discovers sibling checkouts under the parent directory
and skips any missing consumer. CI uses the strict form:

```sh
npm run check:consumers -- --consumer soha --require-all
npm run check:consumers -- --consumer soha-web --require-all
npm run check:consumers -- --consumer soha-cli --require-all
npm run check:consumers -- --consumer soha-agent --require-all
```

Go consumers run in a temporary `go.work` that includes this checkout, so
changes are tested against the local contracts module. `soha-web` runs
`npm ci` and `npm run build` against its `file:../soha-contracts` dependency.

Use `--dry-run` to inspect the discovered commands without executing consumer
builds.

## Compatibility Policy

See [COMPATIBILITY.md](./COMPATIBILITY.md) for the versioning, compatibility,
fixture, baseline, and consumer adoption rules. In short, `0.1.x` changes may
add optional contract surface but must not remove existing public operations,
rename operation IDs, remove existing response statuses, remove schema
properties, add required request properties, or remove enum values.

## Release Shape

The stable TypeScript package name is `@opensoha/contracts`, and the stable Go
module path is `github.com/opensoha/soha-contracts/gen/go/sohaapi`.

Release tags must be `v<package.json version>`. `npm test` runs
`npm run check:version-sync`, which verifies `package.json`,
`package-lock.json`, OpenAPI `info.version`, the Go module path, the Go client
user agent version, and the GitHub release tag when the workflow runs from a
tag.

Release artifacts are prepared and verified with:

```sh
npm run release:artifacts
npm run release:verify -- dist/release/opensoha-contracts-0.1.0.tgz
```

The command writes `dist/release/<npm-tarball>.tgz`, a matching
`.sha256` checksum file, `dist/release/release-manifest.json`, a CycloneDX SBOM
with its checksum, a provenance policy manifest with its checksum, and, when
provided by CI, `dist/release/consumer-matrix.json`. `release:verify` checks
the tarball checksum, manifest metadata, packed `package.json`, every exported
package entrypoint, the SBOM identity and checksum, the provenance policy and
checksum, and the consumer matrix summary.

The tag workflow runs `npm test`, `npm pack --dry-run`, the strict consumer
matrix, requires `NPM_TOKEN`, publishes npm with provenance, verifies the npm
package and Go module tag after publish, and attaches the tarball, checksum,
SBOM, SBOM checksum, provenance policy, provenance policy checksum, manifest,
and consumer matrix to the GitHub release. It then downloads the published
release assets and runs `release:verify` against the downloaded tarball.

The local release manifest records the provenance policy expected from the tag
workflow. It does not prove a remote npm registry attestation by itself; that
evidence is produced and checked only after the tag workflow publishes with
`npm publish --provenance`.

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
