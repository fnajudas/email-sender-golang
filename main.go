package main

import (
	sendmail "email-sender-golang/sendMail"
	"log"
)

func main() {
	log.Println("========== START ==========")

	subject := "Test email sender"
	message := "Hello test"
	err := sendmail.SendMail(subject, message)
	if err != nil {
		log.Fatalf(err.Error())
	}

	log.Println("Email sent!")
}
