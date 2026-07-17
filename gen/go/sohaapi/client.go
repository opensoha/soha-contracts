// Code generated from OpenSoha contracts. DO NOT EDIT.
package sohaapi

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

const defaultUserAgent = "opensoha-contracts/0.1.2"

type Client struct {
	BaseURL    string
	Token      string
	HTTPClient *http.Client
	UserAgent  string
}

type ClientOption func(*Client)

func WithBearerToken(token string) ClientOption {
	return func(c *Client) {
		c.Token = strings.TrimSpace(token)
	}
}

func WithHTTPClient(client *http.Client) ClientOption {
	return func(c *Client) {
		c.HTTPClient = client
	}
}

func WithUserAgent(userAgent string) ClientOption {
	return func(c *Client) {
		c.UserAgent = strings.TrimSpace(userAgent)
	}
}

func NewClient(baseURL string, opts ...ClientOption) *Client {
	c := &Client{
		BaseURL:   strings.TrimSpace(baseURL),
		UserAgent: defaultUserAgent,
	}
	for _, opt := range opts {
		if opt != nil {
			opt(c)
		}
	}
	return c
}

func (c *Client) GetHealthz(ctx context.Context) (GenericObject, error) {
	var out GenericObject
	if err := c.doJSON(ctx, http.MethodGet, "/healthz", false, nil, &out); err != nil {
		return nil, err
	}
	return out, nil
}

func (c *Client) GetReadyz(ctx context.Context) (GenericObject, error) {
	var out GenericObject
	if err := c.doJSON(ctx, http.MethodGet, "/readyz", false, nil, &out); err != nil {
		return nil, err
	}
	return out, nil
}

func (c *Client) GetAgentProviderCatalog(ctx context.Context) (AgentProviderCatalog, error) {
	var out AgentProviderCatalogEnvelope
	if err := c.doJSON(ctx, http.MethodGet, "/ai/agent-providers/catalog", true, nil, &out); err != nil {
		return AgentProviderCatalog{}, err
	}
	return out.Data, nil
}

func (c *Client) GetAgentProviderRuntimeStatus(ctx context.Context) (AgentProviderRuntimeStatus, error) {
	var out AgentProviderRuntimeStatusEnvelope
	if err := c.doJSON(ctx, http.MethodGet, "/ai/agent-providers/runtime-status", true, nil, &out); err != nil {
		return AgentProviderRuntimeStatus{}, err
	}
	return out.Data, nil
}

func (c *Client) GetAgentProviderRegistrySnapshot(ctx context.Context, runnerID string) (AgentProviderRegistrySnapshot, error) {
	var out AgentProviderRegistrySnapshotEnvelope
	query := url.Values{"runnerId": []string{strings.TrimSpace(runnerID)}}
	path := "/ai/agent-providers/registry-snapshot?" + query.Encode()
	if err := c.doJSON(ctx, http.MethodGet, path, true, nil, &out); err != nil {
		return AgentProviderRegistrySnapshot{}, err
	}
	return out.Data, nil
}

func (c *Client) AcknowledgeAgentProviderRegistry(ctx context.Context, req AgentProviderRegistryAcknowledgement) (AgentProviderRegistryAcknowledgement, error) {
	var out AgentProviderRegistryAcknowledgementEnvelope
	if err := c.doJSON(ctx, http.MethodPost, "/ai/agent-providers/registry-acks", true, req, &out); err != nil {
		return AgentProviderRegistryAcknowledgement{}, err
	}
	return out.Data, nil
}

func (c *Client) ListKnowledgeBases(ctx context.Context) ([]KnowledgeBase, error) {
	var out KnowledgeBaseListEnvelope
	if err := c.doJSON(ctx, http.MethodGet, "/ai/knowledge-bases", true, nil, &out); err != nil {
		return nil, err
	}
	return out.Items, nil
}

func (c *Client) GetKnowledgeBase(ctx context.Context, baseID string) (KnowledgeBase, error) {
	var out KnowledgeBaseEnvelope
	path := "/ai/knowledge-bases/" + url.PathEscape(strings.TrimSpace(baseID))
	if err := c.doJSON(ctx, http.MethodGet, path, true, nil, &out); err != nil {
		return KnowledgeBase{}, err
	}
	return out.Data, nil
}

func (c *Client) CreateKnowledgeBase(ctx context.Context, req KnowledgeBaseInput) (KnowledgeBase, error) {
	var out KnowledgeBaseEnvelope
	if err := c.doJSON(ctx, http.MethodPost, "/ai/knowledge-bases", true, req, &out); err != nil {
		return KnowledgeBase{}, err
	}
	return out.Data, nil
}

