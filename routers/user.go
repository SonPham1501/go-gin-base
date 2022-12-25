package routers

import (
	"com.son.server.goginbase/controllers"
	"github.com/gin-gonic/gin"
)

func userRouter(g *gin.RouterGroup)  {
	userGroup := g.Group("user")
	{
		userController := controllers.UserController {}
		userGroup.GET("/:id", userController.Retrieve)
	}
}