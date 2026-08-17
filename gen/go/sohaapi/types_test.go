package sohaapi

import "testing"

func TestKubernetesWorkloadSnapshotRequestGeneratedShape(t *testing.T) {
	request := KubernetesWorkloadSnapshotRequest{
		Namespace:     "default",
		SourceKind:    KubernetesWorkloadSnapshotSourceKindDeployment,
		SourceName:    "reports",
		TargetKind:    KubernetesWorkloadSnapshotTargetKindCronJob,
		TargetName:    "reports-schedule",
		RestartPolicy: KubernetesWorkloadSnapshotRestartPolicyNever,
	}
	if request.Namespace != "default" {
		t.Fatalf("namespace = %q, want default", request.Namespace)
	}
}
