package handlers

import (
	"fmt"
	"net/http"
	"os"
	"pdf-converter/utils"

	"github.com/gin-gonic/gin"
)

func HandleFiles(c *gin.Context) {
	fmt.Println("Parsing form...")
	reader, err := c.Request.MultipartReader()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Erreur de création du lecteur de formulaire"})
		panic(err)
	}
	fmt.Println("OK")

	fmt.Println("Creating temp files from form...")
	tempDirectory, err := utils.FormToFiles(reader)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Erreur de création des fichiers temp"})
		panic(err)
	}
	defer os.RemoveAll(tempDirectory)
	fmt.Println("OK")

	fmt.Println("Parsing files to JPEG...")
	finalDirectory, err := utils.FilesToFinal(tempDirectory)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Erreur de conversion des fichiers temp"})
		panic(err)
	}
	fmt.Println("OK")

	fmt.Println("Parsing JPEG to base64...")
	base64List, err := utils.JPEGtoBase64(finalDirectory + "/")
	fmt.Println("OK")

	c.JSON(http.StatusOK, gin.H{"message": "Conversion des fichiers réussie", "data": len(base64List)})
}
