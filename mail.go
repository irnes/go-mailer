package mailer

import ()

// Mail struct represents an e-Mail
type Mail struct {
	// Sender Name
	FromName string
	// Sender email address
	From string

	// List of recipients
	To  []string
	Cc  []string
	Bcc []string

	// Mail subject as UTF-8 string
	Subject string

	// Headers are the headers
	Headers map[string]string

	// Body provides the actual body of the mail.
	// It has to be UTF-8 encoded, or you must set the Content-Type Header
	Body string
}

// NewMail returns a new Mail struct with Headers initialized to an empty map
func NewMail() *Mail {
	m := new(Mail)
	m.Headers = make(map[string]string)
	return m
}

// SetTo sets mail TO recepients
func (m *Mail) SetTo(addresses ...string) {
	m.To = sliceIt(m.To, addresses)
}

// SetCc sets mail CC recepients
func (m *Mail) SetCc(addresses ...string) {
	m.Cc = sliceIt(m.Cc, addresses)
}

// SetBcc sets mail BCC recepients
func (m *Mail) SetBcc(addresses ...string) {
	m.Bcc = sliceIt(m.Bcc, addresses)
}

// SetHeader sets mail custom headers
func (m *Mail) SetHeader(k, v string) {
	m.Headers[k] = v
}

func sliceIt(slice, add []string) []string {
	if len(slice) == 0 {
		return add
	}
	for _, a := range add {
		slice = append(slice, a)
	}
	return slice
}
