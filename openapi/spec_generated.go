// Code generated from OpenSoha contracts. DO NOT EDIT.

package openapi

import _ "embed"

const (
	Version    = "0.1.4"
	YAMLSHA256 = "09e3b9d0e56d0fa64dec6ca3a631a15cea6c8837fdcd1efe6866d9f2b7163e1a"
	JSONSHA256 = "677e1cf4de13bb9888f4af32c773ff47c8937312bcf04da764477b7b3c711aab"
)

//go:embed soha-api.yaml
var yamlSpec []byte

//go:embed soha-api.json
var jsonSpec []byte

// YAML returns a copy of the canonical OpenAPI YAML document.
func YAML() []byte { return append([]byte(nil), yamlSpec...) }

// JSON returns a copy of the canonical OpenAPI JSON document.
func JSON() []byte { return append([]byte(nil), jsonSpec...) }
