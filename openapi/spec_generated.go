// Code generated from OpenSoha contracts. DO NOT EDIT.

package openapi

import _ "embed"

const (
	Version    = "0.1.17"
	YAMLSHA256 = "43c055eeb546b837a15cdd39bacd5c086d5257bfccc8f9fc137943200e4ea72e"
	JSONSHA256 = "11f96a23fa4302d4518af62b39446ff48204e97c4b9d07e561d81124c2a5e754"
)

//go:embed soha-api.yaml
var yamlSpec []byte

//go:embed soha-api.json
var jsonSpec []byte

// YAML returns a copy of the canonical OpenAPI YAML document.
func YAML() []byte { return append([]byte(nil), yamlSpec...) }

// JSON returns a copy of the canonical OpenAPI JSON document.
func JSON() []byte { return append([]byte(nil), jsonSpec...) }
