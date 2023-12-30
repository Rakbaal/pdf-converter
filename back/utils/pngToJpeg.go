package utils

import (
	"bytes"
	"image/jpeg"
	"image/png"
)

func ToJpeg(imageBytes []byte) ([]byte, error) {

	img, err := png.Decode(bytes.NewReader(imageBytes))
	if err != nil {
		return nil, err
	}

	buffer := new(bytes.Buffer)
	if err := jpeg.Encode(buffer, img, nil); err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}
