package server

import (
	"github.com/ZeineI/sulifa/internal/storage"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type Server struct {
	router  *gin.Engine
	storage *storage.Storage
	cfg     *viper.Viper
	logger  *zap.SugaredLogger
}

func NewServer(db *storage.Storage, logger *zap.SugaredLogger, cfg *viper.Viper) *Server {
	return &Server{
		router:  gin.Default(),
		storage: db,
		logger:  logger,
	}
}

func (server *Server) Run(cfg *viper.Viper) error {

	//endpoints
	server.router.POST("/register", server.register)
	server.router.POST("/login", server.login)

	//websocket endpoints
	server.router.POST("/createroom", server.createRoom)
	server.router.POST("/joinroom", server.joinRoom)

	//listen and serve
	return server.router.Run(cfg.GetString("server.port"))
}
