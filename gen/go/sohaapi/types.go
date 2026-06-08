// Code generated from OpenSoha contracts. DO NOT EDIT.
package sohaapi

import "time"

type AnyValue = any

type GenericObject map[string]any

type GenericDataEnvelope struct {
	Data any `json:"data"`
}

type GenericItemsEnvelope struct {
	Items []any `json:"items"`
}

type ErrorEnvelope struct {
	Error struct {
	Code string `json:"code"`
	Message string `json:"message"`
	RequestID string `json:"request_id,omitempty"`
} `json:"error"`
}

type Principal struct {
	UserID string `json:"userId"`
	UserName string `json:"userName"`
	Email string `json:"email"`
	Roles []string `json:"roles"`
	Teams []string `json:"teams"`
	Projects []string `json:"projects"`
	Tags []string `json:"tags"`
	PermissionKeys []string `json:"permissionKeys,omitempty"`
}

type UserProfile struct {
	UserID string `json:"userId"`
	Username string `json:"username"`
	DisplayName string `json:"displayName"`
	Email string `json:"email"`
	Phone string `json:"phone,omitempty"`
	Status string `json:"status"`
	Roles []string `json:"roles"`
	Teams []string `json:"teams"`
	Projects []string `json:"projects"`
	Tags []string `json:"tags"`
	Identities []map[string]any `json:"identities,omitempty"`
	Sessions []map[string]any `json:"sessions,omitempty"`
	LastLoginAt *time.Time `json:"lastLoginAt,omitempty"`
}

type TokenSet struct {
	AccessToken string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
	TokenType string `json:"tokenType"`
	ExpiresIn int64 `json:"expiresIn"`
	ExpiresAt time.Time `json:"expiresAt"`
}

type AuthResult struct {
	User Principal `json:"user"`
	Tokens TokenSet `json:"tokens"`
}

type AuthProvider struct {
	Type string `json:"type"`
	ID string `json:"id,omitempty"`
	Name string `json:"name"`
	Enabled bool `json:"enabled"`
	LoginURL string `json:"loginUrl,omitempty"`
}

type LoginOptions struct {
	Verification struct {
	SliderEnabled bool `json:"sliderEnabled"`
} `json:"verification"`
}

type PasswordLoginRequest struct {
	Login string `json:"login"`
	Password string `json:"password"`
}

type RefreshRequest struct {
	RefreshToken string `json:"refreshToken,omitempty"`
}

type OIDCExchangeRequest struct {
	Code string `json:"code"`
}

type StreamTicketRequest struct {
	Path string `json:"path"`
}

type StreamTicket struct {
	Ticket string `json:"ticket"`
	ExpiresAt time.Time `json:"expiresAt"`
}

type PrincipalEnvelope struct {
	Data Principal `json:"data"`
}

type UserProfileEnvelope struct {
	Data UserProfile `json:"data"`
}

type AuthResultEnvelope struct {
	Data AuthResult `json:"data"`
}

type AuthProviderListEnvelope struct {
	Items []AuthProvider `json:"items"`
}

type LoginOptionsEnvelope struct {
	Data LoginOptions `json:"data"`
}

type AuthBootstrapEnvelope struct {
	Data struct {
	User Principal `json:"user"`
	CurrentUser Principal `json:"currentUser"`
	PermissionSnapshot map[string]any `json:"permissionSnapshot"`
	Branding map[string]any `json:"branding"`
} `json:"data"`
}

type StreamTicketEnvelope struct {
	Data StreamTicket `json:"data"`
}

type ExecutionTask struct {
	ID string `json:"id"`
	ApplicationID string `json:"applicationId"`
	ApplicationEnvironmentID string `json:"applicationEnvironmentId"`
	TaskKind string `json:"taskKind"`
	ProviderKind string `json:"providerKind"`
	Status string `json:"status"`
	CallbackToken string `json:"callbackToken"`
	Payload map[string]any `json:"payload"`
}

type ExecutionTaskClaimRequest struct {
	AgentID string `json:"agentId"`
	ProviderKinds []string `json:"providerKinds,omitempty"`
	RuntimeEndpoint string `json:"runtimeEndpoint,omitempty"`
}

type ExecutionCallbackRequest struct {
	CallbackToken string `json:"callbackToken"`
	Status string `json:"status"`
	Payload map[string]any `json:"payload"`
}

type DockerOperation struct {
	ID string `json:"id"`
	HostID string `json:"hostId,omitempty"`
	ProjectID string `json:"projectId,omitempty"`
	ServiceID string `json:"serviceId,omitempty"`
	OperationKind string `json:"operationKind"`
	Status string `json:"status"`
	ClaimedByWorkerID string `json:"claimedByWorkerId,omitempty"`
	Payload map[string]any `json:"payload"`
	TimeoutSeconds int `json:"timeoutSeconds"`
}

