package model

import "go.mongodb.org/mongo-driver/v2/bson"

// type Netflix struct {
// 	ID      primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
// 	Movie   string             `json:"movie,omitempty"`
// 	Watched bool               `json:"watched,omitempty"`
// }

type Netflix struct {
	ID      bson.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Movie   string        `json:"movie,omitempty" bson:"movie,omitempty"`
	Watched *bool         `json:"watched,omitempty" bson:"watched,omitempty"`
}
