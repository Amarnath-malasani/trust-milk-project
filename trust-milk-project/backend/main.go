package main

import (
	"fmt"
	"trust-milk-backend/config"
	"trust-milk-backend/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDB()

	router := gin.Default()

	// Serve static files
	router.Static("/static", "../frontend")

	// Serve the main HTML file
	router.GET("/", func(c *gin.Context) {
		c.File("../frontend/index.html")
	})

	// Serve the profile HTML file
	router.GET("/profile", func(c *gin.Context) {
		c.File("../frontend/profile/profile.html")
	})

	// Setup routes
	routes.SetupRoutes(router)

	fmt.Println("Server running on port 8080...")
	router.Run(":8080")
}
