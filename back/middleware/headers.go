package middleware

import "github.com/gin-gonic/gin"

func Headers(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "PUT, POST, GET, DELETE")
	c.Header("Access-Control-Allow-Headers", "content-type")

	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(204)
		return
	}

	c.Next()
}
