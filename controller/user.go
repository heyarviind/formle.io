package controller

import (
	"go-mailapp/config"
	"go-mailapp/db"
	"go-mailapp/models"
	"time"

	"gopkg.in/mgo.v2/bson"
)

// CheckUser if it exists and is verified
func CheckUser(email string) (models.User, bool) {

	var (
		user         models.User
		userverified bool
	)

	// var user *models.User
	if err := db.MgoSession.DB(config.DBName).C("users").Find(bson.M{"email": email}).One(&user); err != nil {
		return user, userverified
	}

	if user.Verified {
		userverified = true
	}

	return user, userverified
}

//CreateUser if does not exist
func CreateUser(email string) bool {
	var user models.User

	user.ID = bson.NewObjectId()
	user.Email = email
	user.Plan = "free"
	user.RegisteredOn = time.Now()

	if err := db.MgoSession.DB(config.DBName).C("users").Insert(&user); err != nil {
		panic(err)
	}

	return true
}

//CheckUserLimit if it crossed the monthly limit
func CheckUserLimit(email string) bool {
	// var user models.User
	currentTime := time.Now()
	dateMonthYear := currentTime.Format("02-01-2006")

	user, _ := CheckUser(email)

	if err := db.MgoSession.DB(config.DBName).C("users").Find(bson.M{"email": email}).One(&user); err != nil {
		return userExists, userverified
	}

	return false
}

// VerifyUser when they click on verify email on first time form submission
func VerifyUser(email string) bool {
	if userExists, _ := CheckUser(email); userExists == true {
		query := bson.M{"email": email}
		change := bson.M{"$set": bson.M{"isverified": true}}
		if err := db.MgoSession.DB(config.DBName).C("users").Update(query, change); err != nil {
			panic(err)
		}
	}

	return true
}
