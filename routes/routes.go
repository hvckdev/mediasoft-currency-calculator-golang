package routes

import (
	"awesomeProject3/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	api := router.Group("/api")
	{
		api.POST("/create", controllers.CreateCurrency)
		api.POST("/convert", controllers.ConvertCurrency)
	}

	return router
}
