package controllers

import (
	"context"
	"log"
	"net/http"
	"tea_ume/ent"
	"tea_ume/ent/user"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

var validate = validator.New()

func SignupHandler(c *gin.Context) {
	var userRequest struct {
		FullName string    `json:"full_name" validate:"required"`
		Email    string    `json:"email" validate:"required,email"`
		Username string    `json:"username" validate:"required"`
		Password string    `json:"password" validate:"required,min=8"`
		Role     user.Role `json:"role" validate:"required,oneof=User Seller"`
	}

	// Bind JSON input to User struct
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := validate.Struct(userRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Hash the user's password using bcrypt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userRequest.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error hashing password"})
		return
	}

	// Create the user in the database using Ent's ORM
	client := c.MustGet("ent").(*ent.Client)

	// Ensure the global Client variable is initialized before using it
	if client == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database client is not initialized"})
		return
	}

	newUser, err := client.User.
		Create().
		SetFullName(userRequest.FullName).
		SetEmail(userRequest.Email).
		SetUsername(userRequest.Username).
		SetPassword(string(hashedPassword)).
		SetRole(userRequest.Role).
		SetCreatedAt(time.Now()).
		Save(context.Background())
	if err != nil {
		log.Printf("Error creating user: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error inserting user into database"})
		return
	}

	// Generate JWT Token
	token, err := GenerateJWT(newUser.ID, newUser.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User created successfully",
		"id":      newUser.ID,
		"token":   token,
	})
}
