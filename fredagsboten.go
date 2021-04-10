package main

import (
	"context"
	"math/rand"
	"os"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
	utils "github.com/jwilsson/go-bot-utils"
	"github.com/slack-go/slack"
)

const fallbackText = "Det är fredag mina bekanta"

func handleRequest(ctx context.Context) (string, error) {
	if utils.IsHoliday(time.Now()) {
		return "", nil
	}

	rand.Seed(time.Now().UnixNano())

	images, err := fetchImages(os.Getenv("DYNAMO_TABLE_NAME"))
	if err != nil {
		return "", err
	}

	index := rand.Intn(len(images))
	block := slack.NewImageBlock(images[index].URL, fallbackText, "", nil)

	err = slack.PostWebhook(
		os.Getenv("SLACK_WEBHOOK_URL"),
		&slack.WebhookMessage{
			Blocks: &slack.Blocks{
				BlockSet: []slack.Block{block},
			},
			Text: fallbackText,
		},
	)

	return "", err
}

func main() {
	lambda.Start(handleRequest)
}