type DockerOperationClaimRequest struct {
	WorkerID string `json:"workerId"`
	AgentID string `json:"agentId"`
	HostIDs []string `json:"hostIds,omitempty"`
	OperationKinds []string `json:"operationKinds,omitempty"`
}

type DockerOperationCallbackRequest struct {
	OperationID string `json:"operationId"`
	WorkerID string `json:"workerId"`
	Status string `json:"status"`
	Payload map[string]any `json:"payload"`
	Logs []string `json:"logs"`
}

type AgentRun struct {
	ID string `json:"id"`
	ProviderID string `json:"providerId"`
	ProviderKind string `json:"providerKind"`
	CapabilityID string `json:"capabilityId"`
	SkillIDs []string `json:"skillIds,omitempty"`
	SessionID string `json:"sessionId,omitempty"`
	Status string `json:"status"`
	Scope map[string]any `json:"scope,omitempty"`
	Toolset map[string]any `json:"toolset,omitempty"`
	ToolBindings []map[string]any `json:"toolBindings,omitempty"`
	SkillBindings []map[string]any `json:"skillBindings,omitempty"`
	Input map[string]any `json:"input,omitempty"`
	Output map[string]any `json:"output,omitempty"`
	CallbackToken string `json:"callbackToken"`
	TimeoutSeconds int `json:"timeoutSeconds"`
	AnalysisArtifacts []map[string]any `json:"analysisArtifacts,omitempty"`
}

type AgentRunClaimRequest struct {
	AgentID string `json:"agentId"`
	ProviderIDs []string `json:"providerIds,omitempty"`
	Kinds []string `json:"kinds,omitempty"`
}

type AgentRunCallbackRequest struct {
	RunID string `json:"runId"`
	CallbackToken string `json:"callbackToken"`
	AgentID string `json:"agentId"`
	Status string `json:"status"`
	Payload map[string]any `json:"payload"`
	ToolExecutions []map[string]any `json:"toolExecutions,omitempty"`
	AnalysisArtifacts []map[string]any `json:"analysisArtifacts,omitempty"`
	ExternalRunID string `json:"externalRunId,omitempty"`
	ErrorMessage string `json:"errorMessage,omitempty"`
}

type AgentRunToolCallRequest struct {
	RunID string `json:"runId"`
	CallbackToken string `json:"callbackToken"`
	AgentID string `json:"agentId"`
	ToolBindingID string `json:"toolBindingId,omitempty"`
	AdapterID string `json:"adapterId,omitempty"`
	ToolName string `json:"toolName,omitempty"`
	Input map[string]any `json:"input,omitempty"`
}

type AgentToolCallResult struct {
	RunID string `json:"runId"`
	ToolExecution map[string]any `json:"toolExecution"`
	Output map[string]any `json:"output,omitempty"`
}

type ExecutionTaskEnvelope struct {
	Data ExecutionTask `json:"data"`
}

type DockerOperationEnvelope struct {
	Data DockerOperation `json:"data"`
}

type AgentRunEnvelope struct {
	Data AgentRun `json:"data"`
}

type AgentToolCallResultEnvelope struct {
	Data AgentToolCallResult `json:"data"`
}

type RiskLevel string

type JSONSchema map[string]any

type ToolCapability struct {
	Name string `json:"name"`
	Title string `json:"title"`
	Description string `json:"description"`
	Domain string `json:"domain"`
	Action string `json:"action"`
	RiskLevel string `json:"riskLevel"`
	PermissionKeys []string `json:"permissionKeys"`
	RequiredScopes []string `json:"requiredScopes,omitempty"`
	MCPAdapterID string `json:"mcpAdapterId,omitempty"`
	MCPToolName string `json:"mcpToolName,omitempty"`
	RequiresApproval bool `json:"requiresApproval"`
	InputSchema JSONSchema `json:"inputSchema,omitempty"`
	OutputSchema JSONSchema `json:"outputSchema,omitempty"`
}

type ResourceCapability struct {
	Name string `json:"name"`
	Description string `json:"description"`
	PermissionKeys []string `json:"permissionKeys"`
	RequiredScopes []string `json:"requiredScopes,omitempty"`
	ContextSchema JSONSchema `json:"contextSchema,omitempty"`
}

type PromptCapability struct {
	Name string `json:"name"`
	Description string `json:"description"`
	PermissionKeys []string `json:"permissionKeys"`
	RequiredScopes []string `json:"requiredScopes,omitempty"`
	ArgumentSchema JSONSchema `json:"argumentSchema,omitempty"`
	ContextSchema JSONSchema `json:"contextSchema,omitempty"`
}

