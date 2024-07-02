package config

import "testing"

func TestYaml(t *testing.T) {
	type Meta struct {
		Lang string `yaml:"lang"`
		Type string `yaml:"type"`
	}

	type Configuration struct {
		Host string `yaml:"host"`
		Port int    `yaml:"port"`
		Meta Meta   `yaml:"meta"`
	}

	var conf Configuration
	if err := Yaml("../testdata/config.yaml", &conf); err != nil {
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