package utils

import (
	"encoding/base64"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

func fileToBase64(base64Array *[]string) filepath.WalkFunc {
	return func(fsElementName string, info fs.FileInfo, err error) error {
		if !info.IsDir() {
			fileBytes, err := os.ReadFile(fsElementName)
			if err != nil {
				return err
			}
			base64File := base64.StdEncoding.EncodeToString(fileBytes)
			base64File = fmt.Sprintf("data:image/%s;base64,%s", GetExtension(fsElementName), base64File)
			*base64Array = append(*base64Array, base64File)
		}
		return nil
	}
}

func JPEGtoBase64(path string) ([]string, error) {
	var base64List []string
	filepath.Walk(path, fileToBase64(&base64List))
	return base64List, nil
}