type SkillCapability struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Category string `json:"category"`
	Description string `json:"description"`
	CapabilityRefs []string `json:"capabilityRefs,omitempty"`
	PermissionKeys []string `json:"permissionKeys,omitempty"`
	RequiredScopes []string `json:"requiredScopes,omitempty"`
}

type CallerContext struct {
	IdentityMode string `json:"identityMode"`
	AIClientID string `json:"aiClientId,omitempty"`
	AIClientName string `json:"aiClientName,omitempty"`
	SkillID string `json:"skillId,omitempty"`
	TokenID string `json:"tokenId,omitempty"`
	SessionID string `json:"sessionId,omitempty"`
	SubjectType string `json:"subjectType,omitempty"`
	SubjectID string `json:"subjectId,omitempty"`
	Source string `json:"source,omitempty"`
}

type AIGatewayManifest struct {
	Name string `json:"name"`
	Version string `json:"version"`
	GeneratedAt time.Time `json:"generatedAt"`
	Principal Principal `json:"principal"`
	Caller CallerContext `json:"caller"`
	PermissionKeys []string `json:"permissionKeys"`
	Tools []ToolCapability `json:"tools"`
	Resources []ResourceCapability `json:"resources,omitempty"`
	Prompts []PromptCapability `json:"prompts,omitempty"`
	Skills []SkillCapability `json:"skills,omitempty"`
	Summary struct {
	ToolCount int `json:"toolCount"`
	ResourceCount int `json:"resourceCount"`
	PromptCount int `json:"promptCount"`
	SkillCount int `json:"skillCount"`
	DeniedCount int `json:"deniedCount"`
} `json:"summary"`
}

type ToolInvocationRequest struct {
	ToolName string `json:"toolName,omitempty"`
	Input map[string]any `json:"input,omitempty"`
	AIClientID string `json:"aiClientId,omitempty"`
	AIClientName string `json:"aiClientName,omitempty"`
	SkillID string `json:"skillId,omitempty"`
	RequestID string `json:"requestId,omitempty"`
}

type ToolInvocationResult struct {
	ToolName string `json:"toolName"`
	RiskLevel string `json:"riskLevel"`
	RequiresApproval bool `json:"requiresApproval"`
	Result string `json:"result"`
	Output any `json:"output,omitempty"`
	RelatedIDs map[string]any `json:"relatedIds,omitempty"`
	Audit map[string]any `json:"audit,omitempty"`
}

type ResourceReadRequest struct {
	Name string `json:"name,omitempty"`
	URI string `json:"uri,omitempty"`
	Context map[string]any `json:"context,omitempty"`
	AIClientID string `json:"aiClientId,omitempty"`
	AIClientName string `json:"aiClientName,omitempty"`
	SkillID string `json:"skillId,omitempty"`
	RequestID string `json:"requestId,omitempty"`
}

type ResourceReadResult struct {
	Name string `json:"name"`
	URI string `json:"uri"`
	MIMEType string `json:"mimeType"`
	Text string `json:"text,omitempty"`
	Data any `json:"data,omitempty"`
	RelatedIDs map[string]any `json:"relatedIds,omitempty"`
	Audit map[string]any `json:"audit,omitempty"`
}

type PromptGetRequest struct {
	Name string `json:"name"`
	Arguments map[string]any `json:"arguments,omitempty"`
	Context map[string]any `json:"context,omitempty"`
	AIClientID string `json:"aiClientId,omitempty"`
	AIClientName string `json:"aiClientName,omitempty"`
	SkillID string `json:"skillId,omitempty"`
	RequestID string `json:"requestId,omitempty"`
}

type PromptMessage struct {
	Role string `json:"role"`
	Content string `json:"content"`
}

type PromptGetResult struct {
	Name string `json:"name"`
	Description string `json:"description"`
	Messages []PromptMessage `json:"messages"`
	RelatedIDs map[string]any `json:"relatedIds,omitempty"`
	Audit map[string]any `json:"audit,omitempty"`
}

type AIGatewayManifestEnvelope struct {
	Data AIGatewayManifest `json:"data"`
}

type ToolInvocationResultEnvelope struct {
	Data ToolInvocationResult `json:"data"`
}

type ResourceReadResultEnvelope struct {
	Data ResourceReadResult `json:"data"`
}

type PromptGetResultEnvelope struct {
	Data PromptGetResult `json:"data"`
}

type OperationStatus struct {
	Status string `json:"status"`
}

