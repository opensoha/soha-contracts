# Compatibility and Versioning

This repository owns the public OpenSoha contract surface. Consumers should
pin released versions and treat this repository as the first place where public
API, schema, and SDK behavior changes are proposed.

## Version Streams

Current stream: `0.1.x`.

- `0.x` releases may still evolve quickly, but compatibility gates must run for
  every change.
- `0.1.x` patch releases may add optional fields, optional query parameters,
  new enum values, new response status codes, new operations, and new schemas.
- `0.1.x` patch releases must not remove existing operations, rename
  `operationId` values, remove response status codes, remove schema properties,
  add required request properties, remove enum values, or narrow existing
  schema types.
- Minor `0.x` releases may include intentional breaking changes only when the
  migration is documented in the release notes and consumers have a clear
  adoption path.
- `1.0.0` starts the stable contract line. After `1.0.0`, breaking changes
  require a major version bump, a migration note, and consumer coordination.

## Version Synchronization

Each release must keep these values aligned:

- `package.json` `version`
- `package-lock.json` root package version
- `openapi/soha-api.yaml` `info.version`
- `gen/go/sohaapi/client.go` default user agent version
- Git tag `v<version>`
- Go module tag for consumers importing
  `github.com/opensoha/soha-contracts/gen/go/sohaapi`

The version sync gate runs with:

```sh
npm run check:version-sync
```

The gate fails when any of the synchronized versions diverge. In GitHub tag
release workflows it also verifies that the pushed tag is exactly
`v<package.json version>`.

Release order for `0.1.x`:

1. Update `package.json`, `package-lock.json`, and OpenAPI `info.version` to
   the same patch version.
2. Regenerate SDKs when OpenAPI changed.
3. Run `npm test`, `npm pack --dry-run`, `npm run release:artifacts`, and
   the consumer matrix.
4. Create tag `v<version>`.
5. Let the release workflow publish npm when `NPM_TOKEN` is configured and
   create the GitHub release with the npm tarball, `.sha256` checksum, and
   release manifest attached.

## Compatibility Baseline

`compat/openapi-0.1.0.snapshot.json` is the current OpenAPI compatibility
baseline. It is generated from `openapi/soha-api.yaml` by:

```sh
npm run check:openapi-breaking:update
```

Do not update the baseline just to make a failing check pass. Update it only
when starting a new compatibility baseline after an intentional versioned break
or when establishing a baseline for a newly stabilized stream.

The breaking-change check currently blocks these high-risk changes:

- Removing a path or HTTP operation.
- Changing an existing `operationId`.
- Removing an existing response status or media type.
- Adding a required request body or required parameter.
- Removing a component schema.
- Removing an existing schema property.
- Adding a required schema property.
- Removing an enum value.
- Changing an existing schema type, format, `$ref`, array item schema, or
  object `additionalProperties` contract in a narrowing direction.

This is a baseline guard, not a full API review substitute. Contract changes
that pass this script can still require consumer tests when semantics change.

## OpenAPI Loose Boundary Baseline

`compat/openapi-loose-boundaries-0.1.0.json` records every current OpenAPI use
of `AnyValue`, `GenericObject`, and bare `additionalProperties: true`. The
check runs with:

```sh
npm run check:openapi-loose-boundaries
```

New loose boundaries must not be added accidentally. If a future contract truly
requires dynamic JSON, add it to the baseline with an intentional review note.
If a loose boundary is tightened, remove it from the baseline in the same
change so the contract surface keeps moving toward named schemas.

## Public Cloud Extension Points

`CloudExtensionPoint` OpenAPI DTOs describe only public extension points that
Cloud can consume through shared contracts, HTTP APIs, SDKs, webhooks, or event
sinks. Billing, quota, and SaaS IAM remain Cloud-only domains and must not be
promoted into this public contracts repository.

For `0.1.x`, patch releases may add optional public extension point fields,
new extension point kinds, new public boundary kinds, and new supported
consumer names. They must not remove or rename `id`, `name`, `kind`,
`boundary`, `contractRefs`, or `excludedCloudOnlyDomains`; must not add new
required fields; and must not remove `billing`, `quota`, or `saas-iam` from
the Cloud-only exclusion enum.

