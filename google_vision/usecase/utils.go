package usecase

import (
	"bytes"
	"fmt"
	"google-vision/domain"
	"image"
	"image/color"
	"image/draw"
	"image/gif"
	"image/jpeg"
	"image/png"

	"cloud.google.com/go/vision/v2/apiv1/visionpb"
	"golang.org/x/image/bmp"
)

// DecodeImage from bytes
func DecodeImage(imageBytes *bytes.Buffer) (img image.Image, err error) {

	// make a deep copy of imageBytes
	imageBytesCopy := bytes.NewBuffer(imageBytes.Bytes())

	// Decode the image to detect its format.
	_, format, err := image.DecodeConfig(imageBytes)
	if err != nil {
		return nil, err
	}

	// Based on the detected format, encode the image.
	switch format {
	case "png":
		img, _, err = image.Decode(imageBytesCopy)
		if err != nil {
			return nil, err
		}
	case "jpeg":
		img, err = jpeg.Decode(imageBytesCopy)
		if err != nil {
			return nil, err
		}
	case "gif":
		img, err = gif.Decode(imageBytesCopy)
		if err != nil {
			return nil, err
		}
	case "jpg":
		img, err = jpeg.Decode(imageBytesCopy)
		if err != nil {
			return nil, err
		}
	case "bmp":
		img, err = bmp.Decode(imageBytesCopy)
		if err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("unsupported format: %s", format)
	}

	return
}

// HexToRGBA converts a hex string to a color.RGBA type.
func HexToRGBA(hex string) (color.RGBA, error) {

	// Convert the hex string to a color.RGBA type.
	var clr color.RGBA
	clr.A = 255 // alpha channel
	_, err := fmt.Sscanf(hex, "#%02x%02x%02x", &clr.R, &clr.G, &clr.B)
	if err != nil {
		return color.RGBA{}, err
	}

	return clr, nil
}

// DrawBoundary from image
func DrawBoundary(imageBytes *bytes.Buffer, annotations []*visionpb.EntityAnnotation, borderColor string, borderSize int) (response domain.DrawBoundaryResponse, err error) {

	// take out the vertices
	var vertices []domain.Vertices
	for _, annotation := range annotations[0].BoundingPoly.Vertices {
		vertices = append(vertices, domain.Vertices{
			X: annotation.X,
			Y: annotation.Y,
		})
	}

	img, err := DecodeImage(imageBytes)
	if err != nil {
		return domain.DrawBoundaryResponse{}, err
	}

	// Create a new RGBA image to draw on.
	bounds := img.Bounds()
	rgba := image.NewRGBA(bounds)
	draw.Draw(rgba, bounds, img, image.Point{0, 0}, draw.Src)

	// Convert the hex string to a color.RGBA type.
	color, err := HexToRGBA(borderColor)
	if err != nil {
		return domain.DrawBoundaryResponse{}, err
	}

	var prevX *int32
	var prevY *int32

	// Draw boundaries using the provided vertices.
	for i := 0; i < len(vertices); i++ {

		// mark the vertices
		// rgba.Set(int(vertices[i].X), int(vertices[i].Y), red)

		currVertex := vertices[i]

		if prevX == nil || prevY == nil {
			prevX = &currVertex.X
			prevY = &currVertex.Y
			continue
		}

		drawLine(rgba, *prevX, *prevY, currVertex.X, currVertex.Y, color, borderSize)
		prevX = &currVertex.X
		prevY = &currVertex.Y

	}

	// connect the last vertex to the first vertex
	drawLine(rgba, *prevX, *prevY, vertices[0].X, vertices[0].Y, color, borderSize)

	// Encode as PNG.
	var buff bytes.Buffer
	err = png.Encode(&buff, rgba)
	if err != nil {
		return domain.DrawBoundaryResponse{}, err
	}

	response.Image = &buff

	return
}

func drawLine(img *image.RGBA, x0, y0, x1, y1 int32, clr color.RGBA, borderSize int) {
	dx := abs(x1 - x0)
	dy := abs(y1 - y0)
	sx := -1
	if x0 < x1 {
		sx = 1
	}
	sy := -1
	if y0 < y1 {
		sy = 1
	}
	err := dx - dy

	for {
		for i := int32(0); i < int32(borderSize); i++ {
			img.Set(int(x0+i), int(y0), clr)
			img.Set(int(x0), int(y0+i), clr)
		}
		if x0 == x1 && y0 == y1 {
			break
		}
		e2 := 2 * err
		if e2 > -dy {
			err -= dy
			x0 += int32(sx)
		}
		if e2 < dx {
			err += dx
			y0 += int32(sy)
		}
	}
}

func abs(x int32) int32 {
	if x < 0 {
		return -x
	}
	return x
}
