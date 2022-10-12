package server

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/ZeineI/sulifa/internal/models"
	"github.com/ZeineI/sulifa/utils/encode"
	"github.com/ZeineI/sulifa/utils/validate"
	"github.com/gin-gonic/gin"
)

func (sv *Server) register(c *gin.Context) {

	req := &models.RegisterReq{}

	if err := c.BindJSON(&req); err != nil {
		sv.logger.Info(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": err.Error(),
		})
		return
	}
	fmt.Println(req)
	if err := validate.ValidateCreds(req.Username, req.Password); err != nil {
		sv.logger.Info(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": err.Error(),
		})
		return
	}

	user := &models.User{
		Username: req.Username,
		Password: encode.GenerateHash(req.Password),
	}

	if err := sv.storage.InsertUser(user); err != nil {
		sv.logger.Info(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "user is created",
	})
}

func (sv *Server) login(c *gin.Context) {

	var req *models.LoginReq

	if err := c.BindJSON(&req); err != nil {
		sv.logger.Info(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": err.Error(),
		})
		return
	}

	user, err := sv.storage.GetUser(req.Username)

	if !encode.ComparePasswordHash(user.Password, req.Password) {
		sv.logger.Info(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": errors.New("Incorrect password"),
		})
		return
	}

	if err != nil {
		sv.logger.Info(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": "Unregister user",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":   "success",
		"message":  "user info respponse",
		"userInfo": user,
	})
}
