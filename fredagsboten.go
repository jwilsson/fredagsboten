package main

import (
	"bytes"
	"context"
	"encoding/json"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
)

type Attachment struct {
	Fallback string `json:"fallback"`
	ImageURL string `json:"image_url"`
}

type Body struct {
	Attachments []*Attachment `json:"attachments"`
}

func handleRequest(ctx context.Context) (string, error) {
	rand.Seed(time.Now().UnixNano())

	images, err := fetchImages(os.Getenv("DYNAMO_TABLE_NAME"))
	if err != nil {
		return "", err
	}

	index := rand.Intn(len(images))
	resBody, err := json.Marshal(Body{
		Attachments: []*Attachment{
			&Attachment{
				Fallback: "Det är fredag mina bekanta",
				ImageURL: images[index].URL,
			},
		},
	})

	if err != nil {
		return "", err
	}

	http.Post(
		os.Getenv("SLACK_WEBHOOK_URL"),
		"application/json",
		bytes.NewBuffer(resBody),
	)

	return "", nil
}

func main() {
	lambda.Start(handleRequest)
}
