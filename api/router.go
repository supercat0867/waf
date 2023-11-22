package api

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
	blacklistApi "waf/api/blacklist"
	userApi "waf/api/user"
	whitelistApi "waf/api/whitelist"
	cfg "waf/config"
	"waf/domain/blacklist"
	"waf/domain/user"
	"waf/domain/whitelist"
	"waf/utils/database"
	"waf/utils/middleware"
)

type Databases struct {
	BlacklistRepository *blacklist.Repository
	WhitelistRepository *whitelist.Repository
	UserRepository      *user.Repository
}

var AppConfig cfg.Config

func CreateDBs() *Databases {
	// 获取配置信息
	config, err := cfg.ReadConfig("./config/config.yaml")
	if err != nil {
		log.Fatalf("配置读取文件错误：%v", err)
	}
	AppConfig = config
	rdb := database.NewRedisDB(config)
	db := database.NewSqliteDB()

	return &Databases{
		BlacklistRepository: blacklist.NewBlacklistRepository(rdb),
		WhitelistRepository: whitelist.NewWhitelistRepository(rdb),
		UserRepository:      user.NewUserRepository(db),
	}

}

// RegisterHandlers 注册控制器
func RegisterHandlers(r *gin.Engine) {
	dbs := *CreateDBs()
	RegisterBlackListHandlers(r, dbs)
	RegisterWhiteListHandlers(r, dbs)
	RegisterUserHandlers(r, dbs)
}

// RegisterBlackListHandlers 注册黑名单控制器
func RegisterBlackListHandlers(r *gin.Engine, dbs Databases) {
	blacklistService := blacklist.NewBlacklistService(*dbs.BlacklistRepository, context.Background())
	blacklistController := blacklistApi.NewBlacklistController(blacklistService)
	blacklistGroup := r.Group("/waf")
	blacklistGroup.POST("blacklist", middleware.AuthUserMiddleware(AppConfig.JwtSetting.SecretKey), blacklistController.AddIPToBlacklist)
	blacklistGroup.DELETE("blacklist", middleware.AuthUserMiddleware(AppConfig.JwtSetting.SecretKey), blacklistController.RemoveIPFromBlacklist)
	blacklistGroup.GET("blacklist", middleware.AuthUserMiddleware(AppConfig.JwtSetting.SecretKey), blacklistController.GetIps)
}

// RegisterWhiteListHandlers 注册白名单控制器
func RegisterWhiteListHandlers(r *gin.Engine, dbs Databases) {
	whitelistService := whitelist.NewWhitelistService(*dbs.WhitelistRepository, context.Background())
	whitelistController := whitelistApi.NewWhitelistController(whitelistService)
	whitelistGroup := r.Group("/waf")
	whitelistGroup.POST("whitelist", middleware.AuthUserMiddleware(AppConfig.JwtSetting.SecretKey), whitelistController.AddIPToWhitelist)
	whitelistGroup.DELETE("whitelist", middleware.AuthUserMiddleware(AppConfig.JwtSetting.SecretKey), whitelistController.RemoveIPFromWhitelist)
	whitelistGroup.GET("whitelist", middleware.AuthUserMiddleware(AppConfig.JwtSetting.SecretKey), whitelistController.GetIps)
}

// RegisterUserHandlers 注册用户控制器
func RegisterUserHandlers(r *gin.Engine, dbs Databases) {
	userService := user.NewUserService(*dbs.UserRepository)
	userController := userApi.NewUserController(userService, &AppConfig)
	userGroup := r.Group("/waf")
	userGroup.POST("/login", userController.Login)
	userGroup.POST("/user", middleware.AuthUserMiddleware(AppConfig.JwtSetting.SecretKey), userController.CreateUser)
	userGroup.PATCH("/user", middleware.AuthUserMiddleware(AppConfig.JwtSetting.SecretKey), userController.ChangePassword)
}
