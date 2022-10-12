package server

import (
	"net/http"

	"github.com/ZeineI/sulifa/internal/models"
	"github.com/ZeineI/sulifa/utils/encode"
	"github.com/ZeineI/sulifa/utils/validate"
	"github.com/gin-gonic/gin"
)

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

	if err := sv.storage.InsertUser(user); err != nil {
		sv.logger.Info(err)
		c.JSON(http.StatusBadRequest, err.Error())
	}

	c.JSON(http.StatusOK, nil)
}

func (sv *Server) login(c *gin.Context) {

	var req *models.LoginReq

	if err := c.BindJSON(&req); err != nil {
		sv.logger.Info(err)
		c.JSON(http.StatusBadRequest, err.Error())
	}

	user, err := sv.storage.GetUser(req.Username)

	if err != nil {
		sv.logger.Info(err)
		c.JSON(http.StatusBadRequest, err.Error())
	}

	c.JSON(http.StatusOK, user)
}
