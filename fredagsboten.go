package main

import (
	"context"
	"math/rand"
	"os"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/slack-go/slack"
)

func handleRequest(ctx context.Context) (string, error) {
	rand.Seed(time.Now().UnixNano())

	images, err := fetchImages(os.Getenv("DYNAMO_TABLE_NAME"))
	if err != nil {
		return "", err
	}

	index := rand.Intn(len(images))
	attachment := slack.Attachment{
		Fallback: "Det är fredag mina bekanta",
		ImageURL: images[index].URL,
	}

	err = slack.PostWebhook(
		os.Getenv("SLACK_WEBHOOK_URL"),
		&slack.WebhookMessage{
			Attachments: []slack.Attachment{attachment},
		},
	)

	return "", err
}

func main() {
	lambda.Start(handleRequest)
}
