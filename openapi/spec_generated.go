// Code generated from OpenSoha contracts. DO NOT EDIT.

package openapi

import _ "embed"

const (
	Version    = "0.1.13"
	YAMLSHA256 = "480571290460492a4f422ecf2c2098720989c31a4413c1bf29fc45b1ecaa5c7b"
	JSONSHA256 = "8327233124b473bdfe2a57582eb66ed9c844630bedd0c947eefb707b71f7d3cf"
)

//go:embed soha-api.yaml
var yamlSpec []byte

//go:embed soha-api.json
var jsonSpec []byte

// YAML returns a copy of the canonical OpenAPI YAML document.
func YAML() []byte { return append([]byte(nil), yamlSpec...) }

// JSON returns a copy of the canonical OpenAPI JSON document.
func JSON() []byte { return append([]byte(nil), jsonSpec...) }
