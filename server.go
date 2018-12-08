package main

import (
	"net/http"

	"github.com/keighl/postmark"
	"github.com/labstack/echo"
)

var serverToken = "0b6a11e4-d012-415d-9ce7-97010d3b6803"
var accountToken = "ec02a12e-7ae0-4df0-86c6-62bb06ba77f0"

//EmailResponse struct
type EmailResponse struct {
	Message string `json:"message" xml:"message"`
}

func main() {
	// Creating a new echo server
	e := echo.New()
	// Defining all the routes here
	e.GET("/", indexRoute)
	e.POST("/:email", doEmail)
	//Starting the server
	e.Logger.Fatal(e.Start(":1420"))
}

func indexRoute(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, world!")
}

func doEmail(c echo.Context) error {

	res := &EmailResponse{
		Message: c.Param("email"),
	}

	// fmt.Printf("%s, %s, %s\n", name, email, message)

	return c.JSON(http.StatusOK, res)
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
