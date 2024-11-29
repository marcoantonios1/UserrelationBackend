package middleware

import (
	"net/http"
	"strings"

	token "userrelation/internals/utils"

	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		splitToken := strings.Split(authHeader, "Bearer ")
		if len(splitToken) != 2 {
			c.JSON(http.StatusNotFound, gin.H{"error": "Invalid/Malformed Authorization Header"})
			c.Abort()
			return
		}
		ClientToken := splitToken[1]
		if ClientToken == "" {
			c.JSON(http.StatusForbidden, gin.H{"error": "No Authorization Header Provided"})
			c.Abort()
			return
		}
		claims, err := token.ValidateToken(ClientToken)
		if err != "" {
			if err == "token is expired" {
				c.JSON(http.StatusUnauthorized, gin.H{"error": err})
				c.Abort()
				return
			} else {
				c.JSON(http.StatusForbidden, gin.H{"error": err})
				c.Abort()
				return
			}
		}
		c.Set("id", claims.Uid)
		c.Set("deviceid", claims.DeviceId)
		c.Set("model", claims.Model)
		c.Set("devicetype", claims.DeviceType)
		c.Set("area1", claims.AdminLevel1)
		c.Set("area2", claims.AdminLevel2)
		c.Set("country", claims.Country)
		c.Set("locality", claims.Locality)
		c.Set("permission", claims.Permission)
		c.Set("os", claims.OpSys)
		c.Set("logged", claims.Logged)
		c.Next()
	}
}
