package main

import (
	"encoding/json"
	"fmt"
	"go-mailapp/config"
	"go-mailapp/controller"
	"go-mailapp/db"
	"go-mailapp/slack"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/keighl/postmark"
	"github.com/unrolled/render"
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
		h := render.New(render.Options{
			Extensions: []string{".html"},
			Layout:     "layout",
		})

		var (
			toReplyEmail string
			toEmail      string
		)

		var params = make(map[string]string)
		var fromURL = r.Host + r.URL.Path

		toEmail = mux.Vars(r)["email"]

		if err := r.ParseForm(); err != nil {
			fmt.Println("unable to parse the form")
			h.HTML(w, http.StatusOK, "wentwrong", nil)
		}

		for key, value := range r.PostForm {
			if key == "_replyto" {
				toReplyEmail = strings.Join(value, "")
			}
			params[strings.Title(key)] = strings.Join(value, "")
		}

		var jsonParams, _ = json.Marshal(params)

		// Validate emails
		if checkEmail(toEmail) {
			if len(toReplyEmail) > 0 && checkEmail(toReplyEmail) {
				user, verified := controller.CheckUser(toEmail)

				if len(user.ID) > 0 {
					if verified {
						//Check the limit
						if controller.CheckUserLimit(toEmail) {
							//Insert data into database
							controller.InsertFormIntoDatbase(toEmail, fromURL, string(jsonParams))
							//TODO Send the email
							h.HTML(w, http.StatusOK, "emailsent", nil)
						} else {
							//TODO Send limit crossed email
							fmt.Println("Sending limit cross email!")
							slack.MakeNotification("[LIMIT] " + toEmail + " reached the monthly limit! 😅")
							controller.InsertFormIntoDatbase(toEmail, fromURL, string(jsonParams))
						}
					} else {
						//Send an email to verify the email
						h.HTML(w, http.StatusOK, "verify-email", nil)
					}
				} else {
					createUser := controller.CreateUser(toEmail)
					if createUser {
						h.HTML(w, http.StatusOK, "verify-email", nil)
					}
				}
			} else {
				//TODO Error, user need to enter valid email
			}

		} else {
			//TODO Invalid email provided
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
