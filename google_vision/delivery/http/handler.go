package http

import (
	"context"
	"google-vision/config"
	"google-vision/domain"
	"google-vision/google_vision/delivery/clients"
	"net/http"

	"github.com/labstack/echo/v4"
)

type GoogleVisionHandler struct {
	GoogleVisionUsecase domain.GoogleVisionUsecase
}

func RegisterGoogleVisionHandler(e *echo.Echo, googleVisionUsecase domain.GoogleVisionUsecase) {
	handler := &GoogleVisionHandler{
		GoogleVisionUsecase: googleVisionUsecase,
	}

	// Routes
	e.POST("/extract-text", handler.ExtractText)
	e.POST("/extract-text-with-boundary", handler.ExtractTextWithBoundary)
	e.POST("/draw-boundary", handler.DrawBoundary)
	e.POST("/detect-labels", handler.DetectLabels)
	e.POST("/detect-object", handler.DetectObject)
	e.POST("/detect-landmark", handler.DetectLandmark)

}

// ExtractText from image
func (gv *GoogleVisionHandler) ExtractText(c echo.Context) error {
	// Get image from request body form-data
	image, err := ReadFormDataFile(c, "image")
	if err != nil {
		return c.JSON(http.StatusBadRequest, domain.ImageNotValidError)
	}

	// Get image bytes
	imageBytes, err := ImageProcessor(image)
	if err != nil {
		return c.JSON(http.StatusBadRequest, domain.ImageNotValidError)
	}

	if imageBytes == nil {
		return c.JSON(http.StatusBadRequest, domain.ImageNotFoundError)
	}

	client, err := clients.NewImageClient(context.Background())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	// Extract text from image
	response, err := gv.GoogleVisionUsecase.ExtractText(client, imageBytes)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	if response.Text == "" {
		return c.JSON(http.StatusBadRequest, domain.NoTextFoundError)
	}

	defer client.Close()

	// Return JSON response
	return c.JSON(http.StatusOK, response)
}

// ExtractTextWithBoundary from image
func (gv *GoogleVisionHandler) ExtractTextWithBoundary(c echo.Context) error {
	// Get image from request body form-data
	image, err := ReadFormDataFile(c, "image")
	if err != nil {
		return c.JSON(http.StatusBadRequest, domain.ImageNotValidError)
	}

	// Get image bytes
	imageBytes, err := ImageProcessor(image)
	if err != nil {
		return c.JSON(http.StatusBadRequest, domain.ImageNotValidError)
	}

	if imageBytes == nil {
		return c.JSON(http.StatusBadRequest, domain.ImageNotFoundError)
	}

	client, err := clients.NewImageClient(context.Background())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	// Extract text from image
	response, err := gv.GoogleVisionUsecase.ExtractTextWithBoundary(client, imageBytes)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	if response.Text == "" {
		return c.JSON(http.StatusBadRequest, domain.NoTextFoundError)
	}

	defer client.Close()

	// Return JSON response
	return c.JSON(http.StatusOK, response)
}

// DetectLabels from image
func (gv *GoogleVisionHandler) DetectLabels(c echo.Context) error {
	// Get image from request body form-data
	image, err := ReadFormDataFile(c, "image")
	if err != nil {
		return c.JSON(http.StatusBadRequest, domain.ImageNotValidError)
	}

	// Get image bytes
	imageBytes, err := ImageProcessor(image)
	if err != nil {
		return c.JSON(http.StatusBadRequest, domain.ImageNotValidError)
	}

	if imageBytes == nil {
		return c.JSON(http.StatusBadRequest, domain.ImageNotFoundError)
	}

	client, err := clients.NewImageClient(context.Background())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	// Extract text from image
	response, err := gv.GoogleVisionUsecase.DetectLabels(client, imageBytes)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	defer client.Close()

	// Return JSON response
	return c.JSON(http.StatusOK, response)
}

// DetectObject from image
func (gv *GoogleVisionHandler) DetectObject(c echo.Context) error {
	// Get image from request body form-data
	image, err := ReadFormDataFile(c, "image")
	if err != nil {
		return c.JSON(http.StatusBadRequest, domain.ImageNotValidError)
	}

	// Get image bytes
	imageBytes, err := ImageProcessor(image)
	if err != nil {
		return c.JSON(http.StatusBadRequest, domain.ImageNotValidError)
	}

	if imageBytes == nil {
		return c.JSON(http.StatusBadRequest, domain.ImageNotFoundError)
	}

	client, err := clients.NewImageClient(context.Background())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	// Extract text from image
	response, err := gv.GoogleVisionUsecase.DetectObject(client, imageBytes)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	defer client.Close()

	// Return JSON response
	return c.JSON(http.StatusOK, response)
}

// DetectLandmark from image
func (gv *GoogleVisionHandler) DetectLandmark(c echo.Context) error {
	// Get image from request body form-data
	image, err := ReadFormDataFile(c, "image")
	if err != nil {
		return c.JSON(http.StatusBadRequest, domain.ImageNotValidError)
	}

	// Get image bytes
	imageBytes, err := ImageProcessor(image)
	if err != nil {
		return c.JSON(http.StatusBadRequest, domain.ImageNotValidError)
	}

	if imageBytes == nil {
		return c.JSON(http.StatusBadRequest, domain.ImageNotFoundError)
	}

	client, err := clients.NewImageClient(context.Background())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	// Extract text from image
	response, err := gv.GoogleVisionUsecase.DetectLandmark(client, imageBytes)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	defer client.Close()

	// Return JSON response
	return c.JSON(http.StatusOK, response)
}

// DrawBoundary from image
func (gv *GoogleVisionHandler) DrawBoundary(c echo.Context) error {

	// Get image from request body form-data
	image, err := ReadFormDataFile(c, "image")
	if err != nil {
		return c.JSON(http.StatusBadRequest, domain.ImageNotValidError)
	}

	// Get Color from request body form-data
	borderColor, err := ReadFormDataString(c, "borderColor")
	if err != nil {
		return c.JSON(http.StatusBadRequest, domain.ColorNotValidError)
	}

	if borderColor == "" {
		// add default red borderColor
		borderColor = config.DefaultColor
	}

	// check if borderColor is a valid hex borderColor
	if !IsHexColor(borderColor) {
		return c.JSON(http.StatusBadRequest, domain.ColorNotValidError)
	}

	// Get borderSize from request body form-data
	borderSize, err := ReadFormDataInt(c, "borderSize")
	if err != nil {
		return c.JSON(http.StatusBadRequest, domain.BorderSizeNotValidError)
	}

	if borderSize == 0 {
		// add default borderSize
		borderSize = config.DefaultBorderSize
	}

	// Get image bytes
	imageBytes, err := ImageProcessor(image)
	if err != nil {
		return c.JSON(http.StatusBadRequest, domain.ImageNotValidError)
	}

	if imageBytes == nil {
		return c.JSON(http.StatusBadRequest, domain.ImageNotFoundError)
	}

	client, err := clients.NewImageClient(context.Background())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	// Extract text from image
	response, err := gv.GoogleVisionUsecase.DrawBoundary(client, imageBytes, borderColor, borderSize)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	defer client.Close()

	// Return a image response
	return c.Stream(http.StatusOK, "image/png", response.Image)
}
