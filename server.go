package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"

	"github.com/keighl/postmark"
)

var serverToken = "0b6a11e4-d012-415d-9ce7-97010d3b6803"
var accountToken = "ec02a12e-7ae0-4df0-86c6-62bb06ba77f0"

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
		fmt.Printf("reached here\n")

		if err := r.ParseForm(); err != nil {
			fmt.Println("got stuck into an error")
			panic(err)
		}

		for key, value := range r.PostForm {
			fmt.Printf("%s = %s\n", key, strings.Join(value, ""))

			if key == "_replyto" {

			}
		}

	}

	w.WriteHeader(http.StatusOK)
}

func sendEmail(toReplyEmail, toEmail, body string) bool {
	client := postmark.NewClient(serverToken, accountToken)
	email := postmark.Email{
		From:       "no-reply@uicard.io",
		To:         toEmail,
		ReplyTo:    toReplyEmail,
		Subject:    "New form submission - UICardio",
		HtmlBody:   "...", //[TODO] place html theme here
		TextBody:   "...",
		Tag:        "pw-reset",
		TrackOpens: true,
	}

	_, err := client.SendEmail(email)
	if err != nil {
		panic(err)
		return false
	}
	return true
}
