package domain

import (
	"bytes"

	vision "cloud.google.com/go/vision/apiv1"
)

type GoogleVisionUsecase interface {
	ExtractText(googleClient *vision.ImageAnnotatorClient, imageBytes *bytes.Buffer) (ExtractTextResponse, error)
	ExtractTextWithBoundary(googleClient *vision.ImageAnnotatorClient, imageBytes *bytes.Buffer) (ExtractTextWithBoundaryResponse, error)
	DetectLabels(googleClient *vision.ImageAnnotatorClient, imageBytes *bytes.Buffer) (DetectLabelsResponse, error)
	DetectLabelsWithConfidence(googleClient *vision.ImageAnnotatorClient, imageBytes *bytes.Buffer) (DetectLabelsWithConfidenceResponse, error)
	DetectObject(googleClient *vision.ImageAnnotatorClient, imageBytes *bytes.Buffer) (DetectObjectResponse, error)
	DetectLandmark(googleClient *vision.ImageAnnotatorClient, imageBytes *bytes.Buffer) (DetectLandmarkResponse, error)
	DrawBoundary(googleClient *vision.ImageAnnotatorClient, imageBytes *bytes.Buffer, color string, borderSize int) (DrawBoundaryResponse, error)
}
