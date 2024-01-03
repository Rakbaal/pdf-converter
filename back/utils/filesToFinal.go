package utils

import (
	"io/fs"
	"net/http"
	"os"
	"path/filepath"
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

func fileToFinal(dest string) filepath.WalkFunc {
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
				return PDFtoJPEG(path, dest)
			case "image/gif":
				return CopyFile(path, dest)
			case "image/jpeg":
				return CopyFile(path, dest)
			default:
				return IMGtoJPEG(path, dest)
			}
		}
		return nil
	}
}

func FilesToFinal(tempDirectory string) (string, error) {
	finalDirectory := tempDirectory + "/final/"
	err := filepath.Walk(tempDirectory, fileToFinal(finalDirectory))
	return finalDirectory, err
}
