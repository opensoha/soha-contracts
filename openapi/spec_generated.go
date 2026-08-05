// Code generated from OpenSoha contracts. DO NOT EDIT.

package openapi

import _ "embed"

const (
	Version    = "0.1.6"
	YAMLSHA256 = "b6388f34cd559f508a70ccdfdc3a01bb33b62c1b5c4631ec8339770a9b1c338a"
	JSONSHA256 = "bff1d46d3f3a083c7a2de06c916b2dc42e5699eeaf4fed2c3e560bf345427985"
)

//go:embed soha-api.yaml
var yamlSpec []byte

//go:embed soha-api.json
var jsonSpec []byte

// YAML returns a copy of the canonical OpenAPI YAML document.
func YAML() []byte { return append([]byte(nil), yamlSpec...) }

// JSON returns a copy of the canonical OpenAPI JSON document.
func JSON() []byte { return append([]byte(nil), jsonSpec...) }
