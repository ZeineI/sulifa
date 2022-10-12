package server

import (
	"net/http"

	"github.com/ZeineI/sulifa/internal/models"
	"github.com/ZeineI/sulifa/pkg/encode"
	"github.com/ZeineI/sulifa/pkg/validate"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

type Server struct {
	router  *gin.Engine
	storage *mongo.Client
	logger  *zap.SugaredLogger
}

func NewServer(db *mongo.Client, logger *zap.SugaredLogger) *Server {
	return &Server{
		router:  gin.Default(),
		storage: db,
		logger:  logger,
	}
}

func (server *Server) Run(cfg *viper.Viper) error {
	server.router.POST("/register", server.register)
	return server.router.Run(cfg.GetString("server.port"))
}

func (sv *Server) register(c *gin.Context) {

	var req *models.RegisterReq

	if err := c.BindJSON(&req); err != nil {
		sv.logger.Info(err)
		c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := validate.ValidateCreds(req.Username, req.Password); err != nil {
		sv.logger.Info(err)
		c.JSON(http.StatusBadRequest, err.Error())
	}

	user := &models.User{
		Username: req.Username,
		Password: encode.GenerateHash(req.Password),
	}

	if err := sv.InsertUser(user); err != nil {
		sv.logger.Info(err)
		c.JSON(http.StatusBadRequest, err.Error())
	}

	c.JSON(http.StatusOK, nil)
}
