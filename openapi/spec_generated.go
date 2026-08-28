// Code generated from OpenSoha contracts. DO NOT EDIT.

package openapi

import _ "embed"

const (
	Version    = "0.1.16"
	YAMLSHA256 = "d9d8645ee5914df5dfaac8b37c2bc92ade27a094cd0c2c1374cfe4f3fd1da664"
	JSONSHA256 = "3a47831135bb3ba3e0b6a9d8d9c2676717a332d9e5ef00677b132b59f0b48521"
)

//go:embed soha-api.yaml
var yamlSpec []byte

//go:embed soha-api.json
var jsonSpec []byte

// YAML returns a copy of the canonical OpenAPI YAML document.
func YAML() []byte { return append([]byte(nil), yamlSpec...) }

// JSON returns a copy of the canonical OpenAPI JSON document.
func JSON() []byte { return append([]byte(nil), jsonSpec...) }
