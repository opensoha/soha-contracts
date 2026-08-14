// Code generated from OpenSoha contracts. DO NOT EDIT.

package openapi

import _ "embed"

const (
	Version    = "0.1.11"
	YAMLSHA256 = "84fc6e95a10bbbf20bc8b2adbda86d824bf664d1ca16c879c37cf1005ed47557"
	JSONSHA256 = "551abf1b2ab9a56f0b59de7d43d1d59c0145d88cf8a017e1c0a937441947fdff"
)

//go:embed soha-api.yaml
var yamlSpec []byte

//go:embed soha-api.json
var jsonSpec []byte

// YAML returns a copy of the canonical OpenAPI YAML document.
func YAML() []byte { return append([]byte(nil), yamlSpec...) }

// JSON returns a copy of the canonical OpenAPI JSON document.
func JSON() []byte { return append([]byte(nil), jsonSpec...) }
