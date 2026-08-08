// Code generated from OpenSoha contracts. DO NOT EDIT.

package openapi

import _ "embed"

const (
	Version    = "0.1.8"
	YAMLSHA256 = "67feba4d2b47ae3982a9feb7a2286be4c856e0d29d63f445be165335acf364ef"
	JSONSHA256 = "dd0200945f41e4d87553ef88fdeb14577b219367768a43487ac3c23f5380eabb"
)

//go:embed soha-api.yaml
var yamlSpec []byte

//go:embed soha-api.json
var jsonSpec []byte

// YAML returns a copy of the canonical OpenAPI YAML document.
func YAML() []byte { return append([]byte(nil), yamlSpec...) }

// JSON returns a copy of the canonical OpenAPI JSON document.
func JSON() []byte { return append([]byte(nil), jsonSpec...) }
