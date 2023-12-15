package domain

import "time"

var (
	IpBlacklist = "Waf_IP_Blacklist"
	IpWhitelist = "Waf_IP_Whitelist"
)

type IpInfo struct {
	IP        string        `json:"ip"`
	ExpiresIn time.Duration `json:"exp"`
}
