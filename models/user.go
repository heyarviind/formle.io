package models

import (
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// User struct
type User struct {
	ID           bson.ObjectId `json:"id" bson:"_id"`
	UUID         string        `json:"uuid" bson:"uuid"`
	Email        string        `json:"email" bson:"email"`
	Password     string        `json:"password" bson:"password"`
	Plan         string        `json:"plan" bson:"plan"`
	Verified     bool          `json:"isverified" bson:"isverified"`
	RegisteredOn time.Time     `json:"registeredOn" bson:"registeredOn"`
}

// UserModelIndex ...
func UserModelIndex() mgo.Index {
	return mgo.Index{
		Key:    []string{"email"},
		Unique: true,
	}
}
