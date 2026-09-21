---
name: soha-contracts
description: Change or review shared OpenSoha API schemas, protocol contracts, generated SDKs, compatibility, and releases. Consumer-only UI changes do not require contract edits.
---

# Soha Contracts

## Purpose

Keep `soha-contracts` the versioned source of truth for behavior shared by
Soha repositories. Change contracts before consumers and preserve stable SDK
import paths.

## Workflow

1. Read `AGENTS.md`, `README.md`, `COMPATIBILITY.md`, the affected schema, its
   fixtures, and the generator or compatibility check before implementation.
   Documentation-only work reads the relevant text and source; reuse unchanged context.
2. Decide whether the source is `openapi/soha-api.yaml`, a focused JSON
   Schema, or a small hand-maintained Go helper package. Do not model private
   implementation details as public contracts.
3. State the changed behavior and preserved invariants: permissions, scope,
   pagination, errors, side effects, idempotency and version compatibility where applicable.
   Add valid and invalid fixtures; keep examples representative and free of secrets.
4. Run `npm run generate` when OpenAPI source changes affect generated artifacts.
   Review generated diffs; never patch generated DTOs to bypass the source schema.
5. Update consumers one repository at a time after the contract gate passes.
6. For public contract changes, run focused compatibility checks, `npm test`,
   and the complete consumer matrix in `.github/workflows/ci.yml` and the
   consumer-check script. The commands below mirror that matrix; update them
   together when the executable matrix changes. Collaboration prose changes
   do not regenerate SDKs or run consumers.

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

## Reproducible Evidence

- Record the contracts SHA and the actual resolved SHA/version of every checked
  consumer. `main` is a selector, not a reproducible result identifier.
- `compat/consumer-revisions.json` describes reviewed selections; inspect the CI
  fallback for unlisted consumers. Do not describe partially pinned checks as
  a fully fixed release combination or replace reviewed pins without evidence.
- Keep fixed-version compatibility evidence separate from latest-branch
  integration evidence. A fixed consumer can pass while a newer one fails.
- Report passing, failing, skipped and unrun checks separately. A missing sibling
  is a verification gap, not a passing consumer; SDK generation or compilation
  alone does not prove the public behavior is correct.

## Verification

```bash
npm test
GOWORK=off go vet ./...
GOWORK=off go run golang.org/x/vuln/cmd/govulncheck@v1.3.0 ./...
npm run check:consumers -- --consumer soha --require-all
npm run check:consumers -- --consumer soha-web --require-all
npm run check:consumers -- --consumer soha-cli --require-all
npm run check:consumers -- --consumer soha-agent --require-all
npm run check:consumers -- --consumer soha-skills --require-all
git diff --check
```

Use Go `1.26.6`. `npm test` includes generation, schema, fixture, compatibility,
package, and Go checks; CI additionally runs `golangci-lint v2.9.0` with
only-new-issues semantics. Every public contract change requires the complete
CI consumer matrix, not only the directly edited consumer. Run release artifact
checks when package or release behavior changes. Preserve existing gates;
resolve documentation/workflow drift rather than weakening either silently.
