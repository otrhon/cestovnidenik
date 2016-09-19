package models

import "gopkg.in/mgo.v2/bson"

type (
	Record struct {
		ID      bson.ObjectId `json:"id" bson:"_id"`
		Heading string
		Content string
	}
)
