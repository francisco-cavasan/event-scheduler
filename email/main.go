package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	mail "github.com/xhit/go-simple-mail/v2"
)

var htmlBody = `
<html>
<head>
   <meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
   <title>Seu pet foi encontrado!</title>
</head>
<body>
   <p>Temos uma boa notícia, seu pet foi encontrado!!</p>
</body>
`

type Post struct {
	address string
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

		var post Post

		err := c.BindJSON(&post)

		if err != nil {
			fmt.Println("error")
			fmt.Println(err)
		}

		email := mail.NewMSG()
		email.SetFrom("From Me <me@host.com>")
		//todo: change to post.address
		email.AddTo("temp@email.com")
		email.SetSubject("Pet encontrado!")
		email.SetBody(mail.TextHTML, htmlBody)

		err = email.Send(smtpClient)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Email Sent Successfully!")
	})

	// Start server
	router.Run(":8081")
}