func (c *Client) UpdateKnowledgeBase(ctx context.Context, baseID string, req KnowledgeBaseInput) (KnowledgeBase, error) {
	var out KnowledgeBaseEnvelope
	path := "/ai/knowledge-bases/" + url.PathEscape(strings.TrimSpace(baseID))
	if err := c.doJSON(ctx, http.MethodPatch, path, true, req, &out); err != nil {
		return KnowledgeBase{}, err
	}
	return out.Data, nil
}

func (c *Client) DeleteKnowledgeBase(ctx context.Context, baseID string) error {
	path := "/ai/knowledge-bases/" + url.PathEscape(strings.TrimSpace(baseID))
	return c.doJSON(ctx, http.MethodDelete, path, true, nil, nil)
}

func (c *Client) ListKnowledgeSources(ctx context.Context, baseID string) ([]KnowledgeSource, error) {
	var out KnowledgeSourceListEnvelope
	path := "/ai/knowledge-bases/" + url.PathEscape(strings.TrimSpace(baseID)) + "/sources"
	if err := c.doJSON(ctx, http.MethodGet, path, true, nil, &out); err != nil {
		return nil, err
	}
	return out.Items, nil
}

func (c *Client) CreateKnowledgeSource(ctx context.Context, baseID string, req KnowledgeSourceInput) (KnowledgeSource, error) {
	var out KnowledgeSourceEnvelope
	path := "/ai/knowledge-bases/" + url.PathEscape(strings.TrimSpace(baseID)) + "/sources"
	if err := c.doJSON(ctx, http.MethodPost, path, true, req, &out); err != nil {
		return KnowledgeSource{}, err
	}
	return out.Data, nil
}

func (c *Client) SyncKnowledgeSource(ctx context.Context, baseID, sourceID string) (KnowledgeSyncRun, error) {
	var out KnowledgeSyncRunEnvelope
	path := "/ai/knowledge-bases/" + url.PathEscape(strings.TrimSpace(baseID)) + "/sources/" + url.PathEscape(strings.TrimSpace(sourceID)) + "/sync"
	if err := c.doJSON(ctx, http.MethodPost, path, true, nil, &out); err != nil {
		return KnowledgeSyncRun{}, err
	}
	return out.Data, nil
}

func (c *Client) ListKnowledgeDocuments(ctx context.Context, baseID string, limit int) ([]KnowledgeDocument, error) {
	var out KnowledgeDocumentListEnvelope
	path := knowledgeListPath(baseID, "documents", limit)
	if err := c.doJSON(ctx, http.MethodGet, path, true, nil, &out); err != nil {
		return nil, err
	}
	return out.Items, nil
}

func (c *Client) ListKnowledgeSyncRuns(ctx context.Context, baseID string, limit int) ([]KnowledgeSyncRun, error) {
	var out KnowledgeSyncRunListEnvelope
	path := knowledgeListPath(baseID, "sync-runs", limit)
	if err := c.doJSON(ctx, http.MethodGet, path, true, nil, &out); err != nil {
		return nil, err
	}
	return out.Items, nil
}

func (c *Client) ListKnowledgeIndexRevisions(ctx context.Context, baseID string, limit int) ([]KnowledgeIndexRevision, error) {
	var out KnowledgeIndexRevisionListEnvelope
	path := knowledgeListPath(baseID, "index-revisions", limit)
	if err := c.doJSON(ctx, http.MethodGet, path, true, nil, &out); err != nil {
		return nil, err
	}
	return out.Items, nil
}

func (c *Client) SearchKnowledge(ctx context.Context, req KnowledgeSearchRequest) (KnowledgeSearchResult, error) {
	var out KnowledgeSearchResultEnvelope
	if err := c.doJSON(ctx, http.MethodPost, "/ai/knowledge/search", true, req, &out); err != nil {
		return KnowledgeSearchResult{}, err
	}
	return out.Data, nil
}

func (c *Client) InspectAIContext(ctx context.Context, req AIContextBuildInput) (AIContextInspection, error) {
	var out AIContextInspectionEnvelope
	if err := c.doJSON(ctx, http.MethodPost, "/ai/context/inspect", true, req, &out); err != nil {
		return AIContextInspection{}, err
	}
	return out.Data, nil
}

func (c *Client) ListEvaluationDatasets(ctx context.Context) ([]EvaluationDataset, error) {
	var out EvaluationDatasetListEnvelope
	if err := c.doJSON(ctx, http.MethodGet, "/ai/evaluations/datasets", true, nil, &out); err != nil {
		return nil, err
	}
	return out.Items, nil
}

func (c *Client) CreateEvaluationDataset(ctx context.Context, req EvaluationDataset) (EvaluationDataset, error) {
	var out EvaluationDatasetEnvelope
	if err := c.doJSON(ctx, http.MethodPost, "/ai/evaluations/datasets", true, req, &out); err != nil {
		return EvaluationDataset{}, err
	}
	return out.Data, nil
}

