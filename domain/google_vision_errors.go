package domain

// ErrorDetails is a struct used for storing response of error details
type ErrorDetails struct {
	Code        string `json:"errorCode"`
	Description string `json:"errorDescription"`
}

// ImageNotFoundError is a struct used for storing response of error details
var ImageNotFoundError = ErrorDetails{
	Code:        "imageNotFound",
	Description: "No image provided",
}

// ImageNotValidError is a struct used for storing response of error details
var ImageNotValidError = ErrorDetails{
	Code:        "imageNotValid",
	Description: "Image is not valid",
}

// NoTextFoundError is a struct used for storing response of error details
var NoTextFoundError = ErrorDetails{
	Code:        "noTextFound",
	Description: "No text found",
}

// ColorNotValidError is a struct used for storing response of error details
var ColorNotValidError = ErrorDetails{
	Code:        "colorNotValid",
	Description: "Color should be a valid hex color",
}

// BorderSizeNotValidError is a struct used for storing response of error details
var BorderSizeNotValidError = ErrorDetails{
	Code:        "borderSizeNotValid",
	Description: "Border size should be a valid natural number",
}
