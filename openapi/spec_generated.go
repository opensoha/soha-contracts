// Code generated from OpenSoha contracts. DO NOT EDIT.

package openapi

import _ "embed"

const (
	Version    = "0.1.18"
	YAMLSHA256 = "002d20841b50eb4b543090877e6c12e01125e30276dc87bddad0a58d381a4ac3"
	JSONSHA256 = "87246533f9a1955e0c4f896dc41e190013da85e9973cfb9b48ce0567c254ee11"
)

//go:embed soha-api.yaml
var yamlSpec []byte

//go:embed soha-api.json
var jsonSpec []byte

// YAML returns a copy of the canonical OpenAPI YAML document.
func YAML() []byte { return append([]byte(nil), yamlSpec...) }

// JSON returns a copy of the canonical OpenAPI JSON document.
func JSON() []byte { return append([]byte(nil), jsonSpec...) }
