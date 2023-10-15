package domain

import (
	"bytes"

	vision "cloud.google.com/go/vision/apiv1"
)

type GoogleVisionUsecase interface {
	ExtractText(googleClient *vision.ImageAnnotatorClient, imageBytes *bytes.Buffer) (ExtractTextResponse, error)
	DetectLabels(googleClient *vision.ImageAnnotatorClient, imageBytes *bytes.Buffer) (DetectLabelsResponse, error)
	DetectObject(googleClient *vision.ImageAnnotatorClient, imageBytes *bytes.Buffer) (DetectObjectResponse, error)
	DetectLandmark(googleClient *vision.ImageAnnotatorClient, imageBytes *bytes.Buffer) (DetectLandmarkResponse, error)
}
