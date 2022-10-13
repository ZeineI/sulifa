package server

import (
	"net/http"

	"github.com/ZeineI/sulifa/internal/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (sv *Server) createRoom(c *gin.Context) {

	var req *models.CreateRoomReq

	if err := c.BindJSON(&req); err != nil {
		sv.logger.Info(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": err.Error(),
		})
		return
	}

	user, err := sv.isAuth(req.Username)

	if err != nil || user == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": "only auth user can create room",
		})
		return
	}

	res, err := sv.storage.GetId(user.Username)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": "hex id error",
		})
		return
	}

	newRoom := &models.Room{
		Player1Id: res.ID.String(),
		Status:    0,
	}

	roomId, err := sv.storage.InsertRoom(newRoom)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "user info respponse",
		"roomID":  roomId,
	})
}

func (sv *Server) joinRoom(c *gin.Context) {

	var req *models.JoinRoomReq

	if err := c.BindJSON(&req); err != nil {
		sv.logger.Info(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": err.Error(),
		})
		return
	}

	docID, err := primitive.ObjectIDFromHex(req.RoomID)

	if err != nil {
		sv.logger.Info(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": err.Error(),
		})
		return
	}

	user, err := sv.isAuth(req.Username)

	if err != nil || user == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": "only auth user can create room",
		})
		return
	}

	room, err := sv.storage.GetRoomByID(docID)
	if err != nil {
		sv.logger.Info(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": err.Error(),
		})
		return
	}

	if room.Status != 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": "room isnt empty",
		})
		return
	}

	res, err := sv.storage.GetId(user.Username)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": "hex id error",
		})
		return
	}

	if res.ID.String() == room.Player1Id {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": "cant join same client who created the room",
		})
		return
	}

	roomInfo, err := sv.storage.UpdateRoomByID(docID, res.ID.String())
	if err != nil {
		sv.logger.Info(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": roomInfo,
	})
}
