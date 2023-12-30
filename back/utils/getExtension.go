package utils

import (
	"strings"
)

func GetExtension(filePath string) string {
	dotIndex := strings.LastIndex(filePath, ".")
	extension := filePath[dotIndex+1:]
	return extension
}
