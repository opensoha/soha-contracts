// Code generated from OpenSoha permission contracts. DO NOT EDIT.

package auth

import _ "embed"

const (
	PermissionCatalogVersion = "2.0.0"
	PermissionCatalogContentSHA256 = "2d98b53189afa67c09a2917325e3b1b1b70aafe00d370bdabaef7d6b11ef493c"
	PermissionCatalogSHA256 = "ae344007ad96fa0b24a0bf74121e1db762a1b904e58b8c8b53a8ac9168abea80"
)

//go:embed permission-catalog.json
var permissionCatalogJSON []byte

// PermissionCatalogJSON returns a copy of the canonical permission catalog.
func PermissionCatalogJSON() []byte { return append([]byte(nil), permissionCatalogJSON...) }
