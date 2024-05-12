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
		userRoutes.GET("/users/viewing/restaurants/total-users-restaurant-relation", controller.CountRestaurantFollowers())
		userRoutes.GET("/user/follow/user/follow-user-user-relation", controller.Follow())
		userRoutes.GET("/user/follow/request/user/follow-user-request-user-relation", controller.Follow())
		userRoutes.GET("/user/unfollow/user/unfollow-user-user-relation", controller.UnFollow())

	}
	incomingRoutes.NoRoute(func(c *gin.Context) {
		// userIP := c.ClientIP()
		// apiEndpoint := c.Request.URL.Path
		// go helpers.KafkaIpLog(context.Background(), userIP, apiEndpoint, http.StatusNotFound, true)
	})
}
