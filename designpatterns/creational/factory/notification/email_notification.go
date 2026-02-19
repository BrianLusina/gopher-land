package notification

type EmailNotification struct{}

// Send sends the email notification with the given message.
func (e *EmailNotification) Send(message string) error {
	// Here you would implement the logic to send an email notification.
	// For demonstration purposes, we'll just print the message.
	println("Sending email notification with message:", message)
	return nil
}
