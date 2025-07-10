package gouse

import "testing"

func TestReadJSON(t *testing.T) {
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
	if err := ReadJSON("./mockdata/config.json", &conf); err != nil {
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

func TestReadTOML(t *testing.T) {
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
	if err := ReadTOML("./mockdata/config.toml", &conf); err != nil {
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

func TestReadYAML(t *testing.T) {
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
	if err := ReadYAML("./mockdata/config.yaml", &conf); err != nil {
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

func TestReadCSV(t *testing.T) {
	records, err := ReadCSV("mockdata/config.csv")
	if err != nil {
		t.Fatalf("failed to read CSV: %v", err)
	}

	if len(records) != 3 {
		t.Fatalf("expected 3 rows (including header), got %d", len(records))
	}

	if records[1][0] != "Alice" {
		t.Errorf("expected 'Alice', got '%s'", records[1][0])
	}

	if records[2][2] != "bob@example.com" {
		t.Errorf("expected 'bob@example.com', got '%s'", records[2][2])
	}
}
