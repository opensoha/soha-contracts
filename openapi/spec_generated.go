// Code generated from OpenSoha contracts. DO NOT EDIT.

package openapi

import _ "embed"

const (
	Version    = "0.1.25"
	YAMLSHA256 = "45dbdf96918df285d1fe95687464e7cabae8d57eda503095cf8996d63705e9a1"
	JSONSHA256 = "f8f37ac29f7a065dc7a2e3738519cb270c854169d17483343443dc2f11988620"
)

//go:embed soha-api.yaml
var yamlSpec []byte

//go:embed soha-api.json
var jsonSpec []byte

// YAML returns a copy of the canonical OpenAPI YAML document.
func YAML() []byte { return append([]byte(nil), yamlSpec...) }

// JSON returns a copy of the canonical OpenAPI JSON document.
func JSON() []byte { return append([]byte(nil), jsonSpec...) }
