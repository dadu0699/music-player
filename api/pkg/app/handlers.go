package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Status is a struct that holds the status of the API.
func Status(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"data": "API running smoothly",
	})
}
