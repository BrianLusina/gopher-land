package globalnotifier

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type SlackNotifierDecorator struct {
	NotifierDecorator
	ChannelID string

}

type SlackMessageOpts struct {
	ChannelID string
	Message string
}

func NewSlackNotifier(notifier INotifier, channelID string) *SlackNotifierDecorator {
	return &SlackNotifierDecorator{NotifierDecorator{notifier}, channelID}
}

func (snd *SlackNotifierDecorator) Send(message string) {
	snd.NotifierDecorator.Send(message)
	snd.sendToSlack(message)
}

func (snd *SlackNotifierDecorator) sendToSlack(message string) {
	fmt.Println("Sending Message to Slack:", message)
	client := &http.Client{}
	reqBody, err := json.Marshal(SlackMessageOpts{snd.ChannelID, message})
	if err != nil {
		fmt.Printf("Error while marshalling: %+v\n", err)
		return
	}
	req, err := http.NewRequest("POST", "https://slack.com/api/chat.postMessage", bytes.NewBuffer(reqBody))
	if err != nil {
		fmt.Printf("Error while making request: %+v\n", err)
		return
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("SLACK_TOKEN")))
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	_, _ = client.Do(req)
}
