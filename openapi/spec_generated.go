// Code generated from OpenSoha contracts. DO NOT EDIT.

package openapi

import _ "embed"

const (
	Version    = "0.1.15"
	YAMLSHA256 = "5de407feba0f68943fbd03103e78da8124177014a400fe483e1d9986312fb247"
	JSONSHA256 = "be2e066a7fd4eec74f5c2d240860fdb7a343ff9a3fb25456982ecaea7eb355c3"
)

//go:embed soha-api.yaml
var yamlSpec []byte

//go:embed soha-api.json
var jsonSpec []byte

// YAML returns a copy of the canonical OpenAPI YAML document.
func YAML() []byte { return append([]byte(nil), yamlSpec...) }

// JSON returns a copy of the canonical OpenAPI JSON document.
func JSON() []byte { return append([]byte(nil), jsonSpec...) }
