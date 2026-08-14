// Code generated from OpenSoha contracts. DO NOT EDIT.

package openapi

import _ "embed"

const (
	Version    = "0.1.10"
	YAMLSHA256 = "4d03aae5117d6ca06b928564b50ce6a131a7bdfec0cb1317e3c80c6416dea55a"
	JSONSHA256 = "f3957a681d1b8aa190e045cd97e3e15242d37051afbd6218be05983e4f5eb25a"
)

//go:embed soha-api.yaml
var yamlSpec []byte

//go:embed soha-api.json
var jsonSpec []byte

// YAML returns a copy of the canonical OpenAPI YAML document.
func YAML() []byte { return append([]byte(nil), yamlSpec...) }

// JSON returns a copy of the canonical OpenAPI JSON document.
func JSON() []byte { return append([]byte(nil), jsonSpec...) }
