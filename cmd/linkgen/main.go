package main

import (
	"linkgen/config"
	"linkgen/linkgen"
	"linkgen/store/mysql"
)

func main() {
	// Start configuration structure with default values
	cfg := config.Config{
		APIPort: "3000",
	}

	// Call configuration load function
	cfg.LoadConfig("config.yaml")

	// Get new instance of store.LinkStore
	linkStore, err := mysql.New(cfg.DBDSN)

	// Panic if there is no storage
	if err != nil {
		panic("no database connection")
	}

	// Creates a new API Service with the port configuration
	linkgenService := linkgen.New(cfg.APIPort, linkStore)

	// Start API Service
	linkgenService.Start()
}
