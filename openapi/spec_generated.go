// Code generated from OpenSoha contracts. DO NOT EDIT.

package openapi

import _ "embed"

const (
	Version    = "0.1.12"
	YAMLSHA256 = "b9dbd5597000f9124790f0df8e9e7616e76f7d7e12df03b72490c3403da4eda5"
	JSONSHA256 = "6f34774013b3e6c9dd0e2c9b3e9d929a68f8f8b659a5909dad64a77821ef24dd"
)

//go:embed soha-api.yaml
var yamlSpec []byte

//go:embed soha-api.json
var jsonSpec []byte

// YAML returns a copy of the canonical OpenAPI YAML document.
func YAML() []byte { return append([]byte(nil), yamlSpec...) }

// JSON returns a copy of the canonical OpenAPI JSON document.
func JSON() []byte { return append([]byte(nil), jsonSpec...) }