## Fixture Policy

Every JSON Schema owned by this repository must have at least one valid and one
invalid fixture under `fixtures/json-schema/<contract>/`.

High-value OpenAPI DTOs must have valid and invalid fixtures under
`fixtures/openapi/<SchemaName>/`. Add fixtures when a schema becomes important
to a cross-repository workflow, especially auth, runner callbacks, AI Gateway,
MCP, plugins, tokens, approvals, and governance.

Fixture validation runs with:

```sh
npm run validate
```

Request/response and manifest examples under `examples/` are also validated
by `npm run validate`. Examples are release documentation assets, so they
should be valid positive payloads rather than negative test cases.

## Ecosystem JSON Schema Baseline

`compat/json-schema-ecosystem-0.1.0.json` protects the public plugin, skill,
and connector JSON Schemas from accidental breaking changes. The check runs
with:

```sh
npm run check:json-schema-compat
```

The baseline blocks removing existing root properties, adding new root required
properties, and removing existing enum values. Optional fields and enum values
may be added in `0.1.x` patch releases when fixtures and compatibility notes
are updated.

## Cluster Capability Matrix

`capabilities/cluster-capability-matrix.schema.json` owns the public
Direct/Agent runtime capability matrix consumed by Core, Web, CLI, Agent,
Skills, Connectors, and Cloud extension points. It describes Kubernetes, Helm,
YAML, logs, exec, port-forward, Docker, and runtime support with explicit
available, partial, and unsupported states.

For `0.1.x`, patch releases may add optional fields, new capability entries,
new categories, new risk levels, and new support diagnostics. They must not
remove or rename `items`, entry `key`, `label`, `category`, `riskLevel`,
`requiresApproval`, `direct`, or `agent`; must not add new required fields; and
must not narrow the existing `available`, `partial`, or `unsupported` status
values. Any `partial` or `unsupported` support state must continue to provide a
human-readable reason so downstream UI and diagnostics can explain degraded
runtime behavior.

## Plugin Manifest

`plugins/plugin-manifest.schema.json` owns the public package manifest for
skills, skill packs, MCP presets, connectors, AI provider adapters, agent
profiles, and gateway policy packs. Marketplace, CLI, Web, Skills, Connectors,
and Cloud consumers should use this schema as the source of truth for package
identity, compatibility declarations, asset references, permissions, secrets,
and integrity metadata.

For `0.1.x`, patch releases may add optional plugin types, optional asset
groups, optional capability groups, optional permission fields, optional secret
metadata, and optional integrity metadata. They must not remove or rename `id`,
`name`, `version`, `publisher`, or `type`; must not add new required fields;
must not narrow the plugin ID pattern; must not remove existing plugin `type`
enum values; and must not allow unknown top-level fields outside explicitly
versioned metadata.

## Skill Manifest

`skills/skill-manifest.schema.json` owns the public structured front matter
contract for official and ecosystem skill assets. Skill repositories may add
repository-local metadata, but cross-repository routing, permission review, and
capability binding should use this schema first.

For `0.1.x`, patch releases may add optional fields, optional provider skill
reference metadata, optional capability reference conventions, and optional
permission or scope annotations. They must not remove or rename `id`, `name`,
or `category`; must not add new required fields; must not narrow the skill ID
pattern; and must not allow unknown top-level fields outside `metadata`.

## Agent Profile

`profiles/agent-profile.schema.json` owns the public agent runtime profile
manifest used by official skills, CLI installers, Web previews, Agent runtime
selection, and Cloud managed fleet configuration. It binds skills, MCP presets,
capabilities, scopes, enabled/disabled tools, execution budgets, and guardrails.

For `0.1.x`, patch releases may add optional modes, optional budget fields,
optional guardrail conventions, and optional metadata. They must not remove or
rename `schemaVersion`, `id`, `name`, `version`, `mode`, `providerKind`,
`skillRefs`, `mcpPresetRefs`, `capabilityRefs`, `platformCapabilityRefs`,
`requiredScopes`, `enabledToolRefs`, or `guardrails`; must not add new required
fields; must not narrow the ID pattern or semantic version pattern; and must
not allow unknown top-level fields outside `metadata`.

