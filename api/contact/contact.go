package contact

import (
	"errors"
	"log/slog"
	"net/mail"
)

type Contact struct {
	logger *slog.Logger
}

type ContactMessageInput struct {
	Name    string  `json:"name"`
	Email   string  `json:"email"`
	Company *string `json:"company"`
	Message string  `json:"message"`
}

func New(logger *slog.Logger) *Contact {
	return &Contact{
		logger: logger,
	}
}

func (c *Contact) SendMessage(input ContactMessageInput) error {
	c.logger.Info("ContactMessageReceived", "name", input.Name, "email", input.Email, "company", input.Company)
	if err := validateInput(input); err != nil {
		return err
	}
	return nil
}

var (
	ErrNameRequired    = errors.New("name is required")
	ErrEmailRequired   = errors.New("email is required")
	ErrMessageRequired = errors.New("message is required")
	ErrInvalidEmail    = errors.New("please enter a valid email address")
)

func validateInput(input ContactMessageInput) error {
	if input.Name == "" {
		return ErrNameRequired
	}
	if input.Email == "" {
		return ErrEmailRequired
	}
	if input.Message == "" {
		return ErrMessageRequired
	}
	if _, err := mail.ParseAddress(input.Email); err != nil {
		return ErrInvalidEmail
	}
	return nil
}
