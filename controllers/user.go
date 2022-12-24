package controllers

import "github.com/gin-gonic/gin"

type IUserController interface {
	Retrieve()
}

type UserController struct {}

func (this *UserController) Retrieve(c *gin.Context)  {
	
}