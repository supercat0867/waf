package api

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

var userIdText = "userId"

// ParseInt 类型转换
func ParseInt(value string, defaultValue int) int {
	if value == "" {
		return defaultValue
	}
	if result, err := strconv.Atoi(value); err == nil {
		return result
	}
	return defaultValue
}

// GetUserId 从content中获取用户id
func GetUserId(g *gin.Context) uint {
	return uint(ParseInt(g.GetString(userIdText), -1))
}
