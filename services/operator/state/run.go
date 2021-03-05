package state

import (
	"context"
	"encoding/json"
	"errors"

	cmn "github.com/bredr/dapr-demo/services/common"
	"github.com/dapr/go-sdk/service/common"
	"github.com/google/uuid"
)

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
