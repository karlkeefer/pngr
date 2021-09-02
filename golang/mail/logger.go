package mail

import "fmt"

// LogMail implements Mailer
type LogMail struct{}

func (m *LogMail) Send(to string, subject string, plain string, h *html) error {
	fmt.Printf("To: %s\nSubject: %s\n%s\n\n", to, subject, plain)
	return nil
}
