package sohaapi

import "testing"

func TestKubernetesWorkloadSnapshotRequestGeneratedShape(t *testing.T) {
	request := KubernetesWorkloadSnapshotRequest{
		Namespace:     "default",
		SourceKind:    KubernetesWorkloadSnapshotSourceKindDeployment,
		SourceName:    "reports",
		TargetKind:    KubernetesWorkloadSnapshotTargetKindCronJob,
		TargetName:    "reports-schedule",
		Inherit:       []KubernetesWorkloadSnapshotInheritance{KubernetesWorkloadSnapshotInheritanceEnvironment},
		RestartPolicy: KubernetesWorkloadSnapshotRestartPolicyNever,
	}
	if request.Namespace != "default" || len(request.Inherit) != 1 {
		t.Fatalf("request = %#v", request)
	}
}

func TestWorkbenchMessageDoneEventRoleGeneratedName(t *testing.T) {
	if WorkbenchMessageDoneEventRoleAssistant != WorkbenchMessageDoneEventRole("assistant") {
		t.Fatalf("role = %q", WorkbenchMessageDoneEventRoleAssistant)
	}
}

func TestWorkbenchMessageDeltaEventLegacyRoleName(t *testing.T) {
	if Assistant != WorkbenchMessageDeltaEventRole("assistant") {
		t.Fatalf("role = %q", Assistant)
	}
}

func TestKubernetesResourceSearchGeneratedShape(t *testing.T) {
	result := KubernetesResourceSearchResult{
		Items: []KubernetesResourceSearchItem{
			{
				Resource: KubernetesResourceRef{
					APIVersion: "v1",
					ClusterID:  "cluster-a",
					Kind:       "Pod",
					Name:       "api-0",
					Namespace:  "production",
					ScopeMode:  KubernetesResourceScopeModeNamespace,
				},
				Status: "Running",
			},
		},
		Truncated: false,
	}
	if len(result.Items) != 1 || result.Items[0].Resource.Name != "api-0" {
		t.Fatalf("result = %#v", result)
	}
}

func TestLegacyEnumAliasesRemainSourceCompatible(t *testing.T) {
	if Allow != ScopeGrantEffect("allow") || Deny != ScopeGrantEffect("deny") {
		t.Fatalf("scope grant effect aliases = %q/%q", Allow, Deny)
	}
	if ListAIGatewayRelayModelCallsParamsStatusSuccess != ListAIGatewayRelayModelCallsParamsStatus("success") ||
		ListAIGatewayRelayModelCallsParamsStatusFailure != ListAIGatewayRelayModelCallsParamsStatus("failure") ||
		ListAIGatewayRelayModelCallsParamsStatusCancelled != ListAIGatewayRelayModelCallsParamsStatus("cancelled") ||
		ListAIGatewayRelayModelCallsParamsStatusClientCancelled != ListAIGatewayRelayModelCallsParamsStatus("client_cancelled") ||
		ListAIGatewayRelayModelCallsParamsStatusRateLimited != ListAIGatewayRelayModelCallsParamsStatus("rate_limited") ||
		ListAIGatewayRelayModelCallsParamsStatusPolicyDenied != ListAIGatewayRelayModelCallsParamsStatus("policy_denied") {
		t.Fatalf("relay model call status aliases changed")
	}
	if TCP != DockerContainerPortInputProtocolTCP || UDP != DockerContainerPortInputProtocolUDP {
		t.Fatalf("docker protocol aliases = %q/%q", TCP, UDP)
	}
	if Critical != MarketplaceAdvisorySeverityCritical ||
		High != MarketplaceAdvisorySeverityHigh ||
		Medium != MarketplaceAdvisorySeverityMedium ||
		Low != MarketplaceAdvisorySeverityLow {
		t.Fatalf("marketplace severity aliases = %q/%q/%q/%q", Critical, High, Medium, Low)
	}
	if ListAIGatewayRelayUpstreamsParamsStatusActive != Active ||
		ListAIGatewayRelayUpstreamsParamsStatusDegraded != Degraded ||
		ListAIGatewayRelayUpstreamsParamsStatusDisabled != Disabled {
		t.Fatalf("relay status aliases changed")
	}
}
