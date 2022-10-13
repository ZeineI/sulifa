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
	insertResult, err := db.User.InsertOne(context.Background(), user)
	if err != nil {
		return fmt.Errorf("Insert user: %s", err)
	}
	fmt.Println(insertResult.InsertedID)
	return nil
}

func (db *Storage) GetUser(userInfo string) (*models.User, error) {
	var user *models.User
	filter := bson.D{{"username", userInfo}}
	if err := db.User.FindOne(context.Background(), filter).Decode(&user); err != nil {
		return nil, fmt.Errorf("Get user: %s", err)
	}
	return user, nil
}
