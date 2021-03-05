package state

import (
	"context"

	dapr "github.com/dapr/go-sdk/client"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const MONGO_URI = "MONGO_URI"

func init() {
	viper.SetDefault(MONGO_URI, "mongodb://localhost:27017")
	viper.AutomaticEnv()
}

type Handler struct {
	Client dapr.Client
	DB     mongo.Collection
}

func New(ctx context.Context, c dapr.Client) *Handler {
	client, err := mongo.Connect(ctx, options.Client().
		ApplyURI(
			viper.GetString(MONGO_URI),
		).
		SetAuth(
			options.Credential{Username: viper.GetString("MONGO_USERNAME"), Password: viper.GetString("MONGO_PASSWORD")},
		),
	)
	if err != nil {
		panic(err)
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		panic(err)
	}
	collection := client.Database("db").Collection("state")
	_, err = collection.Indexes().DropAll(ctx)
	if err != nil {
		panic(err)
	}
	_, err = collection.Indexes().CreateMany(ctx, []mongo.IndexModel{{
		Keys:    bson.D{{"timestampid", -1}},
		Options: options.Index().SetName("timestampDESC"),
	}, {
		Keys:    bson.D{{"timestampid", 1}},
		Options: options.Index().SetName("timestampASC"),
	}})
	if err != nil {
		panic(err)
	}
	go func() {
		<-ctx.Done()
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
	return &Handler{Client: c, DB: *collection}
}
