// Code generated from OpenSoha contracts. DO NOT EDIT.

package openapi

import _ "embed"

const (
	Version    = "0.1.7"
	YAMLSHA256 = "e7075093dbff6e5886f6e8fd8f3d8d2c53734aee49ad0e9e84e7100d85f4eef3"
	JSONSHA256 = "53f23eb3644577ccc32f8c26562240eac3cb4f8b59595fa7e2a0a96ae046d03f"
)

//go:embed soha-api.yaml
var yamlSpec []byte

//go:embed soha-api.json
var jsonSpec []byte

// YAML returns a copy of the canonical OpenAPI YAML document.
func YAML() []byte { return append([]byte(nil), yamlSpec...) }

// JSON returns a copy of the canonical OpenAPI JSON document.
func JSON() []byte { return append([]byte(nil), jsonSpec...) }
