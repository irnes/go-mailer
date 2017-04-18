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
	config := Config{
		Host: host,
		Port: 25,
		User: user,
		Pass: pass,
	}

	mailer := NewMailer(config, false)

	mail := NewMail()
	mail.FromName = "Go Mailer"
	mail.From = user
	mail.SetTo(recipent)

	mail.Subject = "Go SMTP Test"
	mail.Body = "This is a test e-mail from Go Mailer via SMTP"

	if err := mailer.Send(mail); err != nil {
		t.Errorf("Send error: %s", err)
	}
}

func TestSMTPSSL(t *testing.T) {
	config := Config{
		Host: host,
		Port: 465,
		User: user,
		Pass: pass,
	}

	mailer := NewMailer(config, true)

	mail := NewMail()
	mail.FromName = "Go Mailer"
	mail.From = user
	// multiple recipents
	mail.SetTo(recipent)
	mail.SetTo(recipent2)

	mail.Subject = "Go SMTP SSL Test"
	// the rich message body
	mail.Body = `
		<h2>Hello from Go Mailer</h2>
		<p>This is a test e-mail from Go Mailer via SMTP SSL</p>
	`

	if err := mailer.Send(mail); err != nil {
		t.Errorf("Send error: %s", err)
	}
}
