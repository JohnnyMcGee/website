package handler

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/JohnnyMcGee/website/api/contact"
	"github.com/JohnnyMcGee/website/api/contact/message"
)

type ContactHandler struct {
	contact *contact.Contact
	logger  *slog.Logger
}

type ContactResult struct {
	Ok    bool   `json:"ok"`
	Error string `json:"error,omitempty"`
}

func NewContactHandler(contact *contact.Contact, logger *slog.Logger) *ContactHandler {
	return &ContactHandler{
		contact: contact,
		logger:  logger,
	}
}

func (h *ContactHandler) SendMessage(w http.ResponseWriter, r *http.Request) {
	var input message.MessageInput

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		h.logger.Error("FailedToDecodeRequest", "error", err)
		result := ContactResult{Ok: false, Error: "Invalid request body"}
		if err := h.json(w, http.StatusBadRequest, result); err != nil {
			h.logger.Error("FailedToEncodeErrorResponse", "error", err)
		}
		return
	}

	if err := h.contact.SendMessage(input); err != nil {
		h.logger.Error("FailedToSendMessage", "error", err)

		switch err {
		case message.ErrNameRequired, message.ErrEmailRequired, message.ErrMessageRequired, message.ErrInvalidEmail:
			h.logger.Error("ValidationError", "error", err)
			result := ContactResult{Ok: false, Error: err.Error()}
			if err := h.json(w, http.StatusBadRequest, result); err != nil {
				h.logger.Error("FailedToEncodeErrorResponse", "error", err)
			}
			return
		default:
			result := ContactResult{Ok: false, Error: "Failed to process message"}
			if err := h.json(w, http.StatusInternalServerError, result); err != nil {
				h.logger.Error("FailedToEncodeErrorResponse", "error", err)
			}
			return
		}
	}

	h.logger.Info("ContactMessageSent")
	result := ContactResult{Ok: true}
	if err := h.json(w, http.StatusOK, result); err != nil {
		h.logger.Error("FailedToEncodeResponse", "error", err)
	}
}

func (h *ContactHandler) json(w http.ResponseWriter, status int, data any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)
}
