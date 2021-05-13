package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type Image struct {
	URL string `json:"image_url" dynamodbav:"image_url"`
}

func fetchImages(tableName string) (images []Image, err error) {
	session, err := session.NewSession()
	if err != nil {
		return nil, err
	}

	svc := dynamodb.New(session)
	result, err := svc.Scan(&dynamodb.ScanInput{
		TableName: aws.String(tableName),
	})

	if err != nil {
		return nil, err
	}

	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &images)
	if err != nil {
		return nil, err
	}

	return images, nil
}
