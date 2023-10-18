package http

import (
	"mime/multipart"
	"strconv"

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

// ReadFormDataValue is a helper function to get value from request body form-data
func ReadFormDataString(c echo.Context, fieldName string) (value string, err error) {
	// Get file from request body form-data
	value = c.FormValue(fieldName)

	return
}

// ReadFormDataInt is a helper function to get value from request body form-data
func ReadFormDataInt(c echo.Context, fieldName string) (value int, err error) {

	intValue := c.FormValue(fieldName)

	if intValue == "" {
		return 0, nil
	}

	// convert string to int
	value, err = strconv.Atoi(intValue)

	return
}
