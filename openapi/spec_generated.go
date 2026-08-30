// Code generated from OpenSoha contracts. DO NOT EDIT.

package openapi

import _ "embed"

const (
	Version    = "0.1.16"
	YAMLSHA256 = "086073240d6cc038c66a9c8dbe200e9f285a3b85b6bec3d7eecb0b6ccc4ec888"
	JSONSHA256 = "b7773ae196b69bc2076f48bff41d06879d511d9ce43429ca05cebe31c795e6c4"
)

//go:embed soha-api.yaml
var yamlSpec []byte

//go:embed soha-api.json
var jsonSpec []byte

// YAML returns a copy of the canonical OpenAPI YAML document.
func YAML() []byte { return append([]byte(nil), yamlSpec...) }

// JSON returns a copy of the canonical OpenAPI JSON document.
func JSON() []byte { return append([]byte(nil), jsonSpec...) }
