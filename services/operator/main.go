package main

import (
	"context"
	"log"
	"net/http"

	"github.com/bredr/dapr-demo/services/operator/state"
	dapr "github.com/dapr/go-sdk/client"
	"github.com/dapr/go-sdk/service/common"
	daprd "github.com/dapr/go-sdk/service/http"
)

var sub1 = &common.Subscription{
	PubsubName: "redis-pubsub",
	Topic:      "task1",
	Route:      "/task1",
}
var sub2 = &common.Subscription{
	PubsubName: "redis-pubsub",
	Topic:      "task2",
	Route:      "/task2",
}
var sub3 = &common.Subscription{
	PubsubName: "redis-pubsub",
	Topic:      "task3",
	Route:      "/task3",
}

func main() {
	s := daprd.NewService(":8080")

	client, err := dapr.NewClient()
	if err != nil {
		log.Fatalf("error creating dapr client: %v", err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	handler := state.New(ctx, client)
	if err := s.AddTopicEventHandler(sub1, handler.Process); err != nil {
		log.Fatalf("error adding topic handler: %v", err)
	}
	if err := s.AddTopicEventHandler(sub2, handler.Process); err != nil {
		log.Fatalf("error adding topic handler: %v", err)
	}
	if err := s.AddTopicEventHandler(sub3, handler.Process); err != nil {
		log.Fatalf("error adding topic handler: %v", err)
	}
	if err := s.AddServiceInvocationHandler("/run", handler.Run); err != nil {
		log.Fatalf("error adding invocation handler: %v", err)
	}
	if err := s.AddServiceInvocationHandler("/state", handler.List); err != nil {
		log.Fatalf("error adding invocation handler: %v", err)
	}

	if err := s.Start(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("error listening: %v", err)
	}
}
