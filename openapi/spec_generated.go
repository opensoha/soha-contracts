// Code generated from OpenSoha contracts. DO NOT EDIT.

package openapi

import _ "embed"

const (
	Version    = "0.1.16"
	YAMLSHA256 = "bdcd23df5b8d2915e371e81ed9915b9ac30dad08c334e02befd26bd3c407fe15"
	JSONSHA256 = "0bf54dc8e32263490aacab9489ca04804a218618ad52d6fb95561835d3d81cf7"
)

//go:embed soha-api.yaml
var yamlSpec []byte

//go:embed soha-api.json
var jsonSpec []byte

// YAML returns a copy of the canonical OpenAPI YAML document.
func YAML() []byte { return append([]byte(nil), yamlSpec...) }

// JSON returns a copy of the canonical OpenAPI JSON document.
func JSON() []byte { return append([]byte(nil), jsonSpec...) }
