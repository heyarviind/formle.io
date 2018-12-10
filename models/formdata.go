package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

//FormSubmisson struct
type FormSubmisson struct {
	ID       bson.ObjectId `json:"id" bson:"_id"`
	UserID   bson.ObjectId `json:"userId" bson:"userId"`
	FormData string        `json:"formData" bson:"formData"`
	FormURL  string        `json:"formURL" bson:"formURL"`
	Created  time.Time     `json:"created" bson:"created"`
}
