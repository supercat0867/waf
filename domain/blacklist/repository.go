package blacklist

import (
	"context"
	"github.com/go-redis/redis/v8"
	"strings"
	"time"
	"waf/domain"
)

// Repository 黑名单仓库
type Repository struct {
	rdb *redis.Client
	ctx context.Context
}

// NewBlacklistRepository 实例化黑名单仓库
func NewBlacklistRepository(rdb *redis.Client, ctx context.Context) *Repository {
	return &Repository{
		rdb: rdb,
		ctx: ctx,
	}
}

// Add 添加IP至黑名单,并设置过期时间
func (r *Repository) Add(ip string, expiration time.Duration) error {
	// 单独为每个IP地址设置键和过期时间
	key := domain.IpBlacklist + ":" + ip
	// 开启一个redis事务
	_, err := r.rdb.TxPipelined(r.ctx, func(pipe redis.Pipeliner) error {
		// 添加 IP 至黑名单散列，并设置过期时间
		pipe.Set(r.ctx, key, "true", expiration)
		return nil
	})
	return err
}

// Remove 将IP移出黑名单
func (r *Repository) Remove(ip string) error {
	key := domain.IpBlacklist + ":" + ip
	return r.rdb.Del(r.ctx, key).Err()
}

// Get 获取全部黑名单IP以及对应的过期时间
func (r *Repository) Get() ([]domain.IpInfo, error) {
	var blacklist []domain.IpInfo

	// 获取所有黑名单 IP 地址的键
	keys, err := r.rdb.Keys(r.ctx, domain.IpBlacklist+":*").Result()
	if err != nil {
		return nil, err
	}

	for _, key := range keys {
		// 提取 IP 地址
		ip := strings.TrimPrefix(key, domain.IpBlacklist+":")

		// 获取剩余过期时间
		ttl, err := r.rdb.TTL(r.ctx, key).Result()
		if err != nil {
			return nil, err
		}

		blacklist = append(blacklist, domain.IpInfo{IP: ip, ExpiresIn: ttl})
	}
	return blacklist, nil
}
