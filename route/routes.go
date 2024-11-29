package route

import (
	"userrelation/internals/handlers"
	"userrelation/internals/middleware"

	"github.com/gin-gonic/gin"
)

func Routes(incomingRoutes *gin.Engine) {
	userRoutes := incomingRoutes.Group("/api/user")
	{
		userRoutes.Use(middleware.Authentication())
		userRoutes.GET("/users/viewing/users/users-users-relation", handlers.CheckUsersRelationship())
		userRoutes.GET("/users/viewing/search/users/users-users-search-relation", handlers.CheckSearchedUsersRelationship())
		userRoutes.GET("/users/viewing/restaurants/users-restaurant-relation", handlers.CheckRestaurantRelationship())
		userRoutes.GET("/users/viewing/users/total-users-request", handlers.RequestFollow())
		userRoutes.GET("/users/viewing/users/total-users-followers", handlers.ViewFollowers())
		userRoutes.GET("/users/viewing/restaurants/total-restaurants-followers", handlers.ViewFollowedRestaurant())
		userRoutes.GET("/users/viewing/users/total-users-following", handlers.ViewFollowing())
		userRoutes.GET("/users/viewing/users/total-users-mutual", handlers.GetMutualFollowers())
		userRoutes.GET("/user/follow/user/follow-user-user-relation", handlers.Follow())
		userRoutes.GET("/user/unfollow/user/unfollow-user-user-relation", handlers.UnFollow())
		userRoutes.GET("/user/follow/request/user/follow-user-request-user-relation", handlers.FollowRequest())
		userRoutes.GET("/user/follow/request/cancel/user/follow-user-request-user-relation", handlers.CancelRequest())
		userRoutes.GET("/user/follow/request/accept/user/follow-user-request-user-relation", handlers.AcceptRequest())
		userRoutes.GET("/user/follow/request/decline/user/follow-user-request-user-relation", handlers.DeclineRequest())
		userRoutes.GET("/user/follow/restaurant/follow-user-restaurant-relation", handlers.FollowRestaurant())
		userRoutes.GET("/user/unfollow/restaurant/unfollow-user-restaurant-relation", handlers.UnFollowRestaurant())
		userRoutes.GET("/users/viewing/users/total-users-user-request", handlers.GetTotalFollowRequest())
		userRoutes.GET("/users/viewing/users/total-users-user-mutual", handlers.GetMutualFollowersCount())
		userRoutes.POST("/users/reviewing/restaurnats/feedback", handlers.AddFeedback())
		userRoutes.GET("/users/checking/restaurnats/feedback", handlers.CheckIfFeedback())
		userRoutes.GET("/users/viewing/all/restaurnats/feedback", handlers.ViewRestaurantFeedback())
		userRoutes.GET("/users/viewing/total/restaurnats/feedback/stars", handlers.GetStarCounts())
	}
	incomingRoutes.NoRoute(func(c *gin.Context) {
		// userIP := c.ClientIP()
		// apiEndpoint := c.Request.URL.Path
		// go helpers.KafkaIpLog(context.Background(), userIP, apiEndpoint, http.StatusNotFound, true)
	})
}
