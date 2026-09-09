# Network contract semantics

The JSON Schemas in this directory define strict serialized shapes. Every
consumer must also enforce the cross-field rules below before using a message
for authorization, routing, firewall configuration or accounting.

## Runtime messages

- Runtime mTLS certificates carry exactly one Soha URI SAN:
  `spiffe://opensoha.local/network-control/{runtimeKind}/{runtimeId}`.
  Ingest certificates use the separate scope
  `spiffe://opensoha.local/network-ingest/{producerKind}/{producerId}` and may
  not be reused for control. The body identity is only a consistency check.
- `devicePublicKey` is one PKIX `PUBLIC KEY` PEM block and must match the
  public key of the authenticated enrollment certificate.
- A runtime advertising the `wireguard` capability must submit exactly one
  WireGuard public key during enrollment. The corresponding private key is
  generated and retained in runtime-local protected storage. Re-enrollment
  advances the credential/key generation and retires the prior public key.
- Reject a message when `expiresAt` is not later than `occurredAt` or is
  already expired.
- Parse every CIDR with the platform standard library, require canonical
  network form, and reject anything not accepted by `netip.ParsePrefix` or an
  equivalent parser. `network-runtime/v1alpha1` carries IPv4 CIDRs only.
- Resolve `runtimeId` to the authenticated runtime credential. Never trust the
  body to select another runtime.
- Accept `nas.*` messages only from or for a runtime whose registered kind is
  `nas`; the authenticated runtime must also match the active NAS binding for
  `payload.nasId` and its site.
- Treat `nas.authorization.request` as authorization input only. It never
  carries a password, EAP exchange, RADIUS packet or shared secret. The
  `password-compatible` method is lower trust than `eap-tls` and policy may
  only preserve or reduce its access profile.
- For `eap-tls`, resolve the certificate's canonical serial plus authority key
  identifier to one active endpoint credential and derive `deviceId` from that
  credential. Match `subjectId` to the credential; treat `stationId` only as
  NAS evidence, never as device identity. Password-compatible requests identify
  the device explicitly and may not claim client-certificate binding fields.
- Emit RADIUS attributes only for an allow decision. The NAS adapter maps
  `vlanId`, `filterId` and `sessionTimeoutSeconds` to standard RADIUS
  attributes; vendor-specific attributes stay outside this contract.
- Match authorization result `requestId`, session command result `commandId`
  and `sessionId` to the outstanding request or command before changing state.
- Copy the Soha authorization `sessionId` into the standard RADIUS `Class`
  reply attribute and require NAS accounting to echo it as `sessionId` when
  available. Session commands repeat the bound `subjectId` and `deviceId` so
  the adapter can target CoA/Disconnect without querying management APIs.
- If a requested CoA is unsupported, issue a disconnect command when the NAS
  binding supports it; otherwise fail closed and record the unsupported
  enforcement result.
- For endpoint desired state, require every lease `deviceId` to match the
  enrolled device bound to `runtimeId`.
- Require all leases in one desired state to belong to the expected session
  and subject/device binding.
- Require lease `expiresAt` to be later than `issuedAt`, no later than the
  enclosing message/configuration validity, and within the configured maximum
  lease TTL.
- Apply only monotonically newer configuration versions. A rollback is a new
  version whose content intentionally references a prior snapshot.
- Treat `runtimeIntervals` as optional scheduling guidance for backward
  compatibility. When present, both heartbeat and configuration polling use
  the same 30–300 second bounds; clients apply local ±20% jitter.
- Accept `vpn.connect.request` only from an enrolled endpoint bound to the
  requested device. The selected site, NetworkSpace, gateway and NetworkLease
  must all be active and authorized by the loaded policy snapshot.
- `vpn.connect.request.gatewayId` only selects an ingress gateway; it never
  expands the requested site, NetworkSpace, resource or policy scope. When it
  is omitted, the control plane may select an eligible gateway using the same
  authorization checks.
- A gateway without `hubGatewayId` is a topology root. A gateway with
  `hubGatewayId` is a routed spoke of that active root. Spokes may not point to
  another spoke, use SNAT, or form direct spoke-to-spoke peers; traffic between
  spokes transits the root.
- Every `advertisedCidrs` entry must be canonical, belong to an active
  NetworkSpace at that gateway's site, and not overlap another active site or
  gateway overlay prefix. A site link is transport authorization, not user or
  resource authorization.
- WireGuard `allowedIPs` on an endpoint peer identify only that endpoint's
  overlay address. A site peer may additionally carry the explicit overlay and
  advertised prefixes of the remote topology members reachable through that
  peer. Endpoint split routes still come only from active leases, and the
  `0.0.0.0/0` default route is outside v1alpha1.
- Gateway firewall evaluation is ordered and default-deny. An exact active
  ResourceLease allow precedes the matching ProtectedSet deny, which in turn
  precedes a broader NetworkLease allow. Site-link rules carry
  `siteLinkRuntimeId` plus an explicit direction and expire with the peer
  credential; raw site LAN sources remain blocked from protected destinations.
- `routingMode=routed` requires the site LAN to route the overlay CIDR back to
  the gateway. `routingMode=snat` is an explicit deployment fallback and does
  not broaden any lease or firewall rule.
- Treat `credentialReference` only as an opaque runtime-local `secretref:`
  identifier. It is never dereferenced by an untrusted caller and never
  written to logs, telemetry or public responses.
- Validate ProtectedSet before NetworkLease coverage. A broad network lease
  cannot authorize a protected resource without a valid ResourceLease.

## Ingest batches

- Authenticate `producerId` from the transport credential; do not trust the
  body to select another producer.
- Deduplicate by `(producerId, event.id)` for the retention window. Exact
  duplicate objects are rejected by schema, but duplicates with changed
  payloads must also be rejected.
- Accept out-of-order `sequence` values as telemetry while recording gaps and
  regressions. Sequence never changes authorization.
- Require flow `windowEndedAt` to be later than `windowStartedAt` and
  `deniedConnections <= connections`.
- Reject events outside the configured clock-skew and retention window.
- Enforce HTTP body, decompressed body, event-count and per-producer rate
  limits before persistence.
- Accounting and flow events are observability input only. They cannot create,
  renew, broaden or revoke a NetworkLease or ResourceLease.

Consumers must cover these rules with deterministic tests. Failure is
fail-closed for control messages and bounded reject/drop-with-counter behavior
for telemetry.