## MCP Preset

`presets/mcp-preset.schema.json` owns the public MCP preset manifest used by
official skills, CLI installers, Gateway capability binding, Web previews, and
Cloud/local preset packaging. It declares a bounded set of tools, resources,
prompts, scopes, capabilities, risk level, and guardrails.

For `0.1.x`, patch releases may add optional risk levels, optional capability
metadata, optional top-level metadata, and optional tool/resource/prompt fields.
They must not remove or rename `schemaVersion`, `id`, `name`, `version`,
`riskLevel`, `skillRefs`, `platformCapabilityRefs`, `requiredScopes`, `tools`,
or `guardrails`; must not add new required fields; must not narrow the ID
pattern or semantic version pattern; and must not allow unknown top-level
fields outside `metadata`.

## Governance Approval Request

`governance/approval-request.schema.json` owns the public approval request,
trace, decision, stage, and timeline event contract for high-risk Gateway and
platform operations. OpenAPI approval DTOs should stay compatible with this
schema when the same fields are exposed over HTTP.

For `0.1.x`, patch releases may add optional approval fields, optional trace
fields, optional decision metadata, optional stage metadata, and optional
timeline event metadata. They must not remove or rename `id`, `status`,
`strategy`, `actorType`, `actorId`, `toolName`, `riskLevel`,
`requiresApproval`, `summary`, `createdAt`, or `updatedAt`; must not add new
required fields; must not remove existing risk level values; and must not allow
unknown top-level fields outside explicitly modeled metadata objects.

## Governance Audit Log

`governance/audit-log.schema.json` owns the public audit log contract for
Gateway, platform, connector, plugin, skill, and Cloud extension point actions.
OpenAPI audit DTOs and event pipeline audit payloads should stay compatible
with this schema when the same fields are exposed outside a repository.

For `0.1.x`, patch releases may add optional audit fields, optional actor
display metadata, optional resource scope fields, optional request correlation
fields, and optional metadata conventions. They must not remove or rename `id`,
`actorType`, `actorId`, `action`, `result`, `summary`, or `createdAt`; must not
add new required fields; must not remove existing risk level values; and must
not allow unknown top-level fields outside `metadata` or `resourceScope`.

## Connector Event Envelope

`connectors/connector-event-envelope.schema.json` owns the public connector
runtime event sink batch envelope. Connector repositories may mirror this file
for package compatibility, but the canonical `$id` remains
`https://contracts.opensoha.dev/connectors/connector-event-envelope.schema.json`.

For `0.1.x`, patch releases may add optional fields to the wrapper or event
items and optional provider payload conventions, but must not remove or rename
top-level `connectorId` or `events`; must not remove or rename event `id`,
`type`, `source`, `occurredAt`, `subject`, or `payload`; must not add new
required fields, allow unknown wrapper/event fields, narrow the connector ID
pattern, or move provider-specific data out of `payload`.

## Runner Operation Lifecycle

`runner/operation-lifecycle.schema.json` owns the public event contract for
long-running operation lifecycle breadcrumbs across delivery, Docker, VM, AI
provider, agent run, and connector action runners.

For `0.1.x`, patch releases may add optional operation kinds, optional states,
optional evidence kinds, and optional metadata fields. They must not remove or
rename `operationId`, `operationKind`, `runnerId`, `state`, `occurredAt`,
`evidence`, `finalState`, or `callback`; must not add new required fields; and
must not narrow existing lifecycle state names or final status values.

## Consumer Adoption

Recommended adoption flow:

1. Change contracts and fixtures here first.
2. Run `npm test` in this repository.
3. Regenerate SDKs with `npm run generate` when OpenAPI changes.
4. Update consumers to the same contract version.
5. Run consumer build/test gates before release.

Private staging may use sibling checkouts, `go.work`, Go `replace`, or
`file:../soha-contracts`. Public release builds should pin tagged versions or
published npm packages.

The consumer matrix gate is:

```sh
npm run check:consumers
```

Local runs skip missing sibling checkouts by default. CI and release candidate
validation should use `--require-all` or a specific `--consumer <name>` so a
missing checkout is a failure.
