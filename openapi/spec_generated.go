// Code generated from OpenSoha contracts. DO NOT EDIT.

package openapi

import _ "embed"

const (
	Version    = "0.1.16"
	YAMLSHA256 = "7a62fe9c9020f52b935d3c44abcc038c1bb5b3bd3b83eaee6990741108bf731a"
	JSONSHA256 = "6498b2aa19fd2db09d11faf3f85023cd57f24aeb8287502d3bf92b2bcce7a769"
)

//go:embed soha-api.yaml
var yamlSpec []byte

//go:embed soha-api.json
var jsonSpec []byte

// YAML returns a copy of the canonical OpenAPI YAML document.
func YAML() []byte { return append([]byte(nil), yamlSpec...) }

// JSON returns a copy of the canonical OpenAPI JSON document.
func JSON() []byte { return append([]byte(nil), jsonSpec...) }
