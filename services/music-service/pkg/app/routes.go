package app

import "github.com/gin-gonic/gin"

// Routes sets up the routes for the application.
func (server *Server) Routes() *gin.Engine {
	router := server.router

	router.GET("/health-check", server.HealthCheck())

	// v1 := router.Group("/v1/api")
	// {
	// 	music := v1.Group("/music")
	// 	{
	// 		music.GET("/", server.GetAll())
	// 	}
	// }

	return router
}
