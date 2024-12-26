package route

import (
	"userrelation/internals/handlers"
	"userrelation/internals/middleware"

	"github.com/gin-gonic/gin"
)

// func getUserIP(c *gin.Context) string {
// 	// Get the X-Forwarded-For header value
// 	xff := c.GetHeader("X-Forwarded-For")
// 	if xff != "" {
// 		// If there are multiple IPs in the X-Forwarded-For header, use the first one
// 		ips := strings.Split(xff, ",")
// 		if len(ips) > 0 {
// 			return strings.TrimSpace(ips[0])
// 		}
// 	}
// 	// Fallback to the remote address
// 	return c.ClientIP()
// }

func Routes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/", handlers.HealthCheck())
	incomingRoutes.Use(middleware.Authentication())
	prodRoutes := incomingRoutes.Group("/prod/api/user")
	{
		prodRoutes.GET("/users/viewing/users/users-users-relation", handlers.CheckUsersRelationship())
		prodRoutes.GET("/users/viewing/search/users/users-users-search-relation", handlers.CheckSearchedUsersRelationship())
		prodRoutes.GET("/users/viewing/restaurants/users-restaurant-relation", handlers.CheckRestaurantRelationship())
		prodRoutes.GET("/users/viewing/users/total-users-request", handlers.RequestFollow())
		prodRoutes.GET("/users/viewing/users/total-users-followers", handlers.ViewFollowers())
		prodRoutes.GET("/users/viewing/restaurants/total-restaurants-followers", handlers.ViewFollowedRestaurant())
		prodRoutes.GET("/users/viewing/users/total-users-following", handlers.ViewFollowing())
		prodRoutes.GET("/users/viewing/users/total-users-mutual", handlers.GetMutualFollowers())
		prodRoutes.GET("/user/follow/user/follow-user-user-relation", handlers.Follow())
		prodRoutes.GET("/user/unfollow/user/unfollow-user-user-relation", handlers.UnFollow())
		prodRoutes.GET("/user/follow/request/user/follow-user-request-user-relation", handlers.FollowRequest())
		prodRoutes.GET("/user/follow/request/cancel/user/follow-user-request-user-relation", handlers.CancelRequest())
		prodRoutes.GET("/user/follow/request/accept/user/follow-user-request-user-relation", handlers.AcceptRequest())
		prodRoutes.GET("/user/follow/request/decline/user/follow-user-request-user-relation", handlers.DeclineRequest())
		prodRoutes.GET("/user/follow/restaurant/follow-user-restaurant-relation", handlers.FollowRestaurant())
		prodRoutes.GET("/user/unfollow/restaurant/unfollow-user-restaurant-relation", handlers.UnFollowRestaurant())
		prodRoutes.GET("/users/viewing/users/total-users-user-request", handlers.GetTotalFollowRequest())
		prodRoutes.GET("/users/viewing/users/total-users-user-mutual", handlers.GetMutualFollowersCount())
		prodRoutes.POST("/users/reviewing/restaurnats/feedback", handlers.AddFeedback())
		prodRoutes.GET("/users/checking/restaurnats/feedback", handlers.CheckIfFeedback())
		prodRoutes.GET("/users/viewing/all/restaurnats/feedback", handlers.ViewRestaurantFeedback())
		prodRoutes.GET("/users/viewing/total/restaurnats/feedback/stars", handlers.GetStarCounts())
	}
	devRoutes := incomingRoutes.Group("/dev/api/user")
	{
		devRoutes.GET("/users/viewing/users/users-users-relation", handlers.CheckUsersRelationship())
		devRoutes.GET("/users/viewing/search/users/users-users-search-relation", handlers.CheckSearchedUsersRelationship())
		devRoutes.GET("/users/viewing/restaurants/users-restaurant-relation", handlers.CheckRestaurantRelationship())
		devRoutes.GET("/users/viewing/users/total-users-request", handlers.RequestFollow())
		devRoutes.GET("/users/viewing/users/total-users-followers", handlers.ViewFollowers())
		devRoutes.GET("/users/viewing/restaurants/total-restaurants-followers", handlers.ViewFollowedRestaurant())
		devRoutes.GET("/users/viewing/users/total-users-following", handlers.ViewFollowing())
		devRoutes.GET("/users/viewing/users/total-users-mutual", handlers.GetMutualFollowers())
		devRoutes.GET("/user/follow/user/follow-user-user-relation", handlers.Follow())
		devRoutes.GET("/user/unfollow/user/unfollow-user-user-relation", handlers.UnFollow())
		devRoutes.GET("/user/follow/request/user/follow-user-request-user-relation", handlers.FollowRequest())
		devRoutes.GET("/user/follow/request/cancel/user/follow-user-request-user-relation", handlers.CancelRequest())
		devRoutes.GET("/user/follow/request/accept/user/follow-user-request-user-relation", handlers.AcceptRequest())
		devRoutes.GET("/user/follow/request/decline/user/follow-user-request-user-relation", handlers.DeclineRequest())
		devRoutes.GET("/user/follow/restaurant/follow-user-restaurant-relation", handlers.FollowRestaurant())
		devRoutes.GET("/user/unfollow/restaurant/unfollow-user-restaurant-relation", handlers.UnFollowRestaurant())
		devRoutes.GET("/users/viewing/users/total-users-user-request", handlers.GetTotalFollowRequest())
		devRoutes.GET("/users/viewing/users/total-users-user-mutual", handlers.GetMutualFollowersCount())
		devRoutes.POST("/users/reviewing/restaurnats/feedback", handlers.AddFeedback())
		devRoutes.GET("/users/checking/restaurnats/feedback", handlers.CheckIfFeedback())
		devRoutes.GET("/users/viewing/all/restaurnats/feedback", handlers.ViewRestaurantFeedback())
		devRoutes.GET("/users/viewing/total/restaurnats/feedback/stars", handlers.GetStarCounts())
	}
	incomingRoutes.NoRoute(func(c *gin.Context) {
		// userIP := c.ClientIP()
		// apiEndpoint := c.Request.URL.Path
		// go helpers.KafkaIpLog(context.Background(), userIP, apiEndpoint, http.StatusNotFound, true)
	})
}
