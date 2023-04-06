package main

import (
	"context"

	"os"
	"time"

	"github.com/aws/aws-lambda-go/lambda"

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

	images, err := getImages(os.Getenv("DYNAMO_TABLE_NAME"), ctx)
	if err != nil {
		return "", err
	}

	err = sendMessage(os.Getenv("SLACK_WEBHOOK_URL"), images)

	return "", err
}

func main() {
	lambda.Start(handleRequest)
}
