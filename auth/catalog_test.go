package auth

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"testing"
)

func TestPermissionCatalogMetadataMatchesEmbeddedArtifact(t *testing.T) {
	raw := PermissionCatalogJSON()
	var catalog struct {
		CatalogVersion string `json:"catalogVersion"`
		ContentHash    string `json:"contentHash"`
	}
	if err := json.Unmarshal(raw, &catalog); err != nil {
		t.Fatalf("decode permission catalog: %v", err)
	}
	if catalog.CatalogVersion != PermissionCatalogVersion {
		t.Fatalf("catalog version = %q, want %q", catalog.CatalogVersion, PermissionCatalogVersion)
	}
	if catalog.ContentHash != "sha256:"+PermissionCatalogContentSHA256 {
		t.Fatalf("content hash = %q", catalog.ContentHash)
	}
	sum := sha256.Sum256(raw)
	if got := hex.EncodeToString(sum[:]); got != PermissionCatalogSHA256 {
		t.Fatalf("artifact hash = %q, want %q", got, PermissionCatalogSHA256)
	}
}

func TestPermissionCatalogIncludesVirtualizationStorageView(t *testing.T) {
	var catalog struct {
		Permissions []struct {
			Key        string `json:"key"`
			Action     string `json:"action"`
			RiskLevel  string `json:"riskLevel"`
			Status     string `json:"status"`
			Assignable bool   `json:"assignable"`
		} `json:"permissions"`
	}
	if err := json.Unmarshal(PermissionCatalogJSON(), &catalog); err != nil {
		t.Fatalf("decode permission catalog: %v", err)
	}
	for _, permission := range catalog.Permissions {
		if permission.Key == "virtualization.storage.view" {
			if permission.Action != "view" || permission.RiskLevel != "read" || permission.Status != "active" || !permission.Assignable {
				t.Fatalf("unexpected storage permission: %#v", permission)
			}
			return
		}
	}
	t.Fatal("virtualization.storage.view is missing")
}
