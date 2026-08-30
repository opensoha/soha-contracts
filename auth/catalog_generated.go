// Code generated from OpenSoha permission contracts. DO NOT EDIT.

package auth

import _ "embed"

const (
	PermissionCatalogVersion = "2.0.0"
	PermissionCatalogContentSHA256 = "5aa6fa361c2ab8d987415fb1eb93218510deb9359a68212dfe86655dbfaeea5e"
	PermissionCatalogSHA256 = "01c80a71427e3d7f4898212fb1667405c27f1b798318d9eac86ece8262cd188c"
)

//go:embed permission-catalog.json
var permissionCatalogJSON []byte

// PermissionCatalogJSON returns a copy of the canonical permission catalog.
func PermissionCatalogJSON() []byte { return append([]byte(nil), permissionCatalogJSON...) }
