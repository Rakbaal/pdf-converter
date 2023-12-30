package utils

import (
	"io/fs"
	"net/http"
	"os"
	"path/filepath"
	"pdf-converter/utils/toJPEG"
)

func GetFileExt(file *os.File) (string, error) {
	buffer := make([]byte, 512)
	_, err := file.Read(buffer)
	if err != nil {
		return "", err
	}

	mimeType := http.DetectContentType(buffer)

	return mimeType, nil
}

func fileToJpeg(dest string) filepath.WalkFunc {
	return func(path string, info fs.FileInfo, err error) error {
		err = os.MkdirAll(dest, 0777)
		if err != nil {
			return err
		}

		if !info.IsDir() {
			file, err := os.Open(path)
			if err != nil {
				return err
			}

			mimeType, err := GetFileExt(file)
			if err != nil {
				return err
			}

			switch mimeType {
			case "application/pdf":
				err = toJPEG.PDF(path, dest)
				if err != nil {
					return err
				}
				return nil
			default:
				err = toJPEG.IMG(path, dest)
				if err != nil {
					return err
				}
				return nil
			}
		}
		return nil
	}
}

func FilesToJPEG(tempDirectory string) (string, error) {
	jpegDirectory := tempDirectory + "/jpeg"
	err := filepath.Walk(tempDirectory, fileToJpeg(jpegDirectory))
	return jpegDirectory, err
}
