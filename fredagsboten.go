package main

import (
	"context"
	"math/rand"
	"os"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws/session"
	utils "github.com/jwilsson/go-bot-utils"
	"github.com/slack-go/slack"
)

type Image struct {
	URL string `json:"image_url" dynamodbav:"image_url"`
}

const fallbackText = "Det är fredag mina bekanta"

func handleRequest(ctx context.Context) (string, error) {
	now := time.Now()

	if utils.IsHoliday(now) {
		return "", nil
	}

	s, err := session.NewSession()
	if err != nil {
		return "", err
	}

	var images []Image

	err = utils.GetDynamodbData(s, os.Getenv("DYNAMO_TABLE_NAME"), &images)
	if err != nil {
		return "", err
	}

	rand.Seed(now.UnixNano())

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
