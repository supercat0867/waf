package middleware

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"net"
	"strings"
	"sync"
	"waf/config"
	"waf/domain"
	"waf/utils/api"
	"waf/utils/rateLimiter"
)

var ipLimitersMutex sync.Mutex
var ipLimiters = make(map[string]rateLimiter.RateLimiterInterface)

// RateLimitMiddleware 限速中间件
func RateLimitMiddleware(config config.Config) gin.HandlerFunc {
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
		if !limiter.Allow() {
			api.HandleStatusTooManyRequestsError(c, api.ErrStatusTooManyRequests)
			return
		}
		c.Next()
	}
}

var ctx = context.Background()

// 增加函数来检查 IP 是否在范围内
func isIPInRanges(ip string, ranges []string) bool {
	parsedIP := net.ParseIP(ip)
	if parsedIP == nil {
		return false
	}
	for _, r := range ranges {
		if strings.Contains(r, "/") {
			_, cidrNet, err := net.ParseCIDR(r)
			if err != nil {
				continue
			}
			if cidrNet.Contains(parsedIP) {
				return true
			}
		} else if r == ip {
			return true
		}
	}
	return false
}

// IPBlackAndWhiteMiddleware IP黑白名单中间件
func IPBlackAndWhiteMiddleware(rdb *redis.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取客户端ip
		clientIP := c.ClientIP()
		// 检查 IP 是否在黑名单中
		blacklist, err := rdb.SMembers(ctx, domain.IpBlacklistSet).Result()
		if err != nil {
			api.HandleInternalServerError(c, api.ErrInternalServer)
			c.Abort()
			return
		}
		if isIPInRanges(clientIP, blacklist) {
			// 黑名单IP，返回403
			api.HandleStatusForbiddenError(c, api.ErrStatusForbidden)
			c.Abort()
			return
		}
		// 检查 IP 是否在白名单中
		whitelist, err := rdb.SMembers(ctx, domain.IpWhitelistSet).Result()
		if err != nil {
			api.HandleInternalServerError(c, api.ErrInternalServer)
			c.Abort()
			return
		}
		if isIPInRanges(clientIP, whitelist) {
			// 白名单IP，设置一个白名单IP标识，以便绕过后续中间件
			c.Set("isWhitelisted", true)
		}
		c.Next()
	}
}
