package main

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"net/http/httputil"
	"net/url"
	"waf/api"
	cfg "waf/config"
	_ "waf/docs"
	"waf/utils/database"
	"waf/utils/middleware"
)

// @title WAF
// @description Web Application Firewall By Gin
// @version v1.0
func main() {
	// 获取配置信息
	config, err := cfg.ReadConfig("./config/config.yaml")
	if err != nil {
		log.Fatalf("配置读取文件错误：%v", err)
	}
	// 实例化redis连接
	rdb := database.NewRedisDB(config)

	r := gin.Default()

	// 注册API路由
	api.RegisterHandlers(r)

	// API文档
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 黑名单中间件
	r.Use(middleware.BlacklistMiddleware(rdb))
	// 限速中间件
	r.Use(middleware.RateLimitMiddleware(config))

	// 设置反向代理
	target, _ := url.Parse(config.ProxyServer)
	proxy := httputil.NewSingleHostReverseProxy(target)
	r.NoRoute(func(c *gin.Context) {
		proxy.ServeHTTP(c.Writer, c.Request)
	})

	// 启动WAF
	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}
