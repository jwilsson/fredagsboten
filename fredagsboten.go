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

	images := [3]string{
		"https://i.imgur.com/YdiFhzm.png",
		"https://i.imgur.com/XPF1jBP.png",
		"https://i.imgur.com/pqKmIdJ.png",
	}

	jsonBody, _ := json.Marshal(Body{
		Attachments: []*Attachment{
			&Attachment{
				Fallback: "Det är fredag mina bekanta",
				ImageURL: images[rand.Intn(len(images))],
			},
		},
	})

	http.Post(
		os.Getenv("SLACK_WEBHOOK_URL"),
		"application/json",
		bytes.NewBuffer(jsonBody),
	)

	return "", nil
}

func main() {
	lambda.Start(handleRequest)
}
