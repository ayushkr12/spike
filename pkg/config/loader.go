package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

var Cfg *Config

// LoadConfig loads config from given YAML file path
func LoadConfig(path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("failed to read config: %w", err)
	}
	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return fmt.Errorf("failed to parse config: %w", err)
	}
	Cfg = &cfg
	return nil
}

// SaveDefaultConfig writes the default config to a YAML file
func SaveDefaultConfig(path string) error {
	data, err := yaml.Marshal(DefaultConfig())
	if err != nil {
		return fmt.Errorf("failed to marshal default config: %w", err)
	}
	if err := os.WriteFile(path, data, 0644); err != nil {
		return fmt.Errorf("failed to write default config: %w", err)
	}
	return nil
}
