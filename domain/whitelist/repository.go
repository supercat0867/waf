package whitelist

import (
	"context"
	"github.com/go-redis/redis/v8"
	"strings"
	"time"
	"waf/domain"
)

// Repository IP白名单仓库
type Repository struct {
	rdb *redis.Client
	ctx context.Context
}

// NewWhitelistRepository 实例化IP白名单仓库
func NewWhitelistRepository(rdb *redis.Client, ctx context.Context) *Repository {
	return &Repository{
		rdb: rdb,
		ctx: ctx,
	}
}

// Add 添加IP至白名单,并设置过期时间
func (r *Repository) Add(ip string, expiration time.Duration) error {
	// 单独为每个IP地址设置键和过期时间
	key := domain.IpWhitelist + ":" + ip
	// 开启一个redis事务
	_, err := r.rdb.TxPipelined(r.ctx, func(pipe redis.Pipeliner) error {
		// 添加 IP 至白名单散列，并设置过期时间
		pipe.Set(r.ctx, key, "true", expiration)
		return nil
	})
	return err
}

// Remove 将IP移出白名单
func (r *Repository) Remove(ip string) error {
	key := domain.IpWhitelist + ":" + ip
	return r.rdb.Del(r.ctx, key).Err()
}

// Get 获取全部白名单IP以及对应的过期时间
func (r *Repository) Get() ([]domain.IpInfo, error) {
	var whitelist []domain.IpInfo

	// 获取所有白名单 IP 地址的键
	keys, err := r.rdb.Keys(r.ctx, domain.IpWhitelist+":*").Result()
	if err != nil {
		return nil, err
	}

	for _, key := range keys {
		// 提取 IP 地址
		ip := strings.TrimPrefix(key, domain.IpWhitelist+":")

		// 获取剩余过期时间
		ttl, err := r.rdb.TTL(r.ctx, key).Result()
		if err != nil {
			return nil, err
		}

		whitelist = append(whitelist, domain.IpInfo{IP: ip, ExpiresIn: ttl})
	}

	return whitelist, nil
}
