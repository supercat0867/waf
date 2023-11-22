package api

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
	blacklistApi "waf/api/blacklist"
	whitelistApi "waf/api/whitelist"
	cfg "waf/config"
	"waf/domain/blacklist"
	"waf/domain/whitelist"
	"waf/utils/database"
)

type Databases struct {
	BlacklistRepository *blacklist.Repository
	WhitelistRepository *whitelist.Repository
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

	return &Databases{
		BlacklistRepository: blacklist.NewBlacklistRepository(rdb),
		WhitelistRepository: whitelist.NewWhitelistRepository(rdb),
	}

}

// RegisterHandlers 注册控制器
func RegisterHandlers(r *gin.Engine) {
	dbs := *CreateDBs()
	RegisterBlackListHandlers(r, dbs)
	RegisterWhiteListHandlers(r, dbs)
}

// RegisterBlackListHandlers 注册黑名单控制器
func RegisterBlackListHandlers(r *gin.Engine, dbs Databases) {
	blacklistService := blacklist.NewBlacklistService(*dbs.BlacklistRepository, context.Background())
	blacklistController := blacklistApi.NewBlacklistController(blacklistService)
	blacklistGroup := r.Group("/blacklist")
	blacklistGroup.POST("", blacklistController.AddIPToBlacklist)
	blacklistGroup.DELETE("", blacklistController.RemoveIPFromBlacklist)
}

// RegisterWhiteListHandlers 注册白名单控制器
func RegisterWhiteListHandlers(r *gin.Engine, dbs Databases) {
	whitelistService := whitelist.NewWhitelistService(*dbs.WhitelistRepository, context.Background())
	whitelistController := whitelistApi.NewWhitelistController(whitelistService)
	whitelistGroup := r.Group("/whitelist")
	whitelistGroup.POST("", whitelistController.AddIPToWhitelist)
	whitelistGroup.DELETE("", whitelistController.RemoveIPFromWhitelist)
}
