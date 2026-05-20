package handler

import (
	"bytes"
	"encoding/json"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/JohnnyMcGee/website/api/contact"
	"github.com/JohnnyMcGee/website/api/contact/message"
	"github.com/JohnnyMcGee/website/api/contact/notifier"
	"github.com/stretchr/testify/require"
)

func TestContactHandler_SendMessage(t *testing.T) {
	cases := []struct {
		name         string
		input        string
		expectStatus int
		expectResult ContactResult
	}{
		{
			name:         "ValidInput",
			input:        `{ "name": "Bobby Bob", "email": "bobbybobtest0@proton.me", "company": "Bob's Business", "message": "Hello, this is a test message." }`,
			expectStatus: http.StatusOK,
			expectResult: ContactResult{Ok: true},
		},
		{
			name:         "MissingName",
			input:        `{}`,
			expectStatus: http.StatusBadRequest,
			expectResult: ContactResult{Ok: false, Error: "name is required"},
		},
		{
			name:         "MissingEmail",
			input:        `{"name": "Bob" }`,
			expectStatus: http.StatusBadRequest,
			expectResult: ContactResult{Ok: false, Error: "email is required"},
		},
		{
			name:         "MissingMessage",
			input:        `{"name": "Bob", "email": "bobbybobtest0@proton.me" }`,
			expectStatus: http.StatusBadRequest,
			expectResult: ContactResult{Ok: false, Error: "message is required"},
		},
		{
			name:         "InvalidEmail",
			input:        `{ "name": "Bobby Bob", "email": "bademail", "company": "Bob's Business", "message": "Hello, this is a test message." }`,
			expectStatus: http.StatusBadRequest,
			expectResult: ContactResult{Ok: false, Error: "please enter a valid email address"},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
			contactApi := contact.New([]notifier.Notifier{}, logger)
			contactHandler := NewContactHandler(contactApi, logger)

			req := httptest.NewRequest("POST", "/contact", bytes.NewReader([]byte(tc.input)))
			rr := httptest.NewRecorder()
			contactHandler.SendMessage(rr, req)
			require.Equal(t, tc.expectStatus, rr.Code)
			var result ContactResult
			err := json.NewDecoder(rr.Body).Decode(&result)
			require.NoError(t, err)
			require.Equal(t, tc.expectResult, result)
		})
	}
}

type MockNotifier struct {
	called bool
}

func (m *MockNotifier) ContactMessage(input message.MessageInput) error {
	m.called = true
	return nil
}

func TestContactHandler_CallNotifiers(t *testing.T) {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	mockNotifier := &MockNotifier{}
	contactApi := contact.New([]notifier.Notifier{mockNotifier}, logger)
	contactHandler := NewContactHandler(contactApi, logger)

	input := `{ "name": "Bobby Bob", "email": "bobbybobtest0@proton.me", "company": "Bob's Business", "message": "Hello, this is a test message." }`
	req := httptest.NewRequest("POST", "/contact", bytes.NewReader([]byte(input)))
	rr := httptest.NewRecorder()
	contactHandler.SendMessage(rr, req)
	require.True(t, mockNotifier.called, "Expected notifier to be called")
}
