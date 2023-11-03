package middlewares

import (
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

type QpsLimiter struct {
	ClientMap sync.Map // ["127.0.0.1"] : rate/Limiter
}

func NewLimiter() *QpsLimiter {
	return &QpsLimiter{}
}

func (qps *QpsLimiter) allow(remote string, ratelimit *rate.Limiter) bool {

	if limiter, exist := qps.ClientMap.LoadOrStore(remote, ratelimit); exist {
		if ql, ok := limiter.(*rate.Limiter); ok && !ql.Allow() {
			return false
		}
	}

	return true
}

// r: rate per second , b: burst
func (qps *QpsLimiter) RateLimit(r, b int) gin.HandlerFunc {
	return func(c *gin.Context) {
		var limiter = rate.NewLimiter(rate.Limit(r), b)
		if qps.allow(c.RemoteIP(), limiter) {
			c.Next()
			return
		}
		c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
			"error": "too many request of your ip address per second",
		})
	}
}
