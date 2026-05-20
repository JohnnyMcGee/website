package message

import (
	"errors"
	"net/mail"
)

type MessageInput struct {
	Name    string  `json:"name"`
	Email   string  `json:"email"`
	Company *string `json:"company,omitempty"`
	Message string  `json:"message"`
}

var (
	ErrNameRequired    = errors.New("name is required")
	ErrEmailRequired   = errors.New("email is required")
	ErrMessageRequired = errors.New("message is required")
	ErrInvalidEmail    = errors.New("please enter a valid email address")
)

func (i *MessageInput) Validate() error {
	if i.Name == "" {
		return ErrNameRequired
	}
	if i.Email == "" {
		return ErrEmailRequired
	}
	if i.Message == "" {
		return ErrMessageRequired
	}
	if _, err := mail.ParseAddress(i.Email); err != nil {
		return ErrInvalidEmail
	}
	return nil
}
