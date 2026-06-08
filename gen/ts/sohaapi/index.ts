// Code generated from OpenSoha contracts. DO NOT EDIT.

export type AnyObject = Record<string, unknown>
export type ApiResponse<T = unknown> = { data: T }
export type ApiItemsResponse<T = unknown> = { items: T[] }
export type AuthTokens = TokenSet

export type AnyValue = unknown

export type GenericObject = AnyObject

export interface GenericDataEnvelope {
  data: unknown
}

export interface GenericItemsEnvelope {
  items: unknown[]
}

export interface ErrorEnvelope {
  error: {
  code: string
  message: string
  request_id?: string
}
}

export interface Principal {
  userId: string
  userName: string
  email: string
  roles: string[]
  teams: string[]
  projects: string[]
  tags: string[]
  permissionKeys?: string[]
  [key: string]: unknown
}

export interface UserProfile {
  userId: string
  username: string
  displayName: string
  email: string
  phone?: string
  status: string
  roles: string[]
  teams: string[]
  projects: string[]
  tags: string[]
  identities?: AnyObject[]
  sessions?: AnyObject[]
  lastLoginAt?: string
  [key: string]: unknown
}

export interface TokenSet {
  accessToken: string
  refreshToken: string
  tokenType: string
  expiresIn: number
  expiresAt: string
}

export interface AuthResult {
  user: Principal
  tokens: TokenSet
}

export interface AuthProvider {
  type: string
  id?: string
  name: string
  enabled: boolean
  loginUrl?: string
}

export interface LoginOptions {
  verification: {
  sliderEnabled: boolean
}
}

export interface PasswordLoginRequest {
  login: string
  password: string
}

export interface RefreshRequest {
  refreshToken?: string
}

export interface OIDCExchangeRequest {
  code: string
}

export interface StreamTicketRequest {
  path: string
}

export interface StreamTicket {
  ticket: string
  expiresAt: string
}

export interface PrincipalEnvelope {
  data: Principal
}

export interface UserProfileEnvelope {
  data: UserProfile
}

export interface AuthResultEnvelope {
  data: AuthResult
}

export interface AuthProviderListEnvelope {
  items: AuthProvider[]
}

export interface LoginOptionsEnvelope {
  data: LoginOptions
}

export interface AuthBootstrapEnvelope {
  data: {
  user: Principal
  currentUser: Principal
  permissionSnapshot: AnyObject
  branding: AnyObject
  [key: string]: unknown
}
}

export interface StreamTicketEnvelope {
  data: StreamTicket
}

export interface ExecutionTask {
  id: string
  applicationId: string
  applicationEnvironmentId: string
  taskKind: string
  providerKind: string
  status: string
  callbackToken: string
  payload: AnyObject
  [key: string]: unknown
}

export interface ExecutionTaskClaimRequest {
  agentId: string
  providerKinds?: string[]
  runtimeEndpoint?: string
}

export interface ExecutionCallbackRequest {
  callbackToken: string
  status: string
  payload: AnyObject
}

export interface DockerOperation {
  id: string
  hostId?: string
  projectId?: string
  serviceId?: string
  operationKind: string
  status: string
  claimedByWorkerId?: string
  payload: AnyObject
  timeoutSeconds: number
  [key: string]: unknown
}

export interface DockerOperationClaimRequest {
  workerId: string
  agentId: string
  hostIds?: string[]
  operationKinds?: string[]
}

export interface DockerOperationCallbackRequest {
  operationId: string
  workerId: string
  status: string
  payload: AnyObject
  logs: string[]
}

export interface AgentRun {
  id: string
  providerId: string
  providerKind: string
  capabilityId: string
  skillIds?: string[]
  sessionId?: string
  status: string
  scope?: AnyObject
  toolset?: AnyObject
  toolBindings?: AnyObject[]
  skillBindings?: AnyObject[]
  input?: AnyObject
  output?: AnyObject
  callbackToken: string
  timeoutSeconds: number
  analysisArtifacts?: AnyObject[]
  [key: string]: unknown
}

export interface AgentRunClaimRequest {
  agentId: string
  providerIds?: string[]
  kinds?: string[]
}

