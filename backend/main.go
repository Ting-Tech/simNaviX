package main

import (
	"gin/config"
	"gin/model"
	"gin/router"
)

func main() {
	config.ConnectDB()
	config.DB.AutoMigrate(&model.Robot{})

	router := router.SetupRouter()

	router.Run(":8080")
}
