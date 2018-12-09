package controller

import (
	"go-mailapp/config"
	"go-mailapp/db"
	"go-mailapp/models"

	"gopkg.in/mgo.v2/bson"
)

// CheckUser if it exists and is verified
func CheckUser(email string) (bool, bool) {

	var (
		userExists   bool
		userverified bool
	)

	var user *models.User
	if err := db.MgoSession.DB(config.DBName).C("users").Find(bson.M{"email": email}).One(&user); err != nil {
		return userExists, userverified
	}

	if user != nil {
		userExists = true

		if user.Verified {
			userverified = true
		}
	}

	return userExists, userverified
}

//CreateUser if does not exist
func CreateUser(email string) bool {
	var user models.User

	user.ID = bson.NewObjectId()
	user.Email = email

	if err := db.MgoSession.DB(config.DBName).C("users").Insert(&user); err != nil {
		panic(err)
	}

	return true
}

func VerifyUser(email string) bool {

}
