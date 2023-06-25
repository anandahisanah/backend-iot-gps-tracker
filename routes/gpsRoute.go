package routes

import (
	"gps-tracker/controllers"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	router.POST("/gps", controllers.CreateGps)

	return router
}
