// Code generated from OpenSoha permission contracts. DO NOT EDIT.

package auth

import _ "embed"

const (
	PermissionCatalogVersion = "2.0.0"
	PermissionCatalogContentSHA256 = "0efe8f496fa0fb29703584586f9306d04fdb5b9a38eaea78532107698230f609"
	PermissionCatalogSHA256 = "5f1d8b49fb4155473f61ff255002ec7a13e3f63d718d5f9fc592daa93da610df"
)

//go:embed permission-catalog.json
var permissionCatalogJSON []byte

// PermissionCatalogJSON returns a copy of the canonical permission catalog.
func PermissionCatalogJSON() []byte { return append([]byte(nil), permissionCatalogJSON...) }
