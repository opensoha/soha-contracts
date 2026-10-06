// Code generated from OpenSoha contracts. DO NOT EDIT.

package openapi

import _ "embed"

const (
	Version    = "0.1.23"
	YAMLSHA256 = "b919b001e6da4b523d8f30be986bf40453ddd9c5d94126246c8495f1d0e961a3"
	JSONSHA256 = "bf3e1784f7e9398c9215732407ad4af11c928421805ef2fa7df900dd9a48167c"
)

//go:embed soha-api.yaml
var yamlSpec []byte

//go:embed soha-api.json
var jsonSpec []byte

// YAML returns a copy of the canonical OpenAPI YAML document.
func YAML() []byte { return append([]byte(nil), yamlSpec...) }

// JSON returns a copy of the canonical OpenAPI JSON document.
func JSON() []byte { return append([]byte(nil), jsonSpec...) }
