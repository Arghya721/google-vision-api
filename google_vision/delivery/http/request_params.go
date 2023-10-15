package http

import (
	"mime/multipart"

	"github.com/labstack/echo/v4"
)

// ReadFormDataFile is a helper function to get file from request body form-data
func ReadFormDataFile(c echo.Context, fieldName string) (file *multipart.FileHeader, err error) {
	// Get file from request body form-data
	file, err = c.FormFile(fieldName)
	if err != nil {
		return
	}
	return
}
