package storage

import (
	"context"
	"errors"
	"time"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

type Storage struct {
	Client     *mongo.Client
	User       *mongo.Collection
	Authorized *mongo.Collection
	Rooms      *mongo.Collection
}

func NewRepository(cfg *viper.Viper, logger *zap.SugaredLogger) (*Storage, error) {

	client, err := mongo.NewClient(options.Client().ApplyURI(cfg.GetString("db.url")))
	if err != nil {
		logger.Info("NewClient initialization error")
		return nil, errors.New(err.Error())
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	// Create connect
	err = client.Connect(ctx)
	if err != nil {
		logger.Info("Connection to db error")
		return nil, errors.New(err.Error())
	}

	// Check the connection
	err = client.Ping(ctx, nil)
	if err != nil {
		logger.Info("Ping signal delivery error")
		return nil, errors.New(err.Error())
	}

	res := &Storage{
		User:       client.Database("sulifa").Collection("users"),
		Authorized: client.Database("sulifa").Collection("authorized"),
		Rooms:      client.Database("sulifa").Collection("rooms"),
	}

	return res, nil
}
