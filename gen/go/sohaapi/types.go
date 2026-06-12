// Package sohaapi provides primitives to interact with the openapi HTTP API.
//
// Code generated from OpenSoha contracts by github.com/oapi-codegen/oapi-codegen/v2 version v2.7.1 DO NOT EDIT.
package sohaapi

import (
	"encoding/json"
	"fmt"
	"time"
)

const (
	BearerAuthScopes        bearerAuthContextKey        = "bearerAuth.Scopes"
	RuntimeBearerAuthScopes runtimeBearerAuthContextKey = "runtimeBearerAuth.Scopes"
)

// Defines values for ClusterCapabilityStatus.
const (
	Available   ClusterCapabilityStatus = "available"
	Partial     ClusterCapabilityStatus = "partial"
	Unsupported ClusterCapabilityStatus = "unsupported"
)

// Valid indicates whether the value is a known member of the ClusterCapabilityStatus enum.
func (e ClusterCapabilityStatus) Valid() bool {
	switch e {
	case Available:
		return true
	case Partial:
		return true
	case Unsupported:
		return true
	default:
		return false
	}
}

// Defines values for DecideAIGatewayApprovalRequestParamsAction.
const (
	Approve DecideAIGatewayApprovalRequestParamsAction = "approve"
	Cancel  DecideAIGatewayApprovalRequestParamsAction = "cancel"
	Reject  DecideAIGatewayApprovalRequestParamsAction = "reject"
)

// Valid indicates whether the value is a known member of the DecideAIGatewayApprovalRequestParamsAction enum.
func (e DecideAIGatewayApprovalRequestParamsAction) Valid() bool {
	switch e {
	case Approve:
		return true
	case Cancel:
		return true
	case Reject:
		return true
	default:
		return false
	}
}

// AIGatewayManifest defines model for AIGatewayManifest.
type AIGatewayManifest struct {
	Caller         CallerContext        `json:"caller"`
	GeneratedAt    time.Time            `json:"generatedAt"`
	Name           string               `json:"name"`
	PermissionKeys []string             `json:"permissionKeys"`
	Principal      Principal            `json:"principal"`
	Prompts        []PromptCapability   `json:"prompts,omitempty"`
	Resources      []ResourceCapability `json:"resources,omitempty"`
	Skills         []SkillCapability    `json:"skills,omitempty"`
	Summary        struct {
		DeniedCount   int `json:"deniedCount"`
		PromptCount   int `json:"promptCount"`
		ResourceCount int `json:"resourceCount"`
		SkillCount    int `json:"skillCount"`
		ToolCount     int `json:"toolCount"`
	} `json:"summary"`
	Tools   []ToolCapability `json:"tools"`
	Version string           `json:"version"`
}

// AIGatewayManifestEnvelope defines model for AIGatewayManifestEnvelope.
type AIGatewayManifestEnvelope struct {
	Data AIGatewayManifest `json:"data"`
}

// AgentRun defines model for AgentRun.
type AgentRun struct {
	AnalysisArtifacts    []map[string]any `json:"analysisArtifacts,omitempty"`
	CallbackToken        string           `json:"callbackToken"`
	CapabilityID         string           `json:"capabilityId"`
	ID                   string           `json:"id"`
	Input                map[string]any   `json:"input,omitempty"`
	Output               map[string]any   `json:"output,omitempty"`
	ProviderID           string           `json:"providerId"`
	ProviderKind         string           `json:"providerKind"`
	Scope                map[string]any   `json:"scope,omitempty"`
	SessionID            string           `json:"sessionId,omitempty"`
	SkillBindings        []map[string]any `json:"skillBindings,omitempty"`
	SkillIDs             []string         `json:"skillIds,omitempty"`
	Status               string           `json:"status"`
	TimeoutSeconds       int              `json:"timeoutSeconds"`
	ToolBindings         []map[string]any `json:"toolBindings,omitempty"`
	Toolset              map[string]any   `json:"toolset,omitempty"`
	AdditionalProperties map[string]any   `json:"-"`
}

// AgentRunCallbackRequest defines model for AgentRunCallbackRequest.
type AgentRunCallbackRequest struct {
	AgentID           string           `json:"agentId"`
	AnalysisArtifacts []map[string]any `json:"analysisArtifacts,omitempty"`
	CallbackToken     string           `json:"callbackToken"`
	ErrorMessage      string           `json:"errorMessage,omitempty"`
	ExternalRunID     string           `json:"externalRunId,omitempty"`
	Payload           map[string]any   `json:"payload"`
	RunID             string           `json:"runId"`
	Status            string           `json:"status"`
	ToolExecutions    []map[string]any `json:"toolExecutions,omitempty"`
}

// AgentRunClaimRequest defines model for AgentRunClaimRequest.
type AgentRunClaimRequest struct {
	AgentID     string   `json:"agentId"`
	Kinds       []string `json:"kinds,omitempty"`
	ProviderIDs []string `json:"providerIds,omitempty"`
}

// AgentRunEnvelope defines model for AgentRunEnvelope.
type AgentRunEnvelope struct {
	Data AgentRun `json:"data"`
}

// AgentRunToolCallRequest defines model for AgentRunToolCallRequest.
type AgentRunToolCallRequest struct {
	AdapterID     string         `json:"adapterId,omitempty"`
	AgentID       string         `json:"agentId"`
	CallbackToken string         `json:"callbackToken"`
	Input         map[string]any `json:"input,omitempty"`
	RunID         string         `json:"runId"`
	ToolBindingID string         `json:"toolBindingId,omitempty"`
	ToolName      string         `json:"toolName,omitempty"`
}

// AgentToolCallResult defines model for AgentToolCallResult.
type AgentToolCallResult struct {
	Output               map[string]any `json:"output,omitempty"`
	RunID                string         `json:"runId"`
	ToolExecution        map[string]any `json:"toolExecution"`
	AdditionalProperties map[string]any `json:"-"`
}

// AgentToolCallResultEnvelope defines model for AgentToolCallResultEnvelope.
type AgentToolCallResultEnvelope struct {
	Data AgentToolCallResult `json:"data"`
}

// AnyValue defines model for AnyValue.
type AnyValue = any

// ApprovalDecisionInput defines model for ApprovalDecisionInput.
type ApprovalDecisionInput struct {
	Comment string `json:"comment,omitempty"`
}

// ApprovalDecisionResult defines model for ApprovalDecisionResult.
type ApprovalDecisionResult struct {
	Invocation *ToolInvocationResult `json:"invocation,omitempty"`
	Request    ApprovalRequest       `json:"request"`
}

// ApprovalDecisionResultEnvelope defines model for ApprovalDecisionResultEnvelope.
type ApprovalDecisionResultEnvelope struct {
	Data ApprovalDecisionResult `json:"data"`
}

// ApprovalDecisionTrace defines model for ApprovalDecisionTrace.
type ApprovalDecisionTrace struct {
	Comment    string     `json:"comment,omitempty"`
	DecidedAt  *time.Time `json:"decidedAt,omitempty"`
	Result     string     `json:"result,omitempty"`
	Roles      []string   `json:"roles,omitempty"`
	StageIndex int        `json:"stageIndex,omitempty"`
	StageName  string     `json:"stageName,omitempty"`
	Teams      []string   `json:"teams,omitempty"`
	UserID     string     `json:"userId,omitempty"`
	UserName   string     `json:"userName,omitempty"`
}

// ApprovalRequest defines model for ApprovalRequest.
type ApprovalRequest struct {
	ActorID           string         `json:"actorId"`
	ActorName         string         `json:"actorName,omitempty"`
	ActorRoles        []string       `json:"actorRoles,omitempty"`
	ActorTeams        []string       `json:"actorTeams,omitempty"`
	ActorType         string         `json:"actorType"`
	AIClientID        string         `json:"aiClientId,omitempty"`
	AIClientName      string         `json:"aiClientName,omitempty"`
	ApprovalPolicyRef string         `json:"approvalPolicyRef,omitempty"`
	ApprovalTrace     *ApprovalTrace `json:"approvalTrace,omitempty"`
	CreatedAt         time.Time      `json:"createdAt"`
	DecidedAt         *time.Time     `json:"decidedAt,omitempty"`
	DecidedBy         string         `json:"decidedBy,omitempty"`
	DecidedByName     string         `json:"decidedByName,omitempty"`
	DecisionComment   string         `json:"decisionComment,omitempty"`
	ExpiresAt         *time.Time     `json:"expiresAt,omitempty"`
	ID                string         `json:"id"`
	Output            AnyValue       `json:"output,omitempty"`
	PolicyID          string         `json:"policyId,omitempty"`
	RelatedIDs        map[string]any `json:"relatedIds,omitempty"`
	RequestID         string         `json:"requestId,omitempty"`
	RequiresApproval  bool           `json:"requiresApproval"`
	ResourceScope     map[string]any `json:"resourceScope,omitempty"`
	RiskLevel         RiskLevel      `json:"riskLevel"`
	SkillID           string         `json:"skillId,omitempty"`
	SourceIP          string         `json:"sourceIp,omitempty"`
	Status            string         `json:"status"`
	Strategy          string         `json:"strategy"`
	Summary           string         `json:"summary"`
	ToolInput         map[string]any `json:"toolInput,omitempty"`
	ToolName          string         `json:"toolName"`
	UpdatedAt         time.Time      `json:"updatedAt"`
}

// ApprovalRequestListEnvelope defines model for ApprovalRequestListEnvelope.
type ApprovalRequestListEnvelope struct {
	Items []ApprovalRequest `json:"items"`
}

// ApprovalStageTrace defines model for ApprovalStageTrace.
type ApprovalStageTrace struct {
	CompletedAt *time.Time `json:"completedAt,omitempty"`
	Result      string     `json:"result,omitempty"`
	StageIndex  int        `json:"stageIndex,omitempty"`
	StageName   string     `json:"stageName,omitempty"`
}

// ApprovalTimeline defines model for ApprovalTimeline.
type ApprovalTimeline struct {
	Events  []ApprovalTimelineEvent `json:"events,omitempty"`
	Request ApprovalRequest         `json:"request"`
	Trace   *ApprovalTrace          `json:"trace,omitempty"`
}

// ApprovalTimelineEnvelope defines model for ApprovalTimelineEnvelope.
type ApprovalTimelineEnvelope struct {
	Data ApprovalTimeline `json:"data"`
}

// ApprovalTimelineEvent defines model for ApprovalTimelineEvent.
type ApprovalTimelineEvent struct {
	Action     string         `json:"action"`
	ActorID    string         `json:"actorId,omitempty"`
	ActorName  string         `json:"actorName,omitempty"`
	ActorType  string         `json:"actorType,omitempty"`
	CreatedAt  time.Time      `json:"createdAt"`
	ID         string         `json:"id"`
	Kind       string         `json:"kind"`
	Metadata   map[string]any `json:"metadata,omitempty"`
	Result     string         `json:"result"`
	StageIndex int            `json:"stageIndex,omitempty"`
	StageName  string         `json:"stageName,omitempty"`
	Summary    string         `json:"summary,omitempty"`
}

// ApprovalTrace defines model for ApprovalTrace.
type ApprovalTrace struct {
	ApprovalMode           string                  `json:"approvalMode,omitempty"`
	ApprovedCount          int                     `json:"approvedCount,omitempty"`
	CandidateRoles         []string                `json:"candidateRoles,omitempty"`
	CandidateTeams         []string                `json:"candidateTeams,omitempty"`
	CandidateUserIDs       []string                `json:"candidateUserIds,omitempty"`
	CurrentStageIndex      int                     `json:"currentStageIndex,omitempty"`
	CurrentStageName       string                  `json:"currentStageName,omitempty"`
	Decisions              []ApprovalDecisionTrace `json:"decisions,omitempty"`
	ExecutionTaskID        string                  `json:"executionTaskId,omitempty"`
	OnCallCandidateUserIDs []string                `json:"onCallCandidateUserIds,omitempty"`
	PendingRequirements    []string                `json:"pendingRequirements,omitempty"`
	ReleaseBundleID        string                  `json:"releaseBundleId,omitempty"`
	RequiredApprovals      int                     `json:"requiredApprovals,omitempty"`
	RoleApprovedCounts     map[string]int          `json:"roleApprovedCounts,omitempty"`
	SatisfiedRequirements  []string                `json:"satisfiedRequirements,omitempty"`
	StageCount             int                     `json:"stageCount,omitempty"`
	StageHistory           []ApprovalStageTrace    `json:"stageHistory,omitempty"`
	TeamApprovedCounts     map[string]int          `json:"teamApprovedCounts,omitempty"`
	WorkflowRunID          string                  `json:"workflowRunId,omitempty"`
}

// AuditLog defines model for AuditLog.
type AuditLog struct {
	Action        string         `json:"action"`
	ActorID       string         `json:"actorId"`
	ActorName     string         `json:"actorName,omitempty"`
	ActorType     string         `json:"actorType"`
	AIClientID    string         `json:"aiClientId,omitempty"`
	AIClientName  string         `json:"aiClientName,omitempty"`
	CreatedAt     time.Time      `json:"createdAt"`
	ID            string         `json:"id"`
	Metadata      map[string]any `json:"metadata,omitempty"`
	RequestID     string         `json:"requestId,omitempty"`
	ResourceScope map[string]any `json:"resourceScope,omitempty"`
	Result        string         `json:"result"`
	RiskLevel     RiskLevel      `json:"riskLevel,omitempty"`
	SkillID       string         `json:"skillId,omitempty"`
	SourceIP      string         `json:"sourceIp,omitempty"`
	Summary       string         `json:"summary"`
	ToolName      string         `json:"toolName,omitempty"`
}

