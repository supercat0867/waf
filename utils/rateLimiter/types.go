package rateLimiter

import (
	"sync"
	"time"
)

// TokenBucket 令牌桶结构体
type TokenBucket struct {
	mu              sync.Mutex
	Tokens          int       // 当前令牌数
	MaxTokens       int       // 最大令牌数
	TokensPerSecond int       // 令牌生成速率
	LastTime        time.Time // 上个令牌生成时间
	Counter         int       // 计数器
}

// LeakyBucket 漏桶结构体
type LeakyBucket struct {
	mu           sync.Mutex
	Capacity     int           // 桶容量
	Remaining    int           // 当前桶里的水
	LeakInterval time.Duration // 漏水速率
	LastLeakTime time.Time     // 上次漏水时间
	Counter      int           // 计数器
}

// FixedWindow 固定窗口结构体
type FixedWindow struct {
	mu          sync.Mutex
	WindowStart time.Time     // 窗口开始时间
	requests    int           // 窗口当前请求数
	MaxRequests int           // 窗口最大请求数
	WindowSize  time.Duration // 窗口大小
	Counter     int           // 计数器
}

// SlidingWindow 滑动窗口结构体
type SlidingWindow struct {
	mu         sync.Mutex
	timestamps []time.Time   // 时间队列
	Window     time.Duration // 窗口大小
	MaxReq     int           // 最大队列数
	Counter    int           // 计数器
}
