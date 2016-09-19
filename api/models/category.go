package models

import "gopkg.in/mgo.v2/bson"

type (
	Category struct {
		ID      bson.ObjectId `json:"id" bson:"_id"`
		Name    string
		Summary string
		Records []*Record
	}
)
