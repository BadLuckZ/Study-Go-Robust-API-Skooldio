package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

// limit 5 requests/second and limit 5 requests burst size
var limiter = rate.NewLimiter(5, 5)

func RateLimit() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !limiter.Allow() {
			c.AbortWithStatus(http.StatusTooManyRequests)
			return
		}

		c.Next()
	}
}
