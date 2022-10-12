package main

import (
	"log"

	"github.com/ZeineI/sulifa/config"
	"github.com/ZeineI/sulifa/internal/app"
)

func main() {
	// Configuration
	cfg, err := config.ParseConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	// Run
	app.Run(cfg)
}
