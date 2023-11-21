package rateLimiter

import (
	"time"
	"waf/config"
)

// RateLimiterInterface 限速接口
type RateLimiterInterface interface {
	Allow() bool
}

// Allow 令牌桶算法
func (rl *TokenBucket) Allow() bool {
	// 使用互斥锁守护线程
	rl.mu.Lock()
	defer rl.mu.Unlock()
	// 获取当前时间
	now := time.Now()
	// 计算现在距离上次请求过了多少秒
	elapsed := now.Sub(rl.LastTime).Seconds()
	// 将lastTime更新为现在时间，以便下次使用
	rl.LastTime = now
	// 基于经过时间*令牌生成速率来计算需要添加多少令牌
	rl.Tokens += int(elapsed * float64(rl.TokensPerSecond))
	// 当计算结果大于设置的最大令牌数则将当前令牌数设为最大令牌
	if rl.Tokens > rl.MaxTokens {
		rl.Tokens = rl.MaxTokens
	}
	// 当前无令牌，返回false拒绝请求
	if rl.Tokens < 1 {
		return false
	}
	// 执行到这里说明至少存在一个令牌，令牌减1，返回true允许访问
	rl.Tokens--
	return true
}

// Allow 漏桶算法
func (rl *LeakyBucket) Allow() bool {
	// 使用互斥锁守护线程
	rl.mu.Lock()
	defer rl.mu.Unlock()
	// 获取当前时间
	now := time.Now()
	// 计算当前时间和上次漏水发生的时间差
	elapsed := now.Sub(rl.LastLeakTime)
	// 计算漏水量，从上次调用到现在应该漏掉多少水
	leaks := int(elapsed / rl.LeakInterval)
	// 如果桶里的水比漏掉的水多，从桶里减去漏掉的水，否则将桶里的水设置为0
	if leaks > 0 {
		if rl.Remaining > leaks {
			rl.Remaining -= leaks
		} else {
			rl.Remaining = 0
		}
		// 更新上次漏水发生时间为当前时间
		rl.LastLeakTime = now
	}
	// 检查桶是否有足够的空间来容纳新的水（请求），如果有则增加桶里的水，返回true
	if rl.Remaining < rl.Capacity {
		rl.Remaining++
		return true
	}
	// 桶已装满水，返回false拒绝请求
	return false
}

// Allow 固定窗口算法
func (rl *FixedWindow) Allow() bool {
	// 使用互斥锁守护线程
	rl.mu.Lock()
	defer rl.mu.Unlock()
	// 获取当前时间
	now := time.Now()
	// 计算距离上次窗口经过的时间，如果大于设定的窗口值，将上次窗口时间设为当前时间，将请求量重置为0
	if now.Sub(rl.WindowStart) >= rl.WindowSize {
		rl.WindowStart = now
		rl.requests = 0
	}
	// 请求量超过了当前窗口范围设定的最大请求量，返回false
	if rl.requests >= rl.MaxRequests {
		return false
	}
	// 请求量加1，返回true
	rl.requests++
	return true
}

// Allow 滑动窗口算法
func (rl *SlidingWindow) Allow() bool {
	// 使用互斥锁守护线程
	rl.mu.Lock()
	defer rl.mu.Unlock()
	// 获取当前时间
	now := time.Now()
	// 移除早于（当前时间-窗口大小）的时间戳
	var newTimestamps []time.Time
	for _, ts := range rl.timestamps {
		if now.Sub(ts) <= rl.Window {
			newTimestamps = append(newTimestamps, ts)
		}
	}
	// 保存新的时间戳切片
	rl.timestamps = newTimestamps
	// 判断当前请求是否应被允许，如果时间戳队列长度大于最大请求量返回false
	if len(rl.timestamps) >= rl.MaxReq {
		return false
	}
	// 将当前时间戳加入时间戳队列
	rl.timestamps = append(rl.timestamps, now)
	return true
}

// NewRateLimiter 实例化限速器
func NewRateLimiter(config config.Config) RateLimiterInterface {
	rateMode := config.RateLimiterMode
	switch rateMode {
	case 1:
		return &TokenBucket{
			MaxTokens:       config.RateLimiter.TokenBucket.MaxToken,
			TokensPerSecond: config.RateLimiter.TokenBucket.TokenPerSecond,
			Tokens:          config.RateLimiter.TokenBucket.MaxToken,
			LastTime:        time.Now(),
		}
	case 2:
		return &LeakyBucket{
			Capacity:     config.RateLimiter.LeakyBucket.Capacity,
			Remaining:    0,
			LeakInterval: time.Second / time.Duration(config.RateLimiter.LeakyBucket.LeakyPerSecond),
			LastLeakTime: time.Now(),
		}
	case 3:
		return &FixedWindow{
			WindowStart: time.Now(),
			MaxRequests: config.RateLimiter.FixedWindow.MaxRequests,
			WindowSize:  time.Duration(config.RateLimiter.FixedWindow.MaxRequests) * time.Second,
		}
	case 4:
		return &SlidingWindow{
			Window: time.Duration(config.RateLimiter.SlideWindow.WindowSize) * time.Second,
			MaxReq: config.RateLimiter.SlideWindow.MaxRequests,
		}
	default:
		return &TokenBucket{
			MaxTokens:       15,
			TokensPerSecond: 15,
			Tokens:          15,
			LastTime:        time.Now(),
		}
	}
}
