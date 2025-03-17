package controllers

import (
	"context"
	"log"
	"net/http"
	"tea_ume/ent"
	"tea_ume/ent/user"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// LoginHandler authenticates the user
func LoginHandler(c *gin.Context) {
	var loginRequest struct {
		Username string `json:"username" validate:"required"`
		Password string `json:"password" validate:"required,min=8"`
	}

	// Bind JSON input
	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Get database client
	client := c.MustGet("ent").(*ent.Client)

	// Fetch user by username
	existingUser, err := client.User.
		Query().
		Where(user.UsernameEQ(loginRequest.Username)).
		Only(context.Background())

	if err != nil {
		log.Printf("Error finding user: %v", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	// Compare hashed password
	if err := bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(loginRequest.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	// Generate JWT Token
	token, err := GenerateJWT(existingUser.ID, existingUser.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"user_id": existingUser.ID,
		"role":    existingUser.Role,
		"token":   token,
	})
}
