package main

import (
	"fmt"
	"log"
	"net/smtp"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	// Initialize Gin router
	router := gin.Default()

	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := os.Getenv("SMTP_PORT")
	smtpFrom := os.Getenv("SMTP_FROM")
	smtpPassword := os.Getenv("SMTP_PASSWORD")

	router.Use(cors.Default())

	// Authentication.
	auth := smtp.PlainAuth("", smtpFrom, smtpPassword, smtpHost)

	// Define API routes
	router.POST("/send", func(c *gin.Context) {

		var addresses []string
		addresses[0] = c.PostForm("address")
		content := []byte(c.PostForm("content"))

		err := smtp.SendMail(smtpHost+":"+smtpPort, auth, smtpFrom, addresses, content)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("Email Sent Successfully!")
	})

	// Start server
	router.Run(":8081")
}
