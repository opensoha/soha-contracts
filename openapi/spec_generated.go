// Code generated from OpenSoha contracts. DO NOT EDIT.

package openapi

import _ "embed"

const (
	Version    = "0.1.7"
	YAMLSHA256 = "1342c9f391d02d91c336d27c65ae81781f4432cbc5ffc47ea53fad9c688adb0a"
	JSONSHA256 = "1e028d64f794045e22d447b4609b53f9846324fa78ca88818493b6ee80a3b0e1"
)

//go:embed soha-api.yaml
var yamlSpec []byte

//go:embed soha-api.json
var jsonSpec []byte

// YAML returns a copy of the canonical OpenAPI YAML document.
func YAML() []byte { return append([]byte(nil), yamlSpec...) }

// JSON returns a copy of the canonical OpenAPI JSON document.
func JSON() []byte { return append([]byte(nil), jsonSpec...) }
