// Code generated from OpenSoha permission contracts. DO NOT EDIT.

package auth

import _ "embed"

const (
	PermissionCatalogVersion = "2.0.0"
	PermissionCatalogContentSHA256 = "fa8516821cd30722887c6fa4ce7483b7e00ec56a299508f1d8eb8a95eba39188"
	PermissionCatalogSHA256 = "3b89fc97ce6d33cda0f5a5e488e15e4d9614717be1396311b13dd411b9bdddd1"
)

//go:embed permission-catalog.json
var permissionCatalogJSON []byte

// PermissionCatalogJSON returns a copy of the canonical permission catalog.
func PermissionCatalogJSON() []byte { return append([]byte(nil), permissionCatalogJSON...) }
