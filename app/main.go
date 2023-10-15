package main

import (
	googleVisionHandler "google-vision/google_vision/delivery/http"

	_googleVisionUsecase "google-vision/google_vision/usecase"

	"google-vision/config"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var (
	e *echo.Echo
)

func init() {
	// Initialize config
	config.InitializeConfig()

	e = echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
}

func main() {

	// Initialize usecase
	googleVisionUsecase := _googleVisionUsecase.NewGoogleVisionUsecase()

	// Routes
	googleVisionHandler.RegisterGoogleVisionHandler(e, googleVisionUsecase)

	// Start the server
	e.Logger.Fatal(e.Start(":1323"))
}
