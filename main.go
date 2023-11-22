package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"
	"waf/api"
	cfg "waf/config"
	_ "waf/docs"
	"waf/utils/database"
	"waf/utils/graceful"
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

	gin.SetMode(gin.ReleaseMode)
	r := gin.New()

	// 注册API路由
	api.RegisterHandlers(r)

	// API文档
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 注册自定义日志中间件
	registerMiddlewares(r)

	// IP黑白名单中间件
	r.Use(middleware.IPBlackAndWhiteMiddleware(rdb))
	// 限速中间件
	r.Use(middleware.RateLimitMiddleware(config))

	// 设置反向代理
	target, _ := url.Parse(config.ProxyServer)
	proxy := httputil.NewSingleHostReverseProxy(target)
	r.NoRoute(func(c *gin.Context) {
		proxy.ServeHTTP(c.Writer, c.Request)
	})

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", config.Port),
		Handler: r,
	}
	go func() {
		if err = srv.ListenAndServe(); err != nil {
			log.Printf("listen:%s\n", err)
		}
	}()

	wafAsciiArt := `
                               .--.,
         .---.               ,--.'  \
        /. ./|               |  | /\/
     .-'-. ' |    ,--.--.    :  : :
    /___/ \: |   /       \   :  | |-,
 .-'.. '   ' .  .--.  .-. |  |  : :/|
/___/ \:     '   \__\/: . .  |  |  .'
.   \  ' .\      ," .--.; |  '  : '
 \   \   ' \ |  /  /  ,.  |  |  | |
  \   \  |--"  ;  :   .'   \ |  : \
   \   \ |     |  ,     .-./ |  |,'
    '---"       '--''---'     '--'
	`

	log.Println("WAF启动成功")
	fmt.Println(wafAsciiArt)

	graceful.ShutdownGin(srv, time.Second*3)
}

// 注册中间件
func registerMiddlewares(r *gin.Engine) {
	r.Use(
		gin.LoggerWithFormatter(
			func(param gin.LogFormatterParams) string {
				return fmt.Sprintf(
					"%s - [%s]\"%s %s %s %d %s %s\"\n",
					param.ClientIP,
					param.TimeStamp.Format(time.RFC3339),
					param.Method,
					param.Path,
					param.Request.Proto,
					param.StatusCode,
					param.Latency,
					param.ErrorMessage,
				)
			}))
	r.Use(gin.Recovery())
}
