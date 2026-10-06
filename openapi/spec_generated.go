// Code generated from OpenSoha contracts. DO NOT EDIT.

package openapi

import _ "embed"

const (
	Version    = "0.1.22"
	YAMLSHA256 = "5b0140f7820051bcbce9d2b2c0cd875b087a308f71dfdf1701d9052942c464ff"
	JSONSHA256 = "e120a0bb9308d25baac4778d8535b11ea29bb5558ce2eebe35bde87c1c2f2724"
)

//go:embed soha-api.yaml
var yamlSpec []byte

//go:embed soha-api.json
var jsonSpec []byte

// YAML returns a copy of the canonical OpenAPI YAML document.
func YAML() []byte { return append([]byte(nil), yamlSpec...) }

// JSON returns a copy of the canonical OpenAPI JSON document.
func JSON() []byte { return append([]byte(nil), jsonSpec...) }
