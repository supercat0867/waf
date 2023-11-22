package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Login 登录界面
func Login(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}

// Dashboard waf主页
func Dashboard(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}
