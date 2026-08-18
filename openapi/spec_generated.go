// Code generated from OpenSoha contracts. DO NOT EDIT.

package openapi

import _ "embed"

const (
	Version    = "0.1.14"
	YAMLSHA256 = "231a8b25b830df7c602c1356494502235ca8c00f23a4aea56c87c0fd79438663"
	JSONSHA256 = "35983758f2bf5d2b4c2e918ee32f34f384218d882edf42d86c310f2820e2a82f"
)

//go:embed soha-api.yaml
var yamlSpec []byte

//go:embed soha-api.json
var jsonSpec []byte

// YAML returns a copy of the canonical OpenAPI YAML document.
func YAML() []byte { return append([]byte(nil), yamlSpec...) }

// JSON returns a copy of the canonical OpenAPI JSON document.
func JSON() []byte { return append([]byte(nil), jsonSpec...) }
