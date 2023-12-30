package handlers

import (
	"fmt"
	"net/http"
	"pdf-converter/utils"

	"github.com/gin-gonic/gin"
)

func HandleFiles(c *gin.Context) {
	reader, err := c.Request.MultipartReader()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Erreur de création du lecteur de formulaire"})
		panic(err)
	}

	tempDirectory, err := utils.FormToFiles(reader)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Erreur de création des fichiers temp"})
		panic(err)
	}
	// defer os.RemoveAll(tempDirectory)
	fmt.Println(tempDirectory)
	jpegDirectory, err := utils.FilesToJPEG(tempDirectory)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Erreur de conversion des fichiers temp"})
		panic(err)
	}

	base64List, err := utils.JPEGtoBase64(jpegDirectory + "/")
	fmt.Println(len(base64List))
	c.JSON(http.StatusOK, gin.H{"message": "OK"})
}