func (c *Client) ListEvaluationRuns(ctx context.Context) ([]EvaluationRun, error) {
	var out EvaluationRunListEnvelope
	if err := c.doJSON(ctx, http.MethodGet, "/ai/evaluations/runs", true, nil, &out); err != nil {
		return nil, err
	}
	return out.Items, nil
}

func (c *Client) StartEvaluationRun(ctx context.Context, req EvaluationRun) (EvaluationRun, error) {
	var out EvaluationRunEnvelope
	if err := c.doJSON(ctx, http.MethodPost, "/ai/evaluations/runs", true, req, &out); err != nil {
		return EvaluationRun{}, err
	}
	return out.Data, nil
}

func (c *Client) GetEvaluationRun(ctx context.Context, runID string) (EvaluationRun, error) {
	var out EvaluationRunEnvelope
	path := "/ai/evaluations/runs/" + url.PathEscape(strings.TrimSpace(runID))
	if err := c.doJSON(ctx, http.MethodGet, path, true, nil, &out); err != nil {
		return EvaluationRun{}, err
	}
	return out.Data, nil
}

func (c *Client) ListEvaluationRunResults(ctx context.Context, runID string) ([]EvaluationResult, error) {
	var out EvaluationResultListEnvelope
	path := "/ai/evaluations/runs/" + url.PathEscape(strings.TrimSpace(runID)) + "/results"
	if err := c.doJSON(ctx, http.MethodGet, path, true, nil, &out); err != nil {
		return nil, err
	}
	return out.Items, nil
}

func (c *Client) CompleteEvaluationRun(ctx context.Context, runID string, req EvaluationCompleteRunInput) (EvaluationRun, error) {
	var out EvaluationRunEnvelope
	path := "/ai/evaluations/runs/" + url.PathEscape(strings.TrimSpace(runID)) + "/complete"
	if err := c.doJSON(ctx, http.MethodPost, path, true, req, &out); err != nil {
		return EvaluationRun{}, err
	}
	return out.Data, nil
}

func knowledgeListPath(baseID, collection string, limit int) string {
	path := "/ai/knowledge-bases/" + url.PathEscape(strings.TrimSpace(baseID)) + "/" + collection
	if limit > 0 {
		path += "?limit=" + strconv.Itoa(limit)
	}
	return path
}

func (c *Client) ListAuthProviders(ctx context.Context) ([]AuthProvider, error) {
	var out AuthProviderListEnvelope
	if err := c.doJSON(ctx, http.MethodGet, "/auth/providers", false, nil, &out); err != nil {
		return nil, err
	}
	return out.Items, nil
}

func (c *Client) GetAuthLoginOptions(ctx context.Context) (LoginOptions, error) {
	var out LoginOptionsEnvelope
	if err := c.doJSON(ctx, http.MethodGet, "/auth/login-options", false, nil, &out); err != nil {
		return LoginOptions{}, err
	}
	return out.Data, nil
}

func (c *Client) PasswordLogin(ctx context.Context, req PasswordLoginRequest) (AuthResult, error) {
	var out AuthResultEnvelope
	if err := c.doJSON(ctx, http.MethodPost, "/auth/login", false, req, &out); err != nil {
		return AuthResult{}, err
	}
	return out.Data, nil
}

func (c *Client) RefreshAuthSession(ctx context.Context, req *RefreshRequest) (AuthResult, error) {
	var out AuthResultEnvelope
	if err := c.doJSON(ctx, http.MethodPost, "/auth/refresh", false, req, &out); err != nil {
		return AuthResult{}, err
	}
	return out.Data, nil
}

func (c *Client) ExchangeOIDCCode(ctx context.Context, req OIDCExchangeRequest) (AuthResult, error) {
	var out AuthResultEnvelope
	if err := c.doJSON(ctx, http.MethodPost, "/auth/oidc/exchange", false, req, &out); err != nil {
		return AuthResult{}, err
	}
	return out.Data, nil
}

func (c *Client) GetCurrentPrincipal(ctx context.Context) (Principal, error) {
	var out PrincipalEnvelope
	if err := c.doJSON(ctx, http.MethodGet, "/auth/me", true, nil, &out); err != nil {
		return Principal{}, err
	}
	return out.Data, nil
}

func (c *Client) GetCurrentUserProfile(ctx context.Context) (UserProfile, error) {
	var out UserProfileEnvelope
	if err := c.doJSON(ctx, http.MethodGet, "/auth/profile", true, nil, &out); err != nil {
		return UserProfile{}, err
	}
	return out.Data, nil
}

func (c *Client) GetAuthBootstrap(ctx context.Context) (AuthBootstrapEnvelope_Data, error) {
	var out AuthBootstrapEnvelope
	if err := c.doJSON(ctx, http.MethodGet, "/auth/bootstrap", true, nil, &out); err != nil {
		return AuthBootstrapEnvelope_Data{}, err
	}
	return out.Data, nil
}

