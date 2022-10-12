package app

import (
	"log"

	"github.com/ZeineI/sulifa/internal/server"
	logger "github.com/ZeineI/sulifa/pkg/log"
	"github.com/spf13/viper"
)

func Run(cfg *viper.Viper) {
	logger, err := logger.NewLogger()
	if err != nil {
		log.Fatalf("Logger initialization error: %v", err)
	}

	router := server.NewServer()

	if err := router.Run(cfg, logger); err != nil {
		logger.Info(err)
		return
	}
}
