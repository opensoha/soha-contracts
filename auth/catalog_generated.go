// Code generated from OpenSoha permission contracts. DO NOT EDIT.

package auth

import _ "embed"

const (
	PermissionCatalogVersion = "1.2.0"
	PermissionCatalogContentSHA256 = "5d7a468df65eb73a082457b0673b009034b5ed28f6a2fb4885b59a26fb7ca239"
	PermissionCatalogSHA256 = "589f4abadceea081d9ca7b6508f38c55785bdc98496c6c9b2a2055e3abc58dd5"
)

//go:embed permission-catalog.json
var permissionCatalogJSON []byte

// PermissionCatalogJSON returns a copy of the canonical permission catalog.
func PermissionCatalogJSON() []byte { return append([]byte(nil), permissionCatalogJSON...) }
