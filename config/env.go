package config

import (
	"os"
	"reflect"
)

// LoadConfigEnv - Loads Config values from Environment Variables
func (cfg *Config) LoadConfigEnv() {
	av := reflect.ValueOf(cfg)
	v := av.Elem()

	for i := 0; i < v.NumField(); i++ {
		envTag := v.Type().Field(i).Tag.Get("env")
		field := v.Field(i)
		if envTag != "" {
			envValue := os.Getenv(envTag)
			if envValue != "" && field.CanSet() {
				if field.Kind() == reflect.String {
					field.SetString(envValue)
				}
				if field.Kind() == reflect.Bool {
					field.SetBool(envValue == "true")
				}
			}
		}
	}
}
