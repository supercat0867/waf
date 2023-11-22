package validator

import (
	"net"
	"strings"
)

// ValidateIPorCIDR 检查格式是否为Ipv4或CIDR
func ValidateIPorCIDR(ip string) bool {
	if strings.Contains(ip, "/") {
		// 检查 CIDR 格式
		_, _, err := net.ParseCIDR(ip)
		return err == nil
	} else {
		// 检查单个 IP 格式
		parsedIP := net.ParseIP(ip)
		return parsedIP != nil
	}
}