export interface AgentRunCallbackRequest {
  runId: string
  callbackToken: string
  agentId: string
  status: string
  payload: AnyObject
  toolExecutions?: AnyObject[]
  analysisArtifacts?: AnyObject[]
  externalRunId?: string
  errorMessage?: string
}

export interface AgentRunToolCallRequest {
  runId: string
  callbackToken: string
  agentId: string
  toolBindingId?: string
  adapterId?: string
  toolName?: string
  input?: AnyObject
}

export interface AgentToolCallResult {
  runId: string
  toolExecution: AnyObject
  output?: AnyObject
  [key: string]: unknown
}

export interface ExecutionTaskEnvelope {
  data: ExecutionTask
}

export interface DockerOperationEnvelope {
  data: DockerOperation
}

export interface AgentRunEnvelope {
  data: AgentRun
}

export interface AgentToolCallResultEnvelope {
  data: AgentToolCallResult
}

export type RiskLevel = "read" | "analyze" | "mutate" | "execute" | "high"

export type JSONSchema = AnyObject

export interface ToolCapability {
  name: string
  title: string
  description: string
  domain: string
  action: string
  riskLevel: RiskLevel
  permissionKeys: string[]
  requiredScopes?: string[]
  mcpAdapterId?: string
  mcpToolName?: string
  requiresApproval: boolean
  inputSchema?: JSONSchema
  outputSchema?: JSONSchema
}

export interface ResourceCapability {
  name: string
  description: string
  permissionKeys: string[]
  requiredScopes?: string[]
  contextSchema?: JSONSchema
}

export interface PromptCapability {
  name: string
  description: string
  permissionKeys: string[]
  requiredScopes?: string[]
  argumentSchema?: JSONSchema
  contextSchema?: JSONSchema
}

export interface SkillCapability {
  id: string
  name: string
  category: string
  description: string
  capabilityRefs?: string[]
  permissionKeys?: string[]
  requiredScopes?: string[]
}

export interface CallerContext {
  identityMode: string
  aiClientId?: string
  aiClientName?: string
  skillId?: string
  tokenId?: string
  sessionId?: string
  subjectType?: string
  subjectId?: string
  source?: string
}

export interface AIGatewayManifest {
  name: string
  version: string
  generatedAt: string
  principal: Principal
  caller: CallerContext
  permissionKeys: string[]
  tools: ToolCapability[]
  resources?: ResourceCapability[]
  prompts?: PromptCapability[]
  skills?: SkillCapability[]
  summary: {
  toolCount: number
  resourceCount: number
  promptCount: number
  skillCount: number
  deniedCount: number
}
}

export interface ToolInvocationRequest {
  toolName?: string
  input?: AnyObject
  aiClientId?: string
  aiClientName?: string
  skillId?: string
  requestId?: string
}

export interface ToolInvocationResult {
  toolName: string
  riskLevel: RiskLevel
  requiresApproval: boolean
  result: string
  output?: unknown
  relatedIds?: AnyObject
  audit?: AnyObject
  [key: string]: unknown
}

export interface ResourceReadRequest {
  name?: string
  uri?: string
  context?: AnyObject
  aiClientId?: string
  aiClientName?: string
  skillId?: string
  requestId?: string
}

export interface ResourceReadResult {
  name: string
  uri: string
  mimeType: string
  text?: string
  data?: unknown
  relatedIds?: AnyObject
  audit?: AnyObject
  [key: string]: unknown
}

export interface PromptGetRequest {
  name: string
  arguments?: AnyObject
  context?: AnyObject
  aiClientId?: string
  aiClientName?: string
  skillId?: string
  requestId?: string
}

export interface PromptMessage {
  role: string
  content: string
}

export interface PromptGetResult {
  name: string
  description: string
  messages: PromptMessage[]
  relatedIds?: AnyObject
  audit?: AnyObject
  [key: string]: unknown
}

export interface AIGatewayManifestEnvelope {
  data: AIGatewayManifest
}

export interface ToolInvocationResultEnvelope {
  data: ToolInvocationResult
}

export interface ResourceReadResultEnvelope {
  data: ResourceReadResult
}

export interface PromptGetResultEnvelope {
  data: PromptGetResult
}

export interface OperationStatus {
  status: string
}

export interface PluginAssetSnapshot {
  skills?: string[]
  mcpPresets?: string[]
  connectors?: string[]
  aiProviderAdapters?: string[]
  agentProfiles?: string[]
  gatewayPolicyPacks?: string[]
}