// AuditLogListEnvelope defines model for AuditLogListEnvelope.
type AuditLogListEnvelope struct {
	Items []AuditLog `json:"items"`
}

// AuthBootstrapEnvelope defines model for AuthBootstrapEnvelope.
type AuthBootstrapEnvelope struct {
	Data AuthBootstrapEnvelope_Data `json:"data"`
}

// AuthBootstrapEnvelope_Data defines model for AuthBootstrapEnvelope.Data.
type AuthBootstrapEnvelope_Data struct {
	Branding             map[string]any `json:"branding"`
	CurrentUser          Principal      `json:"currentUser"`
	PermissionSnapshot   map[string]any `json:"permissionSnapshot"`
	User                 Principal      `json:"user"`
	AdditionalProperties map[string]any `json:"-"`
}

// AuthProvider defines model for AuthProvider.
type AuthProvider struct {
	Enabled  bool   `json:"enabled"`
	ID       string `json:"id,omitempty"`
	LoginURL string `json:"loginUrl,omitempty"`
	Name     string `json:"name"`
	Type     string `json:"type"`
}

// AuthProviderListEnvelope defines model for AuthProviderListEnvelope.
type AuthProviderListEnvelope struct {
	Items []AuthProvider `json:"items"`
}

// AuthResult defines model for AuthResult.
type AuthResult struct {
	Tokens TokenSet  `json:"tokens"`
	User   Principal `json:"user"`
}

// AuthResultEnvelope defines model for AuthResultEnvelope.
type AuthResultEnvelope struct {
	Data AuthResult `json:"data"`
}

// CallerContext defines model for CallerContext.
type CallerContext struct {
	AIClientID   string `json:"aiClientId,omitempty"`
	AIClientName string `json:"aiClientName,omitempty"`
	IdentityMode string `json:"identityMode"`
	SessionID    string `json:"sessionId,omitempty"`
	SkillID      string `json:"skillId,omitempty"`
	Source       string `json:"source,omitempty"`
	SubjectID    string `json:"subjectId,omitempty"`
	SubjectType  string `json:"subjectType,omitempty"`
	TokenID      string `json:"tokenId,omitempty"`
}

// ClusterCapabilityMatrixEntry defines model for ClusterCapabilityMatrixEntry.
type ClusterCapabilityMatrixEntry struct {
	Agent            ClusterCapabilityModeSupport `json:"agent"`
	Category         string                       `json:"category"`
	Direct           ClusterCapabilityModeSupport `json:"direct"`
	DocsURL          string                       `json:"docsUrl,omitempty"`
	Key              string                       `json:"key"`
	Label            string                       `json:"label"`
	RequiredScopes   []string                     `json:"requiredScopes,omitempty"`
	RequiresApproval bool                         `json:"requiresApproval"`
	RiskLevel        RiskLevel                    `json:"riskLevel"`
}

// ClusterCapabilityMatrixEnvelope defines model for ClusterCapabilityMatrixEnvelope.
type ClusterCapabilityMatrixEnvelope struct {
	Items []ClusterCapabilityMatrixEntry `json:"items"`
}

// ClusterCapabilityModeSupport defines model for ClusterCapabilityModeSupport.
type ClusterCapabilityModeSupport struct {
	Notes  []string                `json:"notes,omitempty"`
	Reason string                  `json:"reason,omitempty"`
	Status ClusterCapabilityStatus `json:"status"`
}

// ClusterCapabilityStatus defines model for ClusterCapabilityStatus.
type ClusterCapabilityStatus string

// CreatedPersonalAccessToken defines model for CreatedPersonalAccessToken.
type CreatedPersonalAccessToken struct {
	Token PersonalAccessToken `json:"token"`
	Value string              `json:"value"`
}

// CreatedPersonalAccessTokenEnvelope defines model for CreatedPersonalAccessTokenEnvelope.
type CreatedPersonalAccessTokenEnvelope struct {
	Data CreatedPersonalAccessToken `json:"data"`
}

// CreatedServiceAccountToken defines model for CreatedServiceAccountToken.
type CreatedServiceAccountToken struct {
	Token ServiceAccountToken `json:"token"`
	Value string              `json:"value"`
}

// CreatedServiceAccountTokenEnvelope defines model for CreatedServiceAccountTokenEnvelope.
type CreatedServiceAccountTokenEnvelope struct {
	Data CreatedServiceAccountToken `json:"data"`
}

// DockerOperation defines model for DockerOperation.
type DockerOperation struct {
	ClaimedByWorkerID    string         `json:"claimedByWorkerId,omitempty"`
	HostID               string         `json:"hostId,omitempty"`
	ID                   string         `json:"id"`
	OperationKind        string         `json:"operationKind"`
	Payload              map[string]any `json:"payload"`
	ProjectID            string         `json:"projectId,omitempty"`
	ServiceID            string         `json:"serviceId,omitempty"`
	Status               string         `json:"status"`
	TimeoutSeconds       int            `json:"timeoutSeconds"`
	AdditionalProperties map[string]any `json:"-"`
}

// DockerOperationCallbackRequest defines model for DockerOperationCallbackRequest.
type DockerOperationCallbackRequest struct {
	Logs        []string       `json:"logs"`
	OperationID string         `json:"operationId"`
	Payload     map[string]any `json:"payload"`
	Status      string         `json:"status"`
	WorkerID    string         `json:"workerId"`
}

// DockerOperationClaimRequest defines model for DockerOperationClaimRequest.
type DockerOperationClaimRequest struct {
	AgentID        string   `json:"agentId"`
	HostIDs        []string `json:"hostIds,omitempty"`
	OperationKinds []string `json:"operationKinds,omitempty"`
	WorkerID       string   `json:"workerId"`
}

// DockerOperationEnvelope defines model for DockerOperationEnvelope.
type DockerOperationEnvelope struct {
	Data DockerOperation `json:"data"`
}

// ErrorEnvelope defines model for ErrorEnvelope.
type ErrorEnvelope struct {
	Error struct {
		Code      string `json:"code"`
		Message   string `json:"message"`
		RequestID string `json:"request_id,omitempty"`
	} `json:"error"`
}

// ExecutionCallbackRequest defines model for ExecutionCallbackRequest.
type ExecutionCallbackRequest struct {
	CallbackToken string         `json:"callbackToken"`
	Payload       map[string]any `json:"payload"`
	Status        string         `json:"status"`
}

// ExecutionTask defines model for ExecutionTask.
type ExecutionTask struct {
	ApplicationEnvironmentID string         `json:"applicationEnvironmentId"`
	ApplicationID            string         `json:"applicationId"`
	CallbackToken            string         `json:"callbackToken"`
	ID                       string         `json:"id"`
	Payload                  map[string]any `json:"payload"`
	ProviderKind             string         `json:"providerKind"`
	Status                   string         `json:"status"`
	TaskKind                 string         `json:"taskKind"`
	AdditionalProperties     map[string]any `json:"-"`
}

// ExecutionTaskClaimRequest defines model for ExecutionTaskClaimRequest.
type ExecutionTaskClaimRequest struct {
	AgentID         string   `json:"agentId"`
	ProviderKinds   []string `json:"providerKinds,omitempty"`
	RuntimeEndpoint string   `json:"runtimeEndpoint,omitempty"`
}

// ExecutionTaskEnvelope defines model for ExecutionTaskEnvelope.
type ExecutionTaskEnvelope struct {
	Data ExecutionTask `json:"data"`
}

// GenericItemsEnvelope defines model for GenericItemsEnvelope.
type GenericItemsEnvelope struct {
	Items []AnyValue `json:"items"`
}

// GenericObject defines model for GenericObject.
type GenericObject map[string]any

// GovernanceApprovalSummary defines model for GovernanceApprovalSummary.
type GovernanceApprovalSummary struct {
	DueSoon                int        `json:"dueSoon"`
	DueSoonRequestIDs      []string   `json:"dueSoonRequestIds,omitempty"`
	NextDueAt              *time.Time `json:"nextDueAt,omitempty"`
	NextDueRequestID       string     `json:"nextDueRequestId,omitempty"`
	OldestPendingHours     int        `json:"oldestPendingHours,omitempty"`
	OldestPendingRequestID string     `json:"oldestPendingRequestId,omitempty"`
	Overdue                int        `json:"overdue"`
	OverdueRequestIDs      []string   `json:"overdueRequestIds,omitempty"`
	Pending                int        `json:"pending"`
	StalePending           int        `json:"stalePending"`
	StalePendingRequestIDs []string   `json:"stalePendingRequestIds,omitempty"`
}

// GovernanceClientSummary defines model for GovernanceClientSummary.
type GovernanceClientSummary struct {
	Active                   int      `json:"active"`
	Disabled                 int      `json:"disabled"`
	PendingApproval          int      `json:"pendingApproval"`
	PendingApprovalClientIDs []string `json:"pendingApprovalClientIds,omitempty"`
	RegistrationApproval     string   `json:"registrationApproval"`
	Total                    int      `json:"total"`
}

// GovernanceFinding defines model for GovernanceFinding.
type GovernanceFinding struct {
	ActorID           string    `json:"actorId,omitempty"`
	ActorType         string    `json:"actorType,omitempty"`
	AIClientID        string    `json:"aiClientId,omitempty"`
	ApprovalRequestID string    `json:"approvalRequestId,omitempty"`
	Count             int       `json:"count,omitempty"`
	GrantID           string    `json:"grantId,omitempty"`
	PolicyID          string    `json:"policyId,omitempty"`
	RiskLevel         RiskLevel `json:"riskLevel,omitempty"`
	Severity          string    `json:"severity"`
	SubjectID         string    `json:"subjectId,omitempty"`
	SubjectType       string    `json:"subjectType,omitempty"`
	Summary           string    `json:"summary"`
	ToolName          string    `json:"toolName,omitempty"`
	Type              string    `json:"type"`
}

// GovernanceHealth defines model for GovernanceHealth.
type GovernanceHealth struct {
	Checks  []GovernanceHealthCheck `json:"checks,omitempty"`
	Message string                  `json:"message"`
	Status  string                  `json:"status"`
}

// GovernanceHealthCheck defines model for GovernanceHealthCheck.
type GovernanceHealthCheck struct {
	Count   int    `json:"count,omitempty"`
	Message string `json:"message"`
	Name    string `json:"name"`
	Status  string `json:"status"`
}

// GovernanceMetricCount defines model for GovernanceMetricCount.
type GovernanceMetricCount struct {
	Count int    `json:"count"`
	Key   string `json:"key"`
}

// GovernanceMetrics defines model for GovernanceMetrics.
type GovernanceMetrics struct {
	DenyCount             int                     `json:"denyCount"`
	DryRunCount           int                     `json:"dryRunCount"`
	FailureCount          int                     `json:"failureCount"`
	PendingApprovalCount  int                     `json:"pendingApprovalCount"`
	RecentActionBreakdown map[string]int          `json:"recentActionBreakdown,omitempty"`
	RecentResultBreakdown map[string]int          `json:"recentResultBreakdown,omitempty"`
	RiskCounts            map[string]int          `json:"riskCounts,omitempty"`
	SuccessCount          int                     `json:"successCount"`
	TopActors             []GovernanceMetricCount `json:"topActors,omitempty"`
	TopAIClients          []GovernanceMetricCount `json:"topAiClients,omitempty"`
	TopTools              []GovernanceMetricCount `json:"topTools,omitempty"`
	TotalCalls            int                     `json:"totalCalls"`
}

// GovernancePolicyCoverage defines model for GovernancePolicyCoverage.
type GovernancePolicyCoverage struct {
	AccessPolicies               int    `json:"accessPolicies"`
	ActiveAccessPolicies         int    `json:"activeAccessPolicies"`
	ActiveSkillBindings          int    `json:"activeSkillBindings"`
	ActiveToolGrants             int    `json:"activeToolGrants"`
	BudgetPolicies               int    `json:"budgetPolicies"`
	BudgetState                  string `json:"budgetState"`
	RateLimitPolicies            int    `json:"rateLimitPolicies"`
	RateLimitState               string `json:"rateLimitState"`
	RedactionPolicies            int    `json:"redactionPolicies"`
	RedactionPolicyState         string `json:"redactionPolicyState"`
	ResourceScopeState           string `json:"resourceScopeState"`
	ResourceScopedAccessPolicies int    `json:"resourceScopedAccessPolicies"`
	ResourceScopedToolGrants     int    `json:"resourceScopedToolGrants"`
	SkillBindings                int    `json:"skillBindings"`
	ToolGrants                   int    `json:"toolGrants"`
}

// GovernanceRecommendationAction defines model for GovernanceRecommendationAction.
type GovernanceRecommendationAction struct {
	Action     string         `json:"action"`
	Count      int            `json:"count,omitempty"`
	Metadata   map[string]any `json:"metadata,omitempty"`
	Refs       []string       `json:"refs,omitempty"`
	Severity   string         `json:"severity"`
	Summary    string         `json:"summary"`
	TargetID   string         `json:"targetId,omitempty"`
	TargetKind string         `json:"targetKind,omitempty"`
	Type       string         `json:"type"`
}

