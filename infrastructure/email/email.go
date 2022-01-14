package email

import (
	"github.com/smokers10/go-diagem.git/infrastructure/config"
	"gopkg.in/gomail.v2"
)

func Fire(to []string) error {
	config := config.ReadConfig().SMTP

	// message & body
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", config.SMTP_Sender_Name)
	mailer.SetHeader("To", to...)
	mailer.SetHeader("Subject", "Test mail")
	mailer.SetBody("text/html", "Hello, <b>have a nice day</b>")

	// dialer
	dialer := gomail.NewDialer(
		config.SMTP_Host,
		config.SMTP_Port,
		config.SMTP_Auth_Username,
		config.SMTP_Auth_Password,
	)

	// kirim
	err := dialer.DialAndSend(mailer)
	if err != nil {
		return err
	}

	return nil
}
