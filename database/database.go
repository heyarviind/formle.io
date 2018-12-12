package database

import (
	"go-mailapp/config"
	"go-mailapp/models"

	"github.com/labstack/gommon/log"
	"gopkg.in/mgo.v2"
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

	// Lets go to database
	db := session.DB(config.DBName)
	userCol := db.C("users")
	if err := userCol.EnsureIndex(models.UserModelIndex()); err != nil {
		panic("duplicate entry of email")
	}

	if userCol == nil {
		panic("User col is nil")
	}

	// if err := userCol.Insert(&user); err != nil {
	// 	panic("unable to insert data")
	// }

	if db == nil {
		log.Errorf("db %v not found, exiting...", config.DBName)
		return
	}
}
