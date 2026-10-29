// Code generated from OpenSoha contracts. DO NOT EDIT.

package openapi

import _ "embed"

const (
	Version    = "0.1.27"
	YAMLSHA256 = "7ec8930963123f5b79dff6dfaddf7200dd0fce51ada6212d47f415d4e7539727"
	JSONSHA256 = "6c1d5c0d4871c5ff6ea82241a02ff8ee4147792c6026a474a9d10a3ef3ca051b"
)

//go:embed soha-api.yaml
var yamlSpec []byte

//go:embed soha-api.json
var jsonSpec []byte

// YAML returns a copy of the canonical OpenAPI YAML document.
func YAML() []byte { return append([]byte(nil), yamlSpec...) }

// JSON returns a copy of the canonical OpenAPI JSON document.
func JSON() []byte { return append([]byte(nil), jsonSpec...) }
