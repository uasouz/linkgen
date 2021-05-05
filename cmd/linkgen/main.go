package main

import (
	"linkgen/api"
	"linkgen/config"
)

func main() {
	// Start configuration structure with default values
	cfg := config.Config{APIPort: "3000"}

	// Call configuration load function
	cfg.LoadConfig("config.yaml")

	// Creates a new API Service with the port configuration
	apiService := api.New(cfg.APIPort)

	// Start API Service
	apiService.Start()
}
