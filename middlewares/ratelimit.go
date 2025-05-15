package middlewares

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

// RateLimiter tracks requests per IP
type RateLimiter struct {
	enabled bool
	limit   int
	window  time.Duration
	visits  map[string][]time.Time
	mu      sync.Mutex
}

// NewRateLimiter configures the rate limiter
func NewRateLimiter(enabled bool, limit int, windowSeconds int) *RateLimiter {
	return &RateLimiter{
		enabled: enabled,
		limit:   limit,
		window:  time.Duration(windowSeconds) * time.Second,
		visits:  make(map[string][]time.Time),
	}
}

// Middleware returns the middleware
func (r *RateLimiter) Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !r.enabled {
			c.Next()
			return
		}

		ip := c.ClientIP()
		now := time.Now()
		windowStart := now.Add(-r.window)

		r.mu.Lock()
		defer r.mu.Unlock()

		times := r.visits[ip]
		var newTimes []time.Time
		for _, t := range times {
			if t.After(windowStart) {
				newTimes = append(newTimes, t)
			}
		}

		if len(newTimes) >= r.limit {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"error": "Rate limit exceeded",
			})
			return
		}

		newTimes = append(newTimes, now)
		r.visits[ip] = newTimes

		c.Next()
	}
}
