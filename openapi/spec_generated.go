// Code generated from OpenSoha contracts. DO NOT EDIT.

package openapi

import _ "embed"

const (
	Version    = "0.1.25"
	YAMLSHA256 = "f6d13bc18f222449df720e54a51bdb741bda0678fdbeccbdbfc190e51fb93dbc"
	JSONSHA256 = "84d9cc1e51b475f662abbb002ceb6f94eb2fcd88686b05cc63aeea1981ec3daa"
)

//go:embed soha-api.yaml
var yamlSpec []byte

//go:embed soha-api.json
var jsonSpec []byte

// YAML returns a copy of the canonical OpenAPI YAML document.
func YAML() []byte { return append([]byte(nil), yamlSpec...) }

// JSON returns a copy of the canonical OpenAPI JSON document.
func JSON() []byte { return append([]byte(nil), jsonSpec...) }
