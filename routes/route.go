package routes

import (
	"backend-iot-gps-tracker/controllers"
	"os"

	"github.com/gin-gonic/gin"
)

var PORT = os.Getenv("PORT")

func StartServer() *gin.Engine {
	router := gin.Default()

	router.GET("/", controllers.GetChat)
	router.POST("/chat", controllers.CreateChat)

	// port
	if PORT == "" {
		PORT = "8080"
	}

	router.Run(":" + PORT)

	return router
}
