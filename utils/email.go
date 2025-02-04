package utils

import (
	"gopkg.in/gomail.v2"
)

func SendReceipt(email, filePath string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", "sabdpp17@gmail.com")
	m.SetHeader("To", email)
	m.SetHeader("Subject", "Ваш чек за премиум подписку")
	m.SetBody("text/plain", "Спасибо за оплату! Ваш чек во вложении.")
	m.Attach(filePath)

	d := gomail.NewDialer("smtp.gmail.com", 587, "sabdpp17@gmail.com", "vili ubad mijt vbqo")
	return d.DialAndSend(m)
}
