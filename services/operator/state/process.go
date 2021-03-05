package state

import (
	"context"
	"fmt"
	"log"
	"time"

	cmn "github.com/bredr/dapr-demo/services/common"
	"github.com/dapr/go-sdk/service/common"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (h *Handler) Process(ctx context.Context, e *common.TopicEvent) (retry bool, err error) {
	log.Printf("event - PubsubName: %s, Topic: %s, ID: %s, Data: %s", e.PubsubName, e.Topic, e.ID, e.Data)
	msg, err := cmn.FromEvent(e)
	if err != nil {
		return false, err
	}

	now := time.Now().Unix()
	state := Run{
		RunID:     msg.ID,
		Stage:     e.Topic,
		Value:     msg.Value,
		Timestamp: primitive.Timestamp{T: uint32(now)},
	}
	result, err := h.DB.InsertOne(ctx, state)
	if err != nil {
		return true, err
	}
	id := result.InsertedID.(primitive.ObjectID)
	_, err = h.DB.UpdateOne(ctx, bson.M{"_id": id},
		bson.D{{"$set", bson.D{{"timestampid", fmt.Sprintf("%d%s", now, id.Hex())}}}},
	)
	if err != nil {
		return true, err
	}
	return false, nil
}
