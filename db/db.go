package db

import (
	"go-mailapp/config"
	"go-mailapp/models"

	"github.com/labstack/gommon/log"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// MgoSession int
var MgoSession *mgo.Session

//Init function
func Init() {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	log.Infof("Successfully connected to database!")
	MgoSession = session

	email := "desiboyarvind@gmail.com"

	user := models.User{
		ID:       bson.NewObjectId(),
		Email:    email,
		Verified: true}

	// var result []User

	// Lets go to database
	db := session.DB(config.DBName)
	userCol := db.C("users")
	if err := userCol.EnsureIndex(models.UserModelIndex()); err != nil {
		panic("duplicate entry of email")
	}

	if userCol == nil {
		panic("User col is nil")
	}

	if err := userCol.Insert(&user); err != nil {
		panic("unable to insert data")
	}

	if db == nil {
		log.Errorf("db %v not found, exiting...", config.DBName)
		return
	}
}
