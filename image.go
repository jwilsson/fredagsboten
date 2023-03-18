package main

import (
	"github.com/aws/aws-sdk-go/aws/session"

	utils "github.com/jwilsson/go-bot-utils"
)

func getImages(tableName string) (images []Image, err error) {
	s, err := session.NewSession()
	if err != nil {
		return images, err
	}

	err = utils.GetDynamodbData(s, tableName, &images)

	return images, err
}
