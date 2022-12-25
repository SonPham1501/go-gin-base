package routers

import (
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.Default()
	router.Use(gin.Recovery())
	// router.Use(middlewares.AuthMiddleware())

	v1 := router.Group("v1")
	{
		userRouter(v1)
	}

	return router
}