type PluginAssetSnapshot struct {
	Skills []string `json:"skills,omitempty"`
	MCPPresets []string `json:"mcpPresets,omitempty"`
	Connectors []string `json:"connectors,omitempty"`
	AIProviderAdapters []string `json:"aiProviderAdapters,omitempty"`
	AgentProfiles []string `json:"agentProfiles,omitempty"`
	GatewayPolicyPacks []string `json:"gatewayPolicyPacks,omitempty"`
}

type PluginCompatibility struct {
	Soha string `json:"soha,omitempty"`
	Contracts string `json:"contracts,omitempty"`
}

type PluginCapabilityRequest struct {
	Tools []string `json:"tools,omitempty"`
	Resources []string `json:"resources,omitempty"`
	Prompts []string `json:"prompts,omitempty"`
	Skills []string `json:"skills,omitempty"`
}

type PluginPermissionRequest struct {
	Required []string `json:"required,omitempty"`
	Domain []string `json:"domain,omitempty"`
}

type PluginSecretRequirement struct {
	Name string `json:"name"`
	Description string `json:"description,omitempty"`
	Required bool `json:"required,omitempty"`
	SecretRef string `json:"secretRef,omitempty"`
}

type PluginIntegrity struct {
	Checksum string `json:"checksum,omitempty"`
	Signature string `json:"signature,omitempty"`
	Verified bool `json:"verified,omitempty"`
	Status string `json:"status,omitempty"`
}

type PluginManifest struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Version string `json:"version"`
	Publisher string `json:"publisher"`
	Type string `json:"type"`
	Description string `json:"description,omitempty"`
	Homepage string `json:"homepage,omitempty"`
	Compatibility *PluginCompatibility `json:"compatibility,omitempty"`
	Assets *PluginAssetSnapshot `json:"assets,omitempty"`
	Capabilities *PluginCapabilityRequest `json:"capabilities,omitempty"`
	Permissions *PluginPermissionRequest `json:"permissions,omitempty"`
	Secrets *struct {
	Required []PluginSecretRequirement `json:"required,omitempty"`
} `json:"secrets,omitempty"`
	Integrity *PluginIntegrity `json:"integrity,omitempty"`
	Metadata map[string]any `json:"metadata,omitempty"`
}

type MarketplacePlugin struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Version string `json:"version"`
	Publisher string `json:"publisher"`
	Type string `json:"type"`
	Summary string `json:"summary,omitempty"`
	Source string `json:"source"`
	RiskLevel string `json:"riskLevel,omitempty"`
	Manifest PluginManifest `json:"manifest"`
	Installed bool `json:"installed,omitempty"`
}

type InstalledPlugin struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Version string `json:"version"`
	Publisher string `json:"publisher"`
	Type string `json:"type"`
	Status string `json:"status"`
	Source string `json:"source"`
	Manifest PluginManifest `json:"manifest"`
	ChecksumStatus string `json:"checksumStatus"`
	SignatureStatus string `json:"signatureStatus,omitempty"`
	RequestedPermissions *PluginPermissionRequest `json:"requestedPermissions,omitempty"`
	ConfiguredSecretRefs map[string]string `json:"configuredSecretRefs,omitempty"`
	InstalledBy string `json:"installedBy"`
	InstalledAt time.Time `json:"installedAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	EnabledAt *time.Time `json:"enabledAt,omitempty"`
	DisabledAt *time.Time `json:"disabledAt,omitempty"`
	Metadata map[string]any `json:"metadata,omitempty"`
}

type PluginInstallRequest struct {
	PluginID string `json:"pluginId,omitempty"`
	Source string `json:"source,omitempty"`
	Manifest *PluginManifest `json:"manifest,omitempty"`
	ExpectedChecksum string `json:"expectedChecksum,omitempty"`
	Enable bool `json:"enable,omitempty"`
}

type PluginConfigRequest struct {
	Enabled *bool `json:"enabled,omitempty"`
	SecretRefs map[string]string `json:"secretRefs,omitempty"`
	Metadata map[string]any `json:"metadata,omitempty"`
}

type PluginManifestEnvelope struct {
	Data PluginManifest `json:"data"`
}

type MarketplacePluginEnvelope struct {
	Data MarketplacePlugin `json:"data"`
}

type MarketplacePluginListEnvelope struct {
	Items []MarketplacePlugin `json:"items"`
}

type InstalledPluginEnvelope struct {
	Data InstalledPlugin `json:"data"`
}

type InstalledPluginListEnvelope struct {
	Items []InstalledPlugin `json:"items"`
}

type PersonalAccessToken struct {
	ID string `json:"id"`
	UserID string `json:"userId"`
	Name string `json:"name"`
	TokenPrefix string `json:"tokenPrefix"`
	Scopes []string `json:"scopes"`
	PermissionKeys []string `json:"permissionKeys"`
	Metadata map[string]any `json:"metadata,omitempty"`
	ExpiresAt *time.Time `json:"expiresAt,omitempty"`
	LastUsedAt *time.Time `json:"lastUsedAt,omitempty"`
	RevokedAt *time.Time `json:"revokedAt,omitempty"`
	CreatedBy string `json:"createdBy"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type PersonalAccessTokenInput struct {
	Name string `json:"name"`
	Scopes []string `json:"scopes"`
	PermissionKeys []string `json:"permissionKeys"`
	Metadata map[string]any `json:"metadata,omitempty"`
	ExpiresAt *time.Time `json:"expiresAt,omitempty"`
}

