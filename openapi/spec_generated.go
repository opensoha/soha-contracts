// Code generated from OpenSoha contracts. DO NOT EDIT.

package openapi

import _ "embed"

const (
	Version    = "0.1.27"
	YAMLSHA256 = "d610eecc45aedba28067f57c7b3f277d60f27c8a95f0231f34fa248264b04a84"
	JSONSHA256 = "cf916820bdf35425a32e1ad78f69d411622a4bc194384c037a542b147597448b"
)

//go:embed soha-api.yaml
var yamlSpec []byte

//go:embed soha-api.json
var jsonSpec []byte

// YAML returns a copy of the canonical OpenAPI YAML document.
func YAML() []byte { return append([]byte(nil), yamlSpec...) }

// JSON returns a copy of the canonical OpenAPI JSON document.
func JSON() []byte { return append([]byte(nil), jsonSpec...) }
