package main

import (
	"context"
	"fmt"
	"log"
	"tea_ume/controllers"
	"tea_ume/ent"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

var Client *ent.Client

// InitDB initializes the database connection using Ent ORM
func InitDB() {
	var err error
	// Connect to the database using Ent's Open method
	Client, err = ent.Open("postgres", "host=ccaml3dimis7eh.cluster-czz5s0kz4scl.eu-west-1.rds.amazonaws.com user=uccosgf6d945th password=p0480eccda7100b9f420a5743ce123b253d7d0dd45734333095bc4e8ff184ee61 dbname=d1s8t9ffge4au7 port=5432 sslmode=require")
	if err != nil {
		log.Fatalf("❌ Failed to connect to database: %v", err)
	}

	// Verify the connection
	if err := Client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("❌ Failed to create schema: %v", err)
	}

	fmt.Println("✅ Successfully connected to PostgreSQL database!")
}

func main() {
	// Initialize the database
	InitDB()
	defer Client.Close()

	// Create a new Gin router
	router := gin.Default()
	router.Use(func(c *gin.Context) {
		c.Set("ent", Client)
		c.Next()
	})

	// Define the /signup route using Gin
	router.POST("/signup", controllers.SignupHandler)
	router.POST("/login", controllers.LoginHandler)

	// Protected Routes (Require JWT Middleware)
	protected := router.Group("/")
	protected.Use(controllers.AuthMiddleware())
	{
	}

	// Start the server
	fmt.Println("Server is running on port 8080...")
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Error starting the server: ", err)
	}
}
