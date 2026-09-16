// Code generated from OpenSoha contracts. DO NOT EDIT.

package openapi

import _ "embed"

const (
	Version    = "0.1.18"
	YAMLSHA256 = "b5d13459fa1732e0cc820708a25cd2354e2a2011a02e1e010e3eeec8e74a8731"
	JSONSHA256 = "1c1a07a168dce1fdff663c35c3b1384c310481fb09cfea4f222ec0ce35a2f2fd"
)

//go:embed soha-api.yaml
var yamlSpec []byte

//go:embed soha-api.json
var jsonSpec []byte

// YAML returns a copy of the canonical OpenAPI YAML document.
func YAML() []byte { return append([]byte(nil), yamlSpec...) }

// JSON returns a copy of the canonical OpenAPI JSON document.
func JSON() []byte { return append([]byte(nil), jsonSpec...) }
