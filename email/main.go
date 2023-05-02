package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	mail "github.com/xhit/go-simple-mail/v2"
)

type Post struct {
	Address string `json:"address" binding:"required"`
	Content string `json:"content" binding:"required"`
}

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

	server := mail.NewSMTPClient()
	server.Host = smtpHost
	server.Port, _ = strconv.Atoi(smtpPort)
	server.Username = smtpFrom
	server.Password = smtpPassword
	server.Encryption = mail.EncryptionTLS

	smtpClient, err := server.Connect()
	if err != nil {
		log.Fatal(err)
	}

	// Define API routes
	router.POST("/send", func(c *gin.Context) {

		// Bind JSON body to Post struct
		var post Post
		err := c.BindJSON(&post)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		htmlBody, err := os.ReadFile("template.html")
		stringHtmlBody := string(htmlBody)

		stringHtmlBody = strings.Replace(stringHtmlBody, "{{content}}", post.Content, 1)

		fmt.Println("send email to", post.Address, "with body", stringHtmlBody)

		email := mail.NewMSG()
		email.SetFrom("From Me <me@host.com>")
		email.AddTo(post.Address)
		email.SetSubject("Email sobre seu pet")
		email.SetBody(mail.TextHTML, stringHtmlBody)

		err = email.Send(smtpClient)
		if err != nil {
			log.Printf(err.Error())
		}
		fmt.Println("Email Sent Successfully!")

		//close the connection
		smtpClient.Close()

	})

	// Start server
	router.Run(":8081")
}
