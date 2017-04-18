package mailer

import (
	"crypto/tls"
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
func NewMailer(config Config, ssl bool) Mailer {
	var mailer Mailer
	if ssl {
		smtpssl := new(SMTPSLL)
		smtpssl.Config = config
		mailer = smtpssl
	} else {
		smtp := new(SMTP)
		smtp.Config = config
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

	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.
	err = smtp.SendMail(server, auth, mail.From, mail.To, mail.Raw())

	return
}

// SMTPSLL client session object for situations where SSL is required from
// the beginning of the connection and using STARTTLS is not appropriate
type SMTPSLL struct {
	Config
}

// Send an email and waits for the process to end, giving proper error feedback
func (m *SMTPSLL) Send(mail *Mail) (err error) {
	// TLS config
	tlsconfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         m.Host,
	}

	// Call tls.Dial instead of smtp.Dial for smtp servers running on 465
	// that require an ssl connection from the very beginning (no starttls)
	server := fmt.Sprintf("%s:%d", m.Host, m.Port)
	conn, err := tls.Dial("tcp", server, tlsconfig)
	if err != nil {
		return
	}
	defer conn.Close()

	c, err := smtp.NewClient(conn, m.Host)
	if err != nil {
		return
	}
	defer c.Quit()

	// Set up authentication information.
	auth := smtp.PlainAuth("", m.User, m.Pass, m.Host)
	if err = c.Auth(auth); err != nil {
		return
	}

	// Set sender and recipents
	if err = c.Mail(mail.From); err != nil {
		return
	}
	for _, rcpt := range mail.To {
		c.Rcpt(rcpt)
	}

	// Data
	w, err := c.Data()
	if err != nil {
		return
	}
	defer w.Close()

	_, err = w.Write(mail.Raw())
	if err != nil {
		return
	}

	return nil
}