// GovernanceRedactionSummary defines model for GovernanceRedactionSummary.
type GovernanceRedactionSummary struct {
	AuditsWithRedaction     int                     `json:"auditsWithRedaction"`
	FieldMatches            int                     `json:"fieldMatches"`
	InputAudits             int                     `json:"inputAudits"`
	OutputAudits            int                     `json:"outputAudits"`
	SecretClassifierMatches int                     `json:"secretClassifierMatches"`
	SensitiveKeyMatches     int                     `json:"sensitiveKeyMatches"`
	SensitiveTextMatches    int                     `json:"sensitiveTextMatches"`
	StructuredSecretMatches int                     `json:"structuredSecretMatches"`
	TopClassifiers          []GovernanceMetricCount `json:"topClassifiers,omitempty"`
	TopFieldPaths           []GovernanceMetricCount `json:"topFieldPaths,omitempty"`
	TopMatchTypes           []GovernanceMetricCount `json:"topMatchTypes,omitempty"`
	TopPolicies             []GovernanceMetricCount `json:"topPolicies,omitempty"`
	TopTargets              []GovernanceMetricCount `json:"topTargets,omitempty"`
	TopTools                []GovernanceMetricCount `json:"topTools,omitempty"`
	TotalMatches            int                     `json:"totalMatches"`
	ValuePatternMatches     int                     `json:"valuePatternMatches"`
}

// GovernanceStatus defines model for GovernanceStatus.
type GovernanceStatus struct {
	Anomalies             []GovernanceFinding              `json:"anomalies,omitempty"`
	Approvals             GovernanceApprovalSummary        `json:"approvals"`
	Clients               GovernanceClientSummary          `json:"clients"`
	GeneratedAt           time.Time                        `json:"generatedAt"`
	Health                GovernanceHealth                 `json:"health"`
	Metadata              map[string]any                   `json:"metadata,omitempty"`
	Metrics               GovernanceMetrics                `json:"metrics"`
	PolicyCoverage        GovernancePolicyCoverage         `json:"policyCoverage"`
	RecommendationActions []GovernanceRecommendationAction `json:"recommendationActions,omitempty"`
	Recommendations       []string                         `json:"recommendations,omitempty"`
	Redaction             GovernanceRedactionSummary       `json:"redaction"`
	Tokens                GovernanceTokenSummary           `json:"tokens"`
	WindowHours           int                              `json:"windowHours"`
}

// GovernanceStatusEnvelope defines model for GovernanceStatusEnvelope.
type GovernanceStatusEnvelope struct {
	Data GovernanceStatus `json:"data"`
}

// GovernanceTokenCounts defines model for GovernanceTokenCounts.
type GovernanceTokenCounts struct {
	Active       int `json:"active"`
	Expired      int `json:"expired"`
	ExpiringSoon int `json:"expiringSoon"`
	NeverUsed    int `json:"neverUsed"`
	Revoked      int `json:"revoked"`
	Stale        int `json:"stale"`
	Total        int `json:"total"`
}

// GovernanceTokenFinding defines model for GovernanceTokenFinding.
type GovernanceTokenFinding struct {
	DaysUntilDue int        `json:"daysUntilDue,omitempty"`
	ExpiresAt    *time.Time `json:"expiresAt,omitempty"`
	ID           string     `json:"id"`
	Kind         string     `json:"kind"`
	LastUsedAt   *time.Time `json:"lastUsedAt,omitempty"`
	Message      string     `json:"message"`
	Name         string     `json:"name"`
	OwnerID      string     `json:"ownerId,omitempty"`
	Severity     string     `json:"severity"`
	StaleDays    int        `json:"staleDays,omitempty"`
	TokenPrefix  string     `json:"tokenPrefix"`
}

// GovernanceTokenSummary defines model for GovernanceTokenSummary.
type GovernanceTokenSummary struct {
	ExpiredActive         []GovernanceTokenFinding `json:"expiredActive,omitempty"`
	ExpiringSoon          []GovernanceTokenFinding `json:"expiringSoon,omitempty"`
	LastUsedTrackingState string                   `json:"lastUsedTrackingState"`
	NeverUsed             []GovernanceTokenFinding `json:"neverUsed,omitempty"`
	PersonalAccessTokens  GovernanceTokenCounts    `json:"personalAccessTokens"`
	ServiceAccountTokens  GovernanceTokenCounts    `json:"serviceAccountTokens"`
	Stale                 []GovernanceTokenFinding `json:"stale,omitempty"`
}

// InstalledPlugin defines model for InstalledPlugin.
type InstalledPlugin struct {
	ChecksumStatus       string                   `json:"checksumStatus"`
	ConfiguredSecretRefs map[string]string        `json:"configuredSecretRefs,omitempty"`
	DisabledAt           *time.Time               `json:"disabledAt,omitempty"`
	EnabledAt            *time.Time               `json:"enabledAt,omitempty"`
	ID                   string                   `json:"id"`
	InstalledAt          time.Time                `json:"installedAt"`
	InstalledBy          string                   `json:"installedBy"`
	Manifest             PluginManifest           `json:"manifest"`
	Metadata             map[string]any           `json:"metadata,omitempty"`
	Name                 string                   `json:"name"`
	Publisher            string                   `json:"publisher"`
	RequestedPermissions *PluginPermissionRequest `json:"requestedPermissions,omitempty"`
	SignatureStatus      string                   `json:"signatureStatus,omitempty"`
	Source               string                   `json:"source"`
	Status               string                   `json:"status"`
	Type                 string                   `json:"type"`
	UpdatedAt            time.Time                `json:"updatedAt"`
	Version              string                   `json:"version"`
}

// InstalledPluginEnvelope defines model for InstalledPluginEnvelope.
type InstalledPluginEnvelope struct {
	Data InstalledPlugin `json:"data"`
}

// InstalledPluginListEnvelope defines model for InstalledPluginListEnvelope.
type InstalledPluginListEnvelope struct {
	Items []InstalledPlugin `json:"items"`
}

// JSONSchema defines model for JSONSchema.
type JSONSchema map[string]any

// LoginOptions defines model for LoginOptions.
type LoginOptions struct {
	Verification struct {
		SliderEnabled bool `json:"sliderEnabled"`
	} `json:"verification"`
}

// LoginOptionsEnvelope defines model for LoginOptionsEnvelope.
type LoginOptionsEnvelope struct {
	Data LoginOptions `json:"data"`
}

// MCPCapability defines model for MCPCapability.
type MCPCapability struct {
	AdapterID   string   `json:"adapterID"`
	Description string   `json:"description"`
	Name        string   `json:"name"`
	Scopes      []string `json:"scopes"`
}

// MCPCapabilityListEnvelope defines model for MCPCapabilityListEnvelope.
type MCPCapabilityListEnvelope struct {
	Items []MCPCapability `json:"items"`
}

// MarketplacePlugin defines model for MarketplacePlugin.
type MarketplacePlugin struct {
	ID        string         `json:"id"`
	Installed bool           `json:"installed,omitempty"`
	Manifest  PluginManifest `json:"manifest"`
	Name      string         `json:"name"`
	Publisher string         `json:"publisher"`
	RiskLevel string         `json:"riskLevel,omitempty"`
	Source    string         `json:"source"`
	Summary   string         `json:"summary,omitempty"`
	Type      string         `json:"type"`
	Version   string         `json:"version"`
}

// MarketplacePluginEnvelope defines model for MarketplacePluginEnvelope.
type MarketplacePluginEnvelope struct {
	Data MarketplacePlugin `json:"data"`
}

// MarketplacePluginListEnvelope defines model for MarketplacePluginListEnvelope.
type MarketplacePluginListEnvelope struct {
	Items []MarketplacePlugin `json:"items"`
}

// OIDCExchangeRequest defines model for OIDCExchangeRequest.
type OIDCExchangeRequest struct {
	Code string `json:"code"`
}

// OperationStatus defines model for OperationStatus.
type OperationStatus struct {
	Status string `json:"status"`
}

// PasswordLoginRequest defines model for PasswordLoginRequest.
type PasswordLoginRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

// PersonalAccessToken defines model for PersonalAccessToken.
type PersonalAccessToken struct {
	CreatedAt      time.Time      `json:"createdAt"`
	CreatedBy      string         `json:"createdBy"`
	ExpiresAt      *time.Time     `json:"expiresAt,omitempty"`
	ID             string         `json:"id"`
	LastUsedAt     *time.Time     `json:"lastUsedAt,omitempty"`
	Metadata       map[string]any `json:"metadata,omitempty"`
	Name           string         `json:"name"`
	PermissionKeys []string       `json:"permissionKeys"`
	RevokedAt      *time.Time     `json:"revokedAt,omitempty"`
	Scopes         []string       `json:"scopes"`
	TokenPrefix    string         `json:"tokenPrefix"`
	UpdatedAt      time.Time      `json:"updatedAt"`
	UserID         string         `json:"userId"`
}

// PersonalAccessTokenInput defines model for PersonalAccessTokenInput.
type PersonalAccessTokenInput struct {
	ExpiresAt      *time.Time     `json:"expiresAt,omitempty"`
	Metadata       map[string]any `json:"metadata,omitempty"`
	Name           string         `json:"name"`
	PermissionKeys []string       `json:"permissionKeys"`
	Scopes         []string       `json:"scopes"`
}

// PersonalAccessTokenListEnvelope defines model for PersonalAccessTokenListEnvelope.
type PersonalAccessTokenListEnvelope struct {
	Items []PersonalAccessToken `json:"items"`
}

// PluginAssetSnapshot defines model for PluginAssetSnapshot.
type PluginAssetSnapshot struct {
	AgentProfiles      []string `json:"agentProfiles,omitempty"`
	AIProviderAdapters []string `json:"aiProviderAdapters,omitempty"`
	Connectors         []string `json:"connectors,omitempty"`
	GatewayPolicyPacks []string `json:"gatewayPolicyPacks,omitempty"`
	MCPPresets         []string `json:"mcpPresets,omitempty"`
	Skills             []string `json:"skills,omitempty"`
}

// PluginCapabilityRequest defines model for PluginCapabilityRequest.
type PluginCapabilityRequest struct {
	Prompts   []string `json:"prompts,omitempty"`
	Resources []string `json:"resources,omitempty"`
	Skills    []string `json:"skills,omitempty"`
	Tools     []string `json:"tools,omitempty"`
}

// PluginCompatibility defines model for PluginCompatibility.
type PluginCompatibility struct {
	Contracts            string         `json:"contracts,omitempty"`
	Soha                 string         `json:"soha,omitempty"`
	AdditionalProperties map[string]any `json:"-"`
}

// PluginConfigRequest defines model for PluginConfigRequest.
type PluginConfigRequest struct {
	Enabled    *bool             `json:"enabled,omitempty"`
	Metadata   map[string]any    `json:"metadata,omitempty"`
	SecretRefs map[string]string `json:"secretRefs,omitempty"`
}

// PluginInstallRequest defines model for PluginInstallRequest.
type PluginInstallRequest struct {
	Enable           bool            `json:"enable,omitempty"`
	ExpectedChecksum string          `json:"expectedChecksum,omitempty"`
	Manifest         *PluginManifest `json:"manifest,omitempty"`
	PluginID         string          `json:"pluginId,omitempty"`
	Source           string          `json:"source,omitempty"`
}

// PluginIntegrity defines model for PluginIntegrity.
type PluginIntegrity struct {
	Checksum  string `json:"checksum,omitempty"`
	Signature string `json:"signature,omitempty"`
	Status    string `json:"status,omitempty"`
	Verified  bool   `json:"verified,omitempty"`
}

// PluginManifest defines model for PluginManifest.
type PluginManifest struct {
	Assets        *PluginAssetSnapshot     `json:"assets,omitempty"`
	Capabilities  *PluginCapabilityRequest `json:"capabilities,omitempty"`
	Compatibility *PluginCompatibility     `json:"compatibility,omitempty"`
	Description   string                   `json:"description,omitempty"`
	Homepage      string                   `json:"homepage,omitempty"`
	ID            string                   `json:"id"`
	Integrity     *PluginIntegrity         `json:"integrity,omitempty"`
	Metadata      map[string]any           `json:"metadata,omitempty"`
	Name          string                   `json:"name"`
	Permissions   *PluginPermissionRequest `json:"permissions,omitempty"`
	Publisher     string                   `json:"publisher"`
	Secrets       *struct {
		Required []PluginSecretRequirement `json:"required,omitempty"`
	} `json:"secrets,omitempty"`
	Type    string `json:"type"`
	Version string `json:"version"`
}

// PluginManifestEnvelope defines model for PluginManifestEnvelope.
type PluginManifestEnvelope struct {
	Data PluginManifest `json:"data"`
}

// PluginPermissionRequest defines model for PluginPermissionRequest.
type PluginPermissionRequest struct {
	Domain   []string `json:"domain,omitempty"`
	Required []string `json:"required,omitempty"`
}

// PluginSecretRequirement defines model for PluginSecretRequirement.
type PluginSecretRequirement struct {
	Description string `json:"description,omitempty"`
	Name        string `json:"name"`
	Required    bool   `json:"required,omitempty"`
	SecretRef   string `json:"secretRef,omitempty"`
}

// Principal defines model for Principal.
type Principal struct {
	Email                string         `json:"email"`
	PermissionKeys       []string       `json:"permissionKeys,omitempty"`
	Projects             []string       `json:"projects"`
	Roles                []string       `json:"roles"`
	Tags                 []string       `json:"tags"`
	Teams                []string       `json:"teams"`
	UserID               string         `json:"userId"`
	UserName             string         `json:"userName"`
	AdditionalProperties map[string]any `json:"-"`
}

// PrincipalEnvelope defines model for PrincipalEnvelope.
type PrincipalEnvelope struct {
	Data Principal `json:"data"`
}

// PromptCapability defines model for PromptCapability.
type PromptCapability struct {
	ArgumentSchema JSONSchema `json:"argumentSchema,omitempty"`
	ContextSchema  JSONSchema `json:"contextSchema,omitempty"`
	Description    string     `json:"description"`
	Name           string     `json:"name"`
	PermissionKeys []string   `json:"permissionKeys"`
	RequiredScopes []string   `json:"requiredScopes,omitempty"`
}

