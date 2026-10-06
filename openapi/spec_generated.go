// Code generated from OpenSoha contracts. DO NOT EDIT.

package openapi

import _ "embed"

const (
	Version    = "0.1.22"
	YAMLSHA256 = "6670ee53d04f37094dab2be377d6d4a49985e2922dc9f0e69f31ca070b63dfbb"
	JSONSHA256 = "fd239c6921505fa065fac7321fd0d1705f515a4b11b40b60811c6f0b8403404e"
)

//go:embed soha-api.yaml
var yamlSpec []byte

//go:embed soha-api.json
var jsonSpec []byte

// YAML returns a copy of the canonical OpenAPI YAML document.
func YAML() []byte { return append([]byte(nil), yamlSpec...) }

// JSON returns a copy of the canonical OpenAPI JSON document.
func JSON() []byte { return append([]byte(nil), jsonSpec...) }
