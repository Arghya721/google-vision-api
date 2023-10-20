package usecase

import (
	"bytes"
	"context"

	vision "cloud.google.com/go/vision/apiv1"
	"cloud.google.com/go/vision/v2/apiv1/visionpb"
)

// ExtractTextFromImage extracts text from image
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

// DetectLabelsFromImage detects labels from image
func DetectLabelsFromImage(googleClient *vision.ImageAnnotatorClient, imageBytes *bytes.Buffer) (annotations []*visionpb.EntityAnnotation, err error) {

	// Create a Vision image object from the base64-encoded image data
	imageObj, err := vision.NewImageFromReader(imageBytes)
	if err != nil {
		return annotations, err
	}

	// Annotate the image
	annotations, err = googleClient.DetectLabels(context.Background(), imageObj, nil, 10)
	if err != nil {
		return annotations, err
	}

	return
}

// DetectObjectFromImage detects object from image
func DetectObjectFromImage(googleClient *vision.ImageAnnotatorClient, imageBytes *bytes.Buffer) (annotations []*visionpb.LocalizedObjectAnnotation, err error) {

	// Create a Vision image object from the base64-encoded image data
	imageObj, err := vision.NewImageFromReader(imageBytes)
	if err != nil {
		return annotations, err
	}

	// Annotate the image
	annotations, err = googleClient.LocalizeObjects(context.Background(), imageObj, nil)
	if err != nil {
		return annotations, err
	}

	return
}
