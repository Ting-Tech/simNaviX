package router

import (
	"gin/controller"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.POST("/path", controller.CalculatePath)

	return router
}
