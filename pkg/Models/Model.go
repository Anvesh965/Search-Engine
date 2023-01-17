package Models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Webpage struct {
	Id       primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Title    string             `json:"title"`
	Keywords []string           `json:"keywords"`
}

