package http

import (
	"bytes"
	"errors"
	"io"
	"mime/multipart"
)

// ImageProcessor is a helper function to get file from request body form-data
func ImageProcessor(image *multipart.FileHeader) (imageBytes *bytes.Buffer, err error) {
	if image != nil {
		file, err := image.Open()
		if err == nil {
			defer file.Close()

			buf := bytes.NewBuffer(nil)
			if _, err := io.Copy(buf, file); err != nil {
				return nil, err
			}
			imageBytes = buf
		}
	}

	if imageBytes == nil {
		return nil, errors.New("no image provided")
	}

	return imageBytes, nil
}
