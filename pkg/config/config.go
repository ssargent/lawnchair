package config

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type Program struct {
	Args        []string          `json:"args,omitempty" yaml:"args,omitempty" toml:"args,omitempty"`
	Description string            `json:"description,omitempty" yaml:"description,omitempty" toml:"description,omitempty"`
	EnvVars     map[string]string `json:"env_vars,omitempty" yaml:"env_vars,omitempty" toml:"env_vars,omitempty"`
	Name        string            `json:"name" yaml:"name" toml:"name"`
	Path        string            `json:"path,omitempty" yaml:"path,omitempty" toml:"path,omitempty"`
}

type Config struct {
	Programs map[string]Program `json:"programs" yaml:"programs" toml:"programs"`
}

func load(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Read all the content into a byte slice
	content, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	c := &Config{}
	// Convert the byte slice to a string
	configTxt := string(content)

	if filepath.Ext(path) == ".yaml" || filepath.Ext(path) == ".yml" {
		if err := yaml.Unmarshal([]byte(configTxt), &c); err != nil {
			return nil, fmt.Errorf("load: %w", err)
		}
	} else if filepath.Ext(path) == ".json" {
		if err := json.Unmarshal([]byte(configTxt), &c); err != nil {
			return nil, fmt.Errorf("load: %w", err)
		}
	}

	return c, nil
}

func New(path string) (*Config, error) {
	return load(path)
}

func (c *Config) AsYaml() (string, error) {
	out, err := yaml.Marshal(c)
	if err != nil {
		return "", fmt.Errorf("AsYaml: %w", err)
	}
	return string(out), nil
}

func (c *Config) AsJson() (string, error) {
	out, err := json.Marshal(c)
	if err != nil {
		return "", fmt.Errorf("AsJson: %w", err)
	}
	return string(out), nil
}
