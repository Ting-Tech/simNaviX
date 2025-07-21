package router

import (
	"gin/controller"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		AllowCredentials: true,
	}))

	// router.Static("/uploads", "./uploads")
	router.GET("/uploads/*filepath", func(c *gin.Context) {
		filepath := c.Param("filepath")
		c.Header("Access-Control-Allow-Origin", "http://localhost:5173")
		c.File("./uploads" + filepath)
	})

	router.POST("/path", controller.CalculatePath)

	router.GET("/robot", controller.GetRobot)
	router.POST("/robot/create", controller.CreateRobot)
	router.POST("/robot/remove", controller.RemoveRobot)
	router.POST("/robot/edit", controller.EditRobot)

	router.POST("/config/space", controller.ConfigSpace)

	return router
}
