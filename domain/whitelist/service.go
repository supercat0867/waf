package whitelist

import "context"

// Service 白名单服务
type Service struct {
	r Repository
}

// NewWhitelistService 实例化白名单服务
func NewWhitelistService(r Repository, ctx context.Context) *Service {
	r.EnsureWhitelistExists(ctx)
	return &Service{
		r: r,
	}
}

// Add 添加IP至白名单
func (c *Service) Add(ip string, ctx context.Context) error {
	err := c.r.Add(ip, ctx)
	if err != nil {
		return ErrIpAdd
	}
	return nil
}

// Remove 移除白名单IP
func (c *Service) Remove(ip string, ctx context.Context) error {
	err := c.r.Remove(ip, ctx)
	if err != nil {
		return ErrIpRemove
	}
	return nil
}

// Get 获取白名单IP
func (c *Service) Get(ctx context.Context) ([]string, error) {
	return c.r.Get(ctx)
}
