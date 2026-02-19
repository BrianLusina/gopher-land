package notification

// Notification is the interface that all notification types must implement
type Notification interface {
	// Send sends the notification with the given message and returns an error if it fails
	Send(message string) error
}
