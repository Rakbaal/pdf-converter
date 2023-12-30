package utils

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
)

func CreateTempFile(tempDirectory string, part *multipart.Part) (*os.File, error) {
	tempFileName := fmt.Sprintf("%s%s", TempName(), filepath.Ext(part.FileName()))
	tempFile, err := os.Create(tempDirectory + "/" + tempFileName)
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(tempFile, part)
	if err != nil {
		return nil, err
	}

	return tempFile, nil
}
