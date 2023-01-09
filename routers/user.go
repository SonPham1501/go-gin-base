package routers

import (
	"com.son.server.goginbase/controllers"
	"github.com/gin-gonic/gin"
)

func userRouter(g *gin.RouterGroup, userController *controllers.UserController)  {
	userGroup := g.Group("user")
	{
		userGroup.GET("/:id", userController.Retrieve)
	}
}