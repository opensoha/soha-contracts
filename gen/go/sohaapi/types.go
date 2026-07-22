// Package sohaapi provides primitives to interact with the openapi HTTP API.
//
// Code generated from OpenSoha contracts by github.com/oapi-codegen/oapi-codegen/v2 version v2.7.1 DO NOT EDIT.
package sohaapi

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/oapi-codegen/runtime"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

const (
	BearerAuthScopes        bearerAuthContextKey        = "bearerAuth.Scopes"
	RuntimeBearerAuthScopes runtimeBearerAuthContextKey = "runtimeBearerAuth.Scopes"
	SohaXAPIKeyScopes       sohaXAPIKeyContextKey       = "sohaXAPIKey.Scopes"
)

// Defines values for AIEnvironmentTemplateBackend.
const (
	AIEnvironmentTemplateBackendContainer  AIEnvironmentTemplateBackend = "container"
	AIEnvironmentTemplateBackendKubernetes AIEnvironmentTemplateBackend = "kubernetes"
)

// Valid indicates whether the value is a known member of the AIEnvironmentTemplateBackend enum.
func (e AIEnvironmentTemplateBackend) Valid() bool {
	switch e {
	case AIEnvironmentTemplateBackendContainer:
		return true
	case AIEnvironmentTemplateBackendKubernetes:
		return true
	default:
		return false
	}
}

// Defines values for AIEnvironmentTemplateIsolationMode.
const (
	AIEnvironmentTemplateIsolationModeDisposableWrite AIEnvironmentTemplateIsolationMode = "disposable-write"
	AIEnvironmentTemplateIsolationModeReadOnly        AIEnvironmentTemplateIsolationMode = "read-only"
)

// Valid indicates whether the value is a known member of the AIEnvironmentTemplateIsolationMode enum.
func (e AIEnvironmentTemplateIsolationMode) Valid() bool {
	switch e {
	case AIEnvironmentTemplateIsolationModeDisposableWrite:
		return true
	case AIEnvironmentTemplateIsolationModeReadOnly:
		return true
	default:
		return false
	}
}

// Defines values for AIEnvironmentTemplateInputBackend.
const (
	AIEnvironmentTemplateInputBackendContainer  AIEnvironmentTemplateInputBackend = "container"
	AIEnvironmentTemplateInputBackendKubernetes AIEnvironmentTemplateInputBackend = "kubernetes"
)

// Valid indicates whether the value is a known member of the AIEnvironmentTemplateInputBackend enum.
func (e AIEnvironmentTemplateInputBackend) Valid() bool {
	switch e {
	case AIEnvironmentTemplateInputBackendContainer:
		return true
	case AIEnvironmentTemplateInputBackendKubernetes:
		return true
	default:
		return false
	}
}

// Defines values for AIEnvironmentTemplateInputIsolationMode.
const (
	AIEnvironmentTemplateInputIsolationModeDisposableWrite AIEnvironmentTemplateInputIsolationMode = "disposable-write"
	AIEnvironmentTemplateInputIsolationModeReadOnly        AIEnvironmentTemplateInputIsolationMode = "read-only"
)

// Valid indicates whether the value is a known member of the AIEnvironmentTemplateInputIsolationMode enum.
func (e AIEnvironmentTemplateInputIsolationMode) Valid() bool {
	switch e {
	case AIEnvironmentTemplateInputIsolationModeDisposableWrite:
		return true
	case AIEnvironmentTemplateInputIsolationModeReadOnly:
		return true
	default:
		return false
	}
}

// Defines values for AIMemoryPolicyInputConsentMode.
const (
	AIMemoryPolicyInputConsentModeDisabled AIMemoryPolicyInputConsentMode = "disabled"
	AIMemoryPolicyInputConsentModeExplicit AIMemoryPolicyInputConsentMode = "explicit"
)

// Valid indicates whether the value is a known member of the AIMemoryPolicyInputConsentMode enum.
func (e AIMemoryPolicyInputConsentMode) Valid() bool {
	switch e {
	case AIMemoryPolicyInputConsentModeDisabled:
		return true
	case AIMemoryPolicyInputConsentModeExplicit:
		return true
	default:
		return false
	}
}

// Defines values for AIMemoryRecordSourceType.
const (
	CuratedExtractor AIMemoryRecordSourceType = "curated_extractor"
	ExplicitUser     AIMemoryRecordSourceType = "explicit_user"
)

// Valid indicates whether the value is a known member of the AIMemoryRecordSourceType enum.
func (e AIMemoryRecordSourceType) Valid() bool {
	switch e {
	case CuratedExtractor:
		return true
	case ExplicitUser:
		return true
	default:
		return false
	}
}

// Defines values for AIProductionOperationKind.
const (
	AIProductionOperationKindBackup       AIProductionOperationKind = "backup"
	AIProductionOperationKindDrill        AIProductionOperationKind = "drill"
	AIProductionOperationKindGc           AIProductionOperationKind = "gc"
	AIProductionOperationKindIndexRebuild AIProductionOperationKind = "index_rebuild"
	AIProductionOperationKindRestore      AIProductionOperationKind = "restore"
)

// Valid indicates whether the value is a known member of the AIProductionOperationKind enum.
func (e AIProductionOperationKind) Valid() bool {
	switch e {
	case AIProductionOperationKindBackup:
		return true
	case AIProductionOperationKindDrill:
		return true
	case AIProductionOperationKindGc:
		return true
	case AIProductionOperationKindIndexRebuild:
		return true
	case AIProductionOperationKindRestore:
		return true
	default:
		return false
	}
}

// Defines values for AIProductionOperationInputKind.
const (
	AIProductionOperationInputKindBackup       AIProductionOperationInputKind = "backup"
	AIProductionOperationInputKindDrill        AIProductionOperationInputKind = "drill"
	AIProductionOperationInputKindIndexRebuild AIProductionOperationInputKind = "index_rebuild"
	AIProductionOperationInputKindRestore      AIProductionOperationInputKind = "restore"
)

// Valid indicates whether the value is a known member of the AIProductionOperationInputKind enum.
func (e AIProductionOperationInputKind) Valid() bool {
	switch e {
	case AIProductionOperationInputKindBackup:
		return true
	case AIProductionOperationInputKindDrill:
		return true
	case AIProductionOperationInputKindIndexRebuild:
		return true
	case AIProductionOperationInputKindRestore:
		return true
	default:
		return false
	}
}

// Defines values for AgentProviderRunnerStatusHealth.
const (
	Healthy   AgentProviderRunnerStatusHealth = "healthy"
	Unhealthy AgentProviderRunnerStatusHealth = "unhealthy"
	Unknown   AgentProviderRunnerStatusHealth = "unknown"
)

// Valid indicates whether the value is a known member of the AgentProviderRunnerStatusHealth enum.
func (e AgentProviderRunnerStatusHealth) Valid() bool {
	switch e {
	case Healthy:
		return true
	case Unhealthy:
		return true
	case Unknown:
		return true
	default:
		return false
	}
}

// Defines values for AgentProviderRuntimeDefinitionKind.
const (
	Cli       AgentProviderRuntimeDefinitionKind = "cli"
	Container AgentProviderRuntimeDefinitionKind = "container"
	Remote    AgentProviderRuntimeDefinitionKind = "remote"
)

// Valid indicates whether the value is a known member of the AgentProviderRuntimeDefinitionKind enum.
func (e AgentProviderRuntimeDefinitionKind) Valid() bool {
	switch e {
	case Cli:
		return true
	case Container:
		return true
	case Remote:
		return true
	default:
		return false
	}
}

// Defines values for AgentRunCallbackWorkbenchAgentStatusEventProviderKind.
const (
	AgentRunCallbackWorkbenchAgentStatusEventProviderKindAnthropic        AgentRunCallbackWorkbenchAgentStatusEventProviderKind = "anthropic"
	AgentRunCallbackWorkbenchAgentStatusEventProviderKindAzureOpenai      AgentRunCallbackWorkbenchAgentStatusEventProviderKind = "azure-openai"
	AgentRunCallbackWorkbenchAgentStatusEventProviderKindCohere           AgentRunCallbackWorkbenchAgentStatusEventProviderKind = "cohere"
	AgentRunCallbackWorkbenchAgentStatusEventProviderKindDeepseek         AgentRunCallbackWorkbenchAgentStatusEventProviderKind = "deepseek"
	AgentRunCallbackWorkbenchAgentStatusEventProviderKindGemini           AgentRunCallbackWorkbenchAgentStatusEventProviderKind = "gemini"
	AgentRunCallbackWorkbenchAgentStatusEventProviderKindGeneral          AgentRunCallbackWorkbenchAgentStatusEventProviderKind = "general"
	AgentRunCallbackWorkbenchAgentStatusEventProviderKindHermes           AgentRunCallbackWorkbenchAgentStatusEventProviderKind = "hermes"
	AgentRunCallbackWorkbenchAgentStatusEventProviderKindInternal         AgentRunCallbackWorkbenchAgentStatusEventProviderKind = "internal"
	AgentRunCallbackWorkbenchAgentStatusEventProviderKindOpenai           AgentRunCallbackWorkbenchAgentStatusEventProviderKind = "openai"
	AgentRunCallbackWorkbenchAgentStatusEventProviderKindOpenaiCompatible AgentRunCallbackWorkbenchAgentStatusEventProviderKind = "openai-compatible"
	AgentRunCallbackWorkbenchAgentStatusEventProviderKindOpenclaw         AgentRunCallbackWorkbenchAgentStatusEventProviderKind = "openclaw"
	AgentRunCallbackWorkbenchAgentStatusEventProviderKindOpenrouter       AgentRunCallbackWorkbenchAgentStatusEventProviderKind = "openrouter"
	AgentRunCallbackWorkbenchAgentStatusEventProviderKindQwen             AgentRunCallbackWorkbenchAgentStatusEventProviderKind = "qwen"
)

// Valid indicates whether the value is a known member of the AgentRunCallbackWorkbenchAgentStatusEventProviderKind enum.
func (e AgentRunCallbackWorkbenchAgentStatusEventProviderKind) Valid() bool {
	switch e {
	case AgentRunCallbackWorkbenchAgentStatusEventProviderKindAnthropic:
		return true
	case AgentRunCallbackWorkbenchAgentStatusEventProviderKindAzureOpenai:
		return true
	case AgentRunCallbackWorkbenchAgentStatusEventProviderKindCohere:
		return true
	case AgentRunCallbackWorkbenchAgentStatusEventProviderKindDeepseek:
		return true
	case AgentRunCallbackWorkbenchAgentStatusEventProviderKindGemini:
		return true
	case AgentRunCallbackWorkbenchAgentStatusEventProviderKindGeneral:
		return true
	case AgentRunCallbackWorkbenchAgentStatusEventProviderKindHermes:
		return true
	case AgentRunCallbackWorkbenchAgentStatusEventProviderKindInternal:
		return true
	case AgentRunCallbackWorkbenchAgentStatusEventProviderKindOpenai:
		return true
	case AgentRunCallbackWorkbenchAgentStatusEventProviderKindOpenaiCompatible:
		return true
	case AgentRunCallbackWorkbenchAgentStatusEventProviderKindOpenclaw:
		return true
	case AgentRunCallbackWorkbenchAgentStatusEventProviderKindOpenrouter:
		return true
	case AgentRunCallbackWorkbenchAgentStatusEventProviderKindQwen:
		return true
	default:
		return false
	}
}

// Defines values for AgentRunCallbackWorkbenchAgentStatusEventStatus.
const (
	AgentRunCallbackWorkbenchAgentStatusEventStatusCancelled AgentRunCallbackWorkbenchAgentStatusEventStatus = "cancelled"
	AgentRunCallbackWorkbenchAgentStatusEventStatusFailed    AgentRunCallbackWorkbenchAgentStatusEventStatus = "failed"
	AgentRunCallbackWorkbenchAgentStatusEventStatusQueued    AgentRunCallbackWorkbenchAgentStatusEventStatus = "queued"
	AgentRunCallbackWorkbenchAgentStatusEventStatusRunning   AgentRunCallbackWorkbenchAgentStatusEventStatus = "running"
	AgentRunCallbackWorkbenchAgentStatusEventStatusSucceeded AgentRunCallbackWorkbenchAgentStatusEventStatus = "succeeded"
)

// Valid indicates whether the value is a known member of the AgentRunCallbackWorkbenchAgentStatusEventStatus enum.
func (e AgentRunCallbackWorkbenchAgentStatusEventStatus) Valid() bool {
	switch e {
	case AgentRunCallbackWorkbenchAgentStatusEventStatusCancelled:
		return true
	case AgentRunCallbackWorkbenchAgentStatusEventStatusFailed:
		return true
	case AgentRunCallbackWorkbenchAgentStatusEventStatusQueued:
		return true
	case AgentRunCallbackWorkbenchAgentStatusEventStatusRunning:
		return true
	case AgentRunCallbackWorkbenchAgentStatusEventStatusSucceeded:
		return true
	default:
		return false
	}
}

// Defines values for AgentRunCallbackWorkbenchAgentStatusEventType.
const (
	AgentRunCallbackWorkbenchAgentStatusEventTypeAgentStatus AgentRunCallbackWorkbenchAgentStatusEventType = "agent.status"
)

// Valid indicates whether the value is a known member of the AgentRunCallbackWorkbenchAgentStatusEventType enum.
func (e AgentRunCallbackWorkbenchAgentStatusEventType) Valid() bool {
	switch e {
	case AgentRunCallbackWorkbenchAgentStatusEventTypeAgentStatus:
		return true
	default:
		return false
	}
}

// Defines values for AgentRunCallbackWorkbenchArtifactUpdatedEventType.
const (
	AgentRunCallbackWorkbenchArtifactUpdatedEventTypeArtifactUpdated AgentRunCallbackWorkbenchArtifactUpdatedEventType = "artifact.updated"
)

// Valid indicates whether the value is a known member of the AgentRunCallbackWorkbenchArtifactUpdatedEventType enum.
func (e AgentRunCallbackWorkbenchArtifactUpdatedEventType) Valid() bool {
	switch e {
	case AgentRunCallbackWorkbenchArtifactUpdatedEventTypeArtifactUpdated:
		return true
	default:
		return false
	}
}

// Defines values for AgentRunCallbackWorkbenchCardCommandEventType.
const (
	AgentRunCallbackWorkbenchCardCommandEventTypeCardCommand AgentRunCallbackWorkbenchCardCommandEventType = "card.command"
)

// Valid indicates whether the value is a known member of the AgentRunCallbackWorkbenchCardCommandEventType enum.
func (e AgentRunCallbackWorkbenchCardCommandEventType) Valid() bool {
	switch e {
	case AgentRunCallbackWorkbenchCardCommandEventTypeCardCommand:
		return true
	default:
		return false
	}
}

// Defines values for AgentRunCallbackWorkbenchErrorEventType.
const (
	AgentRunCallbackWorkbenchErrorEventTypeError AgentRunCallbackWorkbenchErrorEventType = "error"
)

// Valid indicates whether the value is a known member of the AgentRunCallbackWorkbenchErrorEventType enum.
func (e AgentRunCallbackWorkbenchErrorEventType) Valid() bool {
	switch e {
	case AgentRunCallbackWorkbenchErrorEventTypeError:
		return true
	default:
		return false
	}
}

// Defines values for AgentRunCallbackWorkbenchMessageDeltaEventRole.
const (
	AgentRunCallbackWorkbenchMessageDeltaEventRoleAssistant AgentRunCallbackWorkbenchMessageDeltaEventRole = "assistant"
)

// Valid indicates whether the value is a known member of the AgentRunCallbackWorkbenchMessageDeltaEventRole enum.
func (e AgentRunCallbackWorkbenchMessageDeltaEventRole) Valid() bool {
	switch e {
	case AgentRunCallbackWorkbenchMessageDeltaEventRoleAssistant:
		return true
	default:
		return false
	}
}

// Defines values for AgentRunCallbackWorkbenchMessageDeltaEventType.
const (
	AgentRunCallbackWorkbenchMessageDeltaEventTypeMessageDelta AgentRunCallbackWorkbenchMessageDeltaEventType = "message.delta"
)

// Valid indicates whether the value is a known member of the AgentRunCallbackWorkbenchMessageDeltaEventType enum.
func (e AgentRunCallbackWorkbenchMessageDeltaEventType) Valid() bool {
	switch e {
	case AgentRunCallbackWorkbenchMessageDeltaEventTypeMessageDelta:
		return true
	default:
		return false
	}
}

// Defines values for AgentRunCallbackWorkbenchMessageDoneEventRole.
const (
	AgentRunCallbackWorkbenchMessageDoneEventRoleAssistant AgentRunCallbackWorkbenchMessageDoneEventRole = "assistant"
)

// Valid indicates whether the value is a known member of the AgentRunCallbackWorkbenchMessageDoneEventRole enum.
func (e AgentRunCallbackWorkbenchMessageDoneEventRole) Valid() bool {
	switch e {
	case AgentRunCallbackWorkbenchMessageDoneEventRoleAssistant:
		return true
	default:
		return false
	}
}

// Defines values for AgentRunCallbackWorkbenchMessageDoneEventType.
const (
	AgentRunCallbackWorkbenchMessageDoneEventTypeMessageDone AgentRunCallbackWorkbenchMessageDoneEventType = "message.done"
)

// Valid indicates whether the value is a known member of the AgentRunCallbackWorkbenchMessageDoneEventType enum.
func (e AgentRunCallbackWorkbenchMessageDoneEventType) Valid() bool {
	switch e {
	case AgentRunCallbackWorkbenchMessageDoneEventTypeMessageDone:
		return true
	default:
		return false
	}
}

// Defines values for AgentRunCallbackWorkbenchSourceUpdatedEventType.
const (
	AgentRunCallbackWorkbenchSourceUpdatedEventTypeSourceUpdated AgentRunCallbackWorkbenchSourceUpdatedEventType = "source.updated"
)

// Valid indicates whether the value is a known member of the AgentRunCallbackWorkbenchSourceUpdatedEventType enum.
func (e AgentRunCallbackWorkbenchSourceUpdatedEventType) Valid() bool {
	switch e {
	case AgentRunCallbackWorkbenchSourceUpdatedEventTypeSourceUpdated:
		return true
	default:
		return false
	}
}

// Defines values for AgentRunCallbackWorkbenchThinkingDeltaEventType.
const (
	AgentRunCallbackWorkbenchThinkingDeltaEventTypeThinkingDelta AgentRunCallbackWorkbenchThinkingDeltaEventType = "thinking.delta"
)

// Valid indicates whether the value is a known member of the AgentRunCallbackWorkbenchThinkingDeltaEventType enum.
func (e AgentRunCallbackWorkbenchThinkingDeltaEventType) Valid() bool {
	switch e {
	case AgentRunCallbackWorkbenchThinkingDeltaEventTypeThinkingDelta:
		return true
	default:
		return false
	}
}

// Defines values for AgentRunCallbackWorkbenchThinkingDoneEventType.
const (
	AgentRunCallbackWorkbenchThinkingDoneEventTypeThinkingDone AgentRunCallbackWorkbenchThinkingDoneEventType = "thinking.done"
)

// Valid indicates whether the value is a known member of the AgentRunCallbackWorkbenchThinkingDoneEventType enum.
func (e AgentRunCallbackWorkbenchThinkingDoneEventType) Valid() bool {
	switch e {
	case AgentRunCallbackWorkbenchThinkingDoneEventTypeThinkingDone:
		return true
	default:
		return false
	}
}

// Defines values for AgentRunCallbackWorkbenchToolCompletedEventType.
const (
	AgentRunCallbackWorkbenchToolCompletedEventTypeToolCompleted AgentRunCallbackWorkbenchToolCompletedEventType = "tool.completed"
)

// Valid indicates whether the value is a known member of the AgentRunCallbackWorkbenchToolCompletedEventType enum.
func (e AgentRunCallbackWorkbenchToolCompletedEventType) Valid() bool {
	switch e {
	case AgentRunCallbackWorkbenchToolCompletedEventTypeToolCompleted:
		return true
	default:
		return false
	}
}

// Defines values for AgentRunCallbackWorkbenchToolDeltaEventType.
const (
	AgentRunCallbackWorkbenchToolDeltaEventTypeToolDelta AgentRunCallbackWorkbenchToolDeltaEventType = "tool.delta"
)

// Valid indicates whether the value is a known member of the AgentRunCallbackWorkbenchToolDeltaEventType enum.
func (e AgentRunCallbackWorkbenchToolDeltaEventType) Valid() bool {
	switch e {
	case AgentRunCallbackWorkbenchToolDeltaEventTypeToolDelta:
		return true
	default:
		return false
	}
}

// Defines values for AgentRunCallbackWorkbenchToolStartedEventType.
const (
	AgentRunCallbackWorkbenchToolStartedEventTypeToolStarted AgentRunCallbackWorkbenchToolStartedEventType = "tool.started"
)

// Valid indicates whether the value is a known member of the AgentRunCallbackWorkbenchToolStartedEventType enum.
func (e AgentRunCallbackWorkbenchToolStartedEventType) Valid() bool {
	switch e {
	case AgentRunCallbackWorkbenchToolStartedEventTypeToolStarted:
		return true
	default:
		return false
	}
}

// Defines values for ApplicationDeliveryActionKind.
const (
	ApplicationDeliveryActionKindBuild       ApplicationDeliveryActionKind = "build"
	ApplicationDeliveryActionKindBuildDeploy ApplicationDeliveryActionKind = "build_deploy"
	ApplicationDeliveryActionKindDeploy      ApplicationDeliveryActionKind = "deploy"
	ApplicationDeliveryActionKindRollback    ApplicationDeliveryActionKind = "rollback"
	ApplicationDeliveryActionKindVerify      ApplicationDeliveryActionKind = "verify"
	ApplicationDeliveryActionKindWorkflow    ApplicationDeliveryActionKind = "workflow"
)

// Valid indicates whether the value is a known member of the ApplicationDeliveryActionKind enum.
func (e ApplicationDeliveryActionKind) Valid() bool {
	switch e {
	case ApplicationDeliveryActionKindBuild:
		return true
	case ApplicationDeliveryActionKindBuildDeploy:
		return true
	case ApplicationDeliveryActionKindDeploy:
		return true
	case ApplicationDeliveryActionKindRollback:
		return true
	case ApplicationDeliveryActionKindVerify:
		return true
	case ApplicationDeliveryActionKindWorkflow:
		return true
	default:
		return false
	}
}

// Defines values for ApplicationServiceServiceKind.
const (
	ApplicationServiceServiceKindExternalService    ApplicationServiceServiceKind = "external_service"
	ApplicationServiceServiceKindHelmRelease        ApplicationServiceServiceKind = "helm_release"
	ApplicationServiceServiceKindJob                ApplicationServiceServiceKind = "job"
	ApplicationServiceServiceKindKubernetesWorkload ApplicationServiceServiceKind = "kubernetes_workload"
)

// Valid indicates whether the value is a known member of the ApplicationServiceServiceKind enum.
func (e ApplicationServiceServiceKind) Valid() bool {
	switch e {
	case ApplicationServiceServiceKindExternalService:
		return true
	case ApplicationServiceServiceKindHelmRelease:
		return true
	case ApplicationServiceServiceKindJob:
		return true
	case ApplicationServiceServiceKindKubernetesWorkload:
		return true
	default:
		return false
	}
}

// Defines values for ApplicationServiceInputServiceKind.
const (
	ApplicationServiceInputServiceKindExternalService    ApplicationServiceInputServiceKind = "external_service"
	ApplicationServiceInputServiceKindHelmRelease        ApplicationServiceInputServiceKind = "helm_release"
	ApplicationServiceInputServiceKindJob                ApplicationServiceInputServiceKind = "job"
	ApplicationServiceInputServiceKindKubernetesWorkload ApplicationServiceInputServiceKind = "kubernetes_workload"
)

// Valid indicates whether the value is a known member of the ApplicationServiceInputServiceKind enum.
func (e ApplicationServiceInputServiceKind) Valid() bool {
	switch e {
	case ApplicationServiceInputServiceKindExternalService:
		return true
	case ApplicationServiceInputServiceKindHelmRelease:
		return true
	case ApplicationServiceInputServiceKindJob:
		return true
	case ApplicationServiceInputServiceKindKubernetesWorkload:
		return true
	default:
		return false
	}
}

// Defines values for BuildSourceType.
const (
	BuildSourceTypeExternalPipeline      BuildSourceType = "external_pipeline"
	BuildSourceTypePlatformBuildTemplate BuildSourceType = "platform_build_template"
	BuildSourceTypeRepoDockerfile        BuildSourceType = "repo_dockerfile"
)

// Valid indicates whether the value is a known member of the BuildSourceType enum.
func (e BuildSourceType) Valid() bool {
	switch e {
	case BuildSourceTypeExternalPipeline:
		return true
	case BuildSourceTypePlatformBuildTemplate:
		return true
	case BuildSourceTypeRepoDockerfile:
		return true
	default:
		return false
	}
}

// Defines values for BuildSourceConfigBuilderKind.
const (
	BuildSourceConfigBuilderKindBuildx BuildSourceConfigBuilderKind = "buildx"
	BuildSourceConfigBuilderKindDocker BuildSourceConfigBuilderKind = "docker"
	BuildSourceConfigBuilderKindKaniko BuildSourceConfigBuilderKind = "kaniko"
)

// Valid indicates whether the value is a known member of the BuildSourceConfigBuilderKind enum.
func (e BuildSourceConfigBuilderKind) Valid() bool {
	switch e {
	case BuildSourceConfigBuilderKindBuildx:
		return true
	case BuildSourceConfigBuilderKindDocker:
		return true
	case BuildSourceConfigBuilderKindKaniko:
		return true
	default:
		return false
	}
}

// Defines values for BuildSourceInputType.
const (
	BuildSourceInputTypeExternalPipeline      BuildSourceInputType = "external_pipeline"
	BuildSourceInputTypePlatformBuildTemplate BuildSourceInputType = "platform_build_template"
	BuildSourceInputTypeRepoDockerfile        BuildSourceInputType = "repo_dockerfile"
)

// Valid indicates whether the value is a known member of the BuildSourceInputType enum.
func (e BuildSourceInputType) Valid() bool {
	switch e {
	case BuildSourceInputTypeExternalPipeline:
		return true
	case BuildSourceInputTypePlatformBuildTemplate:
		return true
	case BuildSourceInputTypeRepoDockerfile:
		return true
	default:
		return false
	}
}

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

// Defines values for ComputeAccessMode.
const (
	ComputeAccessModeAgentProxy ComputeAccessMode = "agent_proxy"
	ComputeAccessModeDirect     ComputeAccessMode = "direct"
)

// Valid indicates whether the value is a known member of the ComputeAccessMode enum.
func (e ComputeAccessMode) Valid() bool {
	switch e {
	case ComputeAccessModeAgentProxy:
		return true
	case ComputeAccessModeDirect:
		return true
	default:
		return false
	}
}

// Defines values for ComputeAccessSourceType.
const (
	ComputeAccessSourceTypeAgentHost                ComputeAccessSourceType = "agent_host"
	ComputeAccessSourceTypeRuntimeHost              ComputeAccessSourceType = "runtime_host"
	ComputeAccessSourceTypeVirtualizationConnection ComputeAccessSourceType = "virtualization_connection"
)

// Valid indicates whether the value is a known member of the ComputeAccessSourceType enum.
func (e ComputeAccessSourceType) Valid() bool {
	switch e {
	case ComputeAccessSourceTypeAgentHost:
		return true
	case ComputeAccessSourceTypeRuntimeHost:
		return true
	case ComputeAccessSourceTypeVirtualizationConnection:
		return true
	default:
		return false
	}
}

// Defines values for ComputeAttentionSeverity.
const (
	ComputeAttentionSeverityCritical ComputeAttentionSeverity = "critical"
	ComputeAttentionSeverityInfo     ComputeAttentionSeverity = "info"
	ComputeAttentionSeverityWarning  ComputeAttentionSeverity = "warning"
)

// Valid indicates whether the value is a known member of the ComputeAttentionSeverity enum.
func (e ComputeAttentionSeverity) Valid() bool {
	switch e {
	case ComputeAttentionSeverityCritical:
		return true
	case ComputeAttentionSeverityInfo:
		return true
	case ComputeAttentionSeverityWarning:
		return true
	default:
		return false
	}
}

// Defines values for ComputeDomain.
const (
	ComputeDomainAgent            ComputeDomain = "agent"
	ComputeDomainContainerRuntime ComputeDomain = "container_runtime"
	ComputeDomainVirtualization   ComputeDomain = "virtualization"
)

// Valid indicates whether the value is a known member of the ComputeDomain enum.
func (e ComputeDomain) Valid() bool {
	switch e {
	case ComputeDomainAgent:
		return true
	case ComputeDomainContainerRuntime:
		return true
	case ComputeDomainVirtualization:
		return true
	default:
		return false
	}
}

// Defines values for ComputeFeatureID.
const (
	ComputeFeaturePluginProviders ComputeFeatureID = "plugin_providers"
	ComputeFeatureProviderRead    ComputeFeatureID = "provider_read"
	ComputeFeatureProviderWrite   ComputeFeatureID = "provider_write"
	ComputeFeatureRelations       ComputeFeatureID = "relations"
	ComputeFeatureTaskActions     ComputeFeatureID = "task_actions"
	ComputeFeatureWorkbench       ComputeFeatureID = "workbench"
)

// Valid indicates whether the value is a known member of the ComputeFeatureID enum.
func (e ComputeFeatureID) Valid() bool {
	switch e {
	case ComputeFeaturePluginProviders:
		return true
	case ComputeFeatureProviderRead:
		return true
	case ComputeFeatureProviderWrite:
		return true
	case ComputeFeatureRelations:
		return true
	case ComputeFeatureTaskActions:
		return true
	case ComputeFeatureWorkbench:
		return true
	default:
		return false
	}
}

// Defines values for ComputeHealthStatus.
const (
	ComputeHealthStatusDegraded    ComputeHealthStatus = "degraded"
	ComputeHealthStatusHealthy     ComputeHealthStatus = "healthy"
	ComputeHealthStatusPending     ComputeHealthStatus = "pending"
	ComputeHealthStatusUnavailable ComputeHealthStatus = "unavailable"
	ComputeHealthStatusUnknown     ComputeHealthStatus = "unknown"
)

// Valid indicates whether the value is a known member of the ComputeHealthStatus enum.
func (e ComputeHealthStatus) Valid() bool {
	switch e {
	case ComputeHealthStatusDegraded:
		return true
	case ComputeHealthStatusHealthy:
		return true
	case ComputeHealthStatusPending:
		return true
	case ComputeHealthStatusUnavailable:
		return true
	case ComputeHealthStatusUnknown:
		return true
	default:
		return false
	}
}

// Defines values for ComputePluginRuntimeMode.
const (
	ComputePluginRuntimeModeBuiltin      ComputePluginRuntimeMode = "builtin"
	ComputePluginRuntimeModeExternal     ComputePluginRuntimeMode = "external"
	ComputePluginRuntimeModeManaged      ComputePluginRuntimeMode = "managed"
	ComputePluginRuntimeModeManifestOnly ComputePluginRuntimeMode = "manifest_only"
)

// Valid indicates whether the value is a known member of the ComputePluginRuntimeMode enum.
func (e ComputePluginRuntimeMode) Valid() bool {
	switch e {
	case ComputePluginRuntimeModeBuiltin:
		return true
	case ComputePluginRuntimeModeExternal:
		return true
	case ComputePluginRuntimeModeManaged:
		return true
	case ComputePluginRuntimeModeManifestOnly:
		return true
	default:
		return false
	}
}

// Defines values for ComputeProviderActivationLevel.
const (
	ComputeProviderActivationLevelDescriptor ComputeProviderActivationLevel = "descriptor"
	ComputeProviderActivationLevelRead       ComputeProviderActivationLevel = "read"
	ComputeProviderActivationLevelWrite      ComputeProviderActivationLevel = "write"
)

// Valid indicates whether the value is a known member of the ComputeProviderActivationLevel enum.
func (e ComputeProviderActivationLevel) Valid() bool {
	switch e {
	case ComputeProviderActivationLevelDescriptor:
		return true
	case ComputeProviderActivationLevelRead:
		return true
	case ComputeProviderActivationLevelWrite:
		return true
	default:
		return false
	}
}

// Defines values for ComputeProviderDomain.
const (
	ComputeProviderDomainContainerRuntime ComputeProviderDomain = "container_runtime"
	ComputeProviderDomainVirtualization   ComputeProviderDomain = "virtualization"
)

// Valid indicates whether the value is a known member of the ComputeProviderDomain enum.
func (e ComputeProviderDomain) Valid() bool {
	switch e {
	case ComputeProviderDomainContainerRuntime:
		return true
	case ComputeProviderDomainVirtualization:
		return true
	default:
		return false
	}
}

// Defines values for ComputeProviderSource.
const (
	ComputeProviderSourceBuiltin ComputeProviderSource = "builtin"
	ComputeProviderSourcePlugin  ComputeProviderSource = "plugin"
)

// Valid indicates whether the value is a known member of the ComputeProviderSource enum.
func (e ComputeProviderSource) Valid() bool {
	switch e {
	case ComputeProviderSourceBuiltin:
		return true
	case ComputeProviderSourcePlugin:
		return true
	default:
		return false
	}
}

// Defines values for ComputeRelationSource.
const (
	ComputeRelationSourceDerived   ComputeRelationSource = "derived"
	ComputeRelationSourcePersisted ComputeRelationSource = "persisted"
	ComputeRelationSourceProvider  ComputeRelationSource = "provider"
)

// Valid indicates whether the value is a known member of the ComputeRelationSource enum.
func (e ComputeRelationSource) Valid() bool {
	switch e {
	case ComputeRelationSourceDerived:
		return true
	case ComputeRelationSourcePersisted:
		return true
	case ComputeRelationSourceProvider:
		return true
	default:
		return false
	}
}

// Defines values for ComputeRelationType.
const (
	ComputeRelationTypeConnectedBy ComputeRelationType = "connected_by"
	ComputeRelationTypeContains    ComputeRelationType = "contains"
	ComputeRelationTypeDerivedFrom ComputeRelationType = "derived_from"
	ComputeRelationTypeExposes     ComputeRelationType = "exposes"
	ComputeRelationTypeManages     ComputeRelationType = "manages"
	ComputeRelationTypeProvisions  ComputeRelationType = "provisions"
	ComputeRelationTypeRunsOn      ComputeRelationType = "runs_on"
)

// Valid indicates whether the value is a known member of the ComputeRelationType enum.
func (e ComputeRelationType) Valid() bool {
	switch e {
	case ComputeRelationTypeConnectedBy:
		return true
	case ComputeRelationTypeContains:
		return true
	case ComputeRelationTypeDerivedFrom:
		return true
	case ComputeRelationTypeExposes:
		return true
	case ComputeRelationTypeManages:
		return true
	case ComputeRelationTypeProvisions:
		return true
	case ComputeRelationTypeRunsOn:
		return true
	default:
		return false
	}
}

// Defines values for ComputeResourceKind.
const (
	ComputeResourceKindAgentHost   ComputeResourceKind = "agent_host"
	ComputeResourceKindCluster     ComputeResourceKind = "cluster"
	ComputeResourceKindConnection  ComputeResourceKind = "connection"
	ComputeResourceKindContainer   ComputeResourceKind = "container"
	ComputeResourceKindFlavor      ComputeResourceKind = "flavor"
	ComputeResourceKindImage       ComputeResourceKind = "image"
	ComputeResourceKindPort        ComputeResourceKind = "port"
	ComputeResourceKindProject     ComputeResourceKind = "project"
	ComputeResourceKindRuntimeHost ComputeResourceKind = "runtime_host"
	ComputeResourceKindService     ComputeResourceKind = "service"
	ComputeResourceKindTemplate    ComputeResourceKind = "template"
	ComputeResourceKindVM          ComputeResourceKind = "vm"
)

// Valid indicates whether the value is a known member of the ComputeResourceKind enum.
func (e ComputeResourceKind) Valid() bool {
	switch e {
	case ComputeResourceKindAgentHost:
		return true
	case ComputeResourceKindCluster:
		return true
	case ComputeResourceKindConnection:
		return true
	case ComputeResourceKindContainer:
		return true
	case ComputeResourceKindFlavor:
		return true
	case ComputeResourceKindImage:
		return true
	case ComputeResourceKindPort:
		return true
	case ComputeResourceKindProject:
		return true
	case ComputeResourceKindRuntimeHost:
		return true
	case ComputeResourceKindService:
		return true
	case ComputeResourceKindTemplate:
		return true
	case ComputeResourceKindVM:
		return true
	default:
		return false
	}
}

// Defines values for ComputeRolloutStage.
const (
	ComputeRolloutStageAdminPreview ComputeRolloutStage = "admin_preview"
	ComputeRolloutStageDefault      ComputeRolloutStage = "default"
	ComputeRolloutStageDisabled     ComputeRolloutStage = "disabled"
	ComputeRolloutStageHidden       ComputeRolloutStage = "hidden"
)

// Valid indicates whether the value is a known member of the ComputeRolloutStage enum.
func (e ComputeRolloutStage) Valid() bool {
	switch e {
	case ComputeRolloutStageAdminPreview:
		return true
	case ComputeRolloutStageDefault:
		return true
	case ComputeRolloutStageDisabled:
		return true
	case ComputeRolloutStageHidden:
		return true
	default:
		return false
	}
}

// Defines values for ComputeSectionStatus.
const (
	ComputeSectionStatusDegraded    ComputeSectionStatus = "degraded"
	ComputeSectionStatusOK          ComputeSectionStatus = "ok"
	ComputeSectionStatusUnavailable ComputeSectionStatus = "unavailable"
)

// Valid indicates whether the value is a known member of the ComputeSectionStatus enum.
func (e ComputeSectionStatus) Valid() bool {
	switch e {
	case ComputeSectionStatusDegraded:
		return true
	case ComputeSectionStatusOK:
		return true
	case ComputeSectionStatusUnavailable:
		return true
	default:
		return false
	}
}

// Defines values for ComputeTaskAction.
const (
	ComputeTaskActionCancel ComputeTaskAction = "cancel"
	ComputeTaskActionLogs   ComputeTaskAction = "logs"
	ComputeTaskActionRetry  ComputeTaskAction = "retry"
)

// Valid indicates whether the value is a known member of the ComputeTaskAction enum.
func (e ComputeTaskAction) Valid() bool {
	switch e {
	case ComputeTaskActionCancel:
		return true
	case ComputeTaskActionLogs:
		return true
	case ComputeTaskActionRetry:
		return true
	default:
		return false
	}
}

// Defines values for ComputeTaskCategory.
const (
	ComputeTaskCategoryBuild     ComputeTaskCategory = "build"
	ComputeTaskCategoryLifecycle ComputeTaskCategory = "lifecycle"
	ComputeTaskCategoryOperation ComputeTaskCategory = "operation"
	ComputeTaskCategorySync      ComputeTaskCategory = "sync"
)

// Valid indicates whether the value is a known member of the ComputeTaskCategory enum.
func (e ComputeTaskCategory) Valid() bool {
	switch e {
	case ComputeTaskCategoryBuild:
		return true
	case ComputeTaskCategoryLifecycle:
		return true
	case ComputeTaskCategoryOperation:
		return true
	case ComputeTaskCategorySync:
		return true
	default:
		return false
	}
}

// Defines values for ComputeTaskDomain.
const (
	ComputeTaskDomainContainerRuntime ComputeTaskDomain = "container_runtime"
	ComputeTaskDomainVirtualization   ComputeTaskDomain = "virtualization"
)

// Valid indicates whether the value is a known member of the ComputeTaskDomain enum.
func (e ComputeTaskDomain) Valid() bool {
	switch e {
	case ComputeTaskDomainContainerRuntime:
		return true
	case ComputeTaskDomainVirtualization:
		return true
	default:
		return false
	}
}

// Defines values for ComputeTaskStatus.
const (
	ComputeTaskStatusCanceled  ComputeTaskStatus = "canceled"
	ComputeTaskStatusFailed    ComputeTaskStatus = "failed"
	ComputeTaskStatusQueued    ComputeTaskStatus = "queued"
	ComputeTaskStatusRunning   ComputeTaskStatus = "running"
	ComputeTaskStatusSucceeded ComputeTaskStatus = "succeeded"
	ComputeTaskStatusTimeout   ComputeTaskStatus = "timeout"
	ComputeTaskStatusUnknown   ComputeTaskStatus = "unknown"
)

// Valid indicates whether the value is a known member of the ComputeTaskStatus enum.
func (e ComputeTaskStatus) Valid() bool {
	switch e {
	case ComputeTaskStatusCanceled:
		return true
	case ComputeTaskStatusFailed:
		return true
	case ComputeTaskStatusQueued:
		return true
	case ComputeTaskStatusRunning:
		return true
	case ComputeTaskStatusSucceeded:
		return true
	case ComputeTaskStatusTimeout:
		return true
	case ComputeTaskStatusUnknown:
		return true
	default:
		return false
	}
}

// Defines values for DeliveryPlanSource.
const (
	DeliveryPlanSourceAI     DeliveryPlanSource = "ai"
	DeliveryPlanSourceManual DeliveryPlanSource = "manual"
)

// Valid indicates whether the value is a known member of the DeliveryPlanSource enum.
func (e DeliveryPlanSource) Valid() bool {
	switch e {
	case DeliveryPlanSourceAI:
		return true
	case DeliveryPlanSourceManual:
		return true
	default:
		return false
	}
}

// Defines values for DeliveryPlanStatus.
const (
	Confirmed       DeliveryPlanStatus = "confirmed"
	Confirming      DeliveryPlanStatus = "confirming"
	Draft           DeliveryPlanStatus = "draft"
	WaitingApproval DeliveryPlanStatus = "waiting_approval"
)

// Valid indicates whether the value is a known member of the DeliveryPlanStatus enum.
func (e DeliveryPlanStatus) Valid() bool {
	switch e {
	case Confirmed:
		return true
	case Confirming:
		return true
	case Draft:
		return true
	case WaitingApproval:
		return true
	default:
		return false
	}
}

// Defines values for DeliveryPlanInputSource.
const (
	DeliveryPlanInputSourceAI     DeliveryPlanInputSource = "ai"
	DeliveryPlanInputSourceManual DeliveryPlanInputSource = "manual"
)

// Valid indicates whether the value is a known member of the DeliveryPlanInputSource enum.
func (e DeliveryPlanInputSource) Valid() bool {
	switch e {
	case DeliveryPlanInputSourceAI:
		return true
	case DeliveryPlanInputSourceManual:
		return true
	default:
		return false
	}
}

// Defines values for EvaluationExecutorProfileIsolationMode.
const (
	EvaluationExecutorProfileIsolationModeDisposableWrite EvaluationExecutorProfileIsolationMode = "disposable-write"
	EvaluationExecutorProfileIsolationModeReadOnly        EvaluationExecutorProfileIsolationMode = "read-only"
)

// Valid indicates whether the value is a known member of the EvaluationExecutorProfileIsolationMode enum.
func (e EvaluationExecutorProfileIsolationMode) Valid() bool {
	switch e {
	case EvaluationExecutorProfileIsolationModeDisposableWrite:
		return true
	case EvaluationExecutorProfileIsolationModeReadOnly:
		return true
	default:
		return false
	}
}

// Defines values for EvaluationFeedbackInputDisposition.
const (
	EvaluationFeedbackInputDispositionAccepted EvaluationFeedbackInputDisposition = "accepted"
	EvaluationFeedbackInputDispositionDeleted  EvaluationFeedbackInputDisposition = "deleted"
	EvaluationFeedbackInputDispositionPending  EvaluationFeedbackInputDisposition = "pending"
	EvaluationFeedbackInputDispositionRejected EvaluationFeedbackInputDisposition = "rejected"
)

// Valid indicates whether the value is a known member of the EvaluationFeedbackInputDisposition enum.
func (e EvaluationFeedbackInputDisposition) Valid() bool {
	switch e {
	case EvaluationFeedbackInputDispositionAccepted:
		return true
	case EvaluationFeedbackInputDispositionDeleted:
		return true
	case EvaluationFeedbackInputDispositionPending:
		return true
	case EvaluationFeedbackInputDispositionRejected:
		return true
	default:
		return false
	}
}

// Defines values for EvaluationRunStatus.
const (
	EvaluationRunStatusCancelled EvaluationRunStatus = "cancelled"
	EvaluationRunStatusCompleted EvaluationRunStatus = "completed"
	EvaluationRunStatusFailed    EvaluationRunStatus = "failed"
	EvaluationRunStatusQueued    EvaluationRunStatus = "queued"
	EvaluationRunStatusRunning   EvaluationRunStatus = "running"
)

// Valid indicates whether the value is a known member of the EvaluationRunStatus enum.
func (e EvaluationRunStatus) Valid() bool {
	switch e {
	case EvaluationRunStatusCancelled:
		return true
	case EvaluationRunStatusCompleted:
		return true
	case EvaluationRunStatusFailed:
		return true
	case EvaluationRunStatusQueued:
		return true
	case EvaluationRunStatusRunning:
		return true
	default:
		return false
	}
}

// Defines values for KnowledgeAccessScopeVisibility.
const (
	KnowledgeAccessScopeVisibilityPrivate KnowledgeAccessScopeVisibility = "private"
	KnowledgeAccessScopeVisibilityPublic  KnowledgeAccessScopeVisibility = "public"
	KnowledgeAccessScopeVisibilityScoped  KnowledgeAccessScopeVisibility = "scoped"
)

// Valid indicates whether the value is a known member of the KnowledgeAccessScopeVisibility enum.
func (e KnowledgeAccessScopeVisibility) Valid() bool {
	switch e {
	case KnowledgeAccessScopeVisibilityPrivate:
		return true
	case KnowledgeAccessScopeVisibilityPublic:
		return true
	case KnowledgeAccessScopeVisibilityScoped:
		return true
	default:
		return false
	}
}

// Defines values for KnowledgeBaseStatus.
const (
	KnowledgeBaseStatusActive   KnowledgeBaseStatus = "active"
	KnowledgeBaseStatusDisabled KnowledgeBaseStatus = "disabled"
)

// Valid indicates whether the value is a known member of the KnowledgeBaseStatus enum.
func (e KnowledgeBaseStatus) Valid() bool {
	switch e {
	case KnowledgeBaseStatusActive:
		return true
	case KnowledgeBaseStatusDisabled:
		return true
	default:
		return false
	}
}

// Defines values for KnowledgeConnectorKind.
const (
	KnowledgeConnectorKindGit    KnowledgeConnectorKind = "git"
	KnowledgeConnectorKindHTTP   KnowledgeConnectorKind = "http"
	KnowledgeConnectorKindInline KnowledgeConnectorKind = "inline"
	KnowledgeConnectorKindObject KnowledgeConnectorKind = "object"
)

// Valid indicates whether the value is a known member of the KnowledgeConnectorKind enum.
func (e KnowledgeConnectorKind) Valid() bool {
	switch e {
	case KnowledgeConnectorKindGit:
		return true
	case KnowledgeConnectorKindHTTP:
		return true
	case KnowledgeConnectorKindInline:
		return true
	case KnowledgeConnectorKindObject:
		return true
	default:
		return false
	}
}

// Defines values for KnowledgeConnectorInputKind.
const (
	KnowledgeConnectorInputKindGit    KnowledgeConnectorInputKind = "git"
	KnowledgeConnectorInputKindHTTP   KnowledgeConnectorInputKind = "http"
	KnowledgeConnectorInputKindInline KnowledgeConnectorInputKind = "inline"
	KnowledgeConnectorInputKindObject KnowledgeConnectorInputKind = "object"
)

// Valid indicates whether the value is a known member of the KnowledgeConnectorInputKind enum.
func (e KnowledgeConnectorInputKind) Valid() bool {
	switch e {
	case KnowledgeConnectorInputKindGit:
		return true
	case KnowledgeConnectorInputKindHTTP:
		return true
	case KnowledgeConnectorInputKindInline:
		return true
	case KnowledgeConnectorInputKindObject:
		return true
	default:
		return false
	}
}

// Defines values for KnowledgeConnectorValidationKind.
const (
	KnowledgeConnectorValidationKindGit    KnowledgeConnectorValidationKind = "git"
	KnowledgeConnectorValidationKindHTTP   KnowledgeConnectorValidationKind = "http"
	KnowledgeConnectorValidationKindInline KnowledgeConnectorValidationKind = "inline"
	KnowledgeConnectorValidationKindObject KnowledgeConnectorValidationKind = "object"
)

// Valid indicates whether the value is a known member of the KnowledgeConnectorValidationKind enum.
func (e KnowledgeConnectorValidationKind) Valid() bool {
	switch e {
	case KnowledgeConnectorValidationKindGit:
		return true
	case KnowledgeConnectorValidationKindHTTP:
		return true
	case KnowledgeConnectorValidationKindInline:
		return true
	case KnowledgeConnectorValidationKindObject:
		return true
	default:
		return false
	}
}

// Defines values for KnowledgeDocumentStatus.
const (
	KnowledgeDocumentStatusDeleted KnowledgeDocumentStatus = "deleted"
	KnowledgeDocumentStatusFailed  KnowledgeDocumentStatus = "failed"
	KnowledgeDocumentStatusIndexed KnowledgeDocumentStatus = "indexed"
	KnowledgeDocumentStatusPending KnowledgeDocumentStatus = "pending"
)

// Valid indicates whether the value is a known member of the KnowledgeDocumentStatus enum.
func (e KnowledgeDocumentStatus) Valid() bool {
	switch e {
	case KnowledgeDocumentStatusDeleted:
		return true
	case KnowledgeDocumentStatusFailed:
		return true
	case KnowledgeDocumentStatusIndexed:
		return true
	case KnowledgeDocumentStatusPending:
		return true
	default:
		return false
	}
}

// Defines values for KnowledgeIndexRevisionStatus.
const (
	KnowledgeIndexRevisionStatusActive     KnowledgeIndexRevisionStatus = "active"
	KnowledgeIndexRevisionStatusBuilding   KnowledgeIndexRevisionStatus = "building"
	KnowledgeIndexRevisionStatusFailed     KnowledgeIndexRevisionStatus = "failed"
	KnowledgeIndexRevisionStatusSuperseded KnowledgeIndexRevisionStatus = "superseded"
)

// Valid indicates whether the value is a known member of the KnowledgeIndexRevisionStatus enum.
func (e KnowledgeIndexRevisionStatus) Valid() bool {
	switch e {
	case KnowledgeIndexRevisionStatusActive:
		return true
	case KnowledgeIndexRevisionStatusBuilding:
		return true
	case KnowledgeIndexRevisionStatusFailed:
		return true
	case KnowledgeIndexRevisionStatusSuperseded:
		return true
	default:
		return false
	}
}

// Defines values for KnowledgeIngestionJobStatus.
const (
	KnowledgeIngestionJobStatusCancelled  KnowledgeIngestionJobStatus = "cancelled"
	KnowledgeIngestionJobStatusCancelling KnowledgeIngestionJobStatus = "cancelling"
	KnowledgeIngestionJobStatusFailed     KnowledgeIngestionJobStatus = "failed"
	KnowledgeIngestionJobStatusQueued     KnowledgeIngestionJobStatus = "queued"
	KnowledgeIngestionJobStatusRetryWait  KnowledgeIngestionJobStatus = "retry_wait"
	KnowledgeIngestionJobStatusRunning    KnowledgeIngestionJobStatus = "running"
	KnowledgeIngestionJobStatusSucceeded  KnowledgeIngestionJobStatus = "succeeded"
)

// Valid indicates whether the value is a known member of the KnowledgeIngestionJobStatus enum.
func (e KnowledgeIngestionJobStatus) Valid() bool {
	switch e {
	case KnowledgeIngestionJobStatusCancelled:
		return true
	case KnowledgeIngestionJobStatusCancelling:
		return true
	case KnowledgeIngestionJobStatusFailed:
		return true
	case KnowledgeIngestionJobStatusQueued:
		return true
	case KnowledgeIngestionJobStatusRetryWait:
		return true
	case KnowledgeIngestionJobStatusRunning:
		return true
	case KnowledgeIngestionJobStatusSucceeded:
		return true
	default:
		return false
	}
}

// Defines values for KnowledgeIngestionStage.
const (
	Chunking    KnowledgeIngestionStage = "chunking"
	Discovering KnowledgeIngestionStage = "discovering"
	Embedding   KnowledgeIngestionStage = "embedding"
	Fetching    KnowledgeIngestionStage = "fetching"
	Indexing    KnowledgeIngestionStage = "indexing"
	Parsing     KnowledgeIngestionStage = "parsing"
	Publishing  KnowledgeIngestionStage = "publishing"
	Verifying   KnowledgeIngestionStage = "verifying"
)

// Valid indicates whether the value is a known member of the KnowledgeIngestionStage enum.
func (e KnowledgeIngestionStage) Valid() bool {
	switch e {
	case Chunking:
		return true
	case Discovering:
		return true
	case Embedding:
		return true
	case Fetching:
		return true
	case Indexing:
		return true
	case Parsing:
		return true
	case Publishing:
		return true
	case Verifying:
		return true
	default:
		return false
	}
}

// Defines values for KnowledgeSourceKind.
const (
	KnowledgeSourceKindGit    KnowledgeSourceKind = "git"
	KnowledgeSourceKindHTTP   KnowledgeSourceKind = "http"
	KnowledgeSourceKindInline KnowledgeSourceKind = "inline"
	KnowledgeSourceKindObject KnowledgeSourceKind = "object"
)

// Valid indicates whether the value is a known member of the KnowledgeSourceKind enum.
func (e KnowledgeSourceKind) Valid() bool {
	switch e {
	case KnowledgeSourceKindGit:
		return true
	case KnowledgeSourceKindHTTP:
		return true
	case KnowledgeSourceKindInline:
		return true
	case KnowledgeSourceKindObject:
		return true
	default:
		return false
	}
}

// Defines values for KnowledgeSourceStatus.
const (
	KnowledgeSourceStatusFailed  KnowledgeSourceStatus = "failed"
	KnowledgeSourceStatusPending KnowledgeSourceStatus = "pending"
	KnowledgeSourceStatusReady   KnowledgeSourceStatus = "ready"
	KnowledgeSourceStatusSyncing KnowledgeSourceStatus = "syncing"
)

// Valid indicates whether the value is a known member of the KnowledgeSourceStatus enum.
func (e KnowledgeSourceStatus) Valid() bool {
	switch e {
	case KnowledgeSourceStatusFailed:
		return true
	case KnowledgeSourceStatusPending:
		return true
	case KnowledgeSourceStatusReady:
		return true
	case KnowledgeSourceStatusSyncing:
		return true
	default:
		return false
	}
}

// Defines values for KnowledgeSourceInputKind.
const (
	KnowledgeSourceInputKindGit    KnowledgeSourceInputKind = "git"
	KnowledgeSourceInputKindHTTP   KnowledgeSourceInputKind = "http"
	KnowledgeSourceInputKindInline KnowledgeSourceInputKind = "inline"
	KnowledgeSourceInputKindObject KnowledgeSourceInputKind = "object"
)

// Valid indicates whether the value is a known member of the KnowledgeSourceInputKind enum.
func (e KnowledgeSourceInputKind) Valid() bool {
	switch e {
	case KnowledgeSourceInputKindGit:
		return true
	case KnowledgeSourceInputKindHTTP:
		return true
	case KnowledgeSourceInputKindInline:
		return true
	case KnowledgeSourceInputKindObject:
		return true
	default:
		return false
	}
}

// Defines values for KnowledgeSyncRunStatus.
const (
	KnowledgeSyncRunStatusFailed    KnowledgeSyncRunStatus = "failed"
	KnowledgeSyncRunStatusQueued    KnowledgeSyncRunStatus = "queued"
	KnowledgeSyncRunStatusRunning   KnowledgeSyncRunStatus = "running"
	KnowledgeSyncRunStatusSucceeded KnowledgeSyncRunStatus = "succeeded"
)

// Valid indicates whether the value is a known member of the KnowledgeSyncRunStatus enum.
func (e KnowledgeSyncRunStatus) Valid() bool {
	switch e {
	case KnowledgeSyncRunStatusFailed:
		return true
	case KnowledgeSyncRunStatusQueued:
		return true
	case KnowledgeSyncRunStatusRunning:
		return true
	case KnowledgeSyncRunStatusSucceeded:
		return true
	default:
		return false
	}
}

// Defines values for KubernetesResourceAction.
const (
	KubernetesResourceActionCreate KubernetesResourceAction = "create"
	KubernetesResourceActionDelete KubernetesResourceAction = "delete"
	KubernetesResourceActionExec   KubernetesResourceAction = "exec"
	KubernetesResourceActionList   KubernetesResourceAction = "list"
	KubernetesResourceActionUpdate KubernetesResourceAction = "update"
	KubernetesResourceActionView   KubernetesResourceAction = "view"
)

// Valid indicates whether the value is a known member of the KubernetesResourceAction enum.
func (e KubernetesResourceAction) Valid() bool {
	switch e {
	case KubernetesResourceActionCreate:
		return true
	case KubernetesResourceActionDelete:
		return true
	case KubernetesResourceActionExec:
		return true
	case KubernetesResourceActionList:
		return true
	case KubernetesResourceActionUpdate:
		return true
	case KubernetesResourceActionView:
		return true
	default:
		return false
	}
}

// Defines values for KubernetesResourceCapabilityMode.
const (
	Agent  KubernetesResourceCapabilityMode = "agent"
	Direct KubernetesResourceCapabilityMode = "direct"
)

// Valid indicates whether the value is a known member of the KubernetesResourceCapabilityMode enum.
func (e KubernetesResourceCapabilityMode) Valid() bool {
	switch e {
	case Agent:
		return true
	case Direct:
		return true
	default:
		return false
	}
}

// Defines values for KubernetesResourceCreateBatchStatus.
const (
	KubernetesResourceCreateBatchStatusFailed    KubernetesResourceCreateBatchStatus = "failed"
	KubernetesResourceCreateBatchStatusPartial   KubernetesResourceCreateBatchStatus = "partial"
	KubernetesResourceCreateBatchStatusSucceeded KubernetesResourceCreateBatchStatus = "succeeded"
)

// Valid indicates whether the value is a known member of the KubernetesResourceCreateBatchStatus enum.
func (e KubernetesResourceCreateBatchStatus) Valid() bool {
	switch e {
	case KubernetesResourceCreateBatchStatusFailed:
		return true
	case KubernetesResourceCreateBatchStatusPartial:
		return true
	case KubernetesResourceCreateBatchStatusSucceeded:
		return true
	default:
		return false
	}
}

// Defines values for KubernetesResourceCreateErrorCode.
const (
	KubernetesResourceCreateErrorCodeClusterScopedNamespaceIgnored KubernetesResourceCreateErrorCode = "cluster_scoped_namespace_ignored"
	KubernetesResourceCreateErrorCodeHighRiskPermissionRequired    KubernetesResourceCreateErrorCode = "high_risk_permission_required"
	KubernetesResourceCreateErrorCodeMultiDocumentNotAllowed       KubernetesResourceCreateErrorCode = "multi_document_not_allowed"
	KubernetesResourceCreateErrorCodeNamespaceMismatch             KubernetesResourceCreateErrorCode = "namespace_mismatch"
	KubernetesResourceCreateErrorCodeNamespaceRequired             KubernetesResourceCreateErrorCode = "namespace_required"
	KubernetesResourceCreateErrorCodeResourceAlreadyExists         KubernetesResourceCreateErrorCode = "resource_already_exists"
	KubernetesResourceCreateErrorCodeResourceCapabilityUnsupported KubernetesResourceCreateErrorCode = "resource_capability_unsupported"
	KubernetesResourceCreateErrorCodeResourceCreateDenied          KubernetesResourceCreateErrorCode = "resource_create_denied"
	KubernetesResourceCreateErrorCodeResourceCreateFailed          KubernetesResourceCreateErrorCode = "resource_create_failed"
	KubernetesResourceCreateErrorCodeResourceDryRunFailed          KubernetesResourceCreateErrorCode = "resource_dry_run_failed"
	KubernetesResourceCreateErrorCodeResourceKindMismatch          KubernetesResourceCreateErrorCode = "resource_kind_mismatch"
)

// Valid indicates whether the value is a known member of the KubernetesResourceCreateErrorCode enum.
func (e KubernetesResourceCreateErrorCode) Valid() bool {
	switch e {
	case KubernetesResourceCreateErrorCodeClusterScopedNamespaceIgnored:
		return true
	case KubernetesResourceCreateErrorCodeHighRiskPermissionRequired:
		return true
	case KubernetesResourceCreateErrorCodeMultiDocumentNotAllowed:
		return true
	case KubernetesResourceCreateErrorCodeNamespaceMismatch:
		return true
	case KubernetesResourceCreateErrorCodeNamespaceRequired:
		return true
	case KubernetesResourceCreateErrorCodeResourceAlreadyExists:
		return true
	case KubernetesResourceCreateErrorCodeResourceCapabilityUnsupported:
		return true
	case KubernetesResourceCreateErrorCodeResourceCreateDenied:
		return true
	case KubernetesResourceCreateErrorCodeResourceCreateFailed:
		return true
	case KubernetesResourceCreateErrorCodeResourceDryRunFailed:
		return true
	case KubernetesResourceCreateErrorCodeResourceKindMismatch:
		return true
	default:
		return false
	}
}

// Defines values for KubernetesResourceCreateResultStatus.
const (
	KubernetesResourceCreateResultStatusFailed     KubernetesResourceCreateResultStatus = "failed"
	KubernetesResourceCreateResultStatusNotStarted KubernetesResourceCreateResultStatus = "not_started"
	KubernetesResourceCreateResultStatusSucceeded  KubernetesResourceCreateResultStatus = "succeeded"
)

// Valid indicates whether the value is a known member of the KubernetesResourceCreateResultStatus enum.
func (e KubernetesResourceCreateResultStatus) Valid() bool {
	switch e {
	case KubernetesResourceCreateResultStatusFailed:
		return true
	case KubernetesResourceCreateResultStatusNotStarted:
		return true
	case KubernetesResourceCreateResultStatusSucceeded:
		return true
	default:
		return false
	}
}

// Defines values for KubernetesResourceCreateScopeDecisionRequestAction.
const (
	Create KubernetesResourceCreateScopeDecisionRequestAction = "create"
)

// Valid indicates whether the value is a known member of the KubernetesResourceCreateScopeDecisionRequestAction enum.
func (e KubernetesResourceCreateScopeDecisionRequestAction) Valid() bool {
	switch e {
	case Create:
		return true
	default:
		return false
	}
}

// Defines values for KubernetesResourceCreateSource.
const (
	KubernetesResourceCreateSourceForm       KubernetesResourceCreateSource = "form"
	KubernetesResourceCreateSourceGlobalYAML KubernetesResourceCreateSource = "global_yaml"
	KubernetesResourceCreateSourceList       KubernetesResourceCreateSource = "list"
)

// Valid indicates whether the value is a known member of the KubernetesResourceCreateSource enum.
func (e KubernetesResourceCreateSource) Valid() bool {
	switch e {
	case KubernetesResourceCreateSourceForm:
		return true
	case KubernetesResourceCreateSourceGlobalYAML:
		return true
	case KubernetesResourceCreateSourceList:
		return true
	default:
		return false
	}
}

// Defines values for KubernetesResourceDryRunStatus.
const (
	KubernetesResourceDryRunStatusFailed  KubernetesResourceDryRunStatus = "failed"
	KubernetesResourceDryRunStatusPassed  KubernetesResourceDryRunStatus = "passed"
	KubernetesResourceDryRunStatusSkipped KubernetesResourceDryRunStatus = "skipped"
)

// Valid indicates whether the value is a known member of the KubernetesResourceDryRunStatus enum.
func (e KubernetesResourceDryRunStatus) Valid() bool {
	switch e {
	case KubernetesResourceDryRunStatusFailed:
		return true
	case KubernetesResourceDryRunStatusPassed:
		return true
	case KubernetesResourceDryRunStatusSkipped:
		return true
	default:
		return false
	}
}

// Defines values for KubernetesResourceScopeMode.
const (
	KubernetesResourceScopeModeCluster   KubernetesResourceScopeMode = "cluster"
	KubernetesResourceScopeModeNamespace KubernetesResourceScopeMode = "namespace"
)

// Valid indicates whether the value is a known member of the KubernetesResourceScopeMode enum.
func (e KubernetesResourceScopeMode) Valid() bool {
	switch e {
	case KubernetesResourceScopeModeCluster:
		return true
	case KubernetesResourceScopeModeNamespace:
		return true
	default:
		return false
	}
}

// Defines values for LLMCallLogCacheStatus.
const (
	LLMCallLogCacheStatusBypass       LLMCallLogCacheStatus = "bypass"
	LLMCallLogCacheStatusHit          LLMCallLogCacheStatus = "hit"
	LLMCallLogCacheStatusMiss         LLMCallLogCacheStatus = "miss"
	LLMCallLogCacheStatusStaleHit     LLMCallLogCacheStatus = "stale_hit"
	LLMCallLogCacheStatusWrite        LLMCallLogCacheStatus = "write"
	LLMCallLogCacheStatusWriteSkipped LLMCallLogCacheStatus = "write_skipped"
)

// Valid indicates whether the value is a known member of the LLMCallLogCacheStatus enum.
func (e LLMCallLogCacheStatus) Valid() bool {
	switch e {
	case LLMCallLogCacheStatusBypass:
		return true
	case LLMCallLogCacheStatusHit:
		return true
	case LLMCallLogCacheStatusMiss:
		return true
	case LLMCallLogCacheStatusStaleHit:
		return true
	case LLMCallLogCacheStatusWrite:
		return true
	case LLMCallLogCacheStatusWriteSkipped:
		return true
	default:
		return false
	}
}

// Defines values for LLMCallLogProviderKind.
const (
	LLMCallLogProviderKindAnthropic        LLMCallLogProviderKind = "anthropic"
	LLMCallLogProviderKindAzureOpenai      LLMCallLogProviderKind = "azure-openai"
	LLMCallLogProviderKindCohere           LLMCallLogProviderKind = "cohere"
	LLMCallLogProviderKindDeepseek         LLMCallLogProviderKind = "deepseek"
	LLMCallLogProviderKindGemini           LLMCallLogProviderKind = "gemini"
	LLMCallLogProviderKindOpenai           LLMCallLogProviderKind = "openai"
	LLMCallLogProviderKindOpenaiCompatible LLMCallLogProviderKind = "openai-compatible"
	LLMCallLogProviderKindOpenrouter       LLMCallLogProviderKind = "openrouter"
	LLMCallLogProviderKindQwen             LLMCallLogProviderKind = "qwen"
)

// Valid indicates whether the value is a known member of the LLMCallLogProviderKind enum.
func (e LLMCallLogProviderKind) Valid() bool {
	switch e {
	case LLMCallLogProviderKindAnthropic:
		return true
	case LLMCallLogProviderKindAzureOpenai:
		return true
	case LLMCallLogProviderKindCohere:
		return true
	case LLMCallLogProviderKindDeepseek:
		return true
	case LLMCallLogProviderKindGemini:
		return true
	case LLMCallLogProviderKindOpenai:
		return true
	case LLMCallLogProviderKindOpenaiCompatible:
		return true
	case LLMCallLogProviderKindOpenrouter:
		return true
	case LLMCallLogProviderKindQwen:
		return true
	default:
		return false
	}
}

// Defines values for LLMCallLogStatus.
const (
	LLMCallLogStatusCancelled       LLMCallLogStatus = "cancelled"
	LLMCallLogStatusClientCancelled LLMCallLogStatus = "client_cancelled"
	LLMCallLogStatusFailure         LLMCallLogStatus = "failure"
	LLMCallLogStatusPolicyDenied    LLMCallLogStatus = "policy_denied"
	LLMCallLogStatusRateLimited     LLMCallLogStatus = "rate_limited"
	LLMCallLogStatusSuccess         LLMCallLogStatus = "success"
)

// Valid indicates whether the value is a known member of the LLMCallLogStatus enum.
func (e LLMCallLogStatus) Valid() bool {
	switch e {
	case LLMCallLogStatusCancelled:
		return true
	case LLMCallLogStatusClientCancelled:
		return true
	case LLMCallLogStatusFailure:
		return true
	case LLMCallLogStatusPolicyDenied:
		return true
	case LLMCallLogStatusRateLimited:
		return true
	case LLMCallLogStatusSuccess:
		return true
	default:
		return false
	}
}

// Defines values for LLMCallLogTokenKind.
const (
	LLMCallLogTokenKindPersonalAccessToken LLMCallLogTokenKind = "personal_access_token"
	LLMCallLogTokenKindServiceAccountToken LLMCallLogTokenKind = "service_account_token"
)

// Valid indicates whether the value is a known member of the LLMCallLogTokenKind enum.
func (e LLMCallLogTokenKind) Valid() bool {
	switch e {
	case LLMCallLogTokenKindPersonalAccessToken:
		return true
	case LLMCallLogTokenKindServiceAccountToken:
		return true
	default:
		return false
	}
}

// Defines values for LLMModelRouteProviderKind.
const (
	LLMModelRouteProviderKindAnthropic        LLMModelRouteProviderKind = "anthropic"
	LLMModelRouteProviderKindAzureOpenai      LLMModelRouteProviderKind = "azure-openai"
	LLMModelRouteProviderKindCohere           LLMModelRouteProviderKind = "cohere"
	LLMModelRouteProviderKindDeepseek         LLMModelRouteProviderKind = "deepseek"
	LLMModelRouteProviderKindGemini           LLMModelRouteProviderKind = "gemini"
	LLMModelRouteProviderKindOpenai           LLMModelRouteProviderKind = "openai"
	LLMModelRouteProviderKindOpenaiCompatible LLMModelRouteProviderKind = "openai-compatible"
	LLMModelRouteProviderKindOpenrouter       LLMModelRouteProviderKind = "openrouter"
	LLMModelRouteProviderKindQwen             LLMModelRouteProviderKind = "qwen"
)

// Valid indicates whether the value is a known member of the LLMModelRouteProviderKind enum.
func (e LLMModelRouteProviderKind) Valid() bool {
	switch e {
	case LLMModelRouteProviderKindAnthropic:
		return true
	case LLMModelRouteProviderKindAzureOpenai:
		return true
	case LLMModelRouteProviderKindCohere:
		return true
	case LLMModelRouteProviderKindDeepseek:
		return true
	case LLMModelRouteProviderKindGemini:
		return true
	case LLMModelRouteProviderKindOpenai:
		return true
	case LLMModelRouteProviderKindOpenaiCompatible:
		return true
	case LLMModelRouteProviderKindOpenrouter:
		return true
	case LLMModelRouteProviderKindQwen:
		return true
	default:
		return false
	}
}

// Defines values for LLMModelRouteInputProviderKind.
const (
	LLMModelRouteInputProviderKindAnthropic        LLMModelRouteInputProviderKind = "anthropic"
	LLMModelRouteInputProviderKindAzureOpenai      LLMModelRouteInputProviderKind = "azure-openai"
	LLMModelRouteInputProviderKindCohere           LLMModelRouteInputProviderKind = "cohere"
	LLMModelRouteInputProviderKindDeepseek         LLMModelRouteInputProviderKind = "deepseek"
	LLMModelRouteInputProviderKindGemini           LLMModelRouteInputProviderKind = "gemini"
	LLMModelRouteInputProviderKindOpenai           LLMModelRouteInputProviderKind = "openai"
	LLMModelRouteInputProviderKindOpenaiCompatible LLMModelRouteInputProviderKind = "openai-compatible"
	LLMModelRouteInputProviderKindOpenrouter       LLMModelRouteInputProviderKind = "openrouter"
	LLMModelRouteInputProviderKindQwen             LLMModelRouteInputProviderKind = "qwen"
)

// Valid indicates whether the value is a known member of the LLMModelRouteInputProviderKind enum.
func (e LLMModelRouteInputProviderKind) Valid() bool {
	switch e {
	case LLMModelRouteInputProviderKindAnthropic:
		return true
	case LLMModelRouteInputProviderKindAzureOpenai:
		return true
	case LLMModelRouteInputProviderKindCohere:
		return true
	case LLMModelRouteInputProviderKindDeepseek:
		return true
	case LLMModelRouteInputProviderKindGemini:
		return true
	case LLMModelRouteInputProviderKindOpenai:
		return true
	case LLMModelRouteInputProviderKindOpenaiCompatible:
		return true
	case LLMModelRouteInputProviderKindOpenrouter:
		return true
	case LLMModelRouteInputProviderKindQwen:
		return true
	default:
		return false
	}
}

// Defines values for LLMUpstreamProviderKind.
const (
	LLMUpstreamProviderKindAnthropic        LLMUpstreamProviderKind = "anthropic"
	LLMUpstreamProviderKindAzureOpenai      LLMUpstreamProviderKind = "azure-openai"
	LLMUpstreamProviderKindCohere           LLMUpstreamProviderKind = "cohere"
	LLMUpstreamProviderKindDeepseek         LLMUpstreamProviderKind = "deepseek"
	LLMUpstreamProviderKindGemini           LLMUpstreamProviderKind = "gemini"
	LLMUpstreamProviderKindOpenai           LLMUpstreamProviderKind = "openai"
	LLMUpstreamProviderKindOpenaiCompatible LLMUpstreamProviderKind = "openai-compatible"
	LLMUpstreamProviderKindOpenrouter       LLMUpstreamProviderKind = "openrouter"
	LLMUpstreamProviderKindQwen             LLMUpstreamProviderKind = "qwen"
)

// Valid indicates whether the value is a known member of the LLMUpstreamProviderKind enum.
func (e LLMUpstreamProviderKind) Valid() bool {
	switch e {
	case LLMUpstreamProviderKindAnthropic:
		return true
	case LLMUpstreamProviderKindAzureOpenai:
		return true
	case LLMUpstreamProviderKindCohere:
		return true
	case LLMUpstreamProviderKindDeepseek:
		return true
	case LLMUpstreamProviderKindGemini:
		return true
	case LLMUpstreamProviderKindOpenai:
		return true
	case LLMUpstreamProviderKindOpenaiCompatible:
		return true
	case LLMUpstreamProviderKindOpenrouter:
		return true
	case LLMUpstreamProviderKindQwen:
		return true
	default:
		return false
	}
}

// Defines values for LLMUpstreamStatus.
const (
	LLMUpstreamStatusActive   LLMUpstreamStatus = "active"
	LLMUpstreamStatusDegraded LLMUpstreamStatus = "degraded"
	LLMUpstreamStatusDisabled LLMUpstreamStatus = "disabled"
)

// Valid indicates whether the value is a known member of the LLMUpstreamStatus enum.
func (e LLMUpstreamStatus) Valid() bool {
	switch e {
	case LLMUpstreamStatusActive:
		return true
	case LLMUpstreamStatusDegraded:
		return true
	case LLMUpstreamStatusDisabled:
		return true
	default:
		return false
	}
}

// Defines values for LLMUpstreamInputProviderKind.
const (
	LLMUpstreamInputProviderKindAnthropic        LLMUpstreamInputProviderKind = "anthropic"
	LLMUpstreamInputProviderKindAzureOpenai      LLMUpstreamInputProviderKind = "azure-openai"
	LLMUpstreamInputProviderKindCohere           LLMUpstreamInputProviderKind = "cohere"
	LLMUpstreamInputProviderKindDeepseek         LLMUpstreamInputProviderKind = "deepseek"
	LLMUpstreamInputProviderKindGemini           LLMUpstreamInputProviderKind = "gemini"
	LLMUpstreamInputProviderKindOpenai           LLMUpstreamInputProviderKind = "openai"
	LLMUpstreamInputProviderKindOpenaiCompatible LLMUpstreamInputProviderKind = "openai-compatible"
	LLMUpstreamInputProviderKindOpenrouter       LLMUpstreamInputProviderKind = "openrouter"
	LLMUpstreamInputProviderKindQwen             LLMUpstreamInputProviderKind = "qwen"
)

// Valid indicates whether the value is a known member of the LLMUpstreamInputProviderKind enum.
func (e LLMUpstreamInputProviderKind) Valid() bool {
	switch e {
	case LLMUpstreamInputProviderKindAnthropic:
		return true
	case LLMUpstreamInputProviderKindAzureOpenai:
		return true
	case LLMUpstreamInputProviderKindCohere:
		return true
	case LLMUpstreamInputProviderKindDeepseek:
		return true
	case LLMUpstreamInputProviderKindGemini:
		return true
	case LLMUpstreamInputProviderKindOpenai:
		return true
	case LLMUpstreamInputProviderKindOpenaiCompatible:
		return true
	case LLMUpstreamInputProviderKindOpenrouter:
		return true
	case LLMUpstreamInputProviderKindQwen:
		return true
	default:
		return false
	}
}

// Defines values for LLMUpstreamInputStatus.
const (
	LLMUpstreamInputStatusActive   LLMUpstreamInputStatus = "active"
	LLMUpstreamInputStatusDegraded LLMUpstreamInputStatus = "degraded"
	LLMUpstreamInputStatusDisabled LLMUpstreamInputStatus = "disabled"
)

// Valid indicates whether the value is a known member of the LLMUpstreamInputStatus enum.
func (e LLMUpstreamInputStatus) Valid() bool {
	switch e {
	case LLMUpstreamInputStatusActive:
		return true
	case LLMUpstreamInputStatusDegraded:
		return true
	case LLMUpstreamInputStatusDisabled:
		return true
	default:
		return false
	}
}

// Defines values for LLMUpstreamTestRequestEndpoint.
const (
	AudioSpeech         LLMUpstreamTestRequestEndpoint = "audio.speech"
	AudioTranscriptions LLMUpstreamTestRequestEndpoint = "audio.transcriptions"
	AudioTranslations   LLMUpstreamTestRequestEndpoint = "audio.translations"
	ChatCompletions     LLMUpstreamTestRequestEndpoint = "chat.completions"
	ImagesEdits         LLMUpstreamTestRequestEndpoint = "images.edits"
	ImagesGenerations   LLMUpstreamTestRequestEndpoint = "images.generations"
	ImagesVariations    LLMUpstreamTestRequestEndpoint = "images.variations"
	Interactions        LLMUpstreamTestRequestEndpoint = "interactions"
	Messages            LLMUpstreamTestRequestEndpoint = "messages"
	Models              LLMUpstreamTestRequestEndpoint = "models"
	Rerank              LLMUpstreamTestRequestEndpoint = "rerank"
	Responses           LLMUpstreamTestRequestEndpoint = "responses"
)

// Valid indicates whether the value is a known member of the LLMUpstreamTestRequestEndpoint enum.
func (e LLMUpstreamTestRequestEndpoint) Valid() bool {
	switch e {
	case AudioSpeech:
		return true
	case AudioTranscriptions:
		return true
	case AudioTranslations:
		return true
	case ChatCompletions:
		return true
	case ImagesEdits:
		return true
	case ImagesGenerations:
		return true
	case ImagesVariations:
		return true
	case Interactions:
		return true
	case Messages:
		return true
	case Models:
		return true
	case Rerank:
		return true
	case Responses:
		return true
	default:
		return false
	}
}

// Defines values for LLMUpstreamTestResultProviderKind.
const (
	LLMUpstreamTestResultProviderKindAnthropic        LLMUpstreamTestResultProviderKind = "anthropic"
	LLMUpstreamTestResultProviderKindAzureOpenai      LLMUpstreamTestResultProviderKind = "azure-openai"
	LLMUpstreamTestResultProviderKindCohere           LLMUpstreamTestResultProviderKind = "cohere"
	LLMUpstreamTestResultProviderKindDeepseek         LLMUpstreamTestResultProviderKind = "deepseek"
	LLMUpstreamTestResultProviderKindGemini           LLMUpstreamTestResultProviderKind = "gemini"
	LLMUpstreamTestResultProviderKindOpenai           LLMUpstreamTestResultProviderKind = "openai"
	LLMUpstreamTestResultProviderKindOpenaiCompatible LLMUpstreamTestResultProviderKind = "openai-compatible"
	LLMUpstreamTestResultProviderKindOpenrouter       LLMUpstreamTestResultProviderKind = "openrouter"
	LLMUpstreamTestResultProviderKindQwen             LLMUpstreamTestResultProviderKind = "qwen"
)

// Valid indicates whether the value is a known member of the LLMUpstreamTestResultProviderKind enum.
func (e LLMUpstreamTestResultProviderKind) Valid() bool {
	switch e {
	case LLMUpstreamTestResultProviderKindAnthropic:
		return true
	case LLMUpstreamTestResultProviderKindAzureOpenai:
		return true
	case LLMUpstreamTestResultProviderKindCohere:
		return true
	case LLMUpstreamTestResultProviderKindDeepseek:
		return true
	case LLMUpstreamTestResultProviderKindGemini:
		return true
	case LLMUpstreamTestResultProviderKindOpenai:
		return true
	case LLMUpstreamTestResultProviderKindOpenaiCompatible:
		return true
	case LLMUpstreamTestResultProviderKindOpenrouter:
		return true
	case LLMUpstreamTestResultProviderKindQwen:
		return true
	default:
		return false
	}
}

// Defines values for LLMUpstreamTestResultStatus.
const (
	LLMUpstreamTestResultStatusFailure LLMUpstreamTestResultStatus = "failure"
	LLMUpstreamTestResultStatusSuccess LLMUpstreamTestResultStatus = "success"
)

// Valid indicates whether the value is a known member of the LLMUpstreamTestResultStatus enum.
func (e LLMUpstreamTestResultStatus) Valid() bool {
	switch e {
	case LLMUpstreamTestResultStatusFailure:
		return true
	case LLMUpstreamTestResultStatusSuccess:
		return true
	default:
		return false
	}
}

// Defines values for MarketplaceAdvisorySeverity.
const (
	Critical MarketplaceAdvisorySeverity = "critical"
	High     MarketplaceAdvisorySeverity = "high"
	Low      MarketplaceAdvisorySeverity = "low"
	Medium   MarketplaceAdvisorySeverity = "medium"
)

// Valid indicates whether the value is a known member of the MarketplaceAdvisorySeverity enum.
func (e MarketplaceAdvisorySeverity) Valid() bool {
	switch e {
	case Critical:
		return true
	case High:
		return true
	case Low:
		return true
	case Medium:
		return true
	default:
		return false
	}
}

// Defines values for MarketplacePublisherVerificationLevel.
const (
	MarketplacePublisherVerificationLevelCommunity MarketplacePublisherVerificationLevel = "community"
	MarketplacePublisherVerificationLevelOfficial  MarketplacePublisherVerificationLevel = "official"
	MarketplacePublisherVerificationLevelPrivate   MarketplacePublisherVerificationLevel = "private"
	MarketplacePublisherVerificationLevelVerified  MarketplacePublisherVerificationLevel = "verified"
)

// Valid indicates whether the value is a known member of the MarketplacePublisherVerificationLevel enum.
func (e MarketplacePublisherVerificationLevel) Valid() bool {
	switch e {
	case MarketplacePublisherVerificationLevelCommunity:
		return true
	case MarketplacePublisherVerificationLevelOfficial:
		return true
	case MarketplacePublisherVerificationLevelPrivate:
		return true
	case MarketplacePublisherVerificationLevelVerified:
		return true
	default:
		return false
	}
}

// Defines values for OpenAIAudioTranscriptionRequestTimestampGranularities.
const (
	Segment OpenAIAudioTranscriptionRequestTimestampGranularities = "segment"
	Word    OpenAIAudioTranscriptionRequestTimestampGranularities = "word"
)

// Valid indicates whether the value is a known member of the OpenAIAudioTranscriptionRequestTimestampGranularities enum.
func (e OpenAIAudioTranscriptionRequestTimestampGranularities) Valid() bool {
	switch e {
	case Segment:
		return true
	case Word:
		return true
	default:
		return false
	}
}

// Defines values for OpenAIModelsResponseObject.
const (
	List OpenAIModelsResponseObject = "list"
)

// Valid indicates whether the value is a known member of the OpenAIModelsResponseObject enum.
func (e OpenAIModelsResponseObject) Valid() bool {
	switch e {
	case List:
		return true
	default:
		return false
	}
}

// Defines values for PluginComputeContainerRuntimeProviderResourceKinds.
const (
	PluginComputeContainerRuntimeResourceKindContainer   PluginComputeContainerRuntimeProviderResourceKinds = "container"
	PluginComputeContainerRuntimeResourceKindPort        PluginComputeContainerRuntimeProviderResourceKinds = "port"
	PluginComputeContainerRuntimeResourceKindProject     PluginComputeContainerRuntimeProviderResourceKinds = "project"
	PluginComputeContainerRuntimeResourceKindRuntimeHost PluginComputeContainerRuntimeProviderResourceKinds = "runtime_host"
	PluginComputeContainerRuntimeResourceKindService     PluginComputeContainerRuntimeProviderResourceKinds = "service"
	PluginComputeContainerRuntimeResourceKindTemplate    PluginComputeContainerRuntimeProviderResourceKinds = "template"
)

// Valid indicates whether the value is a known member of the PluginComputeContainerRuntimeProviderResourceKinds enum.
func (e PluginComputeContainerRuntimeProviderResourceKinds) Valid() bool {
	switch e {
	case PluginComputeContainerRuntimeResourceKindContainer:
		return true
	case PluginComputeContainerRuntimeResourceKindPort:
		return true
	case PluginComputeContainerRuntimeResourceKindProject:
		return true
	case PluginComputeContainerRuntimeResourceKindRuntimeHost:
		return true
	case PluginComputeContainerRuntimeResourceKindService:
		return true
	case PluginComputeContainerRuntimeResourceKindTemplate:
		return true
	default:
		return false
	}
}

// Defines values for PluginComputeVirtualizationProviderResourceKinds.
const (
	PluginComputeVirtualizationResourceKindCluster    PluginComputeVirtualizationProviderResourceKinds = "cluster"
	PluginComputeVirtualizationResourceKindConnection PluginComputeVirtualizationProviderResourceKinds = "connection"
	PluginComputeVirtualizationResourceKindFlavor     PluginComputeVirtualizationProviderResourceKinds = "flavor"
	PluginComputeVirtualizationResourceKindImage      PluginComputeVirtualizationProviderResourceKinds = "image"
	PluginComputeVirtualizationResourceKindVM         PluginComputeVirtualizationProviderResourceKinds = "vm"
)

// Valid indicates whether the value is a known member of the PluginComputeVirtualizationProviderResourceKinds enum.
func (e PluginComputeVirtualizationProviderResourceKinds) Valid() bool {
	switch e {
	case PluginComputeVirtualizationResourceKindCluster:
		return true
	case PluginComputeVirtualizationResourceKindConnection:
		return true
	case PluginComputeVirtualizationResourceKindFlavor:
		return true
	case PluginComputeVirtualizationResourceKindImage:
		return true
	case PluginComputeVirtualizationResourceKindVM:
		return true
	default:
		return false
	}
}

// Defines values for PluginRuntimeSpecMode.
const (
	ExternalHTTP     PluginRuntimeSpecMode = "external-http"
	ManagedContainer PluginRuntimeSpecMode = "managed-container"
	ManifestOnly     PluginRuntimeSpecMode = "manifest-only"
)

// Valid indicates whether the value is a known member of the PluginRuntimeSpecMode enum.
func (e PluginRuntimeSpecMode) Valid() bool {
	switch e {
	case ExternalHTTP:
		return true
	case ManagedContainer:
		return true
	case ManifestOnly:
		return true
	default:
		return false
	}
}

// Defines values for RepositoryProtocol.
const (
	HTTPS RepositoryProtocol = "https"
	SSH   RepositoryProtocol = "ssh"
)

// Valid indicates whether the value is a known member of the RepositoryProtocol enum.
func (e RepositoryProtocol) Valid() bool {
	switch e {
	case HTTPS:
		return true
	case SSH:
		return true
	default:
		return false
	}
}

// Defines values for RepositoryProvider.
const (
	RepositoryProviderGit    RepositoryProvider = "git"
	RepositoryProviderGitlab RepositoryProvider = "gitlab"
)

// Valid indicates whether the value is a known member of the RepositoryProvider enum.
func (e RepositoryProvider) Valid() bool {
	switch e {
	case RepositoryProviderGit:
		return true
	case RepositoryProviderGitlab:
		return true
	default:
		return false
	}
}

// Defines values for RuntimeConfigApplicationStatus.
const (
	RuntimeConfigApplicationStatusApplied          RuntimeConfigApplicationStatus = "applied"
	RuntimeConfigApplicationStatusApplying         RuntimeConfigApplicationStatus = "applying"
	RuntimeConfigApplicationStatusFailed           RuntimeConfigApplicationStatus = "failed"
	RuntimeConfigApplicationStatusPartiallyApplied RuntimeConfigApplicationStatus = "partially_applied"
	RuntimeConfigApplicationStatusPending          RuntimeConfigApplicationStatus = "pending"
	RuntimeConfigApplicationStatusRestartRequired  RuntimeConfigApplicationStatus = "restart_required"
	RuntimeConfigApplicationStatusRolledBack       RuntimeConfigApplicationStatus = "rolled_back"
)

// Valid indicates whether the value is a known member of the RuntimeConfigApplicationStatus enum.
func (e RuntimeConfigApplicationStatus) Valid() bool {
	switch e {
	case RuntimeConfigApplicationStatusApplied:
		return true
	case RuntimeConfigApplicationStatusApplying:
		return true
	case RuntimeConfigApplicationStatusFailed:
		return true
	case RuntimeConfigApplicationStatusPartiallyApplied:
		return true
	case RuntimeConfigApplicationStatusPending:
		return true
	case RuntimeConfigApplicationStatusRestartRequired:
		return true
	case RuntimeConfigApplicationStatusRolledBack:
		return true
	default:
		return false
	}
}

// Defines values for RuntimeConfigApplyMode.
const (
	RuntimeConfigApplyModeHot         RuntimeConfigApplyMode = "hot"
	RuntimeConfigApplyModeLifecycle   RuntimeConfigApplyMode = "lifecycle"
	RuntimeConfigApplyModeReconfigure RuntimeConfigApplyMode = "reconfigure"
	RuntimeConfigApplyModeRestart     RuntimeConfigApplyMode = "restart"
)

// Valid indicates whether the value is a known member of the RuntimeConfigApplyMode enum.
func (e RuntimeConfigApplyMode) Valid() bool {
	switch e {
	case RuntimeConfigApplyModeHot:
		return true
	case RuntimeConfigApplyModeLifecycle:
		return true
	case RuntimeConfigApplyModeReconfigure:
		return true
	case RuntimeConfigApplyModeRestart:
		return true
	default:
		return false
	}
}

// Defines values for RuntimeConfigIssueSeverity.
const (
	RuntimeConfigIssueSeverityError   RuntimeConfigIssueSeverity = "error"
	RuntimeConfigIssueSeverityWarning RuntimeConfigIssueSeverity = "warning"
)

// Valid indicates whether the value is a known member of the RuntimeConfigIssueSeverity enum.
func (e RuntimeConfigIssueSeverity) Valid() bool {
	switch e {
	case RuntimeConfigIssueSeverityError:
		return true
	case RuntimeConfigIssueSeverityWarning:
		return true
	default:
		return false
	}
}

// Defines values for RuntimeConfigSource.
const (
	RuntimeConfigSourceConfigFile      RuntimeConfigSource = "config_file"
	RuntimeConfigSourceDefault         RuntimeConfigSource = "default"
	RuntimeConfigSourceEnvironment     RuntimeConfigSource = "environment"
	RuntimeConfigSourceRuntimeOverride RuntimeConfigSource = "runtime_override"
	RuntimeConfigSourceSecret          RuntimeConfigSource = "secret"
)

// Valid indicates whether the value is a known member of the RuntimeConfigSource enum.
func (e RuntimeConfigSource) Valid() bool {
	switch e {
	case RuntimeConfigSourceConfigFile:
		return true
	case RuntimeConfigSourceDefault:
		return true
	case RuntimeConfigSourceEnvironment:
		return true
	case RuntimeConfigSourceRuntimeOverride:
		return true
	case RuntimeConfigSourceSecret:
		return true
	default:
		return false
	}
}

// Defines values for RuntimeConfigValueType.
const (
	RuntimeConfigValueTypeBoolean    RuntimeConfigValueType = "boolean"
	RuntimeConfigValueTypeDuration   RuntimeConfigValueType = "duration"
	RuntimeConfigValueTypeInteger    RuntimeConfigValueType = "integer"
	RuntimeConfigValueTypeNumber     RuntimeConfigValueType = "number"
	RuntimeConfigValueTypeString     RuntimeConfigValueType = "string"
	RuntimeConfigValueTypeStringList RuntimeConfigValueType = "string_list"
	RuntimeConfigValueTypeURL        RuntimeConfigValueType = "url"
)

// Valid indicates whether the value is a known member of the RuntimeConfigValueType enum.
func (e RuntimeConfigValueType) Valid() bool {
	switch e {
	case RuntimeConfigValueTypeBoolean:
		return true
	case RuntimeConfigValueTypeDuration:
		return true
	case RuntimeConfigValueTypeInteger:
		return true
	case RuntimeConfigValueTypeNumber:
		return true
	case RuntimeConfigValueTypeString:
		return true
	case RuntimeConfigValueTypeStringList:
		return true
	case RuntimeConfigValueTypeURL:
		return true
	default:
		return false
	}
}

// Defines values for RuntimeNetworkUsageScope.
const (
	NetworkNamespace RuntimeNetworkUsageScope = "network_namespace"
	Process          RuntimeNetworkUsageScope = "process"
	Unavailable      RuntimeNetworkUsageScope = "unavailable"
)

// Valid indicates whether the value is a known member of the RuntimeNetworkUsageScope enum.
func (e RuntimeNetworkUsageScope) Valid() bool {
	switch e {
	case NetworkNamespace:
		return true
	case Process:
		return true
	case Unavailable:
		return true
	default:
		return false
	}
}

// Defines values for SourceConnectionCapabilities.
const (
	Branches     SourceConnectionCapabilities = "branches"
	Files        SourceConnectionCapabilities = "files"
	Repositories SourceConnectionCapabilities = "repositories"
	Tags         SourceConnectionCapabilities = "tags"
)

// Valid indicates whether the value is a known member of the SourceConnectionCapabilities enum.
func (e SourceConnectionCapabilities) Valid() bool {
	switch e {
	case Branches:
		return true
	case Files:
		return true
	case Repositories:
		return true
	case Tags:
		return true
	default:
		return false
	}
}

// Defines values for SourceFileEncoding.
const (
	SourceFileEncodingBase64 SourceFileEncoding = "base64"
	SourceFileEncodingUTF8   SourceFileEncoding = "utf8"
)

// Valid indicates whether the value is a known member of the SourceFileEncoding enum.
func (e SourceFileEncoding) Valid() bool {
	switch e {
	case SourceFileEncodingBase64:
		return true
	case SourceFileEncodingUTF8:
		return true
	default:
		return false
	}
}

// Defines values for SystemIntegrationCategory.
const (
	SystemIntegrationCategoryAI                SystemIntegrationCategory = "ai"
	SystemIntegrationCategoryAPIGateway        SystemIntegrationCategory = "api_gateway"
	SystemIntegrationCategoryCICD              SystemIntegrationCategory = "ci_cd"
	SystemIntegrationCategoryCloud             SystemIntegrationCategory = "cloud"
	SystemIntegrationCategoryCodeQuality       SystemIntegrationCategory = "code_quality"
	SystemIntegrationCategoryConfiguration     SystemIntegrationCategory = "configuration"
	SystemIntegrationCategoryIdentity          SystemIntegrationCategory = "identity"
	SystemIntegrationCategoryMessaging         SystemIntegrationCategory = "messaging"
	SystemIntegrationCategoryMonitoring        SystemIntegrationCategory = "monitoring"
	SystemIntegrationCategoryOther             SystemIntegrationCategory = "other"
	SystemIntegrationCategoryProjectManagement SystemIntegrationCategory = "project_management"
	SystemIntegrationCategorySourceControl     SystemIntegrationCategory = "source_control"
)

// Valid indicates whether the value is a known member of the SystemIntegrationCategory enum.
func (e SystemIntegrationCategory) Valid() bool {
	switch e {
	case SystemIntegrationCategoryAI:
		return true
	case SystemIntegrationCategoryAPIGateway:
		return true
	case SystemIntegrationCategoryCICD:
		return true
	case SystemIntegrationCategoryCloud:
		return true
	case SystemIntegrationCategoryCodeQuality:
		return true
	case SystemIntegrationCategoryConfiguration:
		return true
	case SystemIntegrationCategoryIdentity:
		return true
	case SystemIntegrationCategoryMessaging:
		return true
	case SystemIntegrationCategoryMonitoring:
		return true
	case SystemIntegrationCategoryOther:
		return true
	case SystemIntegrationCategoryProjectManagement:
		return true
	case SystemIntegrationCategorySourceControl:
		return true
	default:
		return false
	}
}

// Defines values for SystemIntegrationHealthStatus.
const (
	SystemIntegrationHealthStatusHealthy   SystemIntegrationHealthStatus = "healthy"
	SystemIntegrationHealthStatusUnhealthy SystemIntegrationHealthStatus = "unhealthy"
	SystemIntegrationHealthStatusUnknown   SystemIntegrationHealthStatus = "unknown"
)

// Valid indicates whether the value is a known member of the SystemIntegrationHealthStatus enum.
func (e SystemIntegrationHealthStatus) Valid() bool {
	switch e {
	case SystemIntegrationHealthStatusHealthy:
		return true
	case SystemIntegrationHealthStatusUnhealthy:
		return true
	case SystemIntegrationHealthStatusUnknown:
		return true
	default:
		return false
	}
}

// Defines values for SystemIntegrationTestStatus.
const (
	SystemIntegrationTestStatusFailed    SystemIntegrationTestStatus = "failed"
	SystemIntegrationTestStatusSucceeded SystemIntegrationTestStatus = "succeeded"
)

// Valid indicates whether the value is a known member of the SystemIntegrationTestStatus enum.
func (e SystemIntegrationTestStatus) Valid() bool {
	switch e {
	case SystemIntegrationTestStatusFailed:
		return true
	case SystemIntegrationTestStatusSucceeded:
		return true
	default:
		return false
	}
}

// Defines values for WorkbenchAgentStatusEventProviderKind.
const (
	WorkbenchAgentStatusEventProviderKindAnthropic        WorkbenchAgentStatusEventProviderKind = "anthropic"
	WorkbenchAgentStatusEventProviderKindAzureOpenai      WorkbenchAgentStatusEventProviderKind = "azure-openai"
	WorkbenchAgentStatusEventProviderKindCohere           WorkbenchAgentStatusEventProviderKind = "cohere"
	WorkbenchAgentStatusEventProviderKindDeepseek         WorkbenchAgentStatusEventProviderKind = "deepseek"
	WorkbenchAgentStatusEventProviderKindGemini           WorkbenchAgentStatusEventProviderKind = "gemini"
	WorkbenchAgentStatusEventProviderKindGeneral          WorkbenchAgentStatusEventProviderKind = "general"
	WorkbenchAgentStatusEventProviderKindHermes           WorkbenchAgentStatusEventProviderKind = "hermes"
	WorkbenchAgentStatusEventProviderKindInternal         WorkbenchAgentStatusEventProviderKind = "internal"
	WorkbenchAgentStatusEventProviderKindOpenai           WorkbenchAgentStatusEventProviderKind = "openai"
	WorkbenchAgentStatusEventProviderKindOpenaiCompatible WorkbenchAgentStatusEventProviderKind = "openai-compatible"
	WorkbenchAgentStatusEventProviderKindOpenclaw         WorkbenchAgentStatusEventProviderKind = "openclaw"
	WorkbenchAgentStatusEventProviderKindOpenrouter       WorkbenchAgentStatusEventProviderKind = "openrouter"
	WorkbenchAgentStatusEventProviderKindQwen             WorkbenchAgentStatusEventProviderKind = "qwen"
)

// Valid indicates whether the value is a known member of the WorkbenchAgentStatusEventProviderKind enum.
func (e WorkbenchAgentStatusEventProviderKind) Valid() bool {
	switch e {
	case WorkbenchAgentStatusEventProviderKindAnthropic:
		return true
	case WorkbenchAgentStatusEventProviderKindAzureOpenai:
		return true
	case WorkbenchAgentStatusEventProviderKindCohere:
		return true
	case WorkbenchAgentStatusEventProviderKindDeepseek:
		return true
	case WorkbenchAgentStatusEventProviderKindGemini:
		return true
	case WorkbenchAgentStatusEventProviderKindGeneral:
		return true
	case WorkbenchAgentStatusEventProviderKindHermes:
		return true
	case WorkbenchAgentStatusEventProviderKindInternal:
		return true
	case WorkbenchAgentStatusEventProviderKindOpenai:
		return true
	case WorkbenchAgentStatusEventProviderKindOpenaiCompatible:
		return true
	case WorkbenchAgentStatusEventProviderKindOpenclaw:
		return true
	case WorkbenchAgentStatusEventProviderKindOpenrouter:
		return true
	case WorkbenchAgentStatusEventProviderKindQwen:
		return true
	default:
		return false
	}
}

// Defines values for WorkbenchAgentStatusEventStatus.
const (
	WorkbenchAgentStatusEventStatusCancelled WorkbenchAgentStatusEventStatus = "cancelled"
	WorkbenchAgentStatusEventStatusFailed    WorkbenchAgentStatusEventStatus = "failed"
	WorkbenchAgentStatusEventStatusQueued    WorkbenchAgentStatusEventStatus = "queued"
	WorkbenchAgentStatusEventStatusRunning   WorkbenchAgentStatusEventStatus = "running"
	WorkbenchAgentStatusEventStatusSucceeded WorkbenchAgentStatusEventStatus = "succeeded"
)

// Valid indicates whether the value is a known member of the WorkbenchAgentStatusEventStatus enum.
func (e WorkbenchAgentStatusEventStatus) Valid() bool {
	switch e {
	case WorkbenchAgentStatusEventStatusCancelled:
		return true
	case WorkbenchAgentStatusEventStatusFailed:
		return true
	case WorkbenchAgentStatusEventStatusQueued:
		return true
	case WorkbenchAgentStatusEventStatusRunning:
		return true
	case WorkbenchAgentStatusEventStatusSucceeded:
		return true
	default:
		return false
	}
}

// Defines values for WorkbenchAgentStatusEventType.
const (
	WorkbenchAgentStatusEventTypeAgentStatus WorkbenchAgentStatusEventType = "agent.status"
)

// Valid indicates whether the value is a known member of the WorkbenchAgentStatusEventType enum.
func (e WorkbenchAgentStatusEventType) Valid() bool {
	switch e {
	case WorkbenchAgentStatusEventTypeAgentStatus:
		return true
	default:
		return false
	}
}

// Defines values for WorkbenchArtifactUpdatedEventType.
const (
	WorkbenchArtifactUpdatedEventTypeArtifactUpdated WorkbenchArtifactUpdatedEventType = "artifact.updated"
)

// Valid indicates whether the value is a known member of the WorkbenchArtifactUpdatedEventType enum.
func (e WorkbenchArtifactUpdatedEventType) Valid() bool {
	switch e {
	case WorkbenchArtifactUpdatedEventTypeArtifactUpdated:
		return true
	default:
		return false
	}
}

// Defines values for WorkbenchCardCommandEventType.
const (
	WorkbenchCardCommandEventTypeCardCommand WorkbenchCardCommandEventType = "card.command"
)

// Valid indicates whether the value is a known member of the WorkbenchCardCommandEventType enum.
func (e WorkbenchCardCommandEventType) Valid() bool {
	switch e {
	case WorkbenchCardCommandEventTypeCardCommand:
		return true
	default:
		return false
	}
}

// Defines values for WorkbenchErrorEventType.
const (
	WorkbenchErrorEventTypeError WorkbenchErrorEventType = "error"
)

// Valid indicates whether the value is a known member of the WorkbenchErrorEventType enum.
func (e WorkbenchErrorEventType) Valid() bool {
	switch e {
	case WorkbenchErrorEventTypeError:
		return true
	default:
		return false
	}
}

// Defines values for WorkbenchGlobalAssistantOpenRequestAction.
const (
	AnalyzePage      WorkbenchGlobalAssistantOpenRequestAction = "analyze-page"
	AnalyzeResource  WorkbenchGlobalAssistantOpenRequestAction = "analyze-resource"
	AnalyzeSelection WorkbenchGlobalAssistantOpenRequestAction = "analyze-selection"
	Open             WorkbenchGlobalAssistantOpenRequestAction = "open"
	OpenWorkbench    WorkbenchGlobalAssistantOpenRequestAction = "open-workbench"
)

// Valid indicates whether the value is a known member of the WorkbenchGlobalAssistantOpenRequestAction enum.
func (e WorkbenchGlobalAssistantOpenRequestAction) Valid() bool {
	switch e {
	case AnalyzePage:
		return true
	case AnalyzeResource:
		return true
	case AnalyzeSelection:
		return true
	case Open:
		return true
	case OpenWorkbench:
		return true
	default:
		return false
	}
}

// Defines values for WorkbenchLaunchContextSourceWorkbench.
const (
	WorkbenchLaunchContextSourceWorkbenchAI             WorkbenchLaunchContextSourceWorkbench = "ai"
	WorkbenchLaunchContextSourceWorkbenchCompute        WorkbenchLaunchContextSourceWorkbench = "compute"
	WorkbenchLaunchContextSourceWorkbenchDelivery       WorkbenchLaunchContextSourceWorkbench = "delivery"
	WorkbenchLaunchContextSourceWorkbenchDocker         WorkbenchLaunchContextSourceWorkbench = "docker"
	WorkbenchLaunchContextSourceWorkbenchMonitoring     WorkbenchLaunchContextSourceWorkbench = "monitoring"
	WorkbenchLaunchContextSourceWorkbenchPlatform       WorkbenchLaunchContextSourceWorkbench = "platform"
	WorkbenchLaunchContextSourceWorkbenchVirtualization WorkbenchLaunchContextSourceWorkbench = "virtualization"
)

// Valid indicates whether the value is a known member of the WorkbenchLaunchContextSourceWorkbench enum.
func (e WorkbenchLaunchContextSourceWorkbench) Valid() bool {
	switch e {
	case WorkbenchLaunchContextSourceWorkbenchAI:
		return true
	case WorkbenchLaunchContextSourceWorkbenchCompute:
		return true
	case WorkbenchLaunchContextSourceWorkbenchDelivery:
		return true
	case WorkbenchLaunchContextSourceWorkbenchDocker:
		return true
	case WorkbenchLaunchContextSourceWorkbenchMonitoring:
		return true
	case WorkbenchLaunchContextSourceWorkbenchPlatform:
		return true
	case WorkbenchLaunchContextSourceWorkbenchVirtualization:
		return true
	default:
		return false
	}
}

// Defines values for WorkbenchMessageDeltaEventRole.
const (
	WorkbenchMessageDeltaEventRoleAssistant WorkbenchMessageDeltaEventRole = "assistant"
)

// Valid indicates whether the value is a known member of the WorkbenchMessageDeltaEventRole enum.
func (e WorkbenchMessageDeltaEventRole) Valid() bool {
	switch e {
	case WorkbenchMessageDeltaEventRoleAssistant:
		return true
	default:
		return false
	}
}

// Defines values for WorkbenchMessageDeltaEventType.
const (
	WorkbenchMessageDeltaEventTypeMessageDelta WorkbenchMessageDeltaEventType = "message.delta"
)

// Valid indicates whether the value is a known member of the WorkbenchMessageDeltaEventType enum.
func (e WorkbenchMessageDeltaEventType) Valid() bool {
	switch e {
	case WorkbenchMessageDeltaEventTypeMessageDelta:
		return true
	default:
		return false
	}
}

// Defines values for WorkbenchMessageDoneEventRole.
const (
	WorkbenchMessageDoneEventRoleAssistant WorkbenchMessageDoneEventRole = "assistant"
)

// Valid indicates whether the value is a known member of the WorkbenchMessageDoneEventRole enum.
func (e WorkbenchMessageDoneEventRole) Valid() bool {
	switch e {
	case WorkbenchMessageDoneEventRoleAssistant:
		return true
	default:
		return false
	}
}

// Defines values for WorkbenchMessageDoneEventType.
const (
	WorkbenchMessageDoneEventTypeMessageDone WorkbenchMessageDoneEventType = "message.done"
)

// Valid indicates whether the value is a known member of the WorkbenchMessageDoneEventType enum.
func (e WorkbenchMessageDoneEventType) Valid() bool {
	switch e {
	case WorkbenchMessageDoneEventTypeMessageDone:
		return true
	default:
		return false
	}
}

// Defines values for WorkbenchSelectionContextKind.
const (
	WorkbenchSelectionContextKindError  WorkbenchSelectionContextKind = "error"
	WorkbenchSelectionContextKindEvent  WorkbenchSelectionContextKind = "event"
	WorkbenchSelectionContextKindLog    WorkbenchSelectionContextKind = "log"
	WorkbenchSelectionContextKindMetric WorkbenchSelectionContextKind = "metric"
	WorkbenchSelectionContextKindPlain  WorkbenchSelectionContextKind = "plain"
	WorkbenchSelectionContextKindYaml   WorkbenchSelectionContextKind = "yaml"
)

// Valid indicates whether the value is a known member of the WorkbenchSelectionContextKind enum.
func (e WorkbenchSelectionContextKind) Valid() bool {
	switch e {
	case WorkbenchSelectionContextKindError:
		return true
	case WorkbenchSelectionContextKindEvent:
		return true
	case WorkbenchSelectionContextKindLog:
		return true
	case WorkbenchSelectionContextKindMetric:
		return true
	case WorkbenchSelectionContextKindPlain:
		return true
	case WorkbenchSelectionContextKindYaml:
		return true
	default:
		return false
	}
}

// Defines values for WorkbenchSourceKind.
const (
	WorkbenchSourceKindAudit    WorkbenchSourceKind = "audit"
	WorkbenchSourceKindDelivery WorkbenchSourceKind = "delivery"
	WorkbenchSourceKindDocument WorkbenchSourceKind = "document"
	WorkbenchSourceKindEvent    WorkbenchSourceKind = "event"
	WorkbenchSourceKindLog      WorkbenchSourceKind = "log"
	WorkbenchSourceKindMetric   WorkbenchSourceKind = "metric"
	WorkbenchSourceKindTrace    WorkbenchSourceKind = "trace"
)

// Valid indicates whether the value is a known member of the WorkbenchSourceKind enum.
func (e WorkbenchSourceKind) Valid() bool {
	switch e {
	case WorkbenchSourceKindAudit:
		return true
	case WorkbenchSourceKindDelivery:
		return true
	case WorkbenchSourceKindDocument:
		return true
	case WorkbenchSourceKindEvent:
		return true
	case WorkbenchSourceKindLog:
		return true
	case WorkbenchSourceKindMetric:
		return true
	case WorkbenchSourceKindTrace:
		return true
	default:
		return false
	}
}

// Defines values for WorkbenchSourceUpdatedEventType.
const (
	WorkbenchSourceUpdatedEventTypeSourceUpdated WorkbenchSourceUpdatedEventType = "source.updated"
)

// Valid indicates whether the value is a known member of the WorkbenchSourceUpdatedEventType enum.
func (e WorkbenchSourceUpdatedEventType) Valid() bool {
	switch e {
	case WorkbenchSourceUpdatedEventTypeSourceUpdated:
		return true
	default:
		return false
	}
}

// Defines values for WorkbenchThinkingDeltaEventType.
const (
	WorkbenchThinkingDeltaEventTypeThinkingDelta WorkbenchThinkingDeltaEventType = "thinking.delta"
)

// Valid indicates whether the value is a known member of the WorkbenchThinkingDeltaEventType enum.
func (e WorkbenchThinkingDeltaEventType) Valid() bool {
	switch e {
	case WorkbenchThinkingDeltaEventTypeThinkingDelta:
		return true
	default:
		return false
	}
}

// Defines values for WorkbenchThinkingDoneEventType.
const (
	WorkbenchThinkingDoneEventTypeThinkingDone WorkbenchThinkingDoneEventType = "thinking.done"
)

// Valid indicates whether the value is a known member of the WorkbenchThinkingDoneEventType enum.
func (e WorkbenchThinkingDoneEventType) Valid() bool {
	switch e {
	case WorkbenchThinkingDoneEventTypeThinkingDone:
		return true
	default:
		return false
	}
}

// Defines values for WorkbenchToolCallStatus.
const (
	WorkbenchToolCallStatusError   WorkbenchToolCallStatus = "error"
	WorkbenchToolCallStatusPending WorkbenchToolCallStatus = "pending"
	WorkbenchToolCallStatusRunning WorkbenchToolCallStatus = "running"
	WorkbenchToolCallStatusSkipped WorkbenchToolCallStatus = "skipped"
	WorkbenchToolCallStatusSuccess WorkbenchToolCallStatus = "success"
)

// Valid indicates whether the value is a known member of the WorkbenchToolCallStatus enum.
func (e WorkbenchToolCallStatus) Valid() bool {
	switch e {
	case WorkbenchToolCallStatusError:
		return true
	case WorkbenchToolCallStatusPending:
		return true
	case WorkbenchToolCallStatusRunning:
		return true
	case WorkbenchToolCallStatusSkipped:
		return true
	case WorkbenchToolCallStatusSuccess:
		return true
	default:
		return false
	}
}

// Defines values for WorkbenchToolCompletedEventType.
const (
	WorkbenchToolCompletedEventTypeToolCompleted WorkbenchToolCompletedEventType = "tool.completed"
)

// Valid indicates whether the value is a known member of the WorkbenchToolCompletedEventType enum.
func (e WorkbenchToolCompletedEventType) Valid() bool {
	switch e {
	case WorkbenchToolCompletedEventTypeToolCompleted:
		return true
	default:
		return false
	}
}

// Defines values for WorkbenchToolDeltaEventType.
const (
	WorkbenchToolDeltaEventTypeToolDelta WorkbenchToolDeltaEventType = "tool.delta"
)

// Valid indicates whether the value is a known member of the WorkbenchToolDeltaEventType enum.
func (e WorkbenchToolDeltaEventType) Valid() bool {
	switch e {
	case WorkbenchToolDeltaEventTypeToolDelta:
		return true
	default:
		return false
	}
}

// Defines values for WorkbenchToolStartedEventType.
const (
	WorkbenchToolStartedEventTypeToolStarted WorkbenchToolStartedEventType = "tool.started"
)

// Valid indicates whether the value is a known member of the WorkbenchToolStartedEventType enum.
func (e WorkbenchToolStartedEventType) Valid() bool {
	switch e {
	case WorkbenchToolStartedEventTypeToolStarted:
		return true
	default:
		return false
	}
}

// Defines values for AgentProviderRolloutAction.
const (
	AgentProviderRolloutActionPause    AgentProviderRolloutAction = "pause"
	AgentProviderRolloutActionResume   AgentProviderRolloutAction = "resume"
	AgentProviderRolloutActionRollback AgentProviderRolloutAction = "rollback"
)

// Valid indicates whether the value is a known member of the AgentProviderRolloutAction enum.
func (e AgentProviderRolloutAction) Valid() bool {
	switch e {
	case AgentProviderRolloutActionPause:
		return true
	case AgentProviderRolloutActionResume:
		return true
	case AgentProviderRolloutActionRollback:
		return true
	default:
		return false
	}
}

// Defines values for OpenAICompatibleProvider.
const (
	OpenAICompatibleProviderAzureOpenai OpenAICompatibleProvider = "azure-openai"
	OpenAICompatibleProviderDeepseek    OpenAICompatibleProvider = "deepseek"
	OpenAICompatibleProviderOpenrouter  OpenAICompatibleProvider = "openrouter"
	OpenAICompatibleProviderQwen        OpenAICompatibleProvider = "qwen"
)

// Valid indicates whether the value is a known member of the OpenAICompatibleProvider enum.
func (e OpenAICompatibleProvider) Valid() bool {
	switch e {
	case OpenAICompatibleProviderAzureOpenai:
		return true
	case OpenAICompatibleProviderDeepseek:
		return true
	case OpenAICompatibleProviderOpenrouter:
		return true
	case OpenAICompatibleProviderQwen:
		return true
	default:
		return false
	}
}

// Defines values for SohaCacheModeHeader.
const (
	SohaCacheModeHeaderBypass   SohaCacheModeHeader = "bypass"
	SohaCacheModeHeaderDefault  SohaCacheModeHeader = "default"
	SohaCacheModeHeaderReadOnly SohaCacheModeHeader = "read-only"
	SohaCacheModeHeaderRefresh  SohaCacheModeHeader = "refresh"
)

// Valid indicates whether the value is a known member of the SohaCacheModeHeader enum.
func (e SohaCacheModeHeader) Valid() bool {
	switch e {
	case SohaCacheModeHeaderBypass:
		return true
	case SohaCacheModeHeaderDefault:
		return true
	case SohaCacheModeHeaderReadOnly:
		return true
	case SohaCacheModeHeaderRefresh:
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

// Defines values for CreateAIGatewayAnthropicMessageParamsXSohaCacheMode.
const (
	CreateAIGatewayAnthropicMessageParamsXSohaCacheModeBypass   CreateAIGatewayAnthropicMessageParamsXSohaCacheMode = "bypass"
	CreateAIGatewayAnthropicMessageParamsXSohaCacheModeDefault  CreateAIGatewayAnthropicMessageParamsXSohaCacheMode = "default"
	CreateAIGatewayAnthropicMessageParamsXSohaCacheModeReadOnly CreateAIGatewayAnthropicMessageParamsXSohaCacheMode = "read-only"
	CreateAIGatewayAnthropicMessageParamsXSohaCacheModeRefresh  CreateAIGatewayAnthropicMessageParamsXSohaCacheMode = "refresh"
)

// Valid indicates whether the value is a known member of the CreateAIGatewayAnthropicMessageParamsXSohaCacheMode enum.
func (e CreateAIGatewayAnthropicMessageParamsXSohaCacheMode) Valid() bool {
	switch e {
	case CreateAIGatewayAnthropicMessageParamsXSohaCacheModeBypass:
		return true
	case CreateAIGatewayAnthropicMessageParamsXSohaCacheModeDefault:
		return true
	case CreateAIGatewayAnthropicMessageParamsXSohaCacheModeReadOnly:
		return true
	case CreateAIGatewayAnthropicMessageParamsXSohaCacheModeRefresh:
		return true
	default:
		return false
	}
}

// Defines values for CreateAIGatewayCohereRerankParamsXSohaCacheMode.
const (
	CreateAIGatewayCohereRerankParamsXSohaCacheModeBypass   CreateAIGatewayCohereRerankParamsXSohaCacheMode = "bypass"
	CreateAIGatewayCohereRerankParamsXSohaCacheModeDefault  CreateAIGatewayCohereRerankParamsXSohaCacheMode = "default"
	CreateAIGatewayCohereRerankParamsXSohaCacheModeReadOnly CreateAIGatewayCohereRerankParamsXSohaCacheMode = "read-only"
	CreateAIGatewayCohereRerankParamsXSohaCacheModeRefresh  CreateAIGatewayCohereRerankParamsXSohaCacheMode = "refresh"
)

// Valid indicates whether the value is a known member of the CreateAIGatewayCohereRerankParamsXSohaCacheMode enum.
func (e CreateAIGatewayCohereRerankParamsXSohaCacheMode) Valid() bool {
	switch e {
	case CreateAIGatewayCohereRerankParamsXSohaCacheModeBypass:
		return true
	case CreateAIGatewayCohereRerankParamsXSohaCacheModeDefault:
		return true
	case CreateAIGatewayCohereRerankParamsXSohaCacheModeReadOnly:
		return true
	case CreateAIGatewayCohereRerankParamsXSohaCacheModeRefresh:
		return true
	default:
		return false
	}
}

// Defines values for CreateAIGatewayGeminiInteractionsParamsXSohaCacheMode.
const (
	CreateAIGatewayGeminiInteractionsParamsXSohaCacheModeBypass   CreateAIGatewayGeminiInteractionsParamsXSohaCacheMode = "bypass"
	CreateAIGatewayGeminiInteractionsParamsXSohaCacheModeDefault  CreateAIGatewayGeminiInteractionsParamsXSohaCacheMode = "default"
	CreateAIGatewayGeminiInteractionsParamsXSohaCacheModeReadOnly CreateAIGatewayGeminiInteractionsParamsXSohaCacheMode = "read-only"
	CreateAIGatewayGeminiInteractionsParamsXSohaCacheModeRefresh  CreateAIGatewayGeminiInteractionsParamsXSohaCacheMode = "refresh"
)

// Valid indicates whether the value is a known member of the CreateAIGatewayGeminiInteractionsParamsXSohaCacheMode enum.
func (e CreateAIGatewayGeminiInteractionsParamsXSohaCacheMode) Valid() bool {
	switch e {
	case CreateAIGatewayGeminiInteractionsParamsXSohaCacheModeBypass:
		return true
	case CreateAIGatewayGeminiInteractionsParamsXSohaCacheModeDefault:
		return true
	case CreateAIGatewayGeminiInteractionsParamsXSohaCacheModeReadOnly:
		return true
	case CreateAIGatewayGeminiInteractionsParamsXSohaCacheModeRefresh:
		return true
	default:
		return false
	}
}

// Defines values for CreateAIGatewayGeminiGenerateContentParamsXSohaCacheMode.
const (
	CreateAIGatewayGeminiGenerateContentParamsXSohaCacheModeBypass   CreateAIGatewayGeminiGenerateContentParamsXSohaCacheMode = "bypass"
	CreateAIGatewayGeminiGenerateContentParamsXSohaCacheModeDefault  CreateAIGatewayGeminiGenerateContentParamsXSohaCacheMode = "default"
	CreateAIGatewayGeminiGenerateContentParamsXSohaCacheModeReadOnly CreateAIGatewayGeminiGenerateContentParamsXSohaCacheMode = "read-only"
	CreateAIGatewayGeminiGenerateContentParamsXSohaCacheModeRefresh  CreateAIGatewayGeminiGenerateContentParamsXSohaCacheMode = "refresh"
)

// Valid indicates whether the value is a known member of the CreateAIGatewayGeminiGenerateContentParamsXSohaCacheMode enum.
func (e CreateAIGatewayGeminiGenerateContentParamsXSohaCacheMode) Valid() bool {
	switch e {
	case CreateAIGatewayGeminiGenerateContentParamsXSohaCacheModeBypass:
		return true
	case CreateAIGatewayGeminiGenerateContentParamsXSohaCacheModeDefault:
		return true
	case CreateAIGatewayGeminiGenerateContentParamsXSohaCacheModeReadOnly:
		return true
	case CreateAIGatewayGeminiGenerateContentParamsXSohaCacheModeRefresh:
		return true
	default:
		return false
	}
}

// Defines values for CreateAIGatewayGeminiStreamGenerateContentParamsXSohaCacheMode.
const (
	CreateAIGatewayGeminiStreamGenerateContentParamsXSohaCacheModeBypass   CreateAIGatewayGeminiStreamGenerateContentParamsXSohaCacheMode = "bypass"
	CreateAIGatewayGeminiStreamGenerateContentParamsXSohaCacheModeDefault  CreateAIGatewayGeminiStreamGenerateContentParamsXSohaCacheMode = "default"
	CreateAIGatewayGeminiStreamGenerateContentParamsXSohaCacheModeReadOnly CreateAIGatewayGeminiStreamGenerateContentParamsXSohaCacheMode = "read-only"
	CreateAIGatewayGeminiStreamGenerateContentParamsXSohaCacheModeRefresh  CreateAIGatewayGeminiStreamGenerateContentParamsXSohaCacheMode = "refresh"
)

// Valid indicates whether the value is a known member of the CreateAIGatewayGeminiStreamGenerateContentParamsXSohaCacheMode enum.
func (e CreateAIGatewayGeminiStreamGenerateContentParamsXSohaCacheMode) Valid() bool {
	switch e {
	case CreateAIGatewayGeminiStreamGenerateContentParamsXSohaCacheModeBypass:
		return true
	case CreateAIGatewayGeminiStreamGenerateContentParamsXSohaCacheModeDefault:
		return true
	case CreateAIGatewayGeminiStreamGenerateContentParamsXSohaCacheModeReadOnly:
		return true
	case CreateAIGatewayGeminiStreamGenerateContentParamsXSohaCacheModeRefresh:
		return true
	default:
		return false
	}
}

// Defines values for CreateAIGatewayOpenAIAudioSpeechParamsXSohaCacheMode.
const (
	CreateAIGatewayOpenAIAudioSpeechParamsXSohaCacheModeBypass   CreateAIGatewayOpenAIAudioSpeechParamsXSohaCacheMode = "bypass"
	CreateAIGatewayOpenAIAudioSpeechParamsXSohaCacheModeDefault  CreateAIGatewayOpenAIAudioSpeechParamsXSohaCacheMode = "default"
	CreateAIGatewayOpenAIAudioSpeechParamsXSohaCacheModeReadOnly CreateAIGatewayOpenAIAudioSpeechParamsXSohaCacheMode = "read-only"
	CreateAIGatewayOpenAIAudioSpeechParamsXSohaCacheModeRefresh  CreateAIGatewayOpenAIAudioSpeechParamsXSohaCacheMode = "refresh"
)

// Valid indicates whether the value is a known member of the CreateAIGatewayOpenAIAudioSpeechParamsXSohaCacheMode enum.
func (e CreateAIGatewayOpenAIAudioSpeechParamsXSohaCacheMode) Valid() bool {
	switch e {
	case CreateAIGatewayOpenAIAudioSpeechParamsXSohaCacheModeBypass:
		return true
	case CreateAIGatewayOpenAIAudioSpeechParamsXSohaCacheModeDefault:
		return true
	case CreateAIGatewayOpenAIAudioSpeechParamsXSohaCacheModeReadOnly:
		return true
	case CreateAIGatewayOpenAIAudioSpeechParamsXSohaCacheModeRefresh:
		return true
	default:
		return false
	}
}

// Defines values for CreateAIGatewayOpenAIAudioTranscriptionParamsXSohaCacheMode.
const (
	CreateAIGatewayOpenAIAudioTranscriptionParamsXSohaCacheModeBypass   CreateAIGatewayOpenAIAudioTranscriptionParamsXSohaCacheMode = "bypass"
	CreateAIGatewayOpenAIAudioTranscriptionParamsXSohaCacheModeDefault  CreateAIGatewayOpenAIAudioTranscriptionParamsXSohaCacheMode = "default"
	CreateAIGatewayOpenAIAudioTranscriptionParamsXSohaCacheModeReadOnly CreateAIGatewayOpenAIAudioTranscriptionParamsXSohaCacheMode = "read-only"
	CreateAIGatewayOpenAIAudioTranscriptionParamsXSohaCacheModeRefresh  CreateAIGatewayOpenAIAudioTranscriptionParamsXSohaCacheMode = "refresh"
)

// Valid indicates whether the value is a known member of the CreateAIGatewayOpenAIAudioTranscriptionParamsXSohaCacheMode enum.
func (e CreateAIGatewayOpenAIAudioTranscriptionParamsXSohaCacheMode) Valid() bool {
	switch e {
	case CreateAIGatewayOpenAIAudioTranscriptionParamsXSohaCacheModeBypass:
		return true
	case CreateAIGatewayOpenAIAudioTranscriptionParamsXSohaCacheModeDefault:
		return true
	case CreateAIGatewayOpenAIAudioTranscriptionParamsXSohaCacheModeReadOnly:
		return true
	case CreateAIGatewayOpenAIAudioTranscriptionParamsXSohaCacheModeRefresh:
		return true
	default:
		return false
	}
}

// Defines values for CreateAIGatewayOpenAIAudioTranslationParamsXSohaCacheMode.
const (
	CreateAIGatewayOpenAIAudioTranslationParamsXSohaCacheModeBypass   CreateAIGatewayOpenAIAudioTranslationParamsXSohaCacheMode = "bypass"
	CreateAIGatewayOpenAIAudioTranslationParamsXSohaCacheModeDefault  CreateAIGatewayOpenAIAudioTranslationParamsXSohaCacheMode = "default"
	CreateAIGatewayOpenAIAudioTranslationParamsXSohaCacheModeReadOnly CreateAIGatewayOpenAIAudioTranslationParamsXSohaCacheMode = "read-only"
	CreateAIGatewayOpenAIAudioTranslationParamsXSohaCacheModeRefresh  CreateAIGatewayOpenAIAudioTranslationParamsXSohaCacheMode = "refresh"
)

// Valid indicates whether the value is a known member of the CreateAIGatewayOpenAIAudioTranslationParamsXSohaCacheMode enum.
func (e CreateAIGatewayOpenAIAudioTranslationParamsXSohaCacheMode) Valid() bool {
	switch e {
	case CreateAIGatewayOpenAIAudioTranslationParamsXSohaCacheModeBypass:
		return true
	case CreateAIGatewayOpenAIAudioTranslationParamsXSohaCacheModeDefault:
		return true
	case CreateAIGatewayOpenAIAudioTranslationParamsXSohaCacheModeReadOnly:
		return true
	case CreateAIGatewayOpenAIAudioTranslationParamsXSohaCacheModeRefresh:
		return true
	default:
		return false
	}
}

// Defines values for CreateAIGatewayOpenAIChatCompletionParamsXSohaCacheMode.
const (
	CreateAIGatewayOpenAIChatCompletionParamsXSohaCacheModeBypass   CreateAIGatewayOpenAIChatCompletionParamsXSohaCacheMode = "bypass"
	CreateAIGatewayOpenAIChatCompletionParamsXSohaCacheModeDefault  CreateAIGatewayOpenAIChatCompletionParamsXSohaCacheMode = "default"
	CreateAIGatewayOpenAIChatCompletionParamsXSohaCacheModeReadOnly CreateAIGatewayOpenAIChatCompletionParamsXSohaCacheMode = "read-only"
	CreateAIGatewayOpenAIChatCompletionParamsXSohaCacheModeRefresh  CreateAIGatewayOpenAIChatCompletionParamsXSohaCacheMode = "refresh"
)

// Valid indicates whether the value is a known member of the CreateAIGatewayOpenAIChatCompletionParamsXSohaCacheMode enum.
func (e CreateAIGatewayOpenAIChatCompletionParamsXSohaCacheMode) Valid() bool {
	switch e {
	case CreateAIGatewayOpenAIChatCompletionParamsXSohaCacheModeBypass:
		return true
	case CreateAIGatewayOpenAIChatCompletionParamsXSohaCacheModeDefault:
		return true
	case CreateAIGatewayOpenAIChatCompletionParamsXSohaCacheModeReadOnly:
		return true
	case CreateAIGatewayOpenAIChatCompletionParamsXSohaCacheModeRefresh:
		return true
	default:
		return false
	}
}

// Defines values for CreateAIGatewayOpenAIEmbeddingParamsXSohaCacheMode.
const (
	CreateAIGatewayOpenAIEmbeddingParamsXSohaCacheModeBypass   CreateAIGatewayOpenAIEmbeddingParamsXSohaCacheMode = "bypass"
	CreateAIGatewayOpenAIEmbeddingParamsXSohaCacheModeDefault  CreateAIGatewayOpenAIEmbeddingParamsXSohaCacheMode = "default"
	CreateAIGatewayOpenAIEmbeddingParamsXSohaCacheModeReadOnly CreateAIGatewayOpenAIEmbeddingParamsXSohaCacheMode = "read-only"
	CreateAIGatewayOpenAIEmbeddingParamsXSohaCacheModeRefresh  CreateAIGatewayOpenAIEmbeddingParamsXSohaCacheMode = "refresh"
)

// Valid indicates whether the value is a known member of the CreateAIGatewayOpenAIEmbeddingParamsXSohaCacheMode enum.
func (e CreateAIGatewayOpenAIEmbeddingParamsXSohaCacheMode) Valid() bool {
	switch e {
	case CreateAIGatewayOpenAIEmbeddingParamsXSohaCacheModeBypass:
		return true
	case CreateAIGatewayOpenAIEmbeddingParamsXSohaCacheModeDefault:
		return true
	case CreateAIGatewayOpenAIEmbeddingParamsXSohaCacheModeReadOnly:
		return true
	case CreateAIGatewayOpenAIEmbeddingParamsXSohaCacheModeRefresh:
		return true
	default:
		return false
	}
}

// Defines values for CreateAIGatewayOpenAIImageEditParamsXSohaCacheMode.
const (
	CreateAIGatewayOpenAIImageEditParamsXSohaCacheModeBypass   CreateAIGatewayOpenAIImageEditParamsXSohaCacheMode = "bypass"
	CreateAIGatewayOpenAIImageEditParamsXSohaCacheModeDefault  CreateAIGatewayOpenAIImageEditParamsXSohaCacheMode = "default"
	CreateAIGatewayOpenAIImageEditParamsXSohaCacheModeReadOnly CreateAIGatewayOpenAIImageEditParamsXSohaCacheMode = "read-only"
	CreateAIGatewayOpenAIImageEditParamsXSohaCacheModeRefresh  CreateAIGatewayOpenAIImageEditParamsXSohaCacheMode = "refresh"
)

// Valid indicates whether the value is a known member of the CreateAIGatewayOpenAIImageEditParamsXSohaCacheMode enum.
func (e CreateAIGatewayOpenAIImageEditParamsXSohaCacheMode) Valid() bool {
	switch e {
	case CreateAIGatewayOpenAIImageEditParamsXSohaCacheModeBypass:
		return true
	case CreateAIGatewayOpenAIImageEditParamsXSohaCacheModeDefault:
		return true
	case CreateAIGatewayOpenAIImageEditParamsXSohaCacheModeReadOnly:
		return true
	case CreateAIGatewayOpenAIImageEditParamsXSohaCacheModeRefresh:
		return true
	default:
		return false
	}
}

// Defines values for CreateAIGatewayOpenAIImageGenerationParamsXSohaCacheMode.
const (
	CreateAIGatewayOpenAIImageGenerationParamsXSohaCacheModeBypass   CreateAIGatewayOpenAIImageGenerationParamsXSohaCacheMode = "bypass"
	CreateAIGatewayOpenAIImageGenerationParamsXSohaCacheModeDefault  CreateAIGatewayOpenAIImageGenerationParamsXSohaCacheMode = "default"
	CreateAIGatewayOpenAIImageGenerationParamsXSohaCacheModeReadOnly CreateAIGatewayOpenAIImageGenerationParamsXSohaCacheMode = "read-only"
	CreateAIGatewayOpenAIImageGenerationParamsXSohaCacheModeRefresh  CreateAIGatewayOpenAIImageGenerationParamsXSohaCacheMode = "refresh"
)

// Valid indicates whether the value is a known member of the CreateAIGatewayOpenAIImageGenerationParamsXSohaCacheMode enum.
func (e CreateAIGatewayOpenAIImageGenerationParamsXSohaCacheMode) Valid() bool {
	switch e {
	case CreateAIGatewayOpenAIImageGenerationParamsXSohaCacheModeBypass:
		return true
	case CreateAIGatewayOpenAIImageGenerationParamsXSohaCacheModeDefault:
		return true
	case CreateAIGatewayOpenAIImageGenerationParamsXSohaCacheModeReadOnly:
		return true
	case CreateAIGatewayOpenAIImageGenerationParamsXSohaCacheModeRefresh:
		return true
	default:
		return false
	}
}

// Defines values for CreateAIGatewayOpenAIImageVariationParamsXSohaCacheMode.
const (
	CreateAIGatewayOpenAIImageVariationParamsXSohaCacheModeBypass   CreateAIGatewayOpenAIImageVariationParamsXSohaCacheMode = "bypass"
	CreateAIGatewayOpenAIImageVariationParamsXSohaCacheModeDefault  CreateAIGatewayOpenAIImageVariationParamsXSohaCacheMode = "default"
	CreateAIGatewayOpenAIImageVariationParamsXSohaCacheModeReadOnly CreateAIGatewayOpenAIImageVariationParamsXSohaCacheMode = "read-only"
	CreateAIGatewayOpenAIImageVariationParamsXSohaCacheModeRefresh  CreateAIGatewayOpenAIImageVariationParamsXSohaCacheMode = "refresh"
)

// Valid indicates whether the value is a known member of the CreateAIGatewayOpenAIImageVariationParamsXSohaCacheMode enum.
func (e CreateAIGatewayOpenAIImageVariationParamsXSohaCacheMode) Valid() bool {
	switch e {
	case CreateAIGatewayOpenAIImageVariationParamsXSohaCacheModeBypass:
		return true
	case CreateAIGatewayOpenAIImageVariationParamsXSohaCacheModeDefault:
		return true
	case CreateAIGatewayOpenAIImageVariationParamsXSohaCacheModeReadOnly:
		return true
	case CreateAIGatewayOpenAIImageVariationParamsXSohaCacheModeRefresh:
		return true
	default:
		return false
	}
}

// Defines values for CreateAIGatewayOpenAIResponseParamsXSohaCacheMode.
const (
	CreateAIGatewayOpenAIResponseParamsXSohaCacheModeBypass   CreateAIGatewayOpenAIResponseParamsXSohaCacheMode = "bypass"
	CreateAIGatewayOpenAIResponseParamsXSohaCacheModeDefault  CreateAIGatewayOpenAIResponseParamsXSohaCacheMode = "default"
	CreateAIGatewayOpenAIResponseParamsXSohaCacheModeReadOnly CreateAIGatewayOpenAIResponseParamsXSohaCacheMode = "read-only"
	CreateAIGatewayOpenAIResponseParamsXSohaCacheModeRefresh  CreateAIGatewayOpenAIResponseParamsXSohaCacheMode = "refresh"
)

// Valid indicates whether the value is a known member of the CreateAIGatewayOpenAIResponseParamsXSohaCacheMode enum.
func (e CreateAIGatewayOpenAIResponseParamsXSohaCacheMode) Valid() bool {
	switch e {
	case CreateAIGatewayOpenAIResponseParamsXSohaCacheModeBypass:
		return true
	case CreateAIGatewayOpenAIResponseParamsXSohaCacheModeDefault:
		return true
	case CreateAIGatewayOpenAIResponseParamsXSohaCacheModeReadOnly:
		return true
	case CreateAIGatewayOpenAIResponseParamsXSohaCacheModeRefresh:
		return true
	default:
		return false
	}
}

// Defines values for CreateAIGatewayOpenAICompatibleProviderAudioSpeechParamsXSohaCacheMode.
const (
	CreateAIGatewayOpenAICompatibleProviderAudioSpeechParamsXSohaCacheModeBypass   CreateAIGatewayOpenAICompatibleProviderAudioSpeechParamsXSohaCacheMode = "bypass"
	CreateAIGatewayOpenAICompatibleProviderAudioSpeechParamsXSohaCacheModeDefault  CreateAIGatewayOpenAICompatibleProviderAudioSpeechParamsXSohaCacheMode = "default"
	CreateAIGatewayOpenAICompatibleProviderAudioSpeechParamsXSohaCacheModeReadOnly CreateAIGatewayOpenAICompatibleProviderAudioSpeechParamsXSohaCacheMode = "read-only"
	CreateAIGatewayOpenAICompatibleProviderAudioSpeechParamsXSohaCacheModeRefresh  CreateAIGatewayOpenAICompatibleProviderAudioSpeechParamsXSohaCacheMode = "refresh"
)

// Valid indicates whether the value is a known member of the CreateAIGatewayOpenAICompatibleProviderAudioSpeechParamsXSohaCacheMode enum.
func (e CreateAIGatewayOpenAICompatibleProviderAudioSpeechParamsXSohaCacheMode) Valid() bool {
	switch e {
	case CreateAIGatewayOpenAICompatibleProviderAudioSpeechParamsXSohaCacheModeBypass:
		return true
	case CreateAIGatewayOpenAICompatibleProviderAudioSpeechParamsXSohaCacheModeDefault:
		return true
	case CreateAIGatewayOpenAICompatibleProviderAudioSpeechParamsXSohaCacheModeReadOnly:
		return true
	case CreateAIGatewayOpenAICompatibleProviderAudioSpeechParamsXSohaCacheModeRefresh:
		return true
	default:
		return false
	}
}

// Defines values for CreateAIGatewayOpenAICompatibleProviderAudioSpeechParamsOpenaiCompatibleProvider.
const (
	CreateAIGatewayOpenAICompatibleProviderAudioSpeechParamsOpenaiCompatibleProviderAzureOpenai CreateAIGatewayOpenAICompatibleProviderAudioSpeechParamsOpenaiCompatibleProvider = "azure-openai"
	CreateAIGatewayOpenAICompatibleProviderAudioSpeechParamsOpenaiCompatibleProviderDeepseek    CreateAIGatewayOpenAICompatibleProviderAudioSpeechParamsOpenaiCompatibleProvider = "deepseek"
	CreateAIGatewayOpenAICompatibleProviderAudioSpeechParamsOpenaiCompatibleProviderOpenrouter  CreateAIGatewayOpenAICompatibleProviderAudioSpeechParamsOpenaiCompatibleProvider = "openrouter"
	CreateAIGatewayOpenAICompatibleProviderAudioSpeechParamsOpenaiCompatibleProviderQwen        CreateAIGatewayOpenAICompatibleProviderAudioSpeechParamsOpenaiCompatibleProvider = "qwen"
)

// Valid indicates whether the value is a known member of the CreateAIGatewayOpenAICompatibleProviderAudioSpeechParamsOpenaiCompatibleProvider enum.
func (e CreateAIGatewayOpenAICompatibleProviderAudioSpeechParamsOpenaiCompatibleProvider) Valid() bool {
	switch e {
	case CreateAIGatewayOpenAICompatibleProviderAudioSpeechParamsOpenaiCompatibleProviderAzureOpenai:
		return true
	case CreateAIGatewayOpenAICompatibleProviderAudioSpeechParamsOpenaiCompatibleProviderDeepseek:
		return true
	case CreateAIGatewayOpenAICompatibleProviderAudioSpeechParamsOpenaiCompatibleProviderOpenrouter:
		return true
	case CreateAIGatewayOpenAICompatibleProviderAudioSpeechParamsOpenaiCompatibleProviderQwen:
		return true
	default:
		return false
	}
}

// Defines values for CreateAIGatewayOpenAICompatibleProviderAudioTranscriptionParamsXSohaCacheMode.
const (
	CreateAIGatewayOpenAICompatibleProviderAudioTranscriptionParamsXSohaCacheModeBypass   CreateAIGatewayOpenAICompatibleProviderAudioTranscriptionParamsXSohaCacheMode = "bypass"
	CreateAIGatewayOpenAICompatibleProviderAudioTranscriptionParamsXSohaCacheModeDefault  CreateAIGatewayOpenAICompatibleProviderAudioTranscriptionParamsXSohaCacheMode = "default"
	CreateAIGatewayOpenAICompatibleProviderAudioTranscriptionParamsXSohaCacheModeReadOnly CreateAIGatewayOpenAICompatibleProviderAudioTranscriptionParamsXSohaCacheMode = "read-only"
	CreateAIGatewayOpenAICompatibleProviderAudioTranscriptionParamsXSohaCacheModeRefresh  CreateAIGatewayOpenAICompatibleProviderAudioTranscriptionParamsXSohaCacheMode = "refresh"
)

// Valid indicates whether the value is a known member of the CreateAIGatewayOpenAICompatibleProviderAudioTranscriptionParamsXSohaCacheMode enum.
func (e CreateAIGatewayOpenAICompatibleProviderAudioTranscriptionParamsXSohaCacheMode) Valid() bool {
	switch e {
	case CreateAIGatewayOpenAICompatibleProviderAudioTranscriptionParamsXSohaCacheModeBypass:
		return true
	case CreateAIGatewayOpenAICompatibleProviderAudioTranscriptionParamsXSohaCacheModeDefault:
		return true
	case CreateAIGatewayOpenAICompatibleProviderAudioTranscriptionParamsXSohaCacheModeReadOnly:
		return true
	case CreateAIGatewayOpenAICompatibleProviderAudioTranscriptionParamsXSohaCacheModeRefresh:
		return true
	default:
		return false
	}
}

// Defines values for CreateAIGatewayOpenAICompatibleProviderAudioTranscriptionParamsOpenaiCompatibleProvider.
const (
	CreateAIGatewayOpenAICompatibleProviderAudioTranscriptionParamsOpenaiCompatibleProviderAzureOpenai CreateAIGatewayOpenAICompatibleProviderAudioTranscriptionParamsOpenaiCompatibleProvider = "azure-openai"
	CreateAIGatewayOpenAICompatibleProviderAudioTranscriptionParamsOpenaiCompatibleProviderDeepseek    CreateAIGatewayOpenAICompatibleProviderAudioTranscriptionParamsOpenaiCompatibleProvider = "deepseek"
	CreateAIGatewayOpenAICompatibleProviderAudioTranscriptionParamsOpenaiCompatibleProviderOpenrouter  CreateAIGatewayOpenAICompatibleProviderAudioTranscriptionParamsOpenaiCompatibleProvider = "openrouter"
	CreateAIGatewayOpenAICompatibleProviderAudioTranscriptionParamsOpenaiCompatibleProviderQwen        CreateAIGatewayOpenAICompatibleProviderAudioTranscriptionParamsOpenaiCompatibleProvider = "qwen"
)

// Valid indicates whether the value is a known member of the CreateAIGatewayOpenAICompatibleProviderAudioTranscriptionParamsOpenaiCompatibleProvider enum.
func (e CreateAIGatewayOpenAICompatibleProviderAudioTranscriptionParamsOpenaiCompatibleProvider) Valid() bool {
	switch e {
	case CreateAIGatewayOpenAICompatibleProviderAudioTranscriptionParamsOpenaiCompatibleProviderAzureOpenai:
		return true
	case CreateAIGatewayOpenAICompatibleProviderAudioTranscriptionParamsOpenaiCompatibleProviderDeepseek:
		return true
	case CreateAIGatewayOpenAICompatibleProviderAudioTranscriptionParamsOpenaiCompatibleProviderOpenrouter:
		return true
	case CreateAIGatewayOpenAICompatibleProviderAudioTranscriptionParamsOpenaiCompatibleProviderQwen:
		return true
	default:
		return false
	}
}

// Defines values for CreateAIGatewayOpenAICompatibleProviderAudioTranslationParamsXSohaCacheMode.
const (
	CreateAIGatewayOpenAICompatibleProviderAudioTranslationParamsXSohaCacheModeBypass   CreateAIGatewayOpenAICompatibleProviderAudioTranslationParamsXSohaCacheMode = "bypass"
	CreateAIGatewayOpenAICompatibleProviderAudioTranslationParamsXSohaCacheModeDefault  CreateAIGatewayOpenAICompatibleProviderAudioTranslationParamsXSohaCacheMode = "default"
	CreateAIGatewayOpenAICompatibleProviderAudioTranslationParamsXSohaCacheModeReadOnly CreateAIGatewayOpenAICompatibleProviderAudioTranslationParamsXSohaCacheMode = "read-only"
	CreateAIGatewayOpenAICompatibleProviderAudioTranslationParamsXSohaCacheModeRefresh  CreateAIGatewayOpenAICompatibleProviderAudioTranslationParamsXSohaCacheMode = "refresh"
)

// Valid indicates whether the value is a known member of the CreateAIGatewayOpenAICompatibleProviderAudioTranslationParamsXSohaCacheMode enum.
func (e CreateAIGatewayOpenAICompatibleProviderAudioTranslationParamsXSohaCacheMode) Valid() bool {
	switch e {
	case CreateAIGatewayOpenAICompatibleProviderAudioTranslationParamsXSohaCacheModeBypass:
		return true
	case CreateAIGatewayOpenAICompatibleProviderAudioTranslationParamsXSohaCacheModeDefault:
		return true
	case CreateAIGatewayOpenAICompatibleProviderAudioTranslationParamsXSohaCacheModeReadOnly:
		return true
	case CreateAIGatewayOpenAICompatibleProviderAudioTranslationParamsXSohaCacheModeRefresh:
		return true
	default:
		return false
	}
}

// Defines values for CreateAIGatewayOpenAICompatibleProviderAudioTranslationParamsOpenaiCompatibleProvider.
const (
	CreateAIGatewayOpenAICompatibleProviderAudioTranslationParamsOpenaiCompatibleProviderAzureOpenai CreateAIGatewayOpenAICompatibleProviderAudioTranslationParamsOpenaiCompatibleProvider = "azure-openai"
	CreateAIGatewayOpenAICompatibleProviderAudioTranslationParamsOpenaiCompatibleProviderDeepseek    CreateAIGatewayOpenAICompatibleProviderAudioTranslationParamsOpenaiCompatibleProvider = "deepseek"
	CreateAIGatewayOpenAICompatibleProviderAudioTranslationParamsOpenaiCompatibleProviderOpenrouter  CreateAIGatewayOpenAICompatibleProviderAudioTranslationParamsOpenaiCompatibleProvider = "openrouter"
	CreateAIGatewayOpenAICompatibleProviderAudioTranslationParamsOpenaiCompatibleProviderQwen        CreateAIGatewayOpenAICompatibleProviderAudioTranslationParamsOpenaiCompatibleProvider = "qwen"
)

// Valid indicates whether the value is a known member of the CreateAIGatewayOpenAICompatibleProviderAudioTranslationParamsOpenaiCompatibleProvider enum.
func (e CreateAIGatewayOpenAICompatibleProviderAudioTranslationParamsOpenaiCompatibleProvider) Valid() bool {
	switch e {
	case CreateAIGatewayOpenAICompatibleProviderAudioTranslationParamsOpenaiCompatibleProviderAzureOpenai:
		return true
	case CreateAIGatewayOpenAICompatibleProviderAudioTranslationParamsOpenaiCompatibleProviderDeepseek:
		return true
	case CreateAIGatewayOpenAICompatibleProviderAudioTranslationParamsOpenaiCompatibleProviderOpenrouter:
		return true
	case CreateAIGatewayOpenAICompatibleProviderAudioTranslationParamsOpenaiCompatibleProviderQwen:
		return true
	default:
		return false
	}
}

// Defines values for CreateAIGatewayOpenAICompatibleProviderChatCompletionParamsXSohaCacheMode.
const (
	CreateAIGatewayOpenAICompatibleProviderChatCompletionParamsXSohaCacheModeBypass   CreateAIGatewayOpenAICompatibleProviderChatCompletionParamsXSohaCacheMode = "bypass"
	CreateAIGatewayOpenAICompatibleProviderChatCompletionParamsXSohaCacheModeDefault  CreateAIGatewayOpenAICompatibleProviderChatCompletionParamsXSohaCacheMode = "default"
	CreateAIGatewayOpenAICompatibleProviderChatCompletionParamsXSohaCacheModeReadOnly CreateAIGatewayOpenAICompatibleProviderChatCompletionParamsXSohaCacheMode = "read-only"
	CreateAIGatewayOpenAICompatibleProviderChatCompletionParamsXSohaCacheModeRefresh  CreateAIGatewayOpenAICompatibleProviderChatCompletionParamsXSohaCacheMode = "refresh"
)

// Valid indicates whether the value is a known member of the CreateAIGatewayOpenAICompatibleProviderChatCompletionParamsXSohaCacheMode enum.
func (e CreateAIGatewayOpenAICompatibleProviderChatCompletionParamsXSohaCacheMode) Valid() bool {
	switch e {
	case CreateAIGatewayOpenAICompatibleProviderChatCompletionParamsXSohaCacheModeBypass:
		return true
	case CreateAIGatewayOpenAICompatibleProviderChatCompletionParamsXSohaCacheModeDefault:
		return true
	case CreateAIGatewayOpenAICompatibleProviderChatCompletionParamsXSohaCacheModeReadOnly:
		return true
	case CreateAIGatewayOpenAICompatibleProviderChatCompletionParamsXSohaCacheModeRefresh:
		return true
	default:
		return false
	}
}

// Defines values for CreateAIGatewayOpenAICompatibleProviderChatCompletionParamsOpenaiCompatibleProvider.
const (
	CreateAIGatewayOpenAICompatibleProviderChatCompletionParamsOpenaiCompatibleProviderAzureOpenai CreateAIGatewayOpenAICompatibleProviderChatCompletionParamsOpenaiCompatibleProvider = "azure-openai"
	CreateAIGatewayOpenAICompatibleProviderChatCompletionParamsOpenaiCompatibleProviderDeepseek    CreateAIGatewayOpenAICompatibleProviderChatCompletionParamsOpenaiCompatibleProvider = "deepseek"
	CreateAIGatewayOpenAICompatibleProviderChatCompletionParamsOpenaiCompatibleProviderOpenrouter  CreateAIGatewayOpenAICompatibleProviderChatCompletionParamsOpenaiCompatibleProvider = "openrouter"
	CreateAIGatewayOpenAICompatibleProviderChatCompletionParamsOpenaiCompatibleProviderQwen        CreateAIGatewayOpenAICompatibleProviderChatCompletionParamsOpenaiCompatibleProvider = "qwen"
)

// Valid indicates whether the value is a known member of the CreateAIGatewayOpenAICompatibleProviderChatCompletionParamsOpenaiCompatibleProvider enum.
func (e CreateAIGatewayOpenAICompatibleProviderChatCompletionParamsOpenaiCompatibleProvider) Valid() bool {
	switch e {
	case CreateAIGatewayOpenAICompatibleProviderChatCompletionParamsOpenaiCompatibleProviderAzureOpenai:
		return true
	case CreateAIGatewayOpenAICompatibleProviderChatCompletionParamsOpenaiCompatibleProviderDeepseek:
		return true
	case CreateAIGatewayOpenAICompatibleProviderChatCompletionParamsOpenaiCompatibleProviderOpenrouter:
		return true
	case CreateAIGatewayOpenAICompatibleProviderChatCompletionParamsOpenaiCompatibleProviderQwen:
		return true
	default:
		return false
	}
}

// Defines values for CreateAIGatewayOpenAICompatibleProviderEmbeddingParamsXSohaCacheMode.
const (
	CreateAIGatewayOpenAICompatibleProviderEmbeddingParamsXSohaCacheModeBypass   CreateAIGatewayOpenAICompatibleProviderEmbeddingParamsXSohaCacheMode = "bypass"
	CreateAIGatewayOpenAICompatibleProviderEmbeddingParamsXSohaCacheModeDefault  CreateAIGatewayOpenAICompatibleProviderEmbeddingParamsXSohaCacheMode = "default"
	CreateAIGatewayOpenAICompatibleProviderEmbeddingParamsXSohaCacheModeReadOnly CreateAIGatewayOpenAICompatibleProviderEmbeddingParamsXSohaCacheMode = "read-only"
	CreateAIGatewayOpenAICompatibleProviderEmbeddingParamsXSohaCacheModeRefresh  CreateAIGatewayOpenAICompatibleProviderEmbeddingParamsXSohaCacheMode = "refresh"
)

// Valid indicates whether the value is a known member of the CreateAIGatewayOpenAICompatibleProviderEmbeddingParamsXSohaCacheMode enum.
func (e CreateAIGatewayOpenAICompatibleProviderEmbeddingParamsXSohaCacheMode) Valid() bool {
	switch e {
	case CreateAIGatewayOpenAICompatibleProviderEmbeddingParamsXSohaCacheModeBypass:
		return true
	case CreateAIGatewayOpenAICompatibleProviderEmbeddingParamsXSohaCacheModeDefault:
		return true
	case CreateAIGatewayOpenAICompatibleProviderEmbeddingParamsXSohaCacheModeReadOnly:
		return true
	case CreateAIGatewayOpenAICompatibleProviderEmbeddingParamsXSohaCacheModeRefresh:
		return true
	default:
		return false
	}
}

// Defines values for CreateAIGatewayOpenAICompatibleProviderEmbeddingParamsOpenaiCompatibleProvider.
const (
	CreateAIGatewayOpenAICompatibleProviderEmbeddingParamsOpenaiCompatibleProviderAzureOpenai CreateAIGatewayOpenAICompatibleProviderEmbeddingParamsOpenaiCompatibleProvider = "azure-openai"
	CreateAIGatewayOpenAICompatibleProviderEmbeddingParamsOpenaiCompatibleProviderDeepseek    CreateAIGatewayOpenAICompatibleProviderEmbeddingParamsOpenaiCompatibleProvider = "deepseek"
	CreateAIGatewayOpenAICompatibleProviderEmbeddingParamsOpenaiCompatibleProviderOpenrouter  CreateAIGatewayOpenAICompatibleProviderEmbeddingParamsOpenaiCompatibleProvider = "openrouter"
	CreateAIGatewayOpenAICompatibleProviderEmbeddingParamsOpenaiCompatibleProviderQwen        CreateAIGatewayOpenAICompatibleProviderEmbeddingParamsOpenaiCompatibleProvider = "qwen"
)

// Valid indicates whether the value is a known member of the CreateAIGatewayOpenAICompatibleProviderEmbeddingParamsOpenaiCompatibleProvider enum.
func (e CreateAIGatewayOpenAICompatibleProviderEmbeddingParamsOpenaiCompatibleProvider) Valid() bool {
	switch e {
	case CreateAIGatewayOpenAICompatibleProviderEmbeddingParamsOpenaiCompatibleProviderAzureOpenai:
		return true
	case CreateAIGatewayOpenAICompatibleProviderEmbeddingParamsOpenaiCompatibleProviderDeepseek:
		return true
	case CreateAIGatewayOpenAICompatibleProviderEmbeddingParamsOpenaiCompatibleProviderOpenrouter:
		return true
	case CreateAIGatewayOpenAICompatibleProviderEmbeddingParamsOpenaiCompatibleProviderQwen:
		return true
	default:
		return false
	}
}

// Defines values for CreateAIGatewayOpenAICompatibleProviderImageEditParamsXSohaCacheMode.
const (
	CreateAIGatewayOpenAICompatibleProviderImageEditParamsXSohaCacheModeBypass   CreateAIGatewayOpenAICompatibleProviderImageEditParamsXSohaCacheMode = "bypass"
	CreateAIGatewayOpenAICompatibleProviderImageEditParamsXSohaCacheModeDefault  CreateAIGatewayOpenAICompatibleProviderImageEditParamsXSohaCacheMode = "default"
	CreateAIGatewayOpenAICompatibleProviderImageEditParamsXSohaCacheModeReadOnly CreateAIGatewayOpenAICompatibleProviderImageEditParamsXSohaCacheMode = "read-only"
	CreateAIGatewayOpenAICompatibleProviderImageEditParamsXSohaCacheModeRefresh  CreateAIGatewayOpenAICompatibleProviderImageEditParamsXSohaCacheMode = "refresh"
)

// Valid indicates whether the value is a known member of the CreateAIGatewayOpenAICompatibleProviderImageEditParamsXSohaCacheMode enum.
func (e CreateAIGatewayOpenAICompatibleProviderImageEditParamsXSohaCacheMode) Valid() bool {
	switch e {
	case CreateAIGatewayOpenAICompatibleProviderImageEditParamsXSohaCacheModeBypass:
		return true
	case CreateAIGatewayOpenAICompatibleProviderImageEditParamsXSohaCacheModeDefault:
		return true
	case CreateAIGatewayOpenAICompatibleProviderImageEditParamsXSohaCacheModeReadOnly:
		return true
	case CreateAIGatewayOpenAICompatibleProviderImageEditParamsXSohaCacheModeRefresh:
		return true
	default:
		return false
	}
}

// Defines values for CreateAIGatewayOpenAICompatibleProviderImageEditParamsOpenaiCompatibleProvider.
const (
	CreateAIGatewayOpenAICompatibleProviderImageEditParamsOpenaiCompatibleProviderAzureOpenai CreateAIGatewayOpenAICompatibleProviderImageEditParamsOpenaiCompatibleProvider = "azure-openai"
	CreateAIGatewayOpenAICompatibleProviderImageEditParamsOpenaiCompatibleProviderDeepseek    CreateAIGatewayOpenAICompatibleProviderImageEditParamsOpenaiCompatibleProvider = "deepseek"
	CreateAIGatewayOpenAICompatibleProviderImageEditParamsOpenaiCompatibleProviderOpenrouter  CreateAIGatewayOpenAICompatibleProviderImageEditParamsOpenaiCompatibleProvider = "openrouter"
	CreateAIGatewayOpenAICompatibleProviderImageEditParamsOpenaiCompatibleProviderQwen        CreateAIGatewayOpenAICompatibleProviderImageEditParamsOpenaiCompatibleProvider = "qwen"
)

// Valid indicates whether the value is a known member of the CreateAIGatewayOpenAICompatibleProviderImageEditParamsOpenaiCompatibleProvider enum.
func (e CreateAIGatewayOpenAICompatibleProviderImageEditParamsOpenaiCompatibleProvider) Valid() bool {
	switch e {
	case CreateAIGatewayOpenAICompatibleProviderImageEditParamsOpenaiCompatibleProviderAzureOpenai:
		return true
	case CreateAIGatewayOpenAICompatibleProviderImageEditParamsOpenaiCompatibleProviderDeepseek:
		return true
	case CreateAIGatewayOpenAICompatibleProviderImageEditParamsOpenaiCompatibleProviderOpenrouter:
		return true
	case CreateAIGatewayOpenAICompatibleProviderImageEditParamsOpenaiCompatibleProviderQwen:
		return true
	default:
		return false
	}
}

// Defines values for CreateAIGatewayOpenAICompatibleProviderImageGenerationParamsXSohaCacheMode.
const (
	CreateAIGatewayOpenAICompatibleProviderImageGenerationParamsXSohaCacheModeBypass   CreateAIGatewayOpenAICompatibleProviderImageGenerationParamsXSohaCacheMode = "bypass"
	CreateAIGatewayOpenAICompatibleProviderImageGenerationParamsXSohaCacheModeDefault  CreateAIGatewayOpenAICompatibleProviderImageGenerationParamsXSohaCacheMode = "default"
	CreateAIGatewayOpenAICompatibleProviderImageGenerationParamsXSohaCacheModeReadOnly CreateAIGatewayOpenAICompatibleProviderImageGenerationParamsXSohaCacheMode = "read-only"
	CreateAIGatewayOpenAICompatibleProviderImageGenerationParamsXSohaCacheModeRefresh  CreateAIGatewayOpenAICompatibleProviderImageGenerationParamsXSohaCacheMode = "refresh"
)

// Valid indicates whether the value is a known member of the CreateAIGatewayOpenAICompatibleProviderImageGenerationParamsXSohaCacheMode enum.
func (e CreateAIGatewayOpenAICompatibleProviderImageGenerationParamsXSohaCacheMode) Valid() bool {
	switch e {
	case CreateAIGatewayOpenAICompatibleProviderImageGenerationParamsXSohaCacheModeBypass:
		return true
	case CreateAIGatewayOpenAICompatibleProviderImageGenerationParamsXSohaCacheModeDefault:
		return true
	case CreateAIGatewayOpenAICompatibleProviderImageGenerationParamsXSohaCacheModeReadOnly:
		return true
	case CreateAIGatewayOpenAICompatibleProviderImageGenerationParamsXSohaCacheModeRefresh:
		return true
	default:
		return false
	}
}

// Defines values for CreateAIGatewayOpenAICompatibleProviderImageGenerationParamsOpenaiCompatibleProvider.
const (
	CreateAIGatewayOpenAICompatibleProviderImageGenerationParamsOpenaiCompatibleProviderAzureOpenai CreateAIGatewayOpenAICompatibleProviderImageGenerationParamsOpenaiCompatibleProvider = "azure-openai"
	CreateAIGatewayOpenAICompatibleProviderImageGenerationParamsOpenaiCompatibleProviderDeepseek    CreateAIGatewayOpenAICompatibleProviderImageGenerationParamsOpenaiCompatibleProvider = "deepseek"
	CreateAIGatewayOpenAICompatibleProviderImageGenerationParamsOpenaiCompatibleProviderOpenrouter  CreateAIGatewayOpenAICompatibleProviderImageGenerationParamsOpenaiCompatibleProvider = "openrouter"
	CreateAIGatewayOpenAICompatibleProviderImageGenerationParamsOpenaiCompatibleProviderQwen        CreateAIGatewayOpenAICompatibleProviderImageGenerationParamsOpenaiCompatibleProvider = "qwen"
)

// Valid indicates whether the value is a known member of the CreateAIGatewayOpenAICompatibleProviderImageGenerationParamsOpenaiCompatibleProvider enum.
func (e CreateAIGatewayOpenAICompatibleProviderImageGenerationParamsOpenaiCompatibleProvider) Valid() bool {
	switch e {
	case CreateAIGatewayOpenAICompatibleProviderImageGenerationParamsOpenaiCompatibleProviderAzureOpenai:
		return true
	case CreateAIGatewayOpenAICompatibleProviderImageGenerationParamsOpenaiCompatibleProviderDeepseek:
		return true
	case CreateAIGatewayOpenAICompatibleProviderImageGenerationParamsOpenaiCompatibleProviderOpenrouter:
		return true
	case CreateAIGatewayOpenAICompatibleProviderImageGenerationParamsOpenaiCompatibleProviderQwen:
		return true
	default:
		return false
	}
}

// Defines values for CreateAIGatewayOpenAICompatibleProviderImageVariationParamsXSohaCacheMode.
const (
	CreateAIGatewayOpenAICompatibleProviderImageVariationParamsXSohaCacheModeBypass   CreateAIGatewayOpenAICompatibleProviderImageVariationParamsXSohaCacheMode = "bypass"
	CreateAIGatewayOpenAICompatibleProviderImageVariationParamsXSohaCacheModeDefault  CreateAIGatewayOpenAICompatibleProviderImageVariationParamsXSohaCacheMode = "default"
	CreateAIGatewayOpenAICompatibleProviderImageVariationParamsXSohaCacheModeReadOnly CreateAIGatewayOpenAICompatibleProviderImageVariationParamsXSohaCacheMode = "read-only"
	CreateAIGatewayOpenAICompatibleProviderImageVariationParamsXSohaCacheModeRefresh  CreateAIGatewayOpenAICompatibleProviderImageVariationParamsXSohaCacheMode = "refresh"
)

// Valid indicates whether the value is a known member of the CreateAIGatewayOpenAICompatibleProviderImageVariationParamsXSohaCacheMode enum.
func (e CreateAIGatewayOpenAICompatibleProviderImageVariationParamsXSohaCacheMode) Valid() bool {
	switch e {
	case CreateAIGatewayOpenAICompatibleProviderImageVariationParamsXSohaCacheModeBypass:
		return true
	case CreateAIGatewayOpenAICompatibleProviderImageVariationParamsXSohaCacheModeDefault:
		return true
	case CreateAIGatewayOpenAICompatibleProviderImageVariationParamsXSohaCacheModeReadOnly:
		return true
	case CreateAIGatewayOpenAICompatibleProviderImageVariationParamsXSohaCacheModeRefresh:
		return true
	default:
		return false
	}
}

// Defines values for CreateAIGatewayOpenAICompatibleProviderImageVariationParamsOpenaiCompatibleProvider.
const (
	CreateAIGatewayOpenAICompatibleProviderImageVariationParamsOpenaiCompatibleProviderAzureOpenai CreateAIGatewayOpenAICompatibleProviderImageVariationParamsOpenaiCompatibleProvider = "azure-openai"
	CreateAIGatewayOpenAICompatibleProviderImageVariationParamsOpenaiCompatibleProviderDeepseek    CreateAIGatewayOpenAICompatibleProviderImageVariationParamsOpenaiCompatibleProvider = "deepseek"
	CreateAIGatewayOpenAICompatibleProviderImageVariationParamsOpenaiCompatibleProviderOpenrouter  CreateAIGatewayOpenAICompatibleProviderImageVariationParamsOpenaiCompatibleProvider = "openrouter"
	CreateAIGatewayOpenAICompatibleProviderImageVariationParamsOpenaiCompatibleProviderQwen        CreateAIGatewayOpenAICompatibleProviderImageVariationParamsOpenaiCompatibleProvider = "qwen"
)

// Valid indicates whether the value is a known member of the CreateAIGatewayOpenAICompatibleProviderImageVariationParamsOpenaiCompatibleProvider enum.
func (e CreateAIGatewayOpenAICompatibleProviderImageVariationParamsOpenaiCompatibleProvider) Valid() bool {
	switch e {
	case CreateAIGatewayOpenAICompatibleProviderImageVariationParamsOpenaiCompatibleProviderAzureOpenai:
		return true
	case CreateAIGatewayOpenAICompatibleProviderImageVariationParamsOpenaiCompatibleProviderDeepseek:
		return true
	case CreateAIGatewayOpenAICompatibleProviderImageVariationParamsOpenaiCompatibleProviderOpenrouter:
		return true
	case CreateAIGatewayOpenAICompatibleProviderImageVariationParamsOpenaiCompatibleProviderQwen:
		return true
	default:
		return false
	}
}

// Defines values for ListAIGatewayOpenAICompatibleProviderModelsParamsOpenaiCompatibleProvider.
const (
	ListAIGatewayOpenAICompatibleProviderModelsParamsOpenaiCompatibleProviderAzureOpenai ListAIGatewayOpenAICompatibleProviderModelsParamsOpenaiCompatibleProvider = "azure-openai"
	ListAIGatewayOpenAICompatibleProviderModelsParamsOpenaiCompatibleProviderDeepseek    ListAIGatewayOpenAICompatibleProviderModelsParamsOpenaiCompatibleProvider = "deepseek"
	ListAIGatewayOpenAICompatibleProviderModelsParamsOpenaiCompatibleProviderOpenrouter  ListAIGatewayOpenAICompatibleProviderModelsParamsOpenaiCompatibleProvider = "openrouter"
	ListAIGatewayOpenAICompatibleProviderModelsParamsOpenaiCompatibleProviderQwen        ListAIGatewayOpenAICompatibleProviderModelsParamsOpenaiCompatibleProvider = "qwen"
)

// Valid indicates whether the value is a known member of the ListAIGatewayOpenAICompatibleProviderModelsParamsOpenaiCompatibleProvider enum.
func (e ListAIGatewayOpenAICompatibleProviderModelsParamsOpenaiCompatibleProvider) Valid() bool {
	switch e {
	case ListAIGatewayOpenAICompatibleProviderModelsParamsOpenaiCompatibleProviderAzureOpenai:
		return true
	case ListAIGatewayOpenAICompatibleProviderModelsParamsOpenaiCompatibleProviderDeepseek:
		return true
	case ListAIGatewayOpenAICompatibleProviderModelsParamsOpenaiCompatibleProviderOpenrouter:
		return true
	case ListAIGatewayOpenAICompatibleProviderModelsParamsOpenaiCompatibleProviderQwen:
		return true
	default:
		return false
	}
}

// Defines values for CreateAIGatewayOpenAICompatibleProviderResponseParamsXSohaCacheMode.
const (
	CreateAIGatewayOpenAICompatibleProviderResponseParamsXSohaCacheModeBypass   CreateAIGatewayOpenAICompatibleProviderResponseParamsXSohaCacheMode = "bypass"
	CreateAIGatewayOpenAICompatibleProviderResponseParamsXSohaCacheModeDefault  CreateAIGatewayOpenAICompatibleProviderResponseParamsXSohaCacheMode = "default"
	CreateAIGatewayOpenAICompatibleProviderResponseParamsXSohaCacheModeReadOnly CreateAIGatewayOpenAICompatibleProviderResponseParamsXSohaCacheMode = "read-only"
	CreateAIGatewayOpenAICompatibleProviderResponseParamsXSohaCacheModeRefresh  CreateAIGatewayOpenAICompatibleProviderResponseParamsXSohaCacheMode = "refresh"
)

// Valid indicates whether the value is a known member of the CreateAIGatewayOpenAICompatibleProviderResponseParamsXSohaCacheMode enum.
func (e CreateAIGatewayOpenAICompatibleProviderResponseParamsXSohaCacheMode) Valid() bool {
	switch e {
	case CreateAIGatewayOpenAICompatibleProviderResponseParamsXSohaCacheModeBypass:
		return true
	case CreateAIGatewayOpenAICompatibleProviderResponseParamsXSohaCacheModeDefault:
		return true
	case CreateAIGatewayOpenAICompatibleProviderResponseParamsXSohaCacheModeReadOnly:
		return true
	case CreateAIGatewayOpenAICompatibleProviderResponseParamsXSohaCacheModeRefresh:
		return true
	default:
		return false
	}
}

// Defines values for CreateAIGatewayOpenAICompatibleProviderResponseParamsOpenaiCompatibleProvider.
const (
	CreateAIGatewayOpenAICompatibleProviderResponseParamsOpenaiCompatibleProviderAzureOpenai CreateAIGatewayOpenAICompatibleProviderResponseParamsOpenaiCompatibleProvider = "azure-openai"
	CreateAIGatewayOpenAICompatibleProviderResponseParamsOpenaiCompatibleProviderDeepseek    CreateAIGatewayOpenAICompatibleProviderResponseParamsOpenaiCompatibleProvider = "deepseek"
	CreateAIGatewayOpenAICompatibleProviderResponseParamsOpenaiCompatibleProviderOpenrouter  CreateAIGatewayOpenAICompatibleProviderResponseParamsOpenaiCompatibleProvider = "openrouter"
	CreateAIGatewayOpenAICompatibleProviderResponseParamsOpenaiCompatibleProviderQwen        CreateAIGatewayOpenAICompatibleProviderResponseParamsOpenaiCompatibleProvider = "qwen"
)

// Valid indicates whether the value is a known member of the CreateAIGatewayOpenAICompatibleProviderResponseParamsOpenaiCompatibleProvider enum.
func (e CreateAIGatewayOpenAICompatibleProviderResponseParamsOpenaiCompatibleProvider) Valid() bool {
	switch e {
	case CreateAIGatewayOpenAICompatibleProviderResponseParamsOpenaiCompatibleProviderAzureOpenai:
		return true
	case CreateAIGatewayOpenAICompatibleProviderResponseParamsOpenaiCompatibleProviderDeepseek:
		return true
	case CreateAIGatewayOpenAICompatibleProviderResponseParamsOpenaiCompatibleProviderOpenrouter:
		return true
	case CreateAIGatewayOpenAICompatibleProviderResponseParamsOpenaiCompatibleProviderQwen:
		return true
	default:
		return false
	}
}

// Defines values for ListAIGatewayRelayModelCallsParamsProviderKind.
const (
	ListAIGatewayRelayModelCallsParamsProviderKindAnthropic        ListAIGatewayRelayModelCallsParamsProviderKind = "anthropic"
	ListAIGatewayRelayModelCallsParamsProviderKindAzureOpenai      ListAIGatewayRelayModelCallsParamsProviderKind = "azure-openai"
	ListAIGatewayRelayModelCallsParamsProviderKindCohere           ListAIGatewayRelayModelCallsParamsProviderKind = "cohere"
	ListAIGatewayRelayModelCallsParamsProviderKindDeepseek         ListAIGatewayRelayModelCallsParamsProviderKind = "deepseek"
	ListAIGatewayRelayModelCallsParamsProviderKindGemini           ListAIGatewayRelayModelCallsParamsProviderKind = "gemini"
	ListAIGatewayRelayModelCallsParamsProviderKindOpenai           ListAIGatewayRelayModelCallsParamsProviderKind = "openai"
	ListAIGatewayRelayModelCallsParamsProviderKindOpenaiCompatible ListAIGatewayRelayModelCallsParamsProviderKind = "openai-compatible"
	ListAIGatewayRelayModelCallsParamsProviderKindOpenrouter       ListAIGatewayRelayModelCallsParamsProviderKind = "openrouter"
	ListAIGatewayRelayModelCallsParamsProviderKindQwen             ListAIGatewayRelayModelCallsParamsProviderKind = "qwen"
)

// Valid indicates whether the value is a known member of the ListAIGatewayRelayModelCallsParamsProviderKind enum.
func (e ListAIGatewayRelayModelCallsParamsProviderKind) Valid() bool {
	switch e {
	case ListAIGatewayRelayModelCallsParamsProviderKindAnthropic:
		return true
	case ListAIGatewayRelayModelCallsParamsProviderKindAzureOpenai:
		return true
	case ListAIGatewayRelayModelCallsParamsProviderKindCohere:
		return true
	case ListAIGatewayRelayModelCallsParamsProviderKindDeepseek:
		return true
	case ListAIGatewayRelayModelCallsParamsProviderKindGemini:
		return true
	case ListAIGatewayRelayModelCallsParamsProviderKindOpenai:
		return true
	case ListAIGatewayRelayModelCallsParamsProviderKindOpenaiCompatible:
		return true
	case ListAIGatewayRelayModelCallsParamsProviderKindOpenrouter:
		return true
	case ListAIGatewayRelayModelCallsParamsProviderKindQwen:
		return true
	default:
		return false
	}
}

// Defines values for ListAIGatewayRelayModelCallsParamsStatus.
const (
	Cancelled       ListAIGatewayRelayModelCallsParamsStatus = "cancelled"
	ClientCancelled ListAIGatewayRelayModelCallsParamsStatus = "client_cancelled"
	Failure         ListAIGatewayRelayModelCallsParamsStatus = "failure"
	PolicyDenied    ListAIGatewayRelayModelCallsParamsStatus = "policy_denied"
	RateLimited     ListAIGatewayRelayModelCallsParamsStatus = "rate_limited"
	Success         ListAIGatewayRelayModelCallsParamsStatus = "success"
)

// Valid indicates whether the value is a known member of the ListAIGatewayRelayModelCallsParamsStatus enum.
func (e ListAIGatewayRelayModelCallsParamsStatus) Valid() bool {
	switch e {
	case Cancelled:
		return true
	case ClientCancelled:
		return true
	case Failure:
		return true
	case PolicyDenied:
		return true
	case RateLimited:
		return true
	case Success:
		return true
	default:
		return false
	}
}

// Defines values for ListAIGatewayRelayModelRoutesParamsProviderKind.
const (
	ListAIGatewayRelayModelRoutesParamsProviderKindAnthropic        ListAIGatewayRelayModelRoutesParamsProviderKind = "anthropic"
	ListAIGatewayRelayModelRoutesParamsProviderKindAzureOpenai      ListAIGatewayRelayModelRoutesParamsProviderKind = "azure-openai"
	ListAIGatewayRelayModelRoutesParamsProviderKindCohere           ListAIGatewayRelayModelRoutesParamsProviderKind = "cohere"
	ListAIGatewayRelayModelRoutesParamsProviderKindDeepseek         ListAIGatewayRelayModelRoutesParamsProviderKind = "deepseek"
	ListAIGatewayRelayModelRoutesParamsProviderKindGemini           ListAIGatewayRelayModelRoutesParamsProviderKind = "gemini"
	ListAIGatewayRelayModelRoutesParamsProviderKindOpenai           ListAIGatewayRelayModelRoutesParamsProviderKind = "openai"
	ListAIGatewayRelayModelRoutesParamsProviderKindOpenaiCompatible ListAIGatewayRelayModelRoutesParamsProviderKind = "openai-compatible"
	ListAIGatewayRelayModelRoutesParamsProviderKindOpenrouter       ListAIGatewayRelayModelRoutesParamsProviderKind = "openrouter"
	ListAIGatewayRelayModelRoutesParamsProviderKindQwen             ListAIGatewayRelayModelRoutesParamsProviderKind = "qwen"
)

// Valid indicates whether the value is a known member of the ListAIGatewayRelayModelRoutesParamsProviderKind enum.
func (e ListAIGatewayRelayModelRoutesParamsProviderKind) Valid() bool {
	switch e {
	case ListAIGatewayRelayModelRoutesParamsProviderKindAnthropic:
		return true
	case ListAIGatewayRelayModelRoutesParamsProviderKindAzureOpenai:
		return true
	case ListAIGatewayRelayModelRoutesParamsProviderKindCohere:
		return true
	case ListAIGatewayRelayModelRoutesParamsProviderKindDeepseek:
		return true
	case ListAIGatewayRelayModelRoutesParamsProviderKindGemini:
		return true
	case ListAIGatewayRelayModelRoutesParamsProviderKindOpenai:
		return true
	case ListAIGatewayRelayModelRoutesParamsProviderKindOpenaiCompatible:
		return true
	case ListAIGatewayRelayModelRoutesParamsProviderKindOpenrouter:
		return true
	case ListAIGatewayRelayModelRoutesParamsProviderKindQwen:
		return true
	default:
		return false
	}
}

// Defines values for ListAIGatewayRelayUpstreamsParamsProviderKind.
const (
	ListAIGatewayRelayUpstreamsParamsProviderKindAnthropic        ListAIGatewayRelayUpstreamsParamsProviderKind = "anthropic"
	ListAIGatewayRelayUpstreamsParamsProviderKindAzureOpenai      ListAIGatewayRelayUpstreamsParamsProviderKind = "azure-openai"
	ListAIGatewayRelayUpstreamsParamsProviderKindCohere           ListAIGatewayRelayUpstreamsParamsProviderKind = "cohere"
	ListAIGatewayRelayUpstreamsParamsProviderKindDeepseek         ListAIGatewayRelayUpstreamsParamsProviderKind = "deepseek"
	ListAIGatewayRelayUpstreamsParamsProviderKindGemini           ListAIGatewayRelayUpstreamsParamsProviderKind = "gemini"
	ListAIGatewayRelayUpstreamsParamsProviderKindOpenai           ListAIGatewayRelayUpstreamsParamsProviderKind = "openai"
	ListAIGatewayRelayUpstreamsParamsProviderKindOpenaiCompatible ListAIGatewayRelayUpstreamsParamsProviderKind = "openai-compatible"
	ListAIGatewayRelayUpstreamsParamsProviderKindOpenrouter       ListAIGatewayRelayUpstreamsParamsProviderKind = "openrouter"
	ListAIGatewayRelayUpstreamsParamsProviderKindQwen             ListAIGatewayRelayUpstreamsParamsProviderKind = "qwen"
)

// Valid indicates whether the value is a known member of the ListAIGatewayRelayUpstreamsParamsProviderKind enum.
func (e ListAIGatewayRelayUpstreamsParamsProviderKind) Valid() bool {
	switch e {
	case ListAIGatewayRelayUpstreamsParamsProviderKindAnthropic:
		return true
	case ListAIGatewayRelayUpstreamsParamsProviderKindAzureOpenai:
		return true
	case ListAIGatewayRelayUpstreamsParamsProviderKindCohere:
		return true
	case ListAIGatewayRelayUpstreamsParamsProviderKindDeepseek:
		return true
	case ListAIGatewayRelayUpstreamsParamsProviderKindGemini:
		return true
	case ListAIGatewayRelayUpstreamsParamsProviderKindOpenai:
		return true
	case ListAIGatewayRelayUpstreamsParamsProviderKindOpenaiCompatible:
		return true
	case ListAIGatewayRelayUpstreamsParamsProviderKindOpenrouter:
		return true
	case ListAIGatewayRelayUpstreamsParamsProviderKindQwen:
		return true
	default:
		return false
	}
}

// Defines values for ListAIGatewayRelayUpstreamsParamsStatus.
const (
	Active   ListAIGatewayRelayUpstreamsParamsStatus = "active"
	Degraded ListAIGatewayRelayUpstreamsParamsStatus = "degraded"
	Disabled ListAIGatewayRelayUpstreamsParamsStatus = "disabled"
)

// Valid indicates whether the value is a known member of the ListAIGatewayRelayUpstreamsParamsStatus enum.
func (e ListAIGatewayRelayUpstreamsParamsStatus) Valid() bool {
	switch e {
	case Active:
		return true
	case Degraded:
		return true
	case Disabled:
		return true
	default:
		return false
	}
}

// Defines values for TransitionAgentProviderRolloutParamsRolloutAction.
const (
	Pause    TransitionAgentProviderRolloutParamsRolloutAction = "pause"
	Resume   TransitionAgentProviderRolloutParamsRolloutAction = "resume"
	Rollback TransitionAgentProviderRolloutParamsRolloutAction = "rollback"
)

// Valid indicates whether the value is a known member of the TransitionAgentProviderRolloutParamsRolloutAction enum.
func (e TransitionAgentProviderRolloutParamsRolloutAction) Valid() bool {
	switch e {
	case Pause:
		return true
	case Resume:
		return true
	case Rollback:
		return true
	default:
		return false
	}
}

// AIContextBudgetUsage defines model for AIContextBudgetUsage.
type AIContextBudgetUsage struct {
	EvidenceItems  int `json:"evidenceItems"`
	EvidenceTokens int `json:"evidenceTokens"`
}

// AIContextBudgets defines model for AIContextBudgets.
type AIContextBudgets struct {
	MaxEvidenceTokens int `json:"maxEvidenceTokens,omitempty"`
	MaxInputTokens    int `json:"maxInputTokens,omitempty"`
	MaxSteps          int `json:"maxSteps,omitempty"`
}

// AIContextBuildInput defines model for AIContextBuildInput.
type AIContextBuildInput struct {
	AgentRunID  string                  `json:"agentRunId,omitempty"`
	Budgets     *AIContextBudgets       `json:"budgets,omitempty"`
	Environment *AIContextEnvironment   `json:"environment,omitempty"`
	Knowledge   AIContextKnowledgeInput `json:"knowledge"`
	Prompt      *AIContextVersionRef    `json:"prompt,omitempty"`
	RequestID   string                  `json:"requestId,omitempty"`
	Session     *AIContextSession       `json:"session,omitempty"`
	SessionID   string                  `json:"sessionId,omitempty"`
	Skills      []AIContextVersionRef   `json:"skills,omitempty"`
	Task        AIContextTask           `json:"task"`
	Tools       []AIContextToolRef      `json:"tools,omitempty"`
}

// AIContextEnvelope defines model for AIContextEnvelope.
type AIContextEnvelope struct {
	AgentRunID     string                  `json:"agentRunId,omitempty"`
	BudgetUsage    AIContextBudgetUsage    `json:"budgetUsage"`
	Budgets        AIContextBudgets        `json:"budgets"`
	Citations      []KnowledgeCitation     `json:"citations,omitempty"`
	ContentHash    string                  `json:"contentHash"`
	CreatedAt      time.Time               `json:"createdAt"`
	Environment    *AIContextEnvironment   `json:"environment,omitempty"`
	Evidence       []AIContextEvidence     `json:"evidence,omitempty"`
	ID             string                  `json:"id"`
	PolicySnapshot AIContextPolicySnapshot `json:"policySnapshot"`
	Principal      AIContextPrincipal      `json:"principal"`
	Prompt         *AIContextVersionRef    `json:"prompt,omitempty"`
	RequestID      string                  `json:"requestId"`
	Session        *AIContextSession       `json:"session,omitempty"`
	SessionID      string                  `json:"sessionId,omitempty"`
	Skills         []AIContextVersionRef   `json:"skills,omitempty"`
	Task           AIContextTask           `json:"task"`
	Tools          []AIContextToolRef      `json:"tools,omitempty"`
	Version        string                  `json:"version"`
}

// AIContextEnvironment defines model for AIContextEnvironment.
type AIContextEnvironment struct {
	Mode            string   `json:"mode,omitempty"`
	ObservationRefs []string `json:"observationRefs,omitempty"`
}

// AIContextEvidence defines model for AIContextEvidence.
type AIContextEvidence struct {
	CitationID string `json:"citationId"`
	Content    string `json:"content"`
	TokenCount int    `json:"tokenCount"`
}

// AIContextInspection defines model for AIContextInspection.
type AIContextInspection struct {
	Envelope        AIContextEnvelope `json:"envelope"`
	RetrievalTimeMs int64             `json:"retrievalTimeMs,omitempty"`
	Sections        []string          `json:"sections"`
	Truncations     []string          `json:"truncations,omitempty"`
}

// AIContextInspectionEnvelope defines model for AIContextInspectionEnvelope.
type AIContextInspectionEnvelope struct {
	Data AIContextInspection `json:"data"`
}

// AIContextKnowledgeInput defines model for AIContextKnowledgeInput.
type AIContextKnowledgeInput struct {
	Enabled          bool     `json:"enabled"`
	KnowledgeBaseIDs []string `json:"knowledgeBaseIds"`
	Query            string   `json:"query,omitempty"`
	TopK             int      `json:"topK,omitempty"`
}

// AIContextPolicySnapshot defines model for AIContextPolicySnapshot.
type AIContextPolicySnapshot struct {
	ID      string `json:"id"`
	Version string `json:"version"`
}

// AIContextPrincipal defines model for AIContextPrincipal.
type AIContextPrincipal struct {
	UserID string `json:"userId"`
}

// AIContextSession defines model for AIContextSession.
type AIContextSession struct {
	RecentMessageRefs []string `json:"recentMessageRefs,omitempty"`
	Summary           string   `json:"summary,omitempty"`
}

// AIContextTask defines model for AIContextTask.
type AIContextTask struct {
	Goal string `json:"goal"`
	Mode string `json:"mode,omitempty"`
}

// AIContextToolRef defines model for AIContextToolRef.
type AIContextToolRef struct {
	Name          string `json:"name"`
	SchemaVersion string `json:"schemaVersion,omitempty"`
}

// AIContextVersionRef defines model for AIContextVersionRef.
type AIContextVersionRef struct {
	ID      string `json:"id,omitempty"`
	Name    string `json:"name,omitempty"`
	Version string `json:"version,omitempty"`
}

// AIEnvironmentLease defines model for AIEnvironmentLease.
type AIEnvironmentLease struct {
	CreatedAt  time.Time  `json:"createdAt"`
	ExpiresAt  *time.Time `json:"expiresAt,omitempty"`
	ID         string     `json:"id"`
	OwnerRef   string     `json:"ownerRef"`
	Status     string     `json:"status"`
	TemplateID string     `json:"templateId"`
	UpdatedAt  time.Time  `json:"updatedAt"`
}

// AIEnvironmentLeaseEnvelope defines model for AIEnvironmentLeaseEnvelope.
type AIEnvironmentLeaseEnvelope struct {
	Data AIEnvironmentLease `json:"data"`
}

// AIEnvironmentLeaseListEnvelope defines model for AIEnvironmentLeaseListEnvelope.
type AIEnvironmentLeaseListEnvelope struct {
	Items []AIEnvironmentLease `json:"items"`
}

// AIEnvironmentTemplate defines model for AIEnvironmentTemplate.
type AIEnvironmentTemplate struct {
	Backend        AIEnvironmentTemplateBackend       `json:"backend"`
	CreatedAt      time.Time                          `json:"createdAt"`
	ID             string                             `json:"id"`
	IsolationMode  AIEnvironmentTemplateIsolationMode `json:"isolationMode"`
	Name           string                             `json:"name"`
	NetworkPolicy  map[string]any                     `json:"networkPolicy,omitempty"`
	ResourcePolicy map[string]any                     `json:"resourcePolicy,omitempty"`
	Status         string                             `json:"status"`
	UpdatedAt      time.Time                          `json:"updatedAt"`
}

// AIEnvironmentTemplateBackend defines model for AIEnvironmentTemplate.Backend.
type AIEnvironmentTemplateBackend string

// AIEnvironmentTemplateIsolationMode defines model for AIEnvironmentTemplate.IsolationMode.
type AIEnvironmentTemplateIsolationMode string

// AIEnvironmentTemplateEnvelope defines model for AIEnvironmentTemplateEnvelope.
type AIEnvironmentTemplateEnvelope struct {
	Data AIEnvironmentTemplate `json:"data"`
}

// AIEnvironmentTemplateInput defines model for AIEnvironmentTemplateInput.
type AIEnvironmentTemplateInput struct {
	Backend        AIEnvironmentTemplateInputBackend       `json:"backend"`
	ID             string                                  `json:"id"`
	IsolationMode  AIEnvironmentTemplateInputIsolationMode `json:"isolationMode"`
	Name           string                                  `json:"name"`
	NetworkPolicy  map[string]any                          `json:"networkPolicy,omitempty"`
	ResourcePolicy map[string]any                          `json:"resourcePolicy,omitempty"`
}

// AIEnvironmentTemplateInputBackend defines model for AIEnvironmentTemplateInput.Backend.
type AIEnvironmentTemplateInputBackend string

// AIEnvironmentTemplateInputIsolationMode defines model for AIEnvironmentTemplateInput.IsolationMode.
type AIEnvironmentTemplateInputIsolationMode string

// AIEnvironmentTemplateListEnvelope defines model for AIEnvironmentTemplateListEnvelope.
type AIEnvironmentTemplateListEnvelope struct {
	Items []AIEnvironmentTemplate `json:"items"`
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

// AIMemoryPolicy defines model for AIMemoryPolicy.
type AIMemoryPolicy struct {
	// DefaultTTL Go time.Duration encoded as nanoseconds.
	DefaultTTL        int64  `json:"defaultTtl"`
	Enabled           bool   `json:"enabled"`
	ExplicitWriteOnly bool   `json:"explicitWriteOnly"`
	ID                string `json:"id"`

	// MaximumTTL Go time.Duration encoded as nanoseconds.
	MaximumTTL        int64    `json:"maximumTtl"`
	MinimumConfidence float32  `json:"minimumConfidence"`
	OwnerTypes        []string `json:"ownerTypes"`
	Version           string   `json:"version"`
}

// AIMemoryPolicyEnvelope defines model for AIMemoryPolicyEnvelope.
type AIMemoryPolicyEnvelope struct {
	Data AIMemoryPolicy `json:"data"`
}

// AIMemoryPolicyInput defines model for AIMemoryPolicyInput.
type AIMemoryPolicyInput struct {
	ConsentMode AIMemoryPolicyInputConsentMode `json:"consentMode"`
	ID          string                         `json:"id"`
	Name        string                         `json:"name"`
	TTLDays     int                            `json:"ttlDays"`
}

// AIMemoryPolicyInputConsentMode defines model for AIMemoryPolicyInput.ConsentMode.
type AIMemoryPolicyInputConsentMode string

// AIMemoryPolicyListEnvelope defines model for AIMemoryPolicyListEnvelope.
type AIMemoryPolicyListEnvelope struct {
	Items []AIMemoryPolicy `json:"items"`
}

// AIMemoryRecord defines model for AIMemoryRecord.
type AIMemoryRecord struct {
	Confidence    float32                  `json:"confidence"`
	CreatedAt     time.Time                `json:"createdAt"`
	DeletedAt     *time.Time               `json:"deletedAt,omitempty"`
	ExpiresAt     *time.Time               `json:"expiresAt,omitempty"`
	Fact          string                   `json:"fact"`
	ID            string                   `json:"id"`
	OwnerID       string                   `json:"ownerId"`
	OwnerType     string                   `json:"ownerType"`
	PolicyVersion string                   `json:"policyVersion"`
	ScopeHash     string                   `json:"scopeHash"`
	SourceRefs    []string                 `json:"sourceRefs"`
	SourceType    AIMemoryRecordSourceType `json:"sourceType"`
	Status        string                   `json:"status"`
	ValidFrom     time.Time                `json:"validFrom"`
}

// AIMemoryRecordSourceType defines model for AIMemoryRecord.SourceType.
type AIMemoryRecordSourceType string

// AIMemoryRecordEnvelope defines model for AIMemoryRecordEnvelope.
type AIMemoryRecordEnvelope struct {
	Data AIMemoryRecord `json:"data"`
}

// AIMemoryRecordListEnvelope defines model for AIMemoryRecordListEnvelope.
type AIMemoryRecordListEnvelope struct {
	Items []AIMemoryRecord `json:"items"`
}

// AIMemoryWriteInput defines model for AIMemoryWriteInput.
type AIMemoryWriteInput struct {
	PolicyID      string         `json:"policyId"`
	PolicyVersion string         `json:"policyVersion"`
	Record        AIMemoryRecord `json:"record"`
}

// AIProductionOperation defines model for AIProductionOperation.
type AIProductionOperation struct {
	Category     string                    `json:"category"`
	CreatedAt    time.Time                 `json:"createdAt"`
	EvidenceRefs []string                  `json:"evidenceRefs,omitempty"`
	ID           string                    `json:"id"`
	Kind         AIProductionOperationKind `json:"kind"`
	RunbookID    string                    `json:"runbookId"`
	Status       string                    `json:"status"`
	TargetRef    string                    `json:"targetRef"`
	UpdatedAt    time.Time                 `json:"updatedAt"`
}

// AIProductionOperationKind defines model for AIProductionOperation.Kind.
type AIProductionOperationKind string

// AIProductionOperationEnvelope defines model for AIProductionOperationEnvelope.
type AIProductionOperationEnvelope struct {
	Data AIProductionOperation `json:"data"`
}

// AIProductionOperationInput defines model for AIProductionOperationInput.
type AIProductionOperationInput struct {
	EvidenceRefs []string                       `json:"evidenceRefs,omitempty"`
	ID           string                         `json:"id,omitempty"`
	Kind         AIProductionOperationInputKind `json:"kind"`
	RunbookID    string                         `json:"runbookId"`
	TargetRef    string                         `json:"targetRef"`
}

// AIProductionOperationInputKind defines model for AIProductionOperationInput.Kind.
type AIProductionOperationInputKind string

// AIProductionOperationListEnvelope defines model for AIProductionOperationListEnvelope.
type AIProductionOperationListEnvelope struct {
	Items []AIProductionOperation `json:"items"`
}

// AIRunbookEvidence defines model for AIRunbookEvidence.
type AIRunbookEvidence struct {
	CreatedAt    time.Time `json:"createdAt"`
	EvidenceRefs []string  `json:"evidenceRefs,omitempty"`
	ID           string    `json:"id"`
	OperationID  string    `json:"operationId"`
	Outcome      string    `json:"outcome"`
	RunbookID    string    `json:"runbookId"`
	Status       string    `json:"status"`
}

// AIRunbookEvidenceListEnvelope defines model for AIRunbookEvidenceListEnvelope.
type AIRunbookEvidenceListEnvelope struct {
	Items []AIRunbookEvidence `json:"items"`
}

// AISettings defines model for AISettings.
type AISettings struct {
	SkillsRegistry []AISkillSettings        `json:"skillsRegistry"`
	WorkbenchModel AIWorkbenchModelSettings `json:"workbenchModel"`
}

// AISettingsEnvelope defines model for AISettingsEnvelope.
type AISettingsEnvelope struct {
	Data AISettings `json:"data"`
}

// AISkillSettings defines model for AISkillSettings.
type AISkillSettings struct {
	BlueprintRefs  []string       `json:"blueprintRefs,omitempty"`
	CapabilityRefs []string       `json:"capabilityRefs,omitempty"`
	Category       string         `json:"category,omitempty"`
	Description    string         `json:"description,omitempty"`
	Enabled        bool           `json:"enabled"`
	ID             string         `json:"id"`
	InputSchema    map[string]any `json:"inputSchema,omitempty"`
	Name           string         `json:"name"`
	OutputSchema   map[string]any `json:"outputSchema,omitempty"`
	OwnerModule    string         `json:"ownerModule,omitempty"`
	ScopeRules     []string       `json:"scopeRules,omitempty"`
	Scopes         []string       `json:"scopes,omitempty"`
}

// AIWorkbenchModelSettings defines model for AIWorkbenchModelSettings.
type AIWorkbenchModelSettings struct {
	// DefaultEndpoint Workbench relay endpoint. Supported values are chat/completions, responses, and messages.
	DefaultEndpoint string `json:"defaultEndpoint,omitempty"`

	// DefaultPublicModel Public model name exposed by AI Gateway model routes and used as the Workbench default.
	DefaultPublicModel string `json:"defaultPublicModel,omitempty"`

	// DefaultRouteID Stable AI Gateway model-route ID used as the Workbench default when set.
	DefaultRouteID string `json:"defaultRouteId,omitempty"`
	Enabled        bool   `json:"enabled"`
}

// AgentProviderCatalog defines model for AgentProviderCatalog.
type AgentProviderCatalog struct {
	CreatedAt     time.Time                 `json:"createdAt"`
	Digest        string                    `json:"digest"`
	Providers     []AgentProviderDefinition `json:"providers"`
	Revision      int64                     `json:"revision"`
	SchemaVersion string                    `json:"schemaVersion"`
}

// AgentProviderCatalogEnvelope defines model for AgentProviderCatalogEnvelope.
type AgentProviderCatalogEnvelope struct {
	Data AgentProviderCatalog `json:"data"`
}

// AgentProviderConformanceResult defines model for AgentProviderConformanceResult.
type AgentProviderConformanceResult struct {
	ProviderID string `json:"providerId"`
	Reason     string `json:"reason,omitempty"`
	Status     string `json:"status"`
}

// AgentProviderConformanceRun defines model for AgentProviderConformanceRun.
type AgentProviderConformanceRun struct {
	CreatedAt      time.Time        `json:"createdAt"`
	EnvironmentRef string           `json:"environmentRef"`
	ID             string           `json:"id"`
	ProviderRef    string           `json:"providerRef"`
	Results        []map[string]any `json:"results,omitempty"`
	Status         string           `json:"status"`
	SuiteVersion   string           `json:"suiteVersion"`
	UpdatedAt      time.Time        `json:"updatedAt"`
}

// AgentProviderConformanceRunEnvelope defines model for AgentProviderConformanceRunEnvelope.
type AgentProviderConformanceRunEnvelope struct {
	Data AgentProviderConformanceRun `json:"data"`
}

// AgentProviderConformanceRunInput defines model for AgentProviderConformanceRunInput.
type AgentProviderConformanceRunInput struct {
	EnvironmentRef string `json:"environmentRef"`
	ID             string `json:"id"`
	ProviderRef    string `json:"providerRef"`
	SuiteVersion   string `json:"suiteVersion"`
}

// AgentProviderConformanceRunListEnvelope defines model for AgentProviderConformanceRunListEnvelope.
type AgentProviderConformanceRunListEnvelope struct {
	Items []AgentProviderConformanceRun `json:"items"`
}

// AgentProviderDefinition defines model for AgentProviderDefinition.
type AgentProviderDefinition struct {
	AdapterProtocol             string                         `json:"adapterProtocol"`
	Capabilities                []string                       `json:"capabilities"`
	DisplayName                 string                         `json:"displayName"`
	Draining                    bool                           `json:"draining,omitempty"`
	ID                          string                         `json:"id"`
	Kind                        string                         `json:"kind"`
	PluginID                    string                         `json:"pluginId"`
	PluginVersion               string                         `json:"pluginVersion"`
	ProviderVersion             string                         `json:"providerVersion"`
	RequiredGatewayCapabilities []string                       `json:"requiredGatewayCapabilities,omitempty"`
	RequiredScopes              []string                       `json:"requiredScopes,omitempty"`
	Runtime                     AgentProviderRuntimeDefinition `json:"runtime"`
	SchemaVersion               string                         `json:"schemaVersion"`
	SecretRefs                  []string                       `json:"secretRefs,omitempty"`
}

// AgentProviderFleetTarget defines model for AgentProviderFleetTarget.
type AgentProviderFleetTarget struct {
	Architectures []string          `json:"architectures,omitempty"`
	Environments  []string          `json:"environments,omitempty"`
	Labels        map[string]string `json:"labels,omitempty"`
	Platforms     []string          `json:"platforms,omitempty"`
}

// AgentProviderRegistryAcknowledgement defines model for AgentProviderRegistryAcknowledgement.
type AgentProviderRegistryAcknowledgement struct {
	Accepted          bool                             `json:"accepted"`
	ActiveRevision    int64                            `json:"activeRevision"`
	ConformanceChecks []AgentProviderConformanceResult `json:"conformanceChecks,omitempty"`
	DesiredRevision   int64                            `json:"desiredRevision,omitempty"`
	LkgRevision       int64                            `json:"lkgRevision,omitempty"`
	ObservedAt        time.Time                        `json:"observedAt"`
	PreviousRevision  int64                            `json:"previousRevision,omitempty"`
	ProviderStatuses  []AgentProviderRunnerStatus      `json:"providerStatuses,omitempty"`
	Reason            string                           `json:"reason,omitempty"`
	Revision          int64                            `json:"revision"`
	RolledBack        bool                             `json:"rolledBack,omitempty"`
	RolloutState      string                           `json:"rolloutState,omitempty"`
	RunnerID          string                           `json:"runnerId"`
	Targeted          bool                             `json:"targeted"`
}

// AgentProviderRegistryAcknowledgementEnvelope defines model for AgentProviderRegistryAcknowledgementEnvelope.
type AgentProviderRegistryAcknowledgementEnvelope struct {
	Data AgentProviderRegistryAcknowledgement `json:"data"`
}

// AgentProviderRegistrySnapshot defines model for AgentProviderRegistrySnapshot.
type AgentProviderRegistrySnapshot struct {
	Digest        string                    `json:"digest,omitempty"`
	FleetTarget   *AgentProviderFleetTarget `json:"fleetTarget,omitempty"`
	IssuedAt      time.Time                 `json:"issuedAt"`
	Providers     []AgentProviderDefinition `json:"providers"`
	Revision      int64                     `json:"revision"`
	SchemaVersion string                    `json:"schemaVersion"`
}

// AgentProviderRegistrySnapshotEnvelope defines model for AgentProviderRegistrySnapshotEnvelope.
type AgentProviderRegistrySnapshotEnvelope struct {
	Data AgentProviderRegistrySnapshot `json:"data"`
}

// AgentProviderRollout defines model for AgentProviderRollout.
type AgentProviderRollout struct {
	CanaryPercent    int                      `json:"canaryPercent"`
	CreatedAt        time.Time                `json:"createdAt"`
	DesiredRevision  int64                    `json:"desiredRevision"`
	ID               string                   `json:"id"`
	Name             string                   `json:"name,omitempty"`
	PreviousRevision int64                    `json:"previousRevision"`
	Status           string                   `json:"status"`
	Target           AgentProviderFleetTarget `json:"target"`
	UpdatedAt        time.Time                `json:"updatedAt"`
}

// AgentProviderRolloutEnvelope defines model for AgentProviderRolloutEnvelope.
type AgentProviderRolloutEnvelope struct {
	Data AgentProviderRollout `json:"data"`
}

// AgentProviderRolloutInput defines model for AgentProviderRolloutInput.
type AgentProviderRolloutInput struct {
	Architectures    []string          `json:"architectures,omitempty"`
	CanaryPercent    int               `json:"canaryPercent"`
	DesiredRevision  int64             `json:"desiredRevision"`
	Environments     []string          `json:"environments,omitempty"`
	ID               string            `json:"id"`
	Labels           map[string]string `json:"labels,omitempty"`
	Name             string            `json:"name,omitempty"`
	Platforms        []string          `json:"platforms,omitempty"`
	PreviousRevision int64             `json:"previousRevision,omitempty"`
}

// AgentProviderRolloutListEnvelope defines model for AgentProviderRolloutListEnvelope.
type AgentProviderRolloutListEnvelope struct {
	Items []AgentProviderRollout `json:"items"`
}

// AgentProviderRunnerStatus defines model for AgentProviderRunnerStatus.
type AgentProviderRunnerStatus struct {
	ActiveRuns      int                             `json:"activeRuns"`
	CatalogRevision int64                           `json:"catalogRevision"`
	Draining        bool                            `json:"draining"`
	Health          AgentProviderRunnerStatusHealth `json:"health"`
	ObservedAt      time.Time                       `json:"observedAt"`
	ProviderID      string                          `json:"providerId"`
	ProviderVersion string                          `json:"providerVersion"`
	Reason          string                          `json:"reason,omitempty"`
}

// AgentProviderRunnerStatusHealth defines model for AgentProviderRunnerStatus.Health.
type AgentProviderRunnerStatusHealth string

// AgentProviderRuntimeDefinition defines model for AgentProviderRuntimeDefinition.
type AgentProviderRuntimeDefinition struct {
	Args             []string                           `json:"args,omitempty"`
	Command          string                             `json:"command,omitempty"`
	Endpoint         string                             `json:"endpoint,omitempty"`
	HealthPath       string                             `json:"healthPath,omitempty"`
	Image            string                             `json:"image,omitempty"`
	Kind             AgentProviderRuntimeDefinitionKind `json:"kind"`
	PromptArg        string                             `json:"promptArg,omitempty"`
	ProviderSkillArg string                             `json:"providerSkillArg,omitempty"`
	SkillArg         string                             `json:"skillArg,omitempty"`
}

// AgentProviderRuntimeDefinitionKind defines model for AgentProviderRuntimeDefinition.Kind.
type AgentProviderRuntimeDefinitionKind string

// AgentProviderRuntimeStatus defines model for AgentProviderRuntimeStatus.
type AgentProviderRuntimeStatus struct {
	Acknowledgements []AgentProviderRegistryAcknowledgement `json:"acknowledgements"`
	CatalogDigest    string                                 `json:"catalogDigest"`
	CatalogRevision  int64                                  `json:"catalogRevision"`
	RunnerCount      int                                    `json:"runnerCount"`
}

// AgentProviderRuntimeStatusEnvelope defines model for AgentProviderRuntimeStatusEnvelope.
type AgentProviderRuntimeStatusEnvelope struct {
	Data AgentProviderRuntimeStatus `json:"data"`
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
	AgentID           string                                 `json:"agentId"`
	AnalysisArtifacts []map[string]any                       `json:"analysisArtifacts,omitempty"`
	CallbackToken     string                                 `json:"callbackToken"`
	ErrorMessage      string                                 `json:"errorMessage,omitempty"`
	Events            []AgentRunCallbackWorkbenchStreamEvent `json:"events,omitempty"`
	ExternalRunID     string                                 `json:"externalRunId,omitempty"`
	Payload           map[string]any                         `json:"payload"`
	RunID             string                                 `json:"runId"`
	Status            string                                 `json:"status"`
	ToolExecutions    []map[string]any                       `json:"toolExecutions,omitempty"`
}

// AgentRunCallbackWorkbenchAgentStatusEvent defines model for AgentRunCallbackWorkbenchAgentStatusEvent.
type AgentRunCallbackWorkbenchAgentStatusEvent struct {
	CreatedAt    *time.Time                                            `json:"createdAt,omitempty"`
	ID           string                                                `json:"id,omitempty"`
	MessageID    string                                                `json:"messageId,omitempty"`
	ProviderID   string                                                `json:"providerId,omitempty"`
	ProviderKind AgentRunCallbackWorkbenchAgentStatusEventProviderKind `json:"providerKind,omitempty"`
	RunID        string                                                `json:"runId,omitempty"`
	Sequence     int                                                   `json:"sequence,omitempty"`
	SessionID    string                                                `json:"sessionId,omitempty"`
	Status       AgentRunCallbackWorkbenchAgentStatusEventStatus       `json:"status"`
	Type         AgentRunCallbackWorkbenchAgentStatusEventType         `json:"type"`
}

// AgentRunCallbackWorkbenchAgentStatusEventProviderKind defines model for AgentRunCallbackWorkbenchAgentStatusEvent.ProviderKind.
type AgentRunCallbackWorkbenchAgentStatusEventProviderKind string

// AgentRunCallbackWorkbenchAgentStatusEventStatus defines model for AgentRunCallbackWorkbenchAgentStatusEvent.Status.
type AgentRunCallbackWorkbenchAgentStatusEventStatus string

// AgentRunCallbackWorkbenchAgentStatusEventType defines model for AgentRunCallbackWorkbenchAgentStatusEvent.Type.
type AgentRunCallbackWorkbenchAgentStatusEventType string

// AgentRunCallbackWorkbenchArtifactUpdatedEvent defines model for AgentRunCallbackWorkbenchArtifactUpdatedEvent.
type AgentRunCallbackWorkbenchArtifactUpdatedEvent struct {
	Artifact  AnyValue                                          `json:"artifact"`
	CreatedAt *time.Time                                        `json:"createdAt,omitempty"`
	ID        string                                            `json:"id,omitempty"`
	MessageID string                                            `json:"messageId,omitempty"`
	RunID     string                                            `json:"runId,omitempty"`
	Sequence  int                                               `json:"sequence,omitempty"`
	SessionID string                                            `json:"sessionId,omitempty"`
	Type      AgentRunCallbackWorkbenchArtifactUpdatedEventType `json:"type"`
}

// AgentRunCallbackWorkbenchArtifactUpdatedEventType defines model for AgentRunCallbackWorkbenchArtifactUpdatedEvent.Type.
type AgentRunCallbackWorkbenchArtifactUpdatedEventType string

// AgentRunCallbackWorkbenchCardCommandEvent defines model for AgentRunCallbackWorkbenchCardCommandEvent.
type AgentRunCallbackWorkbenchCardCommandEvent struct {
	Command   AnyValue                                      `json:"command"`
	CreatedAt *time.Time                                    `json:"createdAt,omitempty"`
	ID        string                                        `json:"id,omitempty"`
	MessageID string                                        `json:"messageId,omitempty"`
	RunID     string                                        `json:"runId,omitempty"`
	Sequence  int                                           `json:"sequence,omitempty"`
	SessionID string                                        `json:"sessionId,omitempty"`
	SurfaceID string                                        `json:"surfaceId"`
	Type      AgentRunCallbackWorkbenchCardCommandEventType `json:"type"`
}

// AgentRunCallbackWorkbenchCardCommandEventType defines model for AgentRunCallbackWorkbenchCardCommandEvent.Type.
type AgentRunCallbackWorkbenchCardCommandEventType string

// AgentRunCallbackWorkbenchErrorEvent defines model for AgentRunCallbackWorkbenchErrorEvent.
type AgentRunCallbackWorkbenchErrorEvent struct {
	Code      string                                  `json:"code,omitempty"`
	CreatedAt *time.Time                              `json:"createdAt,omitempty"`
	ID        string                                  `json:"id,omitempty"`
	Message   string                                  `json:"message"`
	MessageID string                                  `json:"messageId,omitempty"`
	Retryable bool                                    `json:"retryable,omitempty"`
	RunID     string                                  `json:"runId,omitempty"`
	Sequence  int                                     `json:"sequence,omitempty"`
	SessionID string                                  `json:"sessionId,omitempty"`
	Type      AgentRunCallbackWorkbenchErrorEventType `json:"type"`
}

// AgentRunCallbackWorkbenchErrorEventType defines model for AgentRunCallbackWorkbenchErrorEvent.Type.
type AgentRunCallbackWorkbenchErrorEventType string

// AgentRunCallbackWorkbenchMessageDeltaEvent defines model for AgentRunCallbackWorkbenchMessageDeltaEvent.
type AgentRunCallbackWorkbenchMessageDeltaEvent struct {
	ContentDelta string                                         `json:"contentDelta"`
	CreatedAt    *time.Time                                     `json:"createdAt,omitempty"`
	ID           string                                         `json:"id,omitempty"`
	MessageID    string                                         `json:"messageId,omitempty"`
	Role         AgentRunCallbackWorkbenchMessageDeltaEventRole `json:"role"`
	RunID        string                                         `json:"runId,omitempty"`
	Sequence     int                                            `json:"sequence,omitempty"`
	SessionID    string                                         `json:"sessionId,omitempty"`
	Type         AgentRunCallbackWorkbenchMessageDeltaEventType `json:"type"`
}

// AgentRunCallbackWorkbenchMessageDeltaEventRole defines model for AgentRunCallbackWorkbenchMessageDeltaEvent.Role.
type AgentRunCallbackWorkbenchMessageDeltaEventRole string

// AgentRunCallbackWorkbenchMessageDeltaEventType defines model for AgentRunCallbackWorkbenchMessageDeltaEvent.Type.
type AgentRunCallbackWorkbenchMessageDeltaEventType string

// AgentRunCallbackWorkbenchMessageDoneEvent defines model for AgentRunCallbackWorkbenchMessageDoneEvent.
type AgentRunCallbackWorkbenchMessageDoneEvent struct {
	Content   string                                        `json:"content"`
	CreatedAt *time.Time                                    `json:"createdAt,omitempty"`
	ID        string                                        `json:"id,omitempty"`
	MessageID string                                        `json:"messageId,omitempty"`
	Metadata  map[string]any                                `json:"metadata,omitempty"`
	Role      AgentRunCallbackWorkbenchMessageDoneEventRole `json:"role"`
	RunID     string                                        `json:"runId,omitempty"`
	Sequence  int                                           `json:"sequence,omitempty"`
	SessionID string                                        `json:"sessionId,omitempty"`
	Type      AgentRunCallbackWorkbenchMessageDoneEventType `json:"type"`
}

// AgentRunCallbackWorkbenchMessageDoneEventRole defines model for AgentRunCallbackWorkbenchMessageDoneEvent.Role.
type AgentRunCallbackWorkbenchMessageDoneEventRole string

// AgentRunCallbackWorkbenchMessageDoneEventType defines model for AgentRunCallbackWorkbenchMessageDoneEvent.Type.
type AgentRunCallbackWorkbenchMessageDoneEventType string

// AgentRunCallbackWorkbenchSourceUpdatedEvent defines model for AgentRunCallbackWorkbenchSourceUpdatedEvent.
type AgentRunCallbackWorkbenchSourceUpdatedEvent struct {
	CreatedAt *time.Time                                      `json:"createdAt,omitempty"`
	ID        string                                          `json:"id,omitempty"`
	MessageID string                                          `json:"messageId,omitempty"`
	RunID     string                                          `json:"runId,omitempty"`
	Sequence  int                                             `json:"sequence,omitempty"`
	SessionID string                                          `json:"sessionId,omitempty"`
	Source    WorkbenchSource                                 `json:"source"`
	Type      AgentRunCallbackWorkbenchSourceUpdatedEventType `json:"type"`
}

// AgentRunCallbackWorkbenchSourceUpdatedEventType defines model for AgentRunCallbackWorkbenchSourceUpdatedEvent.Type.
type AgentRunCallbackWorkbenchSourceUpdatedEventType string

// AgentRunCallbackWorkbenchStreamEvent Workbench stream event input accepted from external runners during AgentRunCallbackRequest. Canonical stream/replay events use WorkbenchStreamEvent and keep server-owned base fields required.
type AgentRunCallbackWorkbenchStreamEvent struct {
	union json.RawMessage
}

// AgentRunCallbackWorkbenchStreamEventBase Runner-provided Workbench event fields accepted in an agent run callback. The server owns and may assign or override id, sessionId, runId, sequence, and createdAt.
type AgentRunCallbackWorkbenchStreamEventBase struct {
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	ID        string     `json:"id,omitempty"`
	MessageID string     `json:"messageId,omitempty"`
	RunID     string     `json:"runId,omitempty"`
	Sequence  int        `json:"sequence,omitempty"`
	SessionID string     `json:"sessionId,omitempty"`
}

// AgentRunCallbackWorkbenchThinkingDeltaEvent defines model for AgentRunCallbackWorkbenchThinkingDeltaEvent.
type AgentRunCallbackWorkbenchThinkingDeltaEvent struct {
	CreatedAt *time.Time                                      `json:"createdAt,omitempty"`
	ID        string                                          `json:"id,omitempty"`
	MessageID string                                          `json:"messageId,omitempty"`
	RunID     string                                          `json:"runId,omitempty"`
	Sequence  int                                             `json:"sequence,omitempty"`
	SessionID string                                          `json:"sessionId,omitempty"`
	TextDelta string                                          `json:"textDelta"`
	Type      AgentRunCallbackWorkbenchThinkingDeltaEventType `json:"type"`
}

// AgentRunCallbackWorkbenchThinkingDeltaEventType defines model for AgentRunCallbackWorkbenchThinkingDeltaEvent.Type.
type AgentRunCallbackWorkbenchThinkingDeltaEventType string

// AgentRunCallbackWorkbenchThinkingDoneEvent defines model for AgentRunCallbackWorkbenchThinkingDoneEvent.
type AgentRunCallbackWorkbenchThinkingDoneEvent struct {
	Collapsed bool                                           `json:"collapsed"`
	CreatedAt *time.Time                                     `json:"createdAt,omitempty"`
	ID        string                                         `json:"id,omitempty"`
	MessageID string                                         `json:"messageId,omitempty"`
	RunID     string                                         `json:"runId,omitempty"`
	Sequence  int                                            `json:"sequence,omitempty"`
	SessionID string                                         `json:"sessionId,omitempty"`
	Summary   string                                         `json:"summary"`
	Type      AgentRunCallbackWorkbenchThinkingDoneEventType `json:"type"`
}

// AgentRunCallbackWorkbenchThinkingDoneEventType defines model for AgentRunCallbackWorkbenchThinkingDoneEvent.Type.
type AgentRunCallbackWorkbenchThinkingDoneEventType string

// AgentRunCallbackWorkbenchToolCompletedEvent defines model for AgentRunCallbackWorkbenchToolCompletedEvent.
type AgentRunCallbackWorkbenchToolCompletedEvent struct {
	CreatedAt *time.Time                                      `json:"createdAt,omitempty"`
	ID        string                                          `json:"id,omitempty"`
	MessageID string                                          `json:"messageId,omitempty"`
	RunID     string                                          `json:"runId,omitempty"`
	Sequence  int                                             `json:"sequence,omitempty"`
	SessionID string                                          `json:"sessionId,omitempty"`
	ToolCall  WorkbenchToolCall                               `json:"toolCall"`
	Type      AgentRunCallbackWorkbenchToolCompletedEventType `json:"type"`
}

// AgentRunCallbackWorkbenchToolCompletedEventType defines model for AgentRunCallbackWorkbenchToolCompletedEvent.Type.
type AgentRunCallbackWorkbenchToolCompletedEventType string

// AgentRunCallbackWorkbenchToolDeltaEvent defines model for AgentRunCallbackWorkbenchToolDeltaEvent.
type AgentRunCallbackWorkbenchToolDeltaEvent struct {
	CreatedAt   *time.Time                                  `json:"createdAt,omitempty"`
	ID          string                                      `json:"id,omitempty"`
	LogDelta    string                                      `json:"logDelta,omitempty"`
	MessageID   string                                      `json:"messageId,omitempty"`
	OutputDelta string                                      `json:"outputDelta,omitempty"`
	RunID       string                                      `json:"runId,omitempty"`
	Sequence    int                                         `json:"sequence,omitempty"`
	SessionID   string                                      `json:"sessionId,omitempty"`
	ToolCallID  string                                      `json:"toolCallId"`
	Type        AgentRunCallbackWorkbenchToolDeltaEventType `json:"type"`
}

// AgentRunCallbackWorkbenchToolDeltaEventType defines model for AgentRunCallbackWorkbenchToolDeltaEvent.Type.
type AgentRunCallbackWorkbenchToolDeltaEventType string

// AgentRunCallbackWorkbenchToolStartedEvent defines model for AgentRunCallbackWorkbenchToolStartedEvent.
type AgentRunCallbackWorkbenchToolStartedEvent struct {
	CreatedAt *time.Time                                    `json:"createdAt,omitempty"`
	ID        string                                        `json:"id,omitempty"`
	MessageID string                                        `json:"messageId,omitempty"`
	RunID     string                                        `json:"runId,omitempty"`
	Sequence  int                                           `json:"sequence,omitempty"`
	SessionID string                                        `json:"sessionId,omitempty"`
	ToolCall  WorkbenchToolCall                             `json:"toolCall"`
	Type      AgentRunCallbackWorkbenchToolStartedEventType `json:"type"`
}

// AgentRunCallbackWorkbenchToolStartedEventType defines model for AgentRunCallbackWorkbenchToolStartedEvent.Type.
type AgentRunCallbackWorkbenchToolStartedEventType string

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

// AnthropicMessagesRequest defines model for AnthropicMessagesRequest.
type AnthropicMessagesRequest struct {
	MaxTokens            int              `json:"max_tokens"`
	Messages             []map[string]any `json:"messages"`
	Model                string           `json:"model"`
	Stream               bool             `json:"stream,omitempty"`
	System               AnyValue         `json:"system,omitempty"`
	AdditionalProperties map[string]any   `json:"-"`
}

// AnthropicModel defines model for AnthropicModel.
type AnthropicModel struct {
	CreatedAt            string         `json:"created_at,omitempty"`
	DisplayName          string         `json:"display_name,omitempty"`
	ID                   string         `json:"id"`
	Type                 string         `json:"type"`
	AdditionalProperties map[string]any `json:"-"`
}

// AnthropicModelsResponse defines model for AnthropicModelsResponse.
type AnthropicModelsResponse struct {
	Data                 []AnthropicModel `json:"data"`
	FirstID              string           `json:"first_id,omitempty"`
	HasMore              bool             `json:"has_more,omitempty"`
	LastID               string           `json:"last_id,omitempty"`
	AdditionalProperties map[string]any   `json:"-"`
}

// AnyValue defines model for AnyValue.
type AnyValue = any

// Application defines model for Application.
type Application struct {
	BuildContextDir     string         `json:"buildContextDir,omitempty"`
	BuildImage          string         `json:"buildImage,omitempty"`
	BuildSources        []BuildSource  `json:"buildSources,omitempty"`
	BusinessLineID      string         `json:"businessLineId,omitempty"`
	CreatedAt           time.Time      `json:"createdAt"`
	DefaultBranch       string         `json:"defaultBranch,omitempty"`
	DefaultTag          string         `json:"defaultTag,omitempty"`
	Description         string         `json:"description,omitempty"`
	DockerfilePath      string         `json:"dockerfilePath,omitempty"`
	Enabled             bool           `json:"enabled"`
	EnvironmentCount    int            `json:"environmentCount,omitempty"`
	Group               string         `json:"group"`
	ID                  string         `json:"id"`
	Key                 string         `json:"key"`
	Language            string         `json:"language"`
	Metadata            map[string]any `json:"metadata,omitempty"`
	Name                string         `json:"name"`
	OwnerTeam           string         `json:"ownerTeam,omitempty"`
	RepositoryIDs       []string       `json:"repositoryIds,omitempty"`
	RepositoryPath      string         `json:"repositoryPath,omitempty"`
	RepositoryProjectID string         `json:"repositoryProjectId,omitempty"`
	RepositoryProvider  string         `json:"repositoryProvider,omitempty"`
	UpdatedAt           time.Time      `json:"updatedAt"`
}

// ApplicationBindingSummary defines model for ApplicationBindingSummary.
type ApplicationBindingSummary struct {
	ActionKind               string            `json:"actionKind,omitempty"`
	ApplicationEnvironmentID string            `json:"applicationEnvironmentId"`
	BuildPolicy              *BuildPolicy      `json:"buildPolicy,omitempty"`
	BuildSource              *BuildSource      `json:"buildSource,omitempty"`
	BuildSourceID            string            `json:"buildSourceId,omitempty"`
	EnvironmentID            string            `json:"environmentId"`
	EnvironmentKey           string            `json:"environmentKey,omitempty"`
	EnvironmentName          string            `json:"environmentName,omitempty"`
	LatestBuild              *BuildRecord      `json:"latestBuild,omitempty"`
	LatestBundle             *ReleaseBundle    `json:"latestBundle,omitempty"`
	LatestExecutionTask      *ExecutionTask    `json:"latestExecutionTask,omitempty"`
	LatestRelease            *ReleaseRecord    `json:"latestRelease,omitempty"`
	LatestWorkflow           *WorkflowRun      `json:"latestWorkflow,omitempty"`
	RequiresApproval         bool              `json:"requiresApproval"`
	TargetCount              int               `json:"targetCount"`
	Targets                  []ReleaseTarget   `json:"targets,omitempty"`
	WorkflowTemplate         *WorkflowTemplate `json:"workflowTemplate,omitempty"`
	WorkflowTemplateID       string            `json:"workflowTemplateId,omitempty"`
	WorkflowTemplateName     string            `json:"workflowTemplateName,omitempty"`
}

// ApplicationDeliveryActionKind defines model for ApplicationDeliveryActionKind.
type ApplicationDeliveryActionKind string

// ApplicationDeliveryActionRelatedIDs defines model for ApplicationDeliveryActionRelatedIDs.
type ApplicationDeliveryActionRelatedIDs struct {
	ExecutionTaskID string `json:"executionTaskId,omitempty"`
	ReleaseBundleID string `json:"releaseBundleId,omitempty"`
	WorkflowRunID   string `json:"workflowRunId,omitempty"`
}

// ApplicationDeliveryActionRequest defines model for ApplicationDeliveryActionRequest.
type ApplicationDeliveryActionRequest struct {
	Action                   ApplicationDeliveryActionKind `json:"action"`
	ApplicationEnvironmentID string                        `json:"applicationEnvironmentId"`
	BuildArgs                map[string]any                `json:"buildArgs,omitempty"`
	BuildSourceID            string                        `json:"buildSourceId,omitempty"`
	ContainerName            string                        `json:"containerName,omitempty"`
	ImageTag                 string                        `json:"imageTag,omitempty"`
	RefName                  string                        `json:"refName,omitempty"`
	RefType                  string                        `json:"refType,omitempty"`
	ReleaseBundleID          string                        `json:"releaseBundleId,omitempty"`
	ReleaseName              string                        `json:"releaseName,omitempty"`
	TargetID                 string                        `json:"targetId,omitempty"`
	Variables                map[string]any                `json:"variables,omitempty"`
}

// ApplicationDeliveryActionResult defines model for ApplicationDeliveryActionResult.
type ApplicationDeliveryActionResult struct {
	Action                   ApplicationDeliveryActionKind        `json:"action"`
	ApplicationEnvironmentID string                               `json:"applicationEnvironmentId"`
	ApplicationID            string                               `json:"applicationId"`
	Build                    *BuildRecord                         `json:"build,omitempty"`
	RelatedIDs               *ApplicationDeliveryActionRelatedIDs `json:"relatedIds,omitempty"`
	Release                  *ReleaseRecord                       `json:"release,omitempty"`
	Target                   *ReleaseTarget                       `json:"target,omitempty"`
	Workflow                 *WorkflowRun                         `json:"workflow,omitempty"`
}

// ApplicationDeliveryActionResultEnvelope defines model for ApplicationDeliveryActionResultEnvelope.
type ApplicationDeliveryActionResultEnvelope struct {
	Data ApplicationDeliveryActionResult `json:"data"`
}

// ApplicationDeliveryDetail defines model for ApplicationDeliveryDetail.
type ApplicationDeliveryDetail struct {
	Application         Application                 `json:"application"`
	Bindings            []ApplicationBindingSummary `json:"bindings,omitempty"`
	LatestBuild         *BuildRecord                `json:"latestBuild,omitempty"`
	LatestBundle        *ReleaseBundle              `json:"latestBundle,omitempty"`
	LatestExecutionTask *ExecutionTask              `json:"latestExecutionTask,omitempty"`
	LatestRelease       *ReleaseRecord              `json:"latestRelease,omitempty"`
	LatestWorkflow      *WorkflowRun                `json:"latestWorkflow,omitempty"`
}

// ApplicationDeliveryDetailEnvelope defines model for ApplicationDeliveryDetailEnvelope.
type ApplicationDeliveryDetailEnvelope struct {
	Data ApplicationDeliveryDetail `json:"data"`
}

// ApplicationEnvelope defines model for ApplicationEnvelope.
type ApplicationEnvelope struct {
	Data Application `json:"data"`
}

// ApplicationEnvironment defines model for ApplicationEnvironment.
type ApplicationEnvironment struct {
	ApplicationGroup   string            `json:"applicationGroup,omitempty"`
	ApplicationID      string            `json:"applicationId"`
	ArtifactPolicyID   string            `json:"artifactPolicyId,omitempty"`
	BuildPolicy        *BuildPolicy      `json:"buildPolicy,omitempty"`
	BusinessLineID     string            `json:"businessLineId,omitempty"`
	CreatedAt          time.Time         `json:"createdAt"`
	EnvironmentID      string            `json:"environmentId"`
	EnvironmentKey     string            `json:"environmentKey,omitempty"`
	ID                 string            `json:"id"`
	PromotionPolicyID  string            `json:"promotionPolicyId,omitempty"`
	ReleasePolicy      *ReleasePolicy    `json:"releasePolicy,omitempty"`
	ResourceSelector   *ResourceSelector `json:"resourceSelector,omitempty"`
	StrategyProfileID  string            `json:"strategyProfileId,omitempty"`
	Targets            []ReleaseTarget   `json:"targets,omitempty"`
	UpdatedAt          time.Time         `json:"updatedAt"`
	WorkflowTemplate   *WorkflowTemplate `json:"workflowTemplate,omitempty"`
	WorkflowTemplateID string            `json:"workflowTemplateId,omitempty"`
}

// ApplicationEnvironmentDeliveryDetail defines model for ApplicationEnvironmentDeliveryDetail.
type ApplicationEnvironmentDeliveryDetail struct {
	ActionKind          string                 `json:"actionKind,omitempty"`
	Application         Application            `json:"application"`
	Binding             ApplicationEnvironment `json:"binding"`
	BuildSource         *BuildSource           `json:"buildSource,omitempty"`
	Environment         *DeliveryEnvironment   `json:"environment,omitempty"`
	LatestBuild         *BuildRecord           `json:"latestBuild,omitempty"`
	LatestBundle        *ReleaseBundle         `json:"latestBundle,omitempty"`
	LatestExecutionTask *ExecutionTask         `json:"latestExecutionTask,omitempty"`
	LatestRelease       *ReleaseRecord         `json:"latestRelease,omitempty"`
	LatestWorkflow      *WorkflowRun           `json:"latestWorkflow,omitempty"`
	RequiresApproval    bool                   `json:"requiresApproval"`
}

// ApplicationEnvironmentDeliveryDetailEnvelope defines model for ApplicationEnvironmentDeliveryDetailEnvelope.
type ApplicationEnvironmentDeliveryDetailEnvelope struct {
	Data ApplicationEnvironmentDeliveryDetail `json:"data"`
}

// ApplicationEnvironmentEnvelope defines model for ApplicationEnvironmentEnvelope.
type ApplicationEnvironmentEnvelope struct {
	Data ApplicationEnvironment `json:"data"`
}

// ApplicationEnvironmentInput defines model for ApplicationEnvironmentInput.
type ApplicationEnvironmentInput struct {
	ApplicationID      string               `json:"applicationId"`
	ArtifactPolicyID   string               `json:"artifactPolicyId,omitempty"`
	BuildPolicy        *BuildPolicy         `json:"buildPolicy,omitempty"`
	EnvironmentID      string               `json:"environmentId"`
	ID                 string               `json:"id,omitempty"`
	PromotionPolicyID  string               `json:"promotionPolicyId,omitempty"`
	ReleasePolicy      *ReleasePolicy       `json:"releasePolicy,omitempty"`
	ResourceSelector   *ResourceSelector    `json:"resourceSelector,omitempty"`
	StrategyProfileID  string               `json:"strategyProfileId,omitempty"`
	Targets            []ReleaseTargetInput `json:"targets,omitempty"`
	WorkflowTemplateID string               `json:"workflowTemplateId,omitempty"`
}

// ApplicationEnvironmentListEnvelope defines model for ApplicationEnvironmentListEnvelope.
type ApplicationEnvironmentListEnvelope struct {
	Data []ApplicationEnvironment `json:"data"`
}

// ApplicationInput defines model for ApplicationInput.
type ApplicationInput struct {
	BuildContextDir     string             `json:"buildContextDir,omitempty"`
	BuildImage          string             `json:"buildImage,omitempty"`
	BuildSources        []BuildSourceInput `json:"buildSources,omitempty"`
	BusinessLineID      string             `json:"businessLineId,omitempty"`
	DefaultBranch       string             `json:"defaultBranch,omitempty"`
	DefaultTag          string             `json:"defaultTag,omitempty"`
	Description         string             `json:"description,omitempty"`
	DockerfilePath      string             `json:"dockerfilePath,omitempty"`
	Enabled             bool               `json:"enabled"`
	Group               string             `json:"group"`
	ID                  string             `json:"id,omitempty"`
	Key                 string             `json:"key"`
	Language            string             `json:"language"`
	Metadata            map[string]any     `json:"metadata,omitempty"`
	Name                string             `json:"name"`
	OwnerTeam           string             `json:"ownerTeam,omitempty"`
	RepositoryIDs       []string           `json:"repositoryIds,omitempty"`
	RepositoryPath      string             `json:"repositoryPath,omitempty"`
	RepositoryProjectID string             `json:"repositoryProjectId,omitempty"`
	RepositoryProvider  string             `json:"repositoryProvider,omitempty"`
}

// ApplicationListEnvelope defines model for ApplicationListEnvelope.
type ApplicationListEnvelope struct {
	Data []Application `json:"data"`
}

// ApplicationRuntimeDetail defines model for ApplicationRuntimeDetail.
type ApplicationRuntimeDetail struct {
	Application  Application                     `json:"application"`
	Environments []ApplicationRuntimeEnvironment `json:"environments,omitempty"`
}

// ApplicationRuntimeDetailEnvelope defines model for ApplicationRuntimeDetailEnvelope.
type ApplicationRuntimeDetailEnvelope struct {
	Data ApplicationRuntimeDetail `json:"data"`
}

// ApplicationRuntimeEnvironment defines model for ApplicationRuntimeEnvironment.
type ApplicationRuntimeEnvironment struct {
	ActionKind               string                       `json:"actionKind,omitempty"`
	ApplicationEnvironmentID string                       `json:"applicationEnvironmentId"`
	EnvironmentID            string                       `json:"environmentId"`
	EnvironmentKey           string                       `json:"environmentKey,omitempty"`
	EnvironmentName          string                       `json:"environmentName,omitempty"`
	RequiresApproval         bool                         `json:"requiresApproval"`
	ResourceSelector         *ResourceSelector            `json:"resourceSelector,omitempty"`
	Targets                  []ReleaseTarget              `json:"targets,omitempty"`
	Workloads                []ApplicationRuntimeWorkload `json:"workloads,omitempty"`
}

// ApplicationRuntimeWorkload defines model for ApplicationRuntimeWorkload.
type ApplicationRuntimeWorkload struct {
	ApplicationEnvironmentID string            `json:"applicationEnvironmentId"`
	AvailableReplicas        int               `json:"availableReplicas"`
	BuildSource              *BuildSource      `json:"buildSource,omitempty"`
	ClusterID                string            `json:"clusterId"`
	DesiredReplicas          int               `json:"desiredReplicas"`
	Labels                   map[string]string `json:"labels,omitempty"`
	LatestBuild              *BuildRecord      `json:"latestBuild,omitempty"`
	LatestBundle             *ReleaseBundle    `json:"latestBundle,omitempty"`
	LatestExecutionTask      *ExecutionTask    `json:"latestExecutionTask,omitempty"`
	LatestRelease            *ReleaseRecord    `json:"latestRelease,omitempty"`
	LatestWorkflow           *WorkflowRun      `json:"latestWorkflow,omitempty"`
	Namespace                string            `json:"namespace"`
	ReadyReplicas            int               `json:"readyReplicas"`
	Selector                 map[string]string `json:"selector,omitempty"`
	UpdatedReplicas          int               `json:"updatedReplicas"`
	WorkloadKind             string            `json:"workloadKind"`
	WorkloadName             string            `json:"workloadName"`
}

// ApplicationService defines model for ApplicationService.
type ApplicationService struct {
	ApplicationID       string                        `json:"applicationId"`
	BuildSourceID       string                        `json:"buildSourceId,omitempty"`
	Containers          []ApplicationServiceContainer `json:"containers,omitempty"`
	CreatedAt           time.Time                     `json:"createdAt"`
	DefaultBranch       string                        `json:"defaultBranch,omitempty"`
	Description         string                        `json:"description,omitempty"`
	Enabled             bool                          `json:"enabled"`
	ID                  string                        `json:"id"`
	Key                 string                        `json:"key"`
	Metadata            map[string]any                `json:"metadata,omitempty"`
	Name                string                        `json:"name"`
	OwnerTeam           string                        `json:"ownerTeam,omitempty"`
	RepositoryID        string                        `json:"repositoryId,omitempty"`
	RepositoryPath      string                        `json:"repositoryPath,omitempty"`
	RepositoryProjectID string                        `json:"repositoryProjectId,omitempty"`
	RepositoryProvider  string                        `json:"repositoryProvider,omitempty"`
	ServiceKind         ApplicationServiceServiceKind `json:"serviceKind"`
	UpdatedAt           time.Time                     `json:"updatedAt"`
}

// ApplicationServiceServiceKind defines model for ApplicationService.ServiceKind.
type ApplicationServiceServiceKind string

// ApplicationServiceContainer defines model for ApplicationServiceContainer.
type ApplicationServiceContainer struct {
	BuildContextDir    string         `json:"buildContextDir,omitempty"`
	CreatedAt          *time.Time     `json:"createdAt,omitempty"`
	DefaultTagTemplate string         `json:"defaultTagTemplate,omitempty"`
	DockerfilePath     string         `json:"dockerfilePath,omitempty"`
	EnvSchema          map[string]any `json:"envSchema,omitempty"`
	HealthCheck        map[string]any `json:"healthCheck,omitempty"`
	ID                 string         `json:"id"`
	ImageRepository    string         `json:"imageRepository,omitempty"`
	Metadata           map[string]any `json:"metadata,omitempty"`
	Name               string         `json:"name"`
	ResourceProfile    map[string]any `json:"resourceProfile,omitempty"`
	RuntimePorts       []int          `json:"runtimePorts,omitempty"`
	ServiceID          string         `json:"serviceId,omitempty"`
	UpdatedAt          *time.Time     `json:"updatedAt,omitempty"`
}

// ApplicationServiceContainerInput defines model for ApplicationServiceContainerInput.
type ApplicationServiceContainerInput struct {
	BuildContextDir    string         `json:"buildContextDir,omitempty"`
	DefaultTagTemplate string         `json:"defaultTagTemplate,omitempty"`
	DockerfilePath     string         `json:"dockerfilePath,omitempty"`
	EnvSchema          map[string]any `json:"envSchema,omitempty"`
	HealthCheck        map[string]any `json:"healthCheck,omitempty"`
	ID                 string         `json:"id,omitempty"`
	ImageRepository    string         `json:"imageRepository,omitempty"`
	Metadata           map[string]any `json:"metadata,omitempty"`
	Name               string         `json:"name"`
	ResourceProfile    map[string]any `json:"resourceProfile,omitempty"`
	RuntimePorts       []int          `json:"runtimePorts,omitempty"`
}

// ApplicationServiceEnvelope defines model for ApplicationServiceEnvelope.
type ApplicationServiceEnvelope struct {
	Data ApplicationService `json:"data"`
}

// ApplicationServiceInput defines model for ApplicationServiceInput.
type ApplicationServiceInput struct {
	BuildSourceID       string                             `json:"buildSourceId,omitempty"`
	Containers          []ApplicationServiceContainerInput `json:"containers,omitempty"`
	DefaultBranch       string                             `json:"defaultBranch,omitempty"`
	Description         string                             `json:"description,omitempty"`
	Enabled             bool                               `json:"enabled"`
	ID                  string                             `json:"id,omitempty"`
	Key                 string                             `json:"key"`
	Metadata            map[string]any                     `json:"metadata,omitempty"`
	Name                string                             `json:"name"`
	OwnerTeam           string                             `json:"ownerTeam,omitempty"`
	RepositoryID        string                             `json:"repositoryId,omitempty"`
	RepositoryPath      string                             `json:"repositoryPath,omitempty"`
	RepositoryProjectID string                             `json:"repositoryProjectId,omitempty"`
	RepositoryProvider  string                             `json:"repositoryProvider,omitempty"`
	ServiceKind         ApplicationServiceInputServiceKind `json:"serviceKind"`
}

// ApplicationServiceInputServiceKind defines model for ApplicationServiceInput.ServiceKind.
type ApplicationServiceInputServiceKind string

// ApplicationServiceListEnvelope defines model for ApplicationServiceListEnvelope.
type ApplicationServiceListEnvelope struct {
	Data []ApplicationService `json:"data"`
}

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
	IconURL  string `json:"iconUrl,omitempty"`
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

// BuildPolicy defines model for BuildPolicy.
type BuildPolicy struct {
	BuildArgs        map[string]any `json:"buildArgs,omitempty"`
	ImageTagMode     string         `json:"imageTagMode,omitempty"`
	ImageTagTemplate string         `json:"imageTagTemplate,omitempty"`
	RefType          string         `json:"refType,omitempty"`
	RefValue         string         `json:"refValue,omitempty"`
	SourceID         string         `json:"sourceId,omitempty"`
	Variables        map[string]any `json:"variables,omitempty"`
}

// BuildRecord defines model for BuildRecord.
type BuildRecord struct {
	ApplicationID string         `json:"applicationId"`
	CreatedAt     time.Time      `json:"createdAt"`
	FinishedAt    *time.Time     `json:"finishedAt,omitempty"`
	ID            string         `json:"id"`
	Metadata      map[string]any `json:"metadata,omitempty"`
	SourceSystem  string         `json:"sourceSystem"`
	StartedAt     *time.Time     `json:"startedAt,omitempty"`
	Status        string         `json:"status"`
}

// BuildSource defines model for BuildSource.
type BuildSource struct {
	BuildImage string             `json:"buildImage,omitempty"`
	Config     *BuildSourceConfig `json:"config,omitempty"`
	CreatedAt  *time.Time         `json:"createdAt,omitempty"`
	DefaultTag string             `json:"defaultTag,omitempty"`
	Enabled    bool               `json:"enabled"`
	ID         string             `json:"id"`
	IsDefault  bool               `json:"isDefault"`
	Name       string             `json:"name"`
	Type       BuildSourceType    `json:"type"`
	UpdatedAt  *time.Time         `json:"updatedAt,omitempty"`
}

// BuildSourceType defines model for BuildSource.Type.
type BuildSourceType string

// BuildSourceConfig defines model for BuildSourceConfig.
type BuildSourceConfig struct {
	BuildArgs       map[string]BuildSourceConfig_BuildArgs_AdditionalProperties `json:"buildArgs,omitempty"`
	BuildImage      string                                                      `json:"buildImage,omitempty"`
	BuildTemplateID string                                                      `json:"buildTemplateId,omitempty"`
	BuilderKind     BuildSourceConfigBuilderKind                                `json:"builderKind,omitempty"`
	ContextDir      string                                                      `json:"contextDir,omitempty"`
	DefaultTag      string                                                      `json:"defaultTag,omitempty"`
	DockerfilePath  string                                                      `json:"dockerfilePath,omitempty"`
	PipelineRef     string                                                      `json:"pipelineRef,omitempty"`
	ProviderKind    string                                                      `json:"providerKind,omitempty"`
	RepositoryID    string                                                      `json:"repositoryId,omitempty"`
	Variables       map[string]BuildSourceConfig_Variables_AdditionalProperties `json:"variables,omitempty"`
}

// BuildSourceConfigBuildArgs0 defines model for .
type BuildSourceConfigBuildArgs0 = string

// BuildSourceConfigBuildArgs1 defines model for .
type BuildSourceConfigBuildArgs1 = float32

// BuildSourceConfigBuildArgs2 defines model for .
type BuildSourceConfigBuildArgs2 = bool

// BuildSourceConfig_BuildArgs_AdditionalProperties defines model for BuildSourceConfig.buildArgs.AdditionalProperties.
type BuildSourceConfig_BuildArgs_AdditionalProperties struct {
	union json.RawMessage
}

// BuildSourceConfigBuilderKind defines model for BuildSourceConfig.BuilderKind.
type BuildSourceConfigBuilderKind string

// BuildSourceConfigVariables0 defines model for .
type BuildSourceConfigVariables0 = string

// BuildSourceConfigVariables1 defines model for .
type BuildSourceConfigVariables1 = float32

// BuildSourceConfigVariables2 defines model for .
type BuildSourceConfigVariables2 = bool

// BuildSourceConfig_Variables_AdditionalProperties defines model for BuildSourceConfig.variables.AdditionalProperties.
type BuildSourceConfig_Variables_AdditionalProperties struct {
	union json.RawMessage
}

// BuildSourceInput defines model for BuildSourceInput.
type BuildSourceInput struct {
	BuildImage string               `json:"buildImage,omitempty"`
	Config     *BuildSourceConfig   `json:"config,omitempty"`
	DefaultTag string               `json:"defaultTag,omitempty"`
	Enabled    bool                 `json:"enabled"`
	ID         string               `json:"id,omitempty"`
	IsDefault  bool                 `json:"isDefault"`
	Name       string               `json:"name"`
	Type       BuildSourceInputType `json:"type"`
}

// BuildSourceInputType defines model for BuildSourceInput.Type.
type BuildSourceInputType string

// BuildTemplate defines model for BuildTemplate.
type BuildTemplate struct {
	BuildCommands      []string       `json:"buildCommands,omitempty"`
	BuilderKind        string         `json:"builderKind,omitempty"`
	CreatedAt          time.Time      `json:"createdAt"`
	DefaultVariables   map[string]any `json:"defaultVariables,omitempty"`
	Description        string         `json:"description,omitempty"`
	DockerfileTemplate string         `json:"dockerfileTemplate,omitempty"`
	Enabled            bool           `json:"enabled"`
	ID                 string         `json:"id"`
	Key                string         `json:"key"`
	Name               string         `json:"name"`
	UpdatedAt          time.Time      `json:"updatedAt"`
	VariableSchema     map[string]any `json:"variableSchema,omitempty"`
}

// BuildTemplateEnvelope defines model for BuildTemplateEnvelope.
type BuildTemplateEnvelope struct {
	Data BuildTemplate `json:"data"`
}

// BuildTemplateInput defines model for BuildTemplateInput.
type BuildTemplateInput struct {
	BuildCommands      []string       `json:"buildCommands,omitempty"`
	BuilderKind        string         `json:"builderKind,omitempty"`
	DefaultVariables   map[string]any `json:"defaultVariables,omitempty"`
	Description        string         `json:"description,omitempty"`
	DockerfileTemplate string         `json:"dockerfileTemplate,omitempty"`
	Enabled            bool           `json:"enabled"`
	ID                 string         `json:"id,omitempty"`
	Key                string         `json:"key"`
	Name               string         `json:"name"`
	VariableSchema     map[string]any `json:"variableSchema,omitempty"`
}

// BuildTemplateListEnvelope defines model for BuildTemplateListEnvelope.
type BuildTemplateListEnvelope struct {
	Data []BuildTemplate `json:"data"`
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

// CohereRerankRequest defines model for CohereRerankRequest.
type CohereRerankRequest struct {
	Documents       []CohereRerankRequest_Documents_Item `json:"documents"`
	MaxTokensPerDoc int                                  `json:"max_tokens_per_doc,omitempty"`

	// Model Public rerank model ID exposed by Soha; the relay maps it to the configured upstream model.
	Model                string         `json:"model"`
	Query                string         `json:"query"`
	RankFields           []string       `json:"rank_fields,omitempty"`
	TopN                 int            `json:"top_n,omitempty"`
	AdditionalProperties map[string]any `json:"-"`
}

// CohereRerankRequestDocuments0 defines model for .
type CohereRerankRequestDocuments0 = string

// CohereRerankRequestDocuments1 defines model for .
type CohereRerankRequestDocuments1 map[string]any

// CohereRerankRequest_Documents_Item defines model for CohereRerankRequest.documents.Item.
type CohereRerankRequest_Documents_Item struct {
	union json.RawMessage
}

// ComputeAccessMode defines model for ComputeAccessMode.
type ComputeAccessMode string

// ComputeAccessSource defines model for ComputeAccessSource.
type ComputeAccessSource struct {
	AccessMode         ComputeAccessMode       `json:"accessMode,omitempty"`
	AvailableActions   []string                `json:"availableActions,omitempty"`
	ID                 string                  `json:"id"`
	LastObservedAt     *time.Time              `json:"lastObservedAt,omitempty"`
	PluginID           string                  `json:"pluginId,omitempty"`
	PluginVersion      string                  `json:"pluginVersion,omitempty"`
	ProviderGeneration int64                   `json:"providerGeneration,omitempty"`
	ProviderKey        string                  `json:"providerKey,omitempty"`
	ProviderSource     ComputeProviderSource   `json:"providerSource,omitempty"`
	RelatedResources   []ComputeResourceRef    `json:"relatedResources,omitempty"`
	Resource           ComputeResourceRef      `json:"resource"`
	SourceType         ComputeAccessSourceType `json:"sourceType"`
	Status             ComputeHealthStatus     `json:"status"`
}

// ComputeAccessSourceListEnvelope defines model for ComputeAccessSourceListEnvelope.
type ComputeAccessSourceListEnvelope struct {
	Items      []ComputeAccessSource `json:"items"`
	NextCursor string                `json:"nextCursor,omitempty"`
}

// ComputeAccessSourceType defines model for ComputeAccessSourceType.
type ComputeAccessSourceType string

// ComputeAgentOverviewSection defines model for ComputeAgentOverviewSection.
type ComputeAgentOverviewSection struct {
	Status   ComputeSectionStatus `json:"status"`
	Summary  *ComputeAgentSummary `json:"summary,omitempty"`
	Warnings []ComputeWarning     `json:"warnings,omitempty"`
}

// ComputeAgentSummary defines model for ComputeAgentSummary.
type ComputeAgentSummary struct {
	Offline         int `json:"offline"`
	Online          int `json:"online"`
	Total           int `json:"total"`
	VersionMismatch int `json:"versionMismatch"`
}

// ComputeAttention defines model for ComputeAttention.
type ComputeAttention struct {
	Code      string                   `json:"code"`
	Resources []ComputeResourceRef     `json:"resources,omitempty"`
	Severity  ComputeAttentionSeverity `json:"severity"`
	Summary   string                   `json:"summary"`
}

// ComputeAttentionSeverity defines model for ComputeAttention.Severity.
type ComputeAttentionSeverity string

// ComputeCapabilityManifest defines model for ComputeCapabilityManifest.
type ComputeCapabilityManifest struct {
	Features   []ComputeFeatureCapability `json:"features"`
	Generation int64                      `json:"generation"`
}

// ComputeCapabilityManifestEnvelope defines model for ComputeCapabilityManifestEnvelope.
type ComputeCapabilityManifestEnvelope struct {
	Data ComputeCapabilityManifest `json:"data"`
}

// ComputeDomain defines model for ComputeDomain.
type ComputeDomain string

// ComputeFeatureCapability defines model for ComputeFeatureCapability.
type ComputeFeatureCapability struct {
	Enabled            bool                           `json:"enabled"`
	ID                 ComputeFeatureID               `json:"id"`
	MaxActivationLevel ComputeProviderActivationLevel `json:"maxActivationLevel,omitempty"`

	// Reason Redacted explanation when a capability is unavailable to the current principal.
	Reason       string              `json:"reason,omitempty"`
	RolloutStage ComputeRolloutStage `json:"rolloutStage"`
}

// ComputeFeatureID defines model for ComputeFeatureID.
type ComputeFeatureID string

// ComputeHealthStatus defines model for ComputeHealthStatus.
type ComputeHealthStatus string

// ComputeMetadataEntry defines model for ComputeMetadataEntry.
type ComputeMetadataEntry struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// ComputeOverview defines model for ComputeOverview.
type ComputeOverview struct {
	Agents           *ComputeAgentOverviewSection           `json:"agents,omitempty"`
	Attention        []ComputeAttention                     `json:"attention"`
	Partial          bool                                   `json:"partial"`
	ProviderHealth   []ComputeProviderHealth                `json:"providerHealth"`
	RuntimeWorkloads *ComputeRuntimeWorkloadOverviewSection `json:"runtimeWorkloads,omitempty"`
	Runtimes         *ComputeRuntimeOverviewSection         `json:"runtimes,omitempty"`
	Tasks            *ComputeTaskOverviewSection            `json:"tasks,omitempty"`
	Virtualization   *ComputeVirtualizationOverviewSection  `json:"virtualization,omitempty"`
	Warnings         []ComputeWarning                       `json:"warnings"`
}

// ComputeOverviewEnvelope defines model for ComputeOverviewEnvelope.
type ComputeOverviewEnvelope struct {
	Data ComputeOverview `json:"data"`
}

// ComputePluginRuntimeMode defines model for ComputePluginRuntimeMode.
type ComputePluginRuntimeMode string

// ComputeProviderActivationLevel defines model for ComputeProviderActivationLevel.
type ComputeProviderActivationLevel string

// ComputeProviderCapability defines model for ComputeProviderCapability.
type ComputeProviderCapability struct {
	Enabled       bool                           `json:"enabled"`
	ID            string                         `json:"id"`
	Level         ComputeProviderActivationLevel `json:"level"`
	Reason        string                         `json:"reason,omitempty"`
	ResourceKinds []ComputeResourceKind          `json:"resourceKinds"`
}

// ComputeProviderDescriptor defines model for ComputeProviderDescriptor.
type ComputeProviderDescriptor struct {
	ActivationLevel ComputeProviderActivationLevel  `json:"activationLevel"`
	Capabilities    []ComputeProviderCapability     `json:"capabilities"`
	ConfigSchema    JSONSchema                      `json:"configSchema,omitempty"`
	ContractVersion string                          `json:"contractVersion"`
	DisplayName     string                          `json:"displayName"`
	Domain          ComputeProviderDomain           `json:"domain"`
	Generation      int64                           `json:"generation"`
	Health          ComputeProviderHealth           `json:"health"`
	PluginID        string                          `json:"pluginId,omitempty"`
	PluginVersion   string                          `json:"pluginVersion,omitempty"`
	ProviderKey     string                          `json:"providerKey"`
	ResourceKinds   []ComputeResourceKind           `json:"resourceKinds"`
	ResourceSchemas []ComputeProviderResourceSchema `json:"resourceSchemas,omitempty"`
	RuntimeMode     ComputePluginRuntimeMode        `json:"runtimeMode"`
	Source          ComputeProviderSource           `json:"source"`
	StatusMappings  []ComputeProviderStatusMapping  `json:"statusMappings,omitempty"`
	Version         string                          `json:"version"`
}

// ComputeProviderDiscoverRequest defines model for ComputeProviderDiscoverRequest.
type ComputeProviderDiscoverRequest struct {
	Cursor             string `json:"cursor,omitempty"`
	ExpectedGeneration int64  `json:"expectedGeneration"`
	MaxItems           int    `json:"maxItems,omitempty"`
	Scope              string `json:"scope,omitempty"`
}

// ComputeProviderDomain defines model for ComputeProviderDomain.
type ComputeProviderDomain string

// ComputeProviderHealth defines model for ComputeProviderHealth.
type ComputeProviderHealth struct {
	CheckedAt  *time.Time            `json:"checkedAt,omitempty"`
	Code       string                `json:"code,omitempty"`
	Domain     ComputeProviderDomain `json:"domain"`
	Generation int64                 `json:"generation"`

	// Message Redacted health summary.
	Message     string              `json:"message,omitempty"`
	ProviderKey string              `json:"providerKey"`
	Status      ComputeHealthStatus `json:"status"`
}

// ComputeProviderInstance defines model for ComputeProviderInstance.
type ComputeProviderInstance struct {
	AccessMode            ComputeAccessMode           `json:"accessMode"`
	DisplayName           string                      `json:"displayName"`
	EffectiveCapabilities []ComputeProviderCapability `json:"effectiveCapabilities"`
	Enabled               bool                        `json:"enabled"`
	Health                ComputeProviderHealth       `json:"health"`
	InstanceRef           string                      `json:"instanceRef"`
	LastObservedAt        *time.Time                  `json:"lastObservedAt,omitempty"`
	Resource              *ComputeResourceRef         `json:"resource,omitempty"`
	Snapshot              ComputeProviderSnapshot     `json:"snapshot"`
}

// ComputeProviderInstanceEnvelope defines model for ComputeProviderInstanceEnvelope.
type ComputeProviderInstanceEnvelope struct {
	Data ComputeProviderInstance `json:"data"`
}

// ComputeProviderInstanceListEnvelope defines model for ComputeProviderInstanceListEnvelope.
type ComputeProviderInstanceListEnvelope struct {
	Items      []ComputeProviderInstance `json:"items"`
	NextCursor string                    `json:"nextCursor,omitempty"`
}

// ComputeProviderListEnvelope defines model for ComputeProviderListEnvelope.
type ComputeProviderListEnvelope struct {
	Items      []ComputeProviderDescriptor `json:"items"`
	NextCursor string                      `json:"nextCursor,omitempty"`
}

// ComputeProviderReadRequest defines model for ComputeProviderReadRequest.
type ComputeProviderReadRequest struct {
	ExpectedGeneration int64  `json:"expectedGeneration"`
	Scope              string `json:"scope,omitempty"`
}

// ComputeProviderResourceSchema defines model for ComputeProviderResourceSchema.
type ComputeProviderResourceSchema struct {
	Kind   ComputeResourceKind `json:"kind"`
	Schema JSONSchema          `json:"schema"`
}

// ComputeProviderSnapshot defines model for ComputeProviderSnapshot.
type ComputeProviderSnapshot struct {
	ContractVersion string                   `json:"contractVersion"`
	Domain          ComputeProviderDomain    `json:"domain"`
	Generation      int64                    `json:"generation"`
	PluginID        string                   `json:"pluginId,omitempty"`
	PluginVersion   string                   `json:"pluginVersion,omitempty"`
	ProviderKey     string                   `json:"providerKey"`
	RuntimeMode     ComputePluginRuntimeMode `json:"runtimeMode"`
	Source          ComputeProviderSource    `json:"source"`
	Version         string                   `json:"version"`
}

// ComputeProviderSource defines model for ComputeProviderSource.
type ComputeProviderSource string

// ComputeProviderStatusMapping defines model for ComputeProviderStatusMapping.
type ComputeProviderStatusMapping struct {
	NormalizedStatus string `json:"normalizedStatus"`
	ProviderStatus   string `json:"providerStatus"`
}

// ComputeRelationSource defines model for ComputeRelationSource.
type ComputeRelationSource string

// ComputeRelationType defines model for ComputeRelationType.
type ComputeRelationType string

// ComputeResourceActionRequest defines model for ComputeResourceActionRequest.
type ComputeResourceActionRequest struct {
	Metadata []ComputeMetadataEntry `json:"metadata,omitempty"`
	Reason   string                 `json:"reason,omitempty"`
}

// ComputeResourceKind defines model for ComputeResourceKind.
type ComputeResourceKind string

// ComputeResourceRef defines model for ComputeResourceRef.
type ComputeResourceRef struct {
	AccessMode          ComputeAccessMode     `json:"accessMode,omitempty"`
	DisplayName         string                `json:"displayName"`
	Domain              ComputeDomain         `json:"domain"`
	ID                  string                `json:"id"`
	Kind                ComputeResourceKind   `json:"kind"`
	PluginID            string                `json:"pluginId,omitempty"`
	PluginVersion       string                `json:"pluginVersion,omitempty"`
	ProviderGeneration  int64                 `json:"providerGeneration,omitempty"`
	ProviderInstanceRef string                `json:"providerInstanceRef,omitempty"`
	ProviderKey         string                `json:"providerKey,omitempty"`
	ProviderSource      ComputeProviderSource `json:"providerSource,omitempty"`
	Scope               string                `json:"scope,omitempty"`
}

// ComputeResourceRelation defines model for ComputeResourceRelation.
type ComputeResourceRelation struct {
	From               ComputeResourceRef     `json:"from"`
	Metadata           []ComputeMetadataEntry `json:"metadata,omitempty"`
	ObservedAt         time.Time              `json:"observedAt"`
	ProviderGeneration int64                  `json:"providerGeneration,omitempty"`
	Source             ComputeRelationSource  `json:"source"`
	Stale              bool                   `json:"stale,omitempty"`
	To                 ComputeResourceRef     `json:"to"`
	Type               ComputeRelationType    `json:"type"`
}

// ComputeResourceRelationListEnvelope defines model for ComputeResourceRelationListEnvelope.
type ComputeResourceRelationListEnvelope struct {
	Data ComputeResourceRelations `json:"data"`
}

// ComputeResourceRelations defines model for ComputeResourceRelations.
type ComputeResourceRelations struct {
	NextCursor string                    `json:"nextCursor,omitempty"`
	Relations  []ComputeResourceRelation `json:"relations"`
	Resource   ComputeResourceRef        `json:"resource"`
}

// ComputeRolloutStage defines model for ComputeRolloutStage.
type ComputeRolloutStage string

// ComputeRuntimeOverviewSection defines model for ComputeRuntimeOverviewSection.
type ComputeRuntimeOverviewSection struct {
	Status   ComputeSectionStatus   `json:"status"`
	Summary  *ComputeRuntimeSummary `json:"summary,omitempty"`
	Warnings []ComputeWarning       `json:"warnings,omitempty"`
}

// ComputeRuntimeSummary defines model for ComputeRuntimeSummary.
type ComputeRuntimeSummary struct {
	Available    int `json:"available"`
	Error        int `json:"error"`
	Total        int `json:"total"`
	WaitingAgent int `json:"waitingAgent"`
}

// ComputeRuntimeWorkloadOverviewSection defines model for ComputeRuntimeWorkloadOverviewSection.
type ComputeRuntimeWorkloadOverviewSection struct {
	Status   ComputeSectionStatus           `json:"status"`
	Summary  *ComputeRuntimeWorkloadSummary `json:"summary,omitempty"`
	Warnings []ComputeWarning               `json:"warnings,omitempty"`
}

// ComputeRuntimeWorkloadSummary defines model for ComputeRuntimeWorkloadSummary.
type ComputeRuntimeWorkloadSummary struct {
	Containers int `json:"containers"`
	Expiring   int `json:"expiring"`
	Ports      int `json:"ports"`
	Projects   int `json:"projects"`
	Services   int `json:"services"`
}

// ComputeSectionStatus defines model for ComputeSectionStatus.
type ComputeSectionStatus string

// ComputeTaskAction defines model for ComputeTaskAction.
type ComputeTaskAction string

// ComputeTaskCategory defines model for ComputeTaskCategory.
type ComputeTaskCategory string

// ComputeTaskDomain defines model for ComputeTaskDomain.
type ComputeTaskDomain string

// ComputeTaskEnvelope defines model for ComputeTaskEnvelope.
type ComputeTaskEnvelope struct {
	Data ComputeTaskView `json:"data"`
}

// ComputeTaskListEnvelope defines model for ComputeTaskListEnvelope.
type ComputeTaskListEnvelope struct {
	Items      []ComputeTaskView `json:"items"`
	NextCursor string            `json:"nextCursor,omitempty"`
}

// ComputeTaskLog defines model for ComputeTaskLog.
type ComputeTaskLog struct {
	CreatedAt time.Time `json:"createdAt"`
	ID        string    `json:"id"`
	LogLevel  string    `json:"logLevel"`
	Message   string    `json:"message"`

	// Payload Compact JSON representation of optional source-domain log metadata.
	Payload string `json:"payload,omitempty"`
	TaskID  string `json:"taskId"`
}

// ComputeTaskLogListEnvelope defines model for ComputeTaskLogListEnvelope.
type ComputeTaskLogListEnvelope struct {
	Items []ComputeTaskLog `json:"items"`
}

// ComputeTaskMutationRequest defines model for ComputeTaskMutationRequest.
type ComputeTaskMutationRequest struct {
	Reason string `json:"reason,omitempty"`
}

// ComputeTaskOverviewSection defines model for ComputeTaskOverviewSection.
type ComputeTaskOverviewSection struct {
	Status   ComputeSectionStatus `json:"status"`
	Summary  *ComputeTaskSummary  `json:"summary,omitempty"`
	Warnings []ComputeWarning     `json:"warnings,omitempty"`
}

// ComputeTaskStatus defines model for ComputeTaskStatus.
type ComputeTaskStatus string

// ComputeTaskSummary defines model for ComputeTaskSummary.
type ComputeTaskSummary struct {
	Failed  int `json:"failed"`
	Queued  int `json:"queued"`
	Running int `json:"running"`
}

// ComputeTaskView defines model for ComputeTaskView.
type ComputeTaskView struct {
	AttemptCount       int                   `json:"attemptCount"`
	AvailableActions   []ComputeTaskAction   `json:"availableActions"`
	Cancelable         bool                  `json:"cancelable"`
	Category           ComputeTaskCategory   `json:"category"`
	CreatedAt          time.Time             `json:"createdAt"`
	Domain             ComputeTaskDomain     `json:"domain"`
	ErrorCode          string                `json:"errorCode,omitempty"`
	FinishedAt         *time.Time            `json:"finishedAt,omitempty"`
	ID                 string                `json:"id"`
	Kind               string                `json:"kind"`
	NormalizedStatus   ComputeTaskStatus     `json:"normalizedStatus"`
	PluginID           string                `json:"pluginId,omitempty"`
	PluginVersion      string                `json:"pluginVersion,omitempty"`
	ProviderGeneration int64                 `json:"providerGeneration,omitempty"`
	ProviderKey        string                `json:"providerKey,omitempty"`
	ProviderSource     ComputeProviderSource `json:"providerSource,omitempty"`
	RawStatus          string                `json:"rawStatus"`
	RequestedBy        string                `json:"requestedBy,omitempty"`
	Resources          []ComputeResourceRef  `json:"resources"`
	Retryable          bool                  `json:"retryable"`
	SourceID           string                `json:"sourceId"`
	SourceType         string                `json:"sourceType"`
	StartedAt          *time.Time            `json:"startedAt,omitempty"`
	Summary            string                `json:"summary,omitempty"`
	Worker             string                `json:"worker,omitempty"`
}

// ComputeVirtualizationOverviewSection defines model for ComputeVirtualizationOverviewSection.
type ComputeVirtualizationOverviewSection struct {
	Status   ComputeSectionStatus          `json:"status"`
	Summary  *ComputeVirtualizationSummary `json:"summary,omitempty"`
	Warnings []ComputeWarning              `json:"warnings,omitempty"`
}

// ComputeVirtualizationSummary defines model for ComputeVirtualizationSummary.
type ComputeVirtualizationSummary struct {
	ConnectionsDegraded int `json:"connectionsDegraded"`
	ConnectionsHealthy  int `json:"connectionsHealthy"`
	ConnectionsTotal    int `json:"connectionsTotal"`
	ConnectionsUnsynced int `json:"connectionsUnsynced"`
	VmsError            int `json:"vmsError"`
	VmsRunning          int `json:"vmsRunning"`
	VmsStopped          int `json:"vmsStopped"`
	VmsTotal            int `json:"vmsTotal"`
}

// ComputeWarning defines model for ComputeWarning.
type ComputeWarning struct {
	Code string `json:"code"`

	// Message Redacted user-facing warning. Internal errors and credentials are excluded.
	Message string `json:"message,omitempty"`
}

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

// DeliveryEnvironment defines model for DeliveryEnvironment.
type DeliveryEnvironment struct {
	CreatedAt        time.Time `json:"createdAt"`
	Enabled          bool      `json:"enabled"`
	ID               string    `json:"id"`
	IsProduction     bool      `json:"isProduction"`
	Key              string    `json:"key"`
	Name             string    `json:"name"`
	RequiresApproval bool      `json:"requiresApproval"`
	SortOrder        int       `json:"sortOrder"`
	StageLevel       int       `json:"stageLevel"`
	Tier             string    `json:"tier,omitempty"`
	UpdatedAt        time.Time `json:"updatedAt"`
}

// DeliveryPlan defines model for DeliveryPlan.
type DeliveryPlan struct {
	Action                   ApplicationDeliveryActionKind `json:"action"`
	ApplicationEnvironmentID string                        `json:"applicationEnvironmentId"`
	ApplicationID            string                        `json:"applicationId"`
	ApplicationName          string                        `json:"applicationName,omitempty"`
	BuildArgs                map[string]any                `json:"buildArgs,omitempty"`
	BuildSourceID            string                        `json:"buildSourceId,omitempty"`
	ConfirmedAt              *time.Time                    `json:"confirmedAt,omitempty"`
	ContainerName            string                        `json:"containerName,omitempty"`
	CreatedAt                time.Time                     `json:"createdAt"`
	CreatedBy                string                        `json:"createdBy,omitempty"`
	EnvironmentKey           string                        `json:"environmentKey,omitempty"`
	ID                       string                        `json:"id"`
	ImageTag                 string                        `json:"imageTag,omitempty"`
	Impact                   map[string]any                `json:"impact,omitempty"`
	Reason                   string                        `json:"reason,omitempty"`
	RefName                  string                        `json:"refName,omitempty"`
	RefType                  string                        `json:"refType,omitempty"`
	ReleaseBundleID          string                        `json:"releaseBundleId,omitempty"`
	ReleaseName              string                        `json:"releaseName,omitempty"`
	RequiresApproval         bool                          `json:"requiresApproval"`
	RiskLevel                string                        `json:"riskLevel,omitempty"`
	RollbackStrategy         string                        `json:"rollbackStrategy,omitempty"`
	Source                   DeliveryPlanSource            `json:"source"`
	Status                   DeliveryPlanStatus            `json:"status"`
	TargetID                 string                        `json:"targetId,omitempty"`
	TargetSummary            string                        `json:"targetSummary,omitempty"`
	UpdatedAt                time.Time                     `json:"updatedAt"`
	Variables                map[string]any                `json:"variables,omitempty"`
}

// DeliveryPlanSource defines model for DeliveryPlan.Source.
type DeliveryPlanSource string

// DeliveryPlanStatus defines model for DeliveryPlan.Status.
type DeliveryPlanStatus string

// DeliveryPlanConfirmResult defines model for DeliveryPlanConfirmResult.
type DeliveryPlanConfirmResult struct {
	Plan   DeliveryPlan                    `json:"plan"`
	Result ApplicationDeliveryActionResult `json:"result"`
}

// DeliveryPlanConfirmResultEnvelope defines model for DeliveryPlanConfirmResultEnvelope.
type DeliveryPlanConfirmResultEnvelope struct {
	Data DeliveryPlanConfirmResult `json:"data"`
}

// DeliveryPlanEnvelope defines model for DeliveryPlanEnvelope.
type DeliveryPlanEnvelope struct {
	Data DeliveryPlan `json:"data"`
}

// DeliveryPlanInput defines model for DeliveryPlanInput.
type DeliveryPlanInput struct {
	Action                   ApplicationDeliveryActionKind `json:"action"`
	ApplicationEnvironmentID string                        `json:"applicationEnvironmentId"`
	ApplicationID            string                        `json:"applicationId"`
	ApplicationName          string                        `json:"applicationName,omitempty"`
	BuildArgs                map[string]any                `json:"buildArgs,omitempty"`
	BuildSourceID            string                        `json:"buildSourceId,omitempty"`
	ContainerName            string                        `json:"containerName,omitempty"`
	EnvironmentKey           string                        `json:"environmentKey,omitempty"`
	ID                       string                        `json:"id,omitempty"`
	ImageTag                 string                        `json:"imageTag,omitempty"`
	Impact                   map[string]any                `json:"impact,omitempty"`
	Reason                   string                        `json:"reason,omitempty"`
	RefName                  string                        `json:"refName,omitempty"`
	RefType                  string                        `json:"refType,omitempty"`
	ReleaseBundleID          string                        `json:"releaseBundleId,omitempty"`
	ReleaseName              string                        `json:"releaseName,omitempty"`
	RequiresApproval         bool                          `json:"requiresApproval,omitempty"`
	RiskLevel                string                        `json:"riskLevel,omitempty"`
	RollbackStrategy         string                        `json:"rollbackStrategy,omitempty"`
	Source                   DeliveryPlanInputSource       `json:"source,omitempty"`
	TargetID                 string                        `json:"targetId,omitempty"`
	TargetSummary            string                        `json:"targetSummary,omitempty"`
	Variables                map[string]any                `json:"variables,omitempty"`
}

// DeliveryPlanInputSource defines model for DeliveryPlanInput.Source.
type DeliveryPlanInputSource string

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

// EvaluationCompleteRunInput defines model for EvaluationCompleteRunInput.
type EvaluationCompleteRunInput struct {
	Outputs []EvaluationSampleOutput `json:"outputs"`
}

// EvaluationDataset defines model for EvaluationDataset.
type EvaluationDataset struct {
	CreatedAt     time.Time                 `json:"createdAt"`
	ID            string                    `json:"id"`
	Name          string                    `json:"name"`
	Samples       []EvaluationDatasetSample `json:"samples"`
	SchemaVersion string                    `json:"schemaVersion"`
	Version       string                    `json:"version"`
}

// EvaluationDatasetEnvelope defines model for EvaluationDatasetEnvelope.
type EvaluationDatasetEnvelope struct {
	Data EvaluationDataset `json:"data"`
}

// EvaluationDatasetListEnvelope defines model for EvaluationDatasetListEnvelope.
type EvaluationDatasetListEnvelope struct {
	Items []EvaluationDataset `json:"items"`
}

// EvaluationDatasetSample defines model for EvaluationDatasetSample.
type EvaluationDatasetSample struct {
	ExpectedFacts    []string `json:"expectedFacts,omitempty"`
	ExpectedSources  []string `json:"expectedSources,omitempty"`
	ForbiddenActions []string `json:"forbiddenActions,omitempty"`
	ID               string   `json:"id"`
	Input            string   `json:"input"`
}

// EvaluationExecuteRunInput defines model for EvaluationExecuteRunInput.
type EvaluationExecuteRunInput struct {
	ExecutorProfileID string `json:"executorProfileId"`
}

// EvaluationExecutorProfile defines model for EvaluationExecutorProfile.
type EvaluationExecutorProfile struct {
	EnvironmentPolicy string                                 `json:"environmentPolicy"`
	ID                string                                 `json:"id"`
	IsolationMode     EvaluationExecutorProfileIsolationMode `json:"isolationMode"`
	MaxCost           float32                                `json:"maxCost"`

	// Timeout Go time.Duration encoded as nanoseconds.
	Timeout         int64  `json:"timeout"`
	ToolSnapshotRef string `json:"toolSnapshotRef,omitempty"`
}

// EvaluationExecutorProfileIsolationMode defines model for EvaluationExecutorProfile.IsolationMode.
type EvaluationExecutorProfileIsolationMode string

// EvaluationExecutorProfileEnvelope defines model for EvaluationExecutorProfileEnvelope.
type EvaluationExecutorProfileEnvelope struct {
	Data EvaluationExecutorProfile `json:"data"`
}

// EvaluationExecutorProfileListEnvelope defines model for EvaluationExecutorProfileListEnvelope.
type EvaluationExecutorProfileListEnvelope struct {
	Items []EvaluationExecutorProfile `json:"items"`
}

// EvaluationFeedback defines model for EvaluationFeedback.
type EvaluationFeedback struct {
	CreatedAt      time.Time `json:"createdAt"`
	DatasetRef     string    `json:"datasetRef,omitempty"`
	Decision       string    `json:"decision"`
	ID             string    `json:"id"`
	LicenseRef     string    `json:"licenseRef"`
	RedactedInput  string    `json:"redactedInput"`
	RedactedOutput string    `json:"redactedOutput"`
	ScopeHash      string    `json:"scopeHash"`
	TraceRef       string    `json:"traceRef"`
}

// EvaluationFeedbackEnvelope defines model for EvaluationFeedbackEnvelope.
type EvaluationFeedbackEnvelope struct {
	Data EvaluationFeedback `json:"data"`
}

// EvaluationFeedbackInput defines model for EvaluationFeedbackInput.
type EvaluationFeedbackInput struct {
	Disposition EvaluationFeedbackInputDisposition `json:"disposition"`
	ID          string                             `json:"id"`
	TraceRef    string                             `json:"traceRef"`
}

// EvaluationFeedbackInputDisposition defines model for EvaluationFeedbackInput.Disposition.
type EvaluationFeedbackInputDisposition string

// EvaluationFeedbackListEnvelope defines model for EvaluationFeedbackListEnvelope.
type EvaluationFeedbackListEnvelope struct {
	Items []EvaluationFeedback `json:"items"`
}

// EvaluationGateDecision defines model for EvaluationGateDecision.
type EvaluationGateDecision struct {
	BaselineRunID  string                 `json:"baselineRunId"`
	CandidateRunID string                 `json:"candidateRunId"`
	Decision       string                 `json:"decision"`
	EvaluatedAt    time.Time              `json:"evaluatedAt"`
	EvidenceRefs   []string               `json:"evidenceRefs"`
	ID             string                 `json:"id"`
	PolicyID       string                 `json:"policyId"`
	PolicyVersion  string                 `json:"policyVersion"`
	Reasons        []EvaluationGateReason `json:"reasons"`
}

// EvaluationGateDecisionEnvelope defines model for EvaluationGateDecisionEnvelope.
type EvaluationGateDecisionEnvelope struct {
	Data EvaluationGateDecision `json:"data"`
}

// EvaluationGateDecisionListEnvelope defines model for EvaluationGateDecisionListEnvelope.
type EvaluationGateDecisionListEnvelope struct {
	Items []EvaluationGateDecision `json:"items"`
}

// EvaluationGatePolicy defines model for EvaluationGatePolicy.
type EvaluationGatePolicy struct {
	Enabled           bool               `json:"enabled"`
	ID                string             `json:"id"`
	MaximumCost       float32            `json:"maximumCost,omitempty"`
	MaximumLatencyMs  int64              `json:"maximumLatencyMs,omitempty"`
	MaximumRegression map[string]float32 `json:"maximumRegression"`
	MinimumScores     map[string]float32 `json:"minimumScores"`
	Version           string             `json:"version"`
}

// EvaluationGatePolicyEnvelope defines model for EvaluationGatePolicyEnvelope.
type EvaluationGatePolicyEnvelope struct {
	Data EvaluationGatePolicy `json:"data"`
}

// EvaluationGatePolicyInput defines model for EvaluationGatePolicyInput.
type EvaluationGatePolicyInput struct {
	ID        string  `json:"id"`
	Metric    string  `json:"metric"`
	Name      string  `json:"name,omitempty"`
	Threshold float32 `json:"threshold"`
	Version   string  `json:"version"`
}

// EvaluationGatePolicyListEnvelope defines model for EvaluationGatePolicyListEnvelope.
type EvaluationGatePolicyListEnvelope struct {
	Items []EvaluationGatePolicy `json:"items"`
}

// EvaluationGateReason defines model for EvaluationGateReason.
type EvaluationGateReason struct {
	Actual   float32 `json:"actual"`
	Code     string  `json:"code"`
	Expected float32 `json:"expected"`
	Metric   string  `json:"metric"`
}

// EvaluationGateRequest defines model for EvaluationGateRequest.
type EvaluationGateRequest struct {
	BaselineRunID  string `json:"baselineRunId"`
	CandidateRunID string `json:"candidateRunId"`
	PolicyID       string `json:"policyId"`
}

// EvaluationReplayPlan defines model for EvaluationReplayPlan.
type EvaluationReplayPlan struct {
	CandidateRefs   map[string]string         `json:"candidateRefs"`
	CreatedAt       time.Time                 `json:"createdAt"`
	ID              string                    `json:"id"`
	Profile         EvaluationExecutorProfile `json:"profile"`
	ReadOnly        bool                      `json:"readOnly"`
	SourceTraceRefs []string                  `json:"sourceTraceRefs"`
}

// EvaluationReplayPlanEnvelope defines model for EvaluationReplayPlanEnvelope.
type EvaluationReplayPlanEnvelope struct {
	Data EvaluationReplayPlan `json:"data"`
}

// EvaluationReplayPlanInput defines model for EvaluationReplayPlanInput.
type EvaluationReplayPlanInput struct {
	BaselineRunID     string `json:"baselineRunId"`
	CandidateRunID    string `json:"candidateRunId"`
	ExecutorProfileID string `json:"executorProfileId"`
	ID                string `json:"id"`
}

// EvaluationReplayPlanListEnvelope defines model for EvaluationReplayPlanListEnvelope.
type EvaluationReplayPlanListEnvelope struct {
	Items []EvaluationReplayPlan `json:"items"`
}

// EvaluationResult defines model for EvaluationResult.
type EvaluationResult struct {
	Actions          []string           `json:"actions,omitempty"`
	FailureReasons   []string           `json:"failureReasons,omitempty"`
	Passed           bool               `json:"passed"`
	ProducedFacts    []string           `json:"producedFacts,omitempty"`
	RetrievedSources []string           `json:"retrievedSources,omitempty"`
	SampleID         string             `json:"sampleId"`
	SchemaVersion    string             `json:"schemaVersion"`
	Scores           map[string]float32 `json:"scores"`
}

// EvaluationResultListEnvelope defines model for EvaluationResultListEnvelope.
type EvaluationResultListEnvelope struct {
	Items []EvaluationResult `json:"items"`
}

// EvaluationRun defines model for EvaluationRun.
type EvaluationRun struct {
	AggregateScores map[string]float32  `json:"aggregateScores,omitempty"`
	CandidateRefs   map[string]string   `json:"candidateRefs"`
	CompletedAt     *time.Time          `json:"completedAt,omitempty"`
	DatasetID       string              `json:"datasetId"`
	DatasetVersion  string              `json:"datasetVersion"`
	ID              string              `json:"id"`
	SchemaVersion   string              `json:"schemaVersion"`
	StartedAt       time.Time           `json:"startedAt"`
	Status          EvaluationRunStatus `json:"status"`
}

// EvaluationRunStatus defines model for EvaluationRun.Status.
type EvaluationRunStatus string

// EvaluationRunEnvelope defines model for EvaluationRunEnvelope.
type EvaluationRunEnvelope struct {
	Data EvaluationRun `json:"data"`
}

// EvaluationRunListEnvelope defines model for EvaluationRunListEnvelope.
type EvaluationRunListEnvelope struct {
	Items []EvaluationRun `json:"items"`
}

// EvaluationSampleAttempt defines model for EvaluationSampleAttempt.
type EvaluationSampleAttempt struct {
	Attempt       int                `json:"attempt"`
	CandidateRefs map[string]string  `json:"candidateRefs"`
	CompletedAt   time.Time          `json:"completedAt"`
	ErrorCode     string             `json:"errorCode,omitempty"`
	LatencyMillis int64              `json:"latencyMillis"`
	RunID         string             `json:"runId"`
	SampleID      string             `json:"sampleId"`
	Scores        map[string]float32 `json:"scores,omitempty"`
	Status        string             `json:"status"`
	TraceRef      string             `json:"traceRef,omitempty"`
	Usage         map[string]float32 `json:"usage,omitempty"`
}

// EvaluationSampleAttemptListEnvelope defines model for EvaluationSampleAttemptListEnvelope.
type EvaluationSampleAttemptListEnvelope struct {
	Items []EvaluationSampleAttempt `json:"items"`
}

// EvaluationSampleOutput defines model for EvaluationSampleOutput.
type EvaluationSampleOutput struct {
	Actions          []string `json:"actions,omitempty"`
	ProducedFacts    []string `json:"producedFacts,omitempty"`
	RetrievedSources []string `json:"retrievedSources,omitempty"`
	SampleID         string   `json:"sampleId"`
}

// ExecutionArtifact defines model for ExecutionArtifact.
type ExecutionArtifact struct {
	ApplicationEnvironmentID string         `json:"applicationEnvironmentId,omitempty"`
	ApplicationID            string         `json:"applicationId,omitempty"`
	CreatedAt                *time.Time     `json:"createdAt,omitempty"`
	Digest                   string         `json:"digest,omitempty"`
	ExecutionTaskID          string         `json:"executionTaskId,omitempty"`
	ID                       string         `json:"id,omitempty"`
	Kind                     string         `json:"kind"`
	Metadata                 map[string]any `json:"metadata,omitempty"`
	ModifiedAt               *time.Time     `json:"modifiedAt,omitempty"`
	Name                     string         `json:"name,omitempty"`
	Path                     string         `json:"path,omitempty"`
	Ref                      string         `json:"ref,omitempty"`
	ReleaseBundleID          string         `json:"releaseBundleId,omitempty"`
	RetentionUntil           *time.Time     `json:"retentionUntil,omitempty"`
	SizeBytes                int64          `json:"sizeBytes,omitempty"`
	Status                   string         `json:"status,omitempty"`
	UpdatedAt                *time.Time     `json:"updatedAt,omitempty"`
	WorkflowNodeID           string         `json:"workflowNodeId,omitempty"`
	WorkflowRunID            string         `json:"workflowRunId,omitempty"`
}

// ExecutionArtifactListEnvelope defines model for ExecutionArtifactListEnvelope.
type ExecutionArtifactListEnvelope struct {
	Data []ExecutionArtifact `json:"data"`
}

// ExecutionCallbackRequest defines model for ExecutionCallbackRequest.
type ExecutionCallbackRequest struct {
	CallbackToken string         `json:"callbackToken"`
	Payload       map[string]any `json:"payload"`
	Status        string         `json:"status"`
}

// ExecutionTask defines model for ExecutionTask.
type ExecutionTask struct {
	ApplicationEnvironmentID string              `json:"applicationEnvironmentId"`
	ApplicationID            string              `json:"applicationId"`
	Artifacts                []ExecutionArtifact `json:"artifacts,omitempty"`
	AttemptCount             int                 `json:"attemptCount,omitempty"`
	CallbackToken            string              `json:"callbackToken"`
	ClaimedByAgentID         string              `json:"claimedByAgentId,omitempty"`
	CreatedAt                *time.Time          `json:"createdAt,omitempty"`
	FinishedAt               *time.Time          `json:"finishedAt,omitempty"`
	ID                       string              `json:"id"`
	LastHeartbeatAt          *time.Time          `json:"lastHeartbeatAt,omitempty"`
	LastRuntimeSeenAt        *time.Time          `json:"lastRuntimeSeenAt,omitempty"`
	LockKey                  string              `json:"lockKey,omitempty"`
	MaxRetries               int                 `json:"maxRetries,omitempty"`
	OperationState           *OperationState     `json:"operationState,omitempty"`
	Payload                  map[string]any      `json:"payload"`
	ProviderKind             string              `json:"providerKind"`
	QueueKey                 string              `json:"queueKey,omitempty"`
	ReleaseBundleID          string              `json:"releaseBundleId,omitempty"`
	Result                   map[string]any      `json:"result,omitempty"`
	RuntimeClusterID         string              `json:"runtimeClusterId,omitempty"`
	RuntimeEndpoint          string              `json:"runtimeEndpoint,omitempty"`
	StartedAt                *time.Time          `json:"startedAt,omitempty"`
	Status                   string              `json:"status"`
	StopTransport            string              `json:"stopTransport,omitempty"`
	TargetKind               string              `json:"targetKind,omitempty"`
	TaskKind                 string              `json:"taskKind"`
	TimeoutSeconds           int                 `json:"timeoutSeconds,omitempty"`
	UpdatedAt                *time.Time          `json:"updatedAt,omitempty"`
	AdditionalProperties     map[string]any      `json:"-"`
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

// ExecutionTaskListEnvelope defines model for ExecutionTaskListEnvelope.
type ExecutionTaskListEnvelope struct {
	Data []ExecutionTask `json:"data"`
}

// GeminiGenerateContentRequest Native Gemini generateContent request body. The model is taken from the path and mapped by Soha model routes. Gemini text, inlineData, fileData, cachedContent, generationConfig, and safetySettings fields are passed through; multimodal audio/image/file inputs bypass response cache and local token estimation only counts text parts when provider usageMetadata is absent.
type GeminiGenerateContentRequest struct {
	Contents             []map[string]any `json:"contents,omitempty"`
	GenerationConfig     map[string]any   `json:"generationConfig,omitempty"`
	SafetySettings       []map[string]any `json:"safetySettings,omitempty"`
	SystemInstruction    map[string]any   `json:"systemInstruction,omitempty"`
	AdditionalProperties map[string]any   `json:"-"`
}

// GeminiInteractionsRequest Native Gemini Interactions request body. The public model is supplied in the body and mapped by Soha model routes before forwarding to the upstream `/interactions` endpoint. This first slice supports non-streaming JSON interactions for native image generation; `stream` and `background` modes are rejected by the relay. Multimodal inputs bypass response cache, and provider usage totals are recorded when present.
type GeminiInteractionsRequest struct {
	// Background Background interactions are not supported by this relay slice.
	Background bool `json:"background,omitempty"`

	// Input Native Gemini Interactions input payload.
	Input interface{} `json:"input"`

	// Model Public Gemini model ID exposed by Soha; the relay rewrites it to the configured upstream model.
	Model string `json:"model"`

	// Stream Streaming interactions are not supported by this relay slice.
	Stream               bool           `json:"stream,omitempty"`
	AdditionalProperties map[string]any `json:"-"`
}

// GeminiModel defines model for GeminiModel.
type GeminiModel struct {
	Description                string         `json:"description,omitempty"`
	DisplayName                string         `json:"displayName,omitempty"`
	Name                       string         `json:"name"`
	SupportedGenerationMethods []string       `json:"supportedGenerationMethods,omitempty"`
	Version                    string         `json:"version,omitempty"`
	AdditionalProperties       map[string]any `json:"-"`
}

// GeminiModelsResponse defines model for GeminiModelsResponse.
type GeminiModelsResponse struct {
	Models               []GeminiModel  `json:"models"`
	AdditionalProperties map[string]any `json:"-"`
}

// GenericDataEnvelope defines model for GenericDataEnvelope.
type GenericDataEnvelope struct {
	Data AnyValue `json:"data"`
}

// GenericItemsEnvelope defines model for GenericItemsEnvelope.
type GenericItemsEnvelope struct {
	Items []AnyValue `json:"items"`
}

// GenericObject defines model for GenericObject.
type GenericObject map[string]any

// GitCommit defines model for GitCommit.
type GitCommit struct {
	AuthorEmail openapi_types.Email `json:"authorEmail,omitempty"`
	AuthorName  string              `json:"authorName,omitempty"`
	CommittedAt time.Time           `json:"committedAt"`
	ID          string              `json:"id"`
	Message     string              `json:"message,omitempty"`
	ShortID     string              `json:"shortId"`
	Title       string              `json:"title"`
	WebURL      string              `json:"webUrl,omitempty"`
}

// GitCommitPageEnvelope defines model for GitCommitPageEnvelope.
type GitCommitPageEnvelope struct {
	HasMore bool        `json:"hasMore"`
	Items   []GitCommit `json:"items"`
	Limit   int         `json:"limit"`
	Page    int         `json:"page"`
}

// GitProject defines model for GitProject.
type GitProject struct {
	DefaultBranch     string `json:"defaultBranch,omitempty"`
	ID                string `json:"id"`
	Name              string `json:"name"`
	Path              string `json:"path"`
	PathWithNamespace string `json:"pathWithNamespace"`
	WebURL            string `json:"webUrl,omitempty"`
}

// GitProjectListEnvelope defines model for GitProjectListEnvelope.
type GitProjectListEnvelope struct {
	Items []GitProject `json:"items"`
}

// GitReference defines model for GitReference.
type GitReference struct {
	CommitSha string     `json:"commitSha,omitempty"`
	Name      string     `json:"name"`
	Protected bool       `json:"protected"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}

// GitReferenceListEnvelope defines model for GitReferenceListEnvelope.
type GitReferenceListEnvelope struct {
	Items []GitReference `json:"items"`
}

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

// KnowledgeAccessScope defines model for KnowledgeAccessScope.
type KnowledgeAccessScope struct {
	Projects   []string                       `json:"projects,omitempty"`
	Roles      []string                       `json:"roles,omitempty"`
	Teams      []string                       `json:"teams,omitempty"`
	Users      []string                       `json:"users,omitempty"`
	Visibility KnowledgeAccessScopeVisibility `json:"visibility"`
}

// KnowledgeAccessScopeVisibility defines model for KnowledgeAccessScope.Visibility.
type KnowledgeAccessScopeVisibility string

// KnowledgeBase defines model for KnowledgeBase.
type KnowledgeBase struct {
	CreatedAt       time.Time                `json:"createdAt"`
	Description     string                   `json:"description,omitempty"`
	ID              string                   `json:"id"`
	Name            string                   `json:"name"`
	OwnerID         string                   `json:"ownerId"`
	RetrievalPolicy KnowledgeRetrievalPolicy `json:"retrievalPolicy"`
	Scope           KnowledgeAccessScope     `json:"scope"`
	Status          KnowledgeBaseStatus      `json:"status"`
	TenantID        string                   `json:"tenantId,omitempty"`
	UpdatedAt       time.Time                `json:"updatedAt"`
	WorkspaceID     string                   `json:"workspaceId,omitempty"`
}

// KnowledgeBaseStatus defines model for KnowledgeBase.Status.
type KnowledgeBaseStatus string

// KnowledgeBaseEnvelope defines model for KnowledgeBaseEnvelope.
type KnowledgeBaseEnvelope struct {
	Data KnowledgeBase `json:"data"`
}

// KnowledgeBaseInput defines model for KnowledgeBaseInput.
type KnowledgeBaseInput struct {
	Description     string                   `json:"description,omitempty"`
	Name            string                   `json:"name"`
	RetrievalPolicy KnowledgeRetrievalPolicy `json:"retrievalPolicy"`
	Scope           KnowledgeAccessScope     `json:"scope"`
	TenantID        string                   `json:"tenantId,omitempty"`
	WorkspaceID     string                   `json:"workspaceId,omitempty"`
}

// KnowledgeBaseListEnvelope defines model for KnowledgeBaseListEnvelope.
type KnowledgeBaseListEnvelope struct {
	Items []KnowledgeBase `json:"items"`
}

// KnowledgeCitation defines model for KnowledgeCitation.
type KnowledgeCitation struct {
	ChunkID         string                  `json:"chunkId"`
	ContentHash     string                  `json:"contentHash"`
	DocumentID      string                  `json:"documentId"`
	DocumentTitle   string                  `json:"documentTitle"`
	ID              string                  `json:"id"`
	KnowledgeBaseID string                  `json:"knowledgeBaseId"`
	Location        KnowledgeSourceLocation `json:"location"`
	Score           float32                 `json:"score"`
	URI             string                  `json:"uri,omitempty"`
}

// KnowledgeConnector defines model for KnowledgeConnector.
type KnowledgeConnector struct {
	Config          map[string]any         `json:"config"`
	CreatedAt       time.Time              `json:"createdAt"`
	ID              string                 `json:"id"`
	Kind            KnowledgeConnectorKind `json:"kind"`
	KnowledgeBaseID string                 `json:"knowledgeBaseId"`
	Name            string                 `json:"name"`
	SecretRef       string                 `json:"secretRef"`
	Status          string                 `json:"status"`
	SyncPolicy      KnowledgeSyncPolicy    `json:"syncPolicy"`
	UpdatedAt       time.Time              `json:"updatedAt"`
	Version         string                 `json:"version,omitempty"`
}

// KnowledgeConnectorKind defines model for KnowledgeConnector.Kind.
type KnowledgeConnectorKind string

// KnowledgeConnectorEnvelope defines model for KnowledgeConnectorEnvelope.
type KnowledgeConnectorEnvelope struct {
	Data KnowledgeConnector `json:"data"`
}

// KnowledgeConnectorInput defines model for KnowledgeConnectorInput.
type KnowledgeConnectorInput struct {
	Config          map[string]any              `json:"config"`
	Kind            KnowledgeConnectorInputKind `json:"kind"`
	KnowledgeBaseID string                      `json:"knowledgeBaseId"`
	Name            string                      `json:"name"`
	SecretRef       string                      `json:"secretRef"`
	SyncPolicy      KnowledgeSyncPolicy         `json:"syncPolicy"`
	Version         string                      `json:"version,omitempty"`
}

// KnowledgeConnectorInputKind defines model for KnowledgeConnectorInput.Kind.
type KnowledgeConnectorInputKind string

// KnowledgeConnectorListEnvelope defines model for KnowledgeConnectorListEnvelope.
type KnowledgeConnectorListEnvelope struct {
	Items []KnowledgeConnector `json:"items"`
}

// KnowledgeConnectorValidation defines model for KnowledgeConnectorValidation.
type KnowledgeConnectorValidation struct {
	ConfigHash  string                           `json:"configHash"`
	Host        string                           `json:"host,omitempty"`
	Kind        KnowledgeConnectorValidationKind `json:"kind"`
	Resource    string                           `json:"resource,omitempty"`
	SecretRef   string                           `json:"secretRef"`
	Valid       bool                             `json:"valid"`
	ValidatedAt time.Time                        `json:"validatedAt"`
	Warnings    []string                         `json:"warnings"`
}

// KnowledgeConnectorValidationKind defines model for KnowledgeConnectorValidation.Kind.
type KnowledgeConnectorValidationKind string

// KnowledgeConnectorValidationEnvelope defines model for KnowledgeConnectorValidationEnvelope.
type KnowledgeConnectorValidationEnvelope struct {
	Data KnowledgeConnectorValidation `json:"data"`
}

// KnowledgeDocument defines model for KnowledgeDocument.
type KnowledgeDocument struct {
	ACL             KnowledgeAccessScope    `json:"acl"`
	ChunkCount      int                     `json:"chunkCount"`
	ContentHash     string                  `json:"contentHash"`
	CreatedAt       time.Time               `json:"createdAt"`
	ExternalID      string                  `json:"externalId"`
	ID              string                  `json:"id"`
	KnowledgeBaseID string                  `json:"knowledgeBaseId"`
	SourceID        string                  `json:"sourceId"`
	Status          KnowledgeDocumentStatus `json:"status"`
	Title           string                  `json:"title"`
	UpdatedAt       time.Time               `json:"updatedAt"`
	URI             string                  `json:"uri,omitempty"`
	Version         string                  `json:"version"`
}

// KnowledgeDocumentStatus defines model for KnowledgeDocument.Status.
type KnowledgeDocumentStatus string

// KnowledgeDocumentListEnvelope defines model for KnowledgeDocumentListEnvelope.
type KnowledgeDocumentListEnvelope struct {
	Items []KnowledgeDocument `json:"items"`
}

// KnowledgeGraphCommunity defines model for KnowledgeGraphCommunity.
type KnowledgeGraphCommunity struct {
	EntityIDs  []string `json:"entityIds"`
	ID         string   `json:"id"`
	SourceRefs []string `json:"sourceRefs"`
	Summary    string   `json:"summary"`
}

// KnowledgeGraphEntity defines model for KnowledgeGraphEntity.
type KnowledgeGraphEntity struct {
	ContentHash string   `json:"contentHash"`
	ID          string   `json:"id"`
	Kind        string   `json:"kind"`
	Name        string   `json:"name"`
	SourceRefs  []string `json:"sourceRefs"`
}

// KnowledgeGraphQueryInput defines model for KnowledgeGraphQueryInput.
type KnowledgeGraphQueryInput struct {
	Limit int    `json:"limit"`
	Mode  string `json:"mode"`
	Query string `json:"query"`
}

// KnowledgeGraphQueryResult defines model for KnowledgeGraphQueryResult.
type KnowledgeGraphQueryResult struct {
	Entities   []KnowledgeGraphEntity    `json:"entities"`
	Mode       string                    `json:"mode"`
	NoAnswer   bool                      `json:"noAnswer"`
	Relations  []KnowledgeGraphRelation  `json:"relations"`
	RevisionID string                    `json:"revisionId"`
	Summaries  []KnowledgeGraphCommunity `json:"summaries"`
}

// KnowledgeGraphQueryResultEnvelope defines model for KnowledgeGraphQueryResultEnvelope.
type KnowledgeGraphQueryResultEnvelope struct {
	Data KnowledgeGraphQueryResult `json:"data"`
}

// KnowledgeGraphRelation defines model for KnowledgeGraphRelation.
type KnowledgeGraphRelation struct {
	Confidence   float32  `json:"confidence"`
	FromEntityID string   `json:"fromEntityId"`
	ID           string   `json:"id"`
	Kind         string   `json:"kind"`
	SourceRefs   []string `json:"sourceRefs"`
	ToEntityID   string   `json:"toEntityId"`
}

// KnowledgeGraphRevision defines model for KnowledgeGraphRevision.
type KnowledgeGraphRevision struct {
	Communities      []KnowledgeGraphCommunity `json:"communities"`
	CreatedAt        time.Time                 `json:"createdAt"`
	Entities         []KnowledgeGraphEntity    `json:"entities"`
	ExtractorVersion string                    `json:"extractorVersion"`
	ID               string                    `json:"id"`
	KnowledgeBaseID  string                    `json:"knowledgeBaseId"`
	PublishedAt      *time.Time                `json:"publishedAt,omitempty"`
	Relations        []KnowledgeGraphRelation  `json:"relations"`
	SourceIndexRef   string                    `json:"sourceIndexRef"`
	Status           string                    `json:"status"`
}

// KnowledgeGraphRevisionEnvelope defines model for KnowledgeGraphRevisionEnvelope.
type KnowledgeGraphRevisionEnvelope struct {
	Data KnowledgeGraphRevision `json:"data"`
}

// KnowledgeGraphRevisionListEnvelope defines model for KnowledgeGraphRevisionListEnvelope.
type KnowledgeGraphRevisionListEnvelope struct {
	Items []KnowledgeGraphRevision `json:"items"`
}

// KnowledgeIndexRevision defines model for KnowledgeIndexRevision.
type KnowledgeIndexRevision struct {
	ActivatedAt     *time.Time                   `json:"activatedAt,omitempty"`
	ChunkCount      int                          `json:"chunkCount"`
	ChunkerVersion  string                       `json:"chunkerVersion"`
	CreatedAt       time.Time                    `json:"createdAt"`
	DocumentCount   int                          `json:"documentCount"`
	EmbeddingModel  string                       `json:"embeddingModel,omitempty"`
	ID              string                       `json:"id"`
	KnowledgeBaseID string                       `json:"knowledgeBaseId"`
	Revision        int                          `json:"revision"`
	Status          KnowledgeIndexRevisionStatus `json:"status"`
}

// KnowledgeIndexRevisionStatus defines model for KnowledgeIndexRevision.Status.
type KnowledgeIndexRevisionStatus string

// KnowledgeIndexRevisionListEnvelope defines model for KnowledgeIndexRevisionListEnvelope.
type KnowledgeIndexRevisionListEnvelope struct {
	Items []KnowledgeIndexRevision `json:"items"`
}

// KnowledgeIngestionCheckpoint defines model for KnowledgeIngestionCheckpoint.
type KnowledgeIngestionCheckpoint struct {
	ChunksStored    int                     `json:"chunksStored"`
	ContentHash     string                  `json:"contentHash,omitempty"`
	Cursor          string                  `json:"cursor,omitempty"`
	DocumentsSeen   int                     `json:"documentsSeen"`
	DocumentsStored int                     `json:"documentsStored"`
	RecordedAt      time.Time               `json:"recordedAt"`
	Stage           KnowledgeIngestionStage `json:"stage"`
}

// KnowledgeIngestionJob defines model for KnowledgeIngestionJob.
type KnowledgeIngestionJob struct {
	Attempt         int                          `json:"attempt"`
	CancelRequested bool                         `json:"cancelRequested"`
	Checkpoint      KnowledgeIngestionCheckpoint `json:"checkpoint"`
	CompletedAt     *time.Time                   `json:"completedAt,omitempty"`
	CreatedAt       time.Time                    `json:"createdAt"`
	Error           string                       `json:"error,omitempty"`
	ErrorCode       string                       `json:"errorCode,omitempty"`
	ID              string                       `json:"id"`
	KnowledgeBaseID string                       `json:"knowledgeBaseId"`
	MaxAttempts     int                          `json:"maxAttempts"`
	NextAttemptAt   *time.Time                   `json:"nextAttemptAt,omitempty"`
	SourceID        string                       `json:"sourceId"`
	Stage           KnowledgeIngestionStage      `json:"stage"`
	Status          KnowledgeIngestionJobStatus  `json:"status"`
	TargetRevision  int64                        `json:"targetRevision"`
	UpdatedAt       time.Time                    `json:"updatedAt"`
}

// KnowledgeIngestionJobStatus defines model for KnowledgeIngestionJob.Status.
type KnowledgeIngestionJobStatus string

// KnowledgeIngestionJobEnvelope defines model for KnowledgeIngestionJobEnvelope.
type KnowledgeIngestionJobEnvelope struct {
	Data KnowledgeIngestionJob `json:"data"`
}

// KnowledgeIngestionJobInput defines model for KnowledgeIngestionJobInput.
type KnowledgeIngestionJobInput struct {
	SourceID string `json:"sourceId"`
}

// KnowledgeIngestionStage defines model for KnowledgeIngestionStage.
type KnowledgeIngestionStage string

// KnowledgeRetrievalPolicy defines model for KnowledgeRetrievalPolicy.
type KnowledgeRetrievalPolicy struct {
	DefaultTopK   int     `json:"defaultTopK"`
	LexicalWeight float32 `json:"lexicalWeight"`
	MaxTopK       int     `json:"maxTopK"`
	MinScore      float32 `json:"minScore"`
	VectorWeight  float32 `json:"vectorWeight"`
}

// KnowledgeSearchFilters defines model for KnowledgeSearchFilters.
type KnowledgeSearchFilters struct {
	DocumentIDs []string `json:"documentIds,omitempty"`
	SourceIDs   []string `json:"sourceIds,omitempty"`
}

// KnowledgeSearchHit defines model for KnowledgeSearchHit.
type KnowledgeSearchHit struct {
	ChunkID         string            `json:"chunkId"`
	Citation        KnowledgeCitation `json:"citation"`
	Content         string            `json:"content"`
	DocumentID      string            `json:"documentId"`
	KnowledgeBaseID string            `json:"knowledgeBaseId"`
	LexicalScore    float32           `json:"lexicalScore"`
	Score           float32           `json:"score"`
	Title           string            `json:"title"`
	VectorScore     float32           `json:"vectorScore"`
}

// KnowledgeSearchRequest defines model for KnowledgeSearchRequest.
type KnowledgeSearchRequest struct {
	Filters          *KnowledgeSearchFilters `json:"filters,omitempty"`
	KnowledgeBaseIDs []string                `json:"knowledgeBaseIds"`
	Query            string                  `json:"query"`
	TopK             int                     `json:"topK,omitempty"`
}

// KnowledgeSearchResult defines model for KnowledgeSearchResult.
type KnowledgeSearchResult struct {
	CandidateCount int                  `json:"candidateCount"`
	Citations      []KnowledgeCitation  `json:"citations"`
	Hits           []KnowledgeSearchHit `json:"hits"`
	NoAnswer       bool                 `json:"noAnswer"`
	Query          string               `json:"query"`
	TimingMs       int64                `json:"timingMs"`
	TraceID        string               `json:"traceId"`
}

// KnowledgeSearchResultEnvelope defines model for KnowledgeSearchResultEnvelope.
type KnowledgeSearchResultEnvelope struct {
	Data KnowledgeSearchResult `json:"data"`
}

// KnowledgeSource defines model for KnowledgeSource.
type KnowledgeSource struct {
	ConfigRef       string                `json:"configRef,omitempty"`
	CreatedAt       time.Time             `json:"createdAt"`
	Cursor          string                `json:"cursor,omitempty"`
	ID              string                `json:"id"`
	Kind            KnowledgeSourceKind   `json:"kind"`
	KnowledgeBaseID string                `json:"knowledgeBaseId"`
	LastError       string                `json:"lastError,omitempty"`
	LastSyncedAt    *time.Time            `json:"lastSyncedAt,omitempty"`
	Name            string                `json:"name"`
	Status          KnowledgeSourceStatus `json:"status"`
	SyncPolicy      KnowledgeSyncPolicy   `json:"syncPolicy"`
	UpdatedAt       time.Time             `json:"updatedAt"`
}

// KnowledgeSourceKind defines model for KnowledgeSource.Kind.
type KnowledgeSourceKind string

// KnowledgeSourceStatus defines model for KnowledgeSource.Status.
type KnowledgeSourceStatus string

// KnowledgeSourceConfig defines model for KnowledgeSourceConfig.
type KnowledgeSourceConfig struct {
	Documents []KnowledgeSourceDocument `json:"documents,omitempty"`
}

// KnowledgeSourceDocument defines model for KnowledgeSourceDocument.
type KnowledgeSourceDocument struct {
	ACL        KnowledgeAccessScope `json:"acl"`
	Content    string               `json:"content"`
	ExternalID string               `json:"externalId"`
	Title      string               `json:"title"`
	URI        string               `json:"uri,omitempty"`
	Version    string               `json:"version,omitempty"`
}

// KnowledgeSourceEnvelope defines model for KnowledgeSourceEnvelope.
type KnowledgeSourceEnvelope struct {
	Data KnowledgeSource `json:"data"`
}

// KnowledgeSourceInput defines model for KnowledgeSourceInput.
type KnowledgeSourceInput struct {
	Config     *KnowledgeSourceConfig   `json:"config,omitempty"`
	ConfigRef  string                   `json:"configRef,omitempty"`
	Kind       KnowledgeSourceInputKind `json:"kind"`
	Name       string                   `json:"name"`
	SyncPolicy KnowledgeSyncPolicy      `json:"syncPolicy"`
}

// KnowledgeSourceInputKind defines model for KnowledgeSourceInput.Kind.
type KnowledgeSourceInputKind string

// KnowledgeSourceListEnvelope defines model for KnowledgeSourceListEnvelope.
type KnowledgeSourceListEnvelope struct {
	Items []KnowledgeSource `json:"items"`
}

// KnowledgeSourceLocation defines model for KnowledgeSourceLocation.
type KnowledgeSourceLocation struct {
	EndByte   int    `json:"endByte,omitempty"`
	StartByte int    `json:"startByte,omitempty"`
	URI       string `json:"uri,omitempty"`
}

// KnowledgeSyncPolicy defines model for KnowledgeSyncPolicy.
type KnowledgeSyncPolicy struct {
	Mode     string `json:"mode"`
	Schedule string `json:"schedule,omitempty"`
}

// KnowledgeSyncRun defines model for KnowledgeSyncRun.
type KnowledgeSyncRun struct {
	ChunksStored    int                    `json:"chunksStored"`
	CompletedAt     *time.Time             `json:"completedAt,omitempty"`
	DocumentsSeen   int                    `json:"documentsSeen"`
	DocumentsStored int                    `json:"documentsStored"`
	Error           string                 `json:"error,omitempty"`
	ID              string                 `json:"id"`
	KnowledgeBaseID string                 `json:"knowledgeBaseId"`
	SourceID        string                 `json:"sourceId"`
	StartedAt       time.Time              `json:"startedAt"`
	Status          KnowledgeSyncRunStatus `json:"status"`
}

// KnowledgeSyncRunStatus defines model for KnowledgeSyncRun.Status.
type KnowledgeSyncRunStatus string

// KnowledgeSyncRunEnvelope defines model for KnowledgeSyncRunEnvelope.
type KnowledgeSyncRunEnvelope struct {
	Data KnowledgeSyncRun `json:"data"`
}

// KnowledgeSyncRunListEnvelope defines model for KnowledgeSyncRunListEnvelope.
type KnowledgeSyncRunListEnvelope struct {
	Items []KnowledgeSyncRun `json:"items"`
}

// KubernetesResourceAction defines model for KubernetesResourceAction.
type KubernetesResourceAction string

// KubernetesResourceAgentCreateDocument defines model for KubernetesResourceAgentCreateDocument.
type KubernetesResourceAgentCreateDocument struct {
	// Content Exactly one YAML or JSON document. The Agent must verify its hash, GVK, name, and namespace against document and resourceRef before using discovery or the dynamic client.
	Content     string                     `json:"content"`
	Document    KubernetesResourceDocument `json:"document"`
	ResourceRef KubernetesResourceRef      `json:"resourceRef"`
}

// KubernetesResourceAgentCreateRequest defines model for KubernetesResourceAgentCreateRequest.
type KubernetesResourceAgentCreateRequest struct {
	Documents   []KubernetesResourceAgentCreateDocument `json:"documents"`
	OperationID string                                  `json:"operationId"`
}

// KubernetesResourceAgentCreateResult defines model for KubernetesResourceAgentCreateResult.
type KubernetesResourceAgentCreateResult struct {
	Items       []KubernetesResourceCreateResultItem `json:"items"`
	OperationID string                               `json:"operationId"`
	Status      KubernetesResourceCreateBatchStatus  `json:"status"`
}

// KubernetesResourceAgentPreflightItem defines model for KubernetesResourceAgentPreflightItem.
type KubernetesResourceAgentPreflightItem struct {
	Document    KubernetesResourceDocument       `json:"document"`
	DryRun      KubernetesResourceDryRunDecision `json:"dryRun"`
	Errors      []KubernetesResourceCreateError  `json:"errors"`
	ResourceRef KubernetesResourceRef            `json:"resourceRef"`
	Warnings    []KubernetesResourceWarning      `json:"warnings"`
}

// KubernetesResourceAgentPreflightResult defines model for KubernetesResourceAgentPreflightResult.
type KubernetesResourceAgentPreflightResult struct {
	Items []KubernetesResourceAgentPreflightItem `json:"items"`
	Ready bool                                   `json:"ready"`
}

// KubernetesResourceAuthorizationDecision defines model for KubernetesResourceAuthorizationDecision.
type KubernetesResourceAuthorizationDecision struct {
	Allowed        bool                           `json:"allowed"`
	AllowedActions []KubernetesResourceAction     `json:"allowedActions"`
	Error          *KubernetesResourceCreateError `json:"error,omitempty"`

	// Reason Redacted explanation when create is denied.
	Reason        string                  `json:"reason,omitempty"`
	ResourceScope KubernetesResourceScope `json:"resourceScope"`
}

// KubernetesResourceCapability defines model for KubernetesResourceCapability.
type KubernetesResourceCapability struct {
	Key  string                           `json:"key"`
	Mode KubernetesResourceCapabilityMode `json:"mode"`

	// Reason Redacted explanation when the capability is partial or unsupported.
	Reason string                  `json:"reason,omitempty"`
	Status ClusterCapabilityStatus `json:"status"`
}

// KubernetesResourceCapabilityMode defines model for KubernetesResourceCapability.Mode.
type KubernetesResourceCapabilityMode string

// KubernetesResourceCreateBatchStatus defines model for KubernetesResourceCreateBatchStatus.
type KubernetesResourceCreateBatchStatus string

// KubernetesResourceCreateError defines model for KubernetesResourceCreateError.
type KubernetesResourceCreateError struct {
	Code KubernetesResourceCreateErrorCode `json:"code"`

	// Field Optional manifest or request field associated with the diagnostic.
	Field string `json:"field,omitempty"`

	// Message Redacted user-facing diagnostic. Manifest contents and Kubernetes credentials are excluded.
	Message string `json:"message"`
}

// KubernetesResourceCreateErrorCode defines model for KubernetesResourceCreateErrorCode.
type KubernetesResourceCreateErrorCode string

// KubernetesResourceCreateRequest defines model for KubernetesResourceCreateRequest.
type KubernetesResourceCreateRequest struct {
	// Content YAML or JSON manifest content. The server performs bounded multi-document decoding and assigns stable zero-based document indexes.
	Content string `json:"content"`

	// DefaultNamespace Fallback namespace for namespace-scoped global YAML documents that omit metadata.namespace; list and form sources use it as their enforced namespace context.
	DefaultNamespace string `json:"defaultNamespace,omitempty"`

	// ExpectedAPIVersion Optional list or form context API version. The server still resolves the manifest GVK.
	ExpectedAPIVersion string `json:"expectedApiVersion,omitempty"`

	// ExpectedKind Required by list and form callers to enforce the single-kind context boundary.
	ExpectedKind string `json:"expectedKind,omitempty"`

	// ResourceGroup Product resource family used for authorization and list-context validation.
	ResourceGroup string                         `json:"resourceGroup,omitempty"`
	Source        KubernetesResourceCreateSource `json:"source"`
}

// KubernetesResourceCreateResult defines model for KubernetesResourceCreateResult.
type KubernetesResourceCreateResult struct {
	ContentHash string                               `json:"contentHash"`
	Items       []KubernetesResourceCreateResultItem `json:"items"`
	OperationID string                               `json:"operationId"`
	Status      KubernetesResourceCreateBatchStatus  `json:"status"`
}

// KubernetesResourceCreateResultEnvelope defines model for KubernetesResourceCreateResultEnvelope.
type KubernetesResourceCreateResultEnvelope struct {
	Data KubernetesResourceCreateResult `json:"data"`
}

// KubernetesResourceCreateResultItem defines model for KubernetesResourceCreateResultItem.
type KubernetesResourceCreateResultItem struct {
	Document KubernetesResourceDocument     `json:"document"`
	Error    *KubernetesResourceCreateError `json:"error,omitempty"`

	// OperationID Optional child operation associated with this resource result.
	OperationID string                               `json:"operationId,omitempty"`
	ResourceRef *KubernetesResourceRef               `json:"resourceRef,omitempty"`
	Status      KubernetesResourceCreateResultStatus `json:"status"`
	Warnings    []KubernetesResourceWarning          `json:"warnings"`
}

// KubernetesResourceCreateResultStatus defines model for KubernetesResourceCreateResultStatus.
type KubernetesResourceCreateResultStatus string

// KubernetesResourceCreateScopeDecision defines model for KubernetesResourceCreateScopeDecision.
type KubernetesResourceCreateScopeDecision struct {
	Allowed        bool                         `json:"allowed"`
	AllowedActions []KubernetesResourceAction   `json:"allowedActions"`
	Capability     KubernetesResourceCapability `json:"capability"`
	Reason         string                       `json:"reason,omitempty"`
	ResourceScope  KubernetesResourceScope      `json:"resourceScope"`
}

// KubernetesResourceCreateScopeDecisionEnvelope defines model for KubernetesResourceCreateScopeDecisionEnvelope.
type KubernetesResourceCreateScopeDecisionEnvelope struct {
	Data KubernetesResourceCreateScopeDecision `json:"data"`
}

// KubernetesResourceCreateScopeDecisionRequest defines model for KubernetesResourceCreateScopeDecisionRequest.
type KubernetesResourceCreateScopeDecisionRequest struct {
	Action        KubernetesResourceCreateScopeDecisionRequestAction `json:"action"`
	APIVersion    string                                             `json:"apiVersion,omitempty"`
	Kind          string                                             `json:"kind"`
	Namespace     string                                             `json:"namespace,omitempty"`
	ResourceGroup string                                             `json:"resourceGroup"`
}

// KubernetesResourceCreateScopeDecisionRequestAction defines model for KubernetesResourceCreateScopeDecisionRequest.Action.
type KubernetesResourceCreateScopeDecisionRequestAction string

// KubernetesResourceCreateSource defines model for KubernetesResourceCreateSource.
type KubernetesResourceCreateSource string

// KubernetesResourceDocument defines model for KubernetesResourceDocument.
type KubernetesResourceDocument struct {
	APIVersion string `json:"apiVersion,omitempty"`

	// ContentHash SHA-256 of the normalized individual document; the document body is never returned.
	ContentHash string `json:"contentHash"`

	// Index Stable zero-based index after empty YAML documents are removed.
	Index int    `json:"index"`
	Kind  string `json:"kind,omitempty"`
	Name  string `json:"name,omitempty"`

	// Namespace Explicit metadata.namespace from the submitted manifest, before scope normalization.
	Namespace string                      `json:"namespace,omitempty"`
	ScopeMode KubernetesResourceScopeMode `json:"scopeMode,omitempty"`
}

// KubernetesResourceDryRunDecision defines model for KubernetesResourceDryRunDecision.
type KubernetesResourceDryRunDecision struct {
	Error  *KubernetesResourceCreateError `json:"error,omitempty"`
	Status KubernetesResourceDryRunStatus `json:"status"`
}

// KubernetesResourceDryRunStatus defines model for KubernetesResourceDryRunStatus.
type KubernetesResourceDryRunStatus string

// KubernetesResourcePreflight defines model for KubernetesResourcePreflight.
type KubernetesResourcePreflight struct {
	// ContentHash SHA-256 of the normalized complete request content, for matching UI state only; it is not an authorization token.
	ContentHash string                            `json:"contentHash"`
	Items       []KubernetesResourcePreflightItem `json:"items"`
	Ready       bool                              `json:"ready"`
}

// KubernetesResourcePreflightEnvelope defines model for KubernetesResourcePreflightEnvelope.
type KubernetesResourcePreflightEnvelope struct {
	Data KubernetesResourcePreflight `json:"data"`
}

// KubernetesResourcePreflightItem defines model for KubernetesResourcePreflightItem.
type KubernetesResourcePreflightItem struct {
	Authorization KubernetesResourceAuthorizationDecision `json:"authorization"`
	Capability    KubernetesResourceCapability            `json:"capability"`
	Document      KubernetesResourceDocument              `json:"document"`
	DryRun        KubernetesResourceDryRunDecision        `json:"dryRun"`
	Errors        []KubernetesResourceCreateError         `json:"errors"`

	// ResolvedNamespace Final namespace used for authorization and creation; omitted for cluster-scoped resources.
	ResolvedNamespace string                      `json:"resolvedNamespace,omitempty"`
	Warnings          []KubernetesResourceWarning `json:"warnings"`
}

// KubernetesResourceRef defines model for KubernetesResourceRef.
type KubernetesResourceRef struct {
	APIVersion string                      `json:"apiVersion"`
	ClusterID  string                      `json:"clusterId"`
	Kind       string                      `json:"kind"`
	Name       string                      `json:"name"`
	Namespace  string                      `json:"namespace,omitempty"`
	ScopeMode  KubernetesResourceScopeMode `json:"scopeMode"`
	UID        string                      `json:"uid,omitempty"`
}

// KubernetesResourceScope defines model for KubernetesResourceScope.
type KubernetesResourceScope struct {
	ClusterIDs []string `json:"clusterIds"`

	// Namespaces Authorized namespace names. Empty for cluster-scoped decisions.
	Namespaces     []string `json:"namespaces"`
	ResourceGroups []string `json:"resourceGroups"`

	// ResourceKinds Kind restrictions applied after resource-group authorization; empty means no additional kind restriction.
	ResourceKinds []string `json:"resourceKinds"`
}

// KubernetesResourceScopeMode defines model for KubernetesResourceScopeMode.
type KubernetesResourceScopeMode string

// KubernetesResourceWarning defines model for KubernetesResourceWarning.
type KubernetesResourceWarning struct {
	Code  KubernetesResourceCreateErrorCode `json:"code"`
	Field string                            `json:"field,omitempty"`

	// Message Redacted user-facing warning describing a non-fatal normalization.
	Message string `json:"message"`
}

// LLMCallLog defines model for LLMCallLog.
type LLMCallLog struct {
	ActorID           string                `json:"actorId,omitempty"`
	ActorName         string                `json:"actorName,omitempty"`
	ActorType         string                `json:"actorType,omitempty"`
	AIClientID        string                `json:"aiClientId,omitempty"`
	CacheStatus       LLMCallLogCacheStatus `json:"cacheStatus,omitempty"`
	CachedReadTokens  int                   `json:"cachedReadTokens,omitempty"`
	CachedWriteTokens int                   `json:"cachedWriteTokens,omitempty"`
	CompletionTokens  int                   `json:"completionTokens,omitempty"`
	CreatedAt         time.Time             `json:"createdAt"`
	DurationMs        int64                 `json:"durationMs,omitempty"`
	Endpoint          string                `json:"endpoint,omitempty"`
	ErrorCode         string                `json:"errorCode,omitempty"`

	// ErrorMessage Redacted error summary.
	ErrorMessage    string `json:"errorMessage,omitempty"`
	EstimatedTokens bool   `json:"estimatedTokens"`
	HTTPStatus      int    `json:"httpStatus,omitempty"`
	ID              string `json:"id"`
	InputBytes      int64  `json:"inputBytes,omitempty"`

	// Metadata Redacted metadata only; prompt bodies, raw provider payloads, full headers, and credentials are not included.
	Metadata        map[string]any         `json:"metadata,omitempty"`
	OutputBytes     int64                  `json:"outputBytes,omitempty"`
	PromptTokens    int                    `json:"promptTokens,omitempty"`
	ProviderKind    LLMCallLogProviderKind `json:"providerKind,omitempty"`
	PublicModel     string                 `json:"publicModel,omitempty"`
	ReasoningTokens int                    `json:"reasoningTokens,omitempty"`
	RequestID       string                 `json:"requestId,omitempty"`
	RouteTrace      map[string]any         `json:"routeTrace,omitempty"`
	SourceIP        string                 `json:"sourceIp,omitempty"`
	Status          LLMCallLogStatus       `json:"status"`
	Stream          bool                   `json:"stream"`
	TokenID         string                 `json:"tokenId,omitempty"`
	TokenKind       LLMCallLogTokenKind    `json:"tokenKind,omitempty"`
	TokenPrefix     string                 `json:"tokenPrefix,omitempty"`
	TotalTokens     int                    `json:"totalTokens,omitempty"`
	TtfbMs          int64                  `json:"ttfbMs,omitempty"`
	TtftMs          int64                  `json:"ttftMs,omitempty"`
	UpstreamID      string                 `json:"upstreamId,omitempty"`
	UpstreamModel   string                 `json:"upstreamModel,omitempty"`
	UpstreamName    string                 `json:"upstreamName,omitempty"`
	UpstreamStatus  int                    `json:"upstreamStatus,omitempty"`
	UserAgent       string                 `json:"userAgent,omitempty"`
}

// LLMCallLogCacheStatus defines model for LLMCallLog.CacheStatus.
type LLMCallLogCacheStatus string

// LLMCallLogProviderKind defines model for LLMCallLog.ProviderKind.
type LLMCallLogProviderKind string

// LLMCallLogStatus defines model for LLMCallLog.Status.
type LLMCallLogStatus string

// LLMCallLogTokenKind defines model for LLMCallLog.TokenKind.
type LLMCallLogTokenKind string

// LLMCallLogListEnvelope defines model for LLMCallLogListEnvelope.
type LLMCallLogListEnvelope struct {
	Items []LLMCallLog `json:"items"`
}

// LLMModelRoute defines model for LLMModelRoute.
type LLMModelRoute struct {
	CachePolicy        map[string]any            `json:"cachePolicy,omitempty"`
	CreatedAt          time.Time                 `json:"createdAt"`
	Enabled            bool                      `json:"enabled"`
	FallbackPolicy     map[string]any            `json:"fallbackPolicy,omitempty"`
	ID                 string                    `json:"id"`
	Metadata           map[string]any            `json:"metadata,omitempty"`
	Priority           int                       `json:"priority"`
	ProviderKind       LLMModelRouteProviderKind `json:"providerKind,omitempty"`
	PublicModel        string                    `json:"publicModel"`
	RateLimitProfileID string                    `json:"rateLimitProfileId,omitempty"`
	RouteGroup         string                    `json:"routeGroup,omitempty"`

	// TransformPolicy Optional relay format conversion policy. Set `mode: convert` and `targetProviderKind` to `openai` or `anthropic` to enable text-only non-streaming OpenAI Chat Completions <-> Anthropic Messages conversion for this route. Streaming, tool/function, multimodal, audio, image, and file payloads are not converted.
	TransformPolicy map[string]any `json:"transformPolicy,omitempty"`
	UpdatedAt       time.Time      `json:"updatedAt"`
	UpstreamID      string         `json:"upstreamId,omitempty"`
	UpstreamModel   string         `json:"upstreamModel"`
	Weight          int            `json:"weight"`
}

// LLMModelRouteProviderKind defines model for LLMModelRoute.ProviderKind.
type LLMModelRouteProviderKind string

// LLMModelRouteEnvelope defines model for LLMModelRouteEnvelope.
type LLMModelRouteEnvelope struct {
	Data LLMModelRoute `json:"data"`
}

// LLMModelRouteInput defines model for LLMModelRouteInput.
type LLMModelRouteInput struct {
	CachePolicy        map[string]any                 `json:"cachePolicy,omitempty"`
	Enabled            bool                           `json:"enabled,omitempty"`
	FallbackPolicy     map[string]any                 `json:"fallbackPolicy,omitempty"`
	ID                 string                         `json:"id,omitempty"`
	Metadata           map[string]any                 `json:"metadata,omitempty"`
	Priority           int                            `json:"priority,omitempty"`
	ProviderKind       LLMModelRouteInputProviderKind `json:"providerKind,omitempty"`
	PublicModel        string                         `json:"publicModel"`
	RateLimitProfileID string                         `json:"rateLimitProfileId,omitempty"`
	RouteGroup         string                         `json:"routeGroup,omitempty"`

	// TransformPolicy Optional relay format conversion policy. Set `mode: convert` and `targetProviderKind` to `openai` or `anthropic` to enable text-only non-streaming OpenAI Chat Completions <-> Anthropic Messages conversion for this route. Streaming, tool/function, multimodal, audio, image, and file payloads are not converted.
	TransformPolicy map[string]any `json:"transformPolicy,omitempty"`
	UpstreamID      string         `json:"upstreamId,omitempty"`
	UpstreamModel   string         `json:"upstreamModel"`
	Weight          int            `json:"weight,omitempty"`
}

// LLMModelRouteInputProviderKind defines model for LLMModelRouteInput.ProviderKind.
type LLMModelRouteInputProviderKind string

// LLMModelRouteListEnvelope defines model for LLMModelRouteListEnvelope.
type LLMModelRouteListEnvelope struct {
	Items []LLMModelRoute `json:"items"`
}

// LLMRelayCachePurgeRequest defines model for LLMRelayCachePurgeRequest.
type LLMRelayCachePurgeRequest struct {
	DryRun      bool       `json:"dryRun,omitempty"`
	OlderThan   *time.Time `json:"olderThan,omitempty"`
	PublicModel string     `json:"publicModel,omitempty"`
	RouteGroup  string     `json:"routeGroup,omitempty"`
	UpstreamID  string     `json:"upstreamId,omitempty"`
}

// LLMRelayCachePurgeResult defines model for LLMRelayCachePurgeResult.
type LLMRelayCachePurgeResult struct {
	DryRun      bool   `json:"dryRun,omitempty"`
	PurgedCount int    `json:"purgedCount"`
	Status      string `json:"status"`
}

// LLMRelayCachePurgeResultEnvelope defines model for LLMRelayCachePurgeResultEnvelope.
type LLMRelayCachePurgeResultEnvelope struct {
	Data LLMRelayCachePurgeResult `json:"data"`
}

// LLMRelayCacheStats defines model for LLMRelayCacheStats.
type LLMRelayCacheStats struct {
	ByModel                   []map[string]any `json:"byModel,omitempty"`
	ByUpstream                []map[string]any `json:"byUpstream,omitempty"`
	GeneratedAt               time.Time        `json:"generatedAt"`
	ProviderCachedReadTokens  int              `json:"providerCachedReadTokens,omitempty"`
	ProviderCachedWriteTokens int              `json:"providerCachedWriteTokens,omitempty"`
	ResponseCacheBypasses     int              `json:"responseCacheBypasses,omitempty"`
	ResponseCacheEnabled      bool             `json:"responseCacheEnabled"`
	ResponseCacheHits         int              `json:"responseCacheHits,omitempty"`
	ResponseCacheMisses       int              `json:"responseCacheMisses,omitempty"`
	ResponseCacheWrites       int              `json:"responseCacheWrites,omitempty"`
	WindowHours               int              `json:"windowHours"`
}

// LLMRelayCacheStatsEnvelope defines model for LLMRelayCacheStatsEnvelope.
type LLMRelayCacheStatsEnvelope struct {
	Data LLMRelayCacheStats `json:"data"`
}

// LLMRelayHealthCheckRun defines model for LLMRelayHealthCheckRun.
type LLMRelayHealthCheckRun struct {
	Checked   int                     `json:"checked"`
	CheckedAt time.Time               `json:"checkedAt"`
	Degraded  int                     `json:"degraded"`
	Failed    int                     `json:"failed"`
	Healthy   int                     `json:"healthy"`
	Recovered int                     `json:"recovered"`
	Results   []LLMUpstreamTestResult `json:"results,omitempty"`
	Skipped   int                     `json:"skipped"`
	Total     int                     `json:"total"`
}

// LLMRelayHealthCheckRunEnvelope defines model for LLMRelayHealthCheckRunEnvelope.
type LLMRelayHealthCheckRunEnvelope struct {
	Data LLMRelayHealthCheckRun `json:"data"`
}

// LLMRelayMetricBreakdown defines model for LLMRelayMetricBreakdown.
type LLMRelayMetricBreakdown struct {
	AverageDurationMs float32 `json:"averageDurationMs,omitempty"`
	AverageTTFBMs     float32 `json:"averageTTFBMs,omitempty"`
	AverageTTFTMs     float32 `json:"averageTTFTMs,omitempty"`
	CachedReadTokens  int     `json:"cachedReadTokens,omitempty"`
	CachedWriteTokens int     `json:"cachedWriteTokens,omitempty"`
	FailureCount      int     `json:"failureCount,omitempty"`
	Key               string  `json:"key"`
	RequestCount      int     `json:"requestCount"`
	StreamCount       int     `json:"streamCount,omitempty"`
	SuccessCount      int     `json:"successCount,omitempty"`
	TotalTokens       int     `json:"totalTokens,omitempty"`
}

// LLMRelayMetricSeriesPoint defines model for LLMRelayMetricSeriesPoint.
type LLMRelayMetricSeriesPoint struct {
	AverageDurationMs float32   `json:"averageDurationMs,omitempty"`
	AverageTTFBMs     float32   `json:"averageTTFBMs,omitempty"`
	AverageTTFTMs     float32   `json:"averageTTFTMs,omitempty"`
	CachedReadTokens  int       `json:"cachedReadTokens,omitempty"`
	CachedWriteTokens int       `json:"cachedWriteTokens,omitempty"`
	CompletionTokens  int       `json:"completionTokens,omitempty"`
	FailureCount      int       `json:"failureCount,omitempty"`
	PolicyDeniedCount int       `json:"policyDeniedCount,omitempty"`
	PromptTokens      int       `json:"promptTokens,omitempty"`
	RateLimitedCount  int       `json:"rateLimitedCount,omitempty"`
	RequestCount      int       `json:"requestCount"`
	SuccessCount      int       `json:"successCount,omitempty"`
	Timestamp         time.Time `json:"timestamp"`
	TotalTokens       int       `json:"totalTokens,omitempty"`
}

// LLMRelayMetrics defines model for LLMRelayMetrics.
type LLMRelayMetrics struct {
	AverageDurationMs    float32                     `json:"averageDurationMs,omitempty"`
	AverageTTFBMs        float32                     `json:"averageTTFBMs,omitempty"`
	AverageTTFTMs        float32                     `json:"averageTTFTMs,omitempty"`
	ByModel              []LLMRelayMetricBreakdown   `json:"byModel,omitempty"`
	ByUpstream           []LLMRelayMetricBreakdown   `json:"byUpstream,omitempty"`
	CachedReadTokens     int                         `json:"cachedReadTokens,omitempty"`
	CachedWriteTokens    int                         `json:"cachedWriteTokens,omitempty"`
	ClientCancelledCount int                         `json:"clientCancelledCount,omitempty"`
	CompletionTokens     int                         `json:"completionTokens,omitempty"`
	FailureCount         int                         `json:"failureCount"`
	GeneratedAt          time.Time                   `json:"generatedAt"`
	PolicyDeniedCount    int                         `json:"policyDeniedCount,omitempty"`
	PromptTokens         int                         `json:"promptTokens,omitempty"`
	RateLimitedCount     int                         `json:"rateLimitedCount,omitempty"`
	ReasoningTokens      int                         `json:"reasoningTokens,omitempty"`
	RequestCount         int                         `json:"requestCount"`
	Series               []LLMRelayMetricSeriesPoint `json:"series,omitempty"`
	StreamCount          int                         `json:"streamCount,omitempty"`
	SuccessCount         int                         `json:"successCount"`
	TotalTokens          int                         `json:"totalTokens,omitempty"`
	WindowHours          int                         `json:"windowHours"`
}

// LLMRelayMetricsEnvelope defines model for LLMRelayMetricsEnvelope.
type LLMRelayMetricsEnvelope struct {
	Data LLMRelayMetrics `json:"data"`
}

// LLMUpstream defines model for LLMUpstream.
type LLMUpstream struct {
	// APIKeyPrefix Display-only prefix for the encrypted upstream API key. The full key is never returned.
	APIKeyPrefix string    `json:"apiKeyPrefix,omitempty"`
	BaseURL      string    `json:"baseUrl"`
	CreatedAt    time.Time `json:"createdAt"`
	CreatedBy    string    `json:"createdBy"`

	// DefaultHeaders Non-sensitive upstream headers only.
	DefaultHeaders map[string]any `json:"defaultHeaders,omitempty"`
	Health         map[string]any `json:"health,omitempty"`
	ID             string         `json:"id"`
	MaxConcurrency int            `json:"maxConcurrency"`
	Metadata       map[string]any `json:"metadata,omitempty"`
	Name           string         `json:"name"`

	// Priority Lower values are selected first.
	Priority             int                     `json:"priority"`
	ProviderKind         LLMUpstreamProviderKind `json:"providerKind"`
	ProxyURL             string                  `json:"proxyUrl,omitempty"`
	Status               LLMUpstreamStatus       `json:"status"`
	StreamTimeoutSeconds int                     `json:"streamTimeoutSeconds"`
	SupportedModels      []string                `json:"supportedModels"`
	TimeoutSeconds       int                     `json:"timeoutSeconds"`
	UpdatedAt            time.Time               `json:"updatedAt"`
	Weight               int                     `json:"weight"`
}

// LLMUpstreamProviderKind defines model for LLMUpstream.ProviderKind.
type LLMUpstreamProviderKind string

// LLMUpstreamStatus defines model for LLMUpstream.Status.
type LLMUpstreamStatus string

// LLMUpstreamEnvelope defines model for LLMUpstreamEnvelope.
type LLMUpstreamEnvelope struct {
	Data LLMUpstream `json:"data"`
}

// LLMUpstreamInput defines model for LLMUpstreamInput.
type LLMUpstreamInput struct {
	// APIKey Upstream provider key. It must be encrypted at rest and never returned by read APIs.
	APIKey               string                       `json:"apiKey,omitempty"`
	BaseURL              string                       `json:"baseUrl"`
	DefaultHeaders       map[string]any               `json:"defaultHeaders,omitempty"`
	Health               map[string]any               `json:"health,omitempty"`
	ID                   string                       `json:"id,omitempty"`
	MaxConcurrency       int                          `json:"maxConcurrency,omitempty"`
	Metadata             map[string]any               `json:"metadata,omitempty"`
	Name                 string                       `json:"name"`
	Priority             int                          `json:"priority,omitempty"`
	ProviderKind         LLMUpstreamInputProviderKind `json:"providerKind"`
	ProxyURL             string                       `json:"proxyUrl,omitempty"`
	Status               LLMUpstreamInputStatus       `json:"status,omitempty"`
	StreamTimeoutSeconds int                          `json:"streamTimeoutSeconds,omitempty"`
	SupportedModels      []string                     `json:"supportedModels,omitempty"`
	TimeoutSeconds       int                          `json:"timeoutSeconds,omitempty"`
	Weight               int                          `json:"weight,omitempty"`
}

// LLMUpstreamInputProviderKind defines model for LLMUpstreamInput.ProviderKind.
type LLMUpstreamInputProviderKind string

// LLMUpstreamInputStatus defines model for LLMUpstreamInput.Status.
type LLMUpstreamInputStatus string

// LLMUpstreamListEnvelope defines model for LLMUpstreamListEnvelope.
type LLMUpstreamListEnvelope struct {
	Items []LLMUpstream `json:"items"`
}

// LLMUpstreamTestRequest defines model for LLMUpstreamTestRequest.
type LLMUpstreamTestRequest struct {
	Endpoint       LLMUpstreamTestRequestEndpoint `json:"endpoint,omitempty"`
	Model          string                         `json:"model,omitempty"`
	TimeoutSeconds int                            `json:"timeoutSeconds,omitempty"`
}

// LLMUpstreamTestRequestEndpoint defines model for LLMUpstreamTestRequest.Endpoint.
type LLMUpstreamTestRequestEndpoint string

// LLMUpstreamTestResult defines model for LLMUpstreamTestResult.
type LLMUpstreamTestResult struct {
	CheckedAt  *time.Time `json:"checkedAt,omitempty"`
	DurationMs int64      `json:"durationMs"`
	ErrorCode  string     `json:"errorCode,omitempty"`

	// ErrorMessage Redacted error summary.
	ErrorMessage   string                            `json:"errorMessage,omitempty"`
	ModelCount     int                               `json:"modelCount,omitempty"`
	ProviderKind   LLMUpstreamTestResultProviderKind `json:"providerKind"`
	Status         LLMUpstreamTestResultStatus       `json:"status"`
	UpstreamStatus int                               `json:"upstreamStatus,omitempty"`
}

// LLMUpstreamTestResultProviderKind defines model for LLMUpstreamTestResult.ProviderKind.
type LLMUpstreamTestResultProviderKind string

// LLMUpstreamTestResultStatus defines model for LLMUpstreamTestResult.Status.
type LLMUpstreamTestResultStatus string

// LLMUpstreamTestResultEnvelope defines model for LLMUpstreamTestResultEnvelope.
type LLMUpstreamTestResultEnvelope struct {
	Data LLMUpstreamTestResult `json:"data"`
}

// LoginOptions defines model for LoginOptions.
type LoginOptions struct {
	LocalPasswordLoginEnabled bool `json:"localPasswordLoginEnabled,omitempty"`
	Verification              struct {
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

// MarketplaceAdvisory defines model for MarketplaceAdvisory.
type MarketplaceAdvisory struct {
	AffectedVersions []string                    `json:"affectedVersions,omitempty"`
	ID               string                      `json:"id"`
	PublishedAt      *time.Time                  `json:"publishedAt,omitempty"`
	Severity         MarketplaceAdvisorySeverity `json:"severity"`
	Summary          string                      `json:"summary"`
	URL              string                      `json:"url,omitempty"`
	Withdrawn        bool                        `json:"withdrawn,omitempty"`
}

// MarketplaceAdvisorySeverity defines model for MarketplaceAdvisory.Severity.
type MarketplaceAdvisorySeverity string

// MarketplaceCatalog defines model for MarketplaceCatalog.
type MarketplaceCatalog struct {
	Advisories    []MarketplaceAdvisory `json:"advisories,omitempty"`
	GeneratedAt   time.Time             `json:"generatedAt"`
	Metadata      map[string]any        `json:"metadata,omitempty"`
	Plugins       []MarketplacePlugin   `json:"plugins"`
	SchemaVersion string                `json:"schemaVersion"`
	SourceID      string                `json:"sourceId"`
	SourceURL     string                `json:"sourceUrl,omitempty"`
}

// MarketplacePlugin defines model for MarketplacePlugin.
type MarketplacePlugin struct {
	Advisories    []MarketplaceAdvisory      `json:"advisories,omitempty"`
	Categories    []string                   `json:"categories,omitempty"`
	Compatibility *PluginCompatibility       `json:"compatibility,omitempty"`
	Deprecated    bool                       `json:"deprecated,omitempty"`
	ID            string                     `json:"id"`
	Installed     bool                       `json:"installed,omitempty"`
	LatestVersion string                     `json:"latestVersion,omitempty"`
	Manifest      PluginManifest             `json:"manifest"`
	Name          string                     `json:"name"`
	Publisher     string                     `json:"publisher"`
	PublisherInfo *MarketplacePublisher      `json:"publisherInfo,omitempty"`
	RiskLevel     string                     `json:"riskLevel,omitempty"`
	Source        string                     `json:"source"`
	SourceID      string                     `json:"sourceId,omitempty"`
	SourceURL     string                     `json:"sourceUrl,omitempty"`
	Summary       string                     `json:"summary,omitempty"`
	Suspended     bool                       `json:"suspended,omitempty"`
	Type          string                     `json:"type"`
	Verified      bool                       `json:"verified,omitempty"`
	Version       string                     `json:"version"`
	Versions      []MarketplacePluginVersion `json:"versions,omitempty"`
}

// MarketplacePluginEnvelope defines model for MarketplacePluginEnvelope.
type MarketplacePluginEnvelope struct {
	Data MarketplacePlugin `json:"data"`
}

// MarketplacePluginListEnvelope defines model for MarketplacePluginListEnvelope.
type MarketplacePluginListEnvelope struct {
	Items []MarketplacePlugin `json:"items"`
}

// MarketplacePluginVersion defines model for MarketplacePluginVersion.
type MarketplacePluginVersion struct {
	Checksum       string         `json:"checksum,omitempty"`
	Deprecated     bool           `json:"deprecated,omitempty"`
	ManifestURL    string         `json:"manifestUrl"`
	Metadata       map[string]any `json:"metadata,omitempty"`
	MinSohaVersion string         `json:"minSohaVersion,omitempty"`
	PackageURL     string         `json:"packageUrl,omitempty"`
	PublishedAt    *time.Time     `json:"publishedAt,omitempty"`
	Signature      string         `json:"signature,omitempty"`
	Suspended      bool           `json:"suspended,omitempty"`
	Version        string         `json:"version"`
}

// MarketplacePublisher defines model for MarketplacePublisher.
type MarketplacePublisher struct {
	ID                string                                `json:"id"`
	Name              string                                `json:"name"`
	URL               string                                `json:"url,omitempty"`
	VerificationLevel MarketplacePublisherVerificationLevel `json:"verificationLevel,omitempty"`
	Verified          bool                                  `json:"verified,omitempty"`
}

// MarketplacePublisherVerificationLevel defines model for MarketplacePublisher.VerificationLevel.
type MarketplacePublisherVerificationLevel string

// MultiAgentBudget defines model for MultiAgentBudget.
type MultiAgentBudget struct {
	DeadlineSeconds int     `json:"deadlineSeconds"`
	MaxCost         float32 `json:"maxCost"`
	MaxSteps        int     `json:"maxSteps"`
	MaxTokens       int     `json:"maxTokens"`
}

// MultiAgentPlan defines model for MultiAgentPlan.
type MultiAgentPlan struct {
	CompletedAt          *time.Time          `json:"completedAt,omitempty"`
	CoordinatorRef       string              `json:"coordinatorRef"`
	CreatedAt            time.Time           `json:"createdAt"`
	ID                   string              `json:"id"`
	MergedOutputRefs     []string            `json:"mergedOutputRefs,omitempty"`
	PrincipalPermissions []string            `json:"principalPermissions"`
	SharedBudget         MultiAgentBudget    `json:"sharedBudget"`
	Status               string              `json:"status"`
	Subtasks             []MultiAgentSubtask `json:"subtasks"`
}

// MultiAgentPlanEnvelope defines model for MultiAgentPlanEnvelope.
type MultiAgentPlanEnvelope struct {
	Data MultiAgentPlan `json:"data"`
}

// MultiAgentPlanListEnvelope defines model for MultiAgentPlanListEnvelope.
type MultiAgentPlanListEnvelope struct {
	Items []MultiAgentPlan `json:"items"`
}

// MultiAgentSubtask defines model for MultiAgentSubtask.
type MultiAgentSubtask struct {
	AgentProfileRef string           `json:"agentProfileRef"`
	Budget          MultiAgentBudget `json:"budget"`
	DependsOn       []string         `json:"dependsOn,omitempty"`
	ErrorCode       string           `json:"errorCode,omitempty"`
	ID              string           `json:"id"`
	Input           string           `json:"input"`
	OutputRef       string           `json:"outputRef,omitempty"`
	PermissionKeys  []string         `json:"permissionKeys"`
	Status          string           `json:"status"`
}

// MultiAgentSubtaskCompletionInput defines model for MultiAgentSubtaskCompletionInput.
type MultiAgentSubtaskCompletionInput struct {
	OutputRef string `json:"outputRef"`
}

// NativeProviderObject Native provider-compatible JSON. Unknown fields are preserved and responses are not wrapped in an OpenSoha envelope.
type NativeProviderObject map[string]any

// OIDCExchangeRequest defines model for OIDCExchangeRequest.
type OIDCExchangeRequest struct {
	Code string `json:"code"`
}

// OpenAIAudioSpeechRequest defines model for OpenAIAudioSpeechRequest.
type OpenAIAudioSpeechRequest struct {
	Input        string `json:"input"`
	Instructions string `json:"instructions,omitempty"`

	// Model Public speech model ID exposed by Soha; the relay maps it to the configured upstream model.
	Model                string         `json:"model"`
	ResponseFormat       string         `json:"response_format,omitempty"`
	Speed                float32        `json:"speed,omitempty"`
	User                 string         `json:"user,omitempty"`
	Voice                string         `json:"voice"`
	AdditionalProperties map[string]any `json:"-"`
}

// OpenAIAudioTranscriptionRequest defines model for OpenAIAudioTranscriptionRequest.
type OpenAIAudioTranscriptionRequest struct {
	File     openapi_types.File `json:"file"`
	Include  []string           `json:"include,omitempty"`
	Language string             `json:"language,omitempty"`

	// Model Public transcription model ID exposed by Soha; the relay maps it to the configured upstream model.
	Model                  string                                                  `json:"model"`
	Prompt                 string                                                  `json:"prompt,omitempty"`
	ResponseFormat         string                                                  `json:"response_format,omitempty"`
	Temperature            float32                                                 `json:"temperature,omitempty"`
	TimestampGranularities []OpenAIAudioTranscriptionRequestTimestampGranularities `json:"timestamp_granularities,omitempty"`
	AdditionalProperties   map[string]any                                          `json:"-"`
}

// OpenAIAudioTranscriptionRequestTimestampGranularities defines model for OpenAIAudioTranscriptionRequest.TimestampGranularities.
type OpenAIAudioTranscriptionRequestTimestampGranularities string

// OpenAIAudioTranslationRequest defines model for OpenAIAudioTranslationRequest.
type OpenAIAudioTranslationRequest struct {
	File openapi_types.File `json:"file"`

	// Model Public translation model ID exposed by Soha; the relay maps it to the configured upstream model.
	Model                string         `json:"model"`
	Prompt               string         `json:"prompt,omitempty"`
	ResponseFormat       string         `json:"response_format,omitempty"`
	Temperature          float32        `json:"temperature,omitempty"`
	AdditionalProperties map[string]any `json:"-"`
}

// OpenAIChatCompletionRequest defines model for OpenAIChatCompletionRequest.
type OpenAIChatCompletionRequest struct {
	Messages             []map[string]any `json:"messages"`
	Model                string           `json:"model"`
	Stream               bool             `json:"stream,omitempty"`
	StreamOptions        map[string]any   `json:"stream_options,omitempty"`
	AdditionalProperties map[string]any   `json:"-"`
}

// OpenAIEmbeddingRequest defines model for OpenAIEmbeddingRequest.
type OpenAIEmbeddingRequest struct {
	Input                AnyValue       `json:"input"`
	Model                string         `json:"model"`
	AdditionalProperties map[string]any `json:"-"`
}

// OpenAIImageEditRequest defines model for OpenAIImageEditRequest.
type OpenAIImageEditRequest struct {
	Image openapi_types.File `json:"image"`
	Mask  openapi_types.File `json:"mask,omitempty"`

	// Model Public image edit model ID exposed by Soha; the relay maps it to the configured upstream model.
	Model                string         `json:"model"`
	N                    int            `json:"n,omitempty"`
	Prompt               string         `json:"prompt"`
	Quality              string         `json:"quality,omitempty"`
	ResponseFormat       string         `json:"response_format,omitempty"`
	Size                 string         `json:"size,omitempty"`
	User                 string         `json:"user,omitempty"`
	AdditionalProperties map[string]any `json:"-"`
}

// OpenAIImageGenerationRequest defines model for OpenAIImageGenerationRequest.
type OpenAIImageGenerationRequest struct {
	Background string `json:"background,omitempty"`

	// Model Public image model ID exposed by Soha; the relay maps it to the configured upstream model.
	Model                string         `json:"model"`
	Moderation           string         `json:"moderation,omitempty"`
	N                    int            `json:"n,omitempty"`
	OutputCompression    int            `json:"output_compression,omitempty"`
	OutputFormat         string         `json:"output_format,omitempty"`
	Prompt               string         `json:"prompt"`
	Quality              string         `json:"quality,omitempty"`
	ResponseFormat       string         `json:"response_format,omitempty"`
	Size                 string         `json:"size,omitempty"`
	Style                string         `json:"style,omitempty"`
	User                 string         `json:"user,omitempty"`
	AdditionalProperties map[string]any `json:"-"`
}

// OpenAIImageVariationRequest defines model for OpenAIImageVariationRequest.
type OpenAIImageVariationRequest struct {
	Image openapi_types.File `json:"image"`

	// Model Public image variation model ID exposed by Soha; the relay maps it to the configured upstream model.
	Model                string         `json:"model"`
	N                    int            `json:"n,omitempty"`
	ResponseFormat       string         `json:"response_format,omitempty"`
	Size                 string         `json:"size,omitempty"`
	User                 string         `json:"user,omitempty"`
	AdditionalProperties map[string]any `json:"-"`
}

// OpenAIModel defines model for OpenAIModel.
type OpenAIModel struct {
	Created              int            `json:"created,omitempty"`
	ID                   string         `json:"id"`
	Object               string         `json:"object"`
	OwnedBy              string         `json:"owned_by,omitempty"`
	AdditionalProperties map[string]any `json:"-"`
}

// OpenAIModelsResponse defines model for OpenAIModelsResponse.
type OpenAIModelsResponse struct {
	Data                 []OpenAIModel              `json:"data"`
	Object               OpenAIModelsResponseObject `json:"object"`
	AdditionalProperties map[string]any             `json:"-"`
}

// OpenAIModelsResponseObject defines model for OpenAIModelsResponse.Object.
type OpenAIModelsResponseObject string

// OpenAIResponseRequest defines model for OpenAIResponseRequest.
type OpenAIResponseRequest struct {
	Input                AnyValue       `json:"input,omitempty"`
	Model                string         `json:"model"`
	Stream               bool           `json:"stream,omitempty"`
	AdditionalProperties map[string]any `json:"-"`
}

// OperationState defines model for OperationState.
type OperationState struct {
	Cancelable             bool       `json:"cancelable"`
	ClaimedByAgentID       string     `json:"claimedByAgentId,omitempty"`
	FailureMessage         string     `json:"failureMessage,omitempty"`
	FailureReason          string     `json:"failureReason,omitempty"`
	FinalStateRecordedAt   *time.Time `json:"finalStateRecordedAt,omitempty"`
	HeartbeatRequired      bool       `json:"heartbeatRequired"`
	HeartbeatStale         bool       `json:"heartbeatStale"`
	LastHeartbeatAt        *time.Time `json:"lastHeartbeatAt,omitempty"`
	LastRuntimeSeenAt      *time.Time `json:"lastRuntimeSeenAt,omitempty"`
	NextHeartbeatDeadline  *time.Time `json:"nextHeartbeatDeadline,omitempty"`
	Phase                  string     `json:"phase"`
	RecommendedNextAction  string     `json:"recommendedNextAction,omitempty"`
	Retryable              bool       `json:"retryable"`
	RuntimeEndpointPresent bool       `json:"runtimeEndpointPresent"`
	Status                 string     `json:"status"`
	Terminal               bool       `json:"terminal"`
	TimeoutSeconds         int        `json:"timeoutSeconds"`
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

// PluginAIExtensions defines model for PluginAIExtensions.
type PluginAIExtensions struct {
	AgentProviders    []PluginExtensionContribution `json:"agentProviders,omitempty"`
	AnalysisWorkflows []PluginExtensionContribution `json:"analysisWorkflows,omitempty"`
	ArtifactRenderers []PluginExtensionContribution `json:"artifactRenderers,omitempty"`
	MCPPresets        []PluginExtensionContribution `json:"mcpPresets,omitempty"`
	ModelProviders    []PluginExtensionContribution `json:"modelProviders,omitempty"`
	SkillPacks        []PluginExtensionContribution `json:"skillPacks,omitempty"`
	ToolProviders     []PluginExtensionContribution `json:"toolProviders,omitempty"`
}

// PluginAlertExtensions defines model for PluginAlertExtensions.
type PluginAlertExtensions struct {
	Enrichers            []PluginExtensionContribution `json:"enrichers,omitempty"`
	EscalationProviders  []PluginExtensionContribution `json:"escalationProviders,omitempty"`
	NotificationChannels []PluginExtensionContribution `json:"notificationChannels,omitempty"`
	Receivers            []PluginExtensionContribution `json:"receivers,omitempty"`
	SilenceAdapters      []PluginExtensionContribution `json:"silenceAdapters,omitempty"`
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

// PluginAuthExtensions defines model for PluginAuthExtensions.
type PluginAuthExtensions struct {
	DirectorySync  []PluginExtensionContribution `json:"directorySync,omitempty"`
	ProfileMappers []PluginExtensionContribution `json:"profileMappers,omitempty"`
	Sources        []PluginExtensionContribution `json:"sources,omitempty"`
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

// PluginComputeContainerRuntimeProvider defines model for PluginComputeContainerRuntimeProvider.
type PluginComputeContainerRuntimeProvider struct {
	ActionRefs      map[string]string                                    `json:"actionRefs,omitempty"`
	ActivationLevel ComputeProviderActivationLevel                       `json:"activationLevel"`
	Capabilities    []string                                             `json:"capabilities"`
	ConfigSchemaRef string                                               `json:"configSchemaRef,omitempty"`
	Description     string                                               `json:"description,omitempty"`
	DisplayName     string                                               `json:"displayName"`
	ProviderKey     string                                               `json:"providerKey"`
	ResourceKinds   []PluginComputeContainerRuntimeProviderResourceKinds `json:"resourceKinds"`
}

// PluginComputeContainerRuntimeProviderResourceKinds defines model for PluginComputeContainerRuntimeProvider.ResourceKinds.
type PluginComputeContainerRuntimeProviderResourceKinds string

// PluginComputeExtensions defines model for PluginComputeExtensions.
type PluginComputeExtensions struct {
	ContainerRuntimeProviders []PluginComputeContainerRuntimeProvider `json:"containerRuntimeProviders,omitempty"`
	VirtualizationProviders   []PluginComputeVirtualizationProvider   `json:"virtualizationProviders,omitempty"`
}

// PluginComputeVirtualizationProvider defines model for PluginComputeVirtualizationProvider.
type PluginComputeVirtualizationProvider struct {
	ActionRefs      map[string]string              `json:"actionRefs,omitempty"`
	ActivationLevel ComputeProviderActivationLevel `json:"activationLevel"`

	// Capabilities Provider-declared VM capabilities. Standard resource mutation keys are vm.resource.cpu.resize, vm.resource.memory.resize, vm.resource.disk.add, vm.resource.disk.resize, vm.resource.network.add, and vm.resource.network.remove.
	Capabilities    []string                                           `json:"capabilities"`
	ConfigSchemaRef string                                             `json:"configSchemaRef,omitempty"`
	Description     string                                             `json:"description,omitempty"`
	DisplayName     string                                             `json:"displayName"`
	ProviderKey     string                                             `json:"providerKey"`
	ResourceKinds   []PluginComputeVirtualizationProviderResourceKinds `json:"resourceKinds"`
}

// PluginComputeVirtualizationProviderResourceKinds defines model for PluginComputeVirtualizationProvider.ResourceKinds.
type PluginComputeVirtualizationProviderResourceKinds string

// PluginConfigRequest defines model for PluginConfigRequest.
type PluginConfigRequest struct {
	Enabled    *bool             `json:"enabled,omitempty"`
	Metadata   map[string]any    `json:"metadata,omitempty"`
	SecretRefs map[string]string `json:"secretRefs,omitempty"`
}

// PluginDeliveryExtensions defines model for PluginDeliveryExtensions.
type PluginDeliveryExtensions struct {
	ArtifactStores   []PluginExtensionContribution `json:"artifactStores,omitempty"`
	BuildProviders   []PluginExtensionContribution `json:"buildProviders,omitempty"`
	DeployStrategies []PluginExtensionContribution `json:"deployStrategies,omitempty"`
	ReleaseGates     []PluginExtensionContribution `json:"releaseGates,omitempty"`
	ScanProviders    []PluginExtensionContribution `json:"scanProviders,omitempty"`
}

// PluginExtensionContribution defines model for PluginExtensionContribution.
type PluginExtensionContribution struct {
	ActionRef      string         `json:"actionRef,omitempty"`
	ConfigSchema   JSONSchema     `json:"configSchema,omitempty"`
	Description    string         `json:"description,omitempty"`
	ID             string         `json:"id"`
	Label          string         `json:"label,omitempty"`
	Match          JSONSchema     `json:"match,omitempty"`
	Metadata       map[string]any `json:"metadata,omitempty"`
	PermissionKeys []string       `json:"permissionKeys,omitempty"`
	ResourceKinds  []string       `json:"resourceKinds,omitempty"`
	UISchema       JSONSchema     `json:"uiSchema,omitempty"`
}

// PluginExtensionPoints defines model for PluginExtensionPoints.
type PluginExtensionPoints struct {
	AI       *PluginAIExtensions       `json:"ai,omitempty"`
	Alerts   *PluginAlertExtensions    `json:"alerts,omitempty"`
	Auth     *PluginAuthExtensions     `json:"auth,omitempty"`
	Compute  *PluginComputeExtensions  `json:"compute,omitempty"`
	Delivery *PluginDeliveryExtensions `json:"delivery,omitempty"`
	Gateway  *PluginGatewayExtensions  `json:"gateway,omitempty"`
	Identity *PluginIdentityExtensions `json:"identity,omitempty"`
	Metrics  *PluginMetricsExtensions  `json:"metrics,omitempty"`
	Resource *PluginResourceExtensions `json:"resource,omitempty"`
	UI       *PluginUIExtensions       `json:"ui,omitempty"`
}

// PluginGatewayExtensions defines model for PluginGatewayExtensions.
type PluginGatewayExtensions struct {
	Policies  []PluginExtensionContribution `json:"policies,omitempty"`
	Prompts   []PluginExtensionContribution `json:"prompts,omitempty"`
	Resources []PluginExtensionContribution `json:"resources,omitempty"`
	Tools     []PluginExtensionContribution `json:"tools,omitempty"`
}

// PluginIdentityExtensions defines model for PluginIdentityExtensions.
type PluginIdentityExtensions struct {
	ApplicationTemplates []PluginExtensionContribution `json:"applicationTemplates,omitempty"`
	ProviderTemplates    []PluginExtensionContribution `json:"providerTemplates,omitempty"`
}

// PluginInstallRequest defines model for PluginInstallRequest.
type PluginInstallRequest struct {
	Enable           bool            `json:"enable,omitempty"`
	ExpectedChecksum string          `json:"expectedChecksum,omitempty"`
	Manifest         *PluginManifest `json:"manifest,omitempty"`
	MarketplaceURL   string          `json:"marketplaceUrl,omitempty"`
	PluginID         string          `json:"pluginId,omitempty"`
	Source           string          `json:"source,omitempty"`
	SourceID         string          `json:"sourceId,omitempty"`
	Version          string          `json:"version,omitempty"`
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
	Assets          *PluginAssetSnapshot     `json:"assets,omitempty"`
	Capabilities    *PluginCapabilityRequest `json:"capabilities,omitempty"`
	Compatibility   *PluginCompatibility     `json:"compatibility,omitempty"`
	ConfigSchema    JSONSchema               `json:"configSchema,omitempty"`
	Description     string                   `json:"description,omitempty"`
	ExtensionPoints *PluginExtensionPoints   `json:"extensionPoints,omitempty"`
	Homepage        string                   `json:"homepage,omitempty"`
	ID              string                   `json:"id"`
	Integrity       *PluginIntegrity         `json:"integrity,omitempty"`
	Metadata        map[string]any           `json:"metadata,omitempty"`
	Name            string                   `json:"name"`
	Permissions     *PluginPermissionRequest `json:"permissions,omitempty"`
	Publisher       string                   `json:"publisher"`
	Runtime         *PluginRuntimeSpec       `json:"runtime,omitempty"`
	Secrets         *struct {
		Required []PluginSecretRequirement `json:"required,omitempty"`
	} `json:"secrets,omitempty"`
	Type    string `json:"type"`
	Version string `json:"version"`
}

// PluginManifestEnvelope defines model for PluginManifestEnvelope.
type PluginManifestEnvelope struct {
	Data PluginManifest `json:"data"`
}

// PluginMetricsExtensions defines model for PluginMetricsExtensions.
type PluginMetricsExtensions struct {
	Definitions []PluginExtensionContribution `json:"definitions,omitempty"`
	Enrichers   []PluginExtensionContribution `json:"enrichers,omitempty"`
	Panels      []PluginExtensionContribution `json:"panels,omitempty"`
	Providers   []PluginExtensionContribution `json:"providers,omitempty"`
}

// PluginPermissionRequest defines model for PluginPermissionRequest.
type PluginPermissionRequest struct {
	Domain   []string `json:"domain,omitempty"`
	Required []string `json:"required,omitempty"`
}

// PluginResourceExtensions defines model for PluginResourceExtensions.
type PluginResourceExtensions struct {
	Actions     []PluginExtensionContribution `json:"actions,omitempty"`
	Diagnostics []PluginExtensionContribution `json:"diagnostics,omitempty"`
	Tabs        []PluginExtensionContribution `json:"tabs,omitempty"`
	Tags        []PluginExtensionContribution `json:"tags,omitempty"`
}

// PluginRuntimeSpec defines model for PluginRuntimeSpec.
type PluginRuntimeSpec struct {
	ActionPath     string                `json:"actionPath,omitempty"`
	Endpoint       string                `json:"endpoint,omitempty"`
	HealthPath     string                `json:"healthPath,omitempty"`
	ManifestPath   string                `json:"manifestPath,omitempty"`
	Metadata       map[string]any        `json:"metadata,omitempty"`
	MetricsPath    string                `json:"metricsPath,omitempty"`
	Mode           PluginRuntimeSpecMode `json:"mode"`
	TimeoutSeconds int                   `json:"timeoutSeconds,omitempty"`
	WebhookPath    string                `json:"webhookPath,omitempty"`
}

// PluginRuntimeSpecMode defines model for PluginRuntimeSpec.Mode.
type PluginRuntimeSpecMode string

// PluginSecretRequirement defines model for PluginSecretRequirement.
type PluginSecretRequirement struct {
	Description string `json:"description,omitempty"`
	Name        string `json:"name"`
	Required    bool   `json:"required,omitempty"`
	SecretRef   string `json:"secretRef,omitempty"`
}

// PluginUIExtensions defines model for PluginUIExtensions.
type PluginUIExtensions struct {
	ActionButtons []PluginExtensionContribution `json:"actionButtons,omitempty"`
	DetailPanels  []PluginExtensionContribution `json:"detailPanels,omitempty"`
	Menus         []PluginExtensionContribution `json:"menus,omitempty"`
	SettingsForms []PluginExtensionContribution `json:"settingsForms,omitempty"`
	StatusCards   []PluginExtensionContribution `json:"statusCards,omitempty"`
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

// ReleaseBundle defines model for ReleaseBundle.
type ReleaseBundle struct {
	ApplicationEnvironmentID string              `json:"applicationEnvironmentId,omitempty"`
	ApplicationID            string              `json:"applicationId"`
	ArtifactDigest           string              `json:"artifactDigest,omitempty"`
	ArtifactRef              string              `json:"artifactRef,omitempty"`
	Artifacts                []ExecutionArtifact `json:"artifacts,omitempty"`
	CreatedAt                time.Time           `json:"createdAt"`
	ID                       string              `json:"id"`
	Metadata                 map[string]any      `json:"metadata,omitempty"`
	SourceType               string              `json:"sourceType"`
	Status                   string              `json:"status"`
	UpdatedAt                time.Time           `json:"updatedAt"`
	Version                  string              `json:"version"`
}

// ReleaseBundleEnvelope defines model for ReleaseBundleEnvelope.
type ReleaseBundleEnvelope struct {
	Data ReleaseBundle `json:"data"`
}

// ReleaseBundleListEnvelope defines model for ReleaseBundleListEnvelope.
type ReleaseBundleListEnvelope struct {
	Data []ReleaseBundle `json:"data"`
}

// ReleasePolicy defines model for ReleasePolicy.
type ReleasePolicy struct {
	ActionKind            string   `json:"actionKind,omitempty"`
	ApproverRoles         []string `json:"approverRoles,omitempty"`
	AutoRollback          bool     `json:"autoRollback,omitempty"`
	RequiresApproval      bool     `json:"requiresApproval,omitempty"`
	RolloutTimeoutSeconds int      `json:"rolloutTimeoutSeconds,omitempty"`
	VerificationMode      string   `json:"verificationMode,omitempty"`
}

// ReleaseRecord defines model for ReleaseRecord.
type ReleaseRecord struct {
	ApplicationID  string         `json:"applicationId"`
	ClusterID      string         `json:"clusterId"`
	CreatedAt      time.Time      `json:"createdAt"`
	DeployedAt     *time.Time     `json:"deployedAt,omitempty"`
	DeploymentName string         `json:"deploymentName"`
	ID             string         `json:"id"`
	Metadata       map[string]any `json:"metadata,omitempty"`
	Namespace      string         `json:"namespace"`
	Status         string         `json:"status"`
}

// ReleaseTarget defines model for ReleaseTarget.
type ReleaseTarget struct {
	ApplicationEnvironmentID string         `json:"applicationEnvironmentId"`
	ClusterID                string         `json:"clusterId"`
	ConfigRef                string         `json:"configRef,omitempty"`
	ContainerName            string         `json:"containerName,omitempty"`
	CreatedAt                *time.Time     `json:"createdAt,omitempty"`
	Enabled                  bool           `json:"enabled"`
	ExecutorKind             string         `json:"executorKind,omitempty"`
	GroupKey                 string         `json:"groupKey,omitempty"`
	ID                       string         `json:"id"`
	Metadata                 map[string]any `json:"metadata,omitempty"`
	Namespace                string         `json:"namespace"`
	RegionKey                string         `json:"regionKey,omitempty"`
	TargetKind               string         `json:"targetKind,omitempty"`
	UpdatedAt                *time.Time     `json:"updatedAt,omitempty"`
	WaveKey                  string         `json:"waveKey,omitempty"`
	WorkloadKind             string         `json:"workloadKind"`
	WorkloadName             string         `json:"workloadName"`
}

// ReleaseTargetInput defines model for ReleaseTargetInput.
type ReleaseTargetInput struct {
	ClusterID     string         `json:"clusterId"`
	ConfigRef     string         `json:"configRef,omitempty"`
	ContainerName string         `json:"containerName,omitempty"`
	Enabled       bool           `json:"enabled"`
	ExecutorKind  string         `json:"executorKind,omitempty"`
	GroupKey      string         `json:"groupKey,omitempty"`
	ID            string         `json:"id,omitempty"`
	Metadata      map[string]any `json:"metadata,omitempty"`
	Namespace     string         `json:"namespace"`
	RegionKey     string         `json:"regionKey,omitempty"`
	TargetKind    string         `json:"targetKind,omitempty"`
	WaveKey       string         `json:"waveKey,omitempty"`
	WorkloadKind  string         `json:"workloadKind"`
	WorkloadName  string         `json:"workloadName"`
}

// Repository defines model for Repository.
type Repository struct {
	ApplicationIDs []string  `json:"applicationIds,omitempty"`
	CreatedAt      time.Time `json:"createdAt"`

	// CredentialRef Opaque server-side credential reference; never a token or private key.
	CredentialRef   string             `json:"credentialRef,omitempty"`
	DefaultBranch   string             `json:"defaultBranch"`
	GitlabProjectID string             `json:"gitlabProjectId,omitempty"`
	ID              string             `json:"id"`
	Name            string             `json:"name"`
	Path            string             `json:"path"`
	Protocol        RepositoryProtocol `json:"protocol"`
	Provider        RepositoryProvider `json:"provider"`
	UpdatedAt       time.Time          `json:"updatedAt"`

	// URL HTTPS or SSH Git URL. Credentials are resolved only through credentialRef.
	URL string `json:"url"`
}

// RepositoryEnvelope defines model for RepositoryEnvelope.
type RepositoryEnvelope struct {
	Data Repository `json:"data"`
}

// RepositoryInput defines model for RepositoryInput.
type RepositoryInput struct {
	ApplicationIDs []string `json:"applicationIds,omitempty"`

	// CredentialRef Opaque server-side credential reference; never a token or private key.
	CredentialRef   string             `json:"credentialRef,omitempty"`
	DefaultBranch   string             `json:"defaultBranch"`
	GitlabProjectID string             `json:"gitlabProjectId,omitempty"`
	Name            string             `json:"name"`
	Path            string             `json:"path"`
	Protocol        RepositoryProtocol `json:"protocol"`
	Provider        RepositoryProvider `json:"provider"`

	// URL HTTPS or SSH Git URL without embedded credentials.
	URL string `json:"url"`
}

// RepositoryListEnvelope defines model for RepositoryListEnvelope.
type RepositoryListEnvelope struct {
	Data []Repository `json:"data"`
}

// RepositoryProtocol defines model for RepositoryProtocol.
type RepositoryProtocol string

// RepositoryProvider defines model for RepositoryProvider.
type RepositoryProvider string

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

// ResourceSelector defines model for ResourceSelector.
type ResourceSelector struct {
	MatchLabels map[string]string `json:"matchLabels,omitempty"`
}

// RiskLevel defines model for RiskLevel.
type RiskLevel = string

// RuntimeCPUUsage defines model for RuntimeCPUUsage.
type RuntimeCPUUsage struct {
	LogicalCores int     `json:"logicalCores"`
	UsagePercent float64 `json:"usagePercent"`
}

// RuntimeConfigApplication defines model for RuntimeConfigApplication.
type RuntimeConfigApplication struct {
	CreatedAt  time.Time                      `json:"createdAt"`
	Error      string                         `json:"error,omitempty"`
	ID         string                         `json:"id"`
	Items      []RuntimeConfigAppliedItem     `json:"items"`
	RevisionID string                         `json:"revisionId"`
	Status     RuntimeConfigApplicationStatus `json:"status"`
	UpdatedAt  time.Time                      `json:"updatedAt"`
	Version    int64                          `json:"version"`
}

// RuntimeConfigApplicationEnvelope defines model for RuntimeConfigApplicationEnvelope.
type RuntimeConfigApplicationEnvelope struct {
	Data RuntimeConfigApplication `json:"data"`
}

// RuntimeConfigApplicationStatus defines model for RuntimeConfigApplicationStatus.
type RuntimeConfigApplicationStatus string

// RuntimeConfigAppliedItem defines model for RuntimeConfigAppliedItem.
type RuntimeConfigAppliedItem struct {
	ApplyMode RuntimeConfigApplyMode         `json:"applyMode"`
	Key       string                         `json:"key"`
	Message   string                         `json:"message,omitempty"`
	Status    RuntimeConfigApplicationStatus `json:"status"`
}

// RuntimeConfigApplyMode defines model for RuntimeConfigApplyMode.
type RuntimeConfigApplyMode string

// RuntimeConfigApplyResult defines model for RuntimeConfigApplyResult.
type RuntimeConfigApplyResult struct {
	Application RuntimeConfigApplication `json:"application"`
	Revision    RuntimeConfigRevision    `json:"revision"`
}

// RuntimeConfigApplyResultEnvelope defines model for RuntimeConfigApplyResultEnvelope.
type RuntimeConfigApplyResultEnvelope struct {
	Data RuntimeConfigApplyResult `json:"data"`
}

// RuntimeConfigChange defines model for RuntimeConfigChange.
type RuntimeConfigChange struct {
	Key string `json:"key"`

	// Reset Removes the runtime override so the next lower-precedence source becomes effective.
	Reset bool               `json:"reset,omitempty"`
	Value RuntimeConfigValue `json:"value,omitempty"`
}

// RuntimeConfigChangeRequest defines model for RuntimeConfigChangeRequest.
type RuntimeConfigChangeRequest struct {
	Changes         []RuntimeConfigChange `json:"changes"`
	ExpectedVersion int64                 `json:"expectedVersion"`
	Reason          string                `json:"reason,omitempty"`
}

// RuntimeConfigIssueSeverity defines model for RuntimeConfigIssueSeverity.
type RuntimeConfigIssueSeverity string

// RuntimeConfigItem defines model for RuntimeConfigItem.
type RuntimeConfigItem struct {
	ApplyMode      RuntimeConfigApplyMode `json:"applyMode"`
	Category       string                 `json:"category,omitempty"`
	DefaultValue   RuntimeConfigValue     `json:"defaultValue,omitempty"`
	Description    string                 `json:"description,omitempty"`
	Editable       bool                   `json:"editable"`
	EffectiveValue RuntimeConfigValue     `json:"effectiveValue,omitempty"`
	Key            string                 `json:"key"`
	Label          string                 `json:"label,omitempty"`
	PendingRestart bool                   `json:"pendingRestart"`
	Sensitive      bool                   `json:"sensitive"`
	Source         RuntimeConfigSource    `json:"source"`
	ValueType      RuntimeConfigValueType `json:"valueType"`
}

// RuntimeConfigRevision defines model for RuntimeConfigRevision.
type RuntimeConfigRevision struct {
	Actor                string                         `json:"actor"`
	Changes              []RuntimeConfigChange          `json:"changes"`
	CreatedAt            time.Time                      `json:"createdAt"`
	ID                   string                         `json:"id"`
	Reason               string                         `json:"reason,omitempty"`
	RollbackOfRevisionID string                         `json:"rollbackOfRevisionId,omitempty"`
	Status               RuntimeConfigApplicationStatus `json:"status"`
	Version              int64                          `json:"version"`
}

// RuntimeConfigRevisionListEnvelope defines model for RuntimeConfigRevisionListEnvelope.
type RuntimeConfigRevisionListEnvelope struct {
	Items []RuntimeConfigRevision `json:"items"`
}

// RuntimeConfigRollbackRequest defines model for RuntimeConfigRollbackRequest.
type RuntimeConfigRollbackRequest struct {
	ExpectedVersion int64  `json:"expectedVersion"`
	Reason          string `json:"reason,omitempty"`
	TargetVersion   int64  `json:"targetVersion"`
}

// RuntimeConfigSnapshot defines model for RuntimeConfigSnapshot.
type RuntimeConfigSnapshot struct {
	ActiveRevisionID string              `json:"activeRevisionId,omitempty"`
	Items            []RuntimeConfigItem `json:"items"`
	PendingRestart   bool                `json:"pendingRestart"`
	Version          int64               `json:"version"`
}

// RuntimeConfigSnapshotEnvelope defines model for RuntimeConfigSnapshotEnvelope.
type RuntimeConfigSnapshotEnvelope struct {
	Data RuntimeConfigSnapshot `json:"data"`
}

// RuntimeConfigSource defines model for RuntimeConfigSource.
type RuntimeConfigSource string

// RuntimeConfigValidatedChange defines model for RuntimeConfigValidatedChange.
type RuntimeConfigValidatedChange struct {
	ApplyMode     RuntimeConfigApplyMode `json:"applyMode"`
	CurrentValue  RuntimeConfigValue     `json:"currentValue"`
	Key           string                 `json:"key"`
	ProposedValue RuntimeConfigValue     `json:"proposedValue"`
}

// RuntimeConfigValidationEnvelope defines model for RuntimeConfigValidationEnvelope.
type RuntimeConfigValidationEnvelope struct {
	Data RuntimeConfigValidationResult `json:"data"`
}

// RuntimeConfigValidationIssue defines model for RuntimeConfigValidationIssue.
type RuntimeConfigValidationIssue struct {
	Code     string                     `json:"code"`
	Key      string                     `json:"key,omitempty"`
	Message  string                     `json:"message"`
	Severity RuntimeConfigIssueSeverity `json:"severity"`
}

// RuntimeConfigValidationResult defines model for RuntimeConfigValidationResult.
type RuntimeConfigValidationResult struct {
	Changes         []RuntimeConfigValidatedChange `json:"changes"`
	CurrentVersion  int64                          `json:"currentVersion"`
	ExpectedVersion int64                          `json:"expectedVersion"`
	Issues          []RuntimeConfigValidationIssue `json:"issues"`
	RequiresRestart bool                           `json:"requiresRestart"`
	Valid           bool                           `json:"valid"`
}

// RuntimeConfigValue defines model for RuntimeConfigValue.
type RuntimeConfigValue = any

// RuntimeConfigValueType defines model for RuntimeConfigValueType.
type RuntimeConfigValueType string

// RuntimeDiskUsage defines model for RuntimeDiskUsage.
type RuntimeDiskUsage struct {
	Available      bool    `json:"available"`
	AvailableBytes int64   `json:"availableBytes"`
	Path           string  `json:"path"`
	TotalBytes     int64   `json:"totalBytes"`
	UsagePercent   float64 `json:"usagePercent"`
	UsedBytes      int64   `json:"usedBytes"`
}

// RuntimeGoUsage defines model for RuntimeGoUsage.
type RuntimeGoUsage struct {
	GcCycles   int64 `json:"gcCycles"`
	Gomaxprocs int   `json:"gomaxprocs"`
	Goroutines int   `json:"goroutines"`
}

// RuntimeMemoryUsage defines model for RuntimeMemoryUsage.
type RuntimeMemoryUsage struct {
	GoReservedBytes  int64   `json:"goReservedBytes"`
	HeapAllocBytes   int64   `json:"heapAllocBytes"`
	HeapSysBytes     int64   `json:"heapSysBytes"`
	HeapUsagePercent float64 `json:"heapUsagePercent"`
}

// RuntimeNetworkUsage defines model for RuntimeNetworkUsage.
type RuntimeNetworkUsage struct {
	Available        bool                     `json:"available"`
	RxBytes          int64                    `json:"rxBytes"`
	RxBytesPerSecond float64                  `json:"rxBytesPerSecond"`
	Scope            RuntimeNetworkUsageScope `json:"scope"`
	TxBytes          int64                    `json:"txBytes"`
	TxBytesPerSecond float64                  `json:"txBytesPerSecond"`
}

// RuntimeNetworkUsageScope defines model for RuntimeNetworkUsage.Scope.
type RuntimeNetworkUsageScope string

// RuntimeResourceSnapshot defines model for RuntimeResourceSnapshot.
type RuntimeResourceSnapshot struct {
	CPU           RuntimeCPUUsage     `json:"cpu"`
	Disk          RuntimeDiskUsage    `json:"disk"`
	GeneratedAt   time.Time           `json:"generatedAt"`
	GoRuntime     RuntimeGoUsage      `json:"goRuntime"`
	Memory        RuntimeMemoryUsage  `json:"memory"`
	Network       RuntimeNetworkUsage `json:"network"`
	Services      RuntimeServiceUsage `json:"services"`
	UptimeSeconds int64               `json:"uptimeSeconds"`
}

// RuntimeResourceSnapshotEnvelope defines model for RuntimeResourceSnapshotEnvelope.
type RuntimeResourceSnapshotEnvelope struct {
	Data RuntimeResourceSnapshot `json:"data"`
}

// RuntimeServiceUsage defines model for RuntimeServiceUsage.
type RuntimeServiceUsage struct {
	Canceled   int64 `json:"canceled"`
	Failed     int64 `json:"failed"`
	QueueDepth int   `json:"queueDepth"`
	Started    int64 `json:"started"`
	Succeeded  int64 `json:"succeeded"`
}

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

// SourceBranch defines model for SourceBranch.
type SourceBranch struct {
	CommitID      string `json:"commitId"`
	DefaultBranch bool   `json:"defaultBranch"`
	Name          string `json:"name"`
}

// SourceBranchListEnvelope defines model for SourceBranchListEnvelope.
type SourceBranchListEnvelope struct {
	Items []SourceBranch `json:"items"`
}

// SourceConnection defines model for SourceConnection.
type SourceConnection struct {
	Capabilities      []SourceConnectionCapabilities `json:"capabilities"`
	DefaultConnection bool                           `json:"defaultConnection"`

	// DisplayEndpoint Redacted human-readable provider endpoint when one is configured.
	DisplayEndpoint string `json:"displayEndpoint,omitempty"`
	Enabled         bool   `json:"enabled"`
	ID              string `json:"id"`
	IntegrationID   string `json:"integrationId"`
	Name            string `json:"name"`
	ProviderType    string `json:"providerType"`
}

// SourceConnectionCapabilities defines model for SourceConnection.Capabilities.
type SourceConnectionCapabilities string

// SourceConnectionEnvelope defines model for SourceConnectionEnvelope.
type SourceConnectionEnvelope struct {
	Data SourceConnection `json:"data"`
}

// SourceConnectionListEnvelope defines model for SourceConnectionListEnvelope.
type SourceConnectionListEnvelope struct {
	Items []SourceConnection `json:"items"`
}

// SourceFile defines model for SourceFile.
type SourceFile struct {
	BlobID string `json:"blobId,omitempty"`

	// Content UTF-8 text or base64 content according to encoding. This API does not expose checkout credentials.
	Content      string             `json:"content"`
	Encoding     SourceFileEncoding `json:"encoding"`
	Path         string             `json:"path"`
	Ref          string             `json:"ref"`
	RepositoryID string             `json:"repositoryId"`
	SizeBytes    int64              `json:"sizeBytes"`
}

// SourceFileEncoding defines model for SourceFileEncoding.
type SourceFileEncoding string

// SourceFileEnvelope defines model for SourceFileEnvelope.
type SourceFileEnvelope struct {
	Data SourceFile `json:"data"`
}

// SourceRepository defines model for SourceRepository.
type SourceRepository struct {
	Archived      bool   `json:"archived"`
	DefaultBranch string `json:"defaultBranch"`
	FullName      string `json:"fullName"`

	// ID Opaque identifier scoped to the source connection.
	ID        string `json:"id"`
	Name      string `json:"name"`
	Namespace string `json:"namespace,omitempty"`
	WebURL    string `json:"webUrl,omitempty"`
}

// SourceRepositoryListEnvelope defines model for SourceRepositoryListEnvelope.
type SourceRepositoryListEnvelope struct {
	Items      []SourceRepository `json:"items"`
	NextCursor string             `json:"nextCursor,omitempty"`
}

// SourceTag defines model for SourceTag.
type SourceTag struct {
	CommitID string `json:"commitId"`
	Name     string `json:"name"`
}

// SourceTagListEnvelope defines model for SourceTagListEnvelope.
type SourceTagListEnvelope struct {
	Items []SourceTag `json:"items"`
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

// SystemIntegration defines model for SystemIntegration.
type SystemIntegration struct {
	Category      SystemIntegrationCategory             `json:"category"`
	Configuration []SystemIntegrationConfigurationField `json:"configuration"`
	CreatedAt     time.Time                             `json:"createdAt"`

	// CredentialKeys Names of configured credentials. Credential values are never returned.
	CredentialKeys []string                      `json:"credentialKeys"`
	Description    string                        `json:"description,omitempty"`
	Enabled        bool                          `json:"enabled"`
	HealthStatus   SystemIntegrationHealthStatus `json:"healthStatus"`
	ID             string                        `json:"id"`
	LastCheckedAt  *time.Time                    `json:"lastCheckedAt,omitempty"`
	Name           string                        `json:"name"`
	ProviderType   string                        `json:"providerType"`
	UpdatedAt      time.Time                     `json:"updatedAt"`
	Version        int64                         `json:"version"`
}

// SystemIntegrationCategory defines model for SystemIntegrationCategory.
type SystemIntegrationCategory string

// SystemIntegrationConfigurationField defines model for SystemIntegrationConfigurationField.
type SystemIntegrationConfigurationField struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// SystemIntegrationCreateRequest defines model for SystemIntegrationCreateRequest.
type SystemIntegrationCreateRequest struct {
	Category      SystemIntegrationCategory             `json:"category"`
	Configuration []SystemIntegrationConfigurationField `json:"configuration,omitempty"`
	Credentials   []SystemIntegrationCredentialInput    `json:"credentials,omitempty"`
	Description   string                                `json:"description,omitempty"`
	Enabled       bool                                  `json:"enabled"`
	Name          string                                `json:"name"`

	// ProviderType Extensible provider key such as gitlab, github, gitea, or zadig.
	ProviderType string `json:"providerType"`
}

// SystemIntegrationCredentialInput defines model for SystemIntegrationCredentialInput.
type SystemIntegrationCredentialInput struct {
	Key   string `json:"key"`
	Value string `json:"value,omitempty"`
}

// SystemIntegrationEnvelope defines model for SystemIntegrationEnvelope.
type SystemIntegrationEnvelope struct {
	Data SystemIntegration `json:"data"`
}

// SystemIntegrationHealthStatus defines model for SystemIntegrationHealthStatus.
type SystemIntegrationHealthStatus string

// SystemIntegrationListEnvelope defines model for SystemIntegrationListEnvelope.
type SystemIntegrationListEnvelope struct {
	Items []SystemIntegration `json:"items"`
}

// SystemIntegrationTestResult defines model for SystemIntegrationTestResult.
type SystemIntegrationTestResult struct {
	Capabilities  []string                    `json:"capabilities"`
	CheckedAt     time.Time                   `json:"checkedAt"`
	IntegrationID string                      `json:"integrationId"`
	LatencyMs     int64                       `json:"latencyMs"`
	Message       string                      `json:"message,omitempty"`
	Status        SystemIntegrationTestStatus `json:"status"`
}

// SystemIntegrationTestResultEnvelope defines model for SystemIntegrationTestResultEnvelope.
type SystemIntegrationTestResultEnvelope struct {
	Data SystemIntegrationTestResult `json:"data"`
}

// SystemIntegrationTestStatus defines model for SystemIntegrationTestStatus.
type SystemIntegrationTestStatus string

// SystemIntegrationUpdateRequest defines model for SystemIntegrationUpdateRequest.
type SystemIntegrationUpdateRequest struct {
	ClearCredentialKeys []string                               `json:"clearCredentialKeys,omitempty"`
	Configuration       *[]SystemIntegrationConfigurationField `json:"configuration,omitempty"`

	// Credentials Credential values to add or replace. Omitted credential keys remain unchanged.
	Credentials     []SystemIntegrationCredentialInput `json:"credentials,omitempty"`
	Description     *string                            `json:"description,omitempty"`
	Enabled         *bool                              `json:"enabled,omitempty"`
	ExpectedVersion int64                              `json:"expectedVersion"`
	Name            string                             `json:"name,omitempty"`
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

// UpdateAISkillsRequest defines model for UpdateAISkillsRequest.
type UpdateAISkillsRequest struct {
	SkillsRegistry []AISkillSettings `json:"skillsRegistry"`
}

// UpdateAIWorkbenchModelRequest defines model for UpdateAIWorkbenchModelRequest.
type UpdateAIWorkbenchModelRequest struct {
	WorkbenchModel AIWorkbenchModelSettings `json:"workbenchModel"`
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

// WorkbenchAgentStatusEvent defines model for WorkbenchAgentStatusEvent.
type WorkbenchAgentStatusEvent struct {
	CreatedAt    time.Time                             `json:"createdAt"`
	ID           string                                `json:"id"`
	MessageID    string                                `json:"messageId,omitempty"`
	ProviderID   string                                `json:"providerId"`
	ProviderKind WorkbenchAgentStatusEventProviderKind `json:"providerKind"`
	RunID        string                                `json:"runId,omitempty"`
	Sequence     int                                   `json:"sequence"`
	SessionID    string                                `json:"sessionId"`
	Status       WorkbenchAgentStatusEventStatus       `json:"status"`
	Type         WorkbenchAgentStatusEventType         `json:"type"`
}

// WorkbenchAgentStatusEventProviderKind defines model for WorkbenchAgentStatusEvent.ProviderKind.
type WorkbenchAgentStatusEventProviderKind string

// WorkbenchAgentStatusEventStatus defines model for WorkbenchAgentStatusEvent.Status.
type WorkbenchAgentStatusEventStatus string

// WorkbenchAgentStatusEventType defines model for WorkbenchAgentStatusEvent.Type.
type WorkbenchAgentStatusEventType string

// WorkbenchArtifactUpdatedEvent defines model for WorkbenchArtifactUpdatedEvent.
type WorkbenchArtifactUpdatedEvent struct {
	Artifact  AnyValue                          `json:"artifact"`
	CreatedAt time.Time                         `json:"createdAt"`
	ID        string                            `json:"id"`
	MessageID string                            `json:"messageId,omitempty"`
	RunID     string                            `json:"runId,omitempty"`
	Sequence  int                               `json:"sequence"`
	SessionID string                            `json:"sessionId"`
	Type      WorkbenchArtifactUpdatedEventType `json:"type"`
}

// WorkbenchArtifactUpdatedEventType defines model for WorkbenchArtifactUpdatedEvent.Type.
type WorkbenchArtifactUpdatedEventType string

// WorkbenchCardCommandEvent defines model for WorkbenchCardCommandEvent.
type WorkbenchCardCommandEvent struct {
	Command   AnyValue                      `json:"command"`
	CreatedAt time.Time                     `json:"createdAt"`
	ID        string                        `json:"id"`
	MessageID string                        `json:"messageId,omitempty"`
	RunID     string                        `json:"runId,omitempty"`
	Sequence  int                           `json:"sequence"`
	SessionID string                        `json:"sessionId"`
	SurfaceID string                        `json:"surfaceId"`
	Type      WorkbenchCardCommandEventType `json:"type"`
}

// WorkbenchCardCommandEventType defines model for WorkbenchCardCommandEvent.Type.
type WorkbenchCardCommandEventType string

// WorkbenchErrorEvent defines model for WorkbenchErrorEvent.
type WorkbenchErrorEvent struct {
	Code      string                  `json:"code,omitempty"`
	CreatedAt time.Time               `json:"createdAt"`
	ID        string                  `json:"id"`
	Message   string                  `json:"message"`
	MessageID string                  `json:"messageId,omitempty"`
	Retryable bool                    `json:"retryable,omitempty"`
	RunID     string                  `json:"runId,omitempty"`
	Sequence  int                     `json:"sequence"`
	SessionID string                  `json:"sessionId"`
	Type      WorkbenchErrorEventType `json:"type"`
}

// WorkbenchErrorEventType defines model for WorkbenchErrorEvent.Type.
type WorkbenchErrorEventType string

// WorkbenchGlobalAssistantEventEnvelope defines model for WorkbenchGlobalAssistantEventEnvelope.
type WorkbenchGlobalAssistantEventEnvelope struct {
	Ok bool `json:"ok"`
}

// WorkbenchGlobalAssistantOpenRequest defines model for WorkbenchGlobalAssistantOpenRequest.
type WorkbenchGlobalAssistantOpenRequest struct {
	Action           WorkbenchGlobalAssistantOpenRequestAction `json:"action"`
	LaunchContext    WorkbenchLaunchContext                    `json:"launchContext"`
	Prompt           string                                    `json:"prompt,omitempty"`
	SelectionContext *WorkbenchSelectionContext                `json:"selectionContext,omitempty"`
	SessionID        string                                    `json:"sessionId,omitempty"`
	Source           string                                    `json:"source,omitempty"`
}

// WorkbenchGlobalAssistantOpenRequestAction defines model for WorkbenchGlobalAssistantOpenRequest.Action.
type WorkbenchGlobalAssistantOpenRequestAction string

// WorkbenchLaunchContext defines model for WorkbenchLaunchContext.
type WorkbenchLaunchContext struct {
	AlertID                    string                                `json:"alertId,omitempty"`
	ApplicationID              string                                `json:"applicationId,omitempty"`
	ClusterID                  string                                `json:"clusterId,omitempty"`
	DockerHostID               string                                `json:"dockerHostId,omitempty"`
	DockerServiceID            string                                `json:"dockerServiceId,omitempty"`
	EntityKind                 string                                `json:"entityKind,omitempty"`
	EntityName                 string                                `json:"entityName,omitempty"`
	Namespace                  string                                `json:"namespace,omitempty"`
	Node                       string                                `json:"node,omitempty"`
	PinnedData                 map[string]any                        `json:"pinnedData,omitempty"`
	Pod                        string                                `json:"pod,omitempty"`
	ReleaseBundleID            string                                `json:"releaseBundleId,omitempty"`
	Service                    string                                `json:"service,omitempty"`
	SourceRoute                string                                `json:"sourceRoute"`
	SourceTitle                string                                `json:"sourceTitle,omitempty"`
	SourceWorkbench            WorkbenchLaunchContextSourceWorkbench `json:"sourceWorkbench"`
	TimeRangeMinutes           int                                   `json:"timeRangeMinutes,omitempty"`
	VirtualizationConnectionID string                                `json:"virtualizationConnectionId,omitempty"`
	VisibleFilters             map[string]any                        `json:"visibleFilters,omitempty"`
	VMID                       string                                `json:"vmId,omitempty"`
	Workload                   string                                `json:"workload,omitempty"`
}

// WorkbenchLaunchContextSourceWorkbench defines model for WorkbenchLaunchContext.SourceWorkbench.
type WorkbenchLaunchContextSourceWorkbench string

// WorkbenchMessageDeltaEvent defines model for WorkbenchMessageDeltaEvent.
type WorkbenchMessageDeltaEvent struct {
	ContentDelta string                         `json:"contentDelta"`
	CreatedAt    time.Time                      `json:"createdAt"`
	ID           string                         `json:"id"`
	MessageID    string                         `json:"messageId,omitempty"`
	Role         WorkbenchMessageDeltaEventRole `json:"role"`
	RunID        string                         `json:"runId,omitempty"`
	Sequence     int                            `json:"sequence"`
	SessionID    string                         `json:"sessionId"`
	Type         WorkbenchMessageDeltaEventType `json:"type"`
}

// WorkbenchMessageDeltaEventRole defines model for WorkbenchMessageDeltaEvent.Role.
type WorkbenchMessageDeltaEventRole string

// WorkbenchMessageDeltaEventType defines model for WorkbenchMessageDeltaEvent.Type.
type WorkbenchMessageDeltaEventType string

// WorkbenchMessageDoneEvent defines model for WorkbenchMessageDoneEvent.
type WorkbenchMessageDoneEvent struct {
	Content   string                        `json:"content"`
	CreatedAt time.Time                     `json:"createdAt"`
	ID        string                        `json:"id"`
	MessageID string                        `json:"messageId,omitempty"`
	Metadata  map[string]any                `json:"metadata,omitempty"`
	Role      WorkbenchMessageDoneEventRole `json:"role"`
	RunID     string                        `json:"runId,omitempty"`
	Sequence  int                           `json:"sequence"`
	SessionID string                        `json:"sessionId"`
	Type      WorkbenchMessageDoneEventType `json:"type"`
}

// WorkbenchMessageDoneEventRole defines model for WorkbenchMessageDoneEvent.Role.
type WorkbenchMessageDoneEventRole string

// WorkbenchMessageDoneEventType defines model for WorkbenchMessageDoneEvent.Type.
type WorkbenchMessageDoneEventType string

// WorkbenchSelectionContext defines model for WorkbenchSelectionContext.
type WorkbenchSelectionContext struct {
	Kind               WorkbenchSelectionContextKind `json:"kind"`
	SourceElementLabel string                        `json:"sourceElementLabel,omitempty"`
	Text               string                        `json:"text"`
}

// WorkbenchSelectionContextKind defines model for WorkbenchSelectionContext.Kind.
type WorkbenchSelectionContextKind string

// WorkbenchSendMessageStreamRequest defines model for WorkbenchSendMessageStreamRequest.
type WorkbenchSendMessageStreamRequest struct {
	AgentProviderID  string                     `json:"agentProviderId,omitempty"`
	Content          string                     `json:"content"`
	LaunchContext    *WorkbenchLaunchContext    `json:"launchContext,omitempty"`
	Mode             string                     `json:"mode,omitempty"`
	PinnedContext    map[string]any             `json:"pinnedContext,omitempty"`
	ScopeOverrides   AnyValue                   `json:"scopeOverrides,omitempty"`
	SelectionContext *WorkbenchSelectionContext `json:"selectionContext,omitempty"`
	Source           string                     `json:"source,omitempty"`
	Toolset          AnyValue                   `json:"toolset,omitempty"`
}

// WorkbenchSource defines model for WorkbenchSource.
type WorkbenchSource struct {
	ID      string              `json:"id"`
	Kind    WorkbenchSourceKind `json:"kind"`
	Summary string              `json:"summary,omitempty"`
	Title   string              `json:"title"`
	URL     string              `json:"url,omitempty"`
}

// WorkbenchSourceKind defines model for WorkbenchSource.Kind.
type WorkbenchSourceKind string

// WorkbenchSourceUpdatedEvent defines model for WorkbenchSourceUpdatedEvent.
type WorkbenchSourceUpdatedEvent struct {
	CreatedAt time.Time                       `json:"createdAt"`
	ID        string                          `json:"id"`
	MessageID string                          `json:"messageId,omitempty"`
	RunID     string                          `json:"runId,omitempty"`
	Sequence  int                             `json:"sequence"`
	SessionID string                          `json:"sessionId"`
	Source    WorkbenchSource                 `json:"source"`
	Type      WorkbenchSourceUpdatedEventType `json:"type"`
}

// WorkbenchSourceUpdatedEventType defines model for WorkbenchSourceUpdatedEvent.Type.
type WorkbenchSourceUpdatedEventType string

// WorkbenchStreamEvent defines model for WorkbenchStreamEvent.
type WorkbenchStreamEvent struct {
	union json.RawMessage
}

// WorkbenchStreamEventBase defines model for WorkbenchStreamEventBase.
type WorkbenchStreamEventBase struct {
	CreatedAt time.Time `json:"createdAt"`
	ID        string    `json:"id"`
	MessageID string    `json:"messageId,omitempty"`
	RunID     string    `json:"runId,omitempty"`
	Sequence  int       `json:"sequence"`
	SessionID string    `json:"sessionId"`
}

// WorkbenchThinkingDeltaEvent defines model for WorkbenchThinkingDeltaEvent.
type WorkbenchThinkingDeltaEvent struct {
	CreatedAt time.Time                       `json:"createdAt"`
	ID        string                          `json:"id"`
	MessageID string                          `json:"messageId,omitempty"`
	RunID     string                          `json:"runId,omitempty"`
	Sequence  int                             `json:"sequence"`
	SessionID string                          `json:"sessionId"`
	TextDelta string                          `json:"textDelta"`
	Type      WorkbenchThinkingDeltaEventType `json:"type"`
}

// WorkbenchThinkingDeltaEventType defines model for WorkbenchThinkingDeltaEvent.Type.
type WorkbenchThinkingDeltaEventType string

// WorkbenchThinkingDoneEvent defines model for WorkbenchThinkingDoneEvent.
type WorkbenchThinkingDoneEvent struct {
	Collapsed bool                           `json:"collapsed"`
	CreatedAt time.Time                      `json:"createdAt"`
	ID        string                         `json:"id"`
	MessageID string                         `json:"messageId,omitempty"`
	RunID     string                         `json:"runId,omitempty"`
	Sequence  int                            `json:"sequence"`
	SessionID string                         `json:"sessionId"`
	Summary   string                         `json:"summary"`
	Type      WorkbenchThinkingDoneEventType `json:"type"`
}

// WorkbenchThinkingDoneEventType defines model for WorkbenchThinkingDoneEvent.Type.
type WorkbenchThinkingDoneEventType string

// WorkbenchToolCall defines model for WorkbenchToolCall.
type WorkbenchToolCall struct {
	AdapterID     string                  `json:"adapterId"`
	ArtifactRefs  []string                `json:"artifactRefs,omitempty"`
	CapabilityID  string                  `json:"capabilityId,omitempty"`
	CompletedAt   *time.Time              `json:"completedAt,omitempty"`
	DurationMs    int                     `json:"durationMs,omitempty"`
	EvidenceRefs  []string                `json:"evidenceRefs,omitempty"`
	ID            string                  `json:"id"`
	InputPreview  AnyValue                `json:"inputPreview,omitempty"`
	OutputPreview AnyValue                `json:"outputPreview,omitempty"`
	SkillID       string                  `json:"skillId,omitempty"`
	SkillName     string                  `json:"skillName,omitempty"`
	StartedAt     *time.Time              `json:"startedAt,omitempty"`
	Status        WorkbenchToolCallStatus `json:"status"`
	Summary       string                  `json:"summary,omitempty"`
	ToolName      string                  `json:"toolName"`
}

// WorkbenchToolCallStatus defines model for WorkbenchToolCall.Status.
type WorkbenchToolCallStatus string

// WorkbenchToolCompletedEvent defines model for WorkbenchToolCompletedEvent.
type WorkbenchToolCompletedEvent struct {
	CreatedAt time.Time                       `json:"createdAt"`
	ID        string                          `json:"id"`
	MessageID string                          `json:"messageId,omitempty"`
	RunID     string                          `json:"runId,omitempty"`
	Sequence  int                             `json:"sequence"`
	SessionID string                          `json:"sessionId"`
	ToolCall  WorkbenchToolCall               `json:"toolCall"`
	Type      WorkbenchToolCompletedEventType `json:"type"`
}

// WorkbenchToolCompletedEventType defines model for WorkbenchToolCompletedEvent.Type.
type WorkbenchToolCompletedEventType string

// WorkbenchToolDeltaEvent defines model for WorkbenchToolDeltaEvent.
type WorkbenchToolDeltaEvent struct {
	CreatedAt   time.Time                   `json:"createdAt"`
	ID          string                      `json:"id"`
	LogDelta    string                      `json:"logDelta,omitempty"`
	MessageID   string                      `json:"messageId,omitempty"`
	OutputDelta string                      `json:"outputDelta,omitempty"`
	RunID       string                      `json:"runId,omitempty"`
	Sequence    int                         `json:"sequence"`
	SessionID   string                      `json:"sessionId"`
	ToolCallID  string                      `json:"toolCallId"`
	Type        WorkbenchToolDeltaEventType `json:"type"`
}

// WorkbenchToolDeltaEventType defines model for WorkbenchToolDeltaEvent.Type.
type WorkbenchToolDeltaEventType string

// WorkbenchToolStartedEvent defines model for WorkbenchToolStartedEvent.
type WorkbenchToolStartedEvent struct {
	CreatedAt time.Time                     `json:"createdAt"`
	ID        string                        `json:"id"`
	MessageID string                        `json:"messageId,omitempty"`
	RunID     string                        `json:"runId,omitempty"`
	Sequence  int                           `json:"sequence"`
	SessionID string                        `json:"sessionId"`
	ToolCall  WorkbenchToolCall             `json:"toolCall"`
	Type      WorkbenchToolStartedEventType `json:"type"`
}

// WorkbenchToolStartedEventType defines model for WorkbenchToolStartedEvent.Type.
type WorkbenchToolStartedEventType string

// WorkflowNodeRun defines model for WorkflowNodeRun.
type WorkflowNodeRun struct {
	FinishedAt string `json:"finishedAt,omitempty"`
	Name       string `json:"name"`
	NodeID     string `json:"nodeId"`
	StartedAt  string `json:"startedAt,omitempty"`
	Status     string `json:"status"`
	Summary    string `json:"summary,omitempty"`
	Type       string `json:"type"`
}

// WorkflowRun defines model for WorkflowRun.
type WorkflowRun struct {
	ApplicationID  string            `json:"applicationId"`
	ClusterID      string            `json:"clusterId,omitempty"`
	CreatedAt      string            `json:"createdAt"`
	DeploymentName string            `json:"deploymentName,omitempty"`
	ID             string            `json:"id"`
	Metadata       map[string]any    `json:"metadata,omitempty"`
	Namespace      string            `json:"namespace,omitempty"`
	NodeRuns       []WorkflowNodeRun `json:"nodeRuns,omitempty"`
	Status         string            `json:"status"`
	Steps          []WorkflowStep    `json:"steps"`
	UpdatedAt      string            `json:"updatedAt"`
	WorkflowName   string            `json:"workflowName"`
}

// WorkflowStep defines model for WorkflowStep.
type WorkflowStep struct {
	Name    string `json:"name"`
	Status  string `json:"status"`
	Summary string `json:"summary,omitempty"`
}

// WorkflowTemplate defines model for WorkflowTemplate.
type WorkflowTemplate struct {
	Category    string         `json:"category,omitempty"`
	CreatedAt   time.Time      `json:"createdAt"`
	Definition  map[string]any `json:"definition,omitempty"`
	Description string         `json:"description,omitempty"`
	Enabled     bool           `json:"enabled"`
	ID          string         `json:"id"`
	Key         string         `json:"key"`
	Name        string         `json:"name"`
	UpdatedAt   time.Time      `json:"updatedAt"`
}

// WorkflowTemplateEnvelope defines model for WorkflowTemplateEnvelope.
type WorkflowTemplateEnvelope struct {
	Data WorkflowTemplate `json:"data"`
}

// WorkflowTemplateInput defines model for WorkflowTemplateInput.
type WorkflowTemplateInput struct {
	Category    string         `json:"category,omitempty"`
	Definition  map[string]any `json:"definition,omitempty"`
	Description string         `json:"description,omitempty"`
	Enabled     bool           `json:"enabled"`
	ID          string         `json:"id,omitempty"`
	Key         string         `json:"key"`
	Name        string         `json:"name"`
}

// WorkflowTemplateListEnvelope defines model for WorkflowTemplateListEnvelope.
type WorkflowTemplateListEnvelope struct {
	Data []WorkflowTemplate `json:"data"`
}

// AIClientID defines model for AIClientID.
type AIClientID = string

// AIClientName defines model for AIClientName.
type AIClientName = string

// AIEnvironmentLeaseID defines model for AIEnvironmentLeaseID.
type AIEnvironmentLeaseID = string

// AIMemoryID defines model for AIMemoryID.
type AIMemoryID = string

// AgentProviderRolloutAction defines model for AgentProviderRolloutAction.
type AgentProviderRolloutAction string

// AgentProviderRolloutID defines model for AgentProviderRolloutID.
type AgentProviderRolloutID = string

// ApplicationEnvironmentID defines model for ApplicationEnvironmentID.
type ApplicationEnvironmentID = string

// ApplicationID defines model for ApplicationID.
type ApplicationID = string

// BuildTemplateID defines model for BuildTemplateID.
type BuildTemplateID = string

// BundleID defines model for BundleID.
type BundleID = string

// ClusterID defines model for ClusterID.
type ClusterID = string

// ComputeActionID defines model for ComputeActionID.
type ComputeActionID = string

// ComputeCursor defines model for ComputeCursor.
type ComputeCursor = string

// ComputeLimit defines model for ComputeLimit.
type ComputeLimit = int

// ComputeProviderInstanceRef defines model for ComputeProviderInstanceRef.
type ComputeProviderInstanceRef = string

// ComputeProviderKey defines model for ComputeProviderKey.
type ComputeProviderKey = string

// ComputeResourceID defines model for ComputeResourceID.
type ComputeResourceID = string

// ComputeTaskID defines model for ComputeTaskID.
type ComputeTaskID = string

// EvaluationRunID defines model for EvaluationRunID.
type EvaluationRunID = string

// GitLabProjectIDQuery defines model for GitLabProjectIDQuery.
type GitLabProjectIDQuery = string

// IdempotencyKey defines model for IdempotencyKey.
type IdempotencyKey = string

// KnowledgeBaseID defines model for KnowledgeBaseID.
type KnowledgeBaseID = string

// KnowledgeConnectorID defines model for KnowledgeConnectorID.
type KnowledgeConnectorID = string

// KnowledgeGraphRevisionID defines model for KnowledgeGraphRevisionID.
type KnowledgeGraphRevisionID = string

// KnowledgeIngestionJobID defines model for KnowledgeIngestionJobID.
type KnowledgeIngestionJobID = string

// KnowledgeListLimit defines model for KnowledgeListLimit.
type KnowledgeListLimit = int

// KnowledgeSourceID defines model for KnowledgeSourceID.
type KnowledgeSourceID = string

// MultiAgentPlanID defines model for MultiAgentPlanID.
type MultiAgentPlanID = string

// MultiAgentSubtaskID defines model for MultiAgentSubtaskID.
type MultiAgentSubtaskID = string

// OIDCCode defines model for OIDCCode.
type OIDCCode = string

// OIDCState defines model for OIDCState.
type OIDCState = string

// OpenAICompatibleProvider defines model for OpenAICompatibleProvider.
type OpenAICompatibleProvider string

// OperationID defines model for OperationID.
type OperationID = string

// PlanID defines model for PlanID.
type PlanID = string

// PluginID defines model for PluginID.
type PluginID = string

// ProviderID defines model for ProviderID.
type ProviderID = string

// RepositoryID defines model for RepositoryID.
type RepositoryID = string

// RouteID defines model for RouteID.
type RouteID = string

// RuntimeConfigApplicationID defines model for RuntimeConfigApplicationID.
type RuntimeConfigApplicationID = string

// ServiceID defines model for ServiceID.
type ServiceID = string

// SessionID defines model for SessionID.
type SessionID = string

// SkillID defines model for SkillID.
type SkillID = string

// SohaCacheModeHeader defines model for SohaCacheModeHeader.
type SohaCacheModeHeader string

// SohaRouteTraceHeader defines model for SohaRouteTraceHeader.
type SohaRouteTraceHeader = bool

// SohaUpstreamIDHeader defines model for SohaUpstreamIDHeader.
type SohaUpstreamIDHeader = string

// Source defines model for Source.
type Source = string

// SourceConnectionID defines model for SourceConnectionID.
type SourceConnectionID = string

// SourceRepositoryID defines model for SourceRepositoryID.
type SourceRepositoryID = string

// SystemIntegrationID defines model for SystemIntegrationID.
type SystemIntegrationID = string

// TaskID defines model for TaskID.
type TaskID = string

// ToolName defines model for ToolName.
type ToolName = string

// UpstreamID defines model for UpstreamID.
type UpstreamID = string

// WorkflowTemplateID defines model for WorkflowTemplateID.
type WorkflowTemplateID = string

// ComputeError defines model for ComputeError.
type ComputeError = ErrorEnvelope

// Error defines model for Error.
type Error = ErrorEnvelope

// bearerAuthContextKey is the context key for bearerAuth security scheme
type bearerAuthContextKey string

// runtimeBearerAuthContextKey is the context key for runtimeBearerAuth security scheme
type runtimeBearerAuthContextKey string

// sohaXAPIKeyContextKey is the context key for sohaXAPIKey security scheme
type sohaXAPIKeyContextKey string

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

// CreateAIGatewayAnthropicMessageParams defines parameters for CreateAIGatewayAnthropicMessage.
type CreateAIGatewayAnthropicMessageParams struct {
	// AnthropicVersion Forwarded Anthropic API version header when present.
	AnthropicVersion string `json:"anthropic-version,omitempty"`

	// XSohaUpstreamID Optional debug override for administrators or tokens explicitly allowed to select an upstream.
	XSohaUpstreamID SohaUpstreamIDHeader `json:"X-Soha-Upstream-ID,omitempty"`

	// XSohaRouteTrace When true, returns redacted Soha route trace response headers for relay managers or tokens with allowRouteTrace explicitly enabled.
	XSohaRouteTrace SohaRouteTraceHeader `json:"X-Soha-Route-Trace,omitempty"`

	// XSohaCacheMode Optional relay cache behavior hint.
	XSohaCacheMode CreateAIGatewayAnthropicMessageParamsXSohaCacheMode `json:"X-Soha-Cache-Mode,omitempty"`
}

// CreateAIGatewayAnthropicMessageParamsXSohaCacheMode defines parameters for CreateAIGatewayAnthropicMessage.
type CreateAIGatewayAnthropicMessageParamsXSohaCacheMode string

// CreateAIGatewayCohereRerankParams defines parameters for CreateAIGatewayCohereRerank.
type CreateAIGatewayCohereRerankParams struct {
	// XSohaUpstreamID Optional debug override for administrators or tokens explicitly allowed to select an upstream.
	XSohaUpstreamID SohaUpstreamIDHeader `json:"X-Soha-Upstream-ID,omitempty"`

	// XSohaRouteTrace When true, returns redacted Soha route trace response headers for relay managers or tokens with allowRouteTrace explicitly enabled.
	XSohaRouteTrace SohaRouteTraceHeader `json:"X-Soha-Route-Trace,omitempty"`

	// XSohaCacheMode Optional relay cache behavior hint.
	XSohaCacheMode CreateAIGatewayCohereRerankParamsXSohaCacheMode `json:"X-Soha-Cache-Mode,omitempty"`
}

// CreateAIGatewayCohereRerankParamsXSohaCacheMode defines parameters for CreateAIGatewayCohereRerank.
type CreateAIGatewayCohereRerankParamsXSohaCacheMode string

// CreateAIGatewayGeminiInteractionsParams defines parameters for CreateAIGatewayGeminiInteractions.
type CreateAIGatewayGeminiInteractionsParams struct {
	// XSohaUpstreamID Optional debug override for administrators or tokens explicitly allowed to select an upstream.
	XSohaUpstreamID SohaUpstreamIDHeader `json:"X-Soha-Upstream-ID,omitempty"`

	// XSohaRouteTrace When true, returns redacted Soha route trace response headers for relay managers or tokens with allowRouteTrace explicitly enabled.
	XSohaRouteTrace SohaRouteTraceHeader `json:"X-Soha-Route-Trace,omitempty"`

	// XSohaCacheMode Optional relay cache behavior hint.
	XSohaCacheMode CreateAIGatewayGeminiInteractionsParamsXSohaCacheMode `json:"X-Soha-Cache-Mode,omitempty"`
}

// CreateAIGatewayGeminiInteractionsParamsXSohaCacheMode defines parameters for CreateAIGatewayGeminiInteractions.
type CreateAIGatewayGeminiInteractionsParamsXSohaCacheMode string

// CreateAIGatewayGeminiGenerateContentParams defines parameters for CreateAIGatewayGeminiGenerateContent.
type CreateAIGatewayGeminiGenerateContentParams struct {
	// XSohaUpstreamID Optional debug override for administrators or tokens explicitly allowed to select an upstream.
	XSohaUpstreamID SohaUpstreamIDHeader `json:"X-Soha-Upstream-ID,omitempty"`

	// XSohaRouteTrace When true, returns redacted Soha route trace response headers for relay managers or tokens with allowRouteTrace explicitly enabled.
	XSohaRouteTrace SohaRouteTraceHeader `json:"X-Soha-Route-Trace,omitempty"`

	// XSohaCacheMode Optional relay cache behavior hint.
	XSohaCacheMode CreateAIGatewayGeminiGenerateContentParamsXSohaCacheMode `json:"X-Soha-Cache-Mode,omitempty"`
}

// CreateAIGatewayGeminiGenerateContentParamsXSohaCacheMode defines parameters for CreateAIGatewayGeminiGenerateContent.
type CreateAIGatewayGeminiGenerateContentParamsXSohaCacheMode string

// CreateAIGatewayGeminiStreamGenerateContentParams defines parameters for CreateAIGatewayGeminiStreamGenerateContent.
type CreateAIGatewayGeminiStreamGenerateContentParams struct {
	// XSohaUpstreamID Optional debug override for administrators or tokens explicitly allowed to select an upstream.
	XSohaUpstreamID SohaUpstreamIDHeader `json:"X-Soha-Upstream-ID,omitempty"`

	// XSohaRouteTrace When true, returns redacted Soha route trace response headers for relay managers or tokens with allowRouteTrace explicitly enabled.
	XSohaRouteTrace SohaRouteTraceHeader `json:"X-Soha-Route-Trace,omitempty"`

	// XSohaCacheMode Optional relay cache behavior hint.
	XSohaCacheMode CreateAIGatewayGeminiStreamGenerateContentParamsXSohaCacheMode `json:"X-Soha-Cache-Mode,omitempty"`
}

// CreateAIGatewayGeminiStreamGenerateContentParamsXSohaCacheMode defines parameters for CreateAIGatewayGeminiStreamGenerateContent.
type CreateAIGatewayGeminiStreamGenerateContentParamsXSohaCacheMode string

// CreateAIGatewayOpenAIAudioSpeechParams defines parameters for CreateAIGatewayOpenAIAudioSpeech.
type CreateAIGatewayOpenAIAudioSpeechParams struct {
	// XSohaUpstreamID Optional debug override for administrators or tokens explicitly allowed to select an upstream.
	XSohaUpstreamID SohaUpstreamIDHeader `json:"X-Soha-Upstream-ID,omitempty"`

	// XSohaRouteTrace When true, returns redacted Soha route trace response headers for relay managers or tokens with allowRouteTrace explicitly enabled.
	XSohaRouteTrace SohaRouteTraceHeader `json:"X-Soha-Route-Trace,omitempty"`

	// XSohaCacheMode Optional relay cache behavior hint.
	XSohaCacheMode CreateAIGatewayOpenAIAudioSpeechParamsXSohaCacheMode `json:"X-Soha-Cache-Mode,omitempty"`
}

// CreateAIGatewayOpenAIAudioSpeechParamsXSohaCacheMode defines parameters for CreateAIGatewayOpenAIAudioSpeech.
type CreateAIGatewayOpenAIAudioSpeechParamsXSohaCacheMode string

// CreateAIGatewayOpenAIAudioTranscriptionParams defines parameters for CreateAIGatewayOpenAIAudioTranscription.
type CreateAIGatewayOpenAIAudioTranscriptionParams struct {
	// XSohaUpstreamID Optional debug override for administrators or tokens explicitly allowed to select an upstream.
	XSohaUpstreamID SohaUpstreamIDHeader `json:"X-Soha-Upstream-ID,omitempty"`

	// XSohaRouteTrace When true, returns redacted Soha route trace response headers for relay managers or tokens with allowRouteTrace explicitly enabled.
	XSohaRouteTrace SohaRouteTraceHeader `json:"X-Soha-Route-Trace,omitempty"`

	// XSohaCacheMode Optional relay cache behavior hint.
	XSohaCacheMode CreateAIGatewayOpenAIAudioTranscriptionParamsXSohaCacheMode `json:"X-Soha-Cache-Mode,omitempty"`
}

// CreateAIGatewayOpenAIAudioTranscriptionParamsXSohaCacheMode defines parameters for CreateAIGatewayOpenAIAudioTranscription.
type CreateAIGatewayOpenAIAudioTranscriptionParamsXSohaCacheMode string

// CreateAIGatewayOpenAIAudioTranslationParams defines parameters for CreateAIGatewayOpenAIAudioTranslation.
type CreateAIGatewayOpenAIAudioTranslationParams struct {
	// XSohaUpstreamID Optional debug override for administrators or tokens explicitly allowed to select an upstream.
	XSohaUpstreamID SohaUpstreamIDHeader `json:"X-Soha-Upstream-ID,omitempty"`

	// XSohaRouteTrace When true, returns redacted Soha route trace response headers for relay managers or tokens with allowRouteTrace explicitly enabled.
	XSohaRouteTrace SohaRouteTraceHeader `json:"X-Soha-Route-Trace,omitempty"`

	// XSohaCacheMode Optional relay cache behavior hint.
	XSohaCacheMode CreateAIGatewayOpenAIAudioTranslationParamsXSohaCacheMode `json:"X-Soha-Cache-Mode,omitempty"`
}

// CreateAIGatewayOpenAIAudioTranslationParamsXSohaCacheMode defines parameters for CreateAIGatewayOpenAIAudioTranslation.
type CreateAIGatewayOpenAIAudioTranslationParamsXSohaCacheMode string

// CreateAIGatewayOpenAIChatCompletionParams defines parameters for CreateAIGatewayOpenAIChatCompletion.
type CreateAIGatewayOpenAIChatCompletionParams struct {
	// XSohaUpstreamID Optional debug override for administrators or tokens explicitly allowed to select an upstream.
	XSohaUpstreamID SohaUpstreamIDHeader `json:"X-Soha-Upstream-ID,omitempty"`

	// XSohaRouteTrace When true, returns redacted Soha route trace response headers for relay managers or tokens with allowRouteTrace explicitly enabled.
	XSohaRouteTrace SohaRouteTraceHeader `json:"X-Soha-Route-Trace,omitempty"`

	// XSohaCacheMode Optional relay cache behavior hint.
	XSohaCacheMode CreateAIGatewayOpenAIChatCompletionParamsXSohaCacheMode `json:"X-Soha-Cache-Mode,omitempty"`
}

// CreateAIGatewayOpenAIChatCompletionParamsXSohaCacheMode defines parameters for CreateAIGatewayOpenAIChatCompletion.
type CreateAIGatewayOpenAIChatCompletionParamsXSohaCacheMode string

// CreateAIGatewayOpenAIEmbeddingParams defines parameters for CreateAIGatewayOpenAIEmbedding.
type CreateAIGatewayOpenAIEmbeddingParams struct {
	// XSohaUpstreamID Optional debug override for administrators or tokens explicitly allowed to select an upstream.
	XSohaUpstreamID SohaUpstreamIDHeader `json:"X-Soha-Upstream-ID,omitempty"`

	// XSohaRouteTrace When true, returns redacted Soha route trace response headers for relay managers or tokens with allowRouteTrace explicitly enabled.
	XSohaRouteTrace SohaRouteTraceHeader `json:"X-Soha-Route-Trace,omitempty"`

	// XSohaCacheMode Optional relay cache behavior hint.
	XSohaCacheMode CreateAIGatewayOpenAIEmbeddingParamsXSohaCacheMode `json:"X-Soha-Cache-Mode,omitempty"`
}

// CreateAIGatewayOpenAIEmbeddingParamsXSohaCacheMode defines parameters for CreateAIGatewayOpenAIEmbedding.
type CreateAIGatewayOpenAIEmbeddingParamsXSohaCacheMode string

// CreateAIGatewayOpenAIImageEditParams defines parameters for CreateAIGatewayOpenAIImageEdit.
type CreateAIGatewayOpenAIImageEditParams struct {
	// XSohaUpstreamID Optional debug override for administrators or tokens explicitly allowed to select an upstream.
	XSohaUpstreamID SohaUpstreamIDHeader `json:"X-Soha-Upstream-ID,omitempty"`

	// XSohaRouteTrace When true, returns redacted Soha route trace response headers for relay managers or tokens with allowRouteTrace explicitly enabled.
	XSohaRouteTrace SohaRouteTraceHeader `json:"X-Soha-Route-Trace,omitempty"`

	// XSohaCacheMode Optional relay cache behavior hint.
	XSohaCacheMode CreateAIGatewayOpenAIImageEditParamsXSohaCacheMode `json:"X-Soha-Cache-Mode,omitempty"`
}

// CreateAIGatewayOpenAIImageEditParamsXSohaCacheMode defines parameters for CreateAIGatewayOpenAIImageEdit.
type CreateAIGatewayOpenAIImageEditParamsXSohaCacheMode string

// CreateAIGatewayOpenAIImageGenerationParams defines parameters for CreateAIGatewayOpenAIImageGeneration.
type CreateAIGatewayOpenAIImageGenerationParams struct {
	// XSohaUpstreamID Optional debug override for administrators or tokens explicitly allowed to select an upstream.
	XSohaUpstreamID SohaUpstreamIDHeader `json:"X-Soha-Upstream-ID,omitempty"`

	// XSohaRouteTrace When true, returns redacted Soha route trace response headers for relay managers or tokens with allowRouteTrace explicitly enabled.
	XSohaRouteTrace SohaRouteTraceHeader `json:"X-Soha-Route-Trace,omitempty"`

	// XSohaCacheMode Optional relay cache behavior hint.
	XSohaCacheMode CreateAIGatewayOpenAIImageGenerationParamsXSohaCacheMode `json:"X-Soha-Cache-Mode,omitempty"`
}

// CreateAIGatewayOpenAIImageGenerationParamsXSohaCacheMode defines parameters for CreateAIGatewayOpenAIImageGeneration.
type CreateAIGatewayOpenAIImageGenerationParamsXSohaCacheMode string

// CreateAIGatewayOpenAIImageVariationParams defines parameters for CreateAIGatewayOpenAIImageVariation.
type CreateAIGatewayOpenAIImageVariationParams struct {
	// XSohaUpstreamID Optional debug override for administrators or tokens explicitly allowed to select an upstream.
	XSohaUpstreamID SohaUpstreamIDHeader `json:"X-Soha-Upstream-ID,omitempty"`

	// XSohaRouteTrace When true, returns redacted Soha route trace response headers for relay managers or tokens with allowRouteTrace explicitly enabled.
	XSohaRouteTrace SohaRouteTraceHeader `json:"X-Soha-Route-Trace,omitempty"`

	// XSohaCacheMode Optional relay cache behavior hint.
	XSohaCacheMode CreateAIGatewayOpenAIImageVariationParamsXSohaCacheMode `json:"X-Soha-Cache-Mode,omitempty"`
}

// CreateAIGatewayOpenAIImageVariationParamsXSohaCacheMode defines parameters for CreateAIGatewayOpenAIImageVariation.
type CreateAIGatewayOpenAIImageVariationParamsXSohaCacheMode string

// ConnectAIGatewayOpenAIRealtimeParams defines parameters for ConnectAIGatewayOpenAIRealtime.
type ConnectAIGatewayOpenAIRealtimeParams struct {
	// Model Public model name to route to an upstream OpenAI Realtime model.
	Model string `form:"model" json:"model"`

	// XSohaUpstreamID Optional debug override for administrators or tokens explicitly allowed to select an upstream.
	XSohaUpstreamID SohaUpstreamIDHeader `json:"X-Soha-Upstream-ID,omitempty"`

	// XSohaRouteTrace When true, returns redacted Soha route trace response headers for relay managers or tokens with allowRouteTrace explicitly enabled.
	XSohaRouteTrace SohaRouteTraceHeader `json:"X-Soha-Route-Trace,omitempty"`
}

// CreateAIGatewayOpenAIResponseParams defines parameters for CreateAIGatewayOpenAIResponse.
type CreateAIGatewayOpenAIResponseParams struct {
	// XSohaUpstreamID Optional debug override for administrators or tokens explicitly allowed to select an upstream.
	XSohaUpstreamID SohaUpstreamIDHeader `json:"X-Soha-Upstream-ID,omitempty"`

	// XSohaRouteTrace When true, returns redacted Soha route trace response headers for relay managers or tokens with allowRouteTrace explicitly enabled.
	XSohaRouteTrace SohaRouteTraceHeader `json:"X-Soha-Route-Trace,omitempty"`

	// XSohaCacheMode Optional relay cache behavior hint.
	XSohaCacheMode CreateAIGatewayOpenAIResponseParamsXSohaCacheMode `json:"X-Soha-Cache-Mode,omitempty"`
}

// CreateAIGatewayOpenAIResponseParamsXSohaCacheMode defines parameters for CreateAIGatewayOpenAIResponse.
type CreateAIGatewayOpenAIResponseParamsXSohaCacheMode string

// CreateAIGatewayOpenAICompatibleProviderAudioSpeechParams defines parameters for CreateAIGatewayOpenAICompatibleProviderAudioSpeech.
type CreateAIGatewayOpenAICompatibleProviderAudioSpeechParams struct {
	// XSohaUpstreamID Optional debug override for administrators or tokens explicitly allowed to select an upstream.
	XSohaUpstreamID SohaUpstreamIDHeader `json:"X-Soha-Upstream-ID,omitempty"`

	// XSohaRouteTrace When true, returns redacted Soha route trace response headers for relay managers or tokens with allowRouteTrace explicitly enabled.
	XSohaRouteTrace SohaRouteTraceHeader `json:"X-Soha-Route-Trace,omitempty"`

	// XSohaCacheMode Optional relay cache behavior hint.
	XSohaCacheMode CreateAIGatewayOpenAICompatibleProviderAudioSpeechParamsXSohaCacheMode `json:"X-Soha-Cache-Mode,omitempty"`
}

// CreateAIGatewayOpenAICompatibleProviderAudioSpeechParamsXSohaCacheMode defines parameters for CreateAIGatewayOpenAICompatibleProviderAudioSpeech.
type CreateAIGatewayOpenAICompatibleProviderAudioSpeechParamsXSohaCacheMode string

// CreateAIGatewayOpenAICompatibleProviderAudioSpeechParamsOpenaiCompatibleProvider defines parameters for CreateAIGatewayOpenAICompatibleProviderAudioSpeech.
type CreateAIGatewayOpenAICompatibleProviderAudioSpeechParamsOpenaiCompatibleProvider string

// CreateAIGatewayOpenAICompatibleProviderAudioTranscriptionParams defines parameters for CreateAIGatewayOpenAICompatibleProviderAudioTranscription.
type CreateAIGatewayOpenAICompatibleProviderAudioTranscriptionParams struct {
	// XSohaUpstreamID Optional debug override for administrators or tokens explicitly allowed to select an upstream.
	XSohaUpstreamID SohaUpstreamIDHeader `json:"X-Soha-Upstream-ID,omitempty"`

	// XSohaRouteTrace When true, returns redacted Soha route trace response headers for relay managers or tokens with allowRouteTrace explicitly enabled.
	XSohaRouteTrace SohaRouteTraceHeader `json:"X-Soha-Route-Trace,omitempty"`

	// XSohaCacheMode Optional relay cache behavior hint.
	XSohaCacheMode CreateAIGatewayOpenAICompatibleProviderAudioTranscriptionParamsXSohaCacheMode `json:"X-Soha-Cache-Mode,omitempty"`
}

// CreateAIGatewayOpenAICompatibleProviderAudioTranscriptionParamsXSohaCacheMode defines parameters for CreateAIGatewayOpenAICompatibleProviderAudioTranscription.
type CreateAIGatewayOpenAICompatibleProviderAudioTranscriptionParamsXSohaCacheMode string

// CreateAIGatewayOpenAICompatibleProviderAudioTranscriptionParamsOpenaiCompatibleProvider defines parameters for CreateAIGatewayOpenAICompatibleProviderAudioTranscription.
type CreateAIGatewayOpenAICompatibleProviderAudioTranscriptionParamsOpenaiCompatibleProvider string

// CreateAIGatewayOpenAICompatibleProviderAudioTranslationParams defines parameters for CreateAIGatewayOpenAICompatibleProviderAudioTranslation.
type CreateAIGatewayOpenAICompatibleProviderAudioTranslationParams struct {
	// XSohaUpstreamID Optional debug override for administrators or tokens explicitly allowed to select an upstream.
	XSohaUpstreamID SohaUpstreamIDHeader `json:"X-Soha-Upstream-ID,omitempty"`

	// XSohaRouteTrace When true, returns redacted Soha route trace response headers for relay managers or tokens with allowRouteTrace explicitly enabled.
	XSohaRouteTrace SohaRouteTraceHeader `json:"X-Soha-Route-Trace,omitempty"`

	// XSohaCacheMode Optional relay cache behavior hint.
	XSohaCacheMode CreateAIGatewayOpenAICompatibleProviderAudioTranslationParamsXSohaCacheMode `json:"X-Soha-Cache-Mode,omitempty"`
}

// CreateAIGatewayOpenAICompatibleProviderAudioTranslationParamsXSohaCacheMode defines parameters for CreateAIGatewayOpenAICompatibleProviderAudioTranslation.
type CreateAIGatewayOpenAICompatibleProviderAudioTranslationParamsXSohaCacheMode string

// CreateAIGatewayOpenAICompatibleProviderAudioTranslationParamsOpenaiCompatibleProvider defines parameters for CreateAIGatewayOpenAICompatibleProviderAudioTranslation.
type CreateAIGatewayOpenAICompatibleProviderAudioTranslationParamsOpenaiCompatibleProvider string

// CreateAIGatewayOpenAICompatibleProviderChatCompletionParams defines parameters for CreateAIGatewayOpenAICompatibleProviderChatCompletion.
type CreateAIGatewayOpenAICompatibleProviderChatCompletionParams struct {
	// XSohaUpstreamID Optional debug override for administrators or tokens explicitly allowed to select an upstream.
	XSohaUpstreamID SohaUpstreamIDHeader `json:"X-Soha-Upstream-ID,omitempty"`

	// XSohaRouteTrace When true, returns redacted Soha route trace response headers for relay managers or tokens with allowRouteTrace explicitly enabled.
	XSohaRouteTrace SohaRouteTraceHeader `json:"X-Soha-Route-Trace,omitempty"`

	// XSohaCacheMode Optional relay cache behavior hint.
	XSohaCacheMode CreateAIGatewayOpenAICompatibleProviderChatCompletionParamsXSohaCacheMode `json:"X-Soha-Cache-Mode,omitempty"`
}

// CreateAIGatewayOpenAICompatibleProviderChatCompletionParamsXSohaCacheMode defines parameters for CreateAIGatewayOpenAICompatibleProviderChatCompletion.
type CreateAIGatewayOpenAICompatibleProviderChatCompletionParamsXSohaCacheMode string

// CreateAIGatewayOpenAICompatibleProviderChatCompletionParamsOpenaiCompatibleProvider defines parameters for CreateAIGatewayOpenAICompatibleProviderChatCompletion.
type CreateAIGatewayOpenAICompatibleProviderChatCompletionParamsOpenaiCompatibleProvider string

// CreateAIGatewayOpenAICompatibleProviderEmbeddingParams defines parameters for CreateAIGatewayOpenAICompatibleProviderEmbedding.
type CreateAIGatewayOpenAICompatibleProviderEmbeddingParams struct {
	// XSohaUpstreamID Optional debug override for administrators or tokens explicitly allowed to select an upstream.
	XSohaUpstreamID SohaUpstreamIDHeader `json:"X-Soha-Upstream-ID,omitempty"`

	// XSohaRouteTrace When true, returns redacted Soha route trace response headers for relay managers or tokens with allowRouteTrace explicitly enabled.
	XSohaRouteTrace SohaRouteTraceHeader `json:"X-Soha-Route-Trace,omitempty"`

	// XSohaCacheMode Optional relay cache behavior hint.
	XSohaCacheMode CreateAIGatewayOpenAICompatibleProviderEmbeddingParamsXSohaCacheMode `json:"X-Soha-Cache-Mode,omitempty"`
}

// CreateAIGatewayOpenAICompatibleProviderEmbeddingParamsXSohaCacheMode defines parameters for CreateAIGatewayOpenAICompatibleProviderEmbedding.
type CreateAIGatewayOpenAICompatibleProviderEmbeddingParamsXSohaCacheMode string

// CreateAIGatewayOpenAICompatibleProviderEmbeddingParamsOpenaiCompatibleProvider defines parameters for CreateAIGatewayOpenAICompatibleProviderEmbedding.
type CreateAIGatewayOpenAICompatibleProviderEmbeddingParamsOpenaiCompatibleProvider string

// CreateAIGatewayOpenAICompatibleProviderImageEditParams defines parameters for CreateAIGatewayOpenAICompatibleProviderImageEdit.
type CreateAIGatewayOpenAICompatibleProviderImageEditParams struct {
	// XSohaUpstreamID Optional debug override for administrators or tokens explicitly allowed to select an upstream.
	XSohaUpstreamID SohaUpstreamIDHeader `json:"X-Soha-Upstream-ID,omitempty"`

	// XSohaRouteTrace When true, returns redacted Soha route trace response headers for relay managers or tokens with allowRouteTrace explicitly enabled.
	XSohaRouteTrace SohaRouteTraceHeader `json:"X-Soha-Route-Trace,omitempty"`

	// XSohaCacheMode Optional relay cache behavior hint.
	XSohaCacheMode CreateAIGatewayOpenAICompatibleProviderImageEditParamsXSohaCacheMode `json:"X-Soha-Cache-Mode,omitempty"`
}

// CreateAIGatewayOpenAICompatibleProviderImageEditParamsXSohaCacheMode defines parameters for CreateAIGatewayOpenAICompatibleProviderImageEdit.
type CreateAIGatewayOpenAICompatibleProviderImageEditParamsXSohaCacheMode string

// CreateAIGatewayOpenAICompatibleProviderImageEditParamsOpenaiCompatibleProvider defines parameters for CreateAIGatewayOpenAICompatibleProviderImageEdit.
type CreateAIGatewayOpenAICompatibleProviderImageEditParamsOpenaiCompatibleProvider string

// CreateAIGatewayOpenAICompatibleProviderImageGenerationParams defines parameters for CreateAIGatewayOpenAICompatibleProviderImageGeneration.
type CreateAIGatewayOpenAICompatibleProviderImageGenerationParams struct {
	// XSohaUpstreamID Optional debug override for administrators or tokens explicitly allowed to select an upstream.
	XSohaUpstreamID SohaUpstreamIDHeader `json:"X-Soha-Upstream-ID,omitempty"`

	// XSohaRouteTrace When true, returns redacted Soha route trace response headers for relay managers or tokens with allowRouteTrace explicitly enabled.
	XSohaRouteTrace SohaRouteTraceHeader `json:"X-Soha-Route-Trace,omitempty"`

	// XSohaCacheMode Optional relay cache behavior hint.
	XSohaCacheMode CreateAIGatewayOpenAICompatibleProviderImageGenerationParamsXSohaCacheMode `json:"X-Soha-Cache-Mode,omitempty"`
}

// CreateAIGatewayOpenAICompatibleProviderImageGenerationParamsXSohaCacheMode defines parameters for CreateAIGatewayOpenAICompatibleProviderImageGeneration.
type CreateAIGatewayOpenAICompatibleProviderImageGenerationParamsXSohaCacheMode string

// CreateAIGatewayOpenAICompatibleProviderImageGenerationParamsOpenaiCompatibleProvider defines parameters for CreateAIGatewayOpenAICompatibleProviderImageGeneration.
type CreateAIGatewayOpenAICompatibleProviderImageGenerationParamsOpenaiCompatibleProvider string

// CreateAIGatewayOpenAICompatibleProviderImageVariationParams defines parameters for CreateAIGatewayOpenAICompatibleProviderImageVariation.
type CreateAIGatewayOpenAICompatibleProviderImageVariationParams struct {
	// XSohaUpstreamID Optional debug override for administrators or tokens explicitly allowed to select an upstream.
	XSohaUpstreamID SohaUpstreamIDHeader `json:"X-Soha-Upstream-ID,omitempty"`

	// XSohaRouteTrace When true, returns redacted Soha route trace response headers for relay managers or tokens with allowRouteTrace explicitly enabled.
	XSohaRouteTrace SohaRouteTraceHeader `json:"X-Soha-Route-Trace,omitempty"`

	// XSohaCacheMode Optional relay cache behavior hint.
	XSohaCacheMode CreateAIGatewayOpenAICompatibleProviderImageVariationParamsXSohaCacheMode `json:"X-Soha-Cache-Mode,omitempty"`
}

// CreateAIGatewayOpenAICompatibleProviderImageVariationParamsXSohaCacheMode defines parameters for CreateAIGatewayOpenAICompatibleProviderImageVariation.
type CreateAIGatewayOpenAICompatibleProviderImageVariationParamsXSohaCacheMode string

// CreateAIGatewayOpenAICompatibleProviderImageVariationParamsOpenaiCompatibleProvider defines parameters for CreateAIGatewayOpenAICompatibleProviderImageVariation.
type CreateAIGatewayOpenAICompatibleProviderImageVariationParamsOpenaiCompatibleProvider string

// ListAIGatewayOpenAICompatibleProviderModelsParamsOpenaiCompatibleProvider defines parameters for ListAIGatewayOpenAICompatibleProviderModels.
type ListAIGatewayOpenAICompatibleProviderModelsParamsOpenaiCompatibleProvider string

// CreateAIGatewayOpenAICompatibleProviderResponseParams defines parameters for CreateAIGatewayOpenAICompatibleProviderResponse.
type CreateAIGatewayOpenAICompatibleProviderResponseParams struct {
	// XSohaUpstreamID Optional debug override for administrators or tokens explicitly allowed to select an upstream.
	XSohaUpstreamID SohaUpstreamIDHeader `json:"X-Soha-Upstream-ID,omitempty"`

	// XSohaRouteTrace When true, returns redacted Soha route trace response headers for relay managers or tokens with allowRouteTrace explicitly enabled.
	XSohaRouteTrace SohaRouteTraceHeader `json:"X-Soha-Route-Trace,omitempty"`

	// XSohaCacheMode Optional relay cache behavior hint.
	XSohaCacheMode CreateAIGatewayOpenAICompatibleProviderResponseParamsXSohaCacheMode `json:"X-Soha-Cache-Mode,omitempty"`
}

// CreateAIGatewayOpenAICompatibleProviderResponseParamsXSohaCacheMode defines parameters for CreateAIGatewayOpenAICompatibleProviderResponse.
type CreateAIGatewayOpenAICompatibleProviderResponseParamsXSohaCacheMode string

// CreateAIGatewayOpenAICompatibleProviderResponseParamsOpenaiCompatibleProvider defines parameters for CreateAIGatewayOpenAICompatibleProviderResponse.
type CreateAIGatewayOpenAICompatibleProviderResponseParamsOpenaiCompatibleProvider string

// ListPersonalAccessTokensParams defines parameters for ListPersonalAccessTokens.
type ListPersonalAccessTokensParams struct {
	Scope  string `form:"scope,omitempty" json:"scope,omitempty"`
	UserID string `form:"userId,omitempty" json:"userId,omitempty"`
}

// GetAIGatewayRelayCacheStatsParams defines parameters for GetAIGatewayRelayCacheStats.
type GetAIGatewayRelayCacheStatsParams struct {
	WindowHours int    `form:"windowHours,omitempty" json:"windowHours,omitempty"`
	PublicModel string `form:"publicModel,omitempty" json:"publicModel,omitempty"`
	UpstreamID  string `form:"upstreamId,omitempty" json:"upstreamId,omitempty"`
}

// GetAIGatewayRelayMetricsParams defines parameters for GetAIGatewayRelayMetrics.
type GetAIGatewayRelayMetricsParams struct {
	WindowHours int    `form:"windowHours,omitempty" json:"windowHours,omitempty"`
	PublicModel string `form:"publicModel,omitempty" json:"publicModel,omitempty"`
	UpstreamID  string `form:"upstreamId,omitempty" json:"upstreamId,omitempty"`
}

// ListAIGatewayRelayModelCallsParams defines parameters for ListAIGatewayRelayModelCalls.
type ListAIGatewayRelayModelCallsParams struct {
	ActorType    string                                         `form:"actorType,omitempty" json:"actorType,omitempty"`
	ActorID      string                                         `form:"actorId,omitempty" json:"actorId,omitempty"`
	TokenID      string                                         `form:"tokenId,omitempty" json:"tokenId,omitempty"`
	PublicModel  string                                         `form:"publicModel,omitempty" json:"publicModel,omitempty"`
	UpstreamID   string                                         `form:"upstreamId,omitempty" json:"upstreamId,omitempty"`
	ProviderKind ListAIGatewayRelayModelCallsParamsProviderKind `form:"providerKind,omitempty" json:"providerKind,omitempty"`
	Status       ListAIGatewayRelayModelCallsParamsStatus       `form:"status,omitempty" json:"status,omitempty"`
	Endpoint     string                                         `form:"endpoint,omitempty" json:"endpoint,omitempty"`
	From         time.Time                                      `form:"from,omitempty" json:"from,omitempty"`
	To           time.Time                                      `form:"to,omitempty" json:"to,omitempty"`
	Limit        int                                            `form:"limit,omitempty" json:"limit,omitempty"`
}

// ListAIGatewayRelayModelCallsParamsProviderKind defines parameters for ListAIGatewayRelayModelCalls.
type ListAIGatewayRelayModelCallsParamsProviderKind string

// ListAIGatewayRelayModelCallsParamsStatus defines parameters for ListAIGatewayRelayModelCalls.
type ListAIGatewayRelayModelCallsParamsStatus string

// ListAIGatewayRelayModelRoutesParams defines parameters for ListAIGatewayRelayModelRoutes.
type ListAIGatewayRelayModelRoutesParams struct {
	PublicModel     string                                          `form:"publicModel,omitempty" json:"publicModel,omitempty"`
	ProviderKind    ListAIGatewayRelayModelRoutesParamsProviderKind `form:"providerKind,omitempty" json:"providerKind,omitempty"`
	UpstreamID      string                                          `form:"upstreamId,omitempty" json:"upstreamId,omitempty"`
	RouteGroup      string                                          `form:"routeGroup,omitempty" json:"routeGroup,omitempty"`
	IncludeDisabled bool                                            `form:"includeDisabled,omitempty" json:"includeDisabled,omitempty"`
}

// ListAIGatewayRelayModelRoutesParamsProviderKind defines parameters for ListAIGatewayRelayModelRoutes.
type ListAIGatewayRelayModelRoutesParamsProviderKind string

// ListAIGatewayRelayUpstreamsParams defines parameters for ListAIGatewayRelayUpstreams.
type ListAIGatewayRelayUpstreamsParams struct {
	ProviderKind ListAIGatewayRelayUpstreamsParamsProviderKind `form:"providerKind,omitempty" json:"providerKind,omitempty"`
	Status       ListAIGatewayRelayUpstreamsParamsStatus       `form:"status,omitempty" json:"status,omitempty"`
	IncludeAll   bool                                          `form:"includeAll,omitempty" json:"includeAll,omitempty"`
}

// ListAIGatewayRelayUpstreamsParamsProviderKind defines parameters for ListAIGatewayRelayUpstreams.
type ListAIGatewayRelayUpstreamsParamsProviderKind string

// ListAIGatewayRelayUpstreamsParamsStatus defines parameters for ListAIGatewayRelayUpstreams.
type ListAIGatewayRelayUpstreamsParamsStatus string

// GetAgentProviderRegistrySnapshotParams defines parameters for GetAgentProviderRegistrySnapshot.
type GetAgentProviderRegistrySnapshotParams struct {
	RunnerID string `form:"runnerId" json:"runnerId"`
}

// TransitionAgentProviderRolloutParamsRolloutAction defines parameters for TransitionAgentProviderRollout.
type TransitionAgentProviderRolloutParamsRolloutAction string

// ListKnowledgeDocumentsParams defines parameters for ListKnowledgeDocuments.
type ListKnowledgeDocumentsParams struct {
	Limit KnowledgeListLimit `form:"limit,omitempty" json:"limit,omitempty"`
}

// ListKnowledgeIndexRevisionsParams defines parameters for ListKnowledgeIndexRevisions.
type ListKnowledgeIndexRevisionsParams struct {
	Limit KnowledgeListLimit `form:"limit,omitempty" json:"limit,omitempty"`
}

// ListKnowledgeSyncRunsParams defines parameters for ListKnowledgeSyncRuns.
type ListKnowledgeSyncRunsParams struct {
	Limit KnowledgeListLimit `form:"limit,omitempty" json:"limit,omitempty"`
}

// ListKnowledgeConnectorsParams defines parameters for ListKnowledgeConnectors.
type ListKnowledgeConnectorsParams struct {
	Limit KnowledgeListLimit `form:"limit,omitempty" json:"limit,omitempty"`
}

// ListAIMemoryRecordsParams defines parameters for ListAIMemoryRecords.
type ListAIMemoryRecordsParams struct {
	OwnerType string `form:"ownerType,omitempty" json:"ownerType,omitempty"`
	OwnerID   string `form:"ownerId,omitempty" json:"ownerId,omitempty"`
}

// ListApplicationEnvironmentsParams defines parameters for ListApplicationEnvironments.
type ListApplicationEnvironmentsParams struct {
	ApplicationID string `form:"applicationId,omitempty" json:"applicationId,omitempty"`
	EnvironmentID string `form:"environmentId,omitempty" json:"environmentId,omitempty"`
}

// ListApplicationsParams defines parameters for ListApplications.
type ListApplicationsParams struct {
	Search string `form:"search,omitempty" json:"search,omitempty"`
	Limit  int    `form:"limit,omitempty" json:"limit,omitempty"`
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

// ExecuteKubernetesResourceCreateParams defines parameters for ExecuteKubernetesResourceCreate.
type ExecuteKubernetesResourceCreateParams struct {
	IdempotencyKey IdempotencyKey `json:"Idempotency-Key"`
}

// ListComputeAccessSourcesParams defines parameters for ListComputeAccessSources.
type ListComputeAccessSourcesParams struct {
	SourceType  ComputeAccessSourceType `form:"sourceType,omitempty" json:"sourceType,omitempty"`
	ProviderKey string                  `form:"providerKey,omitempty" json:"providerKey,omitempty"`
	Cursor      ComputeCursor           `form:"cursor,omitempty" json:"cursor,omitempty"`
	Limit       ComputeLimit            `form:"limit,omitempty" json:"limit,omitempty"`
}

// ListComputeProviderInstancesParams defines parameters for ListComputeProviderInstances.
type ListComputeProviderInstancesParams struct {
	Domain      ComputeProviderDomain `form:"domain,omitempty" json:"domain,omitempty"`
	ProviderKey string                `form:"providerKey,omitempty" json:"providerKey,omitempty"`
	Cursor      ComputeCursor         `form:"cursor,omitempty" json:"cursor,omitempty"`
	Limit       ComputeLimit          `form:"limit,omitempty" json:"limit,omitempty"`
}

// DiscoverComputeProviderInstanceParams defines parameters for DiscoverComputeProviderInstance.
type DiscoverComputeProviderInstanceParams struct {
	IdempotencyKey IdempotencyKey `json:"Idempotency-Key"`
}

// CheckComputeProviderInstanceHealthParams defines parameters for CheckComputeProviderInstanceHealth.
type CheckComputeProviderInstanceHealthParams struct {
	IdempotencyKey IdempotencyKey `json:"Idempotency-Key"`
}

// ListComputeProvidersParams defines parameters for ListComputeProviders.
type ListComputeProvidersParams struct {
	Domain ComputeProviderDomain `form:"domain,omitempty" json:"domain,omitempty"`
	Source ComputeProviderSource `form:"source,omitempty" json:"source,omitempty"`
	Cursor ComputeCursor         `form:"cursor,omitempty" json:"cursor,omitempty"`
	Limit  ComputeLimit          `form:"limit,omitempty" json:"limit,omitempty"`
}

// ExecuteComputeResourceActionParams defines parameters for ExecuteComputeResourceAction.
type ExecuteComputeResourceActionParams struct {
	IdempotencyKey IdempotencyKey `json:"Idempotency-Key"`
}

// ListComputeResourceRelationsParams defines parameters for ListComputeResourceRelations.
type ListComputeResourceRelationsParams struct {
	Cursor ComputeCursor `form:"cursor,omitempty" json:"cursor,omitempty"`
	Limit  ComputeLimit  `form:"limit,omitempty" json:"limit,omitempty"`
}

// ListComputeTasksParams defines parameters for ListComputeTasks.
type ListComputeTasksParams struct {
	Domain      ComputeTaskDomain   `form:"domain,omitempty" json:"domain,omitempty"`
	ProviderKey string              `form:"providerKey,omitempty" json:"providerKey,omitempty"`
	Status      ComputeTaskStatus   `form:"status,omitempty" json:"status,omitempty"`
	Category    ComputeTaskCategory `form:"category,omitempty" json:"category,omitempty"`

	// ResourceKind Exact normalized resource kind referenced by the task.
	ResourceKind string `form:"resourceKind,omitempty" json:"resourceKind,omitempty"`

	// ResourceID Exact normalized resource identifier referenced by the task.
	ResourceID string        `form:"resourceId,omitempty" json:"resourceId,omitempty"`
	Cursor     ComputeCursor `form:"cursor,omitempty" json:"cursor,omitempty"`
	Limit      ComputeLimit  `form:"limit,omitempty" json:"limit,omitempty"`
}

// CancelComputeTaskParams defines parameters for CancelComputeTask.
type CancelComputeTaskParams struct {
	IdempotencyKey IdempotencyKey `json:"Idempotency-Key"`
}

// RetryComputeTaskParams defines parameters for RetryComputeTask.
type RetryComputeTaskParams struct {
	IdempotencyKey IdempotencyKey `json:"Idempotency-Key"`
}

// ListDeliveryArtifactsParams defines parameters for ListDeliveryArtifacts.
type ListDeliveryArtifactsParams struct {
	ApplicationID            string `form:"applicationId,omitempty" json:"applicationId,omitempty"`
	ApplicationEnvironmentID string `form:"applicationEnvironmentId,omitempty" json:"applicationEnvironmentId,omitempty"`
	WorkflowRunID            string `form:"workflowRunId,omitempty" json:"workflowRunId,omitempty"`
	WorkflowNodeID           string `form:"workflowNodeId,omitempty" json:"workflowNodeId,omitempty"`
	ReleaseBundleID          string `form:"releaseBundleId,omitempty" json:"releaseBundleId,omitempty"`
	ExecutionTaskID          string `form:"executionTaskId,omitempty" json:"executionTaskId,omitempty"`
	Kind                     string `form:"kind,omitempty" json:"kind,omitempty"`
	Status                   string `form:"status,omitempty" json:"status,omitempty"`
	Limit                    int    `form:"limit,omitempty" json:"limit,omitempty"`
}

// ListDeliveryExecutionTasksParams defines parameters for ListDeliveryExecutionTasks.
type ListDeliveryExecutionTasksParams struct {
	ApplicationID            string `form:"applicationId,omitempty" json:"applicationId,omitempty"`
	ApplicationEnvironmentID string `form:"applicationEnvironmentId,omitempty" json:"applicationEnvironmentId,omitempty"`
	ReleaseBundleID          string `form:"releaseBundleId,omitempty" json:"releaseBundleId,omitempty"`
	Status                   string `form:"status,omitempty" json:"status,omitempty"`
	ProviderKind             string `form:"providerKind,omitempty" json:"providerKind,omitempty"`
	Limit                    int    `form:"limit,omitempty" json:"limit,omitempty"`
}

// ListReleaseBundlesParams defines parameters for ListReleaseBundles.
type ListReleaseBundlesParams struct {
	ApplicationID            string `form:"applicationId,omitempty" json:"applicationId,omitempty"`
	ApplicationEnvironmentID string `form:"applicationEnvironmentId,omitempty" json:"applicationEnvironmentId,omitempty"`
	Limit                    int    `form:"limit,omitempty" json:"limit,omitempty"`
}

// ListGitLabBranchesParams defines parameters for ListGitLabBranches.
type ListGitLabBranchesParams struct {
	ProjectID GitLabProjectIDQuery `form:"projectId" json:"projectId"`
	Search    string               `form:"search,omitempty" json:"search,omitempty"`
	Limit     int                  `form:"limit,omitempty" json:"limit,omitempty"`
}

// ListGitLabCommitsParams defines parameters for ListGitLabCommits.
type ListGitLabCommitsParams struct {
	ProjectID GitLabProjectIDQuery `form:"projectId" json:"projectId"`
	Search    string               `form:"search,omitempty" json:"search,omitempty"`
	Page      int                  `form:"page,omitempty" json:"page,omitempty"`
	Limit     int                  `form:"limit,omitempty" json:"limit,omitempty"`
}

// ListGitLabProjectsParams defines parameters for ListGitLabProjects.
type ListGitLabProjectsParams struct {
	Search string `form:"search,omitempty" json:"search,omitempty"`
	Limit  int    `form:"limit,omitempty" json:"limit,omitempty"`
}

// ListGitLabTagsParams defines parameters for ListGitLabTags.
type ListGitLabTagsParams struct {
	ProjectID GitLabProjectIDQuery `form:"projectId" json:"projectId"`
	Search    string               `form:"search,omitempty" json:"search,omitempty"`
	Limit     int                  `form:"limit,omitempty" json:"limit,omitempty"`
}

// ListMarketplacePluginsParams defines parameters for ListMarketplacePlugins.
type ListMarketplacePluginsParams struct {
	Q              string `form:"q,omitempty" json:"q,omitempty"`
	Type           string `form:"type,omitempty" json:"type,omitempty"`
	Publisher      string `form:"publisher,omitempty" json:"publisher,omitempty"`
	SourceID       string `form:"sourceId,omitempty" json:"sourceId,omitempty"`
	MarketplaceURL string `form:"marketplaceUrl,omitempty" json:"marketplaceUrl,omitempty"`
	Version        string `form:"version,omitempty" json:"version,omitempty"`
}

// GetMarketplacePluginParams defines parameters for GetMarketplacePlugin.
type GetMarketplacePluginParams struct {
	SourceID       string `form:"sourceId,omitempty" json:"sourceId,omitempty"`
	MarketplaceURL string `form:"marketplaceUrl,omitempty" json:"marketplaceUrl,omitempty"`
	Version        string `form:"version,omitempty" json:"version,omitempty"`
}

// ListRepositoriesParams defines parameters for ListRepositories.
type ListRepositoriesParams struct {
	ApplicationID string `form:"applicationId,omitempty" json:"applicationId,omitempty"`
	Search        string `form:"search,omitempty" json:"search,omitempty"`
	Limit         int    `form:"limit,omitempty" json:"limit,omitempty"`
}

// ListSourceRepositoriesParams defines parameters for ListSourceRepositories.
type ListSourceRepositoriesParams struct {
	Search string `form:"search,omitempty" json:"search,omitempty"`
	Cursor string `form:"cursor,omitempty" json:"cursor,omitempty"`
	Limit  int    `form:"limit,omitempty" json:"limit,omitempty"`
}

// GetSourceRepositoryFileParams defines parameters for GetSourceRepositoryFile.
type GetSourceRepositoryFileParams struct {
	Ref  string `form:"ref" json:"ref"`
	Path string `form:"path" json:"path"`
}

// ListSystemIntegrationsParams defines parameters for ListSystemIntegrations.
type ListSystemIntegrationsParams struct {
	Category     SystemIntegrationCategory `form:"category,omitempty" json:"category,omitempty"`
	ProviderType string                    `form:"providerType,omitempty" json:"providerType,omitempty"`
	Enabled      bool                      `form:"enabled,omitempty" json:"enabled,omitempty"`
}

// DecideAIGatewayApprovalRequestJSONRequestBody defines body for DecideAIGatewayApprovalRequest for application/json ContentType.
type DecideAIGatewayApprovalRequestJSONRequestBody = ApprovalDecisionInput

// CreateAIGatewayAnthropicMessageJSONRequestBody defines body for CreateAIGatewayAnthropicMessage for application/json ContentType.
type CreateAIGatewayAnthropicMessageJSONRequestBody = AnthropicMessagesRequest

// CreateAIGatewayCohereRerankJSONRequestBody defines body for CreateAIGatewayCohereRerank for application/json ContentType.
type CreateAIGatewayCohereRerankJSONRequestBody = CohereRerankRequest

// CreateAIGatewayGeminiInteractionsJSONRequestBody defines body for CreateAIGatewayGeminiInteractions for application/json ContentType.
type CreateAIGatewayGeminiInteractionsJSONRequestBody = GeminiInteractionsRequest

// CreateAIGatewayGeminiGenerateContentJSONRequestBody defines body for CreateAIGatewayGeminiGenerateContent for application/json ContentType.
type CreateAIGatewayGeminiGenerateContentJSONRequestBody = GeminiGenerateContentRequest

// CreateAIGatewayGeminiStreamGenerateContentJSONRequestBody defines body for CreateAIGatewayGeminiStreamGenerateContent for application/json ContentType.
type CreateAIGatewayGeminiStreamGenerateContentJSONRequestBody = GeminiGenerateContentRequest

// CreateAIGatewayOpenAIAudioSpeechJSONRequestBody defines body for CreateAIGatewayOpenAIAudioSpeech for application/json ContentType.
type CreateAIGatewayOpenAIAudioSpeechJSONRequestBody = OpenAIAudioSpeechRequest

// CreateAIGatewayOpenAIAudioTranscriptionMultipartRequestBody defines body for CreateAIGatewayOpenAIAudioTranscription for multipart/form-data ContentType.
type CreateAIGatewayOpenAIAudioTranscriptionMultipartRequestBody = OpenAIAudioTranscriptionRequest

// CreateAIGatewayOpenAIAudioTranslationMultipartRequestBody defines body for CreateAIGatewayOpenAIAudioTranslation for multipart/form-data ContentType.
type CreateAIGatewayOpenAIAudioTranslationMultipartRequestBody = OpenAIAudioTranslationRequest

// CreateAIGatewayOpenAIChatCompletionJSONRequestBody defines body for CreateAIGatewayOpenAIChatCompletion for application/json ContentType.
type CreateAIGatewayOpenAIChatCompletionJSONRequestBody = OpenAIChatCompletionRequest

// CreateAIGatewayOpenAIEmbeddingJSONRequestBody defines body for CreateAIGatewayOpenAIEmbedding for application/json ContentType.
type CreateAIGatewayOpenAIEmbeddingJSONRequestBody = OpenAIEmbeddingRequest

// CreateAIGatewayOpenAIImageEditMultipartRequestBody defines body for CreateAIGatewayOpenAIImageEdit for multipart/form-data ContentType.
type CreateAIGatewayOpenAIImageEditMultipartRequestBody = OpenAIImageEditRequest

// CreateAIGatewayOpenAIImageGenerationJSONRequestBody defines body for CreateAIGatewayOpenAIImageGeneration for application/json ContentType.
type CreateAIGatewayOpenAIImageGenerationJSONRequestBody = OpenAIImageGenerationRequest

// CreateAIGatewayOpenAIImageVariationMultipartRequestBody defines body for CreateAIGatewayOpenAIImageVariation for multipart/form-data ContentType.
type CreateAIGatewayOpenAIImageVariationMultipartRequestBody = OpenAIImageVariationRequest

// CreateAIGatewayOpenAIResponseJSONRequestBody defines body for CreateAIGatewayOpenAIResponse for application/json ContentType.
type CreateAIGatewayOpenAIResponseJSONRequestBody = OpenAIResponseRequest

// CreateAIGatewayOpenAICompatibleProviderAudioSpeechJSONRequestBody defines body for CreateAIGatewayOpenAICompatibleProviderAudioSpeech for application/json ContentType.
type CreateAIGatewayOpenAICompatibleProviderAudioSpeechJSONRequestBody = OpenAIAudioSpeechRequest

// CreateAIGatewayOpenAICompatibleProviderAudioTranscriptionMultipartRequestBody defines body for CreateAIGatewayOpenAICompatibleProviderAudioTranscription for multipart/form-data ContentType.
type CreateAIGatewayOpenAICompatibleProviderAudioTranscriptionMultipartRequestBody = OpenAIAudioTranscriptionRequest

// CreateAIGatewayOpenAICompatibleProviderAudioTranslationMultipartRequestBody defines body for CreateAIGatewayOpenAICompatibleProviderAudioTranslation for multipart/form-data ContentType.
type CreateAIGatewayOpenAICompatibleProviderAudioTranslationMultipartRequestBody = OpenAIAudioTranslationRequest

// CreateAIGatewayOpenAICompatibleProviderChatCompletionJSONRequestBody defines body for CreateAIGatewayOpenAICompatibleProviderChatCompletion for application/json ContentType.
type CreateAIGatewayOpenAICompatibleProviderChatCompletionJSONRequestBody = OpenAIChatCompletionRequest

// CreateAIGatewayOpenAICompatibleProviderEmbeddingJSONRequestBody defines body for CreateAIGatewayOpenAICompatibleProviderEmbedding for application/json ContentType.
type CreateAIGatewayOpenAICompatibleProviderEmbeddingJSONRequestBody = OpenAIEmbeddingRequest

// CreateAIGatewayOpenAICompatibleProviderImageEditMultipartRequestBody defines body for CreateAIGatewayOpenAICompatibleProviderImageEdit for multipart/form-data ContentType.
type CreateAIGatewayOpenAICompatibleProviderImageEditMultipartRequestBody = OpenAIImageEditRequest

// CreateAIGatewayOpenAICompatibleProviderImageGenerationJSONRequestBody defines body for CreateAIGatewayOpenAICompatibleProviderImageGeneration for application/json ContentType.
type CreateAIGatewayOpenAICompatibleProviderImageGenerationJSONRequestBody = OpenAIImageGenerationRequest

// CreateAIGatewayOpenAICompatibleProviderImageVariationMultipartRequestBody defines body for CreateAIGatewayOpenAICompatibleProviderImageVariation for multipart/form-data ContentType.
type CreateAIGatewayOpenAICompatibleProviderImageVariationMultipartRequestBody = OpenAIImageVariationRequest

// CreateAIGatewayOpenAICompatibleProviderResponseJSONRequestBody defines body for CreateAIGatewayOpenAICompatibleProviderResponse for application/json ContentType.
type CreateAIGatewayOpenAICompatibleProviderResponseJSONRequestBody = OpenAIResponseRequest

// CreatePersonalAccessTokenJSONRequestBody defines body for CreatePersonalAccessToken for application/json ContentType.
type CreatePersonalAccessTokenJSONRequestBody = PersonalAccessTokenInput

// RotatePersonalAccessTokenJSONRequestBody defines body for RotatePersonalAccessToken for application/json ContentType.
type RotatePersonalAccessTokenJSONRequestBody = TokenRotationInput

// GetAIGatewayPromptJSONRequestBody defines body for GetAIGatewayPrompt for application/json ContentType.
type GetAIGatewayPromptJSONRequestBody = PromptGetRequest

// PurgeAIGatewayRelayCacheJSONRequestBody defines body for PurgeAIGatewayRelayCache for application/json ContentType.
type PurgeAIGatewayRelayCacheJSONRequestBody = LLMRelayCachePurgeRequest

// CreateAIGatewayRelayModelRouteJSONRequestBody defines body for CreateAIGatewayRelayModelRoute for application/json ContentType.
type CreateAIGatewayRelayModelRouteJSONRequestBody = LLMModelRouteInput

// UpdateAIGatewayRelayModelRouteJSONRequestBody defines body for UpdateAIGatewayRelayModelRoute for application/json ContentType.
type UpdateAIGatewayRelayModelRouteJSONRequestBody = LLMModelRouteInput

// CreateAIGatewayRelayUpstreamJSONRequestBody defines body for CreateAIGatewayRelayUpstream for application/json ContentType.
type CreateAIGatewayRelayUpstreamJSONRequestBody = LLMUpstreamInput

// UpdateAIGatewayRelayUpstreamJSONRequestBody defines body for UpdateAIGatewayRelayUpstream for application/json ContentType.
type UpdateAIGatewayRelayUpstreamJSONRequestBody = LLMUpstreamInput

// TestAIGatewayRelayUpstreamJSONRequestBody defines body for TestAIGatewayRelayUpstream for application/json ContentType.
type TestAIGatewayRelayUpstreamJSONRequestBody = LLMUpstreamTestRequest

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

// CreateAgentProviderConformanceRunJSONRequestBody defines body for CreateAgentProviderConformanceRun for application/json ContentType.
type CreateAgentProviderConformanceRunJSONRequestBody = AgentProviderConformanceRunInput

// AcknowledgeAgentProviderRegistryJSONRequestBody defines body for AcknowledgeAgentProviderRegistry for application/json ContentType.
type AcknowledgeAgentProviderRegistryJSONRequestBody = AgentProviderRegistryAcknowledgement

// CreateAgentProviderRolloutJSONRequestBody defines body for CreateAgentProviderRollout for application/json ContentType.
type CreateAgentProviderRolloutJSONRequestBody = AgentProviderRolloutInput

// CreateMultiAgentPlanJSONRequestBody defines body for CreateMultiAgentPlan for application/json ContentType.
type CreateMultiAgentPlanJSONRequestBody = MultiAgentPlan

// CompleteMultiAgentSubtaskJSONRequestBody defines body for CompleteMultiAgentSubtask for application/json ContentType.
type CompleteMultiAgentSubtaskJSONRequestBody = MultiAgentSubtaskCompletionInput

// InspectAIContextJSONRequestBody defines body for InspectAIContext for application/json ContentType.
type InspectAIContextJSONRequestBody = AIContextBuildInput

// PutAIEnvironmentTemplateJSONRequestBody defines body for PutAIEnvironmentTemplate for application/json ContentType.
type PutAIEnvironmentTemplateJSONRequestBody = AIEnvironmentTemplateInput

// CreateEvaluationDatasetJSONRequestBody defines body for CreateEvaluationDataset for application/json ContentType.
type CreateEvaluationDatasetJSONRequestBody = EvaluationDataset

// PutEvaluationExecutorProfileJSONRequestBody defines body for PutEvaluationExecutorProfile for application/json ContentType.
type PutEvaluationExecutorProfileJSONRequestBody = EvaluationExecutorProfile

// PutEvaluationFeedbackJSONRequestBody defines body for PutEvaluationFeedback for application/json ContentType.
type PutEvaluationFeedbackJSONRequestBody = EvaluationFeedbackInput

// PutEvaluationGatePolicyJSONRequestBody defines body for PutEvaluationGatePolicy for application/json ContentType.
type PutEvaluationGatePolicyJSONRequestBody = EvaluationGatePolicyInput

// EvaluateEvaluationGateJSONRequestBody defines body for EvaluateEvaluationGate for application/json ContentType.
type EvaluateEvaluationGateJSONRequestBody = EvaluationGateRequest

// PutEvaluationReplayPlanJSONRequestBody defines body for PutEvaluationReplayPlan for application/json ContentType.
type PutEvaluationReplayPlanJSONRequestBody = EvaluationReplayPlanInput

// StartEvaluationRunJSONRequestBody defines body for StartEvaluationRun for application/json ContentType.
type StartEvaluationRunJSONRequestBody = EvaluationRun

// CompleteEvaluationRunJSONRequestBody defines body for CompleteEvaluationRun for application/json ContentType.
type CompleteEvaluationRunJSONRequestBody = EvaluationCompleteRunInput

// ExecuteEvaluationRunJSONRequestBody defines body for ExecuteEvaluationRun for application/json ContentType.
type ExecuteEvaluationRunJSONRequestBody = EvaluationExecuteRunInput

// CreateKnowledgeBaseJSONRequestBody defines body for CreateKnowledgeBase for application/json ContentType.
type CreateKnowledgeBaseJSONRequestBody = KnowledgeBaseInput

// UpdateKnowledgeBaseJSONRequestBody defines body for UpdateKnowledgeBase for application/json ContentType.
type UpdateKnowledgeBaseJSONRequestBody = KnowledgeBaseInput

// PutKnowledgeGraphRevisionJSONRequestBody defines body for PutKnowledgeGraphRevision for application/json ContentType.
type PutKnowledgeGraphRevisionJSONRequestBody = KnowledgeGraphRevision

// QueryKnowledgeGraphRevisionJSONRequestBody defines body for QueryKnowledgeGraphRevision for application/json ContentType.
type QueryKnowledgeGraphRevisionJSONRequestBody = KnowledgeGraphQueryInput

// CreateKnowledgeSourceJSONRequestBody defines body for CreateKnowledgeSource for application/json ContentType.
type CreateKnowledgeSourceJSONRequestBody = KnowledgeSourceInput

// CreateKnowledgeIngestionJobJSONRequestBody defines body for CreateKnowledgeIngestionJob for application/json ContentType.
type CreateKnowledgeIngestionJobJSONRequestBody = KnowledgeIngestionJobInput

// CreateKnowledgeConnectorJSONRequestBody defines body for CreateKnowledgeConnector for application/json ContentType.
type CreateKnowledgeConnectorJSONRequestBody = KnowledgeConnectorInput

// SearchKnowledgeJSONRequestBody defines body for SearchKnowledge for application/json ContentType.
type SearchKnowledgeJSONRequestBody = KnowledgeSearchRequest

// PutAIMemoryRecordJSONRequestBody defines body for PutAIMemoryRecord for application/json ContentType.
type PutAIMemoryRecordJSONRequestBody = AIMemoryWriteInput

// PutAIMemoryPolicyJSONRequestBody defines body for PutAIMemoryPolicy for application/json ContentType.
type PutAIMemoryPolicyJSONRequestBody = AIMemoryPolicyInput

// StartAIProductionOperationJSONRequestBody defines body for StartAIProductionOperation for application/json ContentType.
type StartAIProductionOperationJSONRequestBody = AIProductionOperationInput

// CreateApplicationEnvironmentJSONRequestBody defines body for CreateApplicationEnvironment for application/json ContentType.
type CreateApplicationEnvironmentJSONRequestBody = ApplicationEnvironmentInput

// UpdateApplicationEnvironmentJSONRequestBody defines body for UpdateApplicationEnvironment for application/json ContentType.
type UpdateApplicationEnvironmentJSONRequestBody = ApplicationEnvironmentInput

// CreateApplicationJSONRequestBody defines body for CreateApplication for application/json ContentType.
type CreateApplicationJSONRequestBody = ApplicationInput

// UpdateApplicationJSONRequestBody defines body for UpdateApplication for application/json ContentType.
type UpdateApplicationJSONRequestBody = ApplicationInput

// TriggerApplicationDeliveryActionJSONRequestBody defines body for TriggerApplicationDeliveryAction for application/json ContentType.
type TriggerApplicationDeliveryActionJSONRequestBody = ApplicationDeliveryActionRequest

// CreateApplicationServiceJSONRequestBody defines body for CreateApplicationService for application/json ContentType.
type CreateApplicationServiceJSONRequestBody = ApplicationServiceInput

// UpdateApplicationServiceJSONRequestBody defines body for UpdateApplicationService for application/json ContentType.
type UpdateApplicationServiceJSONRequestBody = ApplicationServiceInput

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

// CreateBuildTemplateJSONRequestBody defines body for CreateBuildTemplate for application/json ContentType.
type CreateBuildTemplateJSONRequestBody = BuildTemplateInput

// UpdateBuildTemplateJSONRequestBody defines body for UpdateBuildTemplate for application/json ContentType.
type UpdateBuildTemplateJSONRequestBody = BuildTemplateInput

// ExecuteKubernetesResourceCreateJSONRequestBody defines body for ExecuteKubernetesResourceCreate for application/json ContentType.
type ExecuteKubernetesResourceCreateJSONRequestBody = KubernetesResourceCreateRequest

// PreflightKubernetesResourceCreateJSONRequestBody defines body for PreflightKubernetesResourceCreate for application/json ContentType.
type PreflightKubernetesResourceCreateJSONRequestBody = KubernetesResourceCreateRequest

// DecideKubernetesResourceCreateScopeJSONRequestBody defines body for DecideKubernetesResourceCreateScope for application/json ContentType.
type DecideKubernetesResourceCreateScopeJSONRequestBody = KubernetesResourceCreateScopeDecisionRequest

// DiscoverComputeProviderInstanceJSONRequestBody defines body for DiscoverComputeProviderInstance for application/json ContentType.
type DiscoverComputeProviderInstanceJSONRequestBody = ComputeProviderDiscoverRequest

// CheckComputeProviderInstanceHealthJSONRequestBody defines body for CheckComputeProviderInstanceHealth for application/json ContentType.
type CheckComputeProviderInstanceHealthJSONRequestBody = ComputeProviderReadRequest

// ExecuteComputeResourceActionJSONRequestBody defines body for ExecuteComputeResourceAction for application/json ContentType.
type ExecuteComputeResourceActionJSONRequestBody = ComputeResourceActionRequest

// CancelComputeTaskJSONRequestBody defines body for CancelComputeTask for application/json ContentType.
type CancelComputeTaskJSONRequestBody = ComputeTaskMutationRequest

// RetryComputeTaskJSONRequestBody defines body for RetryComputeTask for application/json ContentType.
type RetryComputeTaskJSONRequestBody = ComputeTaskMutationRequest

// RecordAgentRunCallbackJSONRequestBody defines body for RecordAgentRunCallback for application/json ContentType.
type RecordAgentRunCallbackJSONRequestBody = AgentRunCallbackRequest

// ClaimAgentRunJSONRequestBody defines body for ClaimAgentRun for application/json ContentType.
type ClaimAgentRunJSONRequestBody = AgentRunClaimRequest

// RecordAgentRunToolCallJSONRequestBody defines body for RecordAgentRunToolCall for application/json ContentType.
type RecordAgentRunToolCallJSONRequestBody = AgentRunToolCallRequest

// RecordWorkbenchGlobalAssistantEventJSONRequestBody defines body for RecordWorkbenchGlobalAssistantEvent for application/json ContentType.
type RecordWorkbenchGlobalAssistantEventJSONRequestBody = WorkbenchGlobalAssistantOpenRequest

// StreamWorkbenchSessionMessageJSONRequestBody defines body for StreamWorkbenchSessionMessage for application/json ContentType.
type StreamWorkbenchSessionMessageJSONRequestBody = WorkbenchSendMessageStreamRequest

// RecordExecutionCallbackJSONRequestBody defines body for RecordExecutionCallback for application/json ContentType.
type RecordExecutionCallbackJSONRequestBody = ExecutionCallbackRequest

// ClaimExecutionTaskJSONRequestBody defines body for ClaimExecutionTask for application/json ContentType.
type ClaimExecutionTaskJSONRequestBody = ExecutionTaskClaimRequest

// CreateDeliveryPlanJSONRequestBody defines body for CreateDeliveryPlan for application/json ContentType.
type CreateDeliveryPlanJSONRequestBody = DeliveryPlanInput

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

// CreateRepositoryJSONRequestBody defines body for CreateRepository for application/json ContentType.
type CreateRepositoryJSONRequestBody = RepositoryInput

// UpdateRepositoryJSONRequestBody defines body for UpdateRepository for application/json ContentType.
type UpdateRepositoryJSONRequestBody = RepositoryInput

// UpdateAISkillsRegistryJSONRequestBody defines body for UpdateAISkillsRegistry for application/json ContentType.
type UpdateAISkillsRegistryJSONRequestBody = UpdateAISkillsRequest

// UpdateAIWorkbenchModelSettingsJSONRequestBody defines body for UpdateAIWorkbenchModelSettings for application/json ContentType.
type UpdateAIWorkbenchModelSettingsJSONRequestBody = UpdateAIWorkbenchModelRequest

// ApplyRuntimeConfigJSONRequestBody defines body for ApplyRuntimeConfig for application/json ContentType.
type ApplyRuntimeConfigJSONRequestBody = RuntimeConfigChangeRequest

// RollbackRuntimeConfigJSONRequestBody defines body for RollbackRuntimeConfig for application/json ContentType.
type RollbackRuntimeConfigJSONRequestBody = RuntimeConfigRollbackRequest

// ValidateRuntimeConfigJSONRequestBody defines body for ValidateRuntimeConfig for application/json ContentType.
type ValidateRuntimeConfigJSONRequestBody = RuntimeConfigChangeRequest

// CreateSystemIntegrationJSONRequestBody defines body for CreateSystemIntegration for application/json ContentType.
type CreateSystemIntegrationJSONRequestBody = SystemIntegrationCreateRequest

// UpdateSystemIntegrationJSONRequestBody defines body for UpdateSystemIntegration for application/json ContentType.
type UpdateSystemIntegrationJSONRequestBody = SystemIntegrationUpdateRequest

// CreateWorkflowTemplateJSONRequestBody defines body for CreateWorkflowTemplate for application/json ContentType.
type CreateWorkflowTemplateJSONRequestBody = WorkflowTemplateInput

// UpdateWorkflowTemplateJSONRequestBody defines body for UpdateWorkflowTemplate for application/json ContentType.
type UpdateWorkflowTemplateJSONRequestBody = WorkflowTemplateInput

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

// Getter for additional properties for AnthropicMessagesRequest. Returns the specified
// element and whether it was found
func (a AnthropicMessagesRequest) Get(fieldName string) (value interface{}, found bool) {
	if a.AdditionalProperties != nil {
		value, found = a.AdditionalProperties[fieldName]
	}
	return
}

// Setter for additional properties for AnthropicMessagesRequest
func (a *AnthropicMessagesRequest) Set(fieldName string, value interface{}) {
	if a.AdditionalProperties == nil {
		a.AdditionalProperties = make(map[string]any)
	}
	a.AdditionalProperties[fieldName] = value
}

// Override default JSON handling for AnthropicMessagesRequest to handle AdditionalProperties
func (a *AnthropicMessagesRequest) UnmarshalJSON(b []byte) error {
	object := make(map[string]json.RawMessage)
	err := json.Unmarshal(b, &object)
	if err != nil {
		return err
	}

	if raw, found := object["max_tokens"]; found {
		err = json.Unmarshal(raw, &a.MaxTokens)
		if err != nil {
			return fmt.Errorf("error reading 'max_tokens': %w", err)
		}
		delete(object, "max_tokens")
	}

	if raw, found := object["messages"]; found {
		err = json.Unmarshal(raw, &a.Messages)
		if err != nil {
			return fmt.Errorf("error reading 'messages': %w", err)
		}
		delete(object, "messages")
	}

	if raw, found := object["model"]; found {
		err = json.Unmarshal(raw, &a.Model)
		if err != nil {
			return fmt.Errorf("error reading 'model': %w", err)
		}
		delete(object, "model")
	}

	if raw, found := object["stream"]; found {
		err = json.Unmarshal(raw, &a.Stream)
		if err != nil {
			return fmt.Errorf("error reading 'stream': %w", err)
		}
		delete(object, "stream")
	}

	if raw, found := object["system"]; found {
		err = json.Unmarshal(raw, &a.System)
		if err != nil {
			return fmt.Errorf("error reading 'system': %w", err)
		}
		delete(object, "system")
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

// Override default JSON handling for AnthropicMessagesRequest to handle AdditionalProperties
func (a AnthropicMessagesRequest) MarshalJSON() ([]byte, error) {
	var err error
	object := make(map[string]json.RawMessage)

	object["max_tokens"], err = json.Marshal(a.MaxTokens)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'max_tokens': %w", err)
	}

	if a.Messages != nil {
		object["messages"], err = json.Marshal(a.Messages)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'messages': %w", err)
		}
	}

	object["model"], err = json.Marshal(a.Model)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'model': %w", err)
	}

	object["stream"], err = json.Marshal(a.Stream)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'stream': %w", err)
	}

	object["system"], err = json.Marshal(a.System)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'system': %w", err)
	}

	for fieldName, field := range a.AdditionalProperties {
		object[fieldName], err = json.Marshal(field)
		if err != nil {
			return nil, fmt.Errorf("error marshaling '%s': %w", fieldName, err)
		}
	}
	return json.Marshal(object)
}

// Getter for additional properties for AnthropicModel. Returns the specified
// element and whether it was found
func (a AnthropicModel) Get(fieldName string) (value interface{}, found bool) {
	if a.AdditionalProperties != nil {
		value, found = a.AdditionalProperties[fieldName]
	}
	return
}

// Setter for additional properties for AnthropicModel
func (a *AnthropicModel) Set(fieldName string, value interface{}) {
	if a.AdditionalProperties == nil {
		a.AdditionalProperties = make(map[string]any)
	}
	a.AdditionalProperties[fieldName] = value
}

// Override default JSON handling for AnthropicModel to handle AdditionalProperties
func (a *AnthropicModel) UnmarshalJSON(b []byte) error {
	object := make(map[string]json.RawMessage)
	err := json.Unmarshal(b, &object)
	if err != nil {
		return err
	}

	if raw, found := object["created_at"]; found {
		err = json.Unmarshal(raw, &a.CreatedAt)
		if err != nil {
			return fmt.Errorf("error reading 'created_at': %w", err)
		}
		delete(object, "created_at")
	}

	if raw, found := object["display_name"]; found {
		err = json.Unmarshal(raw, &a.DisplayName)
		if err != nil {
			return fmt.Errorf("error reading 'display_name': %w", err)
		}
		delete(object, "display_name")
	}

	if raw, found := object["id"]; found {
		err = json.Unmarshal(raw, &a.ID)
		if err != nil {
			return fmt.Errorf("error reading 'id': %w", err)
		}
		delete(object, "id")
	}

	if raw, found := object["type"]; found {
		err = json.Unmarshal(raw, &a.Type)
		if err != nil {
			return fmt.Errorf("error reading 'type': %w", err)
		}
		delete(object, "type")
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

// Override default JSON handling for AnthropicModel to handle AdditionalProperties
func (a AnthropicModel) MarshalJSON() ([]byte, error) {
	var err error
	object := make(map[string]json.RawMessage)

	object["created_at"], err = json.Marshal(a.CreatedAt)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'created_at': %w", err)
	}

	object["display_name"], err = json.Marshal(a.DisplayName)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'display_name': %w", err)
	}

	object["id"], err = json.Marshal(a.ID)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'id': %w", err)
	}

	object["type"], err = json.Marshal(a.Type)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'type': %w", err)
	}

	for fieldName, field := range a.AdditionalProperties {
		object[fieldName], err = json.Marshal(field)
		if err != nil {
			return nil, fmt.Errorf("error marshaling '%s': %w", fieldName, err)
		}
	}
	return json.Marshal(object)
}

// Getter for additional properties for AnthropicModelsResponse. Returns the specified
// element and whether it was found
func (a AnthropicModelsResponse) Get(fieldName string) (value interface{}, found bool) {
	if a.AdditionalProperties != nil {
		value, found = a.AdditionalProperties[fieldName]
	}
	return
}

// Setter for additional properties for AnthropicModelsResponse
func (a *AnthropicModelsResponse) Set(fieldName string, value interface{}) {
	if a.AdditionalProperties == nil {
		a.AdditionalProperties = make(map[string]any)
	}
	a.AdditionalProperties[fieldName] = value
}

// Override default JSON handling for AnthropicModelsResponse to handle AdditionalProperties
func (a *AnthropicModelsResponse) UnmarshalJSON(b []byte) error {
	object := make(map[string]json.RawMessage)
	err := json.Unmarshal(b, &object)
	if err != nil {
		return err
	}

	if raw, found := object["data"]; found {
		err = json.Unmarshal(raw, &a.Data)
		if err != nil {
			return fmt.Errorf("error reading 'data': %w", err)
		}
		delete(object, "data")
	}

	if raw, found := object["first_id"]; found {
		err = json.Unmarshal(raw, &a.FirstID)
		if err != nil {
			return fmt.Errorf("error reading 'first_id': %w", err)
		}
		delete(object, "first_id")
	}

	if raw, found := object["has_more"]; found {
		err = json.Unmarshal(raw, &a.HasMore)
		if err != nil {
			return fmt.Errorf("error reading 'has_more': %w", err)
		}
		delete(object, "has_more")
	}

	if raw, found := object["last_id"]; found {
		err = json.Unmarshal(raw, &a.LastID)
		if err != nil {
			return fmt.Errorf("error reading 'last_id': %w", err)
		}
		delete(object, "last_id")
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

// Override default JSON handling for AnthropicModelsResponse to handle AdditionalProperties
func (a AnthropicModelsResponse) MarshalJSON() ([]byte, error) {
	var err error
	object := make(map[string]json.RawMessage)

	if a.Data != nil {
		object["data"], err = json.Marshal(a.Data)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'data': %w", err)
		}
	}

	object["first_id"], err = json.Marshal(a.FirstID)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'first_id': %w", err)
	}

	object["has_more"], err = json.Marshal(a.HasMore)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'has_more': %w", err)
	}

	object["last_id"], err = json.Marshal(a.LastID)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'last_id': %w", err)
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

// Getter for additional properties for CohereRerankRequest. Returns the specified
// element and whether it was found
func (a CohereRerankRequest) Get(fieldName string) (value interface{}, found bool) {
	if a.AdditionalProperties != nil {
		value, found = a.AdditionalProperties[fieldName]
	}
	return
}

// Setter for additional properties for CohereRerankRequest
func (a *CohereRerankRequest) Set(fieldName string, value interface{}) {
	if a.AdditionalProperties == nil {
		a.AdditionalProperties = make(map[string]any)
	}
	a.AdditionalProperties[fieldName] = value
}

// Override default JSON handling for CohereRerankRequest to handle AdditionalProperties
func (a *CohereRerankRequest) UnmarshalJSON(b []byte) error {
	object := make(map[string]json.RawMessage)
	err := json.Unmarshal(b, &object)
	if err != nil {
		return err
	}

	if raw, found := object["documents"]; found {
		err = json.Unmarshal(raw, &a.Documents)
		if err != nil {
			return fmt.Errorf("error reading 'documents': %w", err)
		}
		delete(object, "documents")
	}

	if raw, found := object["max_tokens_per_doc"]; found {
		err = json.Unmarshal(raw, &a.MaxTokensPerDoc)
		if err != nil {
			return fmt.Errorf("error reading 'max_tokens_per_doc': %w", err)
		}
		delete(object, "max_tokens_per_doc")
	}

	if raw, found := object["model"]; found {
		err = json.Unmarshal(raw, &a.Model)
		if err != nil {
			return fmt.Errorf("error reading 'model': %w", err)
		}
		delete(object, "model")
	}

	if raw, found := object["query"]; found {
		err = json.Unmarshal(raw, &a.Query)
		if err != nil {
			return fmt.Errorf("error reading 'query': %w", err)
		}
		delete(object, "query")
	}

	if raw, found := object["rank_fields"]; found {
		err = json.Unmarshal(raw, &a.RankFields)
		if err != nil {
			return fmt.Errorf("error reading 'rank_fields': %w", err)
		}
		delete(object, "rank_fields")
	}

	if raw, found := object["top_n"]; found {
		err = json.Unmarshal(raw, &a.TopN)
		if err != nil {
			return fmt.Errorf("error reading 'top_n': %w", err)
		}
		delete(object, "top_n")
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

// Override default JSON handling for CohereRerankRequest to handle AdditionalProperties
func (a CohereRerankRequest) MarshalJSON() ([]byte, error) {
	var err error
	object := make(map[string]json.RawMessage)

	if a.Documents != nil {
		object["documents"], err = json.Marshal(a.Documents)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'documents': %w", err)
		}
	}

	object["max_tokens_per_doc"], err = json.Marshal(a.MaxTokensPerDoc)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'max_tokens_per_doc': %w", err)
	}

	object["model"], err = json.Marshal(a.Model)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'model': %w", err)
	}

	object["query"], err = json.Marshal(a.Query)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'query': %w", err)
	}

	if a.RankFields != nil {
		object["rank_fields"], err = json.Marshal(a.RankFields)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'rank_fields': %w", err)
		}
	}

	object["top_n"], err = json.Marshal(a.TopN)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'top_n': %w", err)
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

	if raw, found := object["artifacts"]; found {
		err = json.Unmarshal(raw, &a.Artifacts)
		if err != nil {
			return fmt.Errorf("error reading 'artifacts': %w", err)
		}
		delete(object, "artifacts")
	}

	if raw, found := object["attemptCount"]; found {
		err = json.Unmarshal(raw, &a.AttemptCount)
		if err != nil {
			return fmt.Errorf("error reading 'attemptCount': %w", err)
		}
		delete(object, "attemptCount")
	}

	if raw, found := object["callbackToken"]; found {
		err = json.Unmarshal(raw, &a.CallbackToken)
		if err != nil {
			return fmt.Errorf("error reading 'callbackToken': %w", err)
		}
		delete(object, "callbackToken")
	}

	if raw, found := object["claimedByAgentId"]; found {
		err = json.Unmarshal(raw, &a.ClaimedByAgentID)
		if err != nil {
			return fmt.Errorf("error reading 'claimedByAgentId': %w", err)
		}
		delete(object, "claimedByAgentId")
	}

	if raw, found := object["createdAt"]; found {
		err = json.Unmarshal(raw, &a.CreatedAt)
		if err != nil {
			return fmt.Errorf("error reading 'createdAt': %w", err)
		}
		delete(object, "createdAt")
	}

	if raw, found := object["finishedAt"]; found {
		err = json.Unmarshal(raw, &a.FinishedAt)
		if err != nil {
			return fmt.Errorf("error reading 'finishedAt': %w", err)
		}
		delete(object, "finishedAt")
	}

	if raw, found := object["id"]; found {
		err = json.Unmarshal(raw, &a.ID)
		if err != nil {
			return fmt.Errorf("error reading 'id': %w", err)
		}
		delete(object, "id")
	}

	if raw, found := object["lastHeartbeatAt"]; found {
		err = json.Unmarshal(raw, &a.LastHeartbeatAt)
		if err != nil {
			return fmt.Errorf("error reading 'lastHeartbeatAt': %w", err)
		}
		delete(object, "lastHeartbeatAt")
	}

	if raw, found := object["lastRuntimeSeenAt"]; found {
		err = json.Unmarshal(raw, &a.LastRuntimeSeenAt)
		if err != nil {
			return fmt.Errorf("error reading 'lastRuntimeSeenAt': %w", err)
		}
		delete(object, "lastRuntimeSeenAt")
	}

	if raw, found := object["lockKey"]; found {
		err = json.Unmarshal(raw, &a.LockKey)
		if err != nil {
			return fmt.Errorf("error reading 'lockKey': %w", err)
		}
		delete(object, "lockKey")
	}

	if raw, found := object["maxRetries"]; found {
		err = json.Unmarshal(raw, &a.MaxRetries)
		if err != nil {
			return fmt.Errorf("error reading 'maxRetries': %w", err)
		}
		delete(object, "maxRetries")
	}

	if raw, found := object["operationState"]; found {
		err = json.Unmarshal(raw, &a.OperationState)
		if err != nil {
			return fmt.Errorf("error reading 'operationState': %w", err)
		}
		delete(object, "operationState")
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

	if raw, found := object["queueKey"]; found {
		err = json.Unmarshal(raw, &a.QueueKey)
		if err != nil {
			return fmt.Errorf("error reading 'queueKey': %w", err)
		}
		delete(object, "queueKey")
	}

	if raw, found := object["releaseBundleId"]; found {
		err = json.Unmarshal(raw, &a.ReleaseBundleID)
		if err != nil {
			return fmt.Errorf("error reading 'releaseBundleId': %w", err)
		}
		delete(object, "releaseBundleId")
	}

	if raw, found := object["result"]; found {
		err = json.Unmarshal(raw, &a.Result)
		if err != nil {
			return fmt.Errorf("error reading 'result': %w", err)
		}
		delete(object, "result")
	}

	if raw, found := object["runtimeClusterId"]; found {
		err = json.Unmarshal(raw, &a.RuntimeClusterID)
		if err != nil {
			return fmt.Errorf("error reading 'runtimeClusterId': %w", err)
		}
		delete(object, "runtimeClusterId")
	}

	if raw, found := object["runtimeEndpoint"]; found {
		err = json.Unmarshal(raw, &a.RuntimeEndpoint)
		if err != nil {
			return fmt.Errorf("error reading 'runtimeEndpoint': %w", err)
		}
		delete(object, "runtimeEndpoint")
	}

	if raw, found := object["startedAt"]; found {
		err = json.Unmarshal(raw, &a.StartedAt)
		if err != nil {
			return fmt.Errorf("error reading 'startedAt': %w", err)
		}
		delete(object, "startedAt")
	}

	if raw, found := object["status"]; found {
		err = json.Unmarshal(raw, &a.Status)
		if err != nil {
			return fmt.Errorf("error reading 'status': %w", err)
		}
		delete(object, "status")
	}

	if raw, found := object["stopTransport"]; found {
		err = json.Unmarshal(raw, &a.StopTransport)
		if err != nil {
			return fmt.Errorf("error reading 'stopTransport': %w", err)
		}
		delete(object, "stopTransport")
	}

	if raw, found := object["targetKind"]; found {
		err = json.Unmarshal(raw, &a.TargetKind)
		if err != nil {
			return fmt.Errorf("error reading 'targetKind': %w", err)
		}
		delete(object, "targetKind")
	}

	if raw, found := object["taskKind"]; found {
		err = json.Unmarshal(raw, &a.TaskKind)
		if err != nil {
			return fmt.Errorf("error reading 'taskKind': %w", err)
		}
		delete(object, "taskKind")
	}

	if raw, found := object["timeoutSeconds"]; found {
		err = json.Unmarshal(raw, &a.TimeoutSeconds)
		if err != nil {
			return fmt.Errorf("error reading 'timeoutSeconds': %w", err)
		}
		delete(object, "timeoutSeconds")
	}

	if raw, found := object["updatedAt"]; found {
		err = json.Unmarshal(raw, &a.UpdatedAt)
		if err != nil {
			return fmt.Errorf("error reading 'updatedAt': %w", err)
		}
		delete(object, "updatedAt")
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

	if a.Artifacts != nil {
		object["artifacts"], err = json.Marshal(a.Artifacts)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'artifacts': %w", err)
		}
	}

	object["attemptCount"], err = json.Marshal(a.AttemptCount)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'attemptCount': %w", err)
	}

	object["callbackToken"], err = json.Marshal(a.CallbackToken)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'callbackToken': %w", err)
	}

	object["claimedByAgentId"], err = json.Marshal(a.ClaimedByAgentID)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'claimedByAgentId': %w", err)
	}

	if a.CreatedAt != nil {
		object["createdAt"], err = json.Marshal(a.CreatedAt)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'createdAt': %w", err)
		}
	}

	if a.FinishedAt != nil {
		object["finishedAt"], err = json.Marshal(a.FinishedAt)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'finishedAt': %w", err)
		}
	}

	object["id"], err = json.Marshal(a.ID)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'id': %w", err)
	}

	if a.LastHeartbeatAt != nil {
		object["lastHeartbeatAt"], err = json.Marshal(a.LastHeartbeatAt)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'lastHeartbeatAt': %w", err)
		}
	}

	if a.LastRuntimeSeenAt != nil {
		object["lastRuntimeSeenAt"], err = json.Marshal(a.LastRuntimeSeenAt)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'lastRuntimeSeenAt': %w", err)
		}
	}

	object["lockKey"], err = json.Marshal(a.LockKey)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'lockKey': %w", err)
	}

	object["maxRetries"], err = json.Marshal(a.MaxRetries)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'maxRetries': %w", err)
	}

	if a.OperationState != nil {
		object["operationState"], err = json.Marshal(a.OperationState)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'operationState': %w", err)
		}
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

	object["queueKey"], err = json.Marshal(a.QueueKey)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'queueKey': %w", err)
	}

	object["releaseBundleId"], err = json.Marshal(a.ReleaseBundleID)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'releaseBundleId': %w", err)
	}

	if a.Result != nil {
		object["result"], err = json.Marshal(a.Result)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'result': %w", err)
		}
	}

	object["runtimeClusterId"], err = json.Marshal(a.RuntimeClusterID)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'runtimeClusterId': %w", err)
	}

	object["runtimeEndpoint"], err = json.Marshal(a.RuntimeEndpoint)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'runtimeEndpoint': %w", err)
	}

	if a.StartedAt != nil {
		object["startedAt"], err = json.Marshal(a.StartedAt)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'startedAt': %w", err)
		}
	}

	object["status"], err = json.Marshal(a.Status)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'status': %w", err)
	}

	object["stopTransport"], err = json.Marshal(a.StopTransport)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'stopTransport': %w", err)
	}

	object["targetKind"], err = json.Marshal(a.TargetKind)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'targetKind': %w", err)
	}

	object["taskKind"], err = json.Marshal(a.TaskKind)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'taskKind': %w", err)
	}

	object["timeoutSeconds"], err = json.Marshal(a.TimeoutSeconds)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'timeoutSeconds': %w", err)
	}

	if a.UpdatedAt != nil {
		object["updatedAt"], err = json.Marshal(a.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'updatedAt': %w", err)
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

// Getter for additional properties for GeminiGenerateContentRequest. Returns the specified
// element and whether it was found
func (a GeminiGenerateContentRequest) Get(fieldName string) (value interface{}, found bool) {
	if a.AdditionalProperties != nil {
		value, found = a.AdditionalProperties[fieldName]
	}
	return
}

// Setter for additional properties for GeminiGenerateContentRequest
func (a *GeminiGenerateContentRequest) Set(fieldName string, value interface{}) {
	if a.AdditionalProperties == nil {
		a.AdditionalProperties = make(map[string]any)
	}
	a.AdditionalProperties[fieldName] = value
}

// Override default JSON handling for GeminiGenerateContentRequest to handle AdditionalProperties
func (a *GeminiGenerateContentRequest) UnmarshalJSON(b []byte) error {
	object := make(map[string]json.RawMessage)
	err := json.Unmarshal(b, &object)
	if err != nil {
		return err
	}

	if raw, found := object["contents"]; found {
		err = json.Unmarshal(raw, &a.Contents)
		if err != nil {
			return fmt.Errorf("error reading 'contents': %w", err)
		}
		delete(object, "contents")
	}

	if raw, found := object["generationConfig"]; found {
		err = json.Unmarshal(raw, &a.GenerationConfig)
		if err != nil {
			return fmt.Errorf("error reading 'generationConfig': %w", err)
		}
		delete(object, "generationConfig")
	}

	if raw, found := object["safetySettings"]; found {
		err = json.Unmarshal(raw, &a.SafetySettings)
		if err != nil {
			return fmt.Errorf("error reading 'safetySettings': %w", err)
		}
		delete(object, "safetySettings")
	}

	if raw, found := object["systemInstruction"]; found {
		err = json.Unmarshal(raw, &a.SystemInstruction)
		if err != nil {
			return fmt.Errorf("error reading 'systemInstruction': %w", err)
		}
		delete(object, "systemInstruction")
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

// Override default JSON handling for GeminiGenerateContentRequest to handle AdditionalProperties
func (a GeminiGenerateContentRequest) MarshalJSON() ([]byte, error) {
	var err error
	object := make(map[string]json.RawMessage)

	if a.Contents != nil {
		object["contents"], err = json.Marshal(a.Contents)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'contents': %w", err)
		}
	}

	if a.GenerationConfig != nil {
		object["generationConfig"], err = json.Marshal(a.GenerationConfig)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'generationConfig': %w", err)
		}
	}

	if a.SafetySettings != nil {
		object["safetySettings"], err = json.Marshal(a.SafetySettings)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'safetySettings': %w", err)
		}
	}

	if a.SystemInstruction != nil {
		object["systemInstruction"], err = json.Marshal(a.SystemInstruction)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'systemInstruction': %w", err)
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

// Getter for additional properties for GeminiInteractionsRequest. Returns the specified
// element and whether it was found
func (a GeminiInteractionsRequest) Get(fieldName string) (value interface{}, found bool) {
	if a.AdditionalProperties != nil {
		value, found = a.AdditionalProperties[fieldName]
	}
	return
}

// Setter for additional properties for GeminiInteractionsRequest
func (a *GeminiInteractionsRequest) Set(fieldName string, value interface{}) {
	if a.AdditionalProperties == nil {
		a.AdditionalProperties = make(map[string]any)
	}
	a.AdditionalProperties[fieldName] = value
}

// Override default JSON handling for GeminiInteractionsRequest to handle AdditionalProperties
func (a *GeminiInteractionsRequest) UnmarshalJSON(b []byte) error {
	object := make(map[string]json.RawMessage)
	err := json.Unmarshal(b, &object)
	if err != nil {
		return err
	}

	if raw, found := object["background"]; found {
		err = json.Unmarshal(raw, &a.Background)
		if err != nil {
			return fmt.Errorf("error reading 'background': %w", err)
		}
		delete(object, "background")
	}

	if raw, found := object["input"]; found {
		err = json.Unmarshal(raw, &a.Input)
		if err != nil {
			return fmt.Errorf("error reading 'input': %w", err)
		}
		delete(object, "input")
	}

	if raw, found := object["model"]; found {
		err = json.Unmarshal(raw, &a.Model)
		if err != nil {
			return fmt.Errorf("error reading 'model': %w", err)
		}
		delete(object, "model")
	}

	if raw, found := object["stream"]; found {
		err = json.Unmarshal(raw, &a.Stream)
		if err != nil {
			return fmt.Errorf("error reading 'stream': %w", err)
		}
		delete(object, "stream")
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

// Override default JSON handling for GeminiInteractionsRequest to handle AdditionalProperties
func (a GeminiInteractionsRequest) MarshalJSON() ([]byte, error) {
	var err error
	object := make(map[string]json.RawMessage)

	object["background"], err = json.Marshal(a.Background)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'background': %w", err)
	}

	object["input"], err = json.Marshal(a.Input)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'input': %w", err)
	}

	object["model"], err = json.Marshal(a.Model)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'model': %w", err)
	}

	object["stream"], err = json.Marshal(a.Stream)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'stream': %w", err)
	}

	for fieldName, field := range a.AdditionalProperties {
		object[fieldName], err = json.Marshal(field)
		if err != nil {
			return nil, fmt.Errorf("error marshaling '%s': %w", fieldName, err)
		}
	}
	return json.Marshal(object)
}

// Getter for additional properties for GeminiModel. Returns the specified
// element and whether it was found
func (a GeminiModel) Get(fieldName string) (value interface{}, found bool) {
	if a.AdditionalProperties != nil {
		value, found = a.AdditionalProperties[fieldName]
	}
	return
}

// Setter for additional properties for GeminiModel
func (a *GeminiModel) Set(fieldName string, value interface{}) {
	if a.AdditionalProperties == nil {
		a.AdditionalProperties = make(map[string]any)
	}
	a.AdditionalProperties[fieldName] = value
}

// Override default JSON handling for GeminiModel to handle AdditionalProperties
func (a *GeminiModel) UnmarshalJSON(b []byte) error {
	object := make(map[string]json.RawMessage)
	err := json.Unmarshal(b, &object)
	if err != nil {
		return err
	}

	if raw, found := object["description"]; found {
		err = json.Unmarshal(raw, &a.Description)
		if err != nil {
			return fmt.Errorf("error reading 'description': %w", err)
		}
		delete(object, "description")
	}

	if raw, found := object["displayName"]; found {
		err = json.Unmarshal(raw, &a.DisplayName)
		if err != nil {
			return fmt.Errorf("error reading 'displayName': %w", err)
		}
		delete(object, "displayName")
	}

	if raw, found := object["name"]; found {
		err = json.Unmarshal(raw, &a.Name)
		if err != nil {
			return fmt.Errorf("error reading 'name': %w", err)
		}
		delete(object, "name")
	}

	if raw, found := object["supportedGenerationMethods"]; found {
		err = json.Unmarshal(raw, &a.SupportedGenerationMethods)
		if err != nil {
			return fmt.Errorf("error reading 'supportedGenerationMethods': %w", err)
		}
		delete(object, "supportedGenerationMethods")
	}

	if raw, found := object["version"]; found {
		err = json.Unmarshal(raw, &a.Version)
		if err != nil {
			return fmt.Errorf("error reading 'version': %w", err)
		}
		delete(object, "version")
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

// Override default JSON handling for GeminiModel to handle AdditionalProperties
func (a GeminiModel) MarshalJSON() ([]byte, error) {
	var err error
	object := make(map[string]json.RawMessage)

	object["description"], err = json.Marshal(a.Description)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'description': %w", err)
	}

	object["displayName"], err = json.Marshal(a.DisplayName)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'displayName': %w", err)
	}

	object["name"], err = json.Marshal(a.Name)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'name': %w", err)
	}

	if a.SupportedGenerationMethods != nil {
		object["supportedGenerationMethods"], err = json.Marshal(a.SupportedGenerationMethods)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'supportedGenerationMethods': %w", err)
		}
	}

	object["version"], err = json.Marshal(a.Version)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'version': %w", err)
	}

	for fieldName, field := range a.AdditionalProperties {
		object[fieldName], err = json.Marshal(field)
		if err != nil {
			return nil, fmt.Errorf("error marshaling '%s': %w", fieldName, err)
		}
	}
	return json.Marshal(object)
}

// Getter for additional properties for GeminiModelsResponse. Returns the specified
// element and whether it was found
func (a GeminiModelsResponse) Get(fieldName string) (value interface{}, found bool) {
	if a.AdditionalProperties != nil {
		value, found = a.AdditionalProperties[fieldName]
	}
	return
}

// Setter for additional properties for GeminiModelsResponse
func (a *GeminiModelsResponse) Set(fieldName string, value interface{}) {
	if a.AdditionalProperties == nil {
		a.AdditionalProperties = make(map[string]any)
	}
	a.AdditionalProperties[fieldName] = value
}

// Override default JSON handling for GeminiModelsResponse to handle AdditionalProperties
func (a *GeminiModelsResponse) UnmarshalJSON(b []byte) error {
	object := make(map[string]json.RawMessage)
	err := json.Unmarshal(b, &object)
	if err != nil {
		return err
	}

	if raw, found := object["models"]; found {
		err = json.Unmarshal(raw, &a.Models)
		if err != nil {
			return fmt.Errorf("error reading 'models': %w", err)
		}
		delete(object, "models")
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

// Override default JSON handling for GeminiModelsResponse to handle AdditionalProperties
func (a GeminiModelsResponse) MarshalJSON() ([]byte, error) {
	var err error
	object := make(map[string]json.RawMessage)

	if a.Models != nil {
		object["models"], err = json.Marshal(a.Models)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'models': %w", err)
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

// Getter for additional properties for OpenAIAudioSpeechRequest. Returns the specified
// element and whether it was found
func (a OpenAIAudioSpeechRequest) Get(fieldName string) (value interface{}, found bool) {
	if a.AdditionalProperties != nil {
		value, found = a.AdditionalProperties[fieldName]
	}
	return
}

// Setter for additional properties for OpenAIAudioSpeechRequest
func (a *OpenAIAudioSpeechRequest) Set(fieldName string, value interface{}) {
	if a.AdditionalProperties == nil {
		a.AdditionalProperties = make(map[string]any)
	}
	a.AdditionalProperties[fieldName] = value
}

// Override default JSON handling for OpenAIAudioSpeechRequest to handle AdditionalProperties
func (a *OpenAIAudioSpeechRequest) UnmarshalJSON(b []byte) error {
	object := make(map[string]json.RawMessage)
	err := json.Unmarshal(b, &object)
	if err != nil {
		return err
	}

	if raw, found := object["input"]; found {
		err = json.Unmarshal(raw, &a.Input)
		if err != nil {
			return fmt.Errorf("error reading 'input': %w", err)
		}
		delete(object, "input")
	}

	if raw, found := object["instructions"]; found {
		err = json.Unmarshal(raw, &a.Instructions)
		if err != nil {
			return fmt.Errorf("error reading 'instructions': %w", err)
		}
		delete(object, "instructions")
	}

	if raw, found := object["model"]; found {
		err = json.Unmarshal(raw, &a.Model)
		if err != nil {
			return fmt.Errorf("error reading 'model': %w", err)
		}
		delete(object, "model")
	}

	if raw, found := object["response_format"]; found {
		err = json.Unmarshal(raw, &a.ResponseFormat)
		if err != nil {
			return fmt.Errorf("error reading 'response_format': %w", err)
		}
		delete(object, "response_format")
	}

	if raw, found := object["speed"]; found {
		err = json.Unmarshal(raw, &a.Speed)
		if err != nil {
			return fmt.Errorf("error reading 'speed': %w", err)
		}
		delete(object, "speed")
	}

	if raw, found := object["user"]; found {
		err = json.Unmarshal(raw, &a.User)
		if err != nil {
			return fmt.Errorf("error reading 'user': %w", err)
		}
		delete(object, "user")
	}

	if raw, found := object["voice"]; found {
		err = json.Unmarshal(raw, &a.Voice)
		if err != nil {
			return fmt.Errorf("error reading 'voice': %w", err)
		}
		delete(object, "voice")
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

// Override default JSON handling for OpenAIAudioSpeechRequest to handle AdditionalProperties
func (a OpenAIAudioSpeechRequest) MarshalJSON() ([]byte, error) {
	var err error
	object := make(map[string]json.RawMessage)

	object["input"], err = json.Marshal(a.Input)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'input': %w", err)
	}

	object["instructions"], err = json.Marshal(a.Instructions)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'instructions': %w", err)
	}

	object["model"], err = json.Marshal(a.Model)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'model': %w", err)
	}

	object["response_format"], err = json.Marshal(a.ResponseFormat)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'response_format': %w", err)
	}

	object["speed"], err = json.Marshal(a.Speed)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'speed': %w", err)
	}

	object["user"], err = json.Marshal(a.User)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'user': %w", err)
	}

	object["voice"], err = json.Marshal(a.Voice)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'voice': %w", err)
	}

	for fieldName, field := range a.AdditionalProperties {
		object[fieldName], err = json.Marshal(field)
		if err != nil {
			return nil, fmt.Errorf("error marshaling '%s': %w", fieldName, err)
		}
	}
	return json.Marshal(object)
}

// Getter for additional properties for OpenAIAudioTranscriptionRequest. Returns the specified
// element and whether it was found
func (a OpenAIAudioTranscriptionRequest) Get(fieldName string) (value interface{}, found bool) {
	if a.AdditionalProperties != nil {
		value, found = a.AdditionalProperties[fieldName]
	}
	return
}

// Setter for additional properties for OpenAIAudioTranscriptionRequest
func (a *OpenAIAudioTranscriptionRequest) Set(fieldName string, value interface{}) {
	if a.AdditionalProperties == nil {
		a.AdditionalProperties = make(map[string]any)
	}
	a.AdditionalProperties[fieldName] = value
}

// Override default JSON handling for OpenAIAudioTranscriptionRequest to handle AdditionalProperties
func (a *OpenAIAudioTranscriptionRequest) UnmarshalJSON(b []byte) error {
	object := make(map[string]json.RawMessage)
	err := json.Unmarshal(b, &object)
	if err != nil {
		return err
	}

	if raw, found := object["file"]; found {
		err = json.Unmarshal(raw, &a.File)
		if err != nil {
			return fmt.Errorf("error reading 'file': %w", err)
		}
		delete(object, "file")
	}

	if raw, found := object["include"]; found {
		err = json.Unmarshal(raw, &a.Include)
		if err != nil {
			return fmt.Errorf("error reading 'include': %w", err)
		}
		delete(object, "include")
	}

	if raw, found := object["language"]; found {
		err = json.Unmarshal(raw, &a.Language)
		if err != nil {
			return fmt.Errorf("error reading 'language': %w", err)
		}
		delete(object, "language")
	}

	if raw, found := object["model"]; found {
		err = json.Unmarshal(raw, &a.Model)
		if err != nil {
			return fmt.Errorf("error reading 'model': %w", err)
		}
		delete(object, "model")
	}

	if raw, found := object["prompt"]; found {
		err = json.Unmarshal(raw, &a.Prompt)
		if err != nil {
			return fmt.Errorf("error reading 'prompt': %w", err)
		}
		delete(object, "prompt")
	}

	if raw, found := object["response_format"]; found {
		err = json.Unmarshal(raw, &a.ResponseFormat)
		if err != nil {
			return fmt.Errorf("error reading 'response_format': %w", err)
		}
		delete(object, "response_format")
	}

	if raw, found := object["temperature"]; found {
		err = json.Unmarshal(raw, &a.Temperature)
		if err != nil {
			return fmt.Errorf("error reading 'temperature': %w", err)
		}
		delete(object, "temperature")
	}

	if raw, found := object["timestamp_granularities"]; found {
		err = json.Unmarshal(raw, &a.TimestampGranularities)
		if err != nil {
			return fmt.Errorf("error reading 'timestamp_granularities': %w", err)
		}
		delete(object, "timestamp_granularities")
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

// Override default JSON handling for OpenAIAudioTranscriptionRequest to handle AdditionalProperties
func (a OpenAIAudioTranscriptionRequest) MarshalJSON() ([]byte, error) {
	var err error
	object := make(map[string]json.RawMessage)

	object["file"], err = json.Marshal(a.File)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'file': %w", err)
	}

	if a.Include != nil {
		object["include"], err = json.Marshal(a.Include)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'include': %w", err)
		}
	}

	object["language"], err = json.Marshal(a.Language)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'language': %w", err)
	}

	object["model"], err = json.Marshal(a.Model)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'model': %w", err)
	}

	object["prompt"], err = json.Marshal(a.Prompt)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'prompt': %w", err)
	}

	object["response_format"], err = json.Marshal(a.ResponseFormat)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'response_format': %w", err)
	}

	object["temperature"], err = json.Marshal(a.Temperature)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'temperature': %w", err)
	}

	if a.TimestampGranularities != nil {
		object["timestamp_granularities"], err = json.Marshal(a.TimestampGranularities)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'timestamp_granularities': %w", err)
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

// Getter for additional properties for OpenAIAudioTranslationRequest. Returns the specified
// element and whether it was found
func (a OpenAIAudioTranslationRequest) Get(fieldName string) (value interface{}, found bool) {
	if a.AdditionalProperties != nil {
		value, found = a.AdditionalProperties[fieldName]
	}
	return
}

// Setter for additional properties for OpenAIAudioTranslationRequest
func (a *OpenAIAudioTranslationRequest) Set(fieldName string, value interface{}) {
	if a.AdditionalProperties == nil {
		a.AdditionalProperties = make(map[string]any)
	}
	a.AdditionalProperties[fieldName] = value
}

// Override default JSON handling for OpenAIAudioTranslationRequest to handle AdditionalProperties
func (a *OpenAIAudioTranslationRequest) UnmarshalJSON(b []byte) error {
	object := make(map[string]json.RawMessage)
	err := json.Unmarshal(b, &object)
	if err != nil {
		return err
	}

	if raw, found := object["file"]; found {
		err = json.Unmarshal(raw, &a.File)
		if err != nil {
			return fmt.Errorf("error reading 'file': %w", err)
		}
		delete(object, "file")
	}

	if raw, found := object["model"]; found {
		err = json.Unmarshal(raw, &a.Model)
		if err != nil {
			return fmt.Errorf("error reading 'model': %w", err)
		}
		delete(object, "model")
	}

	if raw, found := object["prompt"]; found {
		err = json.Unmarshal(raw, &a.Prompt)
		if err != nil {
			return fmt.Errorf("error reading 'prompt': %w", err)
		}
		delete(object, "prompt")
	}

	if raw, found := object["response_format"]; found {
		err = json.Unmarshal(raw, &a.ResponseFormat)
		if err != nil {
			return fmt.Errorf("error reading 'response_format': %w", err)
		}
		delete(object, "response_format")
	}

	if raw, found := object["temperature"]; found {
		err = json.Unmarshal(raw, &a.Temperature)
		if err != nil {
			return fmt.Errorf("error reading 'temperature': %w", err)
		}
		delete(object, "temperature")
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

// Override default JSON handling for OpenAIAudioTranslationRequest to handle AdditionalProperties
func (a OpenAIAudioTranslationRequest) MarshalJSON() ([]byte, error) {
	var err error
	object := make(map[string]json.RawMessage)

	object["file"], err = json.Marshal(a.File)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'file': %w", err)
	}

	object["model"], err = json.Marshal(a.Model)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'model': %w", err)
	}

	object["prompt"], err = json.Marshal(a.Prompt)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'prompt': %w", err)
	}

	object["response_format"], err = json.Marshal(a.ResponseFormat)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'response_format': %w", err)
	}

	object["temperature"], err = json.Marshal(a.Temperature)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'temperature': %w", err)
	}

	for fieldName, field := range a.AdditionalProperties {
		object[fieldName], err = json.Marshal(field)
		if err != nil {
			return nil, fmt.Errorf("error marshaling '%s': %w", fieldName, err)
		}
	}
	return json.Marshal(object)
}

// Getter for additional properties for OpenAIChatCompletionRequest. Returns the specified
// element and whether it was found
func (a OpenAIChatCompletionRequest) Get(fieldName string) (value interface{}, found bool) {
	if a.AdditionalProperties != nil {
		value, found = a.AdditionalProperties[fieldName]
	}
	return
}

// Setter for additional properties for OpenAIChatCompletionRequest
func (a *OpenAIChatCompletionRequest) Set(fieldName string, value interface{}) {
	if a.AdditionalProperties == nil {
		a.AdditionalProperties = make(map[string]any)
	}
	a.AdditionalProperties[fieldName] = value
}

// Override default JSON handling for OpenAIChatCompletionRequest to handle AdditionalProperties
func (a *OpenAIChatCompletionRequest) UnmarshalJSON(b []byte) error {
	object := make(map[string]json.RawMessage)
	err := json.Unmarshal(b, &object)
	if err != nil {
		return err
	}

	if raw, found := object["messages"]; found {
		err = json.Unmarshal(raw, &a.Messages)
		if err != nil {
			return fmt.Errorf("error reading 'messages': %w", err)
		}
		delete(object, "messages")
	}

	if raw, found := object["model"]; found {
		err = json.Unmarshal(raw, &a.Model)
		if err != nil {
			return fmt.Errorf("error reading 'model': %w", err)
		}
		delete(object, "model")
	}

	if raw, found := object["stream"]; found {
		err = json.Unmarshal(raw, &a.Stream)
		if err != nil {
			return fmt.Errorf("error reading 'stream': %w", err)
		}
		delete(object, "stream")
	}

	if raw, found := object["stream_options"]; found {
		err = json.Unmarshal(raw, &a.StreamOptions)
		if err != nil {
			return fmt.Errorf("error reading 'stream_options': %w", err)
		}
		delete(object, "stream_options")
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

// Override default JSON handling for OpenAIChatCompletionRequest to handle AdditionalProperties
func (a OpenAIChatCompletionRequest) MarshalJSON() ([]byte, error) {
	var err error
	object := make(map[string]json.RawMessage)

	if a.Messages != nil {
		object["messages"], err = json.Marshal(a.Messages)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'messages': %w", err)
		}
	}

	object["model"], err = json.Marshal(a.Model)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'model': %w", err)
	}

	object["stream"], err = json.Marshal(a.Stream)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'stream': %w", err)
	}

	if a.StreamOptions != nil {
		object["stream_options"], err = json.Marshal(a.StreamOptions)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'stream_options': %w", err)
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

// Getter for additional properties for OpenAIEmbeddingRequest. Returns the specified
// element and whether it was found
func (a OpenAIEmbeddingRequest) Get(fieldName string) (value interface{}, found bool) {
	if a.AdditionalProperties != nil {
		value, found = a.AdditionalProperties[fieldName]
	}
	return
}

// Setter for additional properties for OpenAIEmbeddingRequest
func (a *OpenAIEmbeddingRequest) Set(fieldName string, value interface{}) {
	if a.AdditionalProperties == nil {
		a.AdditionalProperties = make(map[string]any)
	}
	a.AdditionalProperties[fieldName] = value
}

// Override default JSON handling for OpenAIEmbeddingRequest to handle AdditionalProperties
func (a *OpenAIEmbeddingRequest) UnmarshalJSON(b []byte) error {
	object := make(map[string]json.RawMessage)
	err := json.Unmarshal(b, &object)
	if err != nil {
		return err
	}

	if raw, found := object["input"]; found {
		err = json.Unmarshal(raw, &a.Input)
		if err != nil {
			return fmt.Errorf("error reading 'input': %w", err)
		}
		delete(object, "input")
	}

	if raw, found := object["model"]; found {
		err = json.Unmarshal(raw, &a.Model)
		if err != nil {
			return fmt.Errorf("error reading 'model': %w", err)
		}
		delete(object, "model")
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

// Override default JSON handling for OpenAIEmbeddingRequest to handle AdditionalProperties
func (a OpenAIEmbeddingRequest) MarshalJSON() ([]byte, error) {
	var err error
	object := make(map[string]json.RawMessage)

	object["input"], err = json.Marshal(a.Input)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'input': %w", err)
	}

	object["model"], err = json.Marshal(a.Model)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'model': %w", err)
	}

	for fieldName, field := range a.AdditionalProperties {
		object[fieldName], err = json.Marshal(field)
		if err != nil {
			return nil, fmt.Errorf("error marshaling '%s': %w", fieldName, err)
		}
	}
	return json.Marshal(object)
}

// Getter for additional properties for OpenAIImageEditRequest. Returns the specified
// element and whether it was found
func (a OpenAIImageEditRequest) Get(fieldName string) (value interface{}, found bool) {
	if a.AdditionalProperties != nil {
		value, found = a.AdditionalProperties[fieldName]
	}
	return
}

// Setter for additional properties for OpenAIImageEditRequest
func (a *OpenAIImageEditRequest) Set(fieldName string, value interface{}) {
	if a.AdditionalProperties == nil {
		a.AdditionalProperties = make(map[string]any)
	}
	a.AdditionalProperties[fieldName] = value
}

// Override default JSON handling for OpenAIImageEditRequest to handle AdditionalProperties
func (a *OpenAIImageEditRequest) UnmarshalJSON(b []byte) error {
	object := make(map[string]json.RawMessage)
	err := json.Unmarshal(b, &object)
	if err != nil {
		return err
	}

	if raw, found := object["image"]; found {
		err = json.Unmarshal(raw, &a.Image)
		if err != nil {
			return fmt.Errorf("error reading 'image': %w", err)
		}
		delete(object, "image")
	}

	if raw, found := object["mask"]; found {
		err = json.Unmarshal(raw, &a.Mask)
		if err != nil {
			return fmt.Errorf("error reading 'mask': %w", err)
		}
		delete(object, "mask")
	}

	if raw, found := object["model"]; found {
		err = json.Unmarshal(raw, &a.Model)
		if err != nil {
			return fmt.Errorf("error reading 'model': %w", err)
		}
		delete(object, "model")
	}

	if raw, found := object["n"]; found {
		err = json.Unmarshal(raw, &a.N)
		if err != nil {
			return fmt.Errorf("error reading 'n': %w", err)
		}
		delete(object, "n")
	}

	if raw, found := object["prompt"]; found {
		err = json.Unmarshal(raw, &a.Prompt)
		if err != nil {
			return fmt.Errorf("error reading 'prompt': %w", err)
		}
		delete(object, "prompt")
	}

	if raw, found := object["quality"]; found {
		err = json.Unmarshal(raw, &a.Quality)
		if err != nil {
			return fmt.Errorf("error reading 'quality': %w", err)
		}
		delete(object, "quality")
	}

	if raw, found := object["response_format"]; found {
		err = json.Unmarshal(raw, &a.ResponseFormat)
		if err != nil {
			return fmt.Errorf("error reading 'response_format': %w", err)
		}
		delete(object, "response_format")
	}

	if raw, found := object["size"]; found {
		err = json.Unmarshal(raw, &a.Size)
		if err != nil {
			return fmt.Errorf("error reading 'size': %w", err)
		}
		delete(object, "size")
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

// Override default JSON handling for OpenAIImageEditRequest to handle AdditionalProperties
func (a OpenAIImageEditRequest) MarshalJSON() ([]byte, error) {
	var err error
	object := make(map[string]json.RawMessage)

	object["image"], err = json.Marshal(a.Image)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'image': %w", err)
	}

	object["mask"], err = json.Marshal(a.Mask)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'mask': %w", err)
	}

	object["model"], err = json.Marshal(a.Model)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'model': %w", err)
	}

	object["n"], err = json.Marshal(a.N)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'n': %w", err)
	}

	object["prompt"], err = json.Marshal(a.Prompt)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'prompt': %w", err)
	}

	object["quality"], err = json.Marshal(a.Quality)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'quality': %w", err)
	}

	object["response_format"], err = json.Marshal(a.ResponseFormat)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'response_format': %w", err)
	}

	object["size"], err = json.Marshal(a.Size)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'size': %w", err)
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

// Getter for additional properties for OpenAIImageGenerationRequest. Returns the specified
// element and whether it was found
func (a OpenAIImageGenerationRequest) Get(fieldName string) (value interface{}, found bool) {
	if a.AdditionalProperties != nil {
		value, found = a.AdditionalProperties[fieldName]
	}
	return
}

// Setter for additional properties for OpenAIImageGenerationRequest
func (a *OpenAIImageGenerationRequest) Set(fieldName string, value interface{}) {
	if a.AdditionalProperties == nil {
		a.AdditionalProperties = make(map[string]any)
	}
	a.AdditionalProperties[fieldName] = value
}

// Override default JSON handling for OpenAIImageGenerationRequest to handle AdditionalProperties
func (a *OpenAIImageGenerationRequest) UnmarshalJSON(b []byte) error {
	object := make(map[string]json.RawMessage)
	err := json.Unmarshal(b, &object)
	if err != nil {
		return err
	}

	if raw, found := object["background"]; found {
		err = json.Unmarshal(raw, &a.Background)
		if err != nil {
			return fmt.Errorf("error reading 'background': %w", err)
		}
		delete(object, "background")
	}

	if raw, found := object["model"]; found {
		err = json.Unmarshal(raw, &a.Model)
		if err != nil {
			return fmt.Errorf("error reading 'model': %w", err)
		}
		delete(object, "model")
	}

	if raw, found := object["moderation"]; found {
		err = json.Unmarshal(raw, &a.Moderation)
		if err != nil {
			return fmt.Errorf("error reading 'moderation': %w", err)
		}
		delete(object, "moderation")
	}

	if raw, found := object["n"]; found {
		err = json.Unmarshal(raw, &a.N)
		if err != nil {
			return fmt.Errorf("error reading 'n': %w", err)
		}
		delete(object, "n")
	}

	if raw, found := object["output_compression"]; found {
		err = json.Unmarshal(raw, &a.OutputCompression)
		if err != nil {
			return fmt.Errorf("error reading 'output_compression': %w", err)
		}
		delete(object, "output_compression")
	}

	if raw, found := object["output_format"]; found {
		err = json.Unmarshal(raw, &a.OutputFormat)
		if err != nil {
			return fmt.Errorf("error reading 'output_format': %w", err)
		}
		delete(object, "output_format")
	}

	if raw, found := object["prompt"]; found {
		err = json.Unmarshal(raw, &a.Prompt)
		if err != nil {
			return fmt.Errorf("error reading 'prompt': %w", err)
		}
		delete(object, "prompt")
	}

	if raw, found := object["quality"]; found {
		err = json.Unmarshal(raw, &a.Quality)
		if err != nil {
			return fmt.Errorf("error reading 'quality': %w", err)
		}
		delete(object, "quality")
	}

	if raw, found := object["response_format"]; found {
		err = json.Unmarshal(raw, &a.ResponseFormat)
		if err != nil {
			return fmt.Errorf("error reading 'response_format': %w", err)
		}
		delete(object, "response_format")
	}

	if raw, found := object["size"]; found {
		err = json.Unmarshal(raw, &a.Size)
		if err != nil {
			return fmt.Errorf("error reading 'size': %w", err)
		}
		delete(object, "size")
	}

	if raw, found := object["style"]; found {
		err = json.Unmarshal(raw, &a.Style)
		if err != nil {
			return fmt.Errorf("error reading 'style': %w", err)
		}
		delete(object, "style")
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

// Override default JSON handling for OpenAIImageGenerationRequest to handle AdditionalProperties
func (a OpenAIImageGenerationRequest) MarshalJSON() ([]byte, error) {
	var err error
	object := make(map[string]json.RawMessage)

	object["background"], err = json.Marshal(a.Background)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'background': %w", err)
	}

	object["model"], err = json.Marshal(a.Model)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'model': %w", err)
	}

	object["moderation"], err = json.Marshal(a.Moderation)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'moderation': %w", err)
	}

	object["n"], err = json.Marshal(a.N)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'n': %w", err)
	}

	object["output_compression"], err = json.Marshal(a.OutputCompression)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'output_compression': %w", err)
	}

	object["output_format"], err = json.Marshal(a.OutputFormat)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'output_format': %w", err)
	}

	object["prompt"], err = json.Marshal(a.Prompt)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'prompt': %w", err)
	}

	object["quality"], err = json.Marshal(a.Quality)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'quality': %w", err)
	}

	object["response_format"], err = json.Marshal(a.ResponseFormat)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'response_format': %w", err)
	}

	object["size"], err = json.Marshal(a.Size)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'size': %w", err)
	}

	object["style"], err = json.Marshal(a.Style)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'style': %w", err)
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

// Getter for additional properties for OpenAIImageVariationRequest. Returns the specified
// element and whether it was found
func (a OpenAIImageVariationRequest) Get(fieldName string) (value interface{}, found bool) {
	if a.AdditionalProperties != nil {
		value, found = a.AdditionalProperties[fieldName]
	}
	return
}

// Setter for additional properties for OpenAIImageVariationRequest
func (a *OpenAIImageVariationRequest) Set(fieldName string, value interface{}) {
	if a.AdditionalProperties == nil {
		a.AdditionalProperties = make(map[string]any)
	}
	a.AdditionalProperties[fieldName] = value
}

// Override default JSON handling for OpenAIImageVariationRequest to handle AdditionalProperties
func (a *OpenAIImageVariationRequest) UnmarshalJSON(b []byte) error {
	object := make(map[string]json.RawMessage)
	err := json.Unmarshal(b, &object)
	if err != nil {
		return err
	}

	if raw, found := object["image"]; found {
		err = json.Unmarshal(raw, &a.Image)
		if err != nil {
			return fmt.Errorf("error reading 'image': %w", err)
		}
		delete(object, "image")
	}

	if raw, found := object["model"]; found {
		err = json.Unmarshal(raw, &a.Model)
		if err != nil {
			return fmt.Errorf("error reading 'model': %w", err)
		}
		delete(object, "model")
	}

	if raw, found := object["n"]; found {
		err = json.Unmarshal(raw, &a.N)
		if err != nil {
			return fmt.Errorf("error reading 'n': %w", err)
		}
		delete(object, "n")
	}

	if raw, found := object["response_format"]; found {
		err = json.Unmarshal(raw, &a.ResponseFormat)
		if err != nil {
			return fmt.Errorf("error reading 'response_format': %w", err)
		}
		delete(object, "response_format")
	}

	if raw, found := object["size"]; found {
		err = json.Unmarshal(raw, &a.Size)
		if err != nil {
			return fmt.Errorf("error reading 'size': %w", err)
		}
		delete(object, "size")
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

// Override default JSON handling for OpenAIImageVariationRequest to handle AdditionalProperties
func (a OpenAIImageVariationRequest) MarshalJSON() ([]byte, error) {
	var err error
	object := make(map[string]json.RawMessage)

	object["image"], err = json.Marshal(a.Image)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'image': %w", err)
	}

	object["model"], err = json.Marshal(a.Model)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'model': %w", err)
	}

	object["n"], err = json.Marshal(a.N)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'n': %w", err)
	}

	object["response_format"], err = json.Marshal(a.ResponseFormat)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'response_format': %w", err)
	}

	object["size"], err = json.Marshal(a.Size)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'size': %w", err)
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

// Getter for additional properties for OpenAIModel. Returns the specified
// element and whether it was found
func (a OpenAIModel) Get(fieldName string) (value interface{}, found bool) {
	if a.AdditionalProperties != nil {
		value, found = a.AdditionalProperties[fieldName]
	}
	return
}

// Setter for additional properties for OpenAIModel
func (a *OpenAIModel) Set(fieldName string, value interface{}) {
	if a.AdditionalProperties == nil {
		a.AdditionalProperties = make(map[string]any)
	}
	a.AdditionalProperties[fieldName] = value
}

// Override default JSON handling for OpenAIModel to handle AdditionalProperties
func (a *OpenAIModel) UnmarshalJSON(b []byte) error {
	object := make(map[string]json.RawMessage)
	err := json.Unmarshal(b, &object)
	if err != nil {
		return err
	}

	if raw, found := object["created"]; found {
		err = json.Unmarshal(raw, &a.Created)
		if err != nil {
			return fmt.Errorf("error reading 'created': %w", err)
		}
		delete(object, "created")
	}

	if raw, found := object["id"]; found {
		err = json.Unmarshal(raw, &a.ID)
		if err != nil {
			return fmt.Errorf("error reading 'id': %w", err)
		}
		delete(object, "id")
	}

	if raw, found := object["object"]; found {
		err = json.Unmarshal(raw, &a.Object)
		if err != nil {
			return fmt.Errorf("error reading 'object': %w", err)
		}
		delete(object, "object")
	}

	if raw, found := object["owned_by"]; found {
		err = json.Unmarshal(raw, &a.OwnedBy)
		if err != nil {
			return fmt.Errorf("error reading 'owned_by': %w", err)
		}
		delete(object, "owned_by")
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

// Override default JSON handling for OpenAIModel to handle AdditionalProperties
func (a OpenAIModel) MarshalJSON() ([]byte, error) {
	var err error
	object := make(map[string]json.RawMessage)

	object["created"], err = json.Marshal(a.Created)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'created': %w", err)
	}

	object["id"], err = json.Marshal(a.ID)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'id': %w", err)
	}

	object["object"], err = json.Marshal(a.Object)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'object': %w", err)
	}

	object["owned_by"], err = json.Marshal(a.OwnedBy)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'owned_by': %w", err)
	}

	for fieldName, field := range a.AdditionalProperties {
		object[fieldName], err = json.Marshal(field)
		if err != nil {
			return nil, fmt.Errorf("error marshaling '%s': %w", fieldName, err)
		}
	}
	return json.Marshal(object)
}

// Getter for additional properties for OpenAIModelsResponse. Returns the specified
// element and whether it was found
func (a OpenAIModelsResponse) Get(fieldName string) (value interface{}, found bool) {
	if a.AdditionalProperties != nil {
		value, found = a.AdditionalProperties[fieldName]
	}
	return
}

// Setter for additional properties for OpenAIModelsResponse
func (a *OpenAIModelsResponse) Set(fieldName string, value interface{}) {
	if a.AdditionalProperties == nil {
		a.AdditionalProperties = make(map[string]any)
	}
	a.AdditionalProperties[fieldName] = value
}

// Override default JSON handling for OpenAIModelsResponse to handle AdditionalProperties
func (a *OpenAIModelsResponse) UnmarshalJSON(b []byte) error {
	object := make(map[string]json.RawMessage)
	err := json.Unmarshal(b, &object)
	if err != nil {
		return err
	}

	if raw, found := object["data"]; found {
		err = json.Unmarshal(raw, &a.Data)
		if err != nil {
			return fmt.Errorf("error reading 'data': %w", err)
		}
		delete(object, "data")
	}

	if raw, found := object["object"]; found {
		err = json.Unmarshal(raw, &a.Object)
		if err != nil {
			return fmt.Errorf("error reading 'object': %w", err)
		}
		delete(object, "object")
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

// Override default JSON handling for OpenAIModelsResponse to handle AdditionalProperties
func (a OpenAIModelsResponse) MarshalJSON() ([]byte, error) {
	var err error
	object := make(map[string]json.RawMessage)

	if a.Data != nil {
		object["data"], err = json.Marshal(a.Data)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'data': %w", err)
		}
	}

	object["object"], err = json.Marshal(a.Object)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'object': %w", err)
	}

	for fieldName, field := range a.AdditionalProperties {
		object[fieldName], err = json.Marshal(field)
		if err != nil {
			return nil, fmt.Errorf("error marshaling '%s': %w", fieldName, err)
		}
	}
	return json.Marshal(object)
}

// Getter for additional properties for OpenAIResponseRequest. Returns the specified
// element and whether it was found
func (a OpenAIResponseRequest) Get(fieldName string) (value interface{}, found bool) {
	if a.AdditionalProperties != nil {
		value, found = a.AdditionalProperties[fieldName]
	}
	return
}

// Setter for additional properties for OpenAIResponseRequest
func (a *OpenAIResponseRequest) Set(fieldName string, value interface{}) {
	if a.AdditionalProperties == nil {
		a.AdditionalProperties = make(map[string]any)
	}
	a.AdditionalProperties[fieldName] = value
}

// Override default JSON handling for OpenAIResponseRequest to handle AdditionalProperties
func (a *OpenAIResponseRequest) UnmarshalJSON(b []byte) error {
	object := make(map[string]json.RawMessage)
	err := json.Unmarshal(b, &object)
	if err != nil {
		return err
	}

	if raw, found := object["input"]; found {
		err = json.Unmarshal(raw, &a.Input)
		if err != nil {
			return fmt.Errorf("error reading 'input': %w", err)
		}
		delete(object, "input")
	}

	if raw, found := object["model"]; found {
		err = json.Unmarshal(raw, &a.Model)
		if err != nil {
			return fmt.Errorf("error reading 'model': %w", err)
		}
		delete(object, "model")
	}

	if raw, found := object["stream"]; found {
		err = json.Unmarshal(raw, &a.Stream)
		if err != nil {
			return fmt.Errorf("error reading 'stream': %w", err)
		}
		delete(object, "stream")
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

// Override default JSON handling for OpenAIResponseRequest to handle AdditionalProperties
func (a OpenAIResponseRequest) MarshalJSON() ([]byte, error) {
	var err error
	object := make(map[string]json.RawMessage)

	object["input"], err = json.Marshal(a.Input)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'input': %w", err)
	}

	object["model"], err = json.Marshal(a.Model)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'model': %w", err)
	}

	object["stream"], err = json.Marshal(a.Stream)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'stream': %w", err)
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

// AsAgentRunCallbackWorkbenchMessageDeltaEvent returns the union data inside the AgentRunCallbackWorkbenchStreamEvent as a AgentRunCallbackWorkbenchMessageDeltaEvent
func (t AgentRunCallbackWorkbenchStreamEvent) AsAgentRunCallbackWorkbenchMessageDeltaEvent() (AgentRunCallbackWorkbenchMessageDeltaEvent, error) {
	var body AgentRunCallbackWorkbenchMessageDeltaEvent
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromAgentRunCallbackWorkbenchMessageDeltaEvent overwrites any union data inside the AgentRunCallbackWorkbenchStreamEvent as the provided AgentRunCallbackWorkbenchMessageDeltaEvent
func (t *AgentRunCallbackWorkbenchStreamEvent) FromAgentRunCallbackWorkbenchMessageDeltaEvent(v AgentRunCallbackWorkbenchMessageDeltaEvent) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeAgentRunCallbackWorkbenchMessageDeltaEvent performs a merge with any union data inside the AgentRunCallbackWorkbenchStreamEvent, using the provided AgentRunCallbackWorkbenchMessageDeltaEvent
func (t *AgentRunCallbackWorkbenchStreamEvent) MergeAgentRunCallbackWorkbenchMessageDeltaEvent(v AgentRunCallbackWorkbenchMessageDeltaEvent) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsAgentRunCallbackWorkbenchMessageDoneEvent returns the union data inside the AgentRunCallbackWorkbenchStreamEvent as a AgentRunCallbackWorkbenchMessageDoneEvent
func (t AgentRunCallbackWorkbenchStreamEvent) AsAgentRunCallbackWorkbenchMessageDoneEvent() (AgentRunCallbackWorkbenchMessageDoneEvent, error) {
	var body AgentRunCallbackWorkbenchMessageDoneEvent
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromAgentRunCallbackWorkbenchMessageDoneEvent overwrites any union data inside the AgentRunCallbackWorkbenchStreamEvent as the provided AgentRunCallbackWorkbenchMessageDoneEvent
func (t *AgentRunCallbackWorkbenchStreamEvent) FromAgentRunCallbackWorkbenchMessageDoneEvent(v AgentRunCallbackWorkbenchMessageDoneEvent) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeAgentRunCallbackWorkbenchMessageDoneEvent performs a merge with any union data inside the AgentRunCallbackWorkbenchStreamEvent, using the provided AgentRunCallbackWorkbenchMessageDoneEvent
func (t *AgentRunCallbackWorkbenchStreamEvent) MergeAgentRunCallbackWorkbenchMessageDoneEvent(v AgentRunCallbackWorkbenchMessageDoneEvent) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsAgentRunCallbackWorkbenchThinkingDeltaEvent returns the union data inside the AgentRunCallbackWorkbenchStreamEvent as a AgentRunCallbackWorkbenchThinkingDeltaEvent
func (t AgentRunCallbackWorkbenchStreamEvent) AsAgentRunCallbackWorkbenchThinkingDeltaEvent() (AgentRunCallbackWorkbenchThinkingDeltaEvent, error) {
	var body AgentRunCallbackWorkbenchThinkingDeltaEvent
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromAgentRunCallbackWorkbenchThinkingDeltaEvent overwrites any union data inside the AgentRunCallbackWorkbenchStreamEvent as the provided AgentRunCallbackWorkbenchThinkingDeltaEvent
func (t *AgentRunCallbackWorkbenchStreamEvent) FromAgentRunCallbackWorkbenchThinkingDeltaEvent(v AgentRunCallbackWorkbenchThinkingDeltaEvent) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeAgentRunCallbackWorkbenchThinkingDeltaEvent performs a merge with any union data inside the AgentRunCallbackWorkbenchStreamEvent, using the provided AgentRunCallbackWorkbenchThinkingDeltaEvent
func (t *AgentRunCallbackWorkbenchStreamEvent) MergeAgentRunCallbackWorkbenchThinkingDeltaEvent(v AgentRunCallbackWorkbenchThinkingDeltaEvent) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsAgentRunCallbackWorkbenchThinkingDoneEvent returns the union data inside the AgentRunCallbackWorkbenchStreamEvent as a AgentRunCallbackWorkbenchThinkingDoneEvent
func (t AgentRunCallbackWorkbenchStreamEvent) AsAgentRunCallbackWorkbenchThinkingDoneEvent() (AgentRunCallbackWorkbenchThinkingDoneEvent, error) {
	var body AgentRunCallbackWorkbenchThinkingDoneEvent
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromAgentRunCallbackWorkbenchThinkingDoneEvent overwrites any union data inside the AgentRunCallbackWorkbenchStreamEvent as the provided AgentRunCallbackWorkbenchThinkingDoneEvent
func (t *AgentRunCallbackWorkbenchStreamEvent) FromAgentRunCallbackWorkbenchThinkingDoneEvent(v AgentRunCallbackWorkbenchThinkingDoneEvent) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeAgentRunCallbackWorkbenchThinkingDoneEvent performs a merge with any union data inside the AgentRunCallbackWorkbenchStreamEvent, using the provided AgentRunCallbackWorkbenchThinkingDoneEvent
func (t *AgentRunCallbackWorkbenchStreamEvent) MergeAgentRunCallbackWorkbenchThinkingDoneEvent(v AgentRunCallbackWorkbenchThinkingDoneEvent) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsAgentRunCallbackWorkbenchAgentStatusEvent returns the union data inside the AgentRunCallbackWorkbenchStreamEvent as a AgentRunCallbackWorkbenchAgentStatusEvent
func (t AgentRunCallbackWorkbenchStreamEvent) AsAgentRunCallbackWorkbenchAgentStatusEvent() (AgentRunCallbackWorkbenchAgentStatusEvent, error) {
	var body AgentRunCallbackWorkbenchAgentStatusEvent
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromAgentRunCallbackWorkbenchAgentStatusEvent overwrites any union data inside the AgentRunCallbackWorkbenchStreamEvent as the provided AgentRunCallbackWorkbenchAgentStatusEvent
func (t *AgentRunCallbackWorkbenchStreamEvent) FromAgentRunCallbackWorkbenchAgentStatusEvent(v AgentRunCallbackWorkbenchAgentStatusEvent) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeAgentRunCallbackWorkbenchAgentStatusEvent performs a merge with any union data inside the AgentRunCallbackWorkbenchStreamEvent, using the provided AgentRunCallbackWorkbenchAgentStatusEvent
func (t *AgentRunCallbackWorkbenchStreamEvent) MergeAgentRunCallbackWorkbenchAgentStatusEvent(v AgentRunCallbackWorkbenchAgentStatusEvent) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsAgentRunCallbackWorkbenchToolStartedEvent returns the union data inside the AgentRunCallbackWorkbenchStreamEvent as a AgentRunCallbackWorkbenchToolStartedEvent
func (t AgentRunCallbackWorkbenchStreamEvent) AsAgentRunCallbackWorkbenchToolStartedEvent() (AgentRunCallbackWorkbenchToolStartedEvent, error) {
	var body AgentRunCallbackWorkbenchToolStartedEvent
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromAgentRunCallbackWorkbenchToolStartedEvent overwrites any union data inside the AgentRunCallbackWorkbenchStreamEvent as the provided AgentRunCallbackWorkbenchToolStartedEvent
func (t *AgentRunCallbackWorkbenchStreamEvent) FromAgentRunCallbackWorkbenchToolStartedEvent(v AgentRunCallbackWorkbenchToolStartedEvent) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeAgentRunCallbackWorkbenchToolStartedEvent performs a merge with any union data inside the AgentRunCallbackWorkbenchStreamEvent, using the provided AgentRunCallbackWorkbenchToolStartedEvent
func (t *AgentRunCallbackWorkbenchStreamEvent) MergeAgentRunCallbackWorkbenchToolStartedEvent(v AgentRunCallbackWorkbenchToolStartedEvent) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsAgentRunCallbackWorkbenchToolDeltaEvent returns the union data inside the AgentRunCallbackWorkbenchStreamEvent as a AgentRunCallbackWorkbenchToolDeltaEvent
func (t AgentRunCallbackWorkbenchStreamEvent) AsAgentRunCallbackWorkbenchToolDeltaEvent() (AgentRunCallbackWorkbenchToolDeltaEvent, error) {
	var body AgentRunCallbackWorkbenchToolDeltaEvent
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromAgentRunCallbackWorkbenchToolDeltaEvent overwrites any union data inside the AgentRunCallbackWorkbenchStreamEvent as the provided AgentRunCallbackWorkbenchToolDeltaEvent
func (t *AgentRunCallbackWorkbenchStreamEvent) FromAgentRunCallbackWorkbenchToolDeltaEvent(v AgentRunCallbackWorkbenchToolDeltaEvent) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeAgentRunCallbackWorkbenchToolDeltaEvent performs a merge with any union data inside the AgentRunCallbackWorkbenchStreamEvent, using the provided AgentRunCallbackWorkbenchToolDeltaEvent
func (t *AgentRunCallbackWorkbenchStreamEvent) MergeAgentRunCallbackWorkbenchToolDeltaEvent(v AgentRunCallbackWorkbenchToolDeltaEvent) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsAgentRunCallbackWorkbenchToolCompletedEvent returns the union data inside the AgentRunCallbackWorkbenchStreamEvent as a AgentRunCallbackWorkbenchToolCompletedEvent
func (t AgentRunCallbackWorkbenchStreamEvent) AsAgentRunCallbackWorkbenchToolCompletedEvent() (AgentRunCallbackWorkbenchToolCompletedEvent, error) {
	var body AgentRunCallbackWorkbenchToolCompletedEvent
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromAgentRunCallbackWorkbenchToolCompletedEvent overwrites any union data inside the AgentRunCallbackWorkbenchStreamEvent as the provided AgentRunCallbackWorkbenchToolCompletedEvent
func (t *AgentRunCallbackWorkbenchStreamEvent) FromAgentRunCallbackWorkbenchToolCompletedEvent(v AgentRunCallbackWorkbenchToolCompletedEvent) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeAgentRunCallbackWorkbenchToolCompletedEvent performs a merge with any union data inside the AgentRunCallbackWorkbenchStreamEvent, using the provided AgentRunCallbackWorkbenchToolCompletedEvent
func (t *AgentRunCallbackWorkbenchStreamEvent) MergeAgentRunCallbackWorkbenchToolCompletedEvent(v AgentRunCallbackWorkbenchToolCompletedEvent) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsAgentRunCallbackWorkbenchArtifactUpdatedEvent returns the union data inside the AgentRunCallbackWorkbenchStreamEvent as a AgentRunCallbackWorkbenchArtifactUpdatedEvent
func (t AgentRunCallbackWorkbenchStreamEvent) AsAgentRunCallbackWorkbenchArtifactUpdatedEvent() (AgentRunCallbackWorkbenchArtifactUpdatedEvent, error) {
	var body AgentRunCallbackWorkbenchArtifactUpdatedEvent
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromAgentRunCallbackWorkbenchArtifactUpdatedEvent overwrites any union data inside the AgentRunCallbackWorkbenchStreamEvent as the provided AgentRunCallbackWorkbenchArtifactUpdatedEvent
func (t *AgentRunCallbackWorkbenchStreamEvent) FromAgentRunCallbackWorkbenchArtifactUpdatedEvent(v AgentRunCallbackWorkbenchArtifactUpdatedEvent) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeAgentRunCallbackWorkbenchArtifactUpdatedEvent performs a merge with any union data inside the AgentRunCallbackWorkbenchStreamEvent, using the provided AgentRunCallbackWorkbenchArtifactUpdatedEvent
func (t *AgentRunCallbackWorkbenchStreamEvent) MergeAgentRunCallbackWorkbenchArtifactUpdatedEvent(v AgentRunCallbackWorkbenchArtifactUpdatedEvent) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsAgentRunCallbackWorkbenchSourceUpdatedEvent returns the union data inside the AgentRunCallbackWorkbenchStreamEvent as a AgentRunCallbackWorkbenchSourceUpdatedEvent
func (t AgentRunCallbackWorkbenchStreamEvent) AsAgentRunCallbackWorkbenchSourceUpdatedEvent() (AgentRunCallbackWorkbenchSourceUpdatedEvent, error) {
	var body AgentRunCallbackWorkbenchSourceUpdatedEvent
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromAgentRunCallbackWorkbenchSourceUpdatedEvent overwrites any union data inside the AgentRunCallbackWorkbenchStreamEvent as the provided AgentRunCallbackWorkbenchSourceUpdatedEvent
func (t *AgentRunCallbackWorkbenchStreamEvent) FromAgentRunCallbackWorkbenchSourceUpdatedEvent(v AgentRunCallbackWorkbenchSourceUpdatedEvent) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeAgentRunCallbackWorkbenchSourceUpdatedEvent performs a merge with any union data inside the AgentRunCallbackWorkbenchStreamEvent, using the provided AgentRunCallbackWorkbenchSourceUpdatedEvent
func (t *AgentRunCallbackWorkbenchStreamEvent) MergeAgentRunCallbackWorkbenchSourceUpdatedEvent(v AgentRunCallbackWorkbenchSourceUpdatedEvent) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsAgentRunCallbackWorkbenchCardCommandEvent returns the union data inside the AgentRunCallbackWorkbenchStreamEvent as a AgentRunCallbackWorkbenchCardCommandEvent
func (t AgentRunCallbackWorkbenchStreamEvent) AsAgentRunCallbackWorkbenchCardCommandEvent() (AgentRunCallbackWorkbenchCardCommandEvent, error) {
	var body AgentRunCallbackWorkbenchCardCommandEvent
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromAgentRunCallbackWorkbenchCardCommandEvent overwrites any union data inside the AgentRunCallbackWorkbenchStreamEvent as the provided AgentRunCallbackWorkbenchCardCommandEvent
func (t *AgentRunCallbackWorkbenchStreamEvent) FromAgentRunCallbackWorkbenchCardCommandEvent(v AgentRunCallbackWorkbenchCardCommandEvent) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeAgentRunCallbackWorkbenchCardCommandEvent performs a merge with any union data inside the AgentRunCallbackWorkbenchStreamEvent, using the provided AgentRunCallbackWorkbenchCardCommandEvent
func (t *AgentRunCallbackWorkbenchStreamEvent) MergeAgentRunCallbackWorkbenchCardCommandEvent(v AgentRunCallbackWorkbenchCardCommandEvent) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsAgentRunCallbackWorkbenchErrorEvent returns the union data inside the AgentRunCallbackWorkbenchStreamEvent as a AgentRunCallbackWorkbenchErrorEvent
func (t AgentRunCallbackWorkbenchStreamEvent) AsAgentRunCallbackWorkbenchErrorEvent() (AgentRunCallbackWorkbenchErrorEvent, error) {
	var body AgentRunCallbackWorkbenchErrorEvent
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromAgentRunCallbackWorkbenchErrorEvent overwrites any union data inside the AgentRunCallbackWorkbenchStreamEvent as the provided AgentRunCallbackWorkbenchErrorEvent
func (t *AgentRunCallbackWorkbenchStreamEvent) FromAgentRunCallbackWorkbenchErrorEvent(v AgentRunCallbackWorkbenchErrorEvent) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeAgentRunCallbackWorkbenchErrorEvent performs a merge with any union data inside the AgentRunCallbackWorkbenchStreamEvent, using the provided AgentRunCallbackWorkbenchErrorEvent
func (t *AgentRunCallbackWorkbenchStreamEvent) MergeAgentRunCallbackWorkbenchErrorEvent(v AgentRunCallbackWorkbenchErrorEvent) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

func (t AgentRunCallbackWorkbenchStreamEvent) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *AgentRunCallbackWorkbenchStreamEvent) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsBuildSourceConfigBuildArgs0 returns the union data inside the BuildSourceConfig_BuildArgs_AdditionalProperties as a BuildSourceConfigBuildArgs0
func (t BuildSourceConfig_BuildArgs_AdditionalProperties) AsBuildSourceConfigBuildArgs0() (BuildSourceConfigBuildArgs0, error) {
	var body BuildSourceConfigBuildArgs0
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromBuildSourceConfigBuildArgs0 overwrites any union data inside the BuildSourceConfig_BuildArgs_AdditionalProperties as the provided BuildSourceConfigBuildArgs0
func (t *BuildSourceConfig_BuildArgs_AdditionalProperties) FromBuildSourceConfigBuildArgs0(v BuildSourceConfigBuildArgs0) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeBuildSourceConfigBuildArgs0 performs a merge with any union data inside the BuildSourceConfig_BuildArgs_AdditionalProperties, using the provided BuildSourceConfigBuildArgs0
func (t *BuildSourceConfig_BuildArgs_AdditionalProperties) MergeBuildSourceConfigBuildArgs0(v BuildSourceConfigBuildArgs0) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsBuildSourceConfigBuildArgs1 returns the union data inside the BuildSourceConfig_BuildArgs_AdditionalProperties as a BuildSourceConfigBuildArgs1
func (t BuildSourceConfig_BuildArgs_AdditionalProperties) AsBuildSourceConfigBuildArgs1() (BuildSourceConfigBuildArgs1, error) {
	var body BuildSourceConfigBuildArgs1
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromBuildSourceConfigBuildArgs1 overwrites any union data inside the BuildSourceConfig_BuildArgs_AdditionalProperties as the provided BuildSourceConfigBuildArgs1
func (t *BuildSourceConfig_BuildArgs_AdditionalProperties) FromBuildSourceConfigBuildArgs1(v BuildSourceConfigBuildArgs1) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeBuildSourceConfigBuildArgs1 performs a merge with any union data inside the BuildSourceConfig_BuildArgs_AdditionalProperties, using the provided BuildSourceConfigBuildArgs1
func (t *BuildSourceConfig_BuildArgs_AdditionalProperties) MergeBuildSourceConfigBuildArgs1(v BuildSourceConfigBuildArgs1) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsBuildSourceConfigBuildArgs2 returns the union data inside the BuildSourceConfig_BuildArgs_AdditionalProperties as a BuildSourceConfigBuildArgs2
func (t BuildSourceConfig_BuildArgs_AdditionalProperties) AsBuildSourceConfigBuildArgs2() (BuildSourceConfigBuildArgs2, error) {
	var body BuildSourceConfigBuildArgs2
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromBuildSourceConfigBuildArgs2 overwrites any union data inside the BuildSourceConfig_BuildArgs_AdditionalProperties as the provided BuildSourceConfigBuildArgs2
func (t *BuildSourceConfig_BuildArgs_AdditionalProperties) FromBuildSourceConfigBuildArgs2(v BuildSourceConfigBuildArgs2) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeBuildSourceConfigBuildArgs2 performs a merge with any union data inside the BuildSourceConfig_BuildArgs_AdditionalProperties, using the provided BuildSourceConfigBuildArgs2
func (t *BuildSourceConfig_BuildArgs_AdditionalProperties) MergeBuildSourceConfigBuildArgs2(v BuildSourceConfigBuildArgs2) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

func (t BuildSourceConfig_BuildArgs_AdditionalProperties) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *BuildSourceConfig_BuildArgs_AdditionalProperties) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsBuildSourceConfigVariables0 returns the union data inside the BuildSourceConfig_Variables_AdditionalProperties as a BuildSourceConfigVariables0
func (t BuildSourceConfig_Variables_AdditionalProperties) AsBuildSourceConfigVariables0() (BuildSourceConfigVariables0, error) {
	var body BuildSourceConfigVariables0
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromBuildSourceConfigVariables0 overwrites any union data inside the BuildSourceConfig_Variables_AdditionalProperties as the provided BuildSourceConfigVariables0
func (t *BuildSourceConfig_Variables_AdditionalProperties) FromBuildSourceConfigVariables0(v BuildSourceConfigVariables0) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeBuildSourceConfigVariables0 performs a merge with any union data inside the BuildSourceConfig_Variables_AdditionalProperties, using the provided BuildSourceConfigVariables0
func (t *BuildSourceConfig_Variables_AdditionalProperties) MergeBuildSourceConfigVariables0(v BuildSourceConfigVariables0) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsBuildSourceConfigVariables1 returns the union data inside the BuildSourceConfig_Variables_AdditionalProperties as a BuildSourceConfigVariables1
func (t BuildSourceConfig_Variables_AdditionalProperties) AsBuildSourceConfigVariables1() (BuildSourceConfigVariables1, error) {
	var body BuildSourceConfigVariables1
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromBuildSourceConfigVariables1 overwrites any union data inside the BuildSourceConfig_Variables_AdditionalProperties as the provided BuildSourceConfigVariables1
func (t *BuildSourceConfig_Variables_AdditionalProperties) FromBuildSourceConfigVariables1(v BuildSourceConfigVariables1) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeBuildSourceConfigVariables1 performs a merge with any union data inside the BuildSourceConfig_Variables_AdditionalProperties, using the provided BuildSourceConfigVariables1
func (t *BuildSourceConfig_Variables_AdditionalProperties) MergeBuildSourceConfigVariables1(v BuildSourceConfigVariables1) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsBuildSourceConfigVariables2 returns the union data inside the BuildSourceConfig_Variables_AdditionalProperties as a BuildSourceConfigVariables2
func (t BuildSourceConfig_Variables_AdditionalProperties) AsBuildSourceConfigVariables2() (BuildSourceConfigVariables2, error) {
	var body BuildSourceConfigVariables2
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromBuildSourceConfigVariables2 overwrites any union data inside the BuildSourceConfig_Variables_AdditionalProperties as the provided BuildSourceConfigVariables2
func (t *BuildSourceConfig_Variables_AdditionalProperties) FromBuildSourceConfigVariables2(v BuildSourceConfigVariables2) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeBuildSourceConfigVariables2 performs a merge with any union data inside the BuildSourceConfig_Variables_AdditionalProperties, using the provided BuildSourceConfigVariables2
func (t *BuildSourceConfig_Variables_AdditionalProperties) MergeBuildSourceConfigVariables2(v BuildSourceConfigVariables2) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

func (t BuildSourceConfig_Variables_AdditionalProperties) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *BuildSourceConfig_Variables_AdditionalProperties) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsCohereRerankRequestDocuments0 returns the union data inside the CohereRerankRequest_Documents_Item as a CohereRerankRequestDocuments0
func (t CohereRerankRequest_Documents_Item) AsCohereRerankRequestDocuments0() (CohereRerankRequestDocuments0, error) {
	var body CohereRerankRequestDocuments0
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromCohereRerankRequestDocuments0 overwrites any union data inside the CohereRerankRequest_Documents_Item as the provided CohereRerankRequestDocuments0
func (t *CohereRerankRequest_Documents_Item) FromCohereRerankRequestDocuments0(v CohereRerankRequestDocuments0) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeCohereRerankRequestDocuments0 performs a merge with any union data inside the CohereRerankRequest_Documents_Item, using the provided CohereRerankRequestDocuments0
func (t *CohereRerankRequest_Documents_Item) MergeCohereRerankRequestDocuments0(v CohereRerankRequestDocuments0) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsCohereRerankRequestDocuments1 returns the union data inside the CohereRerankRequest_Documents_Item as a CohereRerankRequestDocuments1
func (t CohereRerankRequest_Documents_Item) AsCohereRerankRequestDocuments1() (CohereRerankRequestDocuments1, error) {
	var body CohereRerankRequestDocuments1
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromCohereRerankRequestDocuments1 overwrites any union data inside the CohereRerankRequest_Documents_Item as the provided CohereRerankRequestDocuments1
func (t *CohereRerankRequest_Documents_Item) FromCohereRerankRequestDocuments1(v CohereRerankRequestDocuments1) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeCohereRerankRequestDocuments1 performs a merge with any union data inside the CohereRerankRequest_Documents_Item, using the provided CohereRerankRequestDocuments1
func (t *CohereRerankRequest_Documents_Item) MergeCohereRerankRequestDocuments1(v CohereRerankRequestDocuments1) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

func (t CohereRerankRequest_Documents_Item) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *CohereRerankRequest_Documents_Item) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsWorkbenchMessageDeltaEvent returns the union data inside the WorkbenchStreamEvent as a WorkbenchMessageDeltaEvent
func (t WorkbenchStreamEvent) AsWorkbenchMessageDeltaEvent() (WorkbenchMessageDeltaEvent, error) {
	var body WorkbenchMessageDeltaEvent
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromWorkbenchMessageDeltaEvent overwrites any union data inside the WorkbenchStreamEvent as the provided WorkbenchMessageDeltaEvent
func (t *WorkbenchStreamEvent) FromWorkbenchMessageDeltaEvent(v WorkbenchMessageDeltaEvent) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeWorkbenchMessageDeltaEvent performs a merge with any union data inside the WorkbenchStreamEvent, using the provided WorkbenchMessageDeltaEvent
func (t *WorkbenchStreamEvent) MergeWorkbenchMessageDeltaEvent(v WorkbenchMessageDeltaEvent) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsWorkbenchMessageDoneEvent returns the union data inside the WorkbenchStreamEvent as a WorkbenchMessageDoneEvent
func (t WorkbenchStreamEvent) AsWorkbenchMessageDoneEvent() (WorkbenchMessageDoneEvent, error) {
	var body WorkbenchMessageDoneEvent
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromWorkbenchMessageDoneEvent overwrites any union data inside the WorkbenchStreamEvent as the provided WorkbenchMessageDoneEvent
func (t *WorkbenchStreamEvent) FromWorkbenchMessageDoneEvent(v WorkbenchMessageDoneEvent) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeWorkbenchMessageDoneEvent performs a merge with any union data inside the WorkbenchStreamEvent, using the provided WorkbenchMessageDoneEvent
func (t *WorkbenchStreamEvent) MergeWorkbenchMessageDoneEvent(v WorkbenchMessageDoneEvent) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsWorkbenchThinkingDeltaEvent returns the union data inside the WorkbenchStreamEvent as a WorkbenchThinkingDeltaEvent
func (t WorkbenchStreamEvent) AsWorkbenchThinkingDeltaEvent() (WorkbenchThinkingDeltaEvent, error) {
	var body WorkbenchThinkingDeltaEvent
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromWorkbenchThinkingDeltaEvent overwrites any union data inside the WorkbenchStreamEvent as the provided WorkbenchThinkingDeltaEvent
func (t *WorkbenchStreamEvent) FromWorkbenchThinkingDeltaEvent(v WorkbenchThinkingDeltaEvent) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeWorkbenchThinkingDeltaEvent performs a merge with any union data inside the WorkbenchStreamEvent, using the provided WorkbenchThinkingDeltaEvent
func (t *WorkbenchStreamEvent) MergeWorkbenchThinkingDeltaEvent(v WorkbenchThinkingDeltaEvent) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsWorkbenchThinkingDoneEvent returns the union data inside the WorkbenchStreamEvent as a WorkbenchThinkingDoneEvent
func (t WorkbenchStreamEvent) AsWorkbenchThinkingDoneEvent() (WorkbenchThinkingDoneEvent, error) {
	var body WorkbenchThinkingDoneEvent
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromWorkbenchThinkingDoneEvent overwrites any union data inside the WorkbenchStreamEvent as the provided WorkbenchThinkingDoneEvent
func (t *WorkbenchStreamEvent) FromWorkbenchThinkingDoneEvent(v WorkbenchThinkingDoneEvent) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeWorkbenchThinkingDoneEvent performs a merge with any union data inside the WorkbenchStreamEvent, using the provided WorkbenchThinkingDoneEvent
func (t *WorkbenchStreamEvent) MergeWorkbenchThinkingDoneEvent(v WorkbenchThinkingDoneEvent) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsWorkbenchAgentStatusEvent returns the union data inside the WorkbenchStreamEvent as a WorkbenchAgentStatusEvent
func (t WorkbenchStreamEvent) AsWorkbenchAgentStatusEvent() (WorkbenchAgentStatusEvent, error) {
	var body WorkbenchAgentStatusEvent
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromWorkbenchAgentStatusEvent overwrites any union data inside the WorkbenchStreamEvent as the provided WorkbenchAgentStatusEvent
func (t *WorkbenchStreamEvent) FromWorkbenchAgentStatusEvent(v WorkbenchAgentStatusEvent) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeWorkbenchAgentStatusEvent performs a merge with any union data inside the WorkbenchStreamEvent, using the provided WorkbenchAgentStatusEvent
func (t *WorkbenchStreamEvent) MergeWorkbenchAgentStatusEvent(v WorkbenchAgentStatusEvent) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsWorkbenchToolStartedEvent returns the union data inside the WorkbenchStreamEvent as a WorkbenchToolStartedEvent
func (t WorkbenchStreamEvent) AsWorkbenchToolStartedEvent() (WorkbenchToolStartedEvent, error) {
	var body WorkbenchToolStartedEvent
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromWorkbenchToolStartedEvent overwrites any union data inside the WorkbenchStreamEvent as the provided WorkbenchToolStartedEvent
func (t *WorkbenchStreamEvent) FromWorkbenchToolStartedEvent(v WorkbenchToolStartedEvent) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeWorkbenchToolStartedEvent performs a merge with any union data inside the WorkbenchStreamEvent, using the provided WorkbenchToolStartedEvent
func (t *WorkbenchStreamEvent) MergeWorkbenchToolStartedEvent(v WorkbenchToolStartedEvent) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsWorkbenchToolDeltaEvent returns the union data inside the WorkbenchStreamEvent as a WorkbenchToolDeltaEvent
func (t WorkbenchStreamEvent) AsWorkbenchToolDeltaEvent() (WorkbenchToolDeltaEvent, error) {
	var body WorkbenchToolDeltaEvent
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromWorkbenchToolDeltaEvent overwrites any union data inside the WorkbenchStreamEvent as the provided WorkbenchToolDeltaEvent
func (t *WorkbenchStreamEvent) FromWorkbenchToolDeltaEvent(v WorkbenchToolDeltaEvent) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeWorkbenchToolDeltaEvent performs a merge with any union data inside the WorkbenchStreamEvent, using the provided WorkbenchToolDeltaEvent
func (t *WorkbenchStreamEvent) MergeWorkbenchToolDeltaEvent(v WorkbenchToolDeltaEvent) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsWorkbenchToolCompletedEvent returns the union data inside the WorkbenchStreamEvent as a WorkbenchToolCompletedEvent
func (t WorkbenchStreamEvent) AsWorkbenchToolCompletedEvent() (WorkbenchToolCompletedEvent, error) {
	var body WorkbenchToolCompletedEvent
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromWorkbenchToolCompletedEvent overwrites any union data inside the WorkbenchStreamEvent as the provided WorkbenchToolCompletedEvent
func (t *WorkbenchStreamEvent) FromWorkbenchToolCompletedEvent(v WorkbenchToolCompletedEvent) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeWorkbenchToolCompletedEvent performs a merge with any union data inside the WorkbenchStreamEvent, using the provided WorkbenchToolCompletedEvent
func (t *WorkbenchStreamEvent) MergeWorkbenchToolCompletedEvent(v WorkbenchToolCompletedEvent) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsWorkbenchArtifactUpdatedEvent returns the union data inside the WorkbenchStreamEvent as a WorkbenchArtifactUpdatedEvent
func (t WorkbenchStreamEvent) AsWorkbenchArtifactUpdatedEvent() (WorkbenchArtifactUpdatedEvent, error) {
	var body WorkbenchArtifactUpdatedEvent
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromWorkbenchArtifactUpdatedEvent overwrites any union data inside the WorkbenchStreamEvent as the provided WorkbenchArtifactUpdatedEvent
func (t *WorkbenchStreamEvent) FromWorkbenchArtifactUpdatedEvent(v WorkbenchArtifactUpdatedEvent) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeWorkbenchArtifactUpdatedEvent performs a merge with any union data inside the WorkbenchStreamEvent, using the provided WorkbenchArtifactUpdatedEvent
func (t *WorkbenchStreamEvent) MergeWorkbenchArtifactUpdatedEvent(v WorkbenchArtifactUpdatedEvent) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsWorkbenchSourceUpdatedEvent returns the union data inside the WorkbenchStreamEvent as a WorkbenchSourceUpdatedEvent
func (t WorkbenchStreamEvent) AsWorkbenchSourceUpdatedEvent() (WorkbenchSourceUpdatedEvent, error) {
	var body WorkbenchSourceUpdatedEvent
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromWorkbenchSourceUpdatedEvent overwrites any union data inside the WorkbenchStreamEvent as the provided WorkbenchSourceUpdatedEvent
func (t *WorkbenchStreamEvent) FromWorkbenchSourceUpdatedEvent(v WorkbenchSourceUpdatedEvent) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeWorkbenchSourceUpdatedEvent performs a merge with any union data inside the WorkbenchStreamEvent, using the provided WorkbenchSourceUpdatedEvent
func (t *WorkbenchStreamEvent) MergeWorkbenchSourceUpdatedEvent(v WorkbenchSourceUpdatedEvent) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsWorkbenchCardCommandEvent returns the union data inside the WorkbenchStreamEvent as a WorkbenchCardCommandEvent
func (t WorkbenchStreamEvent) AsWorkbenchCardCommandEvent() (WorkbenchCardCommandEvent, error) {
	var body WorkbenchCardCommandEvent
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromWorkbenchCardCommandEvent overwrites any union data inside the WorkbenchStreamEvent as the provided WorkbenchCardCommandEvent
func (t *WorkbenchStreamEvent) FromWorkbenchCardCommandEvent(v WorkbenchCardCommandEvent) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeWorkbenchCardCommandEvent performs a merge with any union data inside the WorkbenchStreamEvent, using the provided WorkbenchCardCommandEvent
func (t *WorkbenchStreamEvent) MergeWorkbenchCardCommandEvent(v WorkbenchCardCommandEvent) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsWorkbenchErrorEvent returns the union data inside the WorkbenchStreamEvent as a WorkbenchErrorEvent
func (t WorkbenchStreamEvent) AsWorkbenchErrorEvent() (WorkbenchErrorEvent, error) {
	var body WorkbenchErrorEvent
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromWorkbenchErrorEvent overwrites any union data inside the WorkbenchStreamEvent as the provided WorkbenchErrorEvent
func (t *WorkbenchStreamEvent) FromWorkbenchErrorEvent(v WorkbenchErrorEvent) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeWorkbenchErrorEvent performs a merge with any union data inside the WorkbenchStreamEvent, using the provided WorkbenchErrorEvent
func (t *WorkbenchStreamEvent) MergeWorkbenchErrorEvent(v WorkbenchErrorEvent) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

func (t WorkbenchStreamEvent) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *WorkbenchStreamEvent) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}
