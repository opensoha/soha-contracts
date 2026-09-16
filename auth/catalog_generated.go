// Code generated from OpenSoha permission contracts. DO NOT EDIT.

package auth

import _ "embed"

const (
	PermissionCatalogVersion = "2.0.0"
	PermissionCatalogContentSHA256 = "5b4c9eed71c549d61a6fd538db542681cb2932619d66ab031c74fcbaaa0c9867"
	PermissionCatalogSHA256 = "e1e4bab807f64894c9bbde245394902fb1e82253d5847585102560ea6f853894"
)

//go:embed permission-catalog.json
var permissionCatalogJSON []byte

// PermissionCatalogJSON returns a copy of the canonical permission catalog.
func PermissionCatalogJSON() []byte { return append([]byte(nil), permissionCatalogJSON...) }
