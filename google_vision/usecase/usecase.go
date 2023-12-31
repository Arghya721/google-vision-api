package usecase

import (
	"bytes"
	"context"
	"google-vision/domain"

	vision "cloud.google.com/go/vision/apiv1"
)

type GoogleVisionUsecase struct{}

func NewGoogleVisionUsecase() domain.GoogleVisionUsecase {
	return &GoogleVisionUsecase{}
}

// ExtractText from image
func (u *GoogleVisionUsecase) ExtractText(googleClient *vision.ImageAnnotatorClient, imageBytes *bytes.Buffer) (response domain.ExtractTextResponse, err error) {

	annotations, err := ExtractTextFromImage(googleClient, imageBytes)
	if err != nil {
		return domain.ExtractTextResponse{}, err
	}

	if len(annotations) == 0 {
		return domain.ExtractTextResponse{}, nil
	}

	// get text and locale
	response.Locale = annotations[0].Locale
	response.Text = annotations[0].Description

	return
}

// ExtractTextWithBoundary from image
func (u *GoogleVisionUsecase) ExtractTextWithBoundary(googleClient *vision.ImageAnnotatorClient, imageBytes *bytes.Buffer) (response domain.ExtractTextWithBoundaryResponse, err error) {

	// Create a Vision image object from the base64-encoded image data
	annotations, err := ExtractTextFromImage(googleClient, imageBytes)
	if err != nil {
		return domain.ExtractTextWithBoundaryResponse{}, err
	}

	if len(annotations) == 0 {
		return domain.ExtractTextWithBoundaryResponse{}, nil
	}

	response.Locale = annotations[0].Locale
	response.Text = annotations[0].Description

	for _, annotation := range annotations[0].BoundingPoly.Vertices {
		response.Vertices = append(response.Vertices, domain.Vertices{
			X: annotation.X,
			Y: annotation.Y,
		})
	}

	return
}

// DetectLabels from image
func (u *GoogleVisionUsecase) DetectLabels(googleClient *vision.ImageAnnotatorClient, imageBytes *bytes.Buffer) (response domain.DetectLabelsResponse, err error) {

	annotations, err := DetectLabelsFromImage(googleClient, imageBytes)
	if err != nil {
		return domain.DetectLabelsResponse{}, err
	}

	if len(annotations) == 0 {
		return domain.DetectLabelsResponse{}, nil
	}

	for _, annotation := range annotations {
		response.Labels = append(response.Labels, annotation.Description)
	}

	return
}

// DetectLabelsWithConfidence from image
func (u *GoogleVisionUsecase) DetectLabelsWithConfidence(googleClient *vision.ImageAnnotatorClient, imageBytes *bytes.Buffer) (response domain.DetectLabelsWithConfidenceResponse, err error) {

	annotations, err := DetectLabelsFromImage(googleClient, imageBytes)
	if err != nil {
		return domain.DetectLabelsWithConfidenceResponse{}, err
	}

	if len(annotations) == 0 {
		return domain.DetectLabelsWithConfidenceResponse{}, nil
	}

	for _, annotation := range annotations {
		response.Labels = append(response.Labels, domain.Label{
			Name:       annotation.Description,
			Confidence: annotation.Score * 100,
		})
	}

	return
}

// DetectObject from image
func (u *GoogleVisionUsecase) DetectObject(googleClient *vision.ImageAnnotatorClient, imageBytes *bytes.Buffer) (response domain.DetectObjectResponse, err error) {

	annotations, err := DetectObjectFromImage(googleClient, imageBytes)

	if len(annotations) == 0 {
		return domain.DetectObjectResponse{}, nil
	}

	for _, annotation := range annotations {
		response.Objects = append(response.Objects, annotation.Name)
	}

	return
}

// DetectLandmark from image
func (u *GoogleVisionUsecase) DetectLandmark(googleClient *vision.ImageAnnotatorClient, imageBytes *bytes.Buffer) (response domain.DetectLandmarkResponse, err error) {

	// Create a Vision image object from the base64-encoded image data
	imageObj, err := vision.NewImageFromReader(imageBytes)
	if err != nil {
		return domain.DetectLandmarkResponse{}, err
	}

	// Annotate the image
	annotations, err := googleClient.DetectLandmarks(context.Background(), imageObj, nil, 10)
	if err != nil {
		return domain.DetectLandmarkResponse{}, err
	}

	if len(annotations) == 0 {
		return domain.DetectLandmarkResponse{}, nil
	}

	for _, annotation := range annotations {
		response.Landmarks = append(response.Landmarks, annotation.Description)
	}

	return
}

// DrawBoundary from image
func (u *GoogleVisionUsecase) DrawBoundary(googleClient *vision.ImageAnnotatorClient, imageBytes *bytes.Buffer, borderColor string, borderSize int) (response domain.DrawBoundaryResponse, err error) {

	// make a deep copy of imageBytes
	imageBytesCopy := bytes.NewBuffer(imageBytes.Bytes())

	annotations, err := ExtractTextFromImage(googleClient, imageBytes)
	if err != nil {
		return domain.DrawBoundaryResponse{}, err
	}

	if len(annotations) == 0 {
		return domain.DrawBoundaryResponse{}, nil
	}

	// Draw boundary
	response, err = DrawBoundary(imageBytesCopy, annotations, borderColor, borderSize)
	if err != nil {
		return domain.DrawBoundaryResponse{}, err
	}

	return
}
