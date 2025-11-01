package main

import (
	"math/rand"

	"github.com/slack-go/slack"
)

func sendMessage(webHookUrl string, images []Image) error {
	index := rand.Intn(len(images))
	block := slack.NewImageBlock(images[index].URL, fallbackText, "", nil)

	return slack.PostWebhook(
		webHookUrl,
		&slack.WebhookMessage{
			Blocks: &slack.Blocks{
				BlockSet: []slack.Block{block},
			},
			Text: fallbackText,
		},
	)
}
