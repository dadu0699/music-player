package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HealthCheck is the status of the application.
func (server *Server) HealthCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")

		response := gin.H{
			"status":   http.StatusOK,
			"response": "API running smoothly",
		}
		c.JSON(http.StatusOK, response)
	}
}
