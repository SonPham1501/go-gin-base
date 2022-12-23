package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine  {
	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	v1 := router.Group("v1")
	{
		v1.GET("/get", func(c *gin.Context)  {
			c.JSON(http.StatusOK, gin.H{
				"message": "OK",
			})
		})
	}

	return router
}