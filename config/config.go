package config

import (
	"log"
	"reflect"
)

// Config - Base configuration struct
type Config struct {
	APIPort         string          `yaml:"port" env:"API_PORT"`
	LinkStoreConfig LinkStoreConfig `yaml:"link_store_config"`
}

type LinkStoreConfig struct {
	Type string `yaml:"type" env:"LINK_STORE_TYPE"`
	DSN  string `yaml:"dsn" env:"LINK_STORE_DSN"`
}

// LoadConfig - Loads Config from file and then tries to load from Environment Variables, if there is any Environment variables
// they will override the one set using the file
func (cfg *Config) LoadConfig(cfgFilePath string) {
	// Load configuration from file and from environment
	if err := cfg.LoadConfigFile(cfgFilePath); err != nil {
		log.Println("WARN: Could not load config file")
	}
	cfg.LoadConfigEnv()
}

func (cfg *Config) setValidFields(config *Config) {
	configPtrValue := reflect.ValueOf(config)
	configValue := configPtrValue.Elem()
	cfgPtrValue := reflect.ValueOf(cfg)
	cfgValue := cfgPtrValue.Elem()

	for i := 0; i < configValue.NumField(); i++ {
		configField := configValue.Field(i)
		cfgField := cfgValue.Field(i)
		if cfgField.CanSet() {
			if configField.Kind() == reflect.String && configField.String() != "" {
				cfgField.SetString(configField.String())
			}
			if configField.Kind() == reflect.Bool {
				cfgField.SetBool(configField.Bool())
			}
		}
	}
}
