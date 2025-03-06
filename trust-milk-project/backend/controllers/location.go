package controllers

import (
	"context"
	"net/http"
	"trust-milk-backend/config"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Location struct {
	ID                   primitive.ObjectID `bson:"_id,omitempty"`
	FullName             string             `bson:"fullName"`
	MobileNumber         string             `bson:"mobileNumber"`
	Pincode              string             `bson:"pincode"`
	AddressDetails       string             `bson:"addressDetails"`
	TownCity             string             `bson:"townCity"`
	State                string             `bson:"state"`
	DefaultAddress       bool               `bson:"defaultAddress"`
	DeliveryInstructions string             `bson:"deliveryInstructions"`
	Latitude             float64            `bson:"latitude"`
	Longitude            float64            `bson:"longitude"`
}

func SaveLocation(c *gin.Context) {
	var location Location
	if err := c.BindJSON(&location); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	collection := config.DB.Collection("locations")
	_, err := collection.InsertOne(context.Background(), location)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save location"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true})
}
