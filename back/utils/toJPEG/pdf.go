package toJPEG

import (
	"fmt"
	"path/filepath"

	gofitz "github.com/karmdip-mi/go-fitz"
)

func PDF(fsElementPath string, dest string) error {
	fileName := filepath.Base(fsElementPath)

	doc, err := gofitz.New(fsElementPath)
	if err != nil {
		return err
	}

	for n := 0; n < doc.NumPage(); n++ {
		pageImage, err := doc.Image(n)
		if err != nil {
			return err
		}

		fileName := fmt.Sprintf("/%s-%05d.jpg", fileName, n)
		err = CreateJPEGFile(fileName, pageImage, dest)
		if err != nil {
			return err
		}
	}

	return nil
}
