package utils

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func Trigger(c *gin.Context) {
	directoryName := fmt.Sprintf("./%s", strings.TrimSpace(time.Now().String()))
	os.MkdirAll(directoryName, 700)
	// defer os.RemoveAll(directoryName)

}
