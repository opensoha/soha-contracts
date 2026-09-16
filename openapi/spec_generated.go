// Code generated from OpenSoha contracts. DO NOT EDIT.

package openapi

import _ "embed"

const (
	Version    = "0.1.18"
	YAMLSHA256 = "d614ce858924d94da353c7a9380971a3e80964f3e6c865e8b27108e859500cf4"
	JSONSHA256 = "a98ea9ad99b6f044ce14e42f37b90b42e28347d53e37c727a04210b076aba960"
)

//go:embed soha-api.yaml
var yamlSpec []byte

//go:embed soha-api.json
var jsonSpec []byte

// YAML returns a copy of the canonical OpenAPI YAML document.
func YAML() []byte { return append([]byte(nil), yamlSpec...) }

// JSON returns a copy of the canonical OpenAPI JSON document.
func JSON() []byte { return append([]byte(nil), jsonSpec...) }
