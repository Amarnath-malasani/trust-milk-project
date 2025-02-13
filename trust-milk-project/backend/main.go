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
	routes.SetupRoutes(router)

	fmt.Println("Server running on port 8080...")
	router.Run(":8080")
}
