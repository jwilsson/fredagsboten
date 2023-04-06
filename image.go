package main

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"

	utils "github.com/jwilsson/go-bot-utils"
)

func getImages(ctx context.Context, tableName string) (images []Image, err error) {
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return images, err
	}

	err = utils.GetDynamodbData(ctx, cfg, tableName, &images)

	return images, err
}
