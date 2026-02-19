package notification

// CreateNotification is a factory function that creates and returns a Notification based on the provided notificationType.
func CreateNotification(notificationType string) Notification {
	switch notificationType {
	case "email":
		return &EmailNotification{}
	case "sms":
		return &SmsNotification{}
	default:
		return nil
	}
}