export interface PluginCompatibility {
  soha?: string
  contracts?: string
  [key: string]: unknown
}

export interface PluginCapabilityRequest {
  tools?: string[]
  resources?: string[]
  prompts?: string[]
  skills?: string[]
}

export interface PluginPermissionRequest {
  required?: string[]
  domain?: string[]
}

export interface PluginSecretRequirement {
  name: string
  description?: string
  required?: boolean
  secretRef?: string
}

export interface PluginIntegrity {
  checksum?: string
  signature?: string
  verified?: boolean
  status?: string
}

export interface PluginManifest {
  id: string
  name: string
  version: string
  publisher: string
  type: "skill" | "skill-pack" | "mcp-preset" | "connector" | "ai-provider-adapter" | "agent-profile" | "gateway-policy-pack"
  description?: string
  homepage?: string
  compatibility?: PluginCompatibility
  assets?: PluginAssetSnapshot
  capabilities?: PluginCapabilityRequest
  permissions?: PluginPermissionRequest
  secrets?: {
  required?: PluginSecretRequirement[]
}
  integrity?: PluginIntegrity
  metadata?: AnyObject
}

export interface MarketplacePlugin {
  id: string
  name: string
  version: string
  publisher: string
  type: string
  summary?: string
  source: string
  riskLevel?: string
  manifest: PluginManifest
  installed?: boolean
}

export interface InstalledPlugin {
  id: string
  name: string
  version: string
  publisher: string
  type: string
  status: "enabled" | "disabled"
  source: string
  manifest: PluginManifest
  checksumStatus: string
  signatureStatus?: string
  requestedPermissions?: PluginPermissionRequest
  configuredSecretRefs?: Record<string, string>
  installedBy: string
  installedAt: string
  updatedAt: string
  enabledAt?: string
  disabledAt?: string
  metadata?: AnyObject
}

export interface PluginInstallRequest {
  pluginId?: string
  source?: string
  manifest?: PluginManifest
  expectedChecksum?: string
  enable?: boolean
}

export interface PluginConfigRequest {
  enabled?: boolean
  secretRefs?: Record<string, string>
  metadata?: AnyObject
}

export interface PluginManifestEnvelope {
  data: PluginManifest
}

export interface MarketplacePluginEnvelope {
  data: MarketplacePlugin
}

export interface MarketplacePluginListEnvelope {
  items: MarketplacePlugin[]
}

export interface InstalledPluginEnvelope {
  data: InstalledPlugin
}

export interface InstalledPluginListEnvelope {
  items: InstalledPlugin[]
}

export interface PersonalAccessToken {
  id: string
  userId: string
  name: string
  tokenPrefix: string
  scopes: string[]
  permissionKeys: string[]
  metadata?: AnyObject
  expiresAt?: string
  lastUsedAt?: string
  revokedAt?: string
  createdBy: string
  createdAt: string
  updatedAt: string
}

export interface PersonalAccessTokenInput {
  name: string
  scopes: string[]
  permissionKeys: string[]
  metadata?: AnyObject
  expiresAt?: string
}

export interface CreatedPersonalAccessToken {
  token: PersonalAccessToken
  value: string
}

export interface PersonalAccessTokenListEnvelope {
  items: PersonalAccessToken[]
}

export interface CreatedPersonalAccessTokenEnvelope {
  data: CreatedPersonalAccessToken
}

export interface ServiceAccount {
  id: string
  name: string
  description?: string
  status: string
  ownerUserId?: string
  roleIds: string[]
  teamIds: string[]
  scopeGrantIds: string[]
  metadata?: AnyObject
  createdBy: string
  createdAt: string
  updatedAt: string
}

export interface ServiceAccountInput {
  id: string
  name: string
  description?: string
  status: string
  ownerUserId?: string
  roleIds?: string[]
  teamIds?: string[]
  scopeGrantIds?: string[]
  metadata?: AnyObject
}

export interface ServiceAccountEnvelope {
  data: ServiceAccount
}

export interface ServiceAccountListEnvelope {
  items: ServiceAccount[]
}

