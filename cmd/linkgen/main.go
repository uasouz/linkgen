package main

import (
	"linkgen/config"
	"linkgen/linkgen"
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
	linkgenService := linkgen.New(cfg.APIPort, linkStore)

	// Start API Service
	linkgenService.Start()
}
