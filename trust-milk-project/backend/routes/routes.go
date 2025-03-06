package routes

import (
	"trust-milk-backend/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.POST("/register", controllers.RegisterUser)
	router.POST("/login", controllers.LoginUser)
	router.GET("/logout", controllers.LogoutUser)
	router.POST("/save-location", controllers.SaveLocation)
	router.POST("/reset-password", controllers.ResetPassword)
	router.POST("/delete-account", controllers.DeleteAccount)
	router.GET("/api/profile", controllers.GetProfile)

	protected := router.Group("/api")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.GET("/dashboard", func(c *gin.Context) {
			username, _ := c.Get("username")
			c.JSON(http.StatusOK, gin.H{"message": "Welcome to the dashboard!", "user": username})
		})
	}
}