func (c *Client) LogoutAuthSession(ctx context.Context, req *RefreshRequest) error {
	return c.doJSON(ctx, http.MethodPost, "/auth/logout", true, req, nil)
}

func (c *Client) IssueStreamTicket(ctx context.Context, req StreamTicketRequest) (StreamTicket, error) {
	var out StreamTicketEnvelope
	if err := c.doJSON(ctx, http.MethodPost, "/auth/stream-ticket", true, req, &out); err != nil {
		return StreamTicket{}, err
	}
	return out.Data, nil
}

func (c *Client) ClaimExecutionTask(ctx context.Context, req ExecutionTaskClaimRequest) (ExecutionTask, error) {
	var out ExecutionTaskEnvelope
	if err := c.doJSON(ctx, http.MethodPost, "/delivery/execution-tasks/claim", true, req, &out); err != nil {
		return ExecutionTask{}, err
	}
	return out.Data, nil
}

func (c *Client) GetExecutionTaskRunnerStatus(ctx context.Context, taskID string) (ExecutionTask, error) {
	var out ExecutionTaskEnvelope
	path := "/delivery/execution-tasks/" + url.PathEscape(strings.TrimSpace(taskID)) + "/runner-status"
	if err := c.doJSON(ctx, http.MethodGet, path, true, nil, &out); err != nil {
		return ExecutionTask{}, err
	}
	return out.Data, nil
}

func (c *Client) RecordExecutionCallback(ctx context.Context, req ExecutionCallbackRequest) (ExecutionTask, error) {
	var out ExecutionTaskEnvelope
	if err := c.doJSON(ctx, http.MethodPost, "/delivery/execution-callbacks", false, req, &out); err != nil {
		return ExecutionTask{}, err
	}
	return out.Data, nil
}

func (c *Client) ClaimDockerOperation(ctx context.Context, req DockerOperationClaimRequest) (DockerOperation, error) {
	var out DockerOperationEnvelope
	if err := c.doJSON(ctx, http.MethodPost, "/docker/operations/claim", true, req, &out); err != nil {
		return DockerOperation{}, err
	}
	return out.Data, nil
}

func (c *Client) GetDockerOperationRunnerStatus(ctx context.Context, operationID string) (DockerOperation, error) {
	var out DockerOperationEnvelope
	path := "/docker/operations/" + url.PathEscape(strings.TrimSpace(operationID)) + "/runner-status"
	if err := c.doJSON(ctx, http.MethodGet, path, true, nil, &out); err != nil {
		return DockerOperation{}, err
	}
	return out.Data, nil
}

func (c *Client) RecordDockerOperationCallback(ctx context.Context, req DockerOperationCallbackRequest) (DockerOperation, error) {
	var out DockerOperationEnvelope
	if err := c.doJSON(ctx, http.MethodPost, "/docker/operation-callbacks", true, req, &out); err != nil {
		return DockerOperation{}, err
	}
	return out.Data, nil
}

func (c *Client) ClaimAgentRun(ctx context.Context, req AgentRunClaimRequest) (AgentRun, error) {
	var out AgentRunEnvelope
	if err := c.doJSON(ctx, http.MethodPost, "/copilot/agent-runs/claim", true, req, &out); err != nil {
		return AgentRun{}, err
	}
	return out.Data, nil
}

func (c *Client) RecordAgentRunCallback(ctx context.Context, req AgentRunCallbackRequest) (AgentRun, error) {
	var out AgentRunEnvelope
	if err := c.doJSON(ctx, http.MethodPost, "/copilot/agent-runs/callback", true, req, &out); err != nil {
		return AgentRun{}, err
	}
	return out.Data, nil
}

func (c *Client) RecordAgentRunToolCall(ctx context.Context, req AgentRunToolCallRequest) (AgentToolCallResult, error) {
	var out AgentToolCallResultEnvelope
	if err := c.doJSON(ctx, http.MethodPost, "/copilot/agent-runs/tool-call", true, req, &out); err != nil {
		return AgentToolCallResult{}, err
	}
	return out.Data, nil
}

func (c *Client) ListComputeTasks(ctx context.Context, params ListComputeTasksParams) (ComputeTaskListEnvelope, error) {
	var out ComputeTaskListEnvelope
	if err := c.doJSON(ctx, http.MethodGet, "/compute/tasks"+computeTaskQuery(params), true, nil, &out); err != nil {
		return ComputeTaskListEnvelope{}, err
	}
	return out, nil
}

