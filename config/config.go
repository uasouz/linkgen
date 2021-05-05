package config

import (
	"fmt"
	"reflect"
)

type Config struct {
	APIPort string `yaml:"port" env:"API_PORT"`
}

func (cfg *Config) LoadConfig(cfgFilePath string) {
	// Load configuration from file and from environment
	if err := cfg.LoadConfigFile(cfgFilePath); err != nil {
		fmt.Println("WARN: Could not load config file")
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
