package repository

import (
	"simple-dashboard-server/config"
	"testing"

	"gopkg.in/gomail.v2"
)

type MockDialer struct {
	Host     string
	Port     int
	Username string
	Password string
	SSL      bool
}

func (d *MockDialer) DialAndSend(m ...*gomail.Message) error {
	return nil
}

func TestNotif_NotifEmail(t *testing.T) {
	type arg struct {
		subject string
		to      []string
		body    string
	}

	envMock := config.ENV{
		EmailUsername: "test",
		AppPass:       "test",
		FromEmail:     "test",
	}

	t.Run("success", func(t *testing.T) {
		args := arg{
			subject: "Test",
			to:      []string{"mail@mail.com"},
			body:    "Test",
		}

		notifRepo := NewNotifRepo(envMock, &MockDialer{
			Host:     "smtp.gmail.com",
			Port:     587,
			Username: "test",
			Password: "test",
			SSL:      true,
		})

		err := notifRepo.NotifEmail(args.subject, args.to, args.body)
		if err != nil {
			t.Errorf("Expecting nil, got %v", err)
		}
	})

}