// PromptGetRequest defines model for PromptGetRequest.
type PromptGetRequest struct {
	AIClientID   string         `json:"aiClientId,omitempty"`
	AIClientName string         `json:"aiClientName,omitempty"`
	Arguments    map[string]any `json:"arguments,omitempty"`
	Context      map[string]any `json:"context,omitempty"`
	Name         string         `json:"name"`
	RequestID    string         `json:"requestId,omitempty"`
	SkillID      string         `json:"skillId,omitempty"`
}

// PromptGetResult defines model for PromptGetResult.
type PromptGetResult struct {
	Audit                map[string]any  `json:"audit,omitempty"`
	Description          string          `json:"description"`
	Messages             []PromptMessage `json:"messages"`
	Name                 string          `json:"name"`
	RelatedIDs           map[string]any  `json:"relatedIds,omitempty"`
	AdditionalProperties map[string]any  `json:"-"`
}

// PromptGetResultEnvelope defines model for PromptGetResultEnvelope.
type PromptGetResultEnvelope struct {
	Data PromptGetResult `json:"data"`
}

// PromptMessage defines model for PromptMessage.
type PromptMessage struct {
	Content string `json:"content"`
	Role    string `json:"role"`
}

// RefreshRequest defines model for RefreshRequest.
type RefreshRequest struct {
	RefreshToken string `json:"refreshToken,omitempty"`
}

// ResourceCapability defines model for ResourceCapability.
type ResourceCapability struct {
	ContextSchema  JSONSchema `json:"contextSchema,omitempty"`
	Description    string     `json:"description"`
	Name           string     `json:"name"`
	PermissionKeys []string   `json:"permissionKeys"`
	RequiredScopes []string   `json:"requiredScopes,omitempty"`
}

// ResourceReadRequest defines model for ResourceReadRequest.
type ResourceReadRequest struct {
	AIClientID   string         `json:"aiClientId,omitempty"`
	AIClientName string         `json:"aiClientName,omitempty"`
	Context      map[string]any `json:"context,omitempty"`
	Name         string         `json:"name,omitempty"`
	RequestID    string         `json:"requestId,omitempty"`
	SkillID      string         `json:"skillId,omitempty"`
	URI          string         `json:"uri,omitempty"`
}

// ResourceReadResult defines model for ResourceReadResult.
type ResourceReadResult struct {
	Audit                map[string]any `json:"audit,omitempty"`
	Data                 AnyValue       `json:"data,omitempty"`
	MIMEType             string         `json:"mimeType"`
	Name                 string         `json:"name"`
	RelatedIDs           map[string]any `json:"relatedIds,omitempty"`
	Text                 string         `json:"text,omitempty"`
	URI                  string         `json:"uri"`
	AdditionalProperties map[string]any `json:"-"`
}

// ResourceReadResultEnvelope defines model for ResourceReadResultEnvelope.
type ResourceReadResultEnvelope struct {
	Data ResourceReadResult `json:"data"`
}

// RiskLevel defines model for RiskLevel.
type RiskLevel = string

// ServiceAccount defines model for ServiceAccount.
type ServiceAccount struct {
	CreatedAt     time.Time      `json:"createdAt"`
	CreatedBy     string         `json:"createdBy"`
	Description   string         `json:"description,omitempty"`
	ID            string         `json:"id"`
	Metadata      map[string]any `json:"metadata,omitempty"`
	Name          string         `json:"name"`
	OwnerUserID   string         `json:"ownerUserId,omitempty"`
	RoleIDs       []string       `json:"roleIds"`
	ScopeGrantIDs []string       `json:"scopeGrantIds"`
	Status        string         `json:"status"`
	TeamIDs       []string       `json:"teamIds"`
	UpdatedAt     time.Time      `json:"updatedAt"`
}

// ServiceAccountEnvelope defines model for ServiceAccountEnvelope.
type ServiceAccountEnvelope struct {
	Data ServiceAccount `json:"data"`
}

// ServiceAccountInput defines model for ServiceAccountInput.
type ServiceAccountInput struct {
	Description   string         `json:"description,omitempty"`
	ID            string         `json:"id"`
	Metadata      map[string]any `json:"metadata,omitempty"`
	Name          string         `json:"name"`
	OwnerUserID   string         `json:"ownerUserId,omitempty"`
	RoleIDs       []string       `json:"roleIds,omitempty"`
	ScopeGrantIDs []string       `json:"scopeGrantIds,omitempty"`
	Status        string         `json:"status"`
	TeamIDs       []string       `json:"teamIds,omitempty"`
}

// ServiceAccountListEnvelope defines model for ServiceAccountListEnvelope.
type ServiceAccountListEnvelope struct {
	Items []ServiceAccount `json:"items"`
}

// ServiceAccountToken defines model for ServiceAccountToken.
type ServiceAccountToken struct {
	CreatedAt        time.Time      `json:"createdAt"`
	CreatedBy        string         `json:"createdBy"`
	ExpiresAt        *time.Time     `json:"expiresAt,omitempty"`
	ID               string         `json:"id"`
	LastUsedAt       *time.Time     `json:"lastUsedAt,omitempty"`
	Metadata         map[string]any `json:"metadata,omitempty"`
	Name             string         `json:"name"`
	PermissionKeys   []string       `json:"permissionKeys"`
	RevokedAt        *time.Time     `json:"revokedAt,omitempty"`
	Scopes           []string       `json:"scopes"`
	ServiceAccountID string         `json:"serviceAccountId"`
	TokenPrefix      string         `json:"tokenPrefix"`
	UpdatedAt        time.Time      `json:"updatedAt"`
}

// ServiceAccountTokenInput defines model for ServiceAccountTokenInput.
type ServiceAccountTokenInput struct {
	ExpiresAt      *time.Time     `json:"expiresAt,omitempty"`
	Metadata       map[string]any `json:"metadata,omitempty"`
	Name           string         `json:"name"`
	PermissionKeys []string       `json:"permissionKeys"`
	Scopes         []string       `json:"scopes"`
}

// ServiceAccountTokenListEnvelope defines model for ServiceAccountTokenListEnvelope.
type ServiceAccountTokenListEnvelope struct {
	Items []ServiceAccountToken `json:"items"`
}

// SkillCapability defines model for SkillCapability.
type SkillCapability struct {
	CapabilityRefs []string `json:"capabilityRefs,omitempty"`
	Category       string   `json:"category"`
	Description    string   `json:"description"`
	ID             string   `json:"id"`
	Name           string   `json:"name"`
	PermissionKeys []string `json:"permissionKeys,omitempty"`
	RequiredScopes []string `json:"requiredScopes,omitempty"`
}

// StreamTicket defines model for StreamTicket.
type StreamTicket struct {
	ExpiresAt time.Time `json:"expiresAt"`
	Ticket    string    `json:"ticket"`
}

// StreamTicketEnvelope defines model for StreamTicketEnvelope.
type StreamTicketEnvelope struct {
	Data StreamTicket `json:"data"`
}

// StreamTicketRequest defines model for StreamTicketRequest.
type StreamTicketRequest struct {
	Path string `json:"path"`
}

// TokenRotationInput defines model for TokenRotationInput.
type TokenRotationInput struct {
	ExpiresAt *time.Time `json:"expiresAt,omitempty"`
}

// TokenSet defines model for TokenSet.
type TokenSet struct {
	AccessToken  string    `json:"accessToken"`
	ExpiresAt    time.Time `json:"expiresAt"`
	ExpiresIn    int64     `json:"expiresIn"`
	RefreshToken string    `json:"refreshToken"`
	TokenType    string    `json:"tokenType"`
}

// ToolCapability defines model for ToolCapability.
type ToolCapability struct {
	Action           string     `json:"action"`
	Description      string     `json:"description"`
	Domain           string     `json:"domain"`
	InputSchema      JSONSchema `json:"inputSchema,omitempty"`
	MCPAdapterID     string     `json:"mcpAdapterId,omitempty"`
	MCPToolName      string     `json:"mcpToolName,omitempty"`
	Name             string     `json:"name"`
	OutputSchema     JSONSchema `json:"outputSchema,omitempty"`
	PermissionKeys   []string   `json:"permissionKeys"`
	RequiredScopes   []string   `json:"requiredScopes,omitempty"`
	RequiresApproval bool       `json:"requiresApproval"`
	RiskLevel        RiskLevel  `json:"riskLevel"`
	Title            string     `json:"title"`
}

// ToolInvocationRequest defines model for ToolInvocationRequest.
type ToolInvocationRequest struct {
	AIClientID   string         `json:"aiClientId,omitempty"`
	AIClientName string         `json:"aiClientName,omitempty"`
	Input        map[string]any `json:"input,omitempty"`
	RequestID    string         `json:"requestId,omitempty"`
	SkillID      string         `json:"skillId,omitempty"`
	ToolName     string         `json:"toolName,omitempty"`
}

// ToolInvocationResult defines model for ToolInvocationResult.
type ToolInvocationResult struct {
	Audit                map[string]any `json:"audit,omitempty"`
	Output               AnyValue       `json:"output,omitempty"`
	RelatedIDs           map[string]any `json:"relatedIds,omitempty"`
	RequiresApproval     bool           `json:"requiresApproval"`
	Result               string         `json:"result"`
	RiskLevel            RiskLevel      `json:"riskLevel"`
	ToolName             string         `json:"toolName"`
	AdditionalProperties map[string]any `json:"-"`
}

// ToolInvocationResultEnvelope defines model for ToolInvocationResultEnvelope.
type ToolInvocationResultEnvelope struct {
	Data ToolInvocationResult `json:"data"`
}

// UserProfile defines model for UserProfile.
type UserProfile struct {
	DisplayName          string           `json:"displayName"`
	Email                string           `json:"email"`
	Identities           []map[string]any `json:"identities,omitempty"`
	LastLoginAt          *time.Time       `json:"lastLoginAt,omitempty"`
	Phone                string           `json:"phone,omitempty"`
	Projects             []string         `json:"projects"`
	Roles                []string         `json:"roles"`
	Sessions             []map[string]any `json:"sessions,omitempty"`
	Status               string           `json:"status"`
	Tags                 []string         `json:"tags"`
	Teams                []string         `json:"teams"`
	UserID               string           `json:"userId"`
	Username             string           `json:"username"`
	AdditionalProperties map[string]any   `json:"-"`
}

// UserProfileEnvelope defines model for UserProfileEnvelope.
type UserProfileEnvelope struct {
	Data UserProfile `json:"data"`
}

// AIClientID defines model for AIClientID.
type AIClientID = string

// AIClientName defines model for AIClientName.
type AIClientName = string

// OIDCCode defines model for OIDCCode.
type OIDCCode = string

// OIDCState defines model for OIDCState.
type OIDCState = string

// OperationID defines model for OperationID.
type OperationID = string

// PluginID defines model for PluginID.
type PluginID = string

// ProviderID defines model for ProviderID.
type ProviderID = string

// SkillID defines model for SkillID.
type SkillID = string

// Source defines model for Source.
type Source = string

// TaskID defines model for TaskID.
type TaskID = string

// ToolName defines model for ToolName.
type ToolName = string

// Error defines model for Error.
type Error = ErrorEnvelope

// bearerAuthContextKey is the context key for bearerAuth security scheme
type bearerAuthContextKey string

// runtimeBearerAuthContextKey is the context key for runtimeBearerAuth security scheme
type runtimeBearerAuthContextKey string

// ListAIGatewayApprovalRequestsParams defines parameters for ListAIGatewayApprovalRequests.
type ListAIGatewayApprovalRequestsParams struct {
	// ID Canonical approval request id query parameter. The server also accepts requestID and approvalRequestId.
	ID        string `form:"id,omitempty" json:"id,omitempty"`
	Status    string `form:"status,omitempty" json:"status,omitempty"`
	ActorType string `form:"actorType,omitempty" json:"actorType,omitempty"`

	// ActorID Canonical actor identifier query parameter. The server also accepts actor.
	ActorID    string    `form:"actorId,omitempty" json:"actorId,omitempty"`
	AIClientID string    `form:"aiClientId,omitempty" json:"aiClientId,omitempty"`
	SkillID    string    `form:"skillId,omitempty" json:"skillId,omitempty"`
	ToolName   string    `form:"toolName,omitempty" json:"toolName,omitempty"`
	RiskLevel  RiskLevel `form:"riskLevel,omitempty" json:"riskLevel,omitempty"`
	Strategy   string    `form:"strategy,omitempty" json:"strategy,omitempty"`
	From       time.Time `form:"from,omitempty" json:"from,omitempty"`
	To         time.Time `form:"to,omitempty" json:"to,omitempty"`
	Limit      int       `form:"limit,omitempty" json:"limit,omitempty"`
}

// DecideAIGatewayApprovalRequestParamsAction defines parameters for DecideAIGatewayApprovalRequest.
type DecideAIGatewayApprovalRequestParamsAction string

