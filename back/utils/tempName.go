package utils

import (
	"strings"
	"time"
)

func TempName() string {
	return strings.Replace(time.Now().String(), " ", "", -1)
}
