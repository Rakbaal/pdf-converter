package main

import (
	"pdf-converter/handlers"
	"pdf-converter/middleware"
	"pdf-converter/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.Use(middleware.Headers)

	router.POST("/files", handlers.HandleFiles)

	router.GET("/trigger", utils.Trigger)

	router.Run(":8910")
}