// ListAIGatewayAuditLogsParams defines parameters for ListAIGatewayAuditLogs.
type ListAIGatewayAuditLogsParams struct {
	ActorType string `form:"actorType,omitempty" json:"actorType,omitempty"`

	// ActorID Canonical actor identifier query parameter. The server also accepts actor.
	ActorID           string    `form:"actorId,omitempty" json:"actorId,omitempty"`
	AIClientID        string    `form:"aiClientId,omitempty" json:"aiClientId,omitempty"`
	SkillID           string    `form:"skillId,omitempty" json:"skillId,omitempty"`
	ToolName          string    `form:"toolName,omitempty" json:"toolName,omitempty"`
	ApprovalRequestID string    `form:"approvalRequestId,omitempty" json:"approvalRequestId,omitempty"`
	RiskLevel         RiskLevel `form:"riskLevel,omitempty" json:"riskLevel,omitempty"`
	Result            string    `form:"result,omitempty" json:"result,omitempty"`
	Action            string    `form:"action,omitempty" json:"action,omitempty"`
	From              time.Time `form:"from,omitempty" json:"from,omitempty"`
	To                time.Time `form:"to,omitempty" json:"to,omitempty"`
	Limit             int       `form:"limit,omitempty" json:"limit,omitempty"`
}

// GetAIGatewayCapabilitiesParams defines parameters for GetAIGatewayCapabilities.
type GetAIGatewayCapabilitiesParams struct {
	AIClientID   AIClientID   `form:"aiClientId,omitempty" json:"aiClientId,omitempty"`
	AIClientName AIClientName `form:"aiClientName,omitempty" json:"aiClientName,omitempty"`
	SkillID      SkillID      `form:"skillId,omitempty" json:"skillId,omitempty"`
	Source       Source       `form:"source,omitempty" json:"source,omitempty"`
}

// GetAIGatewayGovernanceStatusParams defines parameters for GetAIGatewayGovernanceStatus.
type GetAIGatewayGovernanceStatusParams struct {
	WindowHours int `form:"windowHours,omitempty" json:"windowHours,omitempty"`
}

// ListPersonalAccessTokensParams defines parameters for ListPersonalAccessTokens.
type ListPersonalAccessTokensParams struct {
	Scope  string `form:"scope,omitempty" json:"scope,omitempty"`
	UserID string `form:"userId,omitempty" json:"userId,omitempty"`
}

// HandleProviderCallbackParams defines parameters for HandleProviderCallback.
type HandleProviderCallbackParams struct {
	Code  OIDCCode  `form:"code,omitempty" json:"code,omitempty"`
	State OIDCState `form:"state,omitempty" json:"state,omitempty"`
}

// HandleOIDCCallbackParams defines parameters for HandleOIDCCallback.
type HandleOIDCCallbackParams struct {
	Code  OIDCCode  `form:"code,omitempty" json:"code,omitempty"`
	State OIDCState `form:"state,omitempty" json:"state,omitempty"`
}

// ListMarketplacePluginsParams defines parameters for ListMarketplacePlugins.
type ListMarketplacePluginsParams struct {
	Q         string `form:"q,omitempty" json:"q,omitempty"`
	Type      string `form:"type,omitempty" json:"type,omitempty"`
	Publisher string `form:"publisher,omitempty" json:"publisher,omitempty"`
}

// DecideAIGatewayApprovalRequestJSONRequestBody defines body for DecideAIGatewayApprovalRequest for application/json ContentType.
type DecideAIGatewayApprovalRequestJSONRequestBody = ApprovalDecisionInput

// CreatePersonalAccessTokenJSONRequestBody defines body for CreatePersonalAccessToken for application/json ContentType.
type CreatePersonalAccessTokenJSONRequestBody = PersonalAccessTokenInput

// RotatePersonalAccessTokenJSONRequestBody defines body for RotatePersonalAccessToken for application/json ContentType.
type RotatePersonalAccessTokenJSONRequestBody = TokenRotationInput

// GetAIGatewayPromptJSONRequestBody defines body for GetAIGatewayPrompt for application/json ContentType.
type GetAIGatewayPromptJSONRequestBody = PromptGetRequest

// ReadAIGatewayResourceJSONRequestBody defines body for ReadAIGatewayResource for application/json ContentType.
type ReadAIGatewayResourceJSONRequestBody = ResourceReadRequest

// RotateServiceAccountTokenJSONRequestBody defines body for RotateServiceAccountToken for application/json ContentType.
type RotateServiceAccountTokenJSONRequestBody = TokenRotationInput

// CreateServiceAccountJSONRequestBody defines body for CreateServiceAccount for application/json ContentType.
type CreateServiceAccountJSONRequestBody = ServiceAccountInput

// CreateServiceAccountTokenJSONRequestBody defines body for CreateServiceAccountToken for application/json ContentType.
type CreateServiceAccountTokenJSONRequestBody = ServiceAccountTokenInput

// InvokeAIGatewayToolJSONRequestBody defines body for InvokeAIGatewayTool for application/json ContentType.
type InvokeAIGatewayToolJSONRequestBody = ToolInvocationRequest

// PasswordLoginJSONRequestBody defines body for PasswordLogin for application/json ContentType.
type PasswordLoginJSONRequestBody = PasswordLoginRequest

// LogoutAuthSessionJSONRequestBody defines body for LogoutAuthSession for application/json ContentType.
type LogoutAuthSessionJSONRequestBody = RefreshRequest

// ExchangeOIDCCodeJSONRequestBody defines body for ExchangeOIDCCode for application/json ContentType.
type ExchangeOIDCCodeJSONRequestBody = OIDCExchangeRequest

// RefreshAuthSessionJSONRequestBody defines body for RefreshAuthSession for application/json ContentType.
type RefreshAuthSessionJSONRequestBody = RefreshRequest

// IssueStreamTicketJSONRequestBody defines body for IssueStreamTicket for application/json ContentType.
type IssueStreamTicketJSONRequestBody = StreamTicketRequest

// RecordAgentRunCallbackJSONRequestBody defines body for RecordAgentRunCallback for application/json ContentType.
type RecordAgentRunCallbackJSONRequestBody = AgentRunCallbackRequest

// ClaimAgentRunJSONRequestBody defines body for ClaimAgentRun for application/json ContentType.
type ClaimAgentRunJSONRequestBody = AgentRunClaimRequest

// RecordAgentRunToolCallJSONRequestBody defines body for RecordAgentRunToolCall for application/json ContentType.
type RecordAgentRunToolCallJSONRequestBody = AgentRunToolCallRequest

// RecordExecutionCallbackJSONRequestBody defines body for RecordExecutionCallback for application/json ContentType.
type RecordExecutionCallbackJSONRequestBody = ExecutionCallbackRequest

// ClaimExecutionTaskJSONRequestBody defines body for ClaimExecutionTask for application/json ContentType.
type ClaimExecutionTaskJSONRequestBody = ExecutionTaskClaimRequest

// RecordDockerOperationCallbackJSONRequestBody defines body for RecordDockerOperationCallback for application/json ContentType.
type RecordDockerOperationCallbackJSONRequestBody = DockerOperationCallbackRequest

// ClaimDockerOperationJSONRequestBody defines body for ClaimDockerOperation for application/json ContentType.
type ClaimDockerOperationJSONRequestBody = DockerOperationClaimRequest

// InstallPluginJSONRequestBody defines body for InstallPlugin for application/json ContentType.
type InstallPluginJSONRequestBody = PluginInstallRequest

// ConfigureInstalledPluginJSONRequestBody defines body for ConfigureInstalledPlugin for application/json ContentType.
type ConfigureInstalledPluginJSONRequestBody = PluginConfigRequest

// UpgradeInstalledPluginJSONRequestBody defines body for UpgradeInstalledPlugin for application/json ContentType.
type UpgradeInstalledPluginJSONRequestBody = PluginInstallRequest

// Getter for additional properties for AgentRun. Returns the specified
// element and whether it was found
func (a AgentRun) Get(fieldName string) (value interface{}, found bool) {
	if a.AdditionalProperties != nil {
		value, found = a.AdditionalProperties[fieldName]
	}
	return
}

// Setter for additional properties for AgentRun
func (a *AgentRun) Set(fieldName string, value interface{}) {
	if a.AdditionalProperties == nil {
		a.AdditionalProperties = make(map[string]any)
	}
	a.AdditionalProperties[fieldName] = value
}

// Override default JSON handling for AgentRun to handle AdditionalProperties
func (a *AgentRun) UnmarshalJSON(b []byte) error {
	object := make(map[string]json.RawMessage)
	err := json.Unmarshal(b, &object)
	if err != nil {
		return err
	}

	if raw, found := object["analysisArtifacts"]; found {
		err = json.Unmarshal(raw, &a.AnalysisArtifacts)
		if err != nil {
			return fmt.Errorf("error reading 'analysisArtifacts': %w", err)
		}
		delete(object, "analysisArtifacts")
	}

	if raw, found := object["callbackToken"]; found {
		err = json.Unmarshal(raw, &a.CallbackToken)
		if err != nil {
			return fmt.Errorf("error reading 'callbackToken': %w", err)
		}
		delete(object, "callbackToken")
	}

	if raw, found := object["capabilityId"]; found {
		err = json.Unmarshal(raw, &a.CapabilityID)
		if err != nil {
			return fmt.Errorf("error reading 'capabilityId': %w", err)
		}
		delete(object, "capabilityId")
	}

	if raw, found := object["id"]; found {
		err = json.Unmarshal(raw, &a.ID)
		if err != nil {
			return fmt.Errorf("error reading 'id': %w", err)
		}
		delete(object, "id")
	}

	if raw, found := object["input"]; found {
		err = json.Unmarshal(raw, &a.Input)
		if err != nil {
			return fmt.Errorf("error reading 'input': %w", err)
		}
		delete(object, "input")
	}

	if raw, found := object["output"]; found {
		err = json.Unmarshal(raw, &a.Output)
		if err != nil {
			return fmt.Errorf("error reading 'output': %w", err)
		}
		delete(object, "output")
	}

	if raw, found := object["providerId"]; found {
		err = json.Unmarshal(raw, &a.ProviderID)
		if err != nil {
			return fmt.Errorf("error reading 'providerId': %w", err)
		}
		delete(object, "providerId")
	}

	if raw, found := object["providerKind"]; found {
		err = json.Unmarshal(raw, &a.ProviderKind)
		if err != nil {
			return fmt.Errorf("error reading 'providerKind': %w", err)
		}
		delete(object, "providerKind")
	}

	if raw, found := object["scope"]; found {
		err = json.Unmarshal(raw, &a.Scope)
		if err != nil {
			return fmt.Errorf("error reading 'scope': %w", err)
		}
		delete(object, "scope")
	}

	if raw, found := object["sessionId"]; found {
		err = json.Unmarshal(raw, &a.SessionID)
		if err != nil {
			return fmt.Errorf("error reading 'sessionId': %w", err)
		}
		delete(object, "sessionId")
	}

	if raw, found := object["skillBindings"]; found {
		err = json.Unmarshal(raw, &a.SkillBindings)
		if err != nil {
			return fmt.Errorf("error reading 'skillBindings': %w", err)
		}
		delete(object, "skillBindings")
	}

	if raw, found := object["skillIds"]; found {
		err = json.Unmarshal(raw, &a.SkillIDs)
		if err != nil {
			return fmt.Errorf("error reading 'skillIds': %w", err)
		}
		delete(object, "skillIds")
	}

	if raw, found := object["status"]; found {
		err = json.Unmarshal(raw, &a.Status)
		if err != nil {
			return fmt.Errorf("error reading 'status': %w", err)
		}
		delete(object, "status")
	}

	if raw, found := object["timeoutSeconds"]; found {
		err = json.Unmarshal(raw, &a.TimeoutSeconds)
		if err != nil {
			return fmt.Errorf("error reading 'timeoutSeconds': %w", err)
		}
		delete(object, "timeoutSeconds")
	}

	if raw, found := object["toolBindings"]; found {
		err = json.Unmarshal(raw, &a.ToolBindings)
		if err != nil {
			return fmt.Errorf("error reading 'toolBindings': %w", err)
		}
		delete(object, "toolBindings")
	}

	if raw, found := object["toolset"]; found {
		err = json.Unmarshal(raw, &a.Toolset)
		if err != nil {
			return fmt.Errorf("error reading 'toolset': %w", err)
		}
		delete(object, "toolset")
	}

	if len(object) != 0 {
		a.AdditionalProperties = make(map[string]any)
		for fieldName, fieldBuf := range object {
			var fieldVal interface{}
			err := json.Unmarshal(fieldBuf, &fieldVal)
			if err != nil {
				return fmt.Errorf("error unmarshaling field %s: %w", fieldName, err)
			}
			a.AdditionalProperties[fieldName] = fieldVal
		}
	}
	return nil
}

// Override default JSON handling for AgentRun to handle AdditionalProperties
func (a AgentRun) MarshalJSON() ([]byte, error) {
	var err error
	object := make(map[string]json.RawMessage)

	if a.AnalysisArtifacts != nil {
		object["analysisArtifacts"], err = json.Marshal(a.AnalysisArtifacts)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'analysisArtifacts': %w", err)
		}
	}

	object["callbackToken"], err = json.Marshal(a.CallbackToken)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'callbackToken': %w", err)
	}

	object["capabilityId"], err = json.Marshal(a.CapabilityID)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'capabilityId': %w", err)
	}

	object["id"], err = json.Marshal(a.ID)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'id': %w", err)
	}

	if a.Input != nil {
		object["input"], err = json.Marshal(a.Input)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'input': %w", err)
		}
	}

	if a.Output != nil {
		object["output"], err = json.Marshal(a.Output)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'output': %w", err)
		}
	}

	object["providerId"], err = json.Marshal(a.ProviderID)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'providerId': %w", err)
	}

	object["providerKind"], err = json.Marshal(a.ProviderKind)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'providerKind': %w", err)
	}

	if a.Scope != nil {
		object["scope"], err = json.Marshal(a.Scope)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'scope': %w", err)
		}
	}

	object["sessionId"], err = json.Marshal(a.SessionID)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'sessionId': %w", err)
	}

	if a.SkillBindings != nil {
		object["skillBindings"], err = json.Marshal(a.SkillBindings)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'skillBindings': %w", err)
		}
	}

	if a.SkillIDs != nil {
		object["skillIds"], err = json.Marshal(a.SkillIDs)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'skillIds': %w", err)
		}
	}

	object["status"], err = json.Marshal(a.Status)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'status': %w", err)
	}

	object["timeoutSeconds"], err = json.Marshal(a.TimeoutSeconds)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'timeoutSeconds': %w", err)
	}

	if a.ToolBindings != nil {
		object["toolBindings"], err = json.Marshal(a.ToolBindings)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'toolBindings': %w", err)
		}
	}

	if a.Toolset != nil {
		object["toolset"], err = json.Marshal(a.Toolset)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'toolset': %w", err)
		}
	}

	for fieldName, field := range a.AdditionalProperties {
		object[fieldName], err = json.Marshal(field)
		if err != nil {
			return nil, fmt.Errorf("error marshaling '%s': %w", fieldName, err)
		}
	}
	return json.Marshal(object)
}

