// Code generated from OpenSoha contracts. DO NOT EDIT.

package openapi

import _ "embed"

const (
	Version    = "0.1.16"
	YAMLSHA256 = "1e34ae057d295a8ffa0b13beb54a8a2a1998b1fed5f949cb57d13360914c4ee0"
	JSONSHA256 = "ca3886e57c2dd1dca5a7dc10c0d1041394a8118cb901f783ebf28eadd0c90f15"
)

//go:embed soha-api.yaml
var yamlSpec []byte

//go:embed soha-api.json
var jsonSpec []byte

// YAML returns a copy of the canonical OpenAPI YAML document.
func YAML() []byte { return append([]byte(nil), yamlSpec...) }

// JSON returns a copy of the canonical OpenAPI JSON document.
func JSON() []byte { return append([]byte(nil), jsonSpec...) }
