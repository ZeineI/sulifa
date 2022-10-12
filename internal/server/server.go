package server

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type Server struct {
	router *gin.Engine
}

func NewServer() *Server {
	return &Server{
		router: gin.Default(),
	}
}

func (server *Server) Run(cfg *viper.Viper, logger *zap.SugaredLogger) error {
	server.router.POST("/register", server.register)
	return server.router.Run()
}

func (sv *Server) register(c *gin.Context) {
}
