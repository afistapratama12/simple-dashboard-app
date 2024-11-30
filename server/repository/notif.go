package repository

import (
	"fmt"
	"simple-dashboard-server/config"

	"gopkg.in/gomail.v2"
)

type NotifRepo interface {
	NotifEmail(subject string, to []string, body string) error
}

type notifRepo struct {
	env config.ENV
}

func NewNotifRepo(env config.ENV) NotifRepo {
	return &notifRepo{
		env: env,
	}
}

func (r *notifRepo) NotifEmail(subject string, to []string, body string) error {
	msg := gomail.NewMessage()
	msg.SetHeader("From", fmt.Sprintf("Simple Dashboard Admin <%s>", r.env.FromEmail))
	msg.SetHeader("To", to...)
	msg.SetHeader("Subject", subject)
	msg.SetBody("text/html", body)

	d := gomail.NewDialer("smtp.gmail.com", 587, r.env.EmailUsername, r.env.AppPass)

	if err := d.DialAndSend(msg); err != nil {
		return err
	}

	return nil
}
