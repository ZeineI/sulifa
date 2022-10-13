package storage

import (
	"context"
	"fmt"

	"github.com/ZeineI/sulifa/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (db *Storage) InsertRoom(room *models.Room) (interface{}, error) {
	insertResult, err := db.Rooms.InsertOne(context.Background(), room)
	if err != nil {
		return 0, fmt.Errorf("Insert room: %s", err)
	}
	return insertResult.InsertedID, nil
}

func (db *Storage) GetRoomByID(id primitive.ObjectID) (*models.Room, error) {
	var room *models.Room
	if err := db.Rooms.FindOne(context.Background(), bson.M{"_id": id}).Decode(&room); err != nil {
		return nil, fmt.Errorf("Get room: %s", err)
	}
	return room, nil
}

func (db *Storage) UpdateRoomByID(id primitive.ObjectID, userId string) (interface{}, error) {
	update := bson.D{{"$set", bson.D{{"status", 1}, {"player2Id", userId}}}}

	result, err := db.Rooms.UpdateOne(context.Background(), bson.M{"_id": id}, update)

	if err != nil {
		return nil, err
	}

	return result.UpsertedID, nil
}
