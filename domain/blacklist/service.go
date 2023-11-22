package blacklist

import "context"

// Service 黑名单服务
type Service struct {
	r Repository
}

// NewBlacklistService 实例化黑名单服务
func NewBlacklistService(r Repository, ctx context.Context) *Service {
	r.EnsureBlacklistExists(ctx)
	return &Service{
		r: r,
	}
}

// Add 添加IP至黑名单
func (c *Service) Add(ip string, ctx context.Context) error {
	err := c.r.Add(ip, ctx)
	if err != nil {
		return ErrIpAdd
	}
	return nil
}

// Remove 移除黑名单IP
func (c *Service) Remove(ip string, ctx context.Context) error {
	err := c.r.Remove(ip, ctx)
	if err != nil {
		return ErrIpRemove
	}
	return nil
}

// Get 获取黑名单IP
func (c *Service) Get(ctx context.Context) ([]string, error) {
	return c.r.Get(ctx)
}
