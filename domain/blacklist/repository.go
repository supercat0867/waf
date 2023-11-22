package blacklist

import (
	"context"
	"github.com/go-redis/redis/v8"
	"log"
	"waf/domain"
)

// Repository 黑名单仓库
type Repository struct {
	rdb *redis.Client
}

// NewBlacklistRepository 实例化黑名单仓库
func NewBlacklistRepository(rdb *redis.Client) *Repository {
	return &Repository{
		rdb: rdb,
	}
}

// EnsureBlacklistExists 检查redis中是否存在黑名单集合，不存在则创建
func (r *Repository) EnsureBlacklistExists(ctx context.Context) {
	// 检查 blacklist 集合是否存在
	exists, err := r.rdb.Exists(ctx, domain.IpBlacklistSet).Result()
	if err != nil {
		log.Printf("黑名单创建失败：%v", err)
	}
	// 如果不存在，可以选择创建它（可选步骤）
	if exists == 0 {
		// 创建一个空的 blacklist 集合
		// 通过添加然后删除一个临时元素来创建集合
		err = r.rdb.SAdd(ctx, domain.IpBlacklistSet, "temporary_ip").Err()
		if err != nil {
			log.Printf("黑名单临时测试数据添加失败：%v", err)
		}
		err = r.rdb.SRem(ctx, domain.IpBlacklistSet, "temporary_ip").Err()
		if err != nil {
			log.Printf("黑名单临时测试数据删除失败：%v", err)
		}
	}
}

// Add 添加IP至黑名单
func (r *Repository) Add(ip string, ctx context.Context) error {
	return r.rdb.SAdd(ctx, domain.IpBlacklistSet, ip).Err()
}

// Remove 将IP移出黑名单
func (r *Repository) Remove(ip string, ctx context.Context) error {
	return r.rdb.SRem(ctx, domain.IpBlacklistSet, ip).Err()
}

// Get 获取IP黑名单列表
func (r *Repository) Get(ctx context.Context) ([]string, error) {
	ips, err := r.rdb.SMembers(ctx, domain.IpBlacklistSet).Result()
	if err != nil {
		return nil, err
	}
	return ips, nil
}