// Getter for additional properties for AgentToolCallResult. Returns the specified
// element and whether it was found
func (a AgentToolCallResult) Get(fieldName string) (value interface{}, found bool) {
	if a.AdditionalProperties != nil {
		value, found = a.AdditionalProperties[fieldName]
	}
	return
}

// Setter for additional properties for AgentToolCallResult
func (a *AgentToolCallResult) Set(fieldName string, value interface{}) {
	if a.AdditionalProperties == nil {
		a.AdditionalProperties = make(map[string]any)
	}
	a.AdditionalProperties[fieldName] = value
}

// Override default JSON handling for AgentToolCallResult to handle AdditionalProperties
func (a *AgentToolCallResult) UnmarshalJSON(b []byte) error {
	object := make(map[string]json.RawMessage)
	err := json.Unmarshal(b, &object)
	if err != nil {
		return err
	}

	if raw, found := object["output"]; found {
		err = json.Unmarshal(raw, &a.Output)
		if err != nil {
			return fmt.Errorf("error reading 'output': %w", err)
		}
		delete(object, "output")
	}

	if raw, found := object["runId"]; found {
		err = json.Unmarshal(raw, &a.RunID)
		if err != nil {
			return fmt.Errorf("error reading 'runId': %w", err)
		}
		delete(object, "runId")
	}

	if raw, found := object["toolExecution"]; found {
		err = json.Unmarshal(raw, &a.ToolExecution)
		if err != nil {
			return fmt.Errorf("error reading 'toolExecution': %w", err)
		}
		delete(object, "toolExecution")
	}

	if len(object) != 0 {
		a.AdditionalProperties = make(map[string]any)
		for fieldName, fieldBuf := range object {
			var fieldVal interface{}
			err := json.Unmarshal(fieldBuf, &fieldVal)
			if err != nil {
				return fmt.Errorf("error unmarshaling field %s: %w", fieldName, err)
			}
			a.AdditionalProperties[fieldName] = fieldVal
		}
	}
	return nil
}

// Override default JSON handling for AgentToolCallResult to handle AdditionalProperties
func (a AgentToolCallResult) MarshalJSON() ([]byte, error) {
	var err error
	object := make(map[string]json.RawMessage)

	if a.Output != nil {
		object["output"], err = json.Marshal(a.Output)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'output': %w", err)
		}
	}

	object["runId"], err = json.Marshal(a.RunID)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'runId': %w", err)
	}

	if a.ToolExecution != nil {
		object["toolExecution"], err = json.Marshal(a.ToolExecution)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'toolExecution': %w", err)
		}
	}

	for fieldName, field := range a.AdditionalProperties {
		object[fieldName], err = json.Marshal(field)
		if err != nil {
			return nil, fmt.Errorf("error marshaling '%s': %w", fieldName, err)
		}
	}
	return json.Marshal(object)
}

// Getter for additional properties for AuthBootstrapEnvelope_Data. Returns the specified
// element and whether it was found
func (a AuthBootstrapEnvelope_Data) Get(fieldName string) (value interface{}, found bool) {
	if a.AdditionalProperties != nil {
		value, found = a.AdditionalProperties[fieldName]
	}
	return
}

// Setter for additional properties for AuthBootstrapEnvelope_Data
func (a *AuthBootstrapEnvelope_Data) Set(fieldName string, value interface{}) {
	if a.AdditionalProperties == nil {
		a.AdditionalProperties = make(map[string]any)
	}
	a.AdditionalProperties[fieldName] = value
}

// Override default JSON handling for AuthBootstrapEnvelope_Data to handle AdditionalProperties
func (a *AuthBootstrapEnvelope_Data) UnmarshalJSON(b []byte) error {
	object := make(map[string]json.RawMessage)
	err := json.Unmarshal(b, &object)
	if err != nil {
		return err
	}

	if raw, found := object["branding"]; found {
		err = json.Unmarshal(raw, &a.Branding)
		if err != nil {
			return fmt.Errorf("error reading 'branding': %w", err)
		}
		delete(object, "branding")
	}

	if raw, found := object["currentUser"]; found {
		err = json.Unmarshal(raw, &a.CurrentUser)
		if err != nil {
			return fmt.Errorf("error reading 'currentUser': %w", err)
		}
		delete(object, "currentUser")
	}

	if raw, found := object["permissionSnapshot"]; found {
		err = json.Unmarshal(raw, &a.PermissionSnapshot)
		if err != nil {
			return fmt.Errorf("error reading 'permissionSnapshot': %w", err)
		}
		delete(object, "permissionSnapshot")
	}

	if raw, found := object["user"]; found {
		err = json.Unmarshal(raw, &a.User)
		if err != nil {
			return fmt.Errorf("error reading 'user': %w", err)
		}
		delete(object, "user")
	}

	if len(object) != 0 {
		a.AdditionalProperties = make(map[string]any)
		for fieldName, fieldBuf := range object {
			var fieldVal interface{}
			err := json.Unmarshal(fieldBuf, &fieldVal)
			if err != nil {
				return fmt.Errorf("error unmarshaling field %s: %w", fieldName, err)
			}
			a.AdditionalProperties[fieldName] = fieldVal
		}
	}
	return nil
}

// Override default JSON handling for AuthBootstrapEnvelope_Data to handle AdditionalProperties
func (a AuthBootstrapEnvelope_Data) MarshalJSON() ([]byte, error) {
	var err error
	object := make(map[string]json.RawMessage)

	if a.Branding != nil {
		object["branding"], err = json.Marshal(a.Branding)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'branding': %w", err)
		}
	}

	object["currentUser"], err = json.Marshal(a.CurrentUser)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'currentUser': %w", err)
	}

	if a.PermissionSnapshot != nil {
		object["permissionSnapshot"], err = json.Marshal(a.PermissionSnapshot)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'permissionSnapshot': %w", err)
		}
	}

	object["user"], err = json.Marshal(a.User)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'user': %w", err)
	}

	for fieldName, field := range a.AdditionalProperties {
		object[fieldName], err = json.Marshal(field)
		if err != nil {
			return nil, fmt.Errorf("error marshaling '%s': %w", fieldName, err)
		}
	}
	return json.Marshal(object)
}

// Getter for additional properties for DockerOperation. Returns the specified
// element and whether it was found
func (a DockerOperation) Get(fieldName string) (value interface{}, found bool) {
	if a.AdditionalProperties != nil {
		value, found = a.AdditionalProperties[fieldName]
	}
	return
}

// Setter for additional properties for DockerOperation
func (a *DockerOperation) Set(fieldName string, value interface{}) {
	if a.AdditionalProperties == nil {
		a.AdditionalProperties = make(map[string]any)
	}
	a.AdditionalProperties[fieldName] = value
}

// Override default JSON handling for DockerOperation to handle AdditionalProperties
func (a *DockerOperation) UnmarshalJSON(b []byte) error {
	object := make(map[string]json.RawMessage)
	err := json.Unmarshal(b, &object)
	if err != nil {
		return err
	}

	if raw, found := object["claimedByWorkerId"]; found {
		err = json.Unmarshal(raw, &a.ClaimedByWorkerID)
		if err != nil {
			return fmt.Errorf("error reading 'claimedByWorkerId': %w", err)
		}
		delete(object, "claimedByWorkerId")
	}

	if raw, found := object["hostId"]; found {
		err = json.Unmarshal(raw, &a.HostID)
		if err != nil {
			return fmt.Errorf("error reading 'hostId': %w", err)
		}
		delete(object, "hostId")
	}

	if raw, found := object["id"]; found {
		err = json.Unmarshal(raw, &a.ID)
		if err != nil {
			return fmt.Errorf("error reading 'id': %w", err)
		}
		delete(object, "id")
	}

	if raw, found := object["operationKind"]; found {
		err = json.Unmarshal(raw, &a.OperationKind)
		if err != nil {
			return fmt.Errorf("error reading 'operationKind': %w", err)
		}
		delete(object, "operationKind")
	}

	if raw, found := object["payload"]; found {
		err = json.Unmarshal(raw, &a.Payload)
		if err != nil {
			return fmt.Errorf("error reading 'payload': %w", err)
		}
		delete(object, "payload")
	}

	if raw, found := object["projectId"]; found {
		err = json.Unmarshal(raw, &a.ProjectID)
		if err != nil {
			return fmt.Errorf("error reading 'projectId': %w", err)
		}
		delete(object, "projectId")
	}

	if raw, found := object["serviceId"]; found {
		err = json.Unmarshal(raw, &a.ServiceID)
		if err != nil {
			return fmt.Errorf("error reading 'serviceId': %w", err)
		}
		delete(object, "serviceId")
	}

	if raw, found := object["status"]; found {
		err = json.Unmarshal(raw, &a.Status)
		if err != nil {
			return fmt.Errorf("error reading 'status': %w", err)
		}
		delete(object, "status")
	}

	if raw, found := object["timeoutSeconds"]; found {
		err = json.Unmarshal(raw, &a.TimeoutSeconds)
		if err != nil {
			return fmt.Errorf("error reading 'timeoutSeconds': %w", err)
		}
		delete(object, "timeoutSeconds")
	}

	if len(object) != 0 {
		a.AdditionalProperties = make(map[string]any)
		for fieldName, fieldBuf := range object {
			var fieldVal interface{}
			err := json.Unmarshal(fieldBuf, &fieldVal)
			if err != nil {
				return fmt.Errorf("error unmarshaling field %s: %w", fieldName, err)
			}
			a.AdditionalProperties[fieldName] = fieldVal
		}
	}
	return nil
}

// Override default JSON handling for DockerOperation to handle AdditionalProperties
func (a DockerOperation) MarshalJSON() ([]byte, error) {
	var err error
	object := make(map[string]json.RawMessage)

	object["claimedByWorkerId"], err = json.Marshal(a.ClaimedByWorkerID)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'claimedByWorkerId': %w", err)
	}

	object["hostId"], err = json.Marshal(a.HostID)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'hostId': %w", err)
	}

	object["id"], err = json.Marshal(a.ID)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'id': %w", err)
	}

	object["operationKind"], err = json.Marshal(a.OperationKind)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'operationKind': %w", err)
	}

	if a.Payload != nil {
		object["payload"], err = json.Marshal(a.Payload)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'payload': %w", err)
		}
	}

	object["projectId"], err = json.Marshal(a.ProjectID)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'projectId': %w", err)
	}

	object["serviceId"], err = json.Marshal(a.ServiceID)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'serviceId': %w", err)
	}

	object["status"], err = json.Marshal(a.Status)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'status': %w", err)
	}

	object["timeoutSeconds"], err = json.Marshal(a.TimeoutSeconds)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'timeoutSeconds': %w", err)
	}

	for fieldName, field := range a.AdditionalProperties {
		object[fieldName], err = json.Marshal(field)
		if err != nil {
			return nil, fmt.Errorf("error marshaling '%s': %w", fieldName, err)
		}
	}
	return json.Marshal(object)
}

// Getter for additional properties for ExecutionTask. Returns the specified
// element and whether it was found
func (a ExecutionTask) Get(fieldName string) (value interface{}, found bool) {
	if a.AdditionalProperties != nil {
		value, found = a.AdditionalProperties[fieldName]
	}
	return
}

// Setter for additional properties for ExecutionTask
func (a *ExecutionTask) Set(fieldName string, value interface{}) {
	if a.AdditionalProperties == nil {
		a.AdditionalProperties = make(map[string]any)
	}
	a.AdditionalProperties[fieldName] = value
}

