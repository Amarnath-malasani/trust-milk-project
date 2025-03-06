package controllers

import (
    "context"
    "net/http"
    "trust-milk-backend/config"
    "trust-milk-backend/models"

    "github.com/gin-gonic/gin"
    "go.mongodb.org/mongo-driver/bson"
)

func GetProfile(c *gin.Context) {
    username, exists := c.Get("username")
    if !exists {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
        return
    }

    var user models.User
    err := config.DB.Collection("users").FindOne(context.Background(), bson.M{"username": username}).Decode(&user)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"success": true, "user": user})
}