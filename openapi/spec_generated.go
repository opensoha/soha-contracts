// Code generated from OpenSoha contracts. DO NOT EDIT.

package openapi

import _ "embed"

const (
	Version    = "0.1.5"
	YAMLSHA256 = "d841f66bab9b28db25f34f8c4b7d7047fba018e82c95acc17b93ff2e76e6d1f3"
	JSONSHA256 = "56df00cdf416382845496bd48cb4ab979accc666fde6481366bc9685a7f70f18"
)

//go:embed soha-api.yaml
var yamlSpec []byte

//go:embed soha-api.json
var jsonSpec []byte

// YAML returns a copy of the canonical OpenAPI YAML document.
func YAML() []byte { return append([]byte(nil), yamlSpec...) }

// JSON returns a copy of the canonical OpenAPI JSON document.
func JSON() []byte { return append([]byte(nil), jsonSpec...) }
