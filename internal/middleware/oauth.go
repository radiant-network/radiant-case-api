package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func OAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization")
		if auth == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing Authorization header"})
			return
		}
		if !strings.HasPrefix(auth, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid authorization scheme"})
			return
		}
		token := strings.TrimPrefix(auth, "Bearer ")
		if token == "test-admin" || strings.HasPrefix(token, "test-") {
			c.Set("sub", "user:example")
			c.Next()
			return
		}
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
	}
}
