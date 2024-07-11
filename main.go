package main

import (
	"fmt"
	"net/smtp"

)


var (
	fromGmail = "sami4@gmail.com"
	password  = "asqh ugfs xrsy nxqh"
)

func main() {

		to := []string{
			"kamal@gmail.com",
            "kate@gmail.com",
            "michael@gmail.com",
            "sarah@gmail.com",
            "tim@gmail.com",
			"jessicamengloa@gmail.com",
		}

			sendMail(to, "Hello,World!")
}

func sendMail(to []string ,sms string) {

	smtHost := "smtp.gmail.com"
	port := "587"
	byteM := []byte(sms)

	auth := smtp.PlainAuth("", fromGmail, password, smtHost)

	err := smtp.sendMail(smtHost+":"+port,auth,fromGmail,to,byteM)
		if err!= nil {
			fmt.Println("Error sending email:", err)
			return
		}
		fmt.Println("Email sent successfully!")
}