package toJPEG

import (
	"image"
	"image/jpeg"
	"os"
	"strings"
)

func removeExtension(fileName string) string {
	dotIndex := strings.LastIndex(fileName, ".")
	if dotIndex != -1 {
		newName := fileName[:dotIndex]
		return newName
	}
	return fileName
}

func CreateJPEGFile(fileName string, image image.Image, dest string) error {
	newName := removeExtension(fileName)
	jpegFile, err := os.Create(dest + "/" + newName + ".jpeg")
	if err != nil {
		return err
	}

	err = jpeg.Encode(jpegFile, image, &jpeg.Options{Quality: jpeg.DefaultQuality})
	if err != nil {
		return err
	}

	err = jpegFile.Close()
	if err != nil {
		return err
	}

	return nil
}
