package globalnotifier

type INotifier interface {
	Send(message string)
}

type BaseNotifier struct {
	message string
}

func (b *BaseNotifier) Send(message string) {
	b.message = message
}

type NotifierDecorator struct {
	notifier INotifier
}

func (nd *NotifierDecorator) Send(message string) {
	nd.notifier.Send(message)
}