// Override default JSON handling for ExecutionTask to handle AdditionalProperties
func (a *ExecutionTask) UnmarshalJSON(b []byte) error {
	object := make(map[string]json.RawMessage)
	err := json.Unmarshal(b, &object)
	if err != nil {
		return err
	}

	if raw, found := object["applicationEnvironmentId"]; found {
		err = json.Unmarshal(raw, &a.ApplicationEnvironmentID)
		if err != nil {
			return fmt.Errorf("error reading 'applicationEnvironmentId': %w", err)
		}
		delete(object, "applicationEnvironmentId")
	}

	if raw, found := object["applicationId"]; found {
		err = json.Unmarshal(raw, &a.ApplicationID)
		if err != nil {
			return fmt.Errorf("error reading 'applicationId': %w", err)
		}
		delete(object, "applicationId")
	}

	if raw, found := object["callbackToken"]; found {
		err = json.Unmarshal(raw, &a.CallbackToken)
		if err != nil {
			return fmt.Errorf("error reading 'callbackToken': %w", err)
		}
		delete(object, "callbackToken")
	}

	if raw, found := object["id"]; found {
		err = json.Unmarshal(raw, &a.ID)
		if err != nil {
			return fmt.Errorf("error reading 'id': %w", err)
		}
		delete(object, "id")
	}

	if raw, found := object["payload"]; found {
		err = json.Unmarshal(raw, &a.Payload)
		if err != nil {
			return fmt.Errorf("error reading 'payload': %w", err)
		}
		delete(object, "payload")
	}

	if raw, found := object["providerKind"]; found {
		err = json.Unmarshal(raw, &a.ProviderKind)
		if err != nil {
			return fmt.Errorf("error reading 'providerKind': %w", err)
		}
		delete(object, "providerKind")
	}

	if raw, found := object["status"]; found {
		err = json.Unmarshal(raw, &a.Status)
		if err != nil {
			return fmt.Errorf("error reading 'status': %w", err)
		}
		delete(object, "status")
	}

	if raw, found := object["taskKind"]; found {
		err = json.Unmarshal(raw, &a.TaskKind)
		if err != nil {
			return fmt.Errorf("error reading 'taskKind': %w", err)
		}
		delete(object, "taskKind")
	}

	if len(object) != 0 {
		a.AdditionalProperties = make(map[string]any)
		for fieldName, fieldBuf := range object {
			var fieldVal interface{}
			err := json.Unmarshal(fieldBuf, &fieldVal)
			if err != nil {
				return fmt.Errorf("error unmarshaling field %s: %w", fieldName, err)
			}
			a.AdditionalProperties[fieldName] = fieldVal
		}
	}
	return nil
}

// Override default JSON handling for ExecutionTask to handle AdditionalProperties
func (a ExecutionTask) MarshalJSON() ([]byte, error) {
	var err error
	object := make(map[string]json.RawMessage)

	object["applicationEnvironmentId"], err = json.Marshal(a.ApplicationEnvironmentID)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'applicationEnvironmentId': %w", err)
	}

	object["applicationId"], err = json.Marshal(a.ApplicationID)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'applicationId': %w", err)
	}

	object["callbackToken"], err = json.Marshal(a.CallbackToken)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'callbackToken': %w", err)
	}

	object["id"], err = json.Marshal(a.ID)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'id': %w", err)
	}

	if a.Payload != nil {
		object["payload"], err = json.Marshal(a.Payload)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'payload': %w", err)
		}
	}

	object["providerKind"], err = json.Marshal(a.ProviderKind)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'providerKind': %w", err)
	}

	object["status"], err = json.Marshal(a.Status)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'status': %w", err)
	}

	object["taskKind"], err = json.Marshal(a.TaskKind)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'taskKind': %w", err)
	}

	for fieldName, field := range a.AdditionalProperties {
		object[fieldName], err = json.Marshal(field)
		if err != nil {
			return nil, fmt.Errorf("error marshaling '%s': %w", fieldName, err)
		}
	}
	return json.Marshal(object)
}

// Getter for additional properties for PluginCompatibility. Returns the specified
// element and whether it was found
func (a PluginCompatibility) Get(fieldName string) (value interface{}, found bool) {
	if a.AdditionalProperties != nil {
		value, found = a.AdditionalProperties[fieldName]
	}
	return
}

// Setter for additional properties for PluginCompatibility
func (a *PluginCompatibility) Set(fieldName string, value interface{}) {
	if a.AdditionalProperties == nil {
		a.AdditionalProperties = make(map[string]any)
	}
	a.AdditionalProperties[fieldName] = value
}

// Override default JSON handling for PluginCompatibility to handle AdditionalProperties
func (a *PluginCompatibility) UnmarshalJSON(b []byte) error {
	object := make(map[string]json.RawMessage)
	err := json.Unmarshal(b, &object)
	if err != nil {
		return err
	}

	if raw, found := object["contracts"]; found {
		err = json.Unmarshal(raw, &a.Contracts)
		if err != nil {
			return fmt.Errorf("error reading 'contracts': %w", err)
		}
		delete(object, "contracts")
	}

	if raw, found := object["soha"]; found {
		err = json.Unmarshal(raw, &a.Soha)
		if err != nil {
			return fmt.Errorf("error reading 'soha': %w", err)
		}
		delete(object, "soha")
	}

	if len(object) != 0 {
		a.AdditionalProperties = make(map[string]any)
		for fieldName, fieldBuf := range object {
			var fieldVal interface{}
			err := json.Unmarshal(fieldBuf, &fieldVal)
			if err != nil {
				return fmt.Errorf("error unmarshaling field %s: %w", fieldName, err)
			}
			a.AdditionalProperties[fieldName] = fieldVal
		}
	}
	return nil
}

// Override default JSON handling for PluginCompatibility to handle AdditionalProperties
func (a PluginCompatibility) MarshalJSON() ([]byte, error) {
	var err error
	object := make(map[string]json.RawMessage)

	object["contracts"], err = json.Marshal(a.Contracts)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'contracts': %w", err)
	}

	object["soha"], err = json.Marshal(a.Soha)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'soha': %w", err)
	}

	for fieldName, field := range a.AdditionalProperties {
		object[fieldName], err = json.Marshal(field)
		if err != nil {
			return nil, fmt.Errorf("error marshaling '%s': %w", fieldName, err)
		}
	}
	return json.Marshal(object)
}

// Getter for additional properties for Principal. Returns the specified
// element and whether it was found
func (a Principal) Get(fieldName string) (value interface{}, found bool) {
	if a.AdditionalProperties != nil {
		value, found = a.AdditionalProperties[fieldName]
	}
	return
}

// Setter for additional properties for Principal
func (a *Principal) Set(fieldName string, value interface{}) {
	if a.AdditionalProperties == nil {
		a.AdditionalProperties = make(map[string]any)
	}
	a.AdditionalProperties[fieldName] = value
}

// Override default JSON handling for Principal to handle AdditionalProperties
func (a *Principal) UnmarshalJSON(b []byte) error {
	object := make(map[string]json.RawMessage)
	err := json.Unmarshal(b, &object)
	if err != nil {
		return err
	}

	if raw, found := object["email"]; found {
		err = json.Unmarshal(raw, &a.Email)
		if err != nil {
			return fmt.Errorf("error reading 'email': %w", err)
		}
		delete(object, "email")
	}

	if raw, found := object["permissionKeys"]; found {
		err = json.Unmarshal(raw, &a.PermissionKeys)
		if err != nil {
			return fmt.Errorf("error reading 'permissionKeys': %w", err)
		}
		delete(object, "permissionKeys")
	}

	if raw, found := object["projects"]; found {
		err = json.Unmarshal(raw, &a.Projects)
		if err != nil {
			return fmt.Errorf("error reading 'projects': %w", err)
		}
		delete(object, "projects")
	}

	if raw, found := object["roles"]; found {
		err = json.Unmarshal(raw, &a.Roles)
		if err != nil {
			return fmt.Errorf("error reading 'roles': %w", err)
		}
		delete(object, "roles")
	}

	if raw, found := object["tags"]; found {
		err = json.Unmarshal(raw, &a.Tags)
		if err != nil {
			return fmt.Errorf("error reading 'tags': %w", err)
		}
		delete(object, "tags")
	}

	if raw, found := object["teams"]; found {
		err = json.Unmarshal(raw, &a.Teams)
		if err != nil {
			return fmt.Errorf("error reading 'teams': %w", err)
		}
		delete(object, "teams")
	}

	if raw, found := object["userId"]; found {
		err = json.Unmarshal(raw, &a.UserID)
		if err != nil {
			return fmt.Errorf("error reading 'userId': %w", err)
		}
		delete(object, "userId")
	}

	if raw, found := object["userName"]; found {
		err = json.Unmarshal(raw, &a.UserName)
		if err != nil {
			return fmt.Errorf("error reading 'userName': %w", err)
		}
		delete(object, "userName")
	}

	if len(object) != 0 {
		a.AdditionalProperties = make(map[string]any)
		for fieldName, fieldBuf := range object {
			var fieldVal interface{}
			err := json.Unmarshal(fieldBuf, &fieldVal)
			if err != nil {
				return fmt.Errorf("error unmarshaling field %s: %w", fieldName, err)
			}
			a.AdditionalProperties[fieldName] = fieldVal
		}
	}
	return nil
}

// Override default JSON handling for Principal to handle AdditionalProperties
func (a Principal) MarshalJSON() ([]byte, error) {
	var err error
	object := make(map[string]json.RawMessage)

	object["email"], err = json.Marshal(a.Email)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'email': %w", err)
	}

	if a.PermissionKeys != nil {
		object["permissionKeys"], err = json.Marshal(a.PermissionKeys)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'permissionKeys': %w", err)
		}
	}

	if a.Projects != nil {
		object["projects"], err = json.Marshal(a.Projects)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'projects': %w", err)
		}
	}

	if a.Roles != nil {
		object["roles"], err = json.Marshal(a.Roles)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'roles': %w", err)
		}
	}

	if a.Tags != nil {
		object["tags"], err = json.Marshal(a.Tags)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'tags': %w", err)
		}
	}

	if a.Teams != nil {
		object["teams"], err = json.Marshal(a.Teams)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'teams': %w", err)
		}
	}

	object["userId"], err = json.Marshal(a.UserID)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'userId': %w", err)
	}

	object["userName"], err = json.Marshal(a.UserName)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'userName': %w", err)
	}

	for fieldName, field := range a.AdditionalProperties {
		object[fieldName], err = json.Marshal(field)
		if err != nil {
			return nil, fmt.Errorf("error marshaling '%s': %w", fieldName, err)
		}
	}
	return json.Marshal(object)
}

// Getter for additional properties for PromptGetResult. Returns the specified
// element and whether it was found
func (a PromptGetResult) Get(fieldName string) (value interface{}, found bool) {
	if a.AdditionalProperties != nil {
		value, found = a.AdditionalProperties[fieldName]
	}
	return
}

// Setter for additional properties for PromptGetResult
func (a *PromptGetResult) Set(fieldName string, value interface{}) {
	if a.AdditionalProperties == nil {
		a.AdditionalProperties = make(map[string]any)
	}
	a.AdditionalProperties[fieldName] = value
}

// Override default JSON handling for PromptGetResult to handle AdditionalProperties
func (a *PromptGetResult) UnmarshalJSON(b []byte) error {
	object := make(map[string]json.RawMessage)
	err := json.Unmarshal(b, &object)
	if err != nil {
		return err
	}

	if raw, found := object["audit"]; found {
		err = json.Unmarshal(raw, &a.Audit)
		if err != nil {
			return fmt.Errorf("error reading 'audit': %w", err)
		}
		delete(object, "audit")
	}

	if raw, found := object["description"]; found {
		err = json.Unmarshal(raw, &a.Description)
		if err != nil {
			return fmt.Errorf("error reading 'description': %w", err)
		}
		delete(object, "description")
	}

	if raw, found := object["messages"]; found {
		err = json.Unmarshal(raw, &a.Messages)
		if err != nil {
			return fmt.Errorf("error reading 'messages': %w", err)
		}
		delete(object, "messages")
	}

	if raw, found := object["name"]; found {
		err = json.Unmarshal(raw, &a.Name)
		if err != nil {
			return fmt.Errorf("error reading 'name': %w", err)
		}
		delete(object, "name")
	}

	if raw, found := object["relatedIds"]; found {
		err = json.Unmarshal(raw, &a.RelatedIDs)
		if err != nil {
			return fmt.Errorf("error reading 'relatedIds': %w", err)
		}
		delete(object, "relatedIds")
	}

	if len(object) != 0 {
		a.AdditionalProperties = make(map[string]any)
		for fieldName, fieldBuf := range object {
			var fieldVal interface{}
			err := json.Unmarshal(fieldBuf, &fieldVal)
			if err != nil {
				return fmt.Errorf("error unmarshaling field %s: %w", fieldName, err)
			}
			a.AdditionalProperties[fieldName] = fieldVal
		}
	}
	return nil
}

// Override default JSON handling for PromptGetResult to handle AdditionalProperties
func (a PromptGetResult) MarshalJSON() ([]byte, error) {
	var err error
	object := make(map[string]json.RawMessage)

	if a.Audit != nil {
		object["audit"], err = json.Marshal(a.Audit)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'audit': %w", err)
		}
	}

	object["description"], err = json.Marshal(a.Description)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'description': %w", err)
	}

	if a.Messages != nil {
		object["messages"], err = json.Marshal(a.Messages)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'messages': %w", err)
		}
	}

	object["name"], err = json.Marshal(a.Name)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'name': %w", err)
	}

	if a.RelatedIDs != nil {
		object["relatedIds"], err = json.Marshal(a.RelatedIDs)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'relatedIds': %w", err)
		}
	}

	for fieldName, field := range a.AdditionalProperties {
		object[fieldName], err = json.Marshal(field)
		if err != nil {
			return nil, fmt.Errorf("error marshaling '%s': %w", fieldName, err)
		}
	}
	return json.Marshal(object)
}

// Getter for additional properties for ResourceReadResult. Returns the specified
// element and whether it was found
func (a ResourceReadResult) Get(fieldName string) (value interface{}, found bool) {
	if a.AdditionalProperties != nil {
		value, found = a.AdditionalProperties[fieldName]
	}
	return
}

// Setter for additional properties for ResourceReadResult
func (a *ResourceReadResult) Set(fieldName string, value interface{}) {
	if a.AdditionalProperties == nil {
		a.AdditionalProperties = make(map[string]any)
	}
	a.AdditionalProperties[fieldName] = value
}

