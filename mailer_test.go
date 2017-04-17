package mailer

import (
	"testing"
)

// Provide a valid smtp config
var (
	host     = "mail.xxx.com"
	user     = "user1@xxx.com"
	pass     = "xxx"
	recipent = "user2@xxx.com"
)

func TestSMTP(t *testing.T) {
	mail := NewMail()
	mail.FromName = "Go Mailer"
	mail.From = user
	mail.Subject = "Go SMTP Test"
	mail.Body = "This is a test mail from Go Mailer via SMTP"
	mail.SetTo(recipent)

	mailer := NewMailer(host, 25, user, pass, false)
	if err := mailer.Send(mail); err != nil {
		t.Errorf("Send error: %s", err)
	}
}

func TestSMTPSSL(t *testing.T) {
	mail := NewMail()
	mail.FromName = "Go Mailer"
	mail.From = user
	mail.Subject = "Go SMTP SSL Test"
	mail.Body = "This is a test from Go Mailer via SMTP SSL"
	mail.SetTo(recipent)

	mailer := NewMailer(host, 465, user, pass, true)
	if err := mailer.Send(mail); err != nil {
		t.Errorf("Send error: %s", err)
	}
}
