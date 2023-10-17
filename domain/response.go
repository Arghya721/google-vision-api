package domain

import "bytes"

type ExtractTextResponse struct {
	Text   string `json:"text"`
	Locale string `json:"locale"`
}

type ExtractTextWithBoundaryResponse struct {
	Text     string     `json:"text"`
	Locale   string     `json:"locale"`
	Vertices []Vertices `json:"vertices"`
}

type DetectLabelsResponse struct {
	Labels []string `json:"labels"`
}

type DetectObjectResponse struct {
	Objects []string `json:"objects"`
}

type DetectLandmarkResponse struct {
	Landmarks []string `json:"landmarks"`
}

type DrawBoundaryResponse struct {
	Image *bytes.Buffer `json:"image"`
}
