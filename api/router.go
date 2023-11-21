package api

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
	blacklist2 "waf/api/blacklist"
	cfg "waf/config"
	"waf/domain/blacklist"
	"waf/utils/database"
)

type Databases struct {
	BlacklistRepository *blacklist.Repository
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
	}

}

// RegisterHandlers 注册所有控制器
func RegisterHandlers(r *gin.Engine) {
	dbs := *CreateDBs()
	RegisterBlackLisHandlers(r, dbs)

}

// RegisterBlackLisHandlers 注册黑名单控制器
func RegisterBlackLisHandlers(r *gin.Engine, dbs Databases) {
	blacklistService := blacklist.NewBlacklistService(*dbs.BlacklistRepository, context.Background())
	blacklistController := blacklist2.NewBlacklistController(blacklistService)
	blacklistGroup := r.Group("/blacklist")
	blacklistGroup.POST("", blacklistController.AddIPToBlacklist)
	blacklistGroup.DELETE("", blacklistController.RemoveIPFromBlacklist)
}
