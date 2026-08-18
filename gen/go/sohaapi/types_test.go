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
