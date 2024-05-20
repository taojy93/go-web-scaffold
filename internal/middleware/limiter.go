package middleware

import (
	"go-web-scaffold/internal/ratelimiter"
	"net/http"

	"github.com/gin-gonic/gin"
)

// LimitMiddleware 限流中间件
func LimitMiddleware(c *gin.Context) {

	if !ratelimiter.GslidingWindow.AddRequest() {
		http.Error(c.Writer, "Too Many Requests", http.StatusTooManyRequests)
		return
	}

	c.Next()

}
