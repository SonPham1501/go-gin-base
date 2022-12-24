package server

import (
	"com.son.server.goginbase/controllers"
	"com.son.server.goginbase/middlewares"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(middlewares.AuthMiddleware())

	v1 := router.Group("v1")
	{
		userGroup := v1.Group("user")
		{
			user := controllers.UserController{}
			userGroup.GET("/:id", user.Retrieve)
		}
	}

	return router
}
