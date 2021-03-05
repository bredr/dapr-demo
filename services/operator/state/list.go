package state

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/dapr/go-sdk/service/common"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (h *Handler) List(ctx context.Context, in *common.InvocationEvent) (out *common.Content, err error) {
	if in == nil {
		err = errors.New("invocation param required")
		return
	}
	var input struct {
		After *string `json:"after"`
		First int     `json:"first"`
	}
	err = json.Unmarshal(in.Data, &input)
	if err != nil {
		return
	}
	s := options.Find()
	s.SetSort(bson.D{{"timestampid", -1}})
	s.SetLimit(int64(input.First))
	if input.First == 0 {
		s.SetLimit(100)
	}
	var results []Run
	var cursor *mongo.Cursor
	if input.After != nil {
		cursor, err = h.DB.Find(ctx, bson.D{{"timestampid", bson.D{{"$lt", *input.After}}}}, s)
	} else {
		cursor, err = h.DB.Find(ctx, bson.D{}, s)
	}
	if err != nil {
		return
	}
	if err = cursor.All(ctx, &results); err != nil {
		return
	}
	var b []byte
	b, err = json.Marshal(&results)
	if err != nil {
		return
	}
	return &common.Content{
		Data:        b,
		ContentType: in.ContentType,
		DataTypeURL: in.DataTypeURL,
	}, nil
}
