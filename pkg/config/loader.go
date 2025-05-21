package config

import (
	"fmt"
	"os"

	log "github.com/ayushkr12/logger"
	"gopkg.in/yaml.v3"
)

var Cfg *Config

// LoadConfig loads config from given YAML file path
func LoadConfig(path string) error {
	_, err := os.Stat(path)
	if err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("failed to check config file: %w", err)
	} else if os.IsNotExist(err) {
		log.Infof("Config file not found, creating default config at %s", path)
		if err := SaveDefaultConfig(path); err != nil {
			return fmt.Errorf("failed to create default config: %w", err)
		}
	}
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