type CreatedPersonalAccessToken struct {
	Token PersonalAccessToken `json:"token"`
	Value string `json:"value"`
}

type PersonalAccessTokenListEnvelope struct {
	Items []PersonalAccessToken `json:"items"`
}

type CreatedPersonalAccessTokenEnvelope struct {
	Data CreatedPersonalAccessToken `json:"data"`
}

type ServiceAccount struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Description string `json:"description,omitempty"`
	Status string `json:"status"`
	OwnerUserID string `json:"ownerUserId,omitempty"`
	RoleIDs []string `json:"roleIds"`
	TeamIDs []string `json:"teamIds"`
	ScopeGrantIDs []string `json:"scopeGrantIds"`
	Metadata map[string]any `json:"metadata,omitempty"`
	CreatedBy string `json:"createdBy"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type ServiceAccountInput struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Description string `json:"description,omitempty"`
	Status string `json:"status"`
	OwnerUserID string `json:"ownerUserId,omitempty"`
	RoleIDs []string `json:"roleIds,omitempty"`
	TeamIDs []string `json:"teamIds,omitempty"`
	ScopeGrantIDs []string `json:"scopeGrantIds,omitempty"`
	Metadata map[string]any `json:"metadata,omitempty"`
}

type ServiceAccountEnvelope struct {
	Data ServiceAccount `json:"data"`
}

type ServiceAccountListEnvelope struct {
	Items []ServiceAccount `json:"items"`
}

type ServiceAccountToken struct {
	ID string `json:"id"`
	ServiceAccountID string `json:"serviceAccountId"`
	Name string `json:"name"`
	TokenPrefix string `json:"tokenPrefix"`
	Scopes []string `json:"scopes"`
	PermissionKeys []string `json:"permissionKeys"`
	Metadata map[string]any `json:"metadata,omitempty"`
	ExpiresAt *time.Time `json:"expiresAt,omitempty"`
	LastUsedAt *time.Time `json:"lastUsedAt,omitempty"`
	RevokedAt *time.Time `json:"revokedAt,omitempty"`
	CreatedBy string `json:"createdBy"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type ServiceAccountTokenInput struct {
	Name string `json:"name"`
	Scopes []string `json:"scopes"`
	PermissionKeys []string `json:"permissionKeys"`
	Metadata map[string]any `json:"metadata,omitempty"`
	ExpiresAt *time.Time `json:"expiresAt,omitempty"`
}

type CreatedServiceAccountToken struct {
	Token ServiceAccountToken `json:"token"`
	Value string `json:"value"`
}

type ServiceAccountTokenListEnvelope struct {
	Items []ServiceAccountToken `json:"items"`
}

type CreatedServiceAccountTokenEnvelope struct {
	Data CreatedServiceAccountToken `json:"data"`
}

type TokenRotationInput struct {
	ExpiresAt *time.Time `json:"expiresAt,omitempty"`
}

type AuditLog struct {
	ID string `json:"id"`
	ActorType string `json:"actorType"`
	ActorID string `json:"actorId"`
	ActorName string `json:"actorName,omitempty"`
	AIClientID string `json:"aiClientId,omitempty"`
	AIClientName string `json:"aiClientName,omitempty"`
	SkillID string `json:"skillId,omitempty"`
	ToolName string `json:"toolName,omitempty"`
	RiskLevel string `json:"riskLevel,omitempty"`
	ResourceScope map[string]any `json:"resourceScope,omitempty"`
	Action string `json:"action"`
	Result string `json:"result"`
	Summary string `json:"summary"`
	RequestID string `json:"requestId,omitempty"`
	SourceIP string `json:"sourceIp,omitempty"`
	Metadata map[string]any `json:"metadata,omitempty"`
	CreatedAt time.Time `json:"createdAt"`
}

type AuditLogListEnvelope struct {
	Items []AuditLog `json:"items"`
}

