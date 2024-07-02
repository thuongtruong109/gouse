package config

import "testing"

func TestToml(t *testing.T) {
	type Meta struct {
		Lang string `toml:"lang"`
		Type string `toml:"type"`
	}

	type configuration struct {
		Host string `toml:"host"`
		Port int    `toml:"port"`
		Meta Meta   `toml:"meta"`
	}

	var conf configuration
	if err := Toml("../testdata/config.toml", &conf); err != nil {
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
