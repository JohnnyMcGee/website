package notifier

import "github.com/JohnnyMcGee/website/api/contact/message"

type Notifier interface {
	ContactMessage(input message.MessageInput) error
}
