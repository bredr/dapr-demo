package task

import (
	"context"
	"log"

	cmn "github.com/bredr/dapr-demo/services/common"
	dapr "github.com/dapr/go-sdk/client"
	"github.com/dapr/go-sdk/service/common"
)

type Handler struct {
	Client          dapr.Client
	OutputTopicName string
}

func (h *Handler) Process(ctx context.Context, e *common.TopicEvent) (retry bool, err error) {
	log.Printf("event - PubsubName: %s, Topic: %s, ID: %s, Data: %s", e.PubsubName, e.Topic, e.ID, e.Data)
	msg, err := cmn.FromEvent(e)
	if err != nil {
		return false, nil
	}
	msg.Value += 5
	if err := h.Client.PublishEvent(ctx, e.PubsubName, h.OutputTopicName, msg.ToData()); err != nil {
		return true, err
	}
	return false, nil
}
