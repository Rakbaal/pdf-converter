package utils

import (
	"fmt"
	"image"
	"os"
	"path/filepath"

	"golang.org/x/exp/slices"
)

func decodePicture(file *os.File) (image.Image, error) {
	handledTypes := []string{"jpeg", "png", "gif", "webp", "jpg"}

	decodedPicture, initialFormat, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	if !slices.Contains(handledTypes, initialFormat) {
		return nil, fmt.Errorf("%s is not a valid format", initialFormat)
	}
	return decodedPicture, nil

}

func IMGtoJPEG(src string, dest string) error {
	fileName := filepath.Base(src)
	file, err := os.Open(src)
	if err != nil {
		return err
	}
	defer file.Close()

	decodedImage, err := decodePicture(file)
	if err != nil {
		return err
	}

	err = CreateJPEGFile(fileName, decodedImage, dest)
	if err != nil {
		return err
	}

	return nil
}
