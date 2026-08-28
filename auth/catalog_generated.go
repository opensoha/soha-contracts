// Code generated from OpenSoha permission contracts. DO NOT EDIT.

package auth

import _ "embed"

const (
	PermissionCatalogVersion = "2.0.0"
	PermissionCatalogContentSHA256 = "0da6a1dcd41f78e1699416a03ed0808f39cd569df6ebc7d242e1b35e5001830c"
	PermissionCatalogSHA256 = "7ef12f0a326d2fa3a85e2be2095817fb5a3932427ae1bf35cca91ed026e32a3a"
)

//go:embed permission-catalog.json
var permissionCatalogJSON []byte

// PermissionCatalogJSON returns a copy of the canonical permission catalog.
func PermissionCatalogJSON() []byte { return append([]byte(nil), permissionCatalogJSON...) }
