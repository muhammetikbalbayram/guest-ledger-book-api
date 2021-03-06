package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	Ctx = context.TODO()
)

func Db() *mongo.Client {
	client, err := mongo.Connect(Ctx, options.Client().ApplyURI(""))
	if err != nil {
		panic(err)
	}

	if err := client.Ping(Ctx, readpref.Primary()); err != nil {
		panic(err)
	}

	return client
}
