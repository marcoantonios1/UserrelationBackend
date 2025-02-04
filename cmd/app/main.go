package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"userrelation/route"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file:", err)
	}
	kafkaURL := os.Getenv("KAFKA_URL")
	fmt.Println("Kafka URL from .env:", kafkaURL)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"https://restaurants.dolfins.co", "https://dolfins.co", "http://localhost:4200"},
		AllowedMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Origin", "Content-Type", "Authorization"},
	})
	handler := c.Handler(router)

	router.GET("/redirect", func(c *gin.Context) {
		androidPackage := "com.yourapp" // Replace this with your Android package name
		userId := c.Query("userId")
		profileLink := "https://dolfins.co/userProfile/" + userId
		playStoreLink := "https://play.google.com/store/apps/details?id=" + androidPackage

		c.HTML(http.StatusOK, "redirect.html", gin.H{
			"AndroidPackage": androidPackage,
			"ProfileLink":    profileLink,
			"PlayStoreLink":  playStoreLink,
		})
	})

	route.Routes(router)

	// Create an HTTP/2 server
	h2s := &http2.Server{}

	// Create a standard HTTP server with h2c handler
	server := &http.Server{
		Addr:    ":" + port,
		Handler: h2c.NewHandler(handler, h2s),
	}

	// Start the server
	log.Printf("Server is running on port %s", port)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Could not listen on %s: %v\n", port, err)
	}
}
