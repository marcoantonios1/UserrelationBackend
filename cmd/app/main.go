package main

import (
	"log"
	"net/http"
	"os"
	"userrelation/route"

	"github.com/gin-gonic/gin"
	"github.com/rs/cors"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router := gin.New()
	router.Use(gin.Logger())

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:4200"},
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

	log.Fatal(http.ListenAndServe(":"+port, handler))
}
