package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type itemPOSTRequest struct {
	Id          int    `json:"id" validate:"required"`
	Name        string `json:"name" validate:"required"`
	Price       int    `json:"price" validate:"required"`
	Type        string `json:"type" validate:"required"`
	Size        string `json:"size" validate:"required,regexp=^[0-9]+\\s?(ml|l|g|kg)$"`
	Quantity    int    `json:"quantity" validate:"required"`
	Description string `json:"description" validate:"required,min=10,max=500"`
	Image       string `json:"image" validate:"required"`
}

func AddItemHandler(c *gin.Context) {

	var itemReq itemPOSTRequest

	// Bind JSON input
	if err := c.ShouldBindJSON(&itemReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := validate.Struct(itemReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
}
