package utils

import (
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"os"
	"strings"
)

func trimExtension(fileName string) (string, string) {
	dotIndex := strings.LastIndex(fileName, ".")
	var extension string = ""
	if dotIndex != -1 {
		extension = GetExtension(fileName)
		newName := fileName[:dotIndex]
		return newName, extension
	}
	return fileName, extension
}

func CreateJPEGFile(fileName string, decodedImage image.Image, dest string) error {
	trimmedName, extension := trimExtension(fileName)
	jpegFile, err := os.Create(dest + "/" + trimmedName + ".jpeg")
	if err != nil {
		return err
	}

	if extension == "png" {
		// Crée un fond blanc pour le jpeg final
		cleanBackgroundJPEG := image.NewRGBA(decodedImage.Bounds())
		draw.Draw(cleanBackgroundJPEG, cleanBackgroundJPEG.Bounds(), &image.Uniform{color.White}, image.Point{}, draw.Src)
		draw.Draw(cleanBackgroundJPEG, cleanBackgroundJPEG.Bounds(), decodedImage, decodedImage.Bounds().Min, draw.Over)
		err = jpeg.Encode(jpegFile, cleanBackgroundJPEG, &jpeg.Options{Quality: jpeg.DefaultQuality})
		if err != nil {
			return err
		}
	} else {
		err = jpeg.Encode(jpegFile, decodedImage, &jpeg.Options{Quality: jpeg.DefaultQuality})
		if err != nil {
			return err
		}
	}

	err = jpegFile.Close()
	if err != nil {
		return err
	}

	return nil
}
