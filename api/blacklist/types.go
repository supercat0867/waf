package blacklist

import "waf/domain"

// AddIPToBlacklistRequest 添加IP至黑名单请求
type AddIPToBlacklistRequest struct {
	IP     string `json:"ip"`  // IP地址
	Expiry int    `json:"exp"` // 过期时间
}

// RemoveIPTiBlacklistRequest 将IP移出黑名单请求
type RemoveIPTiBlacklistRequest struct {
	IP string `json:"ip"`
}

// Response 通用响应
type Response struct {
	Status  int    `json:"status"`
	Message string `json:"msg"`
}

// IpListResponse IP列表响应
type IpListResponse struct {
	Status  int             `json:"status"`
	Count   int             `json:"count"`
	Message string          `json:"msg"`
	Data    []domain.IpInfo `json:"data"`
}
