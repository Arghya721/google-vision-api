package usecase

import (
	"bytes"
	"context"

	vision "cloud.google.com/go/vision/apiv1"
	"cloud.google.com/go/vision/v2/apiv1/visionpb"
)

func ExtractTextFromImage(googleClient *vision.ImageAnnotatorClient, imageBytes *bytes.Buffer) (annotations []*visionpb.EntityAnnotation, err error) {

	// Create a Vision image object from the base64-encoded image data
	imageObj, err := vision.NewImageFromReader(imageBytes)
	if err != nil {
		return annotations, err
	}

	// Annotate the image
	annotations, err = googleClient.DetectTexts(context.Background(), imageObj, nil, 10)
	if err != nil {
		return annotations, err
	}

	return
}
