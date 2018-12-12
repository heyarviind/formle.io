package main

import (
	"encoding/json"
	"fmt"
	"formle/controller"
	"formle/database"
	"formle/slack"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

//EmailResponse struct
type EmailResponse struct {
	Message string `json:"message" xml:"message"`
}

func main() {
	// Connect to database
	database.Init()
	// Close the mongodb session
	// defer db.MgoSession.Close()

	r := mux.NewRouter()

	r.HandleFunc("/", indexRoute)
	r.HandleFunc("/{email}", doEmail).Methods("POST")
	r.HandleFunc("/verify/{uuid}", verifyEmail).Methods("GET")

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

			user, verified := controller.CheckUser(toEmail)
			if len(toReplyEmail) > 0 {
				if checkEmail(toReplyEmail) != true {
					fmt.Println("Invalid to email")
					h.HTML(w, http.StatusOK, "invalid-email", nil)
				}
			}
			if len(user.ID) > 0 {
				if verified {
					//Check the limit
					if controller.CheckUserLimit(toEmail) {
						//Insert data into database
						controller.InsertFormIntoDatbase(toEmail, fromURL, string(jsonParams))
						//TODO Send the email
					} else {
						//TODO Send limit crossed email
						fmt.Println("Sending limit cross email!")
						slack.MakeNotification("[LIMIT] " + toEmail + " reached the monthly limit! 😅")
						controller.InsertFormIntoDatbase(toEmail, fromURL, string(jsonParams))
						// TODO inform user about the limit
						controller.SendLimitCrossedEmail(toEmail)
					}
					h.HTML(w, http.StatusOK, "emailsent", nil)
				} else {
					//Send an email to verify the email
					h.HTML(w, http.StatusOK, "verify-email", nil)
				}
			} else {
				if createUser := controller.CreateUser(toEmail); createUser == true {
					// TODO send verification email
					controller.SendVerificationEmail(toEmail)
					h.HTML(w, http.StatusOK, "verify-email", nil)
				}
			}

		} else {
			//TODO Invalid email provided
		}
	}
}

func verifyEmail(w http.ResponseWriter, r *http.Request) {
	h := render.New(render.Options{
		Extensions: []string{".html"},
		Layout:     "layout",
	})

	var (
		uuid string
	)

	uuid = mux.Vars(r)["uuid"]

	if controller.VerifyUser(uuid) == true {
		h.HTML(w, http.StatusOK, "email-verified", nil)
	} else {
		h.HTML(w, http.StatusOK, "error", nil)
	}

}
