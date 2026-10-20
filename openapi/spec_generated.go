// Code generated from OpenSoha contracts. DO NOT EDIT.

package openapi

import _ "embed"

const (
	Version    = "0.1.25"
	YAMLSHA256 = "5f72e3a41fbf5edf22773a9b85d2841c87c99a0636ad704065b9b1333414bd02"
	JSONSHA256 = "60341d8b2940687e94a3ed239363761e8e50cf36824fda11f344ba6363a17c04"
)

//go:embed soha-api.yaml
var yamlSpec []byte

//go:embed soha-api.json
var jsonSpec []byte

// YAML returns a copy of the canonical OpenAPI YAML document.
func YAML() []byte { return append([]byte(nil), yamlSpec...) }

// JSON returns a copy of the canonical OpenAPI JSON document.
func JSON() []byte { return append([]byte(nil), jsonSpec...) }
