package common

import (
	"encoding/json"

	"github.com/dapr/go-sdk/service/common"
)

type TaskMessage struct {
	ID    string `json:"job"`
	Value int    `json:"value"`
}

func (t *TaskMessage) ToData() []byte {
	b, err := json.Marshal(t)
	if err != nil {
		panic(err)
	}
	return b
}

func FromEvent(e *common.TopicEvent) (t TaskMessage, err error) {
	err = json.Unmarshal([]byte(e.Data.(string)), &t)
	if err != nil {
		return t, err
	}
	return t, nil
}
