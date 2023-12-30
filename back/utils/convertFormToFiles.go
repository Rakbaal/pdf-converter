package utils

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
)

func FormToFiles(reader *multipart.Reader) (string, error) {
	tempDirectory := fmt.Sprintf("./temp/%s", TempName())
	os.MkdirAll(tempDirectory, 0777)

	for {
		part, err := reader.NextPart()
		if err == io.EOF {
			break
		}
		if err != nil {
			return "", err
		}

		file, err := CreateTempFile(tempDirectory, part)
		if err != nil {
			return "", err
		}
		file.Close()
	}

	return tempDirectory, nil
}
