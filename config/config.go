package config

import (
	"gopkg.in/yaml.v2"
	"os"
)

// Config 配置信息
type Config struct {
	Port            int    `yaml:"port"`
	ProxyServer     string `yaml:"proxyServer"`
	RateLimiterMode int    `yaml:"rateLimiterMode"`
	RateLimiter     struct {
		TokenBucket struct {
			MaxToken       int `yaml:"maxToken"`
			TokenPerSecond int `yaml:"tokenPerSecond"`
		} `yaml:"tokenBucket"`
		LeakyBucket struct {
			Capacity       int `yaml:"capacity"`
			LeakyPerSecond int `yaml:"leakyPerSecond"`
		} `yaml:"leakyBucket"`
		FixedWindow struct {
			WindowSize  int `yaml:"windowSize"`
			MaxRequests int `yaml:"maxRequests"`
		} `yaml:"fixedWindow"`
		SlideWindow struct {
			WindowSize  int `yaml:"windowSize"`
			MaxRequests int `yaml:"maxRequests"`
		} `yaml:"slideWindow"`
	} `yaml:"rateLimiter"`
	Redis struct {
		Address  string `yaml:"address"`
		Password string `yaml:"password"`
		Database int    `yaml:"database"`
	} `yaml:"redis"`
}

// ReadConfig 读取配置文件
func ReadConfig(configFile string) (Config, error) {
	// 读取 YAML 配置文件
	data, err := os.ReadFile(configFile)
	if err != nil {
		return Config{}, err
	}
	// 解析 YAML 文件到 Config 结构体
	var config Config
	err = yaml.Unmarshal(data, &config)
	return config, err
}
