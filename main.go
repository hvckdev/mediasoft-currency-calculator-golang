package main

import (
	"awesomeProject3/route"
	"awesomeProject3/routine"
	"github.com/gin-gonic/gin"
)

func main() {
	go routine.UpdateCurrencies()

	router := gin.Default()
	api := router.Group("/api")
	{
		api.POST("/create", route.Create)
		api.POST("/convert", route.Convert)
	}

	router.Run("localhost:2222")
}
