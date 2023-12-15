package whitelist

import (
	"time"
	"waf/domain"
)

// Service 白名单服务
type Service struct {
	r Repository
}

// NewWhitelistService 实例化白名单服务
func NewWhitelistService(r Repository) *Service {
	return &Service{
		r: r,
	}
}

// Add 添加IP至白名单
func (c *Service) Add(ip string, exp time.Duration) error {
	err := c.r.Add(ip, exp)
	if err != nil {
		return ErrIpAdd
	}
	return nil
}

// Remove 移除白名单IP
func (c *Service) Remove(ip string) error {
	err := c.r.Remove(ip)
	if err != nil {
		return ErrIpRemove
	}
	return nil
}

// Get 获取IP白名单列表
func (c *Service) Get() ([]domain.IpInfo, error) {
	return c.r.Get()
}
