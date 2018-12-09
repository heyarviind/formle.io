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
			} else {
				params[strings.Title(key)] = strings.Join(value, "")
			}
		}

		// Validate emails
		if checkEmail(toEmail) && checkEmail(toReplyEmail) {
			// Get the HTML Body
			body := prepareEmailBody(params)

			// Send the mail
			// fmt.Println(body)
			sendEmail(toReplyEmail, toEmail, body)
		} else {
			panic("In valid email provided!")
		}

	}

	w.WriteHeader(http.StatusOK)
}

func sendEmail(toReplyEmail, toEmail, body string) bool {

	fmt.Printf("Reply to email : %s\n", toReplyEmail)
	fmt.Printf("to email : %s\n", toEmail)

	client := postmark.NewClient(serverToken, accountToken)
	email := postmark.Email{
		From:       "no-reply@uicard.io",
		To:         toEmail,
		ReplyTo:    toReplyEmail,
		Subject:    "New form submission - UICardio",
		HtmlBody:   body, //[TODO] place html theme here
		TextBody:   body,
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
