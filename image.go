package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type Image struct {
	Id  string `json:"id" dynamodbav:"id"`
	URL string `json:"url" dynamodbav:"url"`
}

func fetchImages(tableName string) ([]Image, error) {
	svc := dynamodb.New(session.New())
	result, err := svc.Scan(&dynamodb.ScanInput{
		TableName: aws.String(tableName),
	})

	if err != nil {
		return nil, err
	}

	var images []Image

	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &images)
	if err != nil {
		return nil, err
	}

	return images, nil
}
