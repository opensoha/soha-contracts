// Code generated from OpenSoha contracts. DO NOT EDIT.

package openapi

import _ "embed"

const (
	Version    = "0.1.26"
	YAMLSHA256 = "f752a7e2c950d7c18be12c9ef079e138dc7232ad292d15d99d32953035d243ea"
	JSONSHA256 = "6bfd4a394ecb7304fe1c7665408b6e10a4da4793f6fafc7a137ed6fb7baa343b"
)

//go:embed soha-api.yaml
var yamlSpec []byte

//go:embed soha-api.json
var jsonSpec []byte

// YAML returns a copy of the canonical OpenAPI YAML document.
func YAML() []byte { return append([]byte(nil), yamlSpec...) }

// JSON returns a copy of the canonical OpenAPI JSON document.
func JSON() []byte { return append([]byte(nil), jsonSpec...) }
