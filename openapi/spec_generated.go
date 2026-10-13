// Code generated from OpenSoha contracts. DO NOT EDIT.

package openapi

import _ "embed"

const (
	Version    = "0.1.23"
	YAMLSHA256 = "461582ddccf2a255645d856611188848d087a08c000af19222febd3e00e515be"
	JSONSHA256 = "ae41db381eab5bba39e1dced1b4f6a9ee775937a901ae39d06eed15db2629ea8"
)

//go:embed soha-api.yaml
var yamlSpec []byte

//go:embed soha-api.json
var jsonSpec []byte

// YAML returns a copy of the canonical OpenAPI YAML document.
func YAML() []byte { return append([]byte(nil), yamlSpec...) }

// JSON returns a copy of the canonical OpenAPI JSON document.
func JSON() []byte { return append([]byte(nil), jsonSpec...) }
