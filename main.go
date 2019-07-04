package main

import (
	"log"
	"os"
	"time"

	"github.com/nlopes/slack"
)

func main() {
	args := os.Args
	slackToken := args[1]
	channel := args[2]

	api := slack.New(slackToken)

	params := slack.HistoryParameters{
		Count: 1000,
	}
	history, err := api.GetChannelHistory(channel, params)
	if err != nil {
		log.Fatal(err)
	}

	for _, message := range history.Messages {
		api.DeleteMessage(channel, message.Timestamp)
		time.Sleep(840 * time.Millisecond)
	}
}
