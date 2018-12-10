package main

import (
	"fmt"
	"go-mailapp/config"
	"go-mailapp/controller"
	"go-mailapp/db"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/keighl/postmark"
)

//EmailResponse struct
type EmailResponse struct {
	Message string `json:"message" xml:"message"`
}

func main() {
	// Connect to database
	db.Init()
	// Close the mongodb session
	// defer db.MgoSession.Close()

	r := mux.NewRouter()

	r.HandleFunc("/", indexRoute)
	r.HandleFunc("/{email}", doEmail).Methods("POST")
	r.HandleFunc("/verify/{email}/{uuid}", verifyEmail).Methods("GET")

	if err := http.ListenAndServe(":1420", r); err != nil {
		panic(err)
	}
}

func indexRoute(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello World!\n"))
	return
}

func doEmail(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {

		var (
			toReplyEmail string
			toEmail      string
		)

		var params = make(map[string]string)
		var fromURL = string(r.Host) + r.URL.Path

		toEmail = mux.Vars(r)["email"]

		if err := r.ParseForm(); err != nil {
			fmt.Println("unable to parse the form")
			panic(err)
		}

		for key, value := range r.PostForm {
			if key == "_replyto" {
				toReplyEmail = strings.Join(value, "")
			}
			params[strings.Title(key)] = strings.Join(value, "")
		}

		// Validate emails
		if checkEmail(toEmail) {
			fmt.Println("checked email")
			if len(toReplyEmail) > 0 && checkEmail(toReplyEmail) {
				fmt.Println("checkd reply email")
				user, verified := controller.CheckUser(toEmail)

				if len(user.ID) > 0 {
					fmt.Println("user found")
					if verified {
						//Check the limit
						if controller.CheckUserLimit(toEmail) {
							//Send the email
							controller.InsertFormIntoDatbase(toEmail, fromURL)
						}
					} else {
						//Send an email to verify the email
					}
				} else {
					fmt.Println("email not found, creating the email")
					createUser := controller.CreateUser(toEmail)
					if createUser {
						//Redirect user to verify the email
					}
				}
				// Get the HTML Body
				// body := prepareEmailBody(params)

				// if sendEmail(toReplyEmail, toEmail, body) {
				// 	w.WriteHeader(http.StatusOK)
				// } else {
				// 	panic("Something went wrong")
				// }
			} else {
				//Error, user need to enter valid email
			}

		} else {
			panic("In valid email provided!")
		}
	}
}

func verifyEmail(w http.ResponseWriter, r *http.Request) {

}

func sendEmail(toReplyEmail, toEmail, body string) bool {
	client := postmark.NewClient(config.ServerToken, config.AccountToken)
	email := postmark.Email{
		From:       "no-reply@uicard.io",
		To:         toEmail,
		ReplyTo:    toReplyEmail,
		Subject:    "New form submission - UICardio",
		HtmlBody:   body,
		TextBody:   body,
		Tag:        "pw-reset",
		TrackOpens: true,
	}

	_, err := client.SendEmail(email)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}