type ApprovalTrace struct {
	ApprovalMode string `json:"approvalMode,omitempty"`
	CurrentStageIndex int `json:"currentStageIndex,omitempty"`
	CurrentStageName string `json:"currentStageName,omitempty"`
	StageCount int `json:"stageCount,omitempty"`
	ApprovedCount int `json:"approvedCount,omitempty"`
	RequiredApprovals int `json:"requiredApprovals,omitempty"`
	PendingRequirements []string `json:"pendingRequirements,omitempty"`
	SatisfiedRequirements []string `json:"satisfiedRequirements,omitempty"`
	RoleApprovedCounts map[string]int `json:"roleApprovedCounts,omitempty"`
	TeamApprovedCounts map[string]int `json:"teamApprovedCounts,omitempty"`
	CandidateUserIDs []string `json:"candidateUserIds,omitempty"`
	CandidateRoles []string `json:"candidateRoles,omitempty"`
	CandidateTeams []string `json:"candidateTeams,omitempty"`
	OnCallCandidateUserIDs []string `json:"onCallCandidateUserIds,omitempty"`
	WorkflowRunID string `json:"workflowRunId,omitempty"`
	ExecutionTaskID string `json:"executionTaskId,omitempty"`
	ReleaseBundleID string `json:"releaseBundleId,omitempty"`
	Decisions []ApprovalDecisionTrace `json:"decisions,omitempty"`
	StageHistory []ApprovalStageTrace `json:"stageHistory,omitempty"`
}

type ApprovalDecisionTrace struct {
	UserID string `json:"userId,omitempty"`
	UserName string `json:"userName,omitempty"`
	Roles []string `json:"roles,omitempty"`
	Teams []string `json:"teams,omitempty"`
	Result string `json:"result,omitempty"`
	Comment string `json:"comment,omitempty"`
	StageIndex int `json:"stageIndex,omitempty"`
	StageName string `json:"stageName,omitempty"`
	DecidedAt *time.Time `json:"decidedAt,omitempty"`
}

type ApprovalStageTrace struct {
	StageIndex int `json:"stageIndex,omitempty"`
	StageName string `json:"stageName,omitempty"`
	Result string `json:"result,omitempty"`
	CompletedAt *time.Time `json:"completedAt,omitempty"`
}

type ApprovalTimelineEvent struct {
	ID string `json:"id"`
	Kind string `json:"kind"`
	Action string `json:"action"`
	Result string `json:"result"`
	Summary string `json:"summary,omitempty"`
	ActorType string `json:"actorType,omitempty"`
	ActorID string `json:"actorId,omitempty"`
	ActorName string `json:"actorName,omitempty"`
	StageIndex int `json:"stageIndex,omitempty"`
	StageName string `json:"stageName,omitempty"`
	Metadata map[string]any `json:"metadata,omitempty"`
	CreatedAt time.Time `json:"createdAt"`
}

