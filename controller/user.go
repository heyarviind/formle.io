package controller

import (
	"fmt"
	"formle/config"
	"formle/database"
	"formle/models"
	"time"

	uuid "github.com/satori/go.uuid"
	"gopkg.in/mgo.v2/bson"
)

// CheckUser if it exists and is verified
func CheckUser(email string) (*models.User, bool) {

	var (
		user         models.User
		userverified bool
	)

	// var user *models.User
	if err := database.MgoSession.DB(config.DBName).C("users").Find(bson.M{"email": email}).One(&user); err != nil {
		return &user, userverified
	}

	if user.Verified {
		userverified = true
	}

	return &user, userverified
}

//CreateUser if does not exist
func CreateUser(email string) bool {
	var user models.User

	user.ID = bson.NewObjectId()
	user.Email = email
	user.Plan = "free"
	user.RegisteredOn = time.Now()
	user.UUID = uuid.Must(uuid.NewV4()).String()

	if err := database.MgoSession.DB(config.DBName).C("users").Insert(&user); err != nil {
		panic(err)
	}

	return true
}

//CheckUserLimit if it crossed the monthly limit
func CheckUserLimit(email string) bool {
	var (
		result  []bson.M
		isValid bool
	)
	// var user models.User
	currentTime := time.Now()
	currentMonth := currentTime.Month()

	user, _ := CheckUser(email)

	if user.Plan == "free" {
		pipe := []bson.M{{"$project": bson.M{"month": bson.M{"$month": "$created"}, "userId": 1}}, {"$match": bson.M{"userId": user.ID, "month": currentMonth}}, {"$group": bson.M{"_id": "null", "total": bson.M{"$sum": 1}}}}
		database.MgoSession.DB(config.DBName).C("formdata").Pipe(pipe).All(&result)
		if len(result) == 1 {
			count := result[0]["total"]
			if count.(int) < config.Plans[user.Plan] {
				isValid = true
			} else {
				isValid = false
			}
		} else {
			isValid = true
		}
	} else {
		isValid = true
	}
	return isValid
}

//InsertFormIntoDatbase ...
func InsertFormIntoDatbase(email, url, params string) bool {
	var FormData models.FormSubmisson

	user, _ := CheckUser(email)

	FormData.ID = bson.NewObjectId()
	FormData.UserID = user.ID
	FormData.FormData = params
	FormData.Created = time.Now()
	FormData.FormURL = url

	if err := database.MgoSession.DB(config.DBName).C("formdata").Insert(&FormData); err != nil {
		panic(err)
	}

	return true
}

// SendVerificationEmail ...
func SendVerificationEmail(email string) bool {
	user, _ := CheckUser(email)

	if user.Verified == false {
		emailBody := PrepareVerificationEmailBody(user.Email, user.UUID)
		if SendEmail("", email, "Verify your email", emailBody) == true {
			fmt.Println("verification email send!")
		}
	}

	return true
}

//SendLimitCrossedEmail ...
func SendLimitCrossedEmail(email string) bool {
	user, _ := CheckUser(email)

	emailBody := PrepareLimitEmailBody()

	if user.Plan == "free" {
		if CheckUserLimit(email) == false {
			if SendEmail("", email, "Your email limit have Crossed, please upgrade your plan!", emailBody) == true {
				return true
			}
		}
	}

	return true
}

// VerifyUser when they click on verify email on first time form submission
func VerifyUser(uuid string) bool {

	query := bson.M{"uuid": uuid}
	change := bson.M{"$set": bson.M{"isverified": true}}
	if err := database.MgoSession.DB(config.DBName).C("users").Update(query, change); err != nil {
		return false
	}

	return true
}
