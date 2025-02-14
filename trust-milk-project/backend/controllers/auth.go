package controllers

import (
    "context"
    "net/http"
    "trust-milk-backend/config"
    "trust-milk-backend/models"
    "trust-milk-backend/utils"

    "github.com/gin-gonic/gin"
    "github.com/gorilla/sessions"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

var store = sessions.NewCookieStore([]byte("super-secret-key"))

func RegisterUser(c *gin.Context) {
    var user models.User
    if err := c.BindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }

    // Check if username already exists
    var existingUser models.User
    err := config.DB.Collection("users").FindOne(context.Background(), bson.M{"username": user.Username}).Decode(&existingUser)
    if err == nil {
        c.JSON(http.StatusConflict, gin.H{"error": "Username already exists"})
        return
    }

    // Check if mobile number already exists
    err = config.DB.Collection("users").FindOne(context.Background(), bson.M{"mobile": user.Mobile}).Decode(&existingUser)
    if err == nil {
        c.JSON(http.StatusConflict, gin.H{"error": "Mobile number already exists"})
        return
    }

    hashedPassword, err := utils.HashPassword(user.Password)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error hashing password"})
        return
    }
    user.Password = hashedPassword

    user.ID = primitive.NewObjectID()
    _, err = config.DB.Collection("users").InsertOne(context.Background(), user)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create user"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}

func LoginUser(c *gin.Context) {
    var loginData models.User
    if err := c.BindJSON(&loginData); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
        return
    }

    var user models.User
    err := config.DB.Collection("users").FindOne(context.Background(), bson.M{"username": loginData.Username}).Decode(&user)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
        return
    }

    if !utils.CheckPasswordHash(loginData.Password, user.Password) {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
        return
    }

    session, _ := store.Get(c.Request, "session")
    session.Values["username"] = user.Username
    session.Save(c.Request, c.Writer)

    c.JSON(http.StatusOK, gin.H{"message": "Login successful", "username": user.Username})
}

func LogoutUser(c *gin.Context) {
    session, _ := store.Get(c.Request, "session")
    session.Options.MaxAge = -1
    session.Save(c.Request, c.Writer)
    c.JSON(http.StatusOK, gin.H{"message": "Logged out"})
}
