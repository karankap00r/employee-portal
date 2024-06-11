package service

import "net/smtp"

// sendEmail sends an email
func sendEmail(to, subject, body string) error {
	from := "your-email@example.com"
	password := "your-email-password"

	smtpHost := "smtp.example.com"
	smtpPort := "587"

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: " + subject + "\n\n" + body

	auth := smtp.PlainAuth("", from, password, smtpHost)

	return smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{to}, []byte(msg))
}
