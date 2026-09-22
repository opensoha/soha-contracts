// Code generated from OpenSoha contracts. DO NOT EDIT.

package openapi

import _ "embed"

const (
	Version    = "0.1.19"
	YAMLSHA256 = "a1edda91ecc623423ca1875777dd500f12f239ca1571d7276a23d5e6c00c4c40"
	JSONSHA256 = "40f61307c9379b1496b43ff31fe8aebd82f1353100fe799f6d1ebbaefb1c3cd3"
)

//go:embed soha-api.yaml
var yamlSpec []byte

//go:embed soha-api.json
var jsonSpec []byte

// YAML returns a copy of the canonical OpenAPI YAML document.
func YAML() []byte { return append([]byte(nil), yamlSpec...) }

// JSON returns a copy of the canonical OpenAPI JSON document.
func JSON() []byte { return append([]byte(nil), jsonSpec...) }