func (c *Client) GetComputeTask(ctx context.Context, domain ComputeTaskDomain, taskID string) (ComputeTaskView, error) {
	var out ComputeTaskEnvelope
	if err := c.doJSON(ctx, http.MethodGet, computeTaskPath(domain, taskID), true, nil, &out); err != nil {
		return ComputeTaskView{}, err
	}
	return out.Data, nil
}

func (c *Client) ListComputeTaskLogs(ctx context.Context, domain ComputeTaskDomain, taskID string) ([]ComputeTaskLog, error) {
	var out ComputeTaskLogListEnvelope
	if err := c.doJSON(ctx, http.MethodGet, computeTaskPath(domain, taskID)+"/logs", true, nil, &out); err != nil {
		return nil, err
	}
	return out.Items, nil
}

func (c *Client) CancelComputeTask(ctx context.Context, domain ComputeTaskDomain, taskID string, req ComputeTaskMutationRequest) (ComputeTaskView, error) {
	return c.mutateComputeTask(ctx, domain, taskID, "cancel", req)
}

func (c *Client) RetryComputeTask(ctx context.Context, domain ComputeTaskDomain, taskID string, req ComputeTaskMutationRequest) (ComputeTaskView, error) {
	return c.mutateComputeTask(ctx, domain, taskID, "retry", req)
}

func (c *Client) mutateComputeTask(ctx context.Context, domain ComputeTaskDomain, taskID, action string, req ComputeTaskMutationRequest) (ComputeTaskView, error) {
	var out ComputeTaskEnvelope
	path := computeTaskPath(domain, taskID) + "/" + action
	if err := c.doJSON(ctx, http.MethodPost, path, true, req, &out); err != nil {
		return ComputeTaskView{}, err
	}
	return out.Data, nil
}

func (c *Client) ListAIGatewayAuditLogs(ctx context.Context, params ListAIGatewayAuditLogsParams) ([]AuditLog, error) {
	var out AuditLogListEnvelope
	if err := c.doJSON(ctx, http.MethodGet, "/ai-gateway/audit-logs"+auditLogQuery(params), true, nil, &out); err != nil {
		return nil, err
	}
	return out.Items, nil
}

func (c *Client) ListAIGatewayApprovalRequests(ctx context.Context, params ListAIGatewayApprovalRequestsParams) ([]ApprovalRequest, error) {
	var out ApprovalRequestListEnvelope
	if err := c.doJSON(ctx, http.MethodGet, "/ai-gateway/approval-requests"+approvalRequestQuery(params), true, nil, &out); err != nil {
		return nil, err
	}
	return out.Items, nil
}

func (c *Client) GetAIGatewayApprovalTimeline(ctx context.Context, requestID string) (ApprovalTimeline, error) {
	var out ApprovalTimelineEnvelope
	path := "/ai-gateway/approval-requests/" + url.PathEscape(strings.TrimSpace(requestID)) + "/timeline"
	if err := c.doJSON(ctx, http.MethodGet, path, true, nil, &out); err != nil {
		return ApprovalTimeline{}, err
	}
	return out.Data, nil
}

func (c *Client) DecideAIGatewayApprovalRequest(ctx context.Context, requestID string, action DecideAIGatewayApprovalRequestParamsAction, input *ApprovalDecisionInput) (ApprovalDecisionResult, error) {
	var out ApprovalDecisionResultEnvelope
	path := "/ai-gateway/approval-requests/" + url.PathEscape(strings.TrimSpace(requestID)) + "/" + url.PathEscape(string(action))
	if err := c.doJSON(ctx, http.MethodPost, path, true, input, &out); err != nil {
		return ApprovalDecisionResult{}, err
	}
	return out.Data, nil
}

func (c *Client) GetAIGatewayGovernanceStatus(ctx context.Context, params GetAIGatewayGovernanceStatusParams) (GovernanceStatus, error) {
	var out GovernanceStatusEnvelope
	if err := c.doJSON(ctx, http.MethodGet, "/ai-gateway/governance/status"+governanceStatusQuery(params), true, nil, &out); err != nil {
		return GovernanceStatus{}, err
	}
	return out.Data, nil
}

func (c *Client) ListPersonalAccessTokens(ctx context.Context, params ListPersonalAccessTokensParams) ([]PersonalAccessToken, error) {
	var out PersonalAccessTokenListEnvelope
	if err := c.doJSON(ctx, http.MethodGet, "/ai-gateway/personal-access-tokens"+personalAccessTokenQuery(params), true, nil, &out); err != nil {
		return nil, err
	}
	return out.Items, nil
}

func (c *Client) CreatePersonalAccessToken(ctx context.Context, input PersonalAccessTokenInput) (CreatedPersonalAccessToken, error) {
	var out CreatedPersonalAccessTokenEnvelope
	if err := c.doJSON(ctx, http.MethodPost, "/ai-gateway/personal-access-tokens", true, input, &out); err != nil {
		return CreatedPersonalAccessToken{}, err
	}
	return out.Data, nil
}

