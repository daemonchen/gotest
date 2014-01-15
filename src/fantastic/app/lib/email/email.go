package email

import (
	"github.com/scorredoira/email"
	"net/smtp"
)

func Mail(username string, password string, title string, body string) {
	m := email.NewHTMLMessage(title, body)
	m.From = "daemon"
	m.To = []string{"cgyqqcgy@126.com", "cgyqqcgy@gmail.com"}
	// m.Cc = []string{"cc1@example.com", "cc2@example.com"}
	// m.Bcc = []string{"bcc1@example.com", "bcc2@example.com"}

	err := email.Send("smtp.gmail.com:587", smtp.PlainAuth("", username, password, "smtp.gmail.com"), m)
	if err != nil {
		panic(err)
	}
}
