// Code generated from OpenSoha contracts. DO NOT EDIT.

package openapi

import _ "embed"

const (
	Version    = "0.1.14"
	YAMLSHA256 = "beda5dec0bfc417e6b0b2fbbce76c947fd0913e1f981045234bb9fb50757e3f4"
	JSONSHA256 = "933e16a468318ffafbf55e8700208fb11b8c12bac00dbd6b090b3cca1bcd8e69"
)

//go:embed soha-api.yaml
var yamlSpec []byte

//go:embed soha-api.json
var jsonSpec []byte

// YAML returns a copy of the canonical OpenAPI YAML document.
func YAML() []byte { return append([]byte(nil), yamlSpec...) }

// JSON returns a copy of the canonical OpenAPI JSON document.
func JSON() []byte { return append([]byte(nil), jsonSpec...) }
