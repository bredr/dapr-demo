package main

import (
	"log"
	"net/http"

	"github.com/bredr/dapr-demo/services/task1/task"
	dapr "github.com/dapr/go-sdk/client"
	"github.com/dapr/go-sdk/service/common"
	daprd "github.com/dapr/go-sdk/service/http"
)

var sub = &common.Subscription{
	PubsubName: "redis-pubsub",
	Topic:      "task1",
	Route:      "/tasks",
}

func main() {
	s := daprd.NewService(":8080")

	client, err := dapr.NewClient()
	if err != nil {
		log.Fatalf("error creating dapr client: %v", err)
	}

	handler := &task.Handler{Client: client, OutputTopicName: "task2"}
	if err := s.AddTopicEventHandler(sub, handler.Process); err != nil {
		log.Fatalf("error adding topic handler: %v", err)
	}

	if err := s.Start(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("error listening: %v", err)
	}
}
