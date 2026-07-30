---
name: soha-contracts
description: >-
  Implement or review public OpenSoha contracts in OpenAPI, JSON Schema,
  generated Go and TypeScript SDKs, compatibility baselines, examples,
  fixtures, and release artifacts. Use when shared API behavior, agent or
  runner protocols, MCP, skills, plugins, connectors, knowledge, evaluation,
  memory, auth, or cross-repository DTOs change.
---

# Soha Contracts

## Purpose

Keep `soha-contracts` the versioned source of truth for behavior shared by
Soha repositories. Change contracts before consumers and preserve stable SDK
import paths.

## Workflow

1. Read `README.md`, `COMPATIBILITY.md`, the affected schema, its fixtures,
   and the generator or compatibility check before editing.
2. Decide whether the source is `openapi/soha-api.yaml`, a focused JSON
   Schema, or a small hand-maintained Go helper package. Do not model private
   implementation details as public contracts.
3. Add or update valid and invalid fixtures for boundary behavior. Keep examples
   representative and free of secrets.
4. Run `npm run generate` when OpenAPI DTOs change. Review generated diffs;
   never patch generated DTOs to bypass the source schema.
5. Update consumers one repository at a time after the contract gate passes.
6. Run the focused compatibility checks, then `npm test`. Run the consumer
   matrix for every affected consumer.

## Ownership Rules

- `gen/ts/sohaapi/index.ts` and generated Go DTOs come from OpenAPI.
- `gen/go/sohaapi/client.go` is the deliberate hand-maintained exception:
  keep its supported method surface stable and cover additions with client
  tests.
- Keep optional additions backward compatible during `0.1.x`. Do not remove
  operations, statuses, properties, or enum values, and do not add required
  request fields without an explicit version decision.
- Update `compat/**` only for an intentional reviewed baseline change, never
  to silence a failure.
- Cloud-only tenancy, billing, quota, or SaaS operations stay in
  `soha-cloud` unless they are genuine public extension points.
- Consumers use released artifacts, generated SDKs, HTTP, or schemas. They must
  not copy contract definitions or import sibling internals.

## Verification

```bash
npm test
GOWORK=off go vet ./...
GOWORK=off go run golang.org/x/vuln/cmd/govulncheck@v1.3.0 ./...
npm run check:consumers -- --consumer soha --require-all
npm run check:consumers -- --consumer soha-web --require-all
npm run check:consumers -- --consumer soha-cli --require-all
npm run check:consumers -- --consumer soha-agent --require-all
git diff --check
```

Use Go `1.26.5`. `npm test` includes generation, schema, fixture, compatibility, package, and Go checks; CI additionally runs `golangci-lint v2.9.0` with only-new-issues semantics. Every contract change requires all four consumer checks, not only the directly edited consumer. Run release artifact checks when package or release behavior changes.