func (c *Client) RevokePersonalAccessToken(ctx context.Context, tokenID string) (OperationStatus, error) {
	var out OperationStatus
	path := "/ai-gateway/personal-access-tokens/" + url.PathEscape(strings.TrimSpace(tokenID)) + "/revoke"
	if err := c.doJSON(ctx, http.MethodPost, path, true, nil, &out); err != nil {
		return OperationStatus{}, err
	}
	return out, nil
}

func (c *Client) RotatePersonalAccessToken(ctx context.Context, tokenID string, input *TokenRotationInput) (CreatedPersonalAccessToken, error) {
	var out CreatedPersonalAccessTokenEnvelope
	path := "/ai-gateway/personal-access-tokens/" + url.PathEscape(strings.TrimSpace(tokenID)) + "/rotate"
	if err := c.doJSON(ctx, http.MethodPost, path, true, input, &out); err != nil {
		return CreatedPersonalAccessToken{}, err
	}
	return out.Data, nil
}

func (c *Client) ListServiceAccounts(ctx context.Context) ([]ServiceAccount, error) {
	var out ServiceAccountListEnvelope
	if err := c.doJSON(ctx, http.MethodGet, "/ai-gateway/service-accounts", true, nil, &out); err != nil {
		return nil, err
	}
	return out.Items, nil
}

func (c *Client) CreateServiceAccount(ctx context.Context, input ServiceAccountInput) (ServiceAccount, error) {
	var out ServiceAccountEnvelope
	if err := c.doJSON(ctx, http.MethodPost, "/ai-gateway/service-accounts", true, input, &out); err != nil {
		return ServiceAccount{}, err
	}
	return out.Data, nil
}

func (c *Client) ListServiceAccountTokens(ctx context.Context) ([]ServiceAccountToken, error) {
	var out ServiceAccountTokenListEnvelope
	if err := c.doJSON(ctx, http.MethodGet, "/ai-gateway/service-account-tokens", true, nil, &out); err != nil {
		return nil, err
	}
	return out.Items, nil
}

func (c *Client) CreateServiceAccountToken(ctx context.Context, serviceAccountID string, input ServiceAccountTokenInput) (CreatedServiceAccountToken, error) {
	var out CreatedServiceAccountTokenEnvelope
	path := "/ai-gateway/service-accounts/" + url.PathEscape(strings.TrimSpace(serviceAccountID)) + "/tokens"
	if err := c.doJSON(ctx, http.MethodPost, path, true, input, &out); err != nil {
		return CreatedServiceAccountToken{}, err
	}
	return out.Data, nil
}

func (c *Client) RevokeServiceAccountToken(ctx context.Context, tokenID string) (OperationStatus, error) {
	var out OperationStatus
	path := "/ai-gateway/service-account-tokens/" + url.PathEscape(strings.TrimSpace(tokenID)) + "/revoke"
	if err := c.doJSON(ctx, http.MethodPost, path, true, nil, &out); err != nil {
		return OperationStatus{}, err
	}
	return out, nil
}

func (c *Client) RotateServiceAccountToken(ctx context.Context, tokenID string, input *TokenRotationInput) (CreatedServiceAccountToken, error) {
	var out CreatedServiceAccountTokenEnvelope
	path := "/ai-gateway/service-account-tokens/" + url.PathEscape(strings.TrimSpace(tokenID)) + "/rotate"
	if err := c.doJSON(ctx, http.MethodPost, path, true, input, &out); err != nil {
		return CreatedServiceAccountToken{}, err
	}
	return out.Data, nil
}

func (c *Client) ListMCPCapabilities(ctx context.Context) ([]MCPCapability, error) {
	var out MCPCapabilityListEnvelope
	if err := c.doJSON(ctx, http.MethodGet, "/mcp/capabilities", true, nil, &out); err != nil {
		return nil, err
	}
	return out.Items, nil
}

func (c *Client) ListMarketplacePlugins(ctx context.Context, params ListMarketplacePluginsParams) ([]MarketplacePlugin, error) {
	var out MarketplacePluginListEnvelope
	if err := c.doJSON(ctx, http.MethodGet, "/plugins/marketplace"+marketplacePluginQuery(params), true, nil, &out); err != nil {
		return nil, err
	}
	return out.Items, nil
}

func (c *Client) GetMarketplacePlugin(ctx context.Context, pluginID string) (MarketplacePlugin, error) {
	var out MarketplacePluginEnvelope
	path := "/plugins/marketplace/" + url.PathEscape(strings.TrimSpace(pluginID))
	if err := c.doJSON(ctx, http.MethodGet, path, true, nil, &out); err != nil {
		return MarketplacePlugin{}, err
	}
	return out.Data, nil
}

func (c *Client) ListInstalledPlugins(ctx context.Context) ([]InstalledPlugin, error) {
	var out InstalledPluginListEnvelope
	if err := c.doJSON(ctx, http.MethodGet, "/plugins/installed", true, nil, &out); err != nil {
		return nil, err
	}
	return out.Items, nil
}

func (c *Client) InstallPlugin(ctx context.Context, input PluginInstallRequest) (InstalledPlugin, error) {
	var out InstalledPluginEnvelope
	if err := c.doJSON(ctx, http.MethodPost, "/plugins/install", true, input, &out); err != nil {
		return InstalledPlugin{}, err
	}
	return out.Data, nil
}

func (c *Client) GetInstalledPlugin(ctx context.Context, pluginID string) (InstalledPlugin, error) {
	var out InstalledPluginEnvelope
	path := "/plugins/" + url.PathEscape(strings.TrimSpace(pluginID))
	if err := c.doJSON(ctx, http.MethodGet, path, true, nil, &out); err != nil {
		return InstalledPlugin{}, err
	}
	return out.Data, nil
}

func (c *Client) RemoveInstalledPlugin(ctx context.Context, pluginID string) (OperationStatus, error) {
	var out OperationStatus
	path := "/plugins/" + url.PathEscape(strings.TrimSpace(pluginID))
	if err := c.doJSON(ctx, http.MethodDelete, path, true, nil, &out); err != nil {
		return OperationStatus{}, err
	}
	return out, nil
}

func (c *Client) GetInstalledPluginManifest(ctx context.Context, pluginID string) (PluginManifest, error) {
	var out PluginManifestEnvelope
	path := "/plugins/" + url.PathEscape(strings.TrimSpace(pluginID)) + "/manifest"
	if err := c.doJSON(ctx, http.MethodGet, path, true, nil, &out); err != nil {
		return PluginManifest{}, err
	}
	return out.Data, nil
}

func (c *Client) EnableInstalledPlugin(ctx context.Context, pluginID string) (InstalledPlugin, error) {
	var out InstalledPluginEnvelope
	path := "/plugins/" + url.PathEscape(strings.TrimSpace(pluginID)) + "/enable"
	if err := c.doJSON(ctx, http.MethodPost, path, true, nil, &out); err != nil {
		return InstalledPlugin{}, err
	}
	return out.Data, nil
}

func (c *Client) DisableInstalledPlugin(ctx context.Context, pluginID string) (InstalledPlugin, error) {
	var out InstalledPluginEnvelope
	path := "/plugins/" + url.PathEscape(strings.TrimSpace(pluginID)) + "/disable"
	if err := c.doJSON(ctx, http.MethodPost, path, true, nil, &out); err != nil {
		return InstalledPlugin{}, err
	}
	return out.Data, nil
}

func (c *Client) UpgradeInstalledPlugin(ctx context.Context, pluginID string, input *PluginInstallRequest) (InstalledPlugin, error) {
	var out InstalledPluginEnvelope
	path := "/plugins/" + url.PathEscape(strings.TrimSpace(pluginID)) + "/upgrade"
	if err := c.doJSON(ctx, http.MethodPost, path, true, input, &out); err != nil {
		return InstalledPlugin{}, err
	}
	return out.Data, nil
}

func (c *Client) ConfigureInstalledPlugin(ctx context.Context, pluginID string, input PluginConfigRequest) (InstalledPlugin, error) {
	var out InstalledPluginEnvelope
	path := "/plugins/" + url.PathEscape(strings.TrimSpace(pluginID)) + "/config"
	if err := c.doJSON(ctx, http.MethodPut, path, true, input, &out); err != nil {
		return InstalledPlugin{}, err
	}
	return out.Data, nil
}

func (c *Client) doJSON(ctx context.Context, method, path string, useToken bool, body any, out any) error {
	endpoint, err := c.endpoint(path)
	if err != nil {
		return err
	}
	var reader io.Reader
	if body != nil {
		raw, err := json.Marshal(body)
		if err != nil {
			return err
		}
		reader = bytes.NewReader(raw)
	}
	req, err := http.NewRequestWithContext(ctx, method, endpoint, reader)
	if err != nil {
		return err
	}
	req.Header.Set("Accept", "application/json")
	if c.UserAgent != "" {
		req.Header.Set("User-Agent", c.UserAgent)
	}
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	if useToken && strings.TrimSpace(c.Token) != "" {
		req.Header.Set("Authorization", "Bearer "+strings.TrimSpace(c.Token))
	}
	client := c.HTTPClient
	if client == nil {
		client = http.DefaultClient
	}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	raw, err := io.ReadAll(io.LimitReader(resp.Body, 10<<20))
	if err != nil {
		return err
	}
	if resp.StatusCode >= 400 {
		return fmt.Errorf("%s %s failed: %s: %s", method, path, resp.Status, responseErrorMessage(raw))
	}
	if out == nil || len(raw) == 0 {
		return nil
	}
	if err := json.Unmarshal(raw, out); err != nil {
		return fmt.Errorf("decode response: %w", err)
	}
	return nil
}

