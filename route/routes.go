package route

import (
	"userrelation/controller"
	"userrelation/middleware"

	"github.com/gin-gonic/gin"
)

func Routes(incomingRoutes *gin.Engine) {
	userRoutes := incomingRoutes.Group("/api/user")
	{
		userRoutes.Use(middleware.Authentication())
		userRoutes.GET("/users/viewing/users/users-users-relation", controller.CheckUsersRelationship())
		userRoutes.GET("/users/viewing/restaurants/users-restaurant-relation", controller.CheckRestaurantRelationship())

	}
	incomingRoutes.NoRoute(func(c *gin.Context) {
		// userIP := c.ClientIP()
		// apiEndpoint := c.Request.URL.Path
		// go helpers.KafkaIpLog(context.Background(), userIP, apiEndpoint, http.StatusNotFound, true)
	})
}
