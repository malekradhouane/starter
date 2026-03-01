package mailer

import (
	"context"
	"fmt"

	"github.com/mailjet/mailjet-apiv3-go/v4"
)

// Mailer defines minimal email sending capability
// Implementations should be safe for concurrent use
// and respect context cancellation when possible.
type Mailer interface {
	Send(ctx context.Context, fromName, fromEmail, toName, toEmail, subject, textPart, htmlPart string) error
}

// Config holds Mailjet credentials and defaults
// All fields are required to send emails.
type Config struct {
	APIKeyPublic  string
	APIKeyPrivate string
	DefaultFrom   struct {
		Name  string
		Email string
	}
}

// NewMailjet returns a Mailer backed by Mailjet's API client
func NewMailjet(cfg Config) (Mailer, error) {
	if cfg.APIKeyPublic == "" || cfg.APIKeyPrivate == "" {
		return nil, fmt.Errorf("mailjet: API keys are required")
	}
	if cfg.DefaultFrom.Email == "" {
		return nil, fmt.Errorf("mailjet: default from email is required")
	}
	return &mailjetMailer{client: mailjet.NewMailjetClient(cfg.APIKeyPublic, cfg.APIKeyPrivate), cfg: cfg}, nil
}

type mailjetMailer struct {
	client *mailjet.Client
	cfg    Config
}

func (m *mailjetMailer) Send(ctx context.Context, fromName, fromEmail, toName, toEmail, subject, textPart, htmlPart string) error {
	// fallbacks to defaults if not provided
	if fromEmail == "" {
		fromEmail = m.cfg.DefaultFrom.Email
	}
	if fromName == "" {
		fromName = m.cfg.DefaultFrom.Name
	}
	if toEmail == "" {
		return fmt.Errorf("mailjet: toEmail is required")
	}
	if subject == "" {
		return fmt.Errorf("mailjet: subject is required")
	}

	messagesInfo := []mailjet.InfoMessagesV31{
		{
			From: &mailjet.RecipientV31{
				Email: fromEmail,
				Name:  fromName,
			},
			To: &mailjet.RecipientsV31{
				mailjet.RecipientV31{
					Email: toEmail,
					Name:  toName,
				},
			},
			Subject:  subject,
			TextPart: textPart,
			HTMLPart: htmlPart,
		},
	}

	messages := mailjet.MessagesV31{Info: messagesInfo}

	// Mailjet client doesn't accept context directly; in case of cancellation,
	// we can pre-check and return early.
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	_, err := m.client.SendMailV31(&messages)
	return err
}
