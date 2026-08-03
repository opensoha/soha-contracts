package openapi

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"testing"
)

func TestEmbeddedSpecsMatchMetadata(t *testing.T) {
	for _, test := range []struct {
		name string
		body []byte
		want string
	}{
		{name: "yaml", body: YAML(), want: YAMLSHA256},
		{name: "json", body: JSON(), want: JSONSHA256},
	} {
		t.Run(test.name, func(t *testing.T) {
			got := fmt.Sprintf("%x", sha256.Sum256(test.body))
			if got != test.want {
				t.Fatalf("sha256 = %s, want %s", got, test.want)
			}
		})
	}

	var document struct {
		Info struct {
			Version string `json:"version"`
		} `json:"info"`
	}
	if err := json.Unmarshal(JSON(), &document); err != nil {
		t.Fatalf("decode embedded JSON: %v", err)
	}
	if document.Info.Version != Version {
		t.Fatalf("info.version = %q, want %q", document.Info.Version, Version)
	}
}
