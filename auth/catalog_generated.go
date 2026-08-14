// Code generated from OpenSoha permission contracts. DO NOT EDIT.

package auth

import _ "embed"

const (
	PermissionCatalogVersion = "2.0.0"
	PermissionCatalogContentSHA256 = "82626eb53d9cd6b53ee502ae32285392c672f7f186a7dbefb42bf1f700d67e98"
	PermissionCatalogSHA256 = "f563a9be8901002e790f271d589c9f4b8b05dcd8e9b4bba767b835b1cc606b76"
)

//go:embed permission-catalog.json
var permissionCatalogJSON []byte

// PermissionCatalogJSON returns a copy of the canonical permission catalog.
func PermissionCatalogJSON() []byte { return append([]byte(nil), permissionCatalogJSON...) }
