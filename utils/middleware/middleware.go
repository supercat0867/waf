package middleware

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"net/http"
	"sync"
	"waf/config"
	"waf/utils/rateLimiter"
)

var ipLimitersMutex sync.Mutex
var ipLimiters = make(map[string]rateLimiter.RateLimiterInterface)

// RateLimitMiddleware 限速中间件
func RateLimitMiddleware(config config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取客户端ip
		clientIP := c.ClientIP()
		// 给限速实例map加锁
		ipLimitersMutex.Lock()
		limiter, exists := ipLimiters[clientIP]
		if !exists {
			limiter = rateLimiter.NewRateLimiter(config)
			ipLimiters[clientIP] = limiter
		}
		// 解锁
		ipLimitersMutex.Unlock()
		if !limiter.Allow() {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"error": "Too Many Request",
			})
			return
		}
		c.Next()
	}
}

var ctx = context.Background()

// BlacklistMiddleware IP黑名单中间件
func BlacklistMiddleware(rdb *redis.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取客户端ip
		clientIP := c.ClientIP()
		// 从 Redis 检查 IP 是否在黑名单中
		exists, err := rdb.SIsMember(ctx, "blacklist", clientIP).Result()
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": "Internal server error",
			})
			return
		}
		if exists {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error": "Your IP address is blacklisted",
			})
			return
		}
		c.Next()
	}
}
