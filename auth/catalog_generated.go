// Code generated from OpenSoha permission contracts. DO NOT EDIT.

package auth

import _ "embed"

const (
	PermissionCatalogVersion = "2.0.0"
	PermissionCatalogContentSHA256 = "b1a07eef4c26f91735b65ec18799e530271a47d0eed1de60f9a955d11208f9a2"
	PermissionCatalogSHA256 = "e89a89c109f4dca0ba3a5bea754df659aa8d9bf0298344e5f0b9c5cc7520a2d2"
)

//go:embed permission-catalog.json
var permissionCatalogJSON []byte

// PermissionCatalogJSON returns a copy of the canonical permission catalog.
func PermissionCatalogJSON() []byte { return append([]byte(nil), permissionCatalogJSON...) }
