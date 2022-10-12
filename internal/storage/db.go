package storage

import (
	"context"
	"errors"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

func NewRepository(cfg *viper.Viper, logger *zap.SugaredLogger) (*mongo.Client, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(cfg.GetString("database.url")))
	if err != nil {
		logger.Info("NewClient initialization error")
		return nil, errors.New(err.Error())
	}

	// Create connect
	err = client.Connect(context.TODO())
	if err != nil {
		logger.Info("Connection to db error")
		return nil, errors.New(err.Error())
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		logger.Info("Ping signal delivery error")
		return nil, errors.New(err.Error())
	}

	err = client.Disconnect(context.TODO())

	if err != nil {
		logger.Info("Disconnect db")
		return nil, errors.New(err.Error())
	}

	logger.Info("Connection to MongoDB closed")

	return client, nil
}
