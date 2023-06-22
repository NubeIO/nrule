package apirules

import (
	"encoding/json"
	"github.com/jordan-wright/email"
	"github.com/labstack/gommon/log"
	"net/smtp"
)

/*
let body = {
  to: ["a@nube-io.com"],
  subject: "test",
  message: "testing",
  senderAddress: "aa@nube-io.com",
  password: "abc",
};

RQL.SendEmail(body);
*/

type mail struct {
	To            []string
	Cc            []string
	Bcc           []string
	Subject       string
	Message       string
	SenderAddress string
	Password      string
}

func emailBody(body any) (*mail, error) {
	result := &mail{}
	dbByte, err := json.Marshal(body)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(dbByte, &result)
	return result, err
}

func (inst *Client) SendEmail(body any) {
	parsed, err := emailBody(body)
	to := parsed.To
	subject := parsed.Subject
	message := parsed.Message
	senderAddress := parsed.SenderAddress
	password := parsed.Password

	e := email.NewEmail()
	e.From = senderAddress
	e.To = to
	e.Cc = parsed.To
	e.Bcc = parsed.Bcc
	e.Subject = subject
	e.HTML = []byte(message)
	err = e.Send("smtp.gmail.com:587", smtp.PlainAuth("", senderAddress, password, "smtp.gmail.com"))
	if err != nil {
		log.Error(err)
		return
	} else {

	}
}
