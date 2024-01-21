package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetSunTimes(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Sunrise and Sunset times here.",
	})
}
