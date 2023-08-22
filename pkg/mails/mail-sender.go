package mail

import (
	"fmt"
	"net/smtp"

	"github.com/jordan-wright/email"
)

const (
	SMTPAuthAddress   = "smtp.gmail.com"
	SMTPServerAddress = "smtp.gmail.com:587"
)

type Sender interface {
	SendMail(
		subject, content string,
		to, cc, bcc, attachFiles []string,
	) error
}

type MailSender struct {
	name, fromMailAddress, fromMailPassword string
}

func NewMailSender(name string, fromMailAddress string, fromMailPassword string) Sender {
	return &MailSender{
		name:             name,
		fromMailAddress:  fromMailAddress,
		fromMailPassword: fromMailPassword,
	}
}

func (sender *MailSender) SendMail(subject string, content string, to []string, cc []string, bcc []string, attachFiles []string) error {
	e := email.NewEmail()
	e.From = fmt.Sprintf("%s <%s>", sender.name, sender.fromMailAddress)
	e.Subject = subject
	e.HTML = []byte(content)
	e.To = to
	e.Cc = cc
	e.Bcc = bcc

	for _, f := range attachFiles {
		_, err := e.AttachFile(f)
		if err != nil {
			return fmt.Errorf("Failed to attach file %s: %w", f, err)
		}
	}

	smtpAuth := smtp.PlainAuth("", sender.fromMailAddress, sender.fromMailPassword, SMTPAuthAddress)
	return e.Send(SMTPServerAddress, smtpAuth)
}
