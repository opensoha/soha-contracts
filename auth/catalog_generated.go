// Code generated from OpenSoha permission contracts. DO NOT EDIT.

package auth

import _ "embed"

const (
	PermissionCatalogVersion = "2.0.0"
	PermissionCatalogContentSHA256 = "be2c1caf05e680f75b40e6ce017e55f1c35f7977cf045ca3ec562adbd6cb37dc"
	PermissionCatalogSHA256 = "2822fab5a3963270e806884ce46402bb42f3384d5530e6ae0acc8bf253811de5"
)

//go:embed permission-catalog.json
var permissionCatalogJSON []byte

// PermissionCatalogJSON returns a copy of the canonical permission catalog.
func PermissionCatalogJSON() []byte { return append([]byte(nil), permissionCatalogJSON...) }
