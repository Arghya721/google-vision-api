package clients

import (
	"context"
	"google-vision/config"

	vision "cloud.google.com/go/vision/apiv1"
	"google.golang.org/api/option"
)

func NewImageClient(ctx context.Context) (*vision.ImageAnnotatorClient, error) {
	credentialFile := config.GoogleCredentialFileName
	var client *vision.ImageAnnotatorClient
	var err error

	if config.ApplicationEnv == config.ProductionEnv {
		client, err = vision.NewImageAnnotatorClient(ctx)
	} else {
		client, err = vision.NewImageAnnotatorClient(ctx, option.WithCredentialsFile(credentialFile))
	}

	if err != nil {
		return nil, err
	}

	return client, nil
}
