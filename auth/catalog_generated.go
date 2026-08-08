// Code generated from OpenSoha permission contracts. DO NOT EDIT.

package auth

import _ "embed"

const (
	PermissionCatalogVersion = "1.3.0"
	PermissionCatalogContentSHA256 = "182dbfb4b9962dc4a0c344c2ce7622ec4424cf5ff0892a73b0b6806eddf6f2c7"
	PermissionCatalogSHA256 = "f9956b1ee4e3d981a4e93e7078006aa915a5eb18fb2347fa0d81bc8346b8958b"
)

//go:embed permission-catalog.json
var permissionCatalogJSON []byte

// PermissionCatalogJSON returns a copy of the canonical permission catalog.
func PermissionCatalogJSON() []byte { return append([]byte(nil), permissionCatalogJSON...) }
