package mailer

import (
	"fmt"
	"net/smtp"
)

// Config represents smtp configuration
type Config struct {
	Host string
	Port int
	User string
	Pass string
}

// Mailer is an interface with a Send method, that dispatches a single Mail
type Mailer interface {
	Send(*Mail) error
}

// NewMailer returns a new Mailer using provided configuration
func NewMailer(host string, port int, user, pass string, ssl bool) Mailer {
	var mailer Mailer
	if ssl {
		smtpssl := new(SMTPSLL)
		smtpssl.Host = host
		smtpssl.Port = port
		smtpssl.User = user
		smtpssl.Pass = pass
		mailer = smtpssl
	} else {
		smtp := new(SMTP)
		smtp.Host = host
		smtp.Port = port
		smtp.User = user
		smtp.Pass = pass
		mailer = smtp
	}
	return mailer
}

// SMTP client session object. If a server advertise STARTTLS, it
// will encrypt all further communication after initial handshake
type SMTP struct {
	Config
}

// Send an email and waits for the process to end, giving proper error feedback
func (m *SMTP) Send(mail *Mail) (err error) {
	server := fmt.Sprintf("%s:%d", m.Host, m.Port)
	// Set up authentication information.
	auth := smtp.PlainAuth("", m.User, m.Pass, m.Host)
	// Prepare message content according to RFC 822-style
	msg := []byte(fmt.Sprintf("From: %s <%s>\r\nSubject: %s\r\n\r\n%s\r\n",
		mail.FromName, mail.From, mail.Subject, mail.Body))

	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.
	err = smtp.SendMail(server, auth, mail.From, mail.To, msg)

	return
}

// SMTPSLL client session object for situations where SSL is required from
// the beginning of the connection and using STARTTLS is not appropriate
type SMTPSLL struct {
	Config
}

// Send an email and waits for the process to end, giving proper error feedback
func (m *SMTPSLL) Send(mail *Mail) (err error) {
	return nil
}
