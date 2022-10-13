package storage

import (
	"context"
	"fmt"

	"github.com/ZeineI/sulifa/internal/models"
	"go.mongodb.org/mongo-driver/bson"
)

func (db *Storage) InsertRegisteredUser(user *models.User) error {
	insertResult, err := db.User.InsertOne(context.Background(), user)
	if err != nil {
		return fmt.Errorf("Insert user: %s", err)
	}
	fmt.Println(insertResult.InsertedID)
	return nil
}

func (db *Storage) InsertLogedInUser(user *models.User) error {
	insertResult, err := db.Authorized.InsertOne(context.Background(), user)
	if err != nil {
		return fmt.Errorf("Insert user: %s", err)
	}
	fmt.Println(insertResult.InsertedID)
	return nil
}

func (db *Storage) GetUserByName(username string) (*models.User, error) {
	var user *models.User
	filter := bson.D{{"username", username}}
	if err := db.User.FindOne(context.Background(), filter).Decode(&user); err != nil {
		return nil, fmt.Errorf("Get user by name: %s", err)
	}
	return user, nil
}

func (db *Storage) GetUserAuth(username string) (*models.User, error) {
	var user *models.User
	filter := bson.D{{"username", username}}
	if err := db.Authorized.FindOne(context.Background(), filter).Decode(&user); err != nil {
		return nil, fmt.Errorf("Get user by name: %s", err)
	}
	return user, nil
}

func (db *Storage) GetId(username string) (*models.HexId, error) {
	var result *models.HexId
	filter := bson.D{{"username", username}}
	if err := db.Authorized.FindOne(context.Background(), filter).Decode(&result); err != nil {
		return nil, fmt.Errorf("Get user by name: %s", err)
	}
	return result, nil
}
