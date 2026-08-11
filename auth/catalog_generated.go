// Code generated from OpenSoha permission contracts. DO NOT EDIT.

package auth

import _ "embed"

const (
	PermissionCatalogVersion = "1.3.0"
	PermissionCatalogContentSHA256 = "fe738c69a160e58cf919ce6de8380d6427be6f8848cabad7d807ec82fdea8a3b"
	PermissionCatalogSHA256 = "206f12ff27279b1f345a87cd93be82789d1a79df311de9e7a86ab8cf75dde013"
)

//go:embed permission-catalog.json
var permissionCatalogJSON []byte

// PermissionCatalogJSON returns a copy of the canonical permission catalog.
func PermissionCatalogJSON() []byte { return append([]byte(nil), permissionCatalogJSON...) }
