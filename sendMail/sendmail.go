package sendmail

import (
	"fmt"
	"net/smtp"
	"strings"
)

const (
	CONFIG_SENDER_NAME   = "test send email <EMAIL_PENGIRIM>"
	CONFIG_AUTH_EMAIL    = "EMAIL_PENGIRIM"
	CONFIG_AUTH_PASSWORD = "SECRET_PASSWORD"
	CONFIG_SMTP_HOST     = "smtp.gmail.com"
	CONFIG_SMTP_PORT     = 587
)

func SendMail(subject, message string) error {
	to := []string{"EMAIL_TUJUAN"}
	cc := []string{"EMAIL_TUJUAN_CC"}
	body := "From : " + CONFIG_SENDER_NAME + "\n" +
		"To: " + strings.Join(to, ",") + "\n" +
		"Cc: " + strings.Join(to, ",") + "\n" +
		"Subject: " + subject + "\n\n" +
		message

	auth := smtp.PlainAuth("", CONFIG_AUTH_EMAIL, CONFIG_AUTH_PASSWORD, CONFIG_SMTP_HOST)
	smtpAddr := fmt.Sprintf("%s:%d", CONFIG_SMTP_HOST, CONFIG_SMTP_PORT)

	err := smtp.SendMail(smtpAddr, auth, CONFIG_AUTH_EMAIL, append(to, cc...), []byte(body))
	if err != nil {
		return err
	}

	return nil
}
