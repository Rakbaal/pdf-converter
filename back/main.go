package main

import (
	"pdf-converter/handlers"
	"pdf-converter/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.Use(middleware.Headers)

	router.POST("/files", handlers.HandleFiles)

	router.Run(":8910")
}
