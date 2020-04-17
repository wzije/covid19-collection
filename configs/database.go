package configs

import (
	"context"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ctx = context.Background()

func DBConnect() (*mongo.Database, error) {

	clientOptions := options.Client()

	url := "mongodb://" +
		viper.GetString("database.host") +
		":" + viper.GetString("database.port")

	user := viper.GetString("database.user")
	pass := viper.GetString("database.password")

	clientOptions.ApplyURI(url).
		SetAuth(options.Credential{Username: user, Password: pass})

	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Connect(ctx)

	if err != nil {
		return nil, err
	}

	return client.Database(viper.GetString("database.name")), nil
}
