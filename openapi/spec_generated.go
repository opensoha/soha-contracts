// Code generated from OpenSoha contracts. DO NOT EDIT.

package openapi

import _ "embed"

const (
	Version    = "0.1.25"
	YAMLSHA256 = "433f9d7c65aa4c74d321bf4de79af649686aa56744a220995906f778d5515267"
	JSONSHA256 = "864488d29260944844b36b1857fb9801e448a0f389f948b4b8b7690d87e0e33f"
)

//go:embed soha-api.yaml
var yamlSpec []byte

//go:embed soha-api.json
var jsonSpec []byte

// YAML returns a copy of the canonical OpenAPI YAML document.
func YAML() []byte { return append([]byte(nil), yamlSpec...) }

// JSON returns a copy of the canonical OpenAPI JSON document.
func JSON() []byte { return append([]byte(nil), jsonSpec...) }
