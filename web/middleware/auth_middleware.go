package middleware

import (
	"myproject/web/utils"
	"github.com/gin-gonic/gin"
	"strings"
	"net/http"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authheader := c.GetHeader("Authorization")
		if(authheader == "" || !strings.HasPrefix(authheader, "Bearer ")) {
			c.JSON(http.StatusUnauthorized, gin.H{"error":"missing or invalid token"})
			c.Abort()
			return
		}

		tokenStr := strings.TrimPrefix(authheader, "Bearer ")
		claims, err := utils.ParseToken(tokenStr)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			c.Abort()
			return
		}

		c.Set("userID", claims.UserID)
		c.Set("userName", claims.Username)
		c.Next()
	}
}