func (c *Client) endpoint(path string) (string, error) {
	base := strings.TrimRight(strings.TrimSpace(c.BaseURL), "/")
	if base == "" {
		return "", fmt.Errorf("base URL is required")
	}
	if strings.HasSuffix(base, "/api/v1") {
		base = strings.TrimSuffix(base, "/api/v1")
	}
	path = "/" + strings.TrimLeft(path, "/")
	return base + "/api/v1" + path, nil
}

func responseErrorMessage(raw []byte) string {
	var wrapped ErrorEnvelope
	if json.Unmarshal(raw, &wrapped) == nil && wrapped.Error.Message != "" {
		if wrapped.Error.Code != "" {
			return wrapped.Error.Code + ": " + wrapped.Error.Message
		}
		return wrapped.Error.Message
	}
	return strings.TrimSpace(string(raw))
}

func auditLogQuery(params ListAIGatewayAuditLogsParams) string {
	values := url.Values{}
	addQueryString(values, "actorType", params.ActorType)
	addQueryString(values, "actorId", params.ActorID)
	addQueryString(values, "aiClientId", params.AIClientID)
	addQueryString(values, "skillId", params.SkillID)
	addQueryString(values, "toolName", params.ToolName)
	addQueryString(values, "approvalRequestId", params.ApprovalRequestID)
	addQueryString(values, "riskLevel", string(params.RiskLevel))
	addQueryString(values, "result", params.Result)
	addQueryString(values, "action", params.Action)
	addQueryTime(values, "from", params.From)
	addQueryTime(values, "to", params.To)
	addQueryInt(values, "limit", params.Limit)
	return encodeQuery(values)
}

func approvalRequestQuery(params ListAIGatewayApprovalRequestsParams) string {
	values := url.Values{}
	addQueryString(values, "id", params.ID)
	addQueryString(values, "status", params.Status)
	addQueryString(values, "actorType", params.ActorType)
	addQueryString(values, "actorId", params.ActorID)
	addQueryString(values, "aiClientId", params.AIClientID)
	addQueryString(values, "skillId", params.SkillID)
	addQueryString(values, "toolName", params.ToolName)
	addQueryString(values, "riskLevel", string(params.RiskLevel))
	addQueryString(values, "strategy", params.Strategy)
	addQueryTime(values, "from", params.From)
	addQueryTime(values, "to", params.To)
	addQueryInt(values, "limit", params.Limit)
	return encodeQuery(values)
}

func governanceStatusQuery(params GetAIGatewayGovernanceStatusParams) string {
	values := url.Values{}
	addQueryInt(values, "windowHours", params.WindowHours)
	return encodeQuery(values)
}

func personalAccessTokenQuery(params ListPersonalAccessTokensParams) string {
	values := url.Values{}
	addQueryString(values, "scope", params.Scope)
	addQueryString(values, "userId", params.UserID)
	return encodeQuery(values)
}

func marketplacePluginQuery(params ListMarketplacePluginsParams) string {
	values := url.Values{}
	addQueryString(values, "q", params.Q)
	addQueryString(values, "type", params.Type)
	addQueryString(values, "publisher", params.Publisher)
	return encodeQuery(values)
}

func computeTaskPath(domain ComputeTaskDomain, taskID string) string {
	return "/compute/tasks/" + url.PathEscape(string(domain)) + "/" + url.PathEscape(strings.TrimSpace(taskID))
}

func computeTaskQuery(params ListComputeTasksParams) string {
	values := url.Values{}
	addQueryString(values, "domain", string(params.Domain))
	addQueryString(values, "providerKey", params.ProviderKey)
	addQueryString(values, "status", string(params.Status))
	addQueryString(values, "category", string(params.Category))
	addQueryString(values, "resourceKind", params.ResourceKind)
	addQueryString(values, "resourceId", params.ResourceID)
	addQueryString(values, "cursor", string(params.Cursor))
	addQueryInt(values, "limit", int(params.Limit))
	return encodeQuery(values)
}

func addQueryString(values url.Values, key, value string) {
	if value != "" {
		values.Set(key, value)
	}
}

func addQueryInt(values url.Values, key string, value int) {
	if value != 0 {
		values.Set(key, strconv.Itoa(value))
	}
}

func addQueryTime(values url.Values, key string, value time.Time) {
	if !value.IsZero() {
		values.Set(key, value.Format(time.RFC3339Nano))
	}
}

func encodeQuery(values url.Values) string {
	if len(values) == 0 {
		return ""
	}
	return "?" + values.Encode()
}
