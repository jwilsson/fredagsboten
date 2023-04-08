package main

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"

	utils "github.com/jwilsson/go-bot-utils"
)

func getImages(ctx context.Context, tableName string) ([]Image, error) {
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return nil, err
	}

	images, err := utils.GetDynamodbData[Image](ctx, cfg, tableName)

	return images, err
}
