// package main

// import (
// 	"database/sql"
// 	"fmt"
// 	"log"
// 	"tea_ume/controllers"

// 	"github.com/gin-gonic/gin"
// 	_ "github.com/lib/pq"
// 	"gorm.io/driver/postgres"
// 	"gorm.io/gorm"
// )

// // Initialize database connection
// // func initDB() {
// // 	var err error
// // 	// Replace with your actual PostgreSQL credentials
// // 	connStr := "user=uccosgf6d945th dbname=d1s8t9ffge4au7 sslmode=require"
// // 	db, err = sql.Open("postgres", connStr)
// // 	if err != nil {
// // 		log.Fatal("Error opening database: ", err)
// // 	}
// // }

// const (
// 	host     = "ccaml3dimis7eh.cluster-czz5s0kz4scl.eu-west-1.rds.amazonaws.com"
// 	port     = 5432 // Default PostgreSQL port
// 	user     = "uccosgf6d945th"
// 	password = "p0480eccda7100b9f420a5743ce123b253d7d0dd45734333095bc4e8ff184ee61"
// 	dbname   = "d1s8t9ffge4au7"
// )

// // DB variable to be used globally
// var DB *gorm.DB

// func InitDB() {
// 	// Create DSN (Data Source Name) for PostgreSQL
// 	dsn := fmt.Sprintf(
// 		"host=%s user=%s password=%s dbname=%s port=%d sslmode=require",
// 		host, user, password, dbname, port,
// 	)

// 	// Connect to PostgreSQL using GORM
// 	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		log.Fatalf("❌ Failed to connect to database: %v", err)
// 	}

// 	// Assign to global variable
// 	DB = db
// 	fmt.Println("✅ Successfully connected to PostgreSQL database!")

// 	// Check connection
// 	sqlDB, err := DB.DB()
// 	if err != nil {
// 		log.Fatalf("❌ Failed to get database instance: %v", err)
// 	}

// 	err = sqlDB.Ping()
// 	if err != nil {
// 		log.Fatalf("❌ Database ping failed: %v", err)
// 	} else {
// 		fmt.Println("🔗 Database connection verified!")
// 	}
// }

// // Signup handler using Gin
// var db *sql.DB

// func main() {
// 	// Initialize the database
// 	InitDB()

// 	defer db.Close()

// 	// Create a new Gin router
// 	router := gin.Default()

// 	// Define the /signup route using Gin
// 	router.POST("/signup", controllers.SignupHandler)

// 	// Start the server
// 	fmt.Println("Server is running on port 8080...")
// 	if err := router.Run(":8080"); err != nil {
// 		log.Fatal("Error starting the server: ", err)
// 	}
// }

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

	// Start the server
	fmt.Println("Server is running on port 8080...")
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Error starting the server: ", err)
	}
}
