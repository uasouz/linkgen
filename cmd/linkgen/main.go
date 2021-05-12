package main

import (
	"linkgen/api"
	"linkgen/config"
	"linkgen/store/memory"
)

func main() {
	// Start configuration structure with default values
	cfg := config.Config{APIPort: "3000"}

	// Call configuration load function
	cfg.LoadConfig("config.yaml")

	// Get new instance of store.LinkStore
	linkStore := memory.New()

	// Creates a new API Service with the port configuration
	apiService := api.New(cfg.APIPort, linkStore)

	// Start API Service
	apiService.Start()
}
