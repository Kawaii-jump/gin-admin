package router

import (
	"github.com/Kawaii-jump/gin-admin/handler"
	"github.com/Kawaii-jump/gin-admin/middlewares"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func InitRouter() {
	route := gin.Default()

	route.Use(middlewares.Cors())

	route.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1 := route.Group("/api/v1")
	{
		v1.Any("/", handler.HandleRoot)
		v1.POST("/search", handler.HandleSearch)
		v1.POST("/query", handler.HandleQuery)
		v1.POST("/login", handler.HandleLogin)
		// v1.OPTIONS("/login", handler.HandleOption)
	}

	route.Run(":8181")
}
