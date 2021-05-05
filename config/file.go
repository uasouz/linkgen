package config

import (
	"gopkg.in/yaml.v2"
	"os"
)

func (cfg *Config) LoadConfigFile(filepath string) error {
	// Check if file exists
	_, err := os.Stat(filepath)
	if err != nil {
		return err
	}

	// Try opening file
	file, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	var config Config

	err = yaml.NewDecoder(file).Decode(&config)
	if err != nil {
		return err
	}

	if cfg != nil {
		cfg.setValidFields(&config)
	}

	return nil
}
