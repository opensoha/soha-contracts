// Code generated from OpenSoha permission contracts. DO NOT EDIT.

package auth

import _ "embed"

const (
	PermissionCatalogVersion = "2.0.0"
	PermissionCatalogContentSHA256 = "58321138575c22b20933cf2c1db2742b741a123bb8f6a2983bb436e84be4e58a"
	PermissionCatalogSHA256 = "38ac1a2e7087ff4831e8f6390b804e32f918b49b9af3165fdf2df286d64fe3c8"
)

//go:embed permission-catalog.json
var permissionCatalogJSON []byte

// PermissionCatalogJSON returns a copy of the canonical permission catalog.
func PermissionCatalogJSON() []byte { return append([]byte(nil), permissionCatalogJSON...) }
