package handlers

import (
	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {

	resp := "Covid19 Collection v1"

	c.JSON(200, gin.H{
		"code":    200,
		"status":  "OK",
		"message": resp,
	})
}
