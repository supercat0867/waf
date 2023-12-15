package middleware

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"sync"
	"waf/config"
	"waf/domain"
	"waf/utils/api"
	"waf/utils/rateLimiter"
)

var ipLimitersMutex sync.Mutex
var ipLimiters = make(map[string]rateLimiter.RateLimiterInterface)

// RateLimitMiddleware 限速中间件
func RateLimitMiddleware(config config.Config, rdb *redis.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 检查是否为白名单IP
		if skip, _ := c.Get("isWhitelisted"); skip == true {
			c.Next()
			return
		}
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
		if !limiter.Allow(rdb, clientIP) {
			api.HandleStatusTooManyRequestsError(c, api.ErrStatusTooManyRequests)
			return
		}
		c.Next()
	}
}

var ctx = context.Background()

// IPBlackAndWhiteMiddleware IP黑白名单中间件
func IPBlackAndWhiteMiddleware(rdb *redis.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取客户端ip
		clientIP := c.ClientIP()

		// 检查 IP 是否在黑名单中
		blacklistKey := domain.IpBlacklist + ":" + clientIP
		exists, err := rdb.Exists(ctx, blacklistKey).Result()
		if err != nil {
			api.HandleInternalServerError(c, err)
			c.Abort()
			return
		}
		if exists > 0 {
			// 黑名单IP，返回403
			api.HandleStatusForbiddenError(c, api.ErrStatusForbidden)
			c.Abort()
			return
		}

		// 检查 IP 是否在白名单中
		whitelistKey := domain.IpWhitelist + ":" + clientIP
		exists, err = rdb.Exists(ctx, whitelistKey).Result()
		if err != nil {
			api.HandleInternalServerError(c, err)
			c.Abort()
			return
		}
		if exists > 0 {
			// 白名单IP，设置一个白名单IP标识，以便绕过后续中间件
			c.Set("isWhitelisted", true)
		}
		c.Next()
	}
}
