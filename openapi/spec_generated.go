// Code generated from OpenSoha contracts. DO NOT EDIT.

package openapi

import _ "embed"

const (
	Version    = "0.1.16"
	YAMLSHA256 = "8bd8168b786a1cddc1be5cf729efb1bbaa8b4b236e38666d722d2eaec7b23210"
	JSONSHA256 = "2c914e916f3106347080bc4ee3c75736e35b3b4c28b2f310e8bfa3a4ef98a4dc"
)

//go:embed soha-api.yaml
var yamlSpec []byte

//go:embed soha-api.json
var jsonSpec []byte

// YAML returns a copy of the canonical OpenAPI YAML document.
func YAML() []byte { return append([]byte(nil), yamlSpec...) }

// JSON returns a copy of the canonical OpenAPI JSON document.
func JSON() []byte { return append([]byte(nil), jsonSpec...) }