export interface ServiceAccountToken {
  id: string
  serviceAccountId: string
  name: string
  tokenPrefix: string
  scopes: string[]
  permissionKeys: string[]
  metadata?: AnyObject
  expiresAt?: string
  lastUsedAt?: string
  revokedAt?: string
  createdBy: string
  createdAt: string
  updatedAt: string
}

export interface ServiceAccountTokenInput {
  name: string
  scopes: string[]
  permissionKeys: string[]
  metadata?: AnyObject
  expiresAt?: string
}

export interface CreatedServiceAccountToken {
  token: ServiceAccountToken
  value: string
}

export interface ServiceAccountTokenListEnvelope {
  items: ServiceAccountToken[]
}

export interface CreatedServiceAccountTokenEnvelope {
  data: CreatedServiceAccountToken
}

export interface TokenRotationInput {
  expiresAt?: string
}

export interface AuditLog {
  id: string
  actorType: string
  actorId: string
  actorName?: string
  aiClientId?: string
  aiClientName?: string
  skillId?: string
  toolName?: string
  riskLevel?: RiskLevel
  resourceScope?: AnyObject
  action: string
  result: string
  summary: string
  requestId?: string
  sourceIp?: string
  metadata?: AnyObject
  createdAt: string
}

export interface AuditLogListEnvelope {
  items: AuditLog[]
}

export interface ApprovalTrace {
  approvalMode?: string
  currentStageIndex?: number
  currentStageName?: string
  stageCount?: number
  approvedCount?: number
  requiredApprovals?: number
  pendingRequirements?: string[]
  satisfiedRequirements?: string[]
  roleApprovedCounts?: Record<string, number>
  teamApprovedCounts?: Record<string, number>
  candidateUserIds?: string[]
  candidateRoles?: string[]
  candidateTeams?: string[]
  onCallCandidateUserIds?: string[]
  workflowRunId?: string
  executionTaskId?: string
  releaseBundleId?: string
  decisions?: ApprovalDecisionTrace[]
  stageHistory?: ApprovalStageTrace[]
}

export interface ApprovalDecisionTrace {
  userId?: string
  userName?: string
  roles?: string[]
  teams?: string[]
  result?: string
  comment?: string
  stageIndex?: number
  stageName?: string
  decidedAt?: string
}

export interface ApprovalStageTrace {
  stageIndex?: number
  stageName?: string
  result?: string
  completedAt?: string
}

export interface ApprovalTimelineEvent {
  id: string
  kind: string
  action: string
  result: string
  summary?: string
  actorType?: string
  actorId?: string
  actorName?: string
  stageIndex?: number
  stageName?: string
  metadata?: AnyObject
  createdAt: string
}

export interface ApprovalRequest {
  id: string
  status: string
  strategy: string
  policyId?: string
  approvalPolicyRef?: string
  actorType: string
  actorId: string
  actorName?: string
  actorRoles?: string[]
  actorTeams?: string[]
  aiClientId?: string
  aiClientName?: string
  skillId?: string
  toolName: string
  riskLevel: RiskLevel
  requiresApproval: boolean
  resourceScope?: AnyObject
  toolInput?: AnyObject
  relatedIds?: AnyObject
  approvalTrace?: ApprovalTrace
  output?: unknown
  summary: string
  requestId?: string
  sourceIp?: string
  decidedBy?: string
  decidedByName?: string
  decidedAt?: string
  decisionComment?: string
  expiresAt?: string
  createdAt: string
  updatedAt: string
}

export interface ApprovalDecisionInput {
  comment?: string
}

export interface ApprovalDecisionResult {
  request: ApprovalRequest
  invocation?: ToolInvocationResult
}

export interface ApprovalTimeline {
  request: ApprovalRequest
  trace?: ApprovalTrace
  events?: ApprovalTimelineEvent[]
}

export interface ApprovalRequestListEnvelope {
  items: ApprovalRequest[]
}

export interface ApprovalTimelineEnvelope {
  data: ApprovalTimeline
}

export interface ApprovalDecisionResultEnvelope {
  data: ApprovalDecisionResult
}

export interface GovernanceStatus {
  generatedAt: string
  windowHours: number
  health: GovernanceHealth
  metrics: GovernanceMetrics
  tokens: GovernanceTokenSummary
  clients: GovernanceClientSummary
  approvals: GovernanceApprovalSummary
  policyCoverage: GovernancePolicyCoverage
  redaction: GovernanceRedactionSummary
  anomalies?: GovernanceFinding[]
  recommendations?: string[]
  recommendationActions?: GovernanceRecommendationAction[]
  metadata?: AnyObject
}

