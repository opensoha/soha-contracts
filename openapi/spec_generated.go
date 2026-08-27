// Code generated from OpenSoha contracts. DO NOT EDIT.

package openapi

import _ "embed"

const (
	Version    = "0.1.15"
	YAMLSHA256 = "4d0ae8a471e140f7989f28396efae6ce3e8b6f55a0c38ba8b6de955117cf06bc"
	JSONSHA256 = "50c9347444fc95f2e744e8573737285cfda8d5b495009de9f740b5e13bf8d083"
)

//go:embed soha-api.yaml
var yamlSpec []byte

//go:embed soha-api.json
var jsonSpec []byte

// YAML returns a copy of the canonical OpenAPI YAML document.
func YAML() []byte { return append([]byte(nil), yamlSpec...) }

// JSON returns a copy of the canonical OpenAPI JSON document.
func JSON() []byte { return append([]byte(nil), jsonSpec...) }
