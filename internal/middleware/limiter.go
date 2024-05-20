package middleware

import (
	"net/http"
	"sync"
	"time"

	"go-web-scaffold/internal/pkg/ratelimiter"

	"github.com/gin-gonic/gin"
)

type rateLimiterMiddleware struct {
	limiters     map[string]*ratelimiter.SlidingWindow // 路由 和 限流器 的隐射
	requestLimit int
	windowSize   time.Duration
	bucketCount  int
	mutex        sync.Mutex
}

func RateLimiterMiddleware(requestLimit int, windowsSize time.Duration, bucketCount int) func(*gin.Context) {

	return NewRateLimiterMiddleware(
		requestLimit,
		windowsSize*time.Second,
		bucketCount,
	).Middleware()

}

func NewRateLimiterMiddleware(limit int, windowSize time.Duration, bucketCount int) *rateLimiterMiddleware {
	return &rateLimiterMiddleware{
		limiters:     make(map[string]*ratelimiter.SlidingWindow),
		requestLimit: limit,
		windowSize:   windowSize,
		bucketCount:  bucketCount,
	}
}

func (r *rateLimiterMiddleware) getLimiter(key string) *ratelimiter.SlidingWindow {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	if limiter, exists := r.limiters[key]; exists {
		return limiter
	}
	limiter := ratelimiter.NewSlidingWindow(r.windowSize, r.bucketCount)
	r.limiters[key] = limiter
	return limiter
}

func (r *rateLimiterMiddleware) Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		key := c.FullPath()
		limiter := r.getLimiter(key)
		if !limiter.AddRequest(r.requestLimit) {
			c.JSON(http.StatusTooManyRequests, gin.H{"error": "Too Many Requests"})
			c.Abort()
			return
		}
		c.Next()
	}
}
