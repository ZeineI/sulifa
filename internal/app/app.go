package app

import (
	"context"
	"log"

	"github.com/ZeineI/sulifa/internal/server"
	"github.com/ZeineI/sulifa/internal/storage"
	logger "github.com/ZeineI/sulifa/utils/log"
	"github.com/spf13/viper"
)

func Run(cfg *viper.Viper) {
	logger, err := logger.NewLogger()
	if err != nil {
		log.Fatalf("Logger initialization error: %v", err)
	}

	storage, err := storage.NewRepository(cfg, logger)
	if err != nil {
		log.Fatalf("MongoDB initialization error: %v", err)
	}

	defer func() {
		err = storage.Client.Disconnect(context.Background())
		if err != nil {
			logger.Info("Disconnect db")
			return
		}

		logger.Info("Connection to MongoDB closed")
	}()

	router := server.NewServer(storage, logger, cfg)

	if err := router.Run(cfg); err != nil {
		logger.Info(err)
		return
	}
}
