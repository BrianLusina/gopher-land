package notification

type SmsNotification struct{}

// Send sends the sms notification with the given message.
func (e *SmsNotification) Send(message string) error {
	// Here you would implement the logic to send an sms notification.
	// For demonstration purposes, we'll just print the message.
	println("Sending sms notification with message:", message)
	return nil
}
