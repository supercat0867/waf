package web

import (
	"github.com/gin-gonic/gin"
	"waf/web/controller"
)

func RegisterHandle(r *gin.Engine) {
	r.LoadHTMLGlob("templates/*") // 加载 HTML 模板
	webGroup := r.Group("/waf")
	webGroup.GET("/login", controller.Login)
	webGroup.GET("/index", controller.Dashboard)
}
