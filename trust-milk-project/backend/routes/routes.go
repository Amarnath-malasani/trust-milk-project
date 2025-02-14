package routes

import (
    "net/http"
    "trust-milk-backend/controllers"
    "trust-milk-backend/middleware"

    "github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
    router.POST("/register", controllers.RegisterUser)
    router.POST("/login", controllers.LoginUser)
    router.GET("/logout", controllers.LogoutUser)

    protected := router.Group("/api")
    protected.Use(middleware.AuthMiddleware())
    {
        protected.GET("/dashboard", func(c *gin.Context) {
            username, _ := c.Get("username")
            c.JSON(http.StatusOK, gin.H{"message": "Welcome to the dashboard!", "user": username})
        })
    }
}
