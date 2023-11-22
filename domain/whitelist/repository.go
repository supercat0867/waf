package whitelist

import (
	"context"
	"github.com/go-redis/redis/v8"
	"log"
	"waf/domain"
)

// Repository IP白名单仓库
type Repository struct {
	rdb *redis.Client
}

// NewWhitelistRepository 实例化IP白名单仓库
func NewWhitelistRepository(rdb *redis.Client) *Repository {
	return &Repository{
		rdb: rdb,
	}
}

// EnsureWhitelistExists 检查redis中是否存在白名单集合，不存在则创建
func (r *Repository) EnsureWhitelistExists(ctx context.Context) {
	// 检查 blacklist 集合是否存在
	exists, err := r.rdb.Exists(ctx, domain.IpWhitelistSet).Result()
	if err != nil {
		log.Printf("IP白名单创建失败：%v", err)
	}
	// 如果不存在，可以选择创建它（可选步骤）
	if exists == 0 {
		// 创建一个空的 blacklist 集合
		// 通过添加然后删除一个临时元素来创建集合
		err = r.rdb.SAdd(ctx, domain.IpWhitelistSet, "temporary_ip").Err()
		if err != nil {
			log.Printf("IP白名单临时测试数据添加失败：%v", err)
		}
		err = r.rdb.SRem(ctx, domain.IpWhitelistSet, "temporary_ip").Err()
		if err != nil {
			log.Printf("IP白名单临时测试数据删除失败：%v", err)
		}
	}
}

// Add 添加IP至白名单
func (r *Repository) Add(ip string, ctx context.Context) error {
	return r.rdb.SAdd(ctx, domain.IpWhitelistSet, ip).Err()
}

// Remove 将IP移出白名单
func (r *Repository) Remove(ip string, ctx context.Context) error {
	return r.rdb.SRem(ctx, domain.IpWhitelistSet, ip).Err()
}