export interface GovernanceHealth {
  status: string
  message: string
  checks?: GovernanceHealthCheck[]
}

export interface GovernanceHealthCheck {
  name: string
  status: string
  message: string
  count?: number
}

export interface GovernanceMetrics {
  totalCalls: number
  successCount: number
  denyCount: number
  failureCount: number
  pendingApprovalCount: number
  dryRunCount: number
  riskCounts?: Record<string, number>
  topTools?: GovernanceMetricCount[]
  topAiClients?: GovernanceMetricCount[]
  topActors?: GovernanceMetricCount[]
  recentResultBreakdown?: Record<string, number>
  recentActionBreakdown?: Record<string, number>
}

export interface GovernanceMetricCount {
  key: string
  count: number
}

export interface GovernanceRedactionSummary {
  totalMatches: number
  auditsWithRedaction: number
  inputAudits: number
  outputAudits: number
  fieldMatches: number
  sensitiveKeyMatches: number
  sensitiveTextMatches: number
  valuePatternMatches: number
  secretClassifierMatches: number
  structuredSecretMatches: number
  topTargets?: GovernanceMetricCount[]
  topFieldPaths?: GovernanceMetricCount[]
  topMatchTypes?: GovernanceMetricCount[]
  topClassifiers?: GovernanceMetricCount[]
  topPolicies?: GovernanceMetricCount[]
  topTools?: GovernanceMetricCount[]
}

export interface GovernanceTokenSummary {
  personalAccessTokens: GovernanceTokenCounts
  serviceAccountTokens: GovernanceTokenCounts
  expiringSoon?: GovernanceTokenFinding[]
  expiredActive?: GovernanceTokenFinding[]
  stale?: GovernanceTokenFinding[]
  neverUsed?: GovernanceTokenFinding[]
  lastUsedTrackingState: string
}

export interface GovernanceTokenCounts {
  total: number
  active: number
  revoked: number
  expired: number
  expiringSoon: number
  stale: number
  neverUsed: number
}

export interface GovernanceTokenFinding {
  kind: string
  id: string
  name: string
  ownerId?: string
  tokenPrefix: string
  severity: string
  message: string
  expiresAt?: string
  lastUsedAt?: string
  daysUntilDue?: number
  staleDays?: number
}

export interface GovernanceClientSummary {
  total: number
  active: number
  disabled: number
  pendingApproval: number
  registrationApproval: string
  pendingApprovalClientIds?: string[]
}

export interface GovernanceApprovalSummary {
  pending: number
  dueSoon: number
  stalePending: number
  overdue: number
  oldestPendingHours?: number
  oldestPendingRequestId?: string
  nextDueAt?: string
  nextDueRequestId?: string
  dueSoonRequestIds?: string[]
  stalePendingRequestIds?: string[]
  overdueRequestIds?: string[]
}

export interface GovernancePolicyCoverage {
  accessPolicies: number
  toolGrants: number
  skillBindings: number
  activeAccessPolicies: number
  activeToolGrants: number
  activeSkillBindings: number
  budgetPolicies: number
  rateLimitPolicies: number
  redactionPolicies: number
  resourceScopedAccessPolicies: number
  resourceScopedToolGrants: number
  budgetState: string
  rateLimitState: string
  redactionPolicyState: string
  resourceScopeState: string
}

export interface GovernanceFinding {
  type: string
  severity: string
  summary: string
  count?: number
  actorType?: string
  actorId?: string
  subjectType?: string
  subjectId?: string
  aiClientId?: string
  policyId?: string
  approvalRequestId?: string
  grantId?: string
  toolName?: string
  riskLevel?: RiskLevel
}

export interface GovernanceRecommendationAction {
  type: string
  severity: string
  summary: string
  action: string
  targetKind?: string
  targetId?: string
  refs?: string[]
  count?: number
  metadata?: AnyObject
}

export interface GovernanceStatusEnvelope {
  data: GovernanceStatus
}

export interface MCPCapability {
  adapterID: string
  name: string
  description: string
  scopes: string[]
}

export interface MCPCapabilityListEnvelope {
  items: MCPCapability[]
}
