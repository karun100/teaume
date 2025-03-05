package controllers

import (
	"context"
	"log"
	"tea_ume/ent"
	"tea_ume/ent/user"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

//	type User struct {
//		FullName string `json:"full_name" binding:"required"`
//		Email    string `json:"email" binding:"required,email"`
//		Username string `json:"username" binding:"required"`
//		Password string `json:"password" binding:"required"`
//		Role     string `json:"role" binding:"omitempty,oneof=User Seller"`
//	}
// var Client *ent.Client

func SignupHandler(c *gin.Context) {
	var userRequest struct {
		FullName string    `json:"full_name"`
		Email    string    `json:"email"`
		Username string    `json:"username"`
		Password string    `json:"password"`
		Role     user.Role `json:"role"`
	}

	// Bind JSON input to User struct
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	// Validate that the role is either "User" or "Seller"
	if userRequest.Role != user.RoleUser && userRequest.Role != user.RoleSeller {
		c.JSON(400, gin.H{"error": "Invalid role. Allowed values are 'User' or 'Seller'"})
		return
	}

	// Hash the user's password using bcrypt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userRequest.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(500, gin.H{"error": "Error hashing password"})
		return
	}
	// client := ent.FromContext(c)
	// if Client == nil {
	// 	c.JSON(500, gin.H{"error": "Database client not found in context"})
	// 	return
	// }
	// Create the user in the database using Ent's ORM
	client := c.MustGet("ent").(*ent.Client)
	log.Printf("Client in SignupHandler: %+v", client)

	// Ensure the global Client variable is initialized before using it
	if client == nil {
		c.JSON(500, gin.H{"error": "Database client is not initialized"})
		return
	}

	log.Printf("Creating user with FullName=%s, Email=%s, Username=%s, Role=%s",
		userRequest.FullName, userRequest.Email, userRequest.Username, userRequest.Role)

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
		c.JSON(500, gin.H{"error": "Error inserting user into database"})
		return
	}

	c.JSON(201, gin.H{
		"message": "User created successfully",
		"id":      newUser.ID,
	})
}
