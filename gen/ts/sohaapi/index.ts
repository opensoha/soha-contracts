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
    };
    requestBodies: never;
    headers: never;
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
