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
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type Attachment struct {
	Fallback string `json:"fallback"`
	ImageURL string `json:"image_url"`
}

type Body struct {
	Attachments []*Attachment `json:"attachments"`
}

type Image struct {
	Id  string `json:"id" dynamodbav:"id"`
	Url string `json:"url" dynamodbav:"url"`
}

func handleRequest(ctx context.Context) (string, error) {
	rand.Seed(time.Now().UnixNano())

	svc := dynamodb.New(session.New())
	result, err := svc.Scan(&dynamodb.ScanInput{
		TableName: aws.String(os.Getenv("DYNAMO_TABLE_NAME")),
	})

	if err != nil {
		return "", err
	}

	var images []Image
	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &images)

	if err != nil {
		return "", err
	}

	jsonBody, _ := json.Marshal(Body{
		Attachments: []*Attachment{
			&Attachment{
				Fallback: "Det är fredag mina bekanta",
				ImageURL: images[rand.Intn(len(images))].Url,
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
