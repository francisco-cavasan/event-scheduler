package main

import (
	"log"
	"os"
	"where_my_pet_at/server/Controllers"
	"where_my_pet_at/server/Database"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	// Retrieve database credentials from environment variables
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("MYSQL_USER")
	dbPass := os.Getenv("MYSQL_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	// Initialize database connection
	db, err := gorm.Open("mysql", dbUser+":"+dbPass+"@tcp("+dbHost+":"+dbPort+")/"+dbName+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatalf("Error connecting to database: %s", err)
	}

	Database.MigrateDatabaseSchema(db)

	defer db.Close()

	// Initialize PetController with database connection
	petController := &Controllers.PetController{DB: db}
	authController := &Controllers.AuthController{DB: db}

	// Initialize Gin router
	router := gin.Default()

	router.Use(cors.Default())

	// Define API routes
	router.GET("/pets", petController.Index)
	router.POST("/pets", petController.Store)
	router.PUT("/pets/:id", petController.Update)
	router.GET("/pets/:id", petController.Get)
	router.DELETE("/pets/:id", petController.Delete)
	router.POST("/pets/found", petController.AddPetFoundLocation)
	router.POST("/login", authController.Login)
	router.POST("/logout", authController.Logout)
	router.POST("/register", authController.Register)

	// Start server
	router.Run(":8080")
}
