package main

import (
	"encoding/base64"
	"os"
	"testing"
)

func TestHandleLambdaEvent(t *testing.T) {
	img, err := os.ReadFile("template0.png")
	font, err := os.ReadFile("jet.ttf")
	if err != nil {
		t.Fatal("File not found")
	}
	cert := CertificateRequest{
		TemplateImage: base64.StdEncoding.EncodeToString(img),
		Font:          base64.StdEncoding.EncodeToString(font),
		Fields:        []Field{{Name: "Name", XCord: 5 * 309, YCord: 5 * 220, Width: 5 * (617 - 309), Height: 5 * (239 - 220), Value: "BootleBAGagdasgdsagdsgsfasdsadihgsaodihgaosidhgodsaihgosihgosaidhgoisadddddddudsogadsghguf"}},
	}

	a, err := HandleLambdaEvent(cert)
	if err != nil {
		t.Fatal("Handle event error")
	}
	base64.StdEncoding.DecodeString(a.Image)

}
