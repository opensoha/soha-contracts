// Code generated from OpenSoha contracts. DO NOT EDIT.

package openapi

import _ "embed"

const (
	Version    = "0.1.6"
	YAMLSHA256 = "5cd3742695a0bc9ee14c37177431522ca9ea6eac50e87c8b73f3349e635ce569"
	JSONSHA256 = "a1adce76608654fd7cf188c86ca4912ad7f1ace01ea15c1c191ea0a94ddfbc93"
)

//go:embed soha-api.yaml
var yamlSpec []byte

//go:embed soha-api.json
var jsonSpec []byte

// YAML returns a copy of the canonical OpenAPI YAML document.
func YAML() []byte { return append([]byte(nil), yamlSpec...) }

// JSON returns a copy of the canonical OpenAPI JSON document.
func JSON() []byte { return append([]byte(nil), jsonSpec...) }
