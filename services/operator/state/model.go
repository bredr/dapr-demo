package state

import "go.mongodb.org/mongo-driver/bson/primitive"

type Run struct {
	ID          primitive.ObjectID  `bson:"_id,omitempty"`
	TimestampID string              `bson:"timestampid,omitempty"`
	RunID       string              `bson:"runid,omitempty"`
	Value       int                 `bson:"value,omitempty"`
	Stage       string              `bson:"stage,omitempty"`
	Timestamp   primitive.Timestamp `bson:"timestamp,omitempty"`
}
