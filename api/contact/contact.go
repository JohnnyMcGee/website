package contact

import (
	"errors"
	"log/slog"

	"github.com/JohnnyMcGee/website/api/contact/message"
	"github.com/JohnnyMcGee/website/api/contact/notifier"
)

var (
	ErrFailedToSend = errors.New("failed to send message")
)

type Contact struct {
	logger    *slog.Logger
	notifiers []notifier.Notifier
}

func New(notifiers []notifier.Notifier, logger *slog.Logger) *Contact {
	return &Contact{
		notifiers: notifiers,
		logger:    logger,
	}
}

func (c *Contact) SendMessage(input message.MessageInput) error {
	c.logger.Info("ContactMessageReceived")
	if err := input.Validate(); err != nil {
		return err
	}

	errs := make([]error, 0)
	for _, n := range c.notifiers {
		if err := n.ContactMessage(input); err != nil {
			errs = append(errs, err)
		}
	}

	if len(errs) > 0 {
		c.logger.Error("FailedToSendMessage", "errors", errs)
		return ErrFailedToSend
	}

	return nil
}
