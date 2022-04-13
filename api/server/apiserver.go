package server

import (
	"strconv"

	"jjrepos/gonang/api/controllers"
	"jjrepos/gonang/api/database"

	"github.com/gin-gonic/gin"
)

func StartApi(port int) {
	database.Connect()
	router := SetupRoutes()
	router.Run("localhost:" + strconv.Itoa(port))
}

func SetupRoutes() *gin.Engine {
	router := gin.Default()
	router.GET("/books", controllers.FindBooks)
	router.POST("/books", controllers.CreateBook)
	router.GET("/books/:id", controllers.FindBook)
	router.PATCH("/books/:id", controllers.UpdateBook)
	return router
}

// read configuration file and configure application
func Configure(configPath string) {
}
