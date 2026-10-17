// Code generated from OpenSoha contracts. DO NOT EDIT.

package openapi

import _ "embed"

const (
	Version    = "0.1.24"
	YAMLSHA256 = "e1a2f7be61a7d7674f3e3928190a04ddf7ecb906170c5da9199519b50bc89616"
	JSONSHA256 = "cf9f707f6e35f29486eb7a190742124d27c2db82915dd98b192af5aa4a587596"
)

//go:embed soha-api.yaml
var yamlSpec []byte

//go:embed soha-api.json
var jsonSpec []byte

// YAML returns a copy of the canonical OpenAPI YAML document.
func YAML() []byte { return append([]byte(nil), yamlSpec...) }

// JSON returns a copy of the canonical OpenAPI JSON document.
func JSON() []byte { return append([]byte(nil), jsonSpec...) }