type ApprovalRequest struct {
	ID string `json:"id"`
	Status string `json:"status"`
	Strategy string `json:"strategy"`
	PolicyID string `json:"policyId,omitempty"`
	ApprovalPolicyRef string `json:"approvalPolicyRef,omitempty"`
	ActorType string `json:"actorType"`
	ActorID string `json:"actorId"`
	ActorName string `json:"actorName,omitempty"`
	ActorRoles []string `json:"actorRoles,omitempty"`
	ActorTeams []string `json:"actorTeams,omitempty"`
	AIClientID string `json:"aiClientId,omitempty"`
	AIClientName string `json:"aiClientName,omitempty"`
	SkillID string `json:"skillId,omitempty"`
	ToolName string `json:"toolName"`
	RiskLevel string `json:"riskLevel"`
	RequiresApproval bool `json:"requiresApproval"`
	ResourceScope map[string]any `json:"resourceScope,omitempty"`
	ToolInput map[string]any `json:"toolInput,omitempty"`
	RelatedIDs map[string]any `json:"relatedIds,omitempty"`
	ApprovalTrace *ApprovalTrace `json:"approvalTrace,omitempty"`
	Output any `json:"output,omitempty"`
	Summary string `json:"summary"`
	RequestID string `json:"requestId,omitempty"`
	SourceIP string `json:"sourceIp,omitempty"`
	DecidedBy string `json:"decidedBy,omitempty"`
	DecidedByName string `json:"decidedByName,omitempty"`
	DecidedAt *time.Time `json:"decidedAt,omitempty"`
	DecisionComment string `json:"decisionComment,omitempty"`
	ExpiresAt *time.Time `json:"expiresAt,omitempty"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type ApprovalDecisionInput struct {
	Comment string `json:"comment,omitempty"`
}

type ApprovalDecisionResult struct {
	Request ApprovalRequest `json:"request"`
	Invocation *ToolInvocationResult `json:"invocation,omitempty"`
}

type ApprovalTimeline struct {
	Request ApprovalRequest `json:"request"`
	Trace *ApprovalTrace `json:"trace,omitempty"`
	Events []ApprovalTimelineEvent `json:"events,omitempty"`
}

type ApprovalRequestListEnvelope struct {
	Items []ApprovalRequest `json:"items"`
}

type ApprovalTimelineEnvelope struct {
	Data ApprovalTimeline `json:"data"`
}

type ApprovalDecisionResultEnvelope struct {
	Data ApprovalDecisionResult `json:"data"`
}

type GovernanceStatus struct {
	GeneratedAt time.Time `json:"generatedAt"`
	WindowHours int `json:"windowHours"`
	Health GovernanceHealth `json:"health"`
	Metrics GovernanceMetrics `json:"metrics"`
	Tokens GovernanceTokenSummary `json:"tokens"`
	Clients GovernanceClientSummary `json:"clients"`
	Approvals GovernanceApprovalSummary `json:"approvals"`
	PolicyCoverage GovernancePolicyCoverage `json:"policyCoverage"`
	Redaction GovernanceRedactionSummary `json:"redaction"`
	Anomalies []GovernanceFinding `json:"anomalies,omitempty"`
	Recommendations []string `json:"recommendations,omitempty"`
	RecommendationActions []GovernanceRecommendationAction `json:"recommendationActions,omitempty"`
	Metadata map[string]any `json:"metadata,omitempty"`
}

type GovernanceHealth struct {
	Status string `json:"status"`
	Message string `json:"message"`
	Checks []GovernanceHealthCheck `json:"checks,omitempty"`
}

type GovernanceHealthCheck struct {
	Name string `json:"name"`
	Status string `json:"status"`
	Message string `json:"message"`
	Count int `json:"count,omitempty"`
}

type GovernanceMetrics struct {
	TotalCalls int `json:"totalCalls"`
	SuccessCount int `json:"successCount"`
	DenyCount int `json:"denyCount"`
	FailureCount int `json:"failureCount"`
	PendingApprovalCount int `json:"pendingApprovalCount"`
	DryRunCount int `json:"dryRunCount"`
	RiskCounts map[string]int `json:"riskCounts,omitempty"`
	TopTools []GovernanceMetricCount `json:"topTools,omitempty"`
	TopAIClients []GovernanceMetricCount `json:"topAiClients,omitempty"`
	TopActors []GovernanceMetricCount `json:"topActors,omitempty"`
	RecentResultBreakdown map[string]int `json:"recentResultBreakdown,omitempty"`
	RecentActionBreakdown map[string]int `json:"recentActionBreakdown,omitempty"`
}

type GovernanceMetricCount struct {
	Key string `json:"key"`
	Count int `json:"count"`
}

type GovernanceRedactionSummary struct {
	TotalMatches int `json:"totalMatches"`
	AuditsWithRedaction int `json:"auditsWithRedaction"`
	InputAudits int `json:"inputAudits"`
	OutputAudits int `json:"outputAudits"`
	FieldMatches int `json:"fieldMatches"`
	SensitiveKeyMatches int `json:"sensitiveKeyMatches"`
	SensitiveTextMatches int `json:"sensitiveTextMatches"`
	ValuePatternMatches int `json:"valuePatternMatches"`
	SecretClassifierMatches int `json:"secretClassifierMatches"`
	StructuredSecretMatches int `json:"structuredSecretMatches"`
	TopTargets []GovernanceMetricCount `json:"topTargets,omitempty"`
	TopFieldPaths []GovernanceMetricCount `json:"topFieldPaths,omitempty"`
	TopMatchTypes []GovernanceMetricCount `json:"topMatchTypes,omitempty"`
	TopClassifiers []GovernanceMetricCount `json:"topClassifiers,omitempty"`
	TopPolicies []GovernanceMetricCount `json:"topPolicies,omitempty"`
	TopTools []GovernanceMetricCount `json:"topTools,omitempty"`
}

type GovernanceTokenSummary struct {
	PersonalAccessTokens GovernanceTokenCounts `json:"personalAccessTokens"`
	ServiceAccountTokens GovernanceTokenCounts `json:"serviceAccountTokens"`
	ExpiringSoon []GovernanceTokenFinding `json:"expiringSoon,omitempty"`
	ExpiredActive []GovernanceTokenFinding `json:"expiredActive,omitempty"`
	Stale []GovernanceTokenFinding `json:"stale,omitempty"`
	NeverUsed []GovernanceTokenFinding `json:"neverUsed,omitempty"`
	LastUsedTrackingState string `json:"lastUsedTrackingState"`
}

type GovernanceTokenCounts struct {
	Total int `json:"total"`
	Active int `json:"active"`
	Revoked int `json:"revoked"`
	Expired int `json:"expired"`
	ExpiringSoon int `json:"expiringSoon"`
	Stale int `json:"stale"`
	NeverUsed int `json:"neverUsed"`
}

type GovernanceTokenFinding struct {
	Kind string `json:"kind"`
	ID string `json:"id"`
	Name string `json:"name"`
	OwnerID string `json:"ownerId,omitempty"`
	TokenPrefix string `json:"tokenPrefix"`
	Severity string `json:"severity"`
	Message string `json:"message"`
	ExpiresAt *time.Time `json:"expiresAt,omitempty"`
	LastUsedAt *time.Time `json:"lastUsedAt,omitempty"`
	DaysUntilDue int `json:"daysUntilDue,omitempty"`
	StaleDays int `json:"staleDays,omitempty"`
}

type GovernanceClientSummary struct {
	Total int `json:"total"`
	Active int `json:"active"`
	Disabled int `json:"disabled"`
	PendingApproval int `json:"pendingApproval"`
	RegistrationApproval string `json:"registrationApproval"`
	PendingApprovalClientIDs []string `json:"pendingApprovalClientIds,omitempty"`
}

type GovernanceApprovalSummary struct {
	Pending int `json:"pending"`
	DueSoon int `json:"dueSoon"`
	StalePending int `json:"stalePending"`
	Overdue int `json:"overdue"`
	OldestPendingHours int `json:"oldestPendingHours,omitempty"`
	OldestPendingRequestID string `json:"oldestPendingRequestId,omitempty"`
	NextDueAt *time.Time `json:"nextDueAt,omitempty"`
	NextDueRequestID string `json:"nextDueRequestId,omitempty"`
	DueSoonRequestIDs []string `json:"dueSoonRequestIds,omitempty"`
	StalePendingRequestIDs []string `json:"stalePendingRequestIds,omitempty"`
	OverdueRequestIDs []string `json:"overdueRequestIds,omitempty"`
}

type GovernancePolicyCoverage struct {
	AccessPolicies int `json:"accessPolicies"`
	ToolGrants int `json:"toolGrants"`
	SkillBindings int `json:"skillBindings"`
	ActiveAccessPolicies int `json:"activeAccessPolicies"`
	ActiveToolGrants int `json:"activeToolGrants"`
	ActiveSkillBindings int `json:"activeSkillBindings"`
	BudgetPolicies int `json:"budgetPolicies"`
	RateLimitPolicies int `json:"rateLimitPolicies"`
	RedactionPolicies int `json:"redactionPolicies"`
	ResourceScopedAccessPolicies int `json:"resourceScopedAccessPolicies"`
	ResourceScopedToolGrants int `json:"resourceScopedToolGrants"`
	BudgetState string `json:"budgetState"`
	RateLimitState string `json:"rateLimitState"`
	RedactionPolicyState string `json:"redactionPolicyState"`
	ResourceScopeState string `json:"resourceScopeState"`
}

type GovernanceFinding struct {
	Type string `json:"type"`
	Severity string `json:"severity"`
	Summary string `json:"summary"`
	Count int `json:"count,omitempty"`
	ActorType string `json:"actorType,omitempty"`
	ActorID string `json:"actorId,omitempty"`
	SubjectType string `json:"subjectType,omitempty"`
	SubjectID string `json:"subjectId,omitempty"`
	AIClientID string `json:"aiClientId,omitempty"`
	PolicyID string `json:"policyId,omitempty"`
	ApprovalRequestID string `json:"approvalRequestId,omitempty"`
	GrantID string `json:"grantId,omitempty"`
	ToolName string `json:"toolName,omitempty"`
	RiskLevel string `json:"riskLevel,omitempty"`
}

type GovernanceRecommendationAction struct {
	Type string `json:"type"`
	Severity string `json:"severity"`
	Summary string `json:"summary"`
	Action string `json:"action"`
	TargetKind string `json:"targetKind,omitempty"`
	TargetID string `json:"targetId,omitempty"`
	Refs []string `json:"refs,omitempty"`
	Count int `json:"count,omitempty"`
	Metadata map[string]any `json:"metadata,omitempty"`
}

type GovernanceStatusEnvelope struct {
	Data GovernanceStatus `json:"data"`
}

type MCPCapability struct {
	AdapterID string `json:"adapterID"`
	Name string `json:"name"`
	Description string `json:"description"`
	Scopes []string `json:"scopes"`
}

type MCPCapabilityListEnvelope struct {
	Items []MCPCapability `json:"items"`
}
