// Code generated from OpenSoha contracts. DO NOT EDIT.

package openapi

import _ "embed"

const (
	Version    = "0.1.9"
	YAMLSHA256 = "7639ee05b3a0f87467bdd29f322673bbbed29f1a6b357f7758d302cdc3a552bd"
	JSONSHA256 = "edc994a5c25e2b82d3226ab94bfacbae21017e023be166b862733d6d499f6a71"
)

//go:embed soha-api.yaml
var yamlSpec []byte

//go:embed soha-api.json
var jsonSpec []byte

// YAML returns a copy of the canonical OpenAPI YAML document.
func YAML() []byte { return append([]byte(nil), yamlSpec...) }

// JSON returns a copy of the canonical OpenAPI JSON document.
func JSON() []byte { return append([]byte(nil), jsonSpec...) }
