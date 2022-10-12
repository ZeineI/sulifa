package app

import (
	"fmt"
	"log"

	logger "github.com/ZeineI/pkg/log"
	"github.com/spf13/viper"
)

func Run(cfg *viper.Viper) {
	logger, err := logger.NewLogger()
	if err != nil {
		log.Fatalf("Logger initialization error: %v", err)
	}

	fmt.Println(logger)

	// router := server.NewServer()

	// if err := router.Run(cfg, logger, resp); err != nil {
	// 	logger.Info(err)
	// 	return
	// }
}
