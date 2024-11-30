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
	env    config.ENV
	dialer Dialer
}

func NewNotifRepo(env config.ENV, dialer Dialer) NotifRepo {
	return &notifRepo{
		env:    env,
		dialer: dialer,
	}
}

type Dialer interface {
	DialAndSend(m ...*gomail.Message) error
}

func (r *notifRepo) NotifEmail(subject string, to []string, body string) error {
	msg := gomail.NewMessage()
	msg.SetHeader("From", fmt.Sprintf("Simple Dashboard Admin <%s>", r.env.FromEmail))
	msg.SetHeader("To", to...)
	msg.SetHeader("Subject", subject)
	msg.SetBody("text/html", body)

	if err := r.dialer.DialAndSend(msg); err != nil {
		return err
	}

	return nil
}
