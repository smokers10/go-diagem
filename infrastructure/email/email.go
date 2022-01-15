package email

import (
	"fmt"
	"net/smtp"

	"github.com/smokers10/go-diagem.git/infrastructure/config"
	"gopkg.in/gomail.v2"
)

type SMTPEmail struct{}

func SMTP() *SMTPEmail {
	return &SMTPEmail{}
}

func (s *SMTPEmail) Fire(to []string, subject string, template string) error {
	config := config.ReadConfig().SMTP

	// message & body
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", config.SMTP_Sender_Name)
	mailer.SetHeader("To", to...)
	mailer.SetHeader("Subject", subject)
	mailer.SetBody("text/html", template)

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

func (s *SMTPEmail) NativeFire(to []string, subjectEmail string, template string) error {
	config := config.ReadConfig().SMTP

	addr := fmt.Sprintf("%s:%d", config.SMTP_Host, config.SMTP_Port)

	auth := smtp.PlainAuth("", config.SMTP_Auth_Username, config.SMTP_Auth_Password, config.SMTP_Host)

	// construct email body
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	subject := fmt.Sprintf("Subject: %s \n", subjectEmail)
	message := []byte(subject + mime + template)

	err := smtp.SendMail(addr, auth, config.SMTP_Auth_Username, to, message)
	if err != nil {
		return err
	}

	return nil
}

func (s *SMTPEmail) VerificationTemplate(to string, code string) string {
	tmplt := fmt.Sprintf(`
	<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<meta http-equiv="X-UA-Compatible" content="IE=edge">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>Diagem</title>
		<style>
			a{
				background-color: #f0ac3f;
				border-color: #e09112;
				border: none;
				color: white;
				padding: 15px 32px;
				text-align: center;
				text-decoration: none;
				display: inline-block;
				font-size: 16px;
				font-weight: 600;
			}
			body{
				font-family: "Nunito Sans", -apple-system, system-ui, BlinkMacSystemFont, "Segoe UI", Roboto, "Helvetica Neue", Arial, "Noto Sans", sans-serif, "Apple Color Emoji", "Segoe UI Emoji", "Segoe UI Symbol";
				font-size: 16px;
			}
		</style>
	</head>
	<body>
		<center>
			<h1>Hallo, %s</h1>   
			<p>Silahkan gunakan kode verifikasi dibawah pada halaman verifikasi</p>
			<h2>%s</h2>
			<p>jangan sebarluaskan kode verifikasi pada siapapun termasuk pihak kami,jika anda tidak request email ini, abaikan!</p>
			<p>terimakasih! <br> Diagem </p>
		</center>
	</body>
	</html>
	`, to, code)

	return tmplt
}

func (s *SMTPEmail) ForgotPasswordTemplate(to string, code string, redirect string) string {
	tmplt := fmt.Sprintf(`
	<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<meta http-equiv="X-UA-Compatible" content="IE=edge">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>Diagem</title>
		<style>
			a{
				background-color: #f0ac3f;
				border-color: #e09112;
				border: none;
				color: white;
				padding: 15px 32px;
				text-align: center;
				text-decoration: none;
				display: inline-block;
				font-size: 16px;
				font-weight: 600;
			}
			body{
				font-family: "Nunito Sans", -apple-system, system-ui, BlinkMacSystemFont, "Segoe UI", Roboto, "Helvetica Neue", Arial, "Noto Sans", sans-serif, "Apple Color Emoji", "Segoe UI Emoji", "Segoe UI Symbol";
				font-size: 16px;
			}
		</style>
	</head>
	<body>
		<center>
			<h1>Hallo, %s</h1>   
			<p>Silahkan gunakan kode reset password dibawah ini pada halaman reset password</p>
			<h2>%s</h2>
			<a href="%s">Reset Password</a>
			<p>jangan membagikan kode reset password pada siapapun termasuk pihak kami, jika anda tidak request email ini, abaikan!</p>
			<p>terimakasih! <br> Diagem </p>
		</center>
	</body>
	</html>
	`, to, code, redirect)

	return tmplt
}
