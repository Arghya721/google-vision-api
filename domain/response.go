package domain

type ExtractTextResponse struct {
	Text   string `json:"text"`
	Locale string `json:"locale"`
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
