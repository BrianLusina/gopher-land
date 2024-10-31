package globalnotifier

import (
	"fmt"
	"net/http"
	"os"
)

type TelegramNotifierDecorator struct {
	NotifierDecorator
	ChatID string
}

func NewTelegramNotifier(notifier INotifier, chatID string) *TelegramNotifierDecorator {
	return &TelegramNotifierDecorator{NotifierDecorator{notifier}, chatID}
}

func (tnd *TelegramNotifierDecorator) Send(message string) {
	tnd.NotifierDecorator.Send(message)
	tnd.sendToTelegram(message)
}

func (tnd *TelegramNotifierDecorator) sendToTelegram(message string) {
	fmt.Println("Sending Message to Telegram:", message)
	client := &http.Client{}
	botToken := os.Getenv("TELEGRAM_TOKEN")
	req, err := http.NewRequest("GET", fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", botToken), nil)
	if err != nil {
		fmt.Printf("Error while making request: %+v\n", err)
		return
	}
	q := req.URL.Query()
	q.Add("chat_id", tnd.ChatID)
	q.Add("text", message)
	req.URL.RawQuery = q.Encode()
	_, _ = client.Do(req)
}
