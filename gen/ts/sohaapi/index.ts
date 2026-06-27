// Code generated from OpenSoha contracts by openapi-typescript. DO NOT EDIT.

export interface paths {
    "/healthz": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get: operations["getHealthz"];
        put?: never;
        post?: never;
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/readyz": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get: operations["getReadyz"];
        put?: never;
        post?: never;
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/auth/providers": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get: operations["listAuthProviders"];
        put?: never;
        post?: never;
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/auth/login-options": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get: operations["getAuthLoginOptions"];
        put?: never;
        post?: never;
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/auth/login": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get?: never;
        put?: never;
        post: operations["passwordLogin"];
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/auth/refresh": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get?: never;
        put?: never;
        post: operations["refreshAuthSession"];
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/auth/oidc/login": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get: operations["beginOIDCLogin"];
        put?: never;
        post?: never;
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/auth/oidc/callback": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get: operations["handleOIDCCallback"];
        put?: never;
        post?: never;
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/auth/oidc/exchange": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get?: never;
        put?: never;
        post: operations["exchangeOIDCCode"];
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/auth/login/{providerID}/start": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get: operations["beginProviderLogin"];
        put?: never;
        post?: never;
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/auth/login/{providerID}/callback": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get: operations["handleProviderCallback"];
        put?: never;
        post?: never;
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/auth/me": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get: operations["getCurrentPrincipal"];
        put?: never;
        post?: never;
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/auth/profile": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get: operations["getCurrentUserProfile"];
        put?: never;
        post?: never;
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/auth/bootstrap": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get: operations["getAuthBootstrap"];
        put?: never;
        post?: never;
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/auth/logout": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get?: never;
        put?: never;
        post: operations["logoutAuthSession"];
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/auth/stream-ticket": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get?: never;
        put?: never;
        post: operations["issueStreamTicket"];
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/settings/ai": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get: operations["getAISettings"];
        put?: never;
        post?: never;
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/settings/ai/workbench-model": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get?: never;
        put: operations["updateAIWorkbenchModelSettings"];
        post?: never;
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/settings/ai/skills": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get?: never;
        put: operations["updateAISkillsRegistry"];
        post?: never;
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/delivery/execution-tasks/claim": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get?: never;
        put?: never;
        post: operations["claimExecutionTask"];
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/delivery/execution-tasks/{taskID}/runner-status": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get: operations["getExecutionTaskRunnerStatus"];
        put?: never;
        post?: never;
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/delivery/execution-callbacks": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get?: never;
        put?: never;
        post: operations["recordExecutionCallback"];
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/docker/operations/claim": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get?: never;
        put?: never;
        post: operations["claimDockerOperation"];
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/docker/operations/{operationID}/runner-status": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get: operations["getDockerOperationRunnerStatus"];
        put?: never;
        post?: never;
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/docker/operation-callbacks": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get?: never;
        put?: never;
        post: operations["recordDockerOperationCallback"];
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/copilot/agent-runs/claim": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get?: never;
        put?: never;
        post: operations["claimAgentRun"];
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/copilot/agent-runs/callback": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get?: never;
        put?: never;
        post: operations["recordAgentRunCallback"];
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/copilot/agent-runs/tool-call": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get?: never;
        put?: never;
        post: operations["recordAgentRunToolCall"];
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/copilot/agent-runs": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get: operations["listAgentRuns"];
        put?: never;
        post?: never;
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/clusters/capabilities": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get: operations["listClusterCapabilityMatrix"];
        put?: never;
        post?: never;
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/ai-gateway/capabilities": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get: operations["getAIGatewayCapabilities"];
        put?: never;
        post?: never;
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/ai-gateway/tools/{toolName}/invoke": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get?: never;
        put?: never;
        post: operations["invokeAIGatewayTool"];
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/ai-gateway/resources/read": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get?: never;
        put?: never;
        post: operations["readAIGatewayResource"];
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/ai-gateway/prompts/get": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get?: never;
        put?: never;
        post: operations["getAIGatewayPrompt"];
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/ai-gateway/personal-access-tokens": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get: operations["listPersonalAccessTokens"];
        put?: never;
        post: operations["createPersonalAccessToken"];
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/ai-gateway/personal-access-tokens/{tokenID}/revoke": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get?: never;
        put?: never;
        post: operations["revokePersonalAccessToken"];
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/ai-gateway/personal-access-tokens/{tokenID}/rotate": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get?: never;
        put?: never;
        post: operations["rotatePersonalAccessToken"];
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/ai-gateway/service-accounts": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get: operations["listServiceAccounts"];
        put?: never;
        post: operations["createServiceAccount"];
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/ai-gateway/service-account-tokens": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get: operations["listServiceAccountTokens"];
        put?: never;
        post?: never;
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/ai-gateway/service-accounts/{serviceAccountID}/tokens": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get?: never;
        put?: never;
        post: operations["createServiceAccountToken"];
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/ai-gateway/service-account-tokens/{tokenID}/revoke": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get?: never;
        put?: never;
        post: operations["revokeServiceAccountToken"];
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/ai-gateway/service-account-tokens/{tokenID}/rotate": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get?: never;
        put?: never;
        post: operations["rotateServiceAccountToken"];
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/ai-gateway/audit-logs": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get: operations["listAIGatewayAuditLogs"];
        put?: never;
        post?: never;
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/ai-gateway/approval-requests": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get: operations["listAIGatewayApprovalRequests"];
        put?: never;
        post?: never;
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/ai-gateway/approval-requests/{requestID}/timeline": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get: operations["getAIGatewayApprovalTimeline"];
        put?: never;
        post?: never;
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/ai-gateway/approval-requests/{requestID}/{action}": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get?: never;
        put?: never;
        post: operations["decideAIGatewayApprovalRequest"];
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/ai-gateway/governance/status": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get: operations["getAIGatewayGovernanceStatus"];
        put?: never;
        post?: never;
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/ai-gateway/relay/upstreams": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get: operations["listAIGatewayRelayUpstreams"];
        put?: never;
        post: operations["createAIGatewayRelayUpstream"];
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/ai-gateway/relay/upstreams/{upstreamID}": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get?: never;
        put: operations["updateAIGatewayRelayUpstream"];
        post?: never;
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/ai-gateway/relay/upstreams/health-checks/run": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get?: never;
        put?: never;
        post: operations["runAIGatewayRelayHealthChecks"];
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/ai-gateway/relay/upstreams/{upstreamID}/test": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get?: never;
        put?: never;
        post: operations["testAIGatewayRelayUpstream"];
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/ai-gateway/relay/model-routes": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get: operations["listAIGatewayRelayModelRoutes"];
        put?: never;
        post: operations["createAIGatewayRelayModelRoute"];
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/ai-gateway/relay/model-routes/{routeID}": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get?: never;
        put: operations["updateAIGatewayRelayModelRoute"];
        post?: never;
        delete: operations["deleteAIGatewayRelayModelRoute"];
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/ai-gateway/relay/model-calls": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get: operations["listAIGatewayRelayModelCalls"];
        put?: never;
        post?: never;
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/ai-gateway/relay/metrics": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get: operations["getAIGatewayRelayMetrics"];
        put?: never;
        post?: never;
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/ai-gateway/relay/cache/stats": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get: operations["getAIGatewayRelayCacheStats"];
        put?: never;
        post?: never;
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/ai-gateway/relay/cache/purge": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get?: never;
        put?: never;
        post: operations["purgeAIGatewayRelayCache"];
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/ai-gateway/llm/openai/v1/models": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get: operations["listAIGatewayOpenAIModels"];
        put?: never;
        post?: never;
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/ai-gateway/llm/openai/v1/chat/completions": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get?: never;
        put?: never;
        post: operations["createAIGatewayOpenAIChatCompletion"];
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/ai-gateway/llm/openai/v1/responses": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get?: never;
        put?: never;
        post: operations["createAIGatewayOpenAIResponse"];
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/ai-gateway/llm/openai/v1/embeddings": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get?: never;
        put?: never;
        post: operations["createAIGatewayOpenAIEmbedding"];
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/ai-gateway/llm/openai/v1/images/generations": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get?: never;
        put?: never;
        post: operations["createAIGatewayOpenAIImageGeneration"];
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/ai-gateway/llm/openai/v1/images/edits": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get?: never;
        put?: never;
        post: operations["createAIGatewayOpenAIImageEdit"];
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/ai-gateway/llm/openai/v1/images/variations": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get?: never;
        put?: never;
        post: operations["createAIGatewayOpenAIImageVariation"];
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/ai-gateway/llm/openai/v1/audio/speech": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get?: never;
        put?: never;
        post: operations["createAIGatewayOpenAIAudioSpeech"];
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/ai-gateway/llm/openai/v1/audio/transcriptions": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get?: never;
        put?: never;
        post: operations["createAIGatewayOpenAIAudioTranscription"];
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/ai-gateway/llm/openai/v1/audio/translations": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get?: never;
        put?: never;
        post: operations["createAIGatewayOpenAIAudioTranslation"];
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/ai-gateway/llm/openai/v1/realtime": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get: operations["connectAIGatewayOpenAIRealtime"];
        put?: never;
        post?: never;
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/ai-gateway/llm/{openaiCompatibleProvider}/v1/models": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get: operations["listAIGatewayOpenAICompatibleProviderModels"];
        put?: never;
        post?: never;
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/ai-gateway/llm/{openaiCompatibleProvider}/v1/chat/completions": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get?: never;
        put?: never;
        post: operations["createAIGatewayOpenAICompatibleProviderChatCompletion"];
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/ai-gateway/llm/{openaiCompatibleProvider}/v1/responses": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get?: never;
        put?: never;
        post: operations["createAIGatewayOpenAICompatibleProviderResponse"];
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/ai-gateway/llm/{openaiCompatibleProvider}/v1/embeddings": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get?: never;
        put?: never;
        post: operations["createAIGatewayOpenAICompatibleProviderEmbedding"];
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/ai-gateway/llm/{openaiCompatibleProvider}/v1/images/generations": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get?: never;
        put?: never;
        post: operations["createAIGatewayOpenAICompatibleProviderImageGeneration"];
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/ai-gateway/llm/{openaiCompatibleProvider}/v1/images/edits": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get?: never;
        put?: never;
        post: operations["createAIGatewayOpenAICompatibleProviderImageEdit"];
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/ai-gateway/llm/{openaiCompatibleProvider}/v1/images/variations": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get?: never;
        put?: never;
        post: operations["createAIGatewayOpenAICompatibleProviderImageVariation"];
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/ai-gateway/llm/{openaiCompatibleProvider}/v1/audio/speech": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get?: never;
        put?: never;
        post: operations["createAIGatewayOpenAICompatibleProviderAudioSpeech"];
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/ai-gateway/llm/{openaiCompatibleProvider}/v1/audio/transcriptions": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get?: never;
        put?: never;
        post: operations["createAIGatewayOpenAICompatibleProviderAudioTranscription"];
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/ai-gateway/llm/{openaiCompatibleProvider}/v1/audio/translations": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get?: never;
        put?: never;
        post: operations["createAIGatewayOpenAICompatibleProviderAudioTranslation"];
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/ai-gateway/llm/gemini/v1beta/models": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get: operations["listAIGatewayGeminiModels"];
        put?: never;
        post?: never;
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/ai-gateway/llm/gemini/v1beta/models/{geminiModel}:generateContent": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get?: never;
        put?: never;
        post: operations["createAIGatewayGeminiGenerateContent"];
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/ai-gateway/llm/gemini/v1beta/models/{geminiModel}:streamGenerateContent": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get?: never;
        put?: never;
        post: operations["createAIGatewayGeminiStreamGenerateContent"];
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/ai-gateway/llm/gemini/v1beta/interactions": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get?: never;
        put?: never;
        post: operations["createAIGatewayGeminiInteractions"];
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/ai-gateway/llm/cohere/v2/rerank": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get?: never;
        put?: never;
        post: operations["createAIGatewayCohereRerank"];
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/ai-gateway/llm/anthropic/v1/models": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get: operations["listAIGatewayAnthropicModels"];
        put?: never;
        post?: never;
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/ai-gateway/llm/anthropic/v1/messages": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get?: never;
        put?: never;
        post: operations["createAIGatewayAnthropicMessage"];
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/mcp/capabilities": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get: operations["listMCPCapabilities"];
        put?: never;
        post?: never;
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/plugins/marketplace": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get: operations["listMarketplacePlugins"];
        put?: never;
        post?: never;
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/plugins/marketplace/{pluginID}": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get: operations["getMarketplacePlugin"];
        put?: never;
        post?: never;
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/plugins/installed": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get: operations["listInstalledPlugins"];
        put?: never;
        post?: never;
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/plugins/install": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get?: never;
        put?: never;
        post: operations["installPlugin"];
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/plugins/{pluginID}": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get: operations["getInstalledPlugin"];
        put?: never;
        post?: never;
        delete: operations["removeInstalledPlugin"];
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/plugins/{pluginID}/manifest": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get: operations["getInstalledPluginManifest"];
        put?: never;
        post?: never;
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/plugins/{pluginID}/enable": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get?: never;
        put?: never;
        post: operations["enableInstalledPlugin"];
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/plugins/{pluginID}/disable": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get?: never;
        put?: never;
        post: operations["disableInstalledPlugin"];
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/plugins/{pluginID}/upgrade": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get?: never;
        put?: never;
        post: operations["upgradeInstalledPlugin"];
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/plugins/{pluginID}/config": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get?: never;
        put: operations["configureInstalledPlugin"];
        post?: never;
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
}
export type webhooks = Record<string, never>;
export interface components {
    schemas: {
        AnyValue: unknown;
        GenericObject: {
            [key: string]: unknown;
        };
        GenericDataEnvelope: {
            data: components["schemas"]["AnyValue"];
        };
        GenericItemsEnvelope: {
            items: components["schemas"]["AnyValue"][];
        };
        ErrorEnvelope: {
            error: {
                code: string;
                message: string;
                request_id?: string;
            };
        };
        AIWorkbenchModelSettings: {
            /** @description Public model name exposed by AI Gateway model routes and used as the Workbench default. */
            defaultPublicModel?: string;
            /** @description Stable AI Gateway model-route ID used as the Workbench default when set. */
            defaultRouteId?: string;
            /**
             * @description Workbench relay endpoint. Supported values are chat/completions, responses, and messages.
             * @default chat/completions
             */
            defaultEndpoint: string;
            enabled: boolean;
        };
        AISkillSettings: {
            id: string;
            name: string;
            category?: string;
            ownerModule?: string;
            description?: string;
            capabilityRefs?: string[];
            blueprintRefs?: string[];
            inputSchema?: {
                [key: string]: unknown;
            };
            outputSchema?: {
                [key: string]: unknown;
            };
            scopeRules?: string[];
            enabled: boolean;
            scopes?: string[];
        };
        AISettings: {
            workbenchModel: components["schemas"]["AIWorkbenchModelSettings"];
            skillsRegistry: components["schemas"]["AISkillSettings"][];
        };
        AISettingsEnvelope: {
            data: components["schemas"]["AISettings"];
        };
        UpdateAIWorkbenchModelRequest: {
            workbenchModel: components["schemas"]["AIWorkbenchModelSettings"];
        };
        UpdateAISkillsRequest: {
            skillsRegistry: components["schemas"]["AISkillSettings"][];
        };
        Principal: {
            userId: string;
            userName: string;
            email: string;
            roles: string[];
            teams: string[];
            projects: string[];
            tags: string[];
            permissionKeys?: string[];
        } & {
            [key: string]: unknown;
        };
        UserProfile: {
            userId: string;
            username: string;
            displayName: string;
            email: string;
            phone?: string;
            status: string;
            roles: string[];
            teams: string[];
            projects: string[];
            tags: string[];
            identities?: {
                [key: string]: unknown;
            }[];
            sessions?: {
                [key: string]: unknown;
            }[];
            /** Format: date-time */
            lastLoginAt?: string;
        } & {
            [key: string]: unknown;
        };
        TokenSet: {
            accessToken: string;
            refreshToken: string;
            tokenType: string;
            /** Format: int64 */
            expiresIn: number;
            /** Format: date-time */
            expiresAt: string;
        };
        AuthResult: {
            user: components["schemas"]["Principal"];
            tokens: components["schemas"]["TokenSet"];
        };
        AuthProvider: {
            type: string;
            id?: string;
            name: string;
            enabled: boolean;
            loginUrl?: string;
        };
        LoginOptions: {
            verification: {
                sliderEnabled: boolean;
            };
        };
        PasswordLoginRequest: {
            login: string;
            password: string;
        };
        RefreshRequest: {
            refreshToken?: string;
        };
        OIDCExchangeRequest: {
            code: string;
        };
        StreamTicketRequest: {
            path: string;
        };
        StreamTicket: {
            ticket: string;
            /** Format: date-time */
            expiresAt: string;
        };
        PrincipalEnvelope: {
            data: components["schemas"]["Principal"];
        };
        UserProfileEnvelope: {
            data: components["schemas"]["UserProfile"];
        };
        AuthResultEnvelope: {
            data: components["schemas"]["AuthResult"];
        };
        AuthProviderListEnvelope: {
            items: components["schemas"]["AuthProvider"][];
        };
        LoginOptionsEnvelope: {
            data: components["schemas"]["LoginOptions"];
        };
        AuthBootstrapEnvelope: {
            data: {
                user: components["schemas"]["Principal"];
                currentUser: components["schemas"]["Principal"];
                permissionSnapshot: {
                    [key: string]: unknown;
                };
                branding: {
                    [key: string]: unknown;
                };
            } & {
                [key: string]: unknown;
            };
        };
        StreamTicketEnvelope: {
            data: components["schemas"]["StreamTicket"];
        };
        ExecutionTask: {
            id: string;
            applicationId: string;
            applicationEnvironmentId: string;
            taskKind: string;
            providerKind: string;
            status: string;
            callbackToken: string;
            payload: {
                [key: string]: unknown;
            };
        } & {
            [key: string]: unknown;
        };
        ExecutionTaskClaimRequest: {
            agentId: string;
            providerKinds?: string[];
            runtimeEndpoint?: string;
        };
        ExecutionCallbackRequest: {
            callbackToken: string;
            status: string;
            payload: {
                [key: string]: unknown;
            };
        };
        DockerOperation: {
            id: string;
            hostId?: string;
            projectId?: string;
            serviceId?: string;
            operationKind: string;
            status: string;
            claimedByWorkerId?: string;
            payload: {
                [key: string]: unknown;
            };
            timeoutSeconds: number;
        } & {
            [key: string]: unknown;
        };
        DockerOperationClaimRequest: {
            workerId: string;
            agentId: string;
            hostIds?: string[];
            operationKinds?: string[];
        };
        DockerOperationCallbackRequest: {
            operationId: string;
            workerId: string;
            status: string;
            payload: {
                [key: string]: unknown;
            };
            logs: string[];
        };
        AgentRun: {
            id: string;
            providerId: string;
            providerKind: string;
            capabilityId: string;
            skillIds?: string[];
            sessionId?: string;
            status: string;
            scope?: {
                [key: string]: unknown;
            };
            toolset?: {
                [key: string]: unknown;
            };
            toolBindings?: {
                [key: string]: unknown;
            }[];
            skillBindings?: {
                [key: string]: unknown;
            }[];
            input?: {
                [key: string]: unknown;
            };
            output?: {
                [key: string]: unknown;
            };
            callbackToken: string;
            timeoutSeconds: number;
            analysisArtifacts?: {
                [key: string]: unknown;
            }[];
        } & {
            [key: string]: unknown;
        };
        AgentRunClaimRequest: {
            agentId: string;
            providerIds?: string[];
            kinds?: string[];
        };
        AgentRunCallbackRequest: {
            runId: string;
            callbackToken: string;
            agentId: string;
            status: string;
            payload: {
                [key: string]: unknown;
            };
            toolExecutions?: {
                [key: string]: unknown;
            }[];
            analysisArtifacts?: {
                [key: string]: unknown;
            }[];
            externalRunId?: string;
            errorMessage?: string;
        };
        AgentRunToolCallRequest: {
            runId: string;
            callbackToken: string;
            agentId: string;
            toolBindingId?: string;
            adapterId?: string;
            toolName?: string;
            input?: {
                [key: string]: unknown;
            };
        };
        AgentToolCallResult: {
            runId: string;
            toolExecution: {
                [key: string]: unknown;
            };
            output?: {
                [key: string]: unknown;
            };
        } & {
            [key: string]: unknown;
        };
        ExecutionTaskEnvelope: {
            data: components["schemas"]["ExecutionTask"];
        };
        DockerOperationEnvelope: {
            data: components["schemas"]["DockerOperation"];
        };
        AgentRunEnvelope: {
            data: components["schemas"]["AgentRun"];
        };
        AgentToolCallResultEnvelope: {
            data: components["schemas"]["AgentToolCallResult"];
        };
        /** @enum {string} */
        RiskLevel: "read" | "analyze" | "mutate" | "execute" | "high";
        /** @enum {string} */
        ClusterCapabilityStatus: "available" | "partial" | "unsupported";
        ClusterCapabilityModeSupport: {
            status: components["schemas"]["ClusterCapabilityStatus"];
            reason?: string;
            notes?: string[];
        };
        ClusterCapabilityMatrixEntry: {
            key: string;
            label: string;
            category: string;
            requiredScopes?: string[];
            riskLevel: components["schemas"]["RiskLevel"];
            requiresApproval: boolean;
            docsUrl?: string;
            direct: components["schemas"]["ClusterCapabilityModeSupport"];
            agent: components["schemas"]["ClusterCapabilityModeSupport"];
        };
        ClusterCapabilityMatrixEnvelope: {
            items: components["schemas"]["ClusterCapabilityMatrixEntry"][];
        };
        JSONSchema: {
            [key: string]: unknown;
        };
        ToolCapability: {
            name: string;
            title: string;
            description: string;
            domain: string;
            action: string;
            riskLevel: components["schemas"]["RiskLevel"];
            permissionKeys: string[];
            requiredScopes?: string[];
            mcpAdapterId?: string;
            mcpToolName?: string;
            requiresApproval: boolean;
            inputSchema?: components["schemas"]["JSONSchema"];
            outputSchema?: components["schemas"]["JSONSchema"];
        };
        ResourceCapability: {
            name: string;
            description: string;
            permissionKeys: string[];
            requiredScopes?: string[];
            contextSchema?: components["schemas"]["JSONSchema"];
        };
        PromptCapability: {
            name: string;
            description: string;
            permissionKeys: string[];
            requiredScopes?: string[];
            argumentSchema?: components["schemas"]["JSONSchema"];
            contextSchema?: components["schemas"]["JSONSchema"];
        };
        SkillCapability: {
            id: string;
            name: string;
            category: string;
            description: string;
            capabilityRefs?: string[];
            permissionKeys?: string[];
            requiredScopes?: string[];
        };
        CallerContext: {
            identityMode: string;
            aiClientId?: string;
            aiClientName?: string;
            skillId?: string;
            tokenId?: string;
            sessionId?: string;
            subjectType?: string;
            subjectId?: string;
            source?: string;
        };
        AIGatewayManifest: {
            name: string;
            version: string;
            /** Format: date-time */
            generatedAt: string;
            principal: components["schemas"]["Principal"];
            caller: components["schemas"]["CallerContext"];
            permissionKeys: string[];
            tools: components["schemas"]["ToolCapability"][];
            resources?: components["schemas"]["ResourceCapability"][];
            prompts?: components["schemas"]["PromptCapability"][];
            skills?: components["schemas"]["SkillCapability"][];
            summary: {
                toolCount: number;
                resourceCount: number;
                promptCount: number;
                skillCount: number;
                deniedCount: number;
            };
        };
        ToolInvocationRequest: {
            toolName?: string;
            input?: {
                [key: string]: unknown;
            };
            aiClientId?: string;
            aiClientName?: string;
            skillId?: string;
            requestId?: string;
        };
        ToolInvocationResult: {
            toolName: string;
            riskLevel: components["schemas"]["RiskLevel"];
            requiresApproval: boolean;
            result: string;
            output?: components["schemas"]["AnyValue"];
            relatedIds?: {
                [key: string]: unknown;
            };
            audit?: {
                [key: string]: unknown;
            };
        } & {
            [key: string]: unknown;
        };
        ResourceReadRequest: {
            name?: string;
            uri?: string;
            context?: {
                [key: string]: unknown;
            };
            aiClientId?: string;
            aiClientName?: string;
            skillId?: string;
            requestId?: string;
        };
        ResourceReadResult: {
            name: string;
            uri: string;
            mimeType: string;
            text?: string;
            data?: components["schemas"]["AnyValue"];
            relatedIds?: {
                [key: string]: unknown;
            };
            audit?: {
                [key: string]: unknown;
            };
        } & {
            [key: string]: unknown;
        };
        PromptGetRequest: {
            name: string;
            arguments?: {
                [key: string]: unknown;
            };
            context?: {
                [key: string]: unknown;
            };
            aiClientId?: string;
            aiClientName?: string;
            skillId?: string;
            requestId?: string;
        };
        PromptMessage: {
            role: string;
            content: string;
        };
        PromptGetResult: {
            name: string;
            description: string;
            messages: components["schemas"]["PromptMessage"][];
            relatedIds?: {
                [key: string]: unknown;
            };
            audit?: {
                [key: string]: unknown;
            };
        } & {
            [key: string]: unknown;
        };
        AIGatewayManifestEnvelope: {
            data: components["schemas"]["AIGatewayManifest"];
        };
        ToolInvocationResultEnvelope: {
            data: components["schemas"]["ToolInvocationResult"];
        };
        ResourceReadResultEnvelope: {
            data: components["schemas"]["ResourceReadResult"];
        };
        PromptGetResultEnvelope: {
            data: components["schemas"]["PromptGetResult"];
        };
        OperationStatus: {
            status: string;
        };
        LLMTokenMetadata: {
            /** @enum {string} */
            purpose?: "llm-relay";
            allowedModels?: string[];
            allowedProviderKinds?: ("openai" | "anthropic" | "openai-compatible" | "deepseek" | "qwen" | "openrouter" | "azure-openai" | "gemini" | "cohere")[];
            allowedUpstreamIds?: string[];
            allowedIPCIDRs?: string[];
            /** @description Team or organization IDs allowed to use this relay token. Empty means unrestricted unless deniedTeams matches. */
            allowedTeams?: string[];
            /** @description Team or organization IDs denied for this relay token. Deny rules take precedence over allowedTeams. */
            deniedTeams?: string[];
            rateLimitProfileId?: string;
            /** @description Allows X-Soha-Route-Trace response headers when set by a relay manager. */
            allowRouteTrace?: boolean;
            /** @description Allows X-Soha-Upstream-ID debug selection, still constrained by allowedUpstreamIds, when set by a relay manager. */
            allowUpstreamSelection?: boolean;
        } & {
            [key: string]: unknown;
        };
        LLMUpstream: {
            id: string;
            name: string;
            /** @enum {string} */
            providerKind: "openai" | "anthropic" | "openai-compatible" | "deepseek" | "qwen" | "openrouter" | "azure-openai" | "gemini" | "cohere";
            /** Format: uri */
            baseUrl: string;
            /** @description Display-only prefix for the encrypted upstream API key. The full key is never returned. */
            apiKeyPrefix?: string;
            /** @enum {string} */
            status: "active" | "disabled" | "degraded";
            /** @description Lower values are selected first. */
            priority: number;
            weight: number;
            timeoutSeconds: number;
            streamTimeoutSeconds: number;
            maxConcurrency: number;
            supportedModels: string[];
            /** @description Non-sensitive upstream headers only. */
            defaultHeaders?: {
                [key: string]: unknown;
            };
            proxyUrl?: string;
            health?: {
                [key: string]: unknown;
            };
            metadata?: {
                [key: string]: unknown;
            };
            createdBy: string;
            /** Format: date-time */
            createdAt: string;
            /** Format: date-time */
            updatedAt: string;
        };
        LLMUpstreamInput: {
            id?: string;
            name: string;
            /** @enum {string} */
            providerKind: "openai" | "anthropic" | "openai-compatible" | "deepseek" | "qwen" | "openrouter" | "azure-openai" | "gemini" | "cohere";
            /** Format: uri */
            baseUrl: string;
            /** @description Upstream provider key. It must be encrypted at rest and never returned by read APIs. */
            apiKey?: string;
            /** @enum {string} */
            status?: "active" | "disabled" | "degraded";
            priority?: number;
            weight?: number;
            timeoutSeconds?: number;
            streamTimeoutSeconds?: number;
            maxConcurrency?: number;
            supportedModels?: string[];
            defaultHeaders?: {
                [key: string]: unknown;
            };
            proxyUrl?: string;
            health?: {
                [key: string]: unknown;
            };
            metadata?: {
                [key: string]: unknown;
            };
        };
        LLMUpstreamEnvelope: {
            data: components["schemas"]["LLMUpstream"];
        };
        LLMUpstreamListEnvelope: {
            items: components["schemas"]["LLMUpstream"][];
        };
        LLMUpstreamTestRequest: {
            model?: string;
            /** @enum {string} */
            endpoint?: "models" | "chat.completions" | "responses" | "messages" | "rerank" | "interactions" | "images.generations" | "images.edits" | "images.variations" | "audio.speech" | "audio.transcriptions" | "audio.translations";
            timeoutSeconds?: number;
        };
        LLMUpstreamTestResult: {
            /** @enum {string} */
            status: "success" | "failure";
            /** @enum {string} */
            providerKind: "openai" | "anthropic" | "openai-compatible" | "deepseek" | "qwen" | "openrouter" | "azure-openai" | "gemini" | "cohere";
            upstreamStatus?: number;
            /** Format: int64 */
            durationMs: number;
            modelCount?: number;
            errorCode?: string;
            /** @description Redacted error summary. */
            errorMessage?: string;
            /** Format: date-time */
            checkedAt?: string;
        };
        LLMUpstreamTestResultEnvelope: {
            data: components["schemas"]["LLMUpstreamTestResult"];
        };
        LLMRelayHealthCheckRun: {
            /** Format: date-time */
            checkedAt: string;
            total: number;
            checked: number;
            skipped: number;
            healthy: number;
            degraded: number;
            recovered: number;
            failed: number;
            results?: components["schemas"]["LLMUpstreamTestResult"][];
        };
        LLMRelayHealthCheckRunEnvelope: {
            data: components["schemas"]["LLMRelayHealthCheckRun"];
        };
        LLMModelRoute: {
            id: string;
            publicModel: string;
            /** @enum {string} */
            providerKind?: "openai" | "anthropic" | "openai-compatible" | "deepseek" | "qwen" | "openrouter" | "azure-openai" | "gemini" | "cohere";
            upstreamId?: string;
            upstreamModel: string;
            routeGroup?: string;
            priority: number;
            weight: number;
            enabled: boolean;
            /**
             * @description Optional relay format conversion policy. Set `mode: convert` and `targetProviderKind` to `openai` or `anthropic` to enable text-only non-streaming OpenAI Chat Completions <-> Anthropic Messages conversion for this route. Streaming, tool/function, multimodal, audio, image, and file payloads are not converted.
             * @example {
             *       "mode": "convert",
             *       "targetProviderKind": "anthropic"
             *     }
             */
            transformPolicy?: {
                [key: string]: unknown;
            };
            fallbackPolicy?: {
                [key: string]: unknown;
            };
            cachePolicy?: {
                [key: string]: unknown;
            };
            rateLimitProfileId?: string;
            metadata?: {
                [key: string]: unknown;
            };
            /** Format: date-time */
            createdAt: string;
            /** Format: date-time */
            updatedAt: string;
        };
        LLMModelRouteInput: {
            id?: string;
            publicModel: string;
            /** @enum {string} */
            providerKind?: "openai" | "anthropic" | "openai-compatible" | "deepseek" | "qwen" | "openrouter" | "azure-openai" | "gemini" | "cohere";
            upstreamId?: string;
            upstreamModel: string;
            routeGroup?: string;
            priority?: number;
            weight?: number;
            enabled?: boolean;
            /**
             * @description Optional relay format conversion policy. Set `mode: convert` and `targetProviderKind` to `openai` or `anthropic` to enable text-only non-streaming OpenAI Chat Completions <-> Anthropic Messages conversion for this route. Streaming, tool/function, multimodal, audio, image, and file payloads are not converted.
             * @example {
             *       "mode": "convert",
             *       "targetProviderKind": "anthropic"
             *     }
             */
            transformPolicy?: {
                [key: string]: unknown;
            };
            fallbackPolicy?: {
                [key: string]: unknown;
            };
            cachePolicy?: {
                [key: string]: unknown;
            };
            rateLimitProfileId?: string;
            metadata?: {
                [key: string]: unknown;
            };
        };
        LLMModelRouteEnvelope: {
            data: components["schemas"]["LLMModelRoute"];
        };
        LLMModelRouteListEnvelope: {
            items: components["schemas"]["LLMModelRoute"][];
        };
        LLMCallLog: {
            id: string;
            requestId?: string;
            actorType?: string;
            actorId?: string;
            actorName?: string;
            tokenId?: string;
            tokenPrefix?: string;
            /** @enum {string} */
            tokenKind?: "personal_access_token" | "service_account_token";
            aiClientId?: string;
            publicModel?: string;
            upstreamId?: string;
            upstreamName?: string;
            /** @enum {string} */
            providerKind?: "openai" | "anthropic" | "openai-compatible" | "deepseek" | "qwen" | "openrouter" | "azure-openai" | "gemini" | "cohere";
            upstreamModel?: string;
            endpoint?: string;
            stream: boolean;
            /** @enum {string} */
            status: "success" | "failure" | "cancelled" | "client_cancelled" | "rate_limited" | "policy_denied";
            httpStatus?: number;
            upstreamStatus?: number;
            errorCode?: string;
            /** @description Redacted error summary. */
            errorMessage?: string;
            promptTokens?: number;
            completionTokens?: number;
            totalTokens?: number;
            reasoningTokens?: number;
            cachedReadTokens?: number;
            cachedWriteTokens?: number;
            estimatedTokens: boolean;
            /** Format: int64 */
            ttfbMs?: number;
            /** Format: int64 */
            ttftMs?: number;
            /** Format: int64 */
            durationMs?: number;
            /** Format: int64 */
            inputBytes?: number;
            /** Format: int64 */
            outputBytes?: number;
            /** @enum {string} */
            cacheStatus?: "bypass" | "miss" | "hit" | "stale_hit" | "write" | "write_skipped";
            routeTrace?: {
                [key: string]: unknown;
            };
            sourceIp?: string;
            userAgent?: string;
            /** @description Redacted metadata only; prompt bodies, raw provider payloads, full headers, and credentials are not included. */
            metadata?: {
                [key: string]: unknown;
            };
            /** Format: date-time */
            createdAt: string;
        };
        LLMCallLogListEnvelope: {
            items: components["schemas"]["LLMCallLog"][];
        };
        LLMRelayMetricSeriesPoint: {
            /** Format: date-time */
            timestamp: string;
            requestCount: number;
            successCount?: number;
            failureCount?: number;
            rateLimitedCount?: number;
            policyDeniedCount?: number;
            averageTTFBMs?: number;
            averageTTFTMs?: number;
            averageDurationMs?: number;
            promptTokens?: number;
            completionTokens?: number;
            totalTokens?: number;
            cachedReadTokens?: number;
            cachedWriteTokens?: number;
        };
        LLMRelayMetricBreakdown: {
            key: string;
            requestCount: number;
            successCount?: number;
            failureCount?: number;
            streamCount?: number;
            averageTTFBMs?: number;
            averageTTFTMs?: number;
            averageDurationMs?: number;
            totalTokens?: number;
            cachedReadTokens?: number;
            cachedWriteTokens?: number;
        };
        LLMRelayMetrics: {
            /** Format: date-time */
            generatedAt: string;
            windowHours: number;
            requestCount: number;
            successCount: number;
            failureCount: number;
            streamCount?: number;
            clientCancelledCount?: number;
            rateLimitedCount?: number;
            policyDeniedCount?: number;
            averageTTFBMs?: number;
            averageTTFTMs?: number;
            averageDurationMs?: number;
            promptTokens?: number;
            completionTokens?: number;
            totalTokens?: number;
            reasoningTokens?: number;
            cachedReadTokens?: number;
            cachedWriteTokens?: number;
            byModel?: components["schemas"]["LLMRelayMetricBreakdown"][];
            byUpstream?: components["schemas"]["LLMRelayMetricBreakdown"][];
            series?: components["schemas"]["LLMRelayMetricSeriesPoint"][];
        };
        LLMRelayMetricsEnvelope: {
            data: components["schemas"]["LLMRelayMetrics"];
        };
        LLMRelayCacheStats: {
            /** Format: date-time */
            generatedAt: string;
            windowHours: number;
            responseCacheEnabled: boolean;
            responseCacheHits?: number;
            responseCacheMisses?: number;
            responseCacheWrites?: number;
            responseCacheBypasses?: number;
            providerCachedReadTokens?: number;
            providerCachedWriteTokens?: number;
            byModel?: {
                [key: string]: unknown;
            }[];
            byUpstream?: {
                [key: string]: unknown;
            }[];
        };
        LLMRelayCacheStatsEnvelope: {
            data: components["schemas"]["LLMRelayCacheStats"];
        };
        LLMRelayCachePurgeRequest: {
            publicModel?: string;
            upstreamId?: string;
            routeGroup?: string;
            /** Format: date-time */
            olderThan?: string;
            /** @default false */
            dryRun: boolean;
        };
        LLMRelayCachePurgeResult: {
            status: string;
            purgedCount: number;
            dryRun?: boolean;
        };
        LLMRelayCachePurgeResultEnvelope: {
            data: components["schemas"]["LLMRelayCachePurgeResult"];
        };
        /** @description Native provider-compatible JSON. Unknown fields are preserved and responses are not wrapped in an OpenSoha envelope. */
        NativeProviderObject: {
            [key: string]: unknown;
        };
        OpenAIModel: {
            id: string;
            object: string;
            created?: number;
            owned_by?: string;
        } & {
            [key: string]: unknown;
        };
        OpenAIModelsResponse: {
            /** @enum {string} */
            object: "list";
            data: components["schemas"]["OpenAIModel"][];
        } & {
            [key: string]: unknown;
        };
        OpenAIChatCompletionRequest: {
            model: string;
            messages: {
                [key: string]: unknown;
            }[];
            stream?: boolean;
            stream_options?: {
                [key: string]: unknown;
            };
        } & {
            [key: string]: unknown;
        };
        OpenAIResponseRequest: {
            model: string;
            input?: components["schemas"]["AnyValue"];
            stream?: boolean;
        } & {
            [key: string]: unknown;
        };
        OpenAIEmbeddingRequest: {
            model: string;
            input: components["schemas"]["AnyValue"];
        } & {
            [key: string]: unknown;
        };
        OpenAIImageGenerationRequest: {
            /** @description Public image model ID exposed by Soha; the relay maps it to the configured upstream model. */
            model: string;
            prompt: string;
            n?: number;
            size?: string;
            quality?: string;
            response_format?: string;
            style?: string;
            background?: string;
            output_format?: string;
            output_compression?: number;
            moderation?: string;
            user?: string;
        } & {
            [key: string]: unknown;
        };
        OpenAIImageEditRequest: {
            /** Format: binary */
            image: string;
            /** @description Public image edit model ID exposed by Soha; the relay maps it to the configured upstream model. */
            model: string;
            prompt: string;
            /** Format: binary */
            mask?: string;
            n?: number;
            size?: string;
            quality?: string;
            response_format?: string;
            user?: string;
        } & {
            [key: string]: unknown;
        };
        OpenAIImageVariationRequest: {
            /** Format: binary */
            image: string;
            /** @description Public image variation model ID exposed by Soha; the relay maps it to the configured upstream model. */
            model: string;
            n?: number;
            size?: string;
            response_format?: string;
            user?: string;
        } & {
            [key: string]: unknown;
        };
        OpenAIAudioSpeechRequest: {
            /** @description Public speech model ID exposed by Soha; the relay maps it to the configured upstream model. */
            model: string;
            input: string;
            voice: string;
            response_format?: string;
            speed?: number;
            instructions?: string;
            user?: string;
        } & {
            [key: string]: unknown;
        };
        OpenAIAudioTranscriptionRequest: {
            /** Format: binary */
            file: string;
            /** @description Public transcription model ID exposed by Soha; the relay maps it to the configured upstream model. */
            model: string;
            language?: string;
            prompt?: string;
            response_format?: string;
            temperature?: number;
            timestamp_granularities?: ("word" | "segment")[];
            include?: string[];
        } & {
            [key: string]: unknown;
        };
        OpenAIAudioTranslationRequest: {
            /** Format: binary */
            file: string;
            /** @description Public translation model ID exposed by Soha; the relay maps it to the configured upstream model. */
            model: string;
            prompt?: string;
            response_format?: string;
            temperature?: number;
        } & {
            [key: string]: unknown;
        };
        GeminiModel: {
            name: string;
            version?: string;
            displayName?: string;
            description?: string;
            supportedGenerationMethods?: string[];
        } & {
            [key: string]: unknown;
        };
        GeminiModelsResponse: {
            models: components["schemas"]["GeminiModel"][];
        } & {
            [key: string]: unknown;
        };
        /** @description Native Gemini generateContent request body. The model is taken from the path and mapped by Soha model routes. Gemini text, inlineData, fileData, cachedContent, generationConfig, and safetySettings fields are passed through; multimodal audio/image/file inputs bypass response cache and local token estimation only counts text parts when provider usageMetadata is absent. */
        GeminiGenerateContentRequest: {
            contents?: {
                [key: string]: unknown;
            }[];
            systemInstruction?: {
                [key: string]: unknown;
            };
            generationConfig?: {
                [key: string]: unknown;
            };
            safetySettings?: {
                [key: string]: unknown;
            }[];
        } & {
            [key: string]: unknown;
        };
        /** @description Native Gemini Interactions request body. The public model is supplied in the body and mapped by Soha model routes before forwarding to the upstream `/interactions` endpoint. This first slice supports non-streaming JSON interactions for native image generation; `stream` and `background` modes are rejected by the relay. Multimodal inputs bypass response cache, and provider usage totals are recorded when present. */
        GeminiInteractionsRequest: {
            /** @description Public Gemini model ID exposed by Soha; the relay rewrites it to the configured upstream model. */
            model: string;
            /** @description Native Gemini Interactions input payload. */
            input: unknown;
            /** @description Streaming interactions are not supported by this relay slice. */
            stream?: boolean;
            /** @description Background interactions are not supported by this relay slice. */
            background?: boolean;
        } & {
            [key: string]: unknown;
        };
        CohereRerankRequest: {
            /** @description Public rerank model ID exposed by Soha; the relay maps it to the configured upstream model. */
            model: string;
            query: string;
            documents: (string | {
                [key: string]: unknown;
            })[];
            top_n?: number;
            max_tokens_per_doc?: number;
            rank_fields?: string[];
        } & {
            [key: string]: unknown;
        };
        AnthropicModel: {
            id: string;
            type: string;
            display_name?: string;
            created_at?: string;
        } & {
            [key: string]: unknown;
        };
        AnthropicModelsResponse: {
            data: components["schemas"]["AnthropicModel"][];
            has_more?: boolean;
            first_id?: string;
            last_id?: string;
        } & {
            [key: string]: unknown;
        };
        AnthropicMessagesRequest: {
            model: string;
            max_tokens: number;
            messages: {
                [key: string]: unknown;
            }[];
            system?: components["schemas"]["AnyValue"];
            stream?: boolean;
        } & {
            [key: string]: unknown;
        };
        PluginAssetSnapshot: {
            skills?: string[];
            mcpPresets?: string[];
            connectors?: string[];
            aiProviderAdapters?: string[];
            agentProfiles?: string[];
            gatewayPolicyPacks?: string[];
        };
        PluginCompatibility: {
            soha?: string;
            contracts?: string;
        } & {
            [key: string]: unknown;
        };
        PluginCapabilityRequest: {
            tools?: string[];
            resources?: string[];
            prompts?: string[];
            skills?: string[];
        };
        PluginPermissionRequest: {
            required?: string[];
            domain?: string[];
        };
        PluginSecretRequirement: {
            name: string;
            description?: string;
            required?: boolean;
            secretRef?: string;
        };
        PluginIntegrity: {
            checksum?: string;
            signature?: string;
            verified?: boolean;
            status?: string;
        };
        PluginManifest: {
            id: string;
            name: string;
            version: string;
            publisher: string;
            /** @enum {string} */
            type: "skill" | "skill-pack" | "mcp-preset" | "connector" | "ai-provider-adapter" | "agent-profile" | "gateway-policy-pack";
            description?: string;
            homepage?: string;
            compatibility?: components["schemas"]["PluginCompatibility"];
            assets?: components["schemas"]["PluginAssetSnapshot"];
            capabilities?: components["schemas"]["PluginCapabilityRequest"];
            permissions?: components["schemas"]["PluginPermissionRequest"];
            secrets?: {
                required?: components["schemas"]["PluginSecretRequirement"][];
            };
            integrity?: components["schemas"]["PluginIntegrity"];
            metadata?: {
                [key: string]: unknown;
            };
        };
        /** @enum {string} */
        CloudExtensionPointKind: "managed-agent-fleet" | "managed-connector" | "managed-cluster" | "ops-back-office" | "public-contract";
        /** @enum {string} */
        CloudExtensionBoundary: "public-contract" | "http-api" | "sdk" | "webhook" | "event-sink";
        /** @enum {string} */
        CloudOnlyDomain: "billing" | "quota" | "saas-iam";
        CloudExtensionPoint: {
            id: string;
            name: string;
            kind: components["schemas"]["CloudExtensionPointKind"];
            boundary: components["schemas"]["CloudExtensionBoundary"];
            description?: string;
            contractRefs: string[];
            supportedConsumers?: ("core" | "cli" | "web" | "agent" | "skills" | "connectors" | "cloud")[];
            requiredHeaders?: string[];
            eventTypes?: string[];
            excludedCloudOnlyDomains: components["schemas"]["CloudOnlyDomain"][];
        };
        CloudExtensionPointListEnvelope: {
            items: components["schemas"]["CloudExtensionPoint"][];
        };
        MarketplacePlugin: {
            id: string;
            name: string;
            version: string;
            publisher: string;
            type: string;
            summary?: string;
            source: string;
            riskLevel?: string;
            manifest: components["schemas"]["PluginManifest"];
            installed?: boolean;
        };
        InstalledPlugin: {
            id: string;
            name: string;
            version: string;
            publisher: string;
            type: string;
            /** @enum {string} */
            status: "enabled" | "disabled";
            source: string;
            manifest: components["schemas"]["PluginManifest"];
            checksumStatus: string;
            signatureStatus?: string;
            requestedPermissions?: components["schemas"]["PluginPermissionRequest"];
            configuredSecretRefs?: {
                [key: string]: string;
            };
            installedBy: string;
            /** Format: date-time */
            installedAt: string;
            /** Format: date-time */
            updatedAt: string;
            /** Format: date-time */
            enabledAt?: string;
            /** Format: date-time */
            disabledAt?: string;
            metadata?: {
                [key: string]: unknown;
            };
        };
        PluginInstallRequest: {
            pluginId?: string;
            source?: string;
            manifest?: components["schemas"]["PluginManifest"];
            expectedChecksum?: string;
            enable?: boolean;
        };
        PluginConfigRequest: {
            enabled?: boolean;
            secretRefs?: {
                [key: string]: string;
            };
            metadata?: {
                [key: string]: unknown;
            };
        };
        PluginManifestEnvelope: {
            data: components["schemas"]["PluginManifest"];
        };
        MarketplacePluginEnvelope: {
            data: components["schemas"]["MarketplacePlugin"];
        };
        MarketplacePluginListEnvelope: {
            items: components["schemas"]["MarketplacePlugin"][];
        };
        InstalledPluginEnvelope: {
            data: components["schemas"]["InstalledPlugin"];
        };
        InstalledPluginListEnvelope: {
            items: components["schemas"]["InstalledPlugin"][];
        };
        PersonalAccessToken: {
            id: string;
            userId: string;
            name: string;
            tokenPrefix: string;
            scopes: string[];
            permissionKeys: string[];
            metadata?: {
                [key: string]: unknown;
            };
            /** Format: date-time */
            expiresAt?: string;
            /** Format: date-time */
            lastUsedAt?: string;
            /** Format: date-time */
            revokedAt?: string;
            createdBy: string;
            /** Format: date-time */
            createdAt: string;
            /** Format: date-time */
            updatedAt: string;
        };
        PersonalAccessTokenInput: {
            name: string;
            scopes: string[];
            permissionKeys: string[];
            metadata?: {
                [key: string]: unknown;
            };
            /** Format: date-time */
            expiresAt?: string;
        };
        CreatedPersonalAccessToken: {
            token: components["schemas"]["PersonalAccessToken"];
            value: string;
        };
        PersonalAccessTokenListEnvelope: {
            items: components["schemas"]["PersonalAccessToken"][];
        };
        CreatedPersonalAccessTokenEnvelope: {
            data: components["schemas"]["CreatedPersonalAccessToken"];
        };
        ServiceAccount: {
            id: string;
            name: string;
            description?: string;
            status: string;
            ownerUserId?: string;
            roleIds: string[];
            teamIds: string[];
            scopeGrantIds: string[];
            metadata?: {
                [key: string]: unknown;
            };
            createdBy: string;
            /** Format: date-time */
            createdAt: string;
            /** Format: date-time */
            updatedAt: string;
        };
        ServiceAccountInput: {
            id: string;
            name: string;
            description?: string;
            status: string;
            ownerUserId?: string;
            roleIds?: string[];
            teamIds?: string[];
            scopeGrantIds?: string[];
            metadata?: {
                [key: string]: unknown;
            };
        };
        ServiceAccountEnvelope: {
            data: components["schemas"]["ServiceAccount"];
        };
        ServiceAccountListEnvelope: {
            items: components["schemas"]["ServiceAccount"][];
        };
        ServiceAccountToken: {
            id: string;
            serviceAccountId: string;
            name: string;
            tokenPrefix: string;
            scopes: string[];
            permissionKeys: string[];
            metadata?: {
                [key: string]: unknown;
            };
            /** Format: date-time */
            expiresAt?: string;
            /** Format: date-time */
            lastUsedAt?: string;
            /** Format: date-time */
            revokedAt?: string;
            createdBy: string;
            /** Format: date-time */
            createdAt: string;
            /** Format: date-time */
            updatedAt: string;
        };
        ServiceAccountTokenInput: {
            name: string;
            scopes: string[];
            permissionKeys: string[];
            metadata?: {
                [key: string]: unknown;
            };
            /** Format: date-time */
            expiresAt?: string;
        };
        CreatedServiceAccountToken: {
            token: components["schemas"]["ServiceAccountToken"];
            value: string;
        };
        ServiceAccountTokenListEnvelope: {
            items: components["schemas"]["ServiceAccountToken"][];
        };
        CreatedServiceAccountTokenEnvelope: {
            data: components["schemas"]["CreatedServiceAccountToken"];
        };
        TokenRotationInput: {
            /** Format: date-time */
            expiresAt?: string;
        };
        AuditLog: {
            id: string;
            actorType: string;
            actorId: string;
            actorName?: string;
            aiClientId?: string;
            aiClientName?: string;
            skillId?: string;
            toolName?: string;
            riskLevel?: components["schemas"]["RiskLevel"];
            resourceScope?: {
                [key: string]: unknown;
            };
            action: string;
            result: string;
            summary: string;
            requestId?: string;
            sourceIp?: string;
            metadata?: {
                [key: string]: unknown;
            };
            /** Format: date-time */
            createdAt: string;
        };
        AuditLogListEnvelope: {
            items: components["schemas"]["AuditLog"][];
        };
        ApprovalTrace: {
            approvalMode?: string;
            currentStageIndex?: number;
            currentStageName?: string;
            stageCount?: number;
            approvedCount?: number;
            requiredApprovals?: number;
            pendingRequirements?: string[];
            satisfiedRequirements?: string[];
            roleApprovedCounts?: {
                [key: string]: number;
            };
            teamApprovedCounts?: {
                [key: string]: number;
            };
            candidateUserIds?: string[];
            candidateRoles?: string[];
            candidateTeams?: string[];
            onCallCandidateUserIds?: string[];
            workflowRunId?: string;
            executionTaskId?: string;
            releaseBundleId?: string;
            decisions?: components["schemas"]["ApprovalDecisionTrace"][];
            stageHistory?: components["schemas"]["ApprovalStageTrace"][];
        };
        ApprovalDecisionTrace: {
            userId?: string;
            userName?: string;
            roles?: string[];
            teams?: string[];
            result?: string;
            comment?: string;
            stageIndex?: number;
            stageName?: string;
            /** Format: date-time */
            decidedAt?: string;
        };
        ApprovalStageTrace: {
            stageIndex?: number;
            stageName?: string;
            result?: string;
            /** Format: date-time */
            completedAt?: string;
        };
        ApprovalTimelineEvent: {
            id: string;
            kind: string;
            action: string;
            result: string;
            summary?: string;
            actorType?: string;
            actorId?: string;
            actorName?: string;
            stageIndex?: number;
            stageName?: string;
            metadata?: {
                [key: string]: unknown;
            };
            /** Format: date-time */
            createdAt: string;
        };
        ApprovalRequest: {
            id: string;
            status: string;
            strategy: string;
            policyId?: string;
            approvalPolicyRef?: string;
            actorType: string;
            actorId: string;
            actorName?: string;
            actorRoles?: string[];
            actorTeams?: string[];
            aiClientId?: string;
            aiClientName?: string;
            skillId?: string;
            toolName: string;
            riskLevel: components["schemas"]["RiskLevel"];
            requiresApproval: boolean;
            resourceScope?: {
                [key: string]: unknown;
            };
            toolInput?: {
                [key: string]: unknown;
            };
            relatedIds?: {
                [key: string]: unknown;
            };
            approvalTrace?: components["schemas"]["ApprovalTrace"];
            output?: components["schemas"]["AnyValue"];
            summary: string;
            requestId?: string;
            sourceIp?: string;
            decidedBy?: string;
            decidedByName?: string;
            /** Format: date-time */
            decidedAt?: string;
            decisionComment?: string;
            /** Format: date-time */
            expiresAt?: string;
            /** Format: date-time */
            createdAt: string;
            /** Format: date-time */
            updatedAt: string;
        };
        ApprovalDecisionInput: {
            comment?: string;
        };
        ApprovalDecisionResult: {
            request: components["schemas"]["ApprovalRequest"];
            invocation?: components["schemas"]["ToolInvocationResult"];
        };
        ApprovalTimeline: {
            request: components["schemas"]["ApprovalRequest"];
            trace?: components["schemas"]["ApprovalTrace"];
            events?: components["schemas"]["ApprovalTimelineEvent"][];
        };
        ApprovalRequestListEnvelope: {
            items: components["schemas"]["ApprovalRequest"][];
        };
        ApprovalTimelineEnvelope: {
            data: components["schemas"]["ApprovalTimeline"];
        };
        ApprovalDecisionResultEnvelope: {
            data: components["schemas"]["ApprovalDecisionResult"];
        };
        GovernanceStatus: {
            /** Format: date-time */
            generatedAt: string;
            windowHours: number;
            health: components["schemas"]["GovernanceHealth"];
            metrics: components["schemas"]["GovernanceMetrics"];
            tokens: components["schemas"]["GovernanceTokenSummary"];
            clients: components["schemas"]["GovernanceClientSummary"];
            approvals: components["schemas"]["GovernanceApprovalSummary"];
            policyCoverage: components["schemas"]["GovernancePolicyCoverage"];
            redaction: components["schemas"]["GovernanceRedactionSummary"];
            anomalies?: components["schemas"]["GovernanceFinding"][];
            recommendations?: string[];
            recommendationActions?: components["schemas"]["GovernanceRecommendationAction"][];
            metadata?: {
                [key: string]: unknown;
            };
        };
        GovernanceHealth: {
            status: string;
            message: string;
            checks?: components["schemas"]["GovernanceHealthCheck"][];
        };
        GovernanceHealthCheck: {
            name: string;
            status: string;
            message: string;
            count?: number;
        };
        GovernanceMetrics: {
            totalCalls: number;
            successCount: number;
            denyCount: number;
            failureCount: number;
            pendingApprovalCount: number;
            dryRunCount: number;
            riskCounts?: {
                [key: string]: number;
            };
            topTools?: components["schemas"]["GovernanceMetricCount"][];
            topAiClients?: components["schemas"]["GovernanceMetricCount"][];
            topActors?: components["schemas"]["GovernanceMetricCount"][];
            recentResultBreakdown?: {
                [key: string]: number;
            };
            recentActionBreakdown?: {
                [key: string]: number;
            };
        };
        GovernanceMetricCount: {
            key: string;
            count: number;
        };
        GovernanceRedactionSummary: {
            totalMatches: number;
            auditsWithRedaction: number;
            inputAudits: number;
            outputAudits: number;
            fieldMatches: number;
            sensitiveKeyMatches: number;
            sensitiveTextMatches: number;
            valuePatternMatches: number;
            secretClassifierMatches: number;
            structuredSecretMatches: number;
            topTargets?: components["schemas"]["GovernanceMetricCount"][];
            topFieldPaths?: components["schemas"]["GovernanceMetricCount"][];
            topMatchTypes?: components["schemas"]["GovernanceMetricCount"][];
            topClassifiers?: components["schemas"]["GovernanceMetricCount"][];
            topPolicies?: components["schemas"]["GovernanceMetricCount"][];
            topTools?: components["schemas"]["GovernanceMetricCount"][];
        };
        GovernanceTokenSummary: {
            personalAccessTokens: components["schemas"]["GovernanceTokenCounts"];
            serviceAccountTokens: components["schemas"]["GovernanceTokenCounts"];
            expiringSoon?: components["schemas"]["GovernanceTokenFinding"][];
            expiredActive?: components["schemas"]["GovernanceTokenFinding"][];
            stale?: components["schemas"]["GovernanceTokenFinding"][];
            neverUsed?: components["schemas"]["GovernanceTokenFinding"][];
            lastUsedTrackingState: string;
        };
        GovernanceTokenCounts: {
            total: number;
            active: number;
            revoked: number;
            expired: number;
            expiringSoon: number;
            stale: number;
            neverUsed: number;
        };
        GovernanceTokenFinding: {
            kind: string;
            id: string;
            name: string;
            ownerId?: string;
            tokenPrefix: string;
            severity: string;
            message: string;
            /** Format: date-time */
            expiresAt?: string;
            /** Format: date-time */
            lastUsedAt?: string;
            daysUntilDue?: number;
            staleDays?: number;
        };
        GovernanceClientSummary: {
            total: number;
            active: number;
            disabled: number;
            pendingApproval: number;
            registrationApproval: string;
            pendingApprovalClientIds?: string[];
        };
        GovernanceApprovalSummary: {
            pending: number;
            dueSoon: number;
            stalePending: number;
            overdue: number;
            oldestPendingHours?: number;
            oldestPendingRequestId?: string;
            /** Format: date-time */
            nextDueAt?: string;
            nextDueRequestId?: string;
            dueSoonRequestIds?: string[];
            stalePendingRequestIds?: string[];
            overdueRequestIds?: string[];
        };
        GovernancePolicyCoverage: {
            accessPolicies: number;
            toolGrants: number;
            skillBindings: number;
            activeAccessPolicies: number;
            activeToolGrants: number;
            activeSkillBindings: number;
            budgetPolicies: number;
            rateLimitPolicies: number;
            redactionPolicies: number;
            resourceScopedAccessPolicies: number;
            resourceScopedToolGrants: number;
            budgetState: string;
            rateLimitState: string;
            redactionPolicyState: string;
            resourceScopeState: string;
        };
        GovernanceFinding: {
            type: string;
            severity: string;
            summary: string;
            count?: number;
            actorType?: string;
            actorId?: string;
            subjectType?: string;
            subjectId?: string;
            aiClientId?: string;
            policyId?: string;
            approvalRequestId?: string;
            grantId?: string;
            toolName?: string;
            riskLevel?: components["schemas"]["RiskLevel"];
        };
        GovernanceRecommendationAction: {
            type: string;
            severity: string;
            summary: string;
            action: string;
            targetKind?: string;
            targetId?: string;
            refs?: string[];
            count?: number;
            metadata?: {
                [key: string]: unknown;
            };
        };
        GovernanceStatusEnvelope: {
            data: components["schemas"]["GovernanceStatus"];
        };
        MCPCapability: {
            adapterID: string;
            name: string;
            description: string;
            scopes: string[];
        };
        MCPCapabilityListEnvelope: {
            items: components["schemas"]["MCPCapability"][];
        };
    };
    responses: {
        /** @description Standard Soha API error envelope. */
        Error: {
            headers: {
                [name: string]: unknown;
            };
            content: {
                "application/json": components["schemas"]["ErrorEnvelope"];
            };
        };
    };
    parameters: {
        AIClientID: string;
        AIClientName: string;
        OIDCCode: string;
        OIDCState: string;
        OperationID: string;
        ProviderID: string;
        PluginID: string;
        SkillID: string;
        Source: string;
        TaskID: string;
        ToolName: string;
        UpstreamID: string;
        RouteID: string;
        /** @description First-class OpenAI-compatible relay provider kind. Azure OpenAI uses Azure data-plane URL and api-key header adaptation upstream. */
        OpenAICompatibleProvider: "deepseek" | "qwen" | "openrouter" | "azure-openai";
        /** @description Optional debug override for administrators or tokens explicitly allowed to select an upstream. */
        SohaUpstreamIDHeader: string;
        /** @description When true, returns redacted Soha route trace response headers for relay managers or tokens with allowRouteTrace explicitly enabled. */
        SohaRouteTraceHeader: boolean;
        /** @description Optional relay cache behavior hint. */
        SohaCacheModeHeader: "default" | "bypass" | "read-only" | "refresh";
    };
    requestBodies: never;
    headers: {
        /** @description Selected Soha relay route ID returned when X-Soha-Route-Trace is authorized. */
        SohaRouteIDHeader: string;
        /** @description Selected Soha relay upstream ID returned when X-Soha-Route-Trace is authorized. */
        SohaSelectedUpstreamIDHeader: string;
        /** @description Selected upstream provider kind returned when X-Soha-Route-Trace is authorized. */
        SohaProviderKindHeader: string;
        /** @description Public model requested by the client returned when X-Soha-Route-Trace is authorized. */
        SohaPublicModelHeader: string;
        /** @description Upstream model selected by Soha returned when X-Soha-Route-Trace is authorized. */
        SohaUpstreamModelHeader: string;
        /** @description Relay endpoint returned when X-Soha-Route-Trace is authorized. */
        SohaRelayEndpointHeader: string;
        /** @description Whether the relayed request streamed upstream output, returned when X-Soha-Route-Trace is authorized. */
        SohaRelayStreamHeader: boolean;
        /** @description Upstream HTTP status returned when X-Soha-Route-Trace is authorized. */
        SohaUpstreamStatusHeader: string;
        /** @description Relay cache status returned when X-Soha-Route-Trace is authorized. */
        SohaCacheStatusHeader: string;
    };
    pathItems: never;
}
export type AnyObject = Record<string, unknown>;
export type ApiResponse<T = unknown> = { data: T };
export type ApiItemsResponse<T = unknown> = { items: T[] };
export type AuthTokens = TokenSet;
export type AnyValue = components['schemas']['AnyValue'];
export type GenericObject = components['schemas']['GenericObject'];
export type GenericDataEnvelope = components['schemas']['GenericDataEnvelope'];
export type GenericItemsEnvelope = components['schemas']['GenericItemsEnvelope'];
export type ErrorEnvelope = components['schemas']['ErrorEnvelope'];
export type AIWorkbenchModelSettings = components['schemas']['AIWorkbenchModelSettings'];
export type AISkillSettings = components['schemas']['AISkillSettings'];
export type AISettings = components['schemas']['AISettings'];
export type AISettingsEnvelope = components['schemas']['AISettingsEnvelope'];
export type UpdateAIWorkbenchModelRequest = components['schemas']['UpdateAIWorkbenchModelRequest'];
export type UpdateAISkillsRequest = components['schemas']['UpdateAISkillsRequest'];
export type Principal = components['schemas']['Principal'];
export type UserProfile = components['schemas']['UserProfile'];
export type TokenSet = components['schemas']['TokenSet'];
export type AuthResult = components['schemas']['AuthResult'];
export type AuthProvider = components['schemas']['AuthProvider'];
export type LoginOptions = components['schemas']['LoginOptions'];
export type PasswordLoginRequest = components['schemas']['PasswordLoginRequest'];
export type RefreshRequest = components['schemas']['RefreshRequest'];
export type OIDCExchangeRequest = components['schemas']['OIDCExchangeRequest'];
export type StreamTicketRequest = components['schemas']['StreamTicketRequest'];
export type StreamTicket = components['schemas']['StreamTicket'];
export type PrincipalEnvelope = components['schemas']['PrincipalEnvelope'];
export type UserProfileEnvelope = components['schemas']['UserProfileEnvelope'];
export type AuthResultEnvelope = components['schemas']['AuthResultEnvelope'];
export type AuthProviderListEnvelope = components['schemas']['AuthProviderListEnvelope'];
export type LoginOptionsEnvelope = components['schemas']['LoginOptionsEnvelope'];
export type AuthBootstrapEnvelope = components['schemas']['AuthBootstrapEnvelope'];
export type StreamTicketEnvelope = components['schemas']['StreamTicketEnvelope'];
export type ExecutionTask = components['schemas']['ExecutionTask'];
export type ExecutionTaskClaimRequest = components['schemas']['ExecutionTaskClaimRequest'];
export type ExecutionCallbackRequest = components['schemas']['ExecutionCallbackRequest'];
export type DockerOperation = components['schemas']['DockerOperation'];
export type DockerOperationClaimRequest = components['schemas']['DockerOperationClaimRequest'];
export type DockerOperationCallbackRequest = components['schemas']['DockerOperationCallbackRequest'];
export type AgentRun = components['schemas']['AgentRun'];
export type AgentRunClaimRequest = components['schemas']['AgentRunClaimRequest'];
export type AgentRunCallbackRequest = components['schemas']['AgentRunCallbackRequest'];
export type AgentRunToolCallRequest = components['schemas']['AgentRunToolCallRequest'];
export type AgentToolCallResult = components['schemas']['AgentToolCallResult'];
export type ExecutionTaskEnvelope = components['schemas']['ExecutionTaskEnvelope'];
export type DockerOperationEnvelope = components['schemas']['DockerOperationEnvelope'];
export type AgentRunEnvelope = components['schemas']['AgentRunEnvelope'];
export type AgentToolCallResultEnvelope = components['schemas']['AgentToolCallResultEnvelope'];
export type RiskLevel = components['schemas']['RiskLevel'];
export type ClusterCapabilityStatus = components['schemas']['ClusterCapabilityStatus'];
export type ClusterCapabilityModeSupport = components['schemas']['ClusterCapabilityModeSupport'];
export type ClusterCapabilityMatrixEntry = components['schemas']['ClusterCapabilityMatrixEntry'];
export type ClusterCapabilityMatrixEnvelope = components['schemas']['ClusterCapabilityMatrixEnvelope'];
export type JSONSchema = components['schemas']['JSONSchema'];
export type ToolCapability = components['schemas']['ToolCapability'];
export type ResourceCapability = components['schemas']['ResourceCapability'];
export type PromptCapability = components['schemas']['PromptCapability'];
export type SkillCapability = components['schemas']['SkillCapability'];
export type CallerContext = components['schemas']['CallerContext'];
export type AIGatewayManifest = components['schemas']['AIGatewayManifest'];
export type ToolInvocationRequest = components['schemas']['ToolInvocationRequest'];
export type ToolInvocationResult = components['schemas']['ToolInvocationResult'];
export type ResourceReadRequest = components['schemas']['ResourceReadRequest'];
export type ResourceReadResult = components['schemas']['ResourceReadResult'];
export type PromptGetRequest = components['schemas']['PromptGetRequest'];
export type PromptMessage = components['schemas']['PromptMessage'];
export type PromptGetResult = components['schemas']['PromptGetResult'];
export type AIGatewayManifestEnvelope = components['schemas']['AIGatewayManifestEnvelope'];
export type ToolInvocationResultEnvelope = components['schemas']['ToolInvocationResultEnvelope'];
export type ResourceReadResultEnvelope = components['schemas']['ResourceReadResultEnvelope'];
export type PromptGetResultEnvelope = components['schemas']['PromptGetResultEnvelope'];
export type OperationStatus = components['schemas']['OperationStatus'];
export type LLMTokenMetadata = components['schemas']['LLMTokenMetadata'];
export type LLMUpstream = components['schemas']['LLMUpstream'];
export type LLMUpstreamInput = components['schemas']['LLMUpstreamInput'];
export type LLMUpstreamEnvelope = components['schemas']['LLMUpstreamEnvelope'];
export type LLMUpstreamListEnvelope = components['schemas']['LLMUpstreamListEnvelope'];
export type LLMUpstreamTestRequest = components['schemas']['LLMUpstreamTestRequest'];
export type LLMUpstreamTestResult = components['schemas']['LLMUpstreamTestResult'];
export type LLMUpstreamTestResultEnvelope = components['schemas']['LLMUpstreamTestResultEnvelope'];
export type LLMRelayHealthCheckRun = components['schemas']['LLMRelayHealthCheckRun'];
export type LLMRelayHealthCheckRunEnvelope = components['schemas']['LLMRelayHealthCheckRunEnvelope'];
export type LLMModelRoute = components['schemas']['LLMModelRoute'];
export type LLMModelRouteInput = components['schemas']['LLMModelRouteInput'];
export type LLMModelRouteEnvelope = components['schemas']['LLMModelRouteEnvelope'];
export type LLMModelRouteListEnvelope = components['schemas']['LLMModelRouteListEnvelope'];
export type LLMCallLog = components['schemas']['LLMCallLog'];
export type LLMCallLogListEnvelope = components['schemas']['LLMCallLogListEnvelope'];
export type LLMRelayMetricSeriesPoint = components['schemas']['LLMRelayMetricSeriesPoint'];
export type LLMRelayMetricBreakdown = components['schemas']['LLMRelayMetricBreakdown'];
export type LLMRelayMetrics = components['schemas']['LLMRelayMetrics'];
export type LLMRelayMetricsEnvelope = components['schemas']['LLMRelayMetricsEnvelope'];
export type LLMRelayCacheStats = components['schemas']['LLMRelayCacheStats'];
export type LLMRelayCacheStatsEnvelope = components['schemas']['LLMRelayCacheStatsEnvelope'];
export type LLMRelayCachePurgeRequest = components['schemas']['LLMRelayCachePurgeRequest'];
export type LLMRelayCachePurgeResult = components['schemas']['LLMRelayCachePurgeResult'];
export type LLMRelayCachePurgeResultEnvelope = components['schemas']['LLMRelayCachePurgeResultEnvelope'];
export type NativeProviderObject = components['schemas']['NativeProviderObject'];
export type OpenAIModel = components['schemas']['OpenAIModel'];
export type OpenAIModelsResponse = components['schemas']['OpenAIModelsResponse'];
export type OpenAIChatCompletionRequest = components['schemas']['OpenAIChatCompletionRequest'];
export type OpenAIResponseRequest = components['schemas']['OpenAIResponseRequest'];
export type OpenAIEmbeddingRequest = components['schemas']['OpenAIEmbeddingRequest'];
export type OpenAIImageGenerationRequest = components['schemas']['OpenAIImageGenerationRequest'];
export type OpenAIImageEditRequest = components['schemas']['OpenAIImageEditRequest'];
export type OpenAIImageVariationRequest = components['schemas']['OpenAIImageVariationRequest'];
export type OpenAIAudioSpeechRequest = components['schemas']['OpenAIAudioSpeechRequest'];
export type OpenAIAudioTranscriptionRequest = components['schemas']['OpenAIAudioTranscriptionRequest'];
export type OpenAIAudioTranslationRequest = components['schemas']['OpenAIAudioTranslationRequest'];
export type GeminiModel = components['schemas']['GeminiModel'];
export type GeminiModelsResponse = components['schemas']['GeminiModelsResponse'];
export type GeminiGenerateContentRequest = components['schemas']['GeminiGenerateContentRequest'];
export type GeminiInteractionsRequest = components['schemas']['GeminiInteractionsRequest'];
export type CohereRerankRequest = components['schemas']['CohereRerankRequest'];
export type AnthropicModel = components['schemas']['AnthropicModel'];
export type AnthropicModelsResponse = components['schemas']['AnthropicModelsResponse'];
export type AnthropicMessagesRequest = components['schemas']['AnthropicMessagesRequest'];
export type PluginAssetSnapshot = components['schemas']['PluginAssetSnapshot'];
export type PluginCompatibility = components['schemas']['PluginCompatibility'];
export type PluginCapabilityRequest = components['schemas']['PluginCapabilityRequest'];
export type PluginPermissionRequest = components['schemas']['PluginPermissionRequest'];
export type PluginSecretRequirement = components['schemas']['PluginSecretRequirement'];
export type PluginIntegrity = components['schemas']['PluginIntegrity'];
export type PluginManifest = components['schemas']['PluginManifest'];
export type CloudExtensionPointKind = components['schemas']['CloudExtensionPointKind'];
export type CloudExtensionBoundary = components['schemas']['CloudExtensionBoundary'];
export type CloudOnlyDomain = components['schemas']['CloudOnlyDomain'];
export type CloudExtensionPoint = components['schemas']['CloudExtensionPoint'];
export type CloudExtensionPointListEnvelope = components['schemas']['CloudExtensionPointListEnvelope'];
export type MarketplacePlugin = components['schemas']['MarketplacePlugin'];
export type InstalledPlugin = components['schemas']['InstalledPlugin'];
export type PluginInstallRequest = components['schemas']['PluginInstallRequest'];
export type PluginConfigRequest = components['schemas']['PluginConfigRequest'];
export type PluginManifestEnvelope = components['schemas']['PluginManifestEnvelope'];
export type MarketplacePluginEnvelope = components['schemas']['MarketplacePluginEnvelope'];
export type MarketplacePluginListEnvelope = components['schemas']['MarketplacePluginListEnvelope'];
export type InstalledPluginEnvelope = components['schemas']['InstalledPluginEnvelope'];
export type InstalledPluginListEnvelope = components['schemas']['InstalledPluginListEnvelope'];
export type PersonalAccessToken = components['schemas']['PersonalAccessToken'];
export type PersonalAccessTokenInput = components['schemas']['PersonalAccessTokenInput'];
export type CreatedPersonalAccessToken = components['schemas']['CreatedPersonalAccessToken'];
export type PersonalAccessTokenListEnvelope = components['schemas']['PersonalAccessTokenListEnvelope'];
export type CreatedPersonalAccessTokenEnvelope = components['schemas']['CreatedPersonalAccessTokenEnvelope'];
export type ServiceAccount = components['schemas']['ServiceAccount'];
export type ServiceAccountInput = components['schemas']['ServiceAccountInput'];
export type ServiceAccountEnvelope = components['schemas']['ServiceAccountEnvelope'];
export type ServiceAccountListEnvelope = components['schemas']['ServiceAccountListEnvelope'];
export type ServiceAccountToken = components['schemas']['ServiceAccountToken'];
export type ServiceAccountTokenInput = components['schemas']['ServiceAccountTokenInput'];
export type CreatedServiceAccountToken = components['schemas']['CreatedServiceAccountToken'];
export type ServiceAccountTokenListEnvelope = components['schemas']['ServiceAccountTokenListEnvelope'];
export type CreatedServiceAccountTokenEnvelope = components['schemas']['CreatedServiceAccountTokenEnvelope'];
export type TokenRotationInput = components['schemas']['TokenRotationInput'];
export type AuditLog = components['schemas']['AuditLog'];
export type AuditLogListEnvelope = components['schemas']['AuditLogListEnvelope'];
export type ApprovalTrace = components['schemas']['ApprovalTrace'];
export type ApprovalDecisionTrace = components['schemas']['ApprovalDecisionTrace'];
export type ApprovalStageTrace = components['schemas']['ApprovalStageTrace'];
export type ApprovalTimelineEvent = components['schemas']['ApprovalTimelineEvent'];
export type ApprovalRequest = components['schemas']['ApprovalRequest'];
export type ApprovalDecisionInput = components['schemas']['ApprovalDecisionInput'];
export type ApprovalDecisionResult = components['schemas']['ApprovalDecisionResult'];
export type ApprovalTimeline = components['schemas']['ApprovalTimeline'];
export type ApprovalRequestListEnvelope = components['schemas']['ApprovalRequestListEnvelope'];
export type ApprovalTimelineEnvelope = components['schemas']['ApprovalTimelineEnvelope'];
export type ApprovalDecisionResultEnvelope = components['schemas']['ApprovalDecisionResultEnvelope'];
export type GovernanceStatus = components['schemas']['GovernanceStatus'];
export type GovernanceHealth = components['schemas']['GovernanceHealth'];
export type GovernanceHealthCheck = components['schemas']['GovernanceHealthCheck'];
export type GovernanceMetrics = components['schemas']['GovernanceMetrics'];
export type GovernanceMetricCount = components['schemas']['GovernanceMetricCount'];
export type GovernanceRedactionSummary = components['schemas']['GovernanceRedactionSummary'];
export type GovernanceTokenSummary = components['schemas']['GovernanceTokenSummary'];
export type GovernanceTokenCounts = components['schemas']['GovernanceTokenCounts'];
export type GovernanceTokenFinding = components['schemas']['GovernanceTokenFinding'];
export type GovernanceClientSummary = components['schemas']['GovernanceClientSummary'];
export type GovernanceApprovalSummary = components['schemas']['GovernanceApprovalSummary'];
export type GovernancePolicyCoverage = components['schemas']['GovernancePolicyCoverage'];
export type GovernanceFinding = components['schemas']['GovernanceFinding'];
export type GovernanceRecommendationAction = components['schemas']['GovernanceRecommendationAction'];
export type GovernanceStatusEnvelope = components['schemas']['GovernanceStatusEnvelope'];
export type MCPCapability = components['schemas']['MCPCapability'];
export type MCPCapabilityListEnvelope = components['schemas']['MCPCapabilityListEnvelope'];
export type ResponseError = components['responses']['Error'];
export type ParameterAiClientId = components['parameters']['AIClientID'];
export type ParameterAiClientName = components['parameters']['AIClientName'];
export type ParameterOidcCode = components['parameters']['OIDCCode'];
export type ParameterOidcState = components['parameters']['OIDCState'];
export type ParameterOperationId = components['parameters']['OperationID'];
export type ParameterProviderId = components['parameters']['ProviderID'];
export type ParameterPluginId = components['parameters']['PluginID'];
export type ParameterSkillId = components['parameters']['SkillID'];
export type ParameterSource = components['parameters']['Source'];
export type ParameterTaskId = components['parameters']['TaskID'];
export type ParameterToolName = components['parameters']['ToolName'];
export type ParameterUpstreamId = components['parameters']['UpstreamID'];
export type ParameterRouteId = components['parameters']['RouteID'];
export type ParameterOpenAiCompatibleProvider = components['parameters']['OpenAICompatibleProvider'];
export type ParameterSohaUpstreamIdHeader = components['parameters']['SohaUpstreamIDHeader'];
export type ParameterSohaRouteTraceHeader = components['parameters']['SohaRouteTraceHeader'];
export type ParameterSohaCacheModeHeader = components['parameters']['SohaCacheModeHeader'];
export type HeaderSohaRouteIdHeader = components['headers']['SohaRouteIDHeader'];
export type HeaderSohaSelectedUpstreamIdHeader = components['headers']['SohaSelectedUpstreamIDHeader'];
export type HeaderSohaProviderKindHeader = components['headers']['SohaProviderKindHeader'];
export type HeaderSohaPublicModelHeader = components['headers']['SohaPublicModelHeader'];
export type HeaderSohaUpstreamModelHeader = components['headers']['SohaUpstreamModelHeader'];
export type HeaderSohaRelayEndpointHeader = components['headers']['SohaRelayEndpointHeader'];
export type HeaderSohaRelayStreamHeader = components['headers']['SohaRelayStreamHeader'];
export type HeaderSohaUpstreamStatusHeader = components['headers']['SohaUpstreamStatusHeader'];
export type HeaderSohaCacheStatusHeader = components['headers']['SohaCacheStatusHeader'];
export type $defs = Record<string, never>;
export interface operations {
    getHealthz: {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        requestBody?: never;
        responses: {
            /** @description Process is alive. */
            200: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["GenericObject"];
                };
            };
        };
    };
    getReadyz: {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        requestBody?: never;
        responses: {
            /** @description Process is ready. */
            200: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["GenericObject"];
                };
            };
        };
    };
    listAuthProviders: {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        requestBody?: never;
        responses: {
            /** @description Enabled login providers. */
            200: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["AuthProviderListEnvelope"];
                };
            };
        };
    };
    getAuthLoginOptions: {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        requestBody?: never;
        responses: {
            /** @description Login verification options. */
            200: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["LoginOptionsEnvelope"];
                };
            };
        };
    };
    passwordLogin: {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        requestBody: {
            content: {
                "application/json": components["schemas"]["PasswordLoginRequest"];
            };
        };
        responses: {
            /** @description Authenticated session. */
            200: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["AuthResultEnvelope"];
                };
            };
            400: components["responses"]["Error"];
            401: components["responses"]["Error"];
        };
    };
    refreshAuthSession: {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        requestBody?: {
            content: {
                "application/json": components["schemas"]["RefreshRequest"];
            };
        };
        responses: {
            /** @description Refreshed token set. */
            200: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["AuthResultEnvelope"];
                };
            };
            401: components["responses"]["Error"];
        };
    };
    beginOIDCLogin: {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        requestBody?: never;
        responses: {
            /** @description Redirect to the configured OIDC provider. */
            302: {
                headers: {
                    [name: string]: unknown;
                };
                content?: never;
            };
        };
    };
    handleOIDCCallback: {
        parameters: {
            query?: {
                code?: components["parameters"]["OIDCCode"];
                state?: components["parameters"]["OIDCState"];
            };
            header?: never;
            path?: never;
            cookie?: never;
        };
        requestBody?: never;
        responses: {
            /** @description Redirect back to the web console with an exchange code. */
            302: {
                headers: {
                    [name: string]: unknown;
                };
                content?: never;
            };
        };
    };
    exchangeOIDCCode: {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        requestBody: {
            content: {
                "application/json": components["schemas"]["OIDCExchangeRequest"];
            };
        };
        responses: {
            /** @description Authenticated session. */
            200: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["AuthResultEnvelope"];
                };
            };
        };
    };
    beginProviderLogin: {
        parameters: {
            query?: never;
            header?: never;
            path: {
                providerID: components["parameters"]["ProviderID"];
            };
            cookie?: never;
        };
        requestBody?: never;
        responses: {
            /** @description Redirect to the selected identity provider. */
            302: {
                headers: {
                    [name: string]: unknown;
                };
                content?: never;
            };
        };
    };
    handleProviderCallback: {
        parameters: {
            query?: {
                code?: components["parameters"]["OIDCCode"];
                state?: components["parameters"]["OIDCState"];
            };
            header?: never;
            path: {
                providerID: components["parameters"]["ProviderID"];
            };
            cookie?: never;
        };
        requestBody?: never;
        responses: {
            /** @description Redirect back to the web console with an exchange code. */
            302: {
                headers: {
                    [name: string]: unknown;
                };
                content?: never;
            };
        };
    };
    getCurrentPrincipal: {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        requestBody?: never;
        responses: {
            /** @description Current principal. */
            200: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["PrincipalEnvelope"];
                };
            };
            401: components["responses"]["Error"];
        };
    };
    getCurrentUserProfile: {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        requestBody?: never;
        responses: {
            /** @description Current user profile. */
            200: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["UserProfileEnvelope"];
                };
            };
        };
    };
    getAuthBootstrap: {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        requestBody?: never;
        responses: {
            /** @description Principal, permission snapshot, and branding bootstrap data. */
            200: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["AuthBootstrapEnvelope"];
                };
            };
        };
    };
    logoutAuthSession: {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        requestBody?: {
            content: {
                "application/json": components["schemas"]["RefreshRequest"];
            };
        };
        responses: {
            /** @description Session logged out. */
            204: {
                headers: {
                    [name: string]: unknown;
                };
                content?: never;
            };
            401: components["responses"]["Error"];
        };
    };
    issueStreamTicket: {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        requestBody: {
            content: {
                "application/json": components["schemas"]["StreamTicketRequest"];
            };
        };
        responses: {
            /** @description Short-lived stream ticket. */
            200: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["StreamTicketEnvelope"];
                };
            };
        };
    };
    getAISettings: {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        requestBody?: never;
        responses: {
            /** @description AI Workbench settings. Provider credentials and upstream connection details are intentionally excluded; provider ingress is managed by AI Gateway relay upstreams and model routes. */
            200: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["AISettingsEnvelope"];
                };
            };
            403: components["responses"]["Error"];
        };
    };
    updateAIWorkbenchModelSettings: {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        requestBody: {
            content: {
                "application/json": components["schemas"]["UpdateAIWorkbenchModelRequest"];
            };
        };
        responses: {
            /** @description Updated AI Workbench default model and route selection. */
            200: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["AISettingsEnvelope"];
                };
            };
            400: components["responses"]["Error"];
            403: components["responses"]["Error"];
        };
    };
    updateAISkillsRegistry: {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        requestBody: {
            content: {
                "application/json": components["schemas"]["UpdateAISkillsRequest"];
            };
        };
        responses: {
            /** @description Updated AI Workbench skills registry. */
            200: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["AISettingsEnvelope"];
                };
            };
            400: components["responses"]["Error"];
            403: components["responses"]["Error"];
        };
    };
    claimExecutionTask: {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        requestBody: {
            content: {
                "application/json": components["schemas"]["ExecutionTaskClaimRequest"];
            };
        };
        responses: {
            /** @description Claimed execution task, or empty data when none is available. */
            200: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["ExecutionTaskEnvelope"];
                };
            };
        };
    };
    getExecutionTaskRunnerStatus: {
        parameters: {
            query?: never;
            header?: never;
            path: {
                taskID: components["parameters"]["TaskID"];
            };
            cookie?: never;
        };
        requestBody?: never;
        responses: {
            /** @description Runner-visible task status. */
            200: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["ExecutionTaskEnvelope"];
                };
            };
        };
    };
    recordExecutionCallback: {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        requestBody: {
            content: {
                "application/json": components["schemas"]["ExecutionCallbackRequest"];
            };
        };
        responses: {
            /** @description Updated execution task. */
            200: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["ExecutionTaskEnvelope"];
                };
            };
        };
    };
    claimDockerOperation: {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        requestBody: {
            content: {
                "application/json": components["schemas"]["DockerOperationClaimRequest"];
            };
        };
        responses: {
            /** @description Claimed Docker operation, or empty data when none is available. */
            200: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["DockerOperationEnvelope"];
                };
            };
        };
    };
    getDockerOperationRunnerStatus: {
        parameters: {
            query?: never;
            header?: never;
            path: {
                operationID: components["parameters"]["OperationID"];
            };
            cookie?: never;
        };
        requestBody?: never;
        responses: {
            /** @description Runner-visible Docker operation status. */
            200: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["DockerOperationEnvelope"];
                };
            };
        };
    };
    recordDockerOperationCallback: {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        requestBody: {
            content: {
                "application/json": components["schemas"]["DockerOperationCallbackRequest"];
            };
        };
        responses: {
            /** @description Updated Docker operation. */
            200: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["DockerOperationEnvelope"];
                };
            };
        };
    };
    claimAgentRun: {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        requestBody: {
            content: {
                "application/json": components["schemas"]["AgentRunClaimRequest"];
            };
        };
        responses: {
            /** @description Claimed agent run, or empty data when none is available. */
            200: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["AgentRunEnvelope"];
                };
            };
        };
    };
    recordAgentRunCallback: {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        requestBody: {
            content: {
                "application/json": components["schemas"]["AgentRunCallbackRequest"];
            };
        };
        responses: {
            /** @description Updated agent run. */
            200: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["AgentRunEnvelope"];
                };
            };
        };
    };
    recordAgentRunToolCall: {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        requestBody: {
            content: {
                "application/json": components["schemas"]["AgentRunToolCallRequest"];
            };
        };
        responses: {
            /** @description Tool call result recorded by the control plane. */
            200: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["AgentToolCallResultEnvelope"];
                };
            };
        };
    };
    listAgentRuns: {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        requestBody?: never;
        responses: {
            /** @description Agent run list visible to the current user. */
            200: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["GenericItemsEnvelope"];
                };
            };
        };
    };
    listClusterCapabilityMatrix: {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        requestBody?: never;
        responses: {
            /** @description Cluster runtime capability matrix for Direct and Agent modes. */
            200: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["ClusterCapabilityMatrixEnvelope"];
                };
            };
            403: components["responses"]["Error"];
        };
    };
    getAIGatewayCapabilities: {
        parameters: {
            query?: {
                aiClientId?: components["parameters"]["AIClientID"];
                aiClientName?: components["parameters"]["AIClientName"];
                skillId?: components["parameters"]["SkillID"];
                source?: components["parameters"]["Source"];
            };
            header?: never;
            path?: never;
            cookie?: never;
        };
        requestBody?: never;
        responses: {
            /** @description Permission-filtered AI Gateway manifest. */
            200: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["AIGatewayManifestEnvelope"];
                };
            };
            403: components["responses"]["Error"];
        };
    };
    invokeAIGatewayTool: {
        parameters: {
            query?: never;
            header?: never;
            path: {
                toolName: components["parameters"]["ToolName"];
            };
            cookie?: never;
        };
        requestBody: {
            content: {
                "application/json": components["schemas"]["ToolInvocationRequest"];
            };
        };
        responses: {
            /** @description Gateway tool invocation result. */
            200: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["ToolInvocationResultEnvelope"];
                };
            };
            403: components["responses"]["Error"];
        };
    };
    readAIGatewayResource: {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        requestBody: {
            content: {
                "application/json": components["schemas"]["ResourceReadRequest"];
            };
        };
        responses: {
            /** @description Redacted Gateway resource document. */
            200: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["ResourceReadResultEnvelope"];
                };
            };
        };
    };
    getAIGatewayPrompt: {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        requestBody: {
            content: {
                "application/json": components["schemas"]["PromptGetRequest"];
            };
        };
        responses: {
            /** @description Gateway prompt messages assembled by the backend. */
            200: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["PromptGetResultEnvelope"];
                };
            };
        };
    };
    listPersonalAccessTokens: {
        parameters: {
            query?: {
                scope?: string;
                userId?: string;
            };
            header?: never;
            path?: never;
            cookie?: never;
        };
        requestBody?: never;
        responses: {
            /** @description Personal access tokens visible to the current principal. */
            200: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["PersonalAccessTokenListEnvelope"];
                };
            };
        };
    };
    createPersonalAccessToken: {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        requestBody: {
            content: {
                "application/json": components["schemas"]["PersonalAccessTokenInput"];
            };
        };
        responses: {
            /** @description Newly created personal access token and one-time secret value. */
            201: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["CreatedPersonalAccessTokenEnvelope"];
                };
            };
        };
    };
    revokePersonalAccessToken: {
        parameters: {
            query?: never;
            header?: never;
            path: {
                tokenID: string;
            };
            cookie?: never;
        };
        requestBody?: never;
        responses: {
            /** @description Personal access token was revoked. */
            200: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["OperationStatus"];
                };
            };
        };
    };
    rotatePersonalAccessToken: {
        parameters: {
            query?: never;
            header?: never;
            path: {
                tokenID: string;
            };
            cookie?: never;
        };
        requestBody?: {
            content: {
                "application/json": components["schemas"]["TokenRotationInput"];
            };
        };
        responses: {
            /** @description Rotated personal access token and one-time secret value. */
            201: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["CreatedPersonalAccessTokenEnvelope"];
                };
            };
        };
    };
    listServiceAccounts: {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        requestBody?: never;
        responses: {
            /** @description Service accounts visible to the current principal. */
            200: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["ServiceAccountListEnvelope"];
                };
            };
        };
    };
    createServiceAccount: {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        requestBody: {
            content: {
                "application/json": components["schemas"]["ServiceAccountInput"];
            };
        };
        responses: {
            /** @description Newly created service account. */
            201: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["ServiceAccountEnvelope"];
                };
            };
        };
    };
    listServiceAccountTokens: {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        requestBody?: never;
        responses: {
            /** @description Service account tokens visible to the current principal. */
            200: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["ServiceAccountTokenListEnvelope"];
                };
            };
        };
    };
    createServiceAccountToken: {
        parameters: {
            query?: never;
            header?: never;
            path: {
                serviceAccountID: string;
            };
            cookie?: never;
        };
        requestBody: {
            content: {
                "application/json": components["schemas"]["ServiceAccountTokenInput"];
            };
        };
        responses: {
            /** @description Newly created service account token and one-time secret value. */
            201: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["CreatedServiceAccountTokenEnvelope"];
                };
            };
        };
    };
    revokeServiceAccountToken: {
        parameters: {
            query?: never;
            header?: never;
            path: {
                tokenID: string;
            };
            cookie?: never;
        };
        requestBody?: never;
        responses: {
            /** @description Service account token was revoked. */
            200: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["OperationStatus"];
                };
            };
        };
    };
    rotateServiceAccountToken: {
        parameters: {
            query?: never;
            header?: never;
            path: {
                tokenID: string;
            };
            cookie?: never;
        };
        requestBody?: {
            content: {
                "application/json": components["schemas"]["TokenRotationInput"];
            };
        };
        responses: {
            /** @description Rotated service account token and one-time secret value. */
            201: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["CreatedServiceAccountTokenEnvelope"];
                };
            };
        };
    };
    listAIGatewayAuditLogs: {
        parameters: {
            query?: {
                actorType?: string;
                /** @description Canonical actor identifier query parameter. The server also accepts actor. */
                actorId?: string;
                aiClientId?: string;
                skillId?: string;
                toolName?: string;
                approvalRequestId?: string;
                riskLevel?: components["schemas"]["RiskLevel"];
                result?: string;
                action?: string;
                from?: string;
                to?: string;
                limit?: number;
            };
            header?: never;
            path?: never;
            cookie?: never;
        };
        requestBody?: never;
        responses: {
            /** @description AI Gateway audit log entries. */
            200: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["AuditLogListEnvelope"];
                };
            };
        };
    };
    listAIGatewayApprovalRequests: {
        parameters: {
            query?: {
                /** @description Canonical approval request id query parameter. The server also accepts requestID and approvalRequestId. */
                id?: string;
                status?: string;
                actorType?: string;
                /** @description Canonical actor identifier query parameter. The server also accepts actor. */
                actorId?: string;
                aiClientId?: string;
                skillId?: string;
                toolName?: string;
                riskLevel?: components["schemas"]["RiskLevel"];
                strategy?: string;
                from?: string;
                to?: string;
                limit?: number;
            };
            header?: never;
            path?: never;
            cookie?: never;
        };
        requestBody?: never;
        responses: {
            /** @description AI Gateway approval requests. */
            200: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["ApprovalRequestListEnvelope"];
                };
            };
        };
    };
    getAIGatewayApprovalTimeline: {
        parameters: {
            query?: never;
            header?: never;
            path: {
                requestID: string;
            };
            cookie?: never;
        };
        requestBody?: never;
        responses: {
            /** @description Approval request timeline. */
            200: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["ApprovalTimelineEnvelope"];
                };
            };
        };
    };
    decideAIGatewayApprovalRequest: {
        parameters: {
            query?: never;
            header?: never;
            path: {
                requestID: string;
                action: "approve" | "reject" | "cancel";
            };
            cookie?: never;
        };
        requestBody?: {
            content: {
                "application/json": components["schemas"]["ApprovalDecisionInput"];
            };
        };
        responses: {
            /** @description Approval decision result. */
            200: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["ApprovalDecisionResultEnvelope"];
                };
            };
        };
    };
    getAIGatewayGovernanceStatus: {
        parameters: {
            query?: {
                windowHours?: number;
            };
            header?: never;
            path?: never;
            cookie?: never;
        };
        requestBody?: never;
        responses: {
            /** @description AI Gateway governance health and usage summary. */
            200: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["GovernanceStatusEnvelope"];
                };
            };
        };
    };
    listAIGatewayRelayUpstreams: {
        parameters: {
            query?: {
                providerKind?: "openai" | "anthropic" | "openai-compatible" | "deepseek" | "qwen" | "openrouter" | "azure-openai" | "gemini" | "cohere";
                status?: "active" | "disabled" | "degraded";
                includeAll?: boolean;
            };
            header?: never;
            path?: never;
            cookie?: never;
        };
        requestBody?: never;
        responses: {
            /** @description Relay upstreams visible to the current Gateway manager. Full upstream API keys are never returned. */
            200: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["LLMUpstreamListEnvelope"];
                };
            };
            403: components["responses"]["Error"];
        };
    };
    createAIGatewayRelayUpstream: {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        requestBody: {
            content: {
                "application/json": components["schemas"]["LLMUpstreamInput"];
            };
        };
        responses: {
            /** @description Created relay upstream. The stored API key is encrypted and only the display prefix is returned. */
            201: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["LLMUpstreamEnvelope"];
                };
            };
            400: components["responses"]["Error"];
            403: components["responses"]["Error"];
        };
    };
    updateAIGatewayRelayUpstream: {
        parameters: {
            query?: never;
            header?: never;
            path: {
                upstreamID: components["parameters"]["UpstreamID"];
            };
            cookie?: never;
        };
        requestBody: {
            content: {
                "application/json": components["schemas"]["LLMUpstreamInput"];
            };
        };
        responses: {
            /** @description Updated relay upstream. Omit apiKey to keep the existing encrypted key. */
            200: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["LLMUpstreamEnvelope"];
                };
            };
            403: components["responses"]["Error"];
            404: components["responses"]["Error"];
        };
    };
    runAIGatewayRelayHealthChecks: {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        requestBody?: never;
        responses: {
            /** @description Runs relay upstream health checks and applies health-check managed degraded/recovered status transitions. */
            200: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["LLMRelayHealthCheckRunEnvelope"];
                };
            };
            403: components["responses"]["Error"];
        };
    };
    testAIGatewayRelayUpstream: {
        parameters: {
            query?: never;
            header?: never;
            path: {
                upstreamID: components["parameters"]["UpstreamID"];
            };
            cookie?: never;
        };
        requestBody?: {
            content: {
                "application/json": components["schemas"]["LLMUpstreamTestRequest"];
            };
        };
        responses: {
            /** @description Redacted upstream connectivity test result. */
            200: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["LLMUpstreamTestResultEnvelope"];
                };
            };
            403: components["responses"]["Error"];
        };
    };
    listAIGatewayRelayModelRoutes: {
        parameters: {
            query?: {
                publicModel?: string;
                providerKind?: "openai" | "anthropic" | "openai-compatible" | "deepseek" | "qwen" | "openrouter" | "azure-openai" | "gemini" | "cohere";
                upstreamId?: string;
                routeGroup?: string;
                includeDisabled?: boolean;
            };
            header?: never;
            path?: never;
            cookie?: never;
        };
        requestBody?: never;
        responses: {
            /** @description Relay model routes. */
            200: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["LLMModelRouteListEnvelope"];
                };
            };
            403: components["responses"]["Error"];
        };
    };
    createAIGatewayRelayModelRoute: {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        requestBody: {
            content: {
                "application/json": components["schemas"]["LLMModelRouteInput"];
            };
        };
        responses: {
            /** @description Created relay model route. */
            201: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["LLMModelRouteEnvelope"];
                };
            };
            403: components["responses"]["Error"];
        };
    };
    updateAIGatewayRelayModelRoute: {
        parameters: {
            query?: never;
            header?: never;
            path: {
                routeID: components["parameters"]["RouteID"];
            };
            cookie?: never;
        };
        requestBody: {
            content: {
                "application/json": components["schemas"]["LLMModelRouteInput"];
            };
        };
        responses: {
            /** @description Updated relay model route. */
            200: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["LLMModelRouteEnvelope"];
                };
            };
            403: components["responses"]["Error"];
            404: components["responses"]["Error"];
        };
    };
    deleteAIGatewayRelayModelRoute: {
        parameters: {
            query?: never;
            header?: never;
            path: {
                routeID: components["parameters"]["RouteID"];
            };
            cookie?: never;
        };
        requestBody?: never;
        responses: {
            /** @description Deleted relay model route. */
            200: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["OperationStatus"];
                };
            };
            403: components["responses"]["Error"];
            404: components["responses"]["Error"];
        };
    };
    listAIGatewayRelayModelCalls: {
        parameters: {
            query?: {
                actorType?: string;
                actorId?: string;
                tokenId?: string;
                publicModel?: string;
                upstreamId?: string;
                providerKind?: "openai" | "anthropic" | "openai-compatible" | "deepseek" | "qwen" | "openrouter" | "azure-openai" | "gemini" | "cohere";
                status?: "success" | "failure" | "cancelled" | "client_cancelled" | "rate_limited" | "policy_denied";
                endpoint?: string;
                from?: string;
                to?: string;
                limit?: number;
            };
            header?: never;
            path?: never;
            cookie?: never;
        };
        requestBody?: never;
        responses: {
            /** @description Redacted relay model call logs. Prompts, full headers, token values, and upstream API keys are not returned. */
            200: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["LLMCallLogListEnvelope"];
                };
            };
            403: components["responses"]["Error"];
        };
    };
    getAIGatewayRelayMetrics: {
        parameters: {
            query?: {
                windowHours?: number;
                publicModel?: string;
                upstreamId?: string;
            };
            header?: never;
            path?: never;
            cookie?: never;
        };
        requestBody?: never;
        responses: {
            /** @description Relay usage, latency, token, and cache-token metrics. */
            200: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["LLMRelayMetricsEnvelope"];
                };
            };
            403: components["responses"]["Error"];
        };
    };
    getAIGatewayRelayCacheStats: {
        parameters: {
            query?: {
                windowHours?: number;
                publicModel?: string;
                upstreamId?: string;
            };
            header?: never;
            path?: never;
            cookie?: never;
        };
        requestBody?: never;
        responses: {
            /** @description Relay response cache and provider prompt-cache statistics. */
            200: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["LLMRelayCacheStatsEnvelope"];
                };
            };
            403: components["responses"]["Error"];
        };
    };
    purgeAIGatewayRelayCache: {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        requestBody?: {
            content: {
                "application/json": components["schemas"]["LLMRelayCachePurgeRequest"];
            };
        };
        responses: {
            /** @description Relay cache purge accepted. */
            200: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["LLMRelayCachePurgeResultEnvelope"];
                };
            };
            403: components["responses"]["Error"];
        };
    };
    listAIGatewayOpenAIModels: {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        requestBody?: never;
        responses: {
            /** @description Native OpenAI-compatible models response. This endpoint does not use the OpenSoha envelope. */
            200: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["OpenAIModelsResponse"];
                };
            };
            401: components["responses"]["Error"];
            403: components["responses"]["Error"];
        };
    };
    createAIGatewayOpenAIChatCompletion: {
        parameters: {
            query?: never;
            header?: {
                /** @description Optional debug override for administrators or tokens explicitly allowed to select an upstream. */
                "X-Soha-Upstream-ID"?: components["parameters"]["SohaUpstreamIDHeader"];
                /** @description When true, returns redacted Soha route trace response headers for relay managers or tokens with allowRouteTrace explicitly enabled. */
                "X-Soha-Route-Trace"?: components["parameters"]["SohaRouteTraceHeader"];
                /** @description Optional relay cache behavior hint. */
                "X-Soha-Cache-Mode"?: components["parameters"]["SohaCacheModeHeader"];
            };
            path?: never;
            cookie?: never;
        };
        requestBody: {
            content: {
                "application/json": components["schemas"]["OpenAIChatCompletionRequest"];
            };
        };
        responses: {
            /** @description Native OpenAI-compatible chat completion JSON or SSE stream. */
            200: {
                headers: {
                    "X-Soha-Route-ID": components["headers"]["SohaRouteIDHeader"];
                    "X-Soha-Upstream-ID": components["headers"]["SohaSelectedUpstreamIDHeader"];
                    "X-Soha-Provider-Kind": components["headers"]["SohaProviderKindHeader"];
                    "X-Soha-Public-Model": components["headers"]["SohaPublicModelHeader"];
                    "X-Soha-Upstream-Model": components["headers"]["SohaUpstreamModelHeader"];
                    "X-Soha-Relay-Endpoint": components["headers"]["SohaRelayEndpointHeader"];
                    "X-Soha-Relay-Stream": components["headers"]["SohaRelayStreamHeader"];
                    "X-Soha-Upstream-Status": components["headers"]["SohaUpstreamStatusHeader"];
                    "X-Soha-Cache-Status": components["headers"]["SohaCacheStatusHeader"];
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["NativeProviderObject"];
                    "text/event-stream": string;
                };
            };
            401: components["responses"]["Error"];
            403: components["responses"]["Error"];
        };
    };
    createAIGatewayOpenAIResponse: {
        parameters: {
            query?: never;
            header?: {
                /** @description Optional debug override for administrators or tokens explicitly allowed to select an upstream. */
                "X-Soha-Upstream-ID"?: components["parameters"]["SohaUpstreamIDHeader"];
                /** @description When true, returns redacted Soha route trace response headers for relay managers or tokens with allowRouteTrace explicitly enabled. */
                "X-Soha-Route-Trace"?: components["parameters"]["SohaRouteTraceHeader"];
                /** @description Optional relay cache behavior hint. */
                "X-Soha-Cache-Mode"?: components["parameters"]["SohaCacheModeHeader"];
            };
            path?: never;
            cookie?: never;
        };
        requestBody: {
            content: {
                "application/json": components["schemas"]["OpenAIResponseRequest"];
            };
        };
        responses: {
            /** @description Native OpenAI-compatible Responses JSON or SSE stream. */
            200: {
                headers: {
                    "X-Soha-Route-ID": components["headers"]["SohaRouteIDHeader"];
                    "X-Soha-Upstream-ID": components["headers"]["SohaSelectedUpstreamIDHeader"];
                    "X-Soha-Provider-Kind": components["headers"]["SohaProviderKindHeader"];
                    "X-Soha-Public-Model": components["headers"]["SohaPublicModelHeader"];
                    "X-Soha-Upstream-Model": components["headers"]["SohaUpstreamModelHeader"];
                    "X-Soha-Relay-Endpoint": components["headers"]["SohaRelayEndpointHeader"];
                    "X-Soha-Relay-Stream": components["headers"]["SohaRelayStreamHeader"];
                    "X-Soha-Upstream-Status": components["headers"]["SohaUpstreamStatusHeader"];
                    "X-Soha-Cache-Status": components["headers"]["SohaCacheStatusHeader"];
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["NativeProviderObject"];
                    "text/event-stream": string;
                };
            };
            401: components["responses"]["Error"];
            403: components["responses"]["Error"];
        };
    };
    createAIGatewayOpenAIEmbedding: {
        parameters: {
            query?: never;
            header?: {
                /** @description Optional debug override for administrators or tokens explicitly allowed to select an upstream. */
                "X-Soha-Upstream-ID"?: components["parameters"]["SohaUpstreamIDHeader"];
                /** @description When true, returns redacted Soha route trace response headers for relay managers or tokens with allowRouteTrace explicitly enabled. */
                "X-Soha-Route-Trace"?: components["parameters"]["SohaRouteTraceHeader"];
                /** @description Optional relay cache behavior hint. */
                "X-Soha-Cache-Mode"?: components["parameters"]["SohaCacheModeHeader"];
            };
            path?: never;
            cookie?: never;
        };
        requestBody: {
            content: {
                "application/json": components["schemas"]["OpenAIEmbeddingRequest"];
            };
        };
        responses: {
            /** @description Native OpenAI-compatible embeddings response. This endpoint does not use the OpenSoha envelope. */
            200: {
                headers: {
                    "X-Soha-Route-ID": components["headers"]["SohaRouteIDHeader"];
                    "X-Soha-Upstream-ID": components["headers"]["SohaSelectedUpstreamIDHeader"];
                    "X-Soha-Provider-Kind": components["headers"]["SohaProviderKindHeader"];
                    "X-Soha-Public-Model": components["headers"]["SohaPublicModelHeader"];
                    "X-Soha-Upstream-Model": components["headers"]["SohaUpstreamModelHeader"];
                    "X-Soha-Relay-Endpoint": components["headers"]["SohaRelayEndpointHeader"];
                    "X-Soha-Relay-Stream": components["headers"]["SohaRelayStreamHeader"];
                    "X-Soha-Upstream-Status": components["headers"]["SohaUpstreamStatusHeader"];
                    "X-Soha-Cache-Status": components["headers"]["SohaCacheStatusHeader"];
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["NativeProviderObject"];
                };
            };
            401: components["responses"]["Error"];
            403: components["responses"]["Error"];
        };
    };
    createAIGatewayOpenAIImageGeneration: {
        parameters: {
            query?: never;
            header?: {
                /** @description Optional debug override for administrators or tokens explicitly allowed to select an upstream. */
                "X-Soha-Upstream-ID"?: components["parameters"]["SohaUpstreamIDHeader"];
                /** @description When true, returns redacted Soha route trace response headers for relay managers or tokens with allowRouteTrace explicitly enabled. */
                "X-Soha-Route-Trace"?: components["parameters"]["SohaRouteTraceHeader"];
                /** @description Optional relay cache behavior hint. */
                "X-Soha-Cache-Mode"?: components["parameters"]["SohaCacheModeHeader"];
            };
            path?: never;
            cookie?: never;
        };
        requestBody: {
            content: {
                "application/json": components["schemas"]["OpenAIImageGenerationRequest"];
            };
        };
        responses: {
            /** @description Native OpenAI-compatible image generation response. This endpoint does not use the OpenSoha envelope. */
            200: {
                headers: {
                    "X-Soha-Route-ID": components["headers"]["SohaRouteIDHeader"];
                    "X-Soha-Upstream-ID": components["headers"]["SohaSelectedUpstreamIDHeader"];
                    "X-Soha-Provider-Kind": components["headers"]["SohaProviderKindHeader"];
                    "X-Soha-Public-Model": components["headers"]["SohaPublicModelHeader"];
                    "X-Soha-Upstream-Model": components["headers"]["SohaUpstreamModelHeader"];
                    "X-Soha-Relay-Endpoint": components["headers"]["SohaRelayEndpointHeader"];
                    "X-Soha-Relay-Stream": components["headers"]["SohaRelayStreamHeader"];
                    "X-Soha-Upstream-Status": components["headers"]["SohaUpstreamStatusHeader"];
                    "X-Soha-Cache-Status": components["headers"]["SohaCacheStatusHeader"];
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["NativeProviderObject"];
                };
            };
            401: components["responses"]["Error"];
            403: components["responses"]["Error"];
        };
    };
    createAIGatewayOpenAIImageEdit: {
        parameters: {
            query?: never;
            header?: {
                /** @description Optional debug override for administrators or tokens explicitly allowed to select an upstream. */
                "X-Soha-Upstream-ID"?: components["parameters"]["SohaUpstreamIDHeader"];
                /** @description When true, returns redacted Soha route trace response headers for relay managers or tokens with allowRouteTrace explicitly enabled. */
                "X-Soha-Route-Trace"?: components["parameters"]["SohaRouteTraceHeader"];
                /** @description Optional relay cache behavior hint. */
                "X-Soha-Cache-Mode"?: components["parameters"]["SohaCacheModeHeader"];
            };
            path?: never;
            cookie?: never;
        };
        requestBody: {
            content: {
                "multipart/form-data": components["schemas"]["OpenAIImageEditRequest"];
            };
        };
        responses: {
            /** @description Native OpenAI-compatible image edit response. This endpoint does not use the OpenSoha envelope. */
            200: {
                headers: {
                    "X-Soha-Route-ID": components["headers"]["SohaRouteIDHeader"];
                    "X-Soha-Upstream-ID": components["headers"]["SohaSelectedUpstreamIDHeader"];
                    "X-Soha-Provider-Kind": components["headers"]["SohaProviderKindHeader"];
                    "X-Soha-Public-Model": components["headers"]["SohaPublicModelHeader"];
                    "X-Soha-Upstream-Model": components["headers"]["SohaUpstreamModelHeader"];
                    "X-Soha-Relay-Endpoint": components["headers"]["SohaRelayEndpointHeader"];
                    "X-Soha-Relay-Stream": components["headers"]["SohaRelayStreamHeader"];
                    "X-Soha-Upstream-Status": components["headers"]["SohaUpstreamStatusHeader"];
                    "X-Soha-Cache-Status": components["headers"]["SohaCacheStatusHeader"];
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["NativeProviderObject"];
                };
            };
            401: components["responses"]["Error"];
            403: components["responses"]["Error"];
        };
    };
    createAIGatewayOpenAIImageVariation: {
        parameters: {
            query?: never;
            header?: {
                /** @description Optional debug override for administrators or tokens explicitly allowed to select an upstream. */
                "X-Soha-Upstream-ID"?: components["parameters"]["SohaUpstreamIDHeader"];
                /** @description When true, returns redacted Soha route trace response headers for relay managers or tokens with allowRouteTrace explicitly enabled. */
                "X-Soha-Route-Trace"?: components["parameters"]["SohaRouteTraceHeader"];
                /** @description Optional relay cache behavior hint. */
                "X-Soha-Cache-Mode"?: components["parameters"]["SohaCacheModeHeader"];
            };
            path?: never;
            cookie?: never;
        };
        requestBody: {
            content: {
                "multipart/form-data": components["schemas"]["OpenAIImageVariationRequest"];
            };
        };
        responses: {
            /** @description Native OpenAI-compatible image variation response. This endpoint does not use the OpenSoha envelope. */
            200: {
                headers: {
                    "X-Soha-Route-ID": components["headers"]["SohaRouteIDHeader"];
                    "X-Soha-Upstream-ID": components["headers"]["SohaSelectedUpstreamIDHeader"];
                    "X-Soha-Provider-Kind": components["headers"]["SohaProviderKindHeader"];
                    "X-Soha-Public-Model": components["headers"]["SohaPublicModelHeader"];
                    "X-Soha-Upstream-Model": components["headers"]["SohaUpstreamModelHeader"];
                    "X-Soha-Relay-Endpoint": components["headers"]["SohaRelayEndpointHeader"];
                    "X-Soha-Relay-Stream": components["headers"]["SohaRelayStreamHeader"];
                    "X-Soha-Upstream-Status": components["headers"]["SohaUpstreamStatusHeader"];
                    "X-Soha-Cache-Status": components["headers"]["SohaCacheStatusHeader"];
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["NativeProviderObject"];
                };
            };
            401: components["responses"]["Error"];
            403: components["responses"]["Error"];
        };
    };
    createAIGatewayOpenAIAudioSpeech: {
        parameters: {
            query?: never;
            header?: {
                /** @description Optional debug override for administrators or tokens explicitly allowed to select an upstream. */
                "X-Soha-Upstream-ID"?: components["parameters"]["SohaUpstreamIDHeader"];
                /** @description When true, returns redacted Soha route trace response headers for relay managers or tokens with allowRouteTrace explicitly enabled. */
                "X-Soha-Route-Trace"?: components["parameters"]["SohaRouteTraceHeader"];
                /** @description Optional relay cache behavior hint. */
                "X-Soha-Cache-Mode"?: components["parameters"]["SohaCacheModeHeader"];
            };
            path?: never;
            cookie?: never;
        };
        requestBody: {
            content: {
                "application/json": components["schemas"]["OpenAIAudioSpeechRequest"];
            };
        };
        responses: {
            /** @description Native OpenAI-compatible audio speech response. This endpoint returns binary audio and does not use the OpenSoha envelope. */
            200: {
                headers: {
                    "X-Soha-Route-ID": components["headers"]["SohaRouteIDHeader"];
                    "X-Soha-Upstream-ID": components["headers"]["SohaSelectedUpstreamIDHeader"];
                    "X-Soha-Provider-Kind": components["headers"]["SohaProviderKindHeader"];
                    "X-Soha-Public-Model": components["headers"]["SohaPublicModelHeader"];
                    "X-Soha-Upstream-Model": components["headers"]["SohaUpstreamModelHeader"];
                    "X-Soha-Relay-Endpoint": components["headers"]["SohaRelayEndpointHeader"];
                    "X-Soha-Relay-Stream": components["headers"]["SohaRelayStreamHeader"];
                    "X-Soha-Upstream-Status": components["headers"]["SohaUpstreamStatusHeader"];
                    "X-Soha-Cache-Status": components["headers"]["SohaCacheStatusHeader"];
                    [name: string]: unknown;
                };
                content: {
                    "application/octet-stream": string;
                    "audio/mpeg": string;
                    "audio/wav": string;
                    "audio/ogg": string;
                    "audio/flac": string;
                };
            };
            401: components["responses"]["Error"];
            403: components["responses"]["Error"];
        };
    };
    createAIGatewayOpenAIAudioTranscription: {
        parameters: {
            query?: never;
            header?: {
                /** @description Optional debug override for administrators or tokens explicitly allowed to select an upstream. */
                "X-Soha-Upstream-ID"?: components["parameters"]["SohaUpstreamIDHeader"];
                /** @description When true, returns redacted Soha route trace response headers for relay managers or tokens with allowRouteTrace explicitly enabled. */
                "X-Soha-Route-Trace"?: components["parameters"]["SohaRouteTraceHeader"];
                /** @description Optional relay cache behavior hint. */
                "X-Soha-Cache-Mode"?: components["parameters"]["SohaCacheModeHeader"];
            };
            path?: never;
            cookie?: never;
        };
        requestBody: {
            content: {
                "multipart/form-data": components["schemas"]["OpenAIAudioTranscriptionRequest"];
            };
        };
        responses: {
            /** @description Native OpenAI-compatible audio transcription response. This endpoint does not use the OpenSoha envelope. */
            200: {
                headers: {
                    "X-Soha-Route-ID": components["headers"]["SohaRouteIDHeader"];
                    "X-Soha-Upstream-ID": components["headers"]["SohaSelectedUpstreamIDHeader"];
                    "X-Soha-Provider-Kind": components["headers"]["SohaProviderKindHeader"];
                    "X-Soha-Public-Model": components["headers"]["SohaPublicModelHeader"];
                    "X-Soha-Upstream-Model": components["headers"]["SohaUpstreamModelHeader"];
                    "X-Soha-Relay-Endpoint": components["headers"]["SohaRelayEndpointHeader"];
                    "X-Soha-Relay-Stream": components["headers"]["SohaRelayStreamHeader"];
                    "X-Soha-Upstream-Status": components["headers"]["SohaUpstreamStatusHeader"];
                    "X-Soha-Cache-Status": components["headers"]["SohaCacheStatusHeader"];
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["NativeProviderObject"];
                    "text/plain": string;
                    "text/vtt": string;
                    "application/x-subrip": string;
                };
            };
            401: components["responses"]["Error"];
            403: components["responses"]["Error"];
        };
    };
    createAIGatewayOpenAIAudioTranslation: {
        parameters: {
            query?: never;
            header?: {
                /** @description Optional debug override for administrators or tokens explicitly allowed to select an upstream. */
                "X-Soha-Upstream-ID"?: components["parameters"]["SohaUpstreamIDHeader"];
                /** @description When true, returns redacted Soha route trace response headers for relay managers or tokens with allowRouteTrace explicitly enabled. */
                "X-Soha-Route-Trace"?: components["parameters"]["SohaRouteTraceHeader"];
                /** @description Optional relay cache behavior hint. */
                "X-Soha-Cache-Mode"?: components["parameters"]["SohaCacheModeHeader"];
            };
            path?: never;
            cookie?: never;
        };
        requestBody: {
            content: {
                "multipart/form-data": components["schemas"]["OpenAIAudioTranslationRequest"];
            };
        };
        responses: {
            /** @description Native OpenAI-compatible audio translation response. This endpoint does not use the OpenSoha envelope. */
            200: {
                headers: {
                    "X-Soha-Route-ID": components["headers"]["SohaRouteIDHeader"];
                    "X-Soha-Upstream-ID": components["headers"]["SohaSelectedUpstreamIDHeader"];
                    "X-Soha-Provider-Kind": components["headers"]["SohaProviderKindHeader"];
                    "X-Soha-Public-Model": components["headers"]["SohaPublicModelHeader"];
                    "X-Soha-Upstream-Model": components["headers"]["SohaUpstreamModelHeader"];
                    "X-Soha-Relay-Endpoint": components["headers"]["SohaRelayEndpointHeader"];
                    "X-Soha-Relay-Stream": components["headers"]["SohaRelayStreamHeader"];
                    "X-Soha-Upstream-Status": components["headers"]["SohaUpstreamStatusHeader"];
                    "X-Soha-Cache-Status": components["headers"]["SohaCacheStatusHeader"];
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["NativeProviderObject"];
                    "text/plain": string;
                    "text/vtt": string;
                    "application/x-subrip": string;
                };
            };
            401: components["responses"]["Error"];
            403: components["responses"]["Error"];
        };
    };
    connectAIGatewayOpenAIRealtime: {
        parameters: {
            query: {
                /** @description Public model name to route to an upstream OpenAI Realtime model. */
                model: string;
            };
            header?: {
                /** @description Optional debug override for administrators or tokens explicitly allowed to select an upstream. */
                "X-Soha-Upstream-ID"?: components["parameters"]["SohaUpstreamIDHeader"];
                /** @description When true, returns redacted Soha route trace response headers for relay managers or tokens with allowRouteTrace explicitly enabled. */
                "X-Soha-Route-Trace"?: components["parameters"]["SohaRouteTraceHeader"];
            };
            path?: never;
            cookie?: never;
        };
        requestBody?: never;
        responses: {
            /** @description WebSocket upgrade accepted. The connection proxies OpenAI Realtime frames without wrapping provider messages in an OpenSoha envelope. */
            101: {
                headers: {
                    "X-Soha-Route-ID": components["headers"]["SohaRouteIDHeader"];
                    "X-Soha-Upstream-ID": components["headers"]["SohaSelectedUpstreamIDHeader"];
                    "X-Soha-Provider-Kind": components["headers"]["SohaProviderKindHeader"];
                    "X-Soha-Public-Model": components["headers"]["SohaPublicModelHeader"];
                    "X-Soha-Upstream-Model": components["headers"]["SohaUpstreamModelHeader"];
                    "X-Soha-Relay-Endpoint": components["headers"]["SohaRelayEndpointHeader"];
                    "X-Soha-Relay-Stream": components["headers"]["SohaRelayStreamHeader"];
                    "X-Soha-Upstream-Status": components["headers"]["SohaUpstreamStatusHeader"];
                    "X-Soha-Cache-Status": components["headers"]["SohaCacheStatusHeader"];
                    [name: string]: unknown;
                };
                content?: never;
            };
            400: components["responses"]["Error"];
            401: components["responses"]["Error"];
            403: components["responses"]["Error"];
        };
    };
    listAIGatewayOpenAICompatibleProviderModels: {
        parameters: {
            query?: never;
            header?: never;
            path: {
                /** @description First-class OpenAI-compatible relay provider kind. Azure OpenAI uses Azure data-plane URL and api-key header adaptation upstream. */
                openaiCompatibleProvider: components["parameters"]["OpenAICompatibleProvider"];
            };
            cookie?: never;
        };
        requestBody?: never;
        responses: {
            /** @description Native OpenAI-compatible models response for DeepSeek, Qwen, OpenRouter, or Azure OpenAI. This endpoint does not use the OpenSoha envelope. */
            200: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["OpenAIModelsResponse"];
                };
            };
            401: components["responses"]["Error"];
            403: components["responses"]["Error"];
        };
    };
    createAIGatewayOpenAICompatibleProviderChatCompletion: {
        parameters: {
            query?: never;
            header?: {
                /** @description Optional debug override for administrators or tokens explicitly allowed to select an upstream. */
                "X-Soha-Upstream-ID"?: components["parameters"]["SohaUpstreamIDHeader"];
                /** @description When true, returns redacted Soha route trace response headers for relay managers or tokens with allowRouteTrace explicitly enabled. */
                "X-Soha-Route-Trace"?: components["parameters"]["SohaRouteTraceHeader"];
                /** @description Optional relay cache behavior hint. */
                "X-Soha-Cache-Mode"?: components["parameters"]["SohaCacheModeHeader"];
            };
            path: {
                /** @description First-class OpenAI-compatible relay provider kind. Azure OpenAI uses Azure data-plane URL and api-key header adaptation upstream. */
                openaiCompatibleProvider: components["parameters"]["OpenAICompatibleProvider"];
            };
            cookie?: never;
        };
        requestBody: {
            content: {
                "application/json": components["schemas"]["OpenAIChatCompletionRequest"];
            };
        };
        responses: {
            /** @description Native OpenAI-compatible chat completion JSON or SSE stream for DeepSeek, Qwen, OpenRouter, or Azure OpenAI. */
            200: {
                headers: {
                    "X-Soha-Route-ID": components["headers"]["SohaRouteIDHeader"];
                    "X-Soha-Upstream-ID": components["headers"]["SohaSelectedUpstreamIDHeader"];
                    "X-Soha-Provider-Kind": components["headers"]["SohaProviderKindHeader"];
                    "X-Soha-Public-Model": components["headers"]["SohaPublicModelHeader"];
                    "X-Soha-Upstream-Model": components["headers"]["SohaUpstreamModelHeader"];
                    "X-Soha-Relay-Endpoint": components["headers"]["SohaRelayEndpointHeader"];
                    "X-Soha-Relay-Stream": components["headers"]["SohaRelayStreamHeader"];
                    "X-Soha-Upstream-Status": components["headers"]["SohaUpstreamStatusHeader"];
                    "X-Soha-Cache-Status": components["headers"]["SohaCacheStatusHeader"];
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["NativeProviderObject"];
                    "text/event-stream": string;
                };
            };
            401: components["responses"]["Error"];
            403: components["responses"]["Error"];
        };
    };
    createAIGatewayOpenAICompatibleProviderResponse: {
        parameters: {
            query?: never;
            header?: {
                /** @description Optional debug override for administrators or tokens explicitly allowed to select an upstream. */
                "X-Soha-Upstream-ID"?: components["parameters"]["SohaUpstreamIDHeader"];
                /** @description When true, returns redacted Soha route trace response headers for relay managers or tokens with allowRouteTrace explicitly enabled. */
                "X-Soha-Route-Trace"?: components["parameters"]["SohaRouteTraceHeader"];
                /** @description Optional relay cache behavior hint. */
                "X-Soha-Cache-Mode"?: components["parameters"]["SohaCacheModeHeader"];
            };
            path: {
                /** @description First-class OpenAI-compatible relay provider kind. Azure OpenAI uses Azure data-plane URL and api-key header adaptation upstream. */
                openaiCompatibleProvider: components["parameters"]["OpenAICompatibleProvider"];
            };
            cookie?: never;
        };
        requestBody: {
            content: {
                "application/json": components["schemas"]["OpenAIResponseRequest"];
            };
        };
        responses: {
            /** @description Native OpenAI-compatible Responses JSON or SSE stream for DeepSeek, Qwen, OpenRouter, or Azure OpenAI. */
            200: {
                headers: {
                    "X-Soha-Route-ID": components["headers"]["SohaRouteIDHeader"];
                    "X-Soha-Upstream-ID": components["headers"]["SohaSelectedUpstreamIDHeader"];
                    "X-Soha-Provider-Kind": components["headers"]["SohaProviderKindHeader"];
                    "X-Soha-Public-Model": components["headers"]["SohaPublicModelHeader"];
                    "X-Soha-Upstream-Model": components["headers"]["SohaUpstreamModelHeader"];
                    "X-Soha-Relay-Endpoint": components["headers"]["SohaRelayEndpointHeader"];
                    "X-Soha-Relay-Stream": components["headers"]["SohaRelayStreamHeader"];
                    "X-Soha-Upstream-Status": components["headers"]["SohaUpstreamStatusHeader"];
                    "X-Soha-Cache-Status": components["headers"]["SohaCacheStatusHeader"];
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["NativeProviderObject"];
                    "text/event-stream": string;
                };
            };
            401: components["responses"]["Error"];
            403: components["responses"]["Error"];
        };
    };
    createAIGatewayOpenAICompatibleProviderEmbedding: {
        parameters: {
            query?: never;
            header?: {
                /** @description Optional debug override for administrators or tokens explicitly allowed to select an upstream. */
                "X-Soha-Upstream-ID"?: components["parameters"]["SohaUpstreamIDHeader"];
                /** @description When true, returns redacted Soha route trace response headers for relay managers or tokens with allowRouteTrace explicitly enabled. */
                "X-Soha-Route-Trace"?: components["parameters"]["SohaRouteTraceHeader"];
                /** @description Optional relay cache behavior hint. */
                "X-Soha-Cache-Mode"?: components["parameters"]["SohaCacheModeHeader"];
            };
            path: {
                /** @description First-class OpenAI-compatible relay provider kind. Azure OpenAI uses Azure data-plane URL and api-key header adaptation upstream. */
                openaiCompatibleProvider: components["parameters"]["OpenAICompatibleProvider"];
            };
            cookie?: never;
        };
        requestBody: {
            content: {
                "application/json": components["schemas"]["OpenAIEmbeddingRequest"];
            };
        };
        responses: {
            /** @description Native OpenAI-compatible embeddings response for DeepSeek, Qwen, OpenRouter, or Azure OpenAI. This endpoint does not use the OpenSoha envelope. */
            200: {
                headers: {
                    "X-Soha-Route-ID": components["headers"]["SohaRouteIDHeader"];
                    "X-Soha-Upstream-ID": components["headers"]["SohaSelectedUpstreamIDHeader"];
                    "X-Soha-Provider-Kind": components["headers"]["SohaProviderKindHeader"];
                    "X-Soha-Public-Model": components["headers"]["SohaPublicModelHeader"];
                    "X-Soha-Upstream-Model": components["headers"]["SohaUpstreamModelHeader"];
                    "X-Soha-Relay-Endpoint": components["headers"]["SohaRelayEndpointHeader"];
                    "X-Soha-Relay-Stream": components["headers"]["SohaRelayStreamHeader"];
                    "X-Soha-Upstream-Status": components["headers"]["SohaUpstreamStatusHeader"];
                    "X-Soha-Cache-Status": components["headers"]["SohaCacheStatusHeader"];
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["NativeProviderObject"];
                };
            };
            401: components["responses"]["Error"];
            403: components["responses"]["Error"];
        };
    };
    createAIGatewayOpenAICompatibleProviderImageGeneration: {
        parameters: {
            query?: never;
            header?: {
                /** @description Optional debug override for administrators or tokens explicitly allowed to select an upstream. */
                "X-Soha-Upstream-ID"?: components["parameters"]["SohaUpstreamIDHeader"];
                /** @description When true, returns redacted Soha route trace response headers for relay managers or tokens with allowRouteTrace explicitly enabled. */
                "X-Soha-Route-Trace"?: components["parameters"]["SohaRouteTraceHeader"];
                /** @description Optional relay cache behavior hint. */
                "X-Soha-Cache-Mode"?: components["parameters"]["SohaCacheModeHeader"];
            };
            path: {
                /** @description First-class OpenAI-compatible relay provider kind. Azure OpenAI uses Azure data-plane URL and api-key header adaptation upstream. */
                openaiCompatibleProvider: components["parameters"]["OpenAICompatibleProvider"];
            };
            cookie?: never;
        };
        requestBody: {
            content: {
                "application/json": components["schemas"]["OpenAIImageGenerationRequest"];
            };
        };
        responses: {
            /** @description Native OpenAI-compatible image generation response for DeepSeek, Qwen, OpenRouter, or Azure OpenAI. This endpoint does not use the OpenSoha envelope. */
            200: {
                headers: {
                    "X-Soha-Route-ID": components["headers"]["SohaRouteIDHeader"];
                    "X-Soha-Upstream-ID": components["headers"]["SohaSelectedUpstreamIDHeader"];
                    "X-Soha-Provider-Kind": components["headers"]["SohaProviderKindHeader"];
                    "X-Soha-Public-Model": components["headers"]["SohaPublicModelHeader"];
                    "X-Soha-Upstream-Model": components["headers"]["SohaUpstreamModelHeader"];
                    "X-Soha-Relay-Endpoint": components["headers"]["SohaRelayEndpointHeader"];
                    "X-Soha-Relay-Stream": components["headers"]["SohaRelayStreamHeader"];
                    "X-Soha-Upstream-Status": components["headers"]["SohaUpstreamStatusHeader"];
                    "X-Soha-Cache-Status": components["headers"]["SohaCacheStatusHeader"];
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["NativeProviderObject"];
                };
            };
            401: components["responses"]["Error"];
            403: components["responses"]["Error"];
        };
    };
    createAIGatewayOpenAICompatibleProviderImageEdit: {
        parameters: {
            query?: never;
            header?: {
                /** @description Optional debug override for administrators or tokens explicitly allowed to select an upstream. */
                "X-Soha-Upstream-ID"?: components["parameters"]["SohaUpstreamIDHeader"];
                /** @description When true, returns redacted Soha route trace response headers for relay managers or tokens with allowRouteTrace explicitly enabled. */
                "X-Soha-Route-Trace"?: components["parameters"]["SohaRouteTraceHeader"];
                /** @description Optional relay cache behavior hint. */
                "X-Soha-Cache-Mode"?: components["parameters"]["SohaCacheModeHeader"];
            };
            path: {
                /** @description First-class OpenAI-compatible relay provider kind. Azure OpenAI uses Azure data-plane URL and api-key header adaptation upstream. */
                openaiCompatibleProvider: components["parameters"]["OpenAICompatibleProvider"];
            };
            cookie?: never;
        };
        requestBody: {
            content: {
                "multipart/form-data": components["schemas"]["OpenAIImageEditRequest"];
            };
        };
        responses: {
            /** @description Native OpenAI-compatible image edit response for DeepSeek, Qwen, OpenRouter, or Azure OpenAI. This endpoint does not use the OpenSoha envelope. */
            200: {
                headers: {
                    "X-Soha-Route-ID": components["headers"]["SohaRouteIDHeader"];
                    "X-Soha-Upstream-ID": components["headers"]["SohaSelectedUpstreamIDHeader"];
                    "X-Soha-Provider-Kind": components["headers"]["SohaProviderKindHeader"];
                    "X-Soha-Public-Model": components["headers"]["SohaPublicModelHeader"];
                    "X-Soha-Upstream-Model": components["headers"]["SohaUpstreamModelHeader"];
                    "X-Soha-Relay-Endpoint": components["headers"]["SohaRelayEndpointHeader"];
                    "X-Soha-Relay-Stream": components["headers"]["SohaRelayStreamHeader"];
                    "X-Soha-Upstream-Status": components["headers"]["SohaUpstreamStatusHeader"];
                    "X-Soha-Cache-Status": components["headers"]["SohaCacheStatusHeader"];
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["NativeProviderObject"];
                };
            };
            401: components["responses"]["Error"];
            403: components["responses"]["Error"];
        };
    };
    createAIGatewayOpenAICompatibleProviderImageVariation: {
        parameters: {
            query?: never;
            header?: {
                /** @description Optional debug override for administrators or tokens explicitly allowed to select an upstream. */
                "X-Soha-Upstream-ID"?: components["parameters"]["SohaUpstreamIDHeader"];
                /** @description When true, returns redacted Soha route trace response headers for relay managers or tokens with allowRouteTrace explicitly enabled. */
                "X-Soha-Route-Trace"?: components["parameters"]["SohaRouteTraceHeader"];
                /** @description Optional relay cache behavior hint. */
                "X-Soha-Cache-Mode"?: components["parameters"]["SohaCacheModeHeader"];
            };
            path: {
                /** @description First-class OpenAI-compatible relay provider kind. Azure OpenAI uses Azure data-plane URL and api-key header adaptation upstream. */
                openaiCompatibleProvider: components["parameters"]["OpenAICompatibleProvider"];
            };
            cookie?: never;
        };
        requestBody: {
            content: {
                "multipart/form-data": components["schemas"]["OpenAIImageVariationRequest"];
            };
        };
        responses: {
            /** @description Native OpenAI-compatible image variation response for DeepSeek, Qwen, OpenRouter, or Azure OpenAI. This endpoint does not use the OpenSoha envelope. */
            200: {
                headers: {
                    "X-Soha-Route-ID": components["headers"]["SohaRouteIDHeader"];
                    "X-Soha-Upstream-ID": components["headers"]["SohaSelectedUpstreamIDHeader"];
                    "X-Soha-Provider-Kind": components["headers"]["SohaProviderKindHeader"];
                    "X-Soha-Public-Model": components["headers"]["SohaPublicModelHeader"];
                    "X-Soha-Upstream-Model": components["headers"]["SohaUpstreamModelHeader"];
                    "X-Soha-Relay-Endpoint": components["headers"]["SohaRelayEndpointHeader"];
                    "X-Soha-Relay-Stream": components["headers"]["SohaRelayStreamHeader"];
                    "X-Soha-Upstream-Status": components["headers"]["SohaUpstreamStatusHeader"];
                    "X-Soha-Cache-Status": components["headers"]["SohaCacheStatusHeader"];
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["NativeProviderObject"];
                };
            };
            401: components["responses"]["Error"];
            403: components["responses"]["Error"];
        };
    };
    createAIGatewayOpenAICompatibleProviderAudioSpeech: {
        parameters: {
            query?: never;
            header?: {
                /** @description Optional debug override for administrators or tokens explicitly allowed to select an upstream. */
                "X-Soha-Upstream-ID"?: components["parameters"]["SohaUpstreamIDHeader"];
                /** @description When true, returns redacted Soha route trace response headers for relay managers or tokens with allowRouteTrace explicitly enabled. */
                "X-Soha-Route-Trace"?: components["parameters"]["SohaRouteTraceHeader"];
                /** @description Optional relay cache behavior hint. */
                "X-Soha-Cache-Mode"?: components["parameters"]["SohaCacheModeHeader"];
            };
            path: {
                /** @description First-class OpenAI-compatible relay provider kind. Azure OpenAI uses Azure data-plane URL and api-key header adaptation upstream. */
                openaiCompatibleProvider: components["parameters"]["OpenAICompatibleProvider"];
            };
            cookie?: never;
        };
        requestBody: {
            content: {
                "application/json": components["schemas"]["OpenAIAudioSpeechRequest"];
            };
        };
        responses: {
            /** @description Native OpenAI-compatible audio speech response for DeepSeek, Qwen, OpenRouter, or Azure OpenAI. This endpoint returns binary audio and does not use the OpenSoha envelope. */
            200: {
                headers: {
                    "X-Soha-Route-ID": components["headers"]["SohaRouteIDHeader"];
                    "X-Soha-Upstream-ID": components["headers"]["SohaSelectedUpstreamIDHeader"];
                    "X-Soha-Provider-Kind": components["headers"]["SohaProviderKindHeader"];
                    "X-Soha-Public-Model": components["headers"]["SohaPublicModelHeader"];
                    "X-Soha-Upstream-Model": components["headers"]["SohaUpstreamModelHeader"];
                    "X-Soha-Relay-Endpoint": components["headers"]["SohaRelayEndpointHeader"];
                    "X-Soha-Relay-Stream": components["headers"]["SohaRelayStreamHeader"];
                    "X-Soha-Upstream-Status": components["headers"]["SohaUpstreamStatusHeader"];
                    "X-Soha-Cache-Status": components["headers"]["SohaCacheStatusHeader"];
                    [name: string]: unknown;
                };
                content: {
                    "application/octet-stream": string;
                    "audio/mpeg": string;
                    "audio/wav": string;
                    "audio/ogg": string;
                    "audio/flac": string;
                };
            };
            401: components["responses"]["Error"];
            403: components["responses"]["Error"];
        };
    };
    createAIGatewayOpenAICompatibleProviderAudioTranscription: {
        parameters: {
            query?: never;
            header?: {
                /** @description Optional debug override for administrators or tokens explicitly allowed to select an upstream. */
                "X-Soha-Upstream-ID"?: components["parameters"]["SohaUpstreamIDHeader"];
                /** @description When true, returns redacted Soha route trace response headers for relay managers or tokens with allowRouteTrace explicitly enabled. */
                "X-Soha-Route-Trace"?: components["parameters"]["SohaRouteTraceHeader"];
                /** @description Optional relay cache behavior hint. */
                "X-Soha-Cache-Mode"?: components["parameters"]["SohaCacheModeHeader"];
            };
            path: {
                /** @description First-class OpenAI-compatible relay provider kind. Azure OpenAI uses Azure data-plane URL and api-key header adaptation upstream. */
                openaiCompatibleProvider: components["parameters"]["OpenAICompatibleProvider"];
            };
            cookie?: never;
        };
        requestBody: {
            content: {
                "multipart/form-data": components["schemas"]["OpenAIAudioTranscriptionRequest"];
            };
        };
        responses: {
            /** @description Native OpenAI-compatible audio transcription response for DeepSeek, Qwen, OpenRouter, or Azure OpenAI. This endpoint does not use the OpenSoha envelope. */
            200: {
                headers: {
                    "X-Soha-Route-ID": components["headers"]["SohaRouteIDHeader"];
                    "X-Soha-Upstream-ID": components["headers"]["SohaSelectedUpstreamIDHeader"];
                    "X-Soha-Provider-Kind": components["headers"]["SohaProviderKindHeader"];
                    "X-Soha-Public-Model": components["headers"]["SohaPublicModelHeader"];
                    "X-Soha-Upstream-Model": components["headers"]["SohaUpstreamModelHeader"];
                    "X-Soha-Relay-Endpoint": components["headers"]["SohaRelayEndpointHeader"];
                    "X-Soha-Relay-Stream": components["headers"]["SohaRelayStreamHeader"];
                    "X-Soha-Upstream-Status": components["headers"]["SohaUpstreamStatusHeader"];
                    "X-Soha-Cache-Status": components["headers"]["SohaCacheStatusHeader"];
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["NativeProviderObject"];
                    "text/plain": string;
                    "text/vtt": string;
                    "application/x-subrip": string;
                };
            };
            401: components["responses"]["Error"];
            403: components["responses"]["Error"];
        };
    };
    createAIGatewayOpenAICompatibleProviderAudioTranslation: {
        parameters: {
            query?: never;
            header?: {
                /** @description Optional debug override for administrators or tokens explicitly allowed to select an upstream. */
                "X-Soha-Upstream-ID"?: components["parameters"]["SohaUpstreamIDHeader"];
                /** @description When true, returns redacted Soha route trace response headers for relay managers or tokens with allowRouteTrace explicitly enabled. */
                "X-Soha-Route-Trace"?: components["parameters"]["SohaRouteTraceHeader"];
                /** @description Optional relay cache behavior hint. */
                "X-Soha-Cache-Mode"?: components["parameters"]["SohaCacheModeHeader"];
            };
            path: {
                /** @description First-class OpenAI-compatible relay provider kind. Azure OpenAI uses Azure data-plane URL and api-key header adaptation upstream. */
                openaiCompatibleProvider: components["parameters"]["OpenAICompatibleProvider"];
            };
            cookie?: never;
        };
        requestBody: {
            content: {
                "multipart/form-data": components["schemas"]["OpenAIAudioTranslationRequest"];
            };
        };
        responses: {
            /** @description Native OpenAI-compatible audio translation response for DeepSeek, Qwen, OpenRouter, or Azure OpenAI. This endpoint does not use the OpenSoha envelope. */
            200: {
                headers: {
                    "X-Soha-Route-ID": components["headers"]["SohaRouteIDHeader"];
                    "X-Soha-Upstream-ID": components["headers"]["SohaSelectedUpstreamIDHeader"];
                    "X-Soha-Provider-Kind": components["headers"]["SohaProviderKindHeader"];
                    "X-Soha-Public-Model": components["headers"]["SohaPublicModelHeader"];
                    "X-Soha-Upstream-Model": components["headers"]["SohaUpstreamModelHeader"];
                    "X-Soha-Relay-Endpoint": components["headers"]["SohaRelayEndpointHeader"];
                    "X-Soha-Relay-Stream": components["headers"]["SohaRelayStreamHeader"];
                    "X-Soha-Upstream-Status": components["headers"]["SohaUpstreamStatusHeader"];
                    "X-Soha-Cache-Status": components["headers"]["SohaCacheStatusHeader"];
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["NativeProviderObject"];
                    "text/plain": string;
                    "text/vtt": string;
                    "application/x-subrip": string;
                };
            };
            401: components["responses"]["Error"];
            403: components["responses"]["Error"];
        };
    };
    listAIGatewayGeminiModels: {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        requestBody?: never;
        responses: {
            /** @description Native Gemini-compatible models response. This endpoint does not use the OpenSoha envelope. */
            200: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["GeminiModelsResponse"];
                };
            };
            401: components["responses"]["Error"];
            403: components["responses"]["Error"];
        };
    };
    createAIGatewayGeminiGenerateContent: {
        parameters: {
            query?: never;
            header?: {
                /** @description Optional debug override for administrators or tokens explicitly allowed to select an upstream. */
                "X-Soha-Upstream-ID"?: components["parameters"]["SohaUpstreamIDHeader"];
                /** @description When true, returns redacted Soha route trace response headers for relay managers or tokens with allowRouteTrace explicitly enabled. */
                "X-Soha-Route-Trace"?: components["parameters"]["SohaRouteTraceHeader"];
                /** @description Optional relay cache behavior hint. */
                "X-Soha-Cache-Mode"?: components["parameters"]["SohaCacheModeHeader"];
            };
            path: {
                /** @description Public Gemini model ID exposed by Soha; the relay maps it to the configured upstream model. */
                geminiModel: string;
            };
            cookie?: never;
        };
        requestBody: {
            content: {
                "application/json": components["schemas"]["GeminiGenerateContentRequest"];
            };
        };
        responses: {
            /** @description Native Gemini generateContent JSON response. */
            200: {
                headers: {
                    "X-Soha-Route-ID": components["headers"]["SohaRouteIDHeader"];
                    "X-Soha-Upstream-ID": components["headers"]["SohaSelectedUpstreamIDHeader"];
                    "X-Soha-Provider-Kind": components["headers"]["SohaProviderKindHeader"];
                    "X-Soha-Public-Model": components["headers"]["SohaPublicModelHeader"];
                    "X-Soha-Upstream-Model": components["headers"]["SohaUpstreamModelHeader"];
                    "X-Soha-Relay-Endpoint": components["headers"]["SohaRelayEndpointHeader"];
                    "X-Soha-Relay-Stream": components["headers"]["SohaRelayStreamHeader"];
                    "X-Soha-Upstream-Status": components["headers"]["SohaUpstreamStatusHeader"];
                    "X-Soha-Cache-Status": components["headers"]["SohaCacheStatusHeader"];
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["NativeProviderObject"];
                };
            };
            401: components["responses"]["Error"];
            403: components["responses"]["Error"];
        };
    };
    createAIGatewayGeminiStreamGenerateContent: {
        parameters: {
            query?: never;
            header?: {
                /** @description Optional debug override for administrators or tokens explicitly allowed to select an upstream. */
                "X-Soha-Upstream-ID"?: components["parameters"]["SohaUpstreamIDHeader"];
                /** @description When true, returns redacted Soha route trace response headers for relay managers or tokens with allowRouteTrace explicitly enabled. */
                "X-Soha-Route-Trace"?: components["parameters"]["SohaRouteTraceHeader"];
                /** @description Optional relay cache behavior hint. */
                "X-Soha-Cache-Mode"?: components["parameters"]["SohaCacheModeHeader"];
            };
            path: {
                /** @description Public Gemini model ID exposed by Soha; the relay maps it to the configured upstream model. */
                geminiModel: string;
            };
            cookie?: never;
        };
        requestBody: {
            content: {
                "application/json": components["schemas"]["GeminiGenerateContentRequest"];
            };
        };
        responses: {
            /** @description Native Gemini streamGenerateContent SSE stream. */
            200: {
                headers: {
                    "X-Soha-Route-ID": components["headers"]["SohaRouteIDHeader"];
                    "X-Soha-Upstream-ID": components["headers"]["SohaSelectedUpstreamIDHeader"];
                    "X-Soha-Provider-Kind": components["headers"]["SohaProviderKindHeader"];
                    "X-Soha-Public-Model": components["headers"]["SohaPublicModelHeader"];
                    "X-Soha-Upstream-Model": components["headers"]["SohaUpstreamModelHeader"];
                    "X-Soha-Relay-Endpoint": components["headers"]["SohaRelayEndpointHeader"];
                    "X-Soha-Relay-Stream": components["headers"]["SohaRelayStreamHeader"];
                    "X-Soha-Upstream-Status": components["headers"]["SohaUpstreamStatusHeader"];
                    "X-Soha-Cache-Status": components["headers"]["SohaCacheStatusHeader"];
                    [name: string]: unknown;
                };
                content: {
                    "text/event-stream": string;
                };
            };
            401: components["responses"]["Error"];
            403: components["responses"]["Error"];
        };
    };
    createAIGatewayGeminiInteractions: {
        parameters: {
            query?: never;
            header?: {
                /** @description Optional debug override for administrators or tokens explicitly allowed to select an upstream. */
                "X-Soha-Upstream-ID"?: components["parameters"]["SohaUpstreamIDHeader"];
                /** @description When true, returns redacted Soha route trace response headers for relay managers or tokens with allowRouteTrace explicitly enabled. */
                "X-Soha-Route-Trace"?: components["parameters"]["SohaRouteTraceHeader"];
                /** @description Optional relay cache behavior hint. */
                "X-Soha-Cache-Mode"?: components["parameters"]["SohaCacheModeHeader"];
            };
            path?: never;
            cookie?: never;
        };
        requestBody: {
            content: {
                "application/json": components["schemas"]["GeminiInteractionsRequest"];
            };
        };
        responses: {
            /** @description Native Gemini Interactions JSON response. This endpoint does not use the OpenSoha envelope. */
            200: {
                headers: {
                    "X-Soha-Route-ID": components["headers"]["SohaRouteIDHeader"];
                    "X-Soha-Upstream-ID": components["headers"]["SohaSelectedUpstreamIDHeader"];
                    "X-Soha-Provider-Kind": components["headers"]["SohaProviderKindHeader"];
                    "X-Soha-Public-Model": components["headers"]["SohaPublicModelHeader"];
                    "X-Soha-Upstream-Model": components["headers"]["SohaUpstreamModelHeader"];
                    "X-Soha-Relay-Endpoint": components["headers"]["SohaRelayEndpointHeader"];
                    "X-Soha-Relay-Stream": components["headers"]["SohaRelayStreamHeader"];
                    "X-Soha-Upstream-Status": components["headers"]["SohaUpstreamStatusHeader"];
                    "X-Soha-Cache-Status": components["headers"]["SohaCacheStatusHeader"];
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["NativeProviderObject"];
                };
            };
            401: components["responses"]["Error"];
            403: components["responses"]["Error"];
        };
    };
    createAIGatewayCohereRerank: {
        parameters: {
            query?: never;
            header?: {
                /** @description Optional debug override for administrators or tokens explicitly allowed to select an upstream. */
                "X-Soha-Upstream-ID"?: components["parameters"]["SohaUpstreamIDHeader"];
                /** @description When true, returns redacted Soha route trace response headers for relay managers or tokens with allowRouteTrace explicitly enabled. */
                "X-Soha-Route-Trace"?: components["parameters"]["SohaRouteTraceHeader"];
                /** @description Optional relay cache behavior hint. */
                "X-Soha-Cache-Mode"?: components["parameters"]["SohaCacheModeHeader"];
            };
            path?: never;
            cookie?: never;
        };
        requestBody: {
            content: {
                "application/json": components["schemas"]["CohereRerankRequest"];
            };
        };
        responses: {
            /** @description Native Cohere-compatible rerank JSON response. This endpoint does not use the OpenSoha envelope. */
            200: {
                headers: {
                    "X-Soha-Route-ID": components["headers"]["SohaRouteIDHeader"];
                    "X-Soha-Upstream-ID": components["headers"]["SohaSelectedUpstreamIDHeader"];
                    "X-Soha-Provider-Kind": components["headers"]["SohaProviderKindHeader"];
                    "X-Soha-Public-Model": components["headers"]["SohaPublicModelHeader"];
                    "X-Soha-Upstream-Model": components["headers"]["SohaUpstreamModelHeader"];
                    "X-Soha-Relay-Endpoint": components["headers"]["SohaRelayEndpointHeader"];
                    "X-Soha-Relay-Stream": components["headers"]["SohaRelayStreamHeader"];
                    "X-Soha-Upstream-Status": components["headers"]["SohaUpstreamStatusHeader"];
                    "X-Soha-Cache-Status": components["headers"]["SohaCacheStatusHeader"];
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["NativeProviderObject"];
                };
            };
            401: components["responses"]["Error"];
            403: components["responses"]["Error"];
        };
    };
    listAIGatewayAnthropicModels: {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        requestBody?: never;
        responses: {
            /** @description Native Anthropic-compatible models response. This endpoint does not use the OpenSoha envelope. */
            200: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["AnthropicModelsResponse"];
                };
            };
            401: components["responses"]["Error"];
            403: components["responses"]["Error"];
        };
    };
    createAIGatewayAnthropicMessage: {
        parameters: {
            query?: never;
            header?: {
                /** @description Forwarded Anthropic API version header when present. */
                "anthropic-version"?: string;
                /** @description Optional debug override for administrators or tokens explicitly allowed to select an upstream. */
                "X-Soha-Upstream-ID"?: components["parameters"]["SohaUpstreamIDHeader"];
                /** @description When true, returns redacted Soha route trace response headers for relay managers or tokens with allowRouteTrace explicitly enabled. */
                "X-Soha-Route-Trace"?: components["parameters"]["SohaRouteTraceHeader"];
                /** @description Optional relay cache behavior hint. */
                "X-Soha-Cache-Mode"?: components["parameters"]["SohaCacheModeHeader"];
            };
            path?: never;
            cookie?: never;
        };
        requestBody: {
            content: {
                "application/json": components["schemas"]["AnthropicMessagesRequest"];
            };
        };
        responses: {
            /** @description Native Anthropic-compatible message JSON or SSE stream. */
            200: {
                headers: {
                    "X-Soha-Route-ID": components["headers"]["SohaRouteIDHeader"];
                    "X-Soha-Upstream-ID": components["headers"]["SohaSelectedUpstreamIDHeader"];
                    "X-Soha-Provider-Kind": components["headers"]["SohaProviderKindHeader"];
                    "X-Soha-Public-Model": components["headers"]["SohaPublicModelHeader"];
                    "X-Soha-Upstream-Model": components["headers"]["SohaUpstreamModelHeader"];
                    "X-Soha-Relay-Endpoint": components["headers"]["SohaRelayEndpointHeader"];
                    "X-Soha-Relay-Stream": components["headers"]["SohaRelayStreamHeader"];
                    "X-Soha-Upstream-Status": components["headers"]["SohaUpstreamStatusHeader"];
                    "X-Soha-Cache-Status": components["headers"]["SohaCacheStatusHeader"];
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["NativeProviderObject"];
                    "text/event-stream": string;
                };
            };
            401: components["responses"]["Error"];
            403: components["responses"]["Error"];
        };
    };
    listMCPCapabilities: {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        requestBody?: never;
        responses: {
            /** @description MCP adapter capability list. */
            200: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["MCPCapabilityListEnvelope"];
                };
            };
        };
    };
    listMarketplacePlugins: {
        parameters: {
            query?: {
                q?: string;
                type?: string;
                publisher?: string;
            };
            header?: never;
            path?: never;
            cookie?: never;
        };
        requestBody?: never;
        responses: {
            /** @description Marketplace plugins visible to the current principal. */
            200: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["MarketplacePluginListEnvelope"];
                };
            };
        };
    };
    getMarketplacePlugin: {
        parameters: {
            query?: never;
            header?: never;
            path: {
                pluginID: components["parameters"]["PluginID"];
            };
            cookie?: never;
        };
        requestBody?: never;
        responses: {
            /** @description Marketplace plugin detail. */
            200: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["MarketplacePluginEnvelope"];
                };
            };
        };
    };
    listInstalledPlugins: {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        requestBody?: never;
        responses: {
            /** @description Installed plugin records. */
            200: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["InstalledPluginListEnvelope"];
                };
            };
        };
    };
    installPlugin: {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        requestBody: {
            content: {
                "application/json": components["schemas"]["PluginInstallRequest"];
            };
        };
        responses: {
            /** @description Installed plugin snapshot. */
            201: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["InstalledPluginEnvelope"];
                };
            };
        };
    };
    getInstalledPlugin: {
        parameters: {
            query?: never;
            header?: never;
            path: {
                pluginID: components["parameters"]["PluginID"];
            };
            cookie?: never;
        };
        requestBody?: never;
        responses: {
            /** @description Installed plugin record. */
            200: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["InstalledPluginEnvelope"];
                };
            };
        };
    };
    removeInstalledPlugin: {
        parameters: {
            query?: never;
            header?: never;
            path: {
                pluginID: components["parameters"]["PluginID"];
            };
            cookie?: never;
        };
        requestBody?: never;
        responses: {
            /** @description Installed plugin was removed. */
            200: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["OperationStatus"];
                };
            };
        };
    };
    getInstalledPluginManifest: {
        parameters: {
            query?: never;
            header?: never;
            path: {
                pluginID: components["parameters"]["PluginID"];
            };
            cookie?: never;
        };
        requestBody?: never;
        responses: {
            /** @description Installed plugin manifest snapshot. */
            200: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["PluginManifestEnvelope"];
                };
            };
        };
    };
    enableInstalledPlugin: {
        parameters: {
            query?: never;
            header?: never;
            path: {
                pluginID: components["parameters"]["PluginID"];
            };
            cookie?: never;
        };
        requestBody?: never;
        responses: {
            /** @description Plugin enabled. */
            200: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["InstalledPluginEnvelope"];
                };
            };
        };
    };
    disableInstalledPlugin: {
        parameters: {
            query?: never;
            header?: never;
            path: {
                pluginID: components["parameters"]["PluginID"];
            };
            cookie?: never;
        };
        requestBody?: never;
        responses: {
            /** @description Plugin disabled. */
            200: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["InstalledPluginEnvelope"];
                };
            };
        };
    };
    upgradeInstalledPlugin: {
        parameters: {
            query?: never;
            header?: never;
            path: {
                pluginID: components["parameters"]["PluginID"];
            };
            cookie?: never;
        };
        requestBody?: {
            content: {
                "application/json": components["schemas"]["PluginInstallRequest"];
            };
        };
        responses: {
            /** @description Plugin upgraded and manifest snapshot refreshed. */
            200: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["InstalledPluginEnvelope"];
                };
            };
        };
    };
    configureInstalledPlugin: {
        parameters: {
            query?: never;
            header?: never;
            path: {
                pluginID: components["parameters"]["PluginID"];
            };
            cookie?: never;
        };
        requestBody: {
            content: {
                "application/json": components["schemas"]["PluginConfigRequest"];
            };
        };
        responses: {
            /** @description Plugin configuration updated. */
            200: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["InstalledPluginEnvelope"];
                };
            };
        };
    };
}
