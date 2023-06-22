package apirules

import (
	"encoding/json"
	"github.com/jordan-wright/email"
	"github.com/labstack/gommon/log"
	"net/smtp"
)

type mail struct {
	To            []string
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

func (inst *Client) SendEmail(body *mail) {
	body, err := emailBody(body)
	to := body.To
	subject := body.Subject
	message := body.Message
	senderAddress := body.SenderAddress
	password := body.Password

	e := email.NewEmail()
	e.From = senderAddress
	e.To = to
	e.Subject = subject
	e.HTML = []byte(message)
	err = e.Send("smtp.gmail.com:587", smtp.PlainAuth("", senderAddress, password, "smtp.gmail.com"))
	if err != nil {
		log.Error(err)
		return
	} else {

	}
}
