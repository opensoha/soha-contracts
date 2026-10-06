// Code generated from OpenSoha contracts. DO NOT EDIT.

package openapi

import _ "embed"

const (
	Version    = "0.1.21"
	YAMLSHA256 = "4f6a18c89a70f8ce91f27eb8d52eb49a7f5e4db38dc7de27eaa0efceac1c33fe"
	JSONSHA256 = "5e3ba8e83879359042b4df47658cb2df7e67a7bcc941ab2507155a1555c93d25"
)

//go:embed soha-api.yaml
var yamlSpec []byte

//go:embed soha-api.json
var jsonSpec []byte

// YAML returns a copy of the canonical OpenAPI YAML document.
func YAML() []byte { return append([]byte(nil), yamlSpec...) }

// JSON returns a copy of the canonical OpenAPI JSON document.
func JSON() []byte { return append([]byte(nil), jsonSpec...) }
