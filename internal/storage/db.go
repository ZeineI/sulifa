package storage

import (
	"context"
	"errors"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

type Storage struct {
	Client     *mongo.Client
	User       *mongo.Collection
	Authorized *mongo.Collection
}

func NewRepository(cfg *viper.Viper, logger *zap.SugaredLogger) (*Storage, error) {

	client, err := mongo.NewClient(options.Client().ApplyURI(cfg.GetString("db.url")))
	if err != nil {
		logger.Info("NewClient initialization error")
		return nil, errors.New(err.Error())
	}

	// Create connect
	err = client.Connect(context.Background())
	if err != nil {
		logger.Info("Connection to db error")
		return nil, errors.New(err.Error())
	}

	// Check the connection
	err = client.Ping(context.Background(), nil)
	if err != nil {
		logger.Info("Ping signal delivery error")
		return nil, errors.New(err.Error())
	}

	res := &Storage{
		User:       client.Database("sulifa").Collection("users"),
		Authorized: client.Database("sulifa").Collection("authorized"),
	}

	return res, nil
}