// Override default JSON handling for ResourceReadResult to handle AdditionalProperties
func (a *ResourceReadResult) UnmarshalJSON(b []byte) error {
	object := make(map[string]json.RawMessage)
	err := json.Unmarshal(b, &object)
	if err != nil {
		return err
	}

	if raw, found := object["audit"]; found {
		err = json.Unmarshal(raw, &a.Audit)
		if err != nil {
			return fmt.Errorf("error reading 'audit': %w", err)
		}
		delete(object, "audit")
	}

	if raw, found := object["data"]; found {
		err = json.Unmarshal(raw, &a.Data)
		if err != nil {
			return fmt.Errorf("error reading 'data': %w", err)
		}
		delete(object, "data")
	}

	if raw, found := object["mimeType"]; found {
		err = json.Unmarshal(raw, &a.MIMEType)
		if err != nil {
			return fmt.Errorf("error reading 'mimeType': %w", err)
		}
		delete(object, "mimeType")
	}

	if raw, found := object["name"]; found {
		err = json.Unmarshal(raw, &a.Name)
		if err != nil {
			return fmt.Errorf("error reading 'name': %w", err)
		}
		delete(object, "name")
	}

	if raw, found := object["relatedIds"]; found {
		err = json.Unmarshal(raw, &a.RelatedIDs)
		if err != nil {
			return fmt.Errorf("error reading 'relatedIds': %w", err)
		}
		delete(object, "relatedIds")
	}

	if raw, found := object["text"]; found {
		err = json.Unmarshal(raw, &a.Text)
		if err != nil {
			return fmt.Errorf("error reading 'text': %w", err)
		}
		delete(object, "text")
	}

	if raw, found := object["uri"]; found {
		err = json.Unmarshal(raw, &a.URI)
		if err != nil {
			return fmt.Errorf("error reading 'uri': %w", err)
		}
		delete(object, "uri")
	}

	if len(object) != 0 {
		a.AdditionalProperties = make(map[string]any)
		for fieldName, fieldBuf := range object {
			var fieldVal interface{}
			err := json.Unmarshal(fieldBuf, &fieldVal)
			if err != nil {
				return fmt.Errorf("error unmarshaling field %s: %w", fieldName, err)
			}
			a.AdditionalProperties[fieldName] = fieldVal
		}
	}
	return nil
}

// Override default JSON handling for ResourceReadResult to handle AdditionalProperties
func (a ResourceReadResult) MarshalJSON() ([]byte, error) {
	var err error
	object := make(map[string]json.RawMessage)

	if a.Audit != nil {
		object["audit"], err = json.Marshal(a.Audit)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'audit': %w", err)
		}
	}

	object["data"], err = json.Marshal(a.Data)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'data': %w", err)
	}

	object["mimeType"], err = json.Marshal(a.MIMEType)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'mimeType': %w", err)
	}

	object["name"], err = json.Marshal(a.Name)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'name': %w", err)
	}

	if a.RelatedIDs != nil {
		object["relatedIds"], err = json.Marshal(a.RelatedIDs)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'relatedIds': %w", err)
		}
	}

	object["text"], err = json.Marshal(a.Text)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'text': %w", err)
	}

	object["uri"], err = json.Marshal(a.URI)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'uri': %w", err)
	}

	for fieldName, field := range a.AdditionalProperties {
		object[fieldName], err = json.Marshal(field)
		if err != nil {
			return nil, fmt.Errorf("error marshaling '%s': %w", fieldName, err)
		}
	}
	return json.Marshal(object)
}

// Getter for additional properties for ToolInvocationResult. Returns the specified
// element and whether it was found
func (a ToolInvocationResult) Get(fieldName string) (value interface{}, found bool) {
	if a.AdditionalProperties != nil {
		value, found = a.AdditionalProperties[fieldName]
	}
	return
}

// Setter for additional properties for ToolInvocationResult
func (a *ToolInvocationResult) Set(fieldName string, value interface{}) {
	if a.AdditionalProperties == nil {
		a.AdditionalProperties = make(map[string]any)
	}
	a.AdditionalProperties[fieldName] = value
}

// Override default JSON handling for ToolInvocationResult to handle AdditionalProperties
func (a *ToolInvocationResult) UnmarshalJSON(b []byte) error {
	object := make(map[string]json.RawMessage)
	err := json.Unmarshal(b, &object)
	if err != nil {
		return err
	}

	if raw, found := object["audit"]; found {
		err = json.Unmarshal(raw, &a.Audit)
		if err != nil {
			return fmt.Errorf("error reading 'audit': %w", err)
		}
		delete(object, "audit")
	}

	if raw, found := object["output"]; found {
		err = json.Unmarshal(raw, &a.Output)
		if err != nil {
			return fmt.Errorf("error reading 'output': %w", err)
		}
		delete(object, "output")
	}

	if raw, found := object["relatedIds"]; found {
		err = json.Unmarshal(raw, &a.RelatedIDs)
		if err != nil {
			return fmt.Errorf("error reading 'relatedIds': %w", err)
		}
		delete(object, "relatedIds")
	}

	if raw, found := object["requiresApproval"]; found {
		err = json.Unmarshal(raw, &a.RequiresApproval)
		if err != nil {
			return fmt.Errorf("error reading 'requiresApproval': %w", err)
		}
		delete(object, "requiresApproval")
	}

	if raw, found := object["result"]; found {
		err = json.Unmarshal(raw, &a.Result)
		if err != nil {
			return fmt.Errorf("error reading 'result': %w", err)
		}
		delete(object, "result")
	}

	if raw, found := object["riskLevel"]; found {
		err = json.Unmarshal(raw, &a.RiskLevel)
		if err != nil {
			return fmt.Errorf("error reading 'riskLevel': %w", err)
		}
		delete(object, "riskLevel")
	}

	if raw, found := object["toolName"]; found {
		err = json.Unmarshal(raw, &a.ToolName)
		if err != nil {
			return fmt.Errorf("error reading 'toolName': %w", err)
		}
		delete(object, "toolName")
	}

	if len(object) != 0 {
		a.AdditionalProperties = make(map[string]any)
		for fieldName, fieldBuf := range object {
			var fieldVal interface{}
			err := json.Unmarshal(fieldBuf, &fieldVal)
			if err != nil {
				return fmt.Errorf("error unmarshaling field %s: %w", fieldName, err)
			}
			a.AdditionalProperties[fieldName] = fieldVal
		}
	}
	return nil
}

// Override default JSON handling for ToolInvocationResult to handle AdditionalProperties
func (a ToolInvocationResult) MarshalJSON() ([]byte, error) {
	var err error
	object := make(map[string]json.RawMessage)

	if a.Audit != nil {
		object["audit"], err = json.Marshal(a.Audit)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'audit': %w", err)
		}
	}

	object["output"], err = json.Marshal(a.Output)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'output': %w", err)
	}

	if a.RelatedIDs != nil {
		object["relatedIds"], err = json.Marshal(a.RelatedIDs)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'relatedIds': %w", err)
		}
	}

	object["requiresApproval"], err = json.Marshal(a.RequiresApproval)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'requiresApproval': %w", err)
	}

	object["result"], err = json.Marshal(a.Result)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'result': %w", err)
	}

	object["riskLevel"], err = json.Marshal(a.RiskLevel)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'riskLevel': %w", err)
	}

	object["toolName"], err = json.Marshal(a.ToolName)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'toolName': %w", err)
	}

	for fieldName, field := range a.AdditionalProperties {
		object[fieldName], err = json.Marshal(field)
		if err != nil {
			return nil, fmt.Errorf("error marshaling '%s': %w", fieldName, err)
		}
	}
	return json.Marshal(object)
}

// Getter for additional properties for UserProfile. Returns the specified
// element and whether it was found
func (a UserProfile) Get(fieldName string) (value interface{}, found bool) {
	if a.AdditionalProperties != nil {
		value, found = a.AdditionalProperties[fieldName]
	}
	return
}

// Setter for additional properties for UserProfile
func (a *UserProfile) Set(fieldName string, value interface{}) {
	if a.AdditionalProperties == nil {
		a.AdditionalProperties = make(map[string]any)
	}
	a.AdditionalProperties[fieldName] = value
}

// Override default JSON handling for UserProfile to handle AdditionalProperties
func (a *UserProfile) UnmarshalJSON(b []byte) error {
	object := make(map[string]json.RawMessage)
	err := json.Unmarshal(b, &object)
	if err != nil {
		return err
	}

	if raw, found := object["displayName"]; found {
		err = json.Unmarshal(raw, &a.DisplayName)
		if err != nil {
			return fmt.Errorf("error reading 'displayName': %w", err)
		}
		delete(object, "displayName")
	}

	if raw, found := object["email"]; found {
		err = json.Unmarshal(raw, &a.Email)
		if err != nil {
			return fmt.Errorf("error reading 'email': %w", err)
		}
		delete(object, "email")
	}

	if raw, found := object["identities"]; found {
		err = json.Unmarshal(raw, &a.Identities)
		if err != nil {
			return fmt.Errorf("error reading 'identities': %w", err)
		}
		delete(object, "identities")
	}

	if raw, found := object["lastLoginAt"]; found {
		err = json.Unmarshal(raw, &a.LastLoginAt)
		if err != nil {
			return fmt.Errorf("error reading 'lastLoginAt': %w", err)
		}
		delete(object, "lastLoginAt")
	}

	if raw, found := object["phone"]; found {
		err = json.Unmarshal(raw, &a.Phone)
		if err != nil {
			return fmt.Errorf("error reading 'phone': %w", err)
		}
		delete(object, "phone")
	}

	if raw, found := object["projects"]; found {
		err = json.Unmarshal(raw, &a.Projects)
		if err != nil {
			return fmt.Errorf("error reading 'projects': %w", err)
		}
		delete(object, "projects")
	}

	if raw, found := object["roles"]; found {
		err = json.Unmarshal(raw, &a.Roles)
		if err != nil {
			return fmt.Errorf("error reading 'roles': %w", err)
		}
		delete(object, "roles")
	}

	if raw, found := object["sessions"]; found {
		err = json.Unmarshal(raw, &a.Sessions)
		if err != nil {
			return fmt.Errorf("error reading 'sessions': %w", err)
		}
		delete(object, "sessions")
	}

	if raw, found := object["status"]; found {
		err = json.Unmarshal(raw, &a.Status)
		if err != nil {
			return fmt.Errorf("error reading 'status': %w", err)
		}
		delete(object, "status")
	}

	if raw, found := object["tags"]; found {
		err = json.Unmarshal(raw, &a.Tags)
		if err != nil {
			return fmt.Errorf("error reading 'tags': %w", err)
		}
		delete(object, "tags")
	}

	if raw, found := object["teams"]; found {
		err = json.Unmarshal(raw, &a.Teams)
		if err != nil {
			return fmt.Errorf("error reading 'teams': %w", err)
		}
		delete(object, "teams")
	}

	if raw, found := object["userId"]; found {
		err = json.Unmarshal(raw, &a.UserID)
		if err != nil {
			return fmt.Errorf("error reading 'userId': %w", err)
		}
		delete(object, "userId")
	}

	if raw, found := object["username"]; found {
		err = json.Unmarshal(raw, &a.Username)
		if err != nil {
			return fmt.Errorf("error reading 'username': %w", err)
		}
		delete(object, "username")
	}

	if len(object) != 0 {
		a.AdditionalProperties = make(map[string]any)
		for fieldName, fieldBuf := range object {
			var fieldVal interface{}
			err := json.Unmarshal(fieldBuf, &fieldVal)
			if err != nil {
				return fmt.Errorf("error unmarshaling field %s: %w", fieldName, err)
			}
			a.AdditionalProperties[fieldName] = fieldVal
		}
	}
	return nil
}

// Override default JSON handling for UserProfile to handle AdditionalProperties
func (a UserProfile) MarshalJSON() ([]byte, error) {
	var err error
	object := make(map[string]json.RawMessage)

	object["displayName"], err = json.Marshal(a.DisplayName)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'displayName': %w", err)
	}

	object["email"], err = json.Marshal(a.Email)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'email': %w", err)
	}

	if a.Identities != nil {
		object["identities"], err = json.Marshal(a.Identities)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'identities': %w", err)
		}
	}

	if a.LastLoginAt != nil {
		object["lastLoginAt"], err = json.Marshal(a.LastLoginAt)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'lastLoginAt': %w", err)
		}
	}

	object["phone"], err = json.Marshal(a.Phone)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'phone': %w", err)
	}

	if a.Projects != nil {
		object["projects"], err = json.Marshal(a.Projects)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'projects': %w", err)
		}
	}

	if a.Roles != nil {
		object["roles"], err = json.Marshal(a.Roles)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'roles': %w", err)
		}
	}

	if a.Sessions != nil {
		object["sessions"], err = json.Marshal(a.Sessions)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'sessions': %w", err)
		}
	}

	object["status"], err = json.Marshal(a.Status)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'status': %w", err)
	}

	if a.Tags != nil {
		object["tags"], err = json.Marshal(a.Tags)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'tags': %w", err)
		}
	}

	if a.Teams != nil {
		object["teams"], err = json.Marshal(a.Teams)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'teams': %w", err)
		}
	}

	object["userId"], err = json.Marshal(a.UserID)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'userId': %w", err)
	}

	object["username"], err = json.Marshal(a.Username)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'username': %w", err)
	}

	for fieldName, field := range a.AdditionalProperties {
		object[fieldName], err = json.Marshal(field)
		if err != nil {
			return nil, fmt.Errorf("error marshaling '%s': %w", fieldName, err)
		}
	}
	return json.Marshal(object)
}
