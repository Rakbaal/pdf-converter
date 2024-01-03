package utils

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

func CopyFile(src string, dest string) error {
	fileName := filepath.Base(src)
	fmt.Println("triggered")
	srcBytes, err := os.ReadFile(src)
	if err != nil {
		panic(err)
	}

	err = os.WriteFile(dest+fileName, srcBytes, fs.ModePerm)
	if err != nil {
		panic(err)
	}

	return nil
}
