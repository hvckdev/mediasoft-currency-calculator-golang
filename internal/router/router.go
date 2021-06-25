package router

import (
	"awesomeProject3/internal/currency"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	api := router.Group("/api")
	{
		api.POST("/create", currency.Create)
		api.POST("/convert", currency.Convert)
	}

	return router
}
