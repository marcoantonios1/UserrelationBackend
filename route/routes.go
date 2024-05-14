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
		userRoutes.GET("/users/viewing/search/users/users-users-search-relation", controller.CheckSearchedUsersRelationship())
		userRoutes.GET("/users/viewing/restaurants/users-restaurant-relation", controller.CheckRestaurantRelationship())
		userRoutes.GET("/users/viewing/users/total-users-request", controller.RequestFollow())
		userRoutes.GET("/user/follow/user/follow-user-user-relation", controller.Follow())
		userRoutes.GET("/user/unfollow/user/unfollow-user-user-relation", controller.UnFollow())
		userRoutes.GET("/user/follow/request/user/follow-user-request-user-relation", controller.FollowRequest())
		userRoutes.GET("/user/follow/request/cancel/user/follow-user-request-user-relation", controller.CancelRequest())
		userRoutes.GET("/user/follow/request/accept/user/follow-user-request-user-relation", controller.AcceptRequest())
		userRoutes.GET("/user/follow/request/decline/user/follow-user-request-user-relation", controller.DeclineRequest())
		userRoutes.GET("/user/follow/restaurant/follow-user-restaurant-relation", controller.FollowRestaurant())
		userRoutes.GET("/user/unfollow/restaurant/unfollow-user-restaurant-relation", controller.UnFollowRestaurant())
		userRoutes.GET("/users/viewing/users/total-users-user-request", controller.GetTotalFollowRequest())
		//userRoutes.GET("/user/search/user/get/user/info", controller.GetOtherUserAndRelationship())

	}
	incomingRoutes.NoRoute(func(c *gin.Context) {
		// userIP := c.ClientIP()
		// apiEndpoint := c.Request.URL.Path
		// go helpers.KafkaIpLog(context.Background(), userIP, apiEndpoint, http.StatusNotFound, true)
	})
}
