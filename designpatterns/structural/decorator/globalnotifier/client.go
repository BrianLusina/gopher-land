package globalnotifier

func main() {
	baseNotifier := &BaseNotifier{}

	slackNotifier := NewSlackNotifier(baseNotifier, "slack111")
	telegramNotifier := NewTelegramNotifier(baseNotifier, "tele111")
	globalNotifier := NewSlackNotifier(telegramNotifier, "slack111") //or can also be NewTelegramNotifier(slackNotifier, "tele111")

	slackNotifier.Send("Only sent to Slack!")
	telegramNotifier.Send("Only sent to Telegram!")
	globalNotifier.Send("Send to both Slack and Telegram!")
}

