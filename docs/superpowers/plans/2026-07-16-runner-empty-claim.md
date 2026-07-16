# Runner Empty Claim Responses Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Treat an empty runner claim queue as an idle polling result instead of an HTTP 404 error for Delivery and Docker claims.

**Architecture:** Keep repository and application `ErrNotFound` behavior unchanged. Translate only claim-handler `ErrNotFound` results to HTTP 204, retain HTTP 202 with the existing envelope for successful claims, and document both outcomes in the shared OpenAPI contract. The existing SDK and agent already accept bodyless successful responses; regression tests will lock that behavior down.

**Tech Stack:** OpenAPI 3.1, Go 1.25, Gin, generated TypeScript contracts, handwritten Go SDK client.

---

### Task 1: Update the public runner claim contract

**Files:**
- Modify: `openapi/soha-api.yaml`
- Modify: `scripts/validate.mjs`
- Modify: `gen/go/sohaapi/client_test.go`
- Regenerate: `gen/ts/sohaapi/schema.ts`
- Modify if required by compatibility policy: `compat/openapi-0.1.0.snapshot.json`

- [ ] Add structural validation requiring `202` and `204` responses for both claim operations.
- [ ] Run `npm run validate` and confirm the new assertion fails against the old contract.
- [ ] Replace the incorrect `200` response with `202`, and add bodyless `204` for both endpoints.
- [ ] Add Go client tests proving both claim methods return zero-value items without error on `204`.
- [ ] Run `npm run generate`, focused Go tests, generated checks, OpenAPI lint, and compatibility checks.
- [ ] Commit the contracts change.

### Task 2: Map empty backend claims to HTTP 204

**Files:**
- Create: `internal/api/handlers/runner_claim_test.go`
- Modify: `internal/api/handlers/delivery.go`
- Modify: `internal/api/handlers/docker.go`

- [ ] Add handler tests covering empty queue, successful claim, and unexpected error for Delivery and Docker.
- [ ] Run the focused tests and confirm empty-queue cases fail with `404`.
- [ ] Special-case `errors.Is(err, apperrors.ErrNotFound)` in each claim handler and return `204` without registering a Gin error.
- [ ] Keep successful claims at `202` and other errors on the shared error path.
- [ ] Run focused handler tests, the handler package tests, formatting, complexity, and repository diff checks.
- [ ] Commit the backend change.

### Task 3: Lock down Agent idle-poll compatibility

**Files:**
- Modify: `internal/agent/runner/fake_control_plane_test.go`

- [ ] Add execution and Docker claim tests whose fake control plane returns `204` with an empty body.
- [ ] Assert each call reports a normal claim miss, returns no task, and records no successful claim.
- [ ] Run focused runner tests and the runner package race test.
- [ ] Commit the Agent regression tests.

### Task 4: Cross-repository verification

- [ ] Run the full contracts test gate.
- [ ] Run backend full tests, vet, complexity checks, and build checks appropriate to this transport-only change.
- [ ] Run Agent full tests, race tests, vet, and build checks.
- [ ] Run `git diff --check` and inspect status/log in all three repositories.
- [ ] Summarize commits, behavior, and any verification limitations before integration.
