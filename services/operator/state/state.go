package state

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"

	cmn "github.com/bredr/dapr-demo/services/common"
	dapr "github.com/dapr/go-sdk/client"
	"github.com/dapr/go-sdk/service/common"
	"github.com/google/uuid"
)

type Handler struct {
	Client dapr.Client
}

func (h *Handler) Process(ctx context.Context, e *common.TopicEvent) (retry bool, err error) {
	log.Printf("event - PubsubName: %s, Topic: %s, ID: %s, Data: %s", e.PubsubName, e.Topic, e.ID, e.Data)
	msg, err := cmn.FromEvent(e)
	if err != nil {
		return false, err
	}

	err = h.Client.SaveState(ctx, "statestore", fmt.Sprintf("%s:%s", msg.ID, e.Topic), []byte(fmt.Sprint(msg.Value)))
	if err != nil {
		return true, err
	}
	return false, nil
}

func (h *Handler) Run(ctx context.Context, in *common.InvocationEvent) (out *common.Content, err error) {
	if in == nil {
		err = errors.New("invocation param required")
		return
	}
	var input struct {
		Value int `json:"value"`
	}
	err = json.Unmarshal(in.Data, &input)
	if err != nil {
		return
	}
	id := uuid.New().String()
	var b []byte
	b, err = json.Marshal(&cmn.TaskMessage{ID: id, Value: input.Value})
	if err != nil {
		return
	}
	err = h.Client.PublishEvent(ctx, "redis-pubsub", "task1", b)
	if err != nil {
		return
	}
	return &common.Content{
		Data:        b,
		ContentType: in.ContentType,
		DataTypeURL: in.DataTypeURL,
	}, nil
}
