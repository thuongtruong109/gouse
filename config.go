package gouse

import (
	"encoding/csv"
	"encoding/json"
	"os"

	"github.com/kelseyhightower/envconfig"
	"github.com/pelletier/go-toml"
	"gopkg.in/yaml.v2"
)

func ReadJSON[T any](path string, configuration *T) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer func() {
		if err := file.Close(); err != nil {
			panic(err)
		}
	}()

	decoder := json.NewDecoder(file)
	return decoder.Decode(configuration)
}

func ReadTOML[T any](path string, configuration *T) error {
	data, err := toml.LoadFile(path)
	if err != nil {
		return err
	}

	err = data.Unmarshal(configuration)
	if err != nil {
		return err
	}

	return nil
}

func ReadYAML[T any](path string, configuration *T) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}

	defer func() {
		if err := f.Close(); err != nil {
			panic(err)
		}
	}()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(configuration)
	if err != nil {
		return err
	}

	err = envconfig.Process("", configuration)
	if err != nil {
		return err
	}

	return nil
}

func ReadCSV(path string) ([][]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := file.Close(); err != nil {
			panic(err)
		}
	}()

	reader := csv.NewReader(file)
	return reader.ReadAll()
}
