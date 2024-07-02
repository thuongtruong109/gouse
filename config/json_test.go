package config

import "testing"

func TestJson(t *testing.T) {
	type Meta struct {
		Lang string `json:"lang"`
		Type string `json:"type"`
	}

	type configuration struct {
		Host string `json:"host"`
		Port int    `json:"port"`
		Meta Meta   `json:"meta"`
	}

	var conf configuration
	if err := Json("../testdata/config.json", &conf); err != nil {
		t.Fatal(err)
	}

	if conf.Host != "localhost" {
		t.Fatalf("expected host to be 'localhost', got %q", conf.Host)
	}

	if conf.Port != 8080 {
		t.Fatalf("expected port to be 8080, got %d", conf.Port)
	}

	if conf.Meta.Lang != "golang" {
		t.Fatalf("expected meta.lang to be 'golang', got %q", conf.Meta.Lang)
	}

	if conf.Meta.Type != "package" {
		t.Fatalf("expected meta.type to be 'package', got %q", conf.Meta.Type)
	}
}