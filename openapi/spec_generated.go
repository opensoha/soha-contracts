// Code generated from OpenSoha contracts. DO NOT EDIT.

package openapi

import _ "embed"

const (
	Version    = "0.1.25"
	YAMLSHA256 = "8c8746c6c38e220ab0083d7f26a369ccd56c1ba6246cd9d4f79c681cbaa28daa"
	JSONSHA256 = "4df75edaf6d99fa08ad67c43d0fee64403c0d1370459e728d688acf2d3142a48"
)

//go:embed soha-api.yaml
var yamlSpec []byte

//go:embed soha-api.json
var jsonSpec []byte

// YAML returns a copy of the canonical OpenAPI YAML document.
func YAML() []byte { return append([]byte(nil), yamlSpec...) }

// JSON returns a copy of the canonical OpenAPI JSON document.
func JSON() []byte { return append([]byte(nil), jsonSpec...) }
