package main

import (
	"fmt"
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
	r := mux.NewRouter()

	r.HandleFunc("/", indexRoute)
	r.HandleFunc("/{email}", doEmail).Methods("POST")

	if err := http.ListenAndServe(":1420", r); err != nil {
		panic(err)
	}
}

func indexRoute(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello World!\n"))
}

func doEmail(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {

		var (
			toReplyEmail string
			toEmail      string
		)

		var params = make(map[string]string)

		toEmail = mux.Vars(r)["email"]

		if err := r.ParseForm(); err != nil {
			fmt.Println("got stuck into an error")
			panic(err)
		}

		for key, value := range r.PostForm {
			if key == "_replyto" {
				toReplyEmail = strings.Join(value, "")
			}
			params[strings.Title(key)] = strings.Join(value, "")
		}

		// Get the HTML Body
		body := prepareEmailBody(params)

		// Validate emails
		if checkEmail(toEmail) {
			if len(toReplyEmail) > 0 && checkEmail(toReplyEmail) {
				// Send the mail
				// fmt.Println(body)
				if sendEmail(toReplyEmail, toEmail, body) {
					w.WriteHeader(http.StatusOK)
				} else {
					panic("Something went wrong")
				}
			} else {
				if sendEmail("", toEmail, body) {
					w.WriteHeader(http.StatusOK)
				} else {
					panic("Something went wrong")
				}
			}
		} else {
			panic("In valid email provided!")
		}
	}
}

func sendEmail(toReplyEmail, toEmail, body string) bool {

	client := postmark.NewClient(serverToken, accountToken)
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
