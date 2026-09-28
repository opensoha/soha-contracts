// Code generated from OpenSoha contracts. DO NOT EDIT.

package openapi

import _ "embed"

const (
	Version    = "0.1.20"
	YAMLSHA256 = "cbf0b3c1f47da982289c43ccdc76c5eb9e2c8d064aebfe92aa23e4b179ea136b"
	JSONSHA256 = "6862030a09d680756adc6bff8cc1e6bf81318704bf957027f22e1b7655bb471e"
)

//go:embed soha-api.yaml
var yamlSpec []byte

//go:embed soha-api.json
var jsonSpec []byte

// YAML returns a copy of the canonical OpenAPI YAML document.
func YAML() []byte { return append([]byte(nil), yamlSpec...) }

// JSON returns a copy of the canonical OpenAPI JSON document.
func JSON() []byte { return append([]byte(nil), jsonSpec...) }
