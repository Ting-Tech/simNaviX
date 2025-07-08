package main

import "gin/router"

func main() {
	router := router.SetupRouter()

	// Run the server on port 8080
	router.Run(":8080")
}
