// Code generated from OpenSoha permission contracts. DO NOT EDIT.

package auth

import _ "embed"

const (
	PermissionCatalogVersion = "2.0.0"
	PermissionCatalogContentSHA256 = "f744896d362070a1e7cdc1eda89c046bee8132f37c6f323dafdd05644e2dbd3c"
	PermissionCatalogSHA256 = "354ad414691ff0578d805036bbbe80da28d242d0bbabe4699e490af1b6bd87bb"
)

//go:embed permission-catalog.json
var permissionCatalogJSON []byte

// PermissionCatalogJSON returns a copy of the canonical permission catalog.
func PermissionCatalogJSON() []byte { return append([]byte(nil), permissionCatalogJSON...) }
