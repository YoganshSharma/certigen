package main

import (
	"bytes"
	"encoding/base64"
	"image"
	"strings"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/fogleman/gg"
	"github.com/golang/freetype/truetype"
)

type Field struct {
	Name   string  `json:"name"`
	XCord  float64 `json:"xcord"`
	YCord  float64 `json:"ycord"`
	Width  float64 `json:"width"`
	Height float64 `json:"height"`
	Value  string  `json:"value"`
}

type CertificateRequest struct {
	TemplateImage string  `json:"template_image"`
	Fields        []Field `json:"fields"`
	Font          string  `json:"font"`
}

type CertResponse struct {
	Image string `json:"image:"`
}
type ErrorResponse struct {
	Message string `json:"message"`
}

func HandleLambdaEvent(req CertificateRequest) (any, error) {
	// Remove the "data:image/png;base64," prefix if it exists
	encodedImage := strings.TrimPrefix(req.TemplateImage, "data:image/png;base64,")
	encodedFont := strings.TrimPrefix(req.Font, "data:image/png;base64,")

	// Decode the base64 encoded image and font
	decodedImage, err := base64.StdEncoding.DecodeString(encodedImage)
	if err != nil {
		return ErrorResponse{Message: "Not a valid base64 for image"}, err
	}
	decodedFont, err := base64.StdEncoding.DecodeString(encodedFont)
	if err != nil {
		return ErrorResponse{Message: "Not a valid base64 for font"}, err
	}

	// Decode font and image
	img, _, err := image.Decode(bytes.NewReader(decodedImage))
	if err != nil {
		return ErrorResponse{Message: "Not a valid Image format"}, err
	}
	dc := gg.NewContextForImage(img)

	ttfont, err := truetype.Parse(decodedFont)
	if err != nil {
		return ErrorResponse{Message: "Not a valid font"}, err
	}
	dc.SetRGB(0, 0, 0)
	for _, field := range req.Fields {

		pts := field.Height * 96 / 72

		font := truetype.NewFace(ttfont, &truetype.Options{Size: pts})
		dc.SetFontFace(font)
		w, _ := dc.MeasureString(field.Value)
		if w > field.Width {
			height := field.Height * field.Width / w
			dc.LoadFontFace(req.Font, height)
		}

		dc.DrawStringAnchored(field.Value, field.XCord, field.YCord, 0.5, 0.5)
		dc.Fill()

	}
	var buf bytes.Buffer
	dc.EncodePNG(&buf)
	outputImage := base64.StdEncoding.EncodeToString(buf.Bytes())
	return CertResponse{Image: outputImage}, nil
}

func main() {
	lambda.Start(HandleLambdaEvent)
}
