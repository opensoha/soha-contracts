package openapi

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"strings"
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

func TestTopologySuccessResponsePreservesDescription(t *testing.T) {
	var document struct {
		Paths map[string]struct {
			Get struct {
				Responses map[string]map[string]json.RawMessage `json:"responses"`
			} `json:"get"`
		} `json:"paths"`
	}
	if err := json.Unmarshal(JSON(), &document); err != nil {
		t.Fatal(err)
	}
	response := document.Paths["/clusters/{clusterID}/network/topology"].Get.Responses["200"]
	if len(response) == 0 {
		t.Fatal("missing topology success response")
	}
	for key := range response {
		switch key {
		case "description", "headers", "content", "links":
		default:
			if !strings.HasPrefix(key, "x-") {
				t.Errorf("unexpected topology response field %q", key)
			}
		}
	}
	var description string
	if err := json.Unmarshal(response["description"], &description); err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(description, "Kubernetes, Gateway API, Traefik and APISIX topology,") {
		t.Fatalf("topology response description was truncated: %q", description)
	}
}
