package routes

import (
	"backend-iot-gps-tracker/controllers"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var PORT = os.Getenv("PORT")

func StartServer() *gin.Engine {
	router := gin.Default()

	// config CORS
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}
	config.AllowMethods = []string{"GET", "POST"}
	router.Use(cors.New(config))

	router.GET("/", controllers.GetChat)
	router.POST("/chat", controllers.CreateChat)
	router.GET("/chat-url/:lat/:lon", controllers.CreateChatByUrl)
	router.GET("/list", controllers.GetChat)
	router.POST("/delete", controllers.DeleteChat)

	// port
	if PORT == "" {
		PORT = "8080"
	}

	router.Run(":" + PORT)

	return router
}
