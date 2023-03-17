package main

import (
	"context"

	"os"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws/session"
	utils "github.com/jwilsson/go-bot-utils"
)

type Image struct {
	URL string `json:"image_url" dynamodbav:"image_url"`
}

const fallbackText = "Det är fredag mina bekanta"

func handleRequest(ctx context.Context) (string, error) {
	if !utils.ShouldRun(time.Now(), os.Getenv("RUN_AT_TIME"), os.Getenv("TARGET_TIMEZONE")) {
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

	err = sendMessage(os.Getenv("SLACK_WEBHOOK_URL"), images)

	return "", err
}

func main() {
	lambda.Start(handleRequest)
}
