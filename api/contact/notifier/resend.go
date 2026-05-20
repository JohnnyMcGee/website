package notifier

import (
	"log/slog"

	"github.com/JohnnyMcGee/website/api/config"
	"github.com/JohnnyMcGee/website/api/contact/message"
	"github.com/resend/resend-go/v3"
)

type ResendNotifier struct {
	config config.ResendNotifierConfig
	client *resend.Client
	logger *slog.Logger
}

func NewResendNotifier(client *resend.Client, config config.ResendNotifierConfig, logger *slog.Logger) *ResendNotifier {
	return &ResendNotifier{
		config: config,
		client: client,
		logger: logger,
	}
}

func (n *ResendNotifier) ContactMessage(input message.MessageInput) error {
	n.logger.Info("ResendNotifierContactMessage", "name", input.Name, "email", input.Email, "company", input.Company)

	params := &resend.SendEmailRequest{
		From:    n.config.FromEmail,
		To:      []string{n.config.ToEmail},
		Subject: formatSubject(input),
		Html:    formatEmail(input),
	}

	sent, err := n.client.Emails.Send(params)
	if err != nil {
		n.logger.Error("failed to send email", "error", err)
		return err
	}

	n.logger.Info("email sent", "email_id", sent.Id)

	return nil
}

func formatSubject(input message.MessageInput) string {
	var company string
	if input.Company != nil {
		company = " @ " + *input.Company
	}
	return "Contact Form Submission from " + input.Name + company
}

func formatEmail(input message.MessageInput) string {
	var company string
	if input.Company != nil {
		company = *input.Company
	}
	return "<p><strong>Name:</strong> " + input.Name + "</p>" +
		"<p><strong>Email:</strong> " + input.Email + "</p>" +
		"<p><strong>Company:</strong> " + company + "</p>" +
		"<p><strong>Message:</strong><br>" + input.Message + "</p>"

}
