package mailer

import (
	"testing"
)

// Provide a valid smtp config
var (
	host      = "mail.xxx.com"
	user      = "send@xxx.com"
	pass      = "xxx"
	recipent  = "rcpt1@xxx.com"
	recipent2 = "rcpt2@xxx.com"
)

func TestSMTP(t *testing.T) {
	mail := NewMail()
	mail.FromName = "Go Mailer"
	mail.From = user
	mail.Subject = "Go SMTP Test"
	mail.Body = "This is a test e-mail from Go Mailer via SMTP"
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
	// the rich message body
	mail.Body = `
		<h2>Hello from Go Mailer</h2>
		<p>
			This is a test e-mail from Go Mailer via SMTP SSL
		</p>
	`

	// multiple recipents
	mail.SetTo(recipent)
	mail.SetTo(recipent2)

	mailer := NewMailer(host, 465, user, pass, true)
	if err := mailer.Send(mail); err != nil {
		t.Errorf("Send error: %s", err)
	}
}
