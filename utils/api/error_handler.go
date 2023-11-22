package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// HandleError 400错误处理
func HandleError(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, ErrorResponse{Message: err.Error()})
	c.Abort()
	return
}

// HandleInternalServerError 500错误处理
func HandleInternalServerError(c *gin.Context, err error) {
	c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
	c.Abort()
	return
}

// HandleStatusForbiddenError 403错误处理
func HandleStatusForbiddenError(c *gin.Context, err error) {
	c.JSON(http.StatusForbidden, ErrorResponse{Message: err.Error()})
	c.Abort()
	return
}

// HandleStatusTooManyRequestsError 429错误处理
func HandleStatusTooManyRequestsError(c *gin.Context, err error) {
	c.JSON(http.StatusTooManyRequests, ErrorResponse{Message: err.Error()})
	c.Abort()
	return
}
