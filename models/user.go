package models

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// User struct
type User struct {
	ID       bson.ObjectId `json:"id" bson:"_id"`
	Email    string        `json:"email" bson:"email"`
	Verified bool          `json:"isverified" bson:"isverified"`
}

// UserModelIndex ...
func UserModelIndex() mgo.Index {
	return mgo.Index{
		Key:    []string{"email"},
		Unique: true,
	}
}
