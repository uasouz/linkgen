package main

import (
	"linkgen/config"
)

func main() {
	cfg := config.Config{APIPort: "3000"}

	cfg.LoadConfig("config.yaml")

}
