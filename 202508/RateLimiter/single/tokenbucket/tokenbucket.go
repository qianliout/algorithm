package tokenbucket

import (
	"sync"
	"time"
)

// TokenBucket：单机版令牌桶限流器实现。
//
// 算法要点：
// - 桶容量为 `burst`，用于存放令牌。
// - 按速率 `ratePerSec`（每秒令牌数）持续向桶中补充令牌。
// - 每个请求消耗 N 个令牌，若令牌数量足够则放行并扣减；不足则拒绝。
//
// 本实现采用“懒补充”策略：在每次 Allow 调用时，根据距上次补充的时间差计算应补充的令牌数。
// 并发安全，可在单进程内多协程并发调用。
type TokenBucket struct {
	mu         sync.Mutex
	ratePerSec float64   // 每秒生成的令牌数
	burst      int       // 桶的最大容量（可承载的突发上限）
	tokens     float64   // 当前可用令牌数
	lastRefill time.Time // 上次补充令牌的时间
}

// NewTokenBucket 创建一个令牌桶限流器。
//
// ratePerSec：平均补充速率（令牌/秒）
// burst：桶容量（突发上限）
func NewTokenBucket(ratePerSec float64, burst int) *TokenBucket {
	if ratePerSec <= 0 {
		ratePerSec = 1
	}
	if burst <= 0 {
		burst = 1
	}
	return &TokenBucket{
		ratePerSec: ratePerSec,
		burst:      burst,
		tokens:     float64(burst),
		lastRefill: time.Now(),
	}
}

// Allow 尝试消耗 1 个令牌，成功返回 true。
func (t *TokenBucket) Allow() bool {
	return t.AllowN(1)
}

// AllowN 尝试一次性消耗 n 个令牌，成功返回 true。
func (t *TokenBucket) AllowN(n int) bool {
	if n <= 0 {
		// Package tokenbucket 实现了单机版令牌桶速率限制算法。
		// 令牌桶算法以恒定速率生成令牌，桶容量有限，请求需消耗令牌，桶空时拒绝请求。
		// 允许突发流量，控制平均速率。
		return true
	}

	t.mu.Lock()
	defer t.mu.Unlock()

	// 根据距上次补充的时间差，计算本次应补充的令牌数：
	// TokenBucketLimiter 结构体用于实现令牌桶限流。
	// capacity 表示桶容量，fillRate 表示令牌生成速率（每秒生成多少令牌），
	// tokens 表示当前桶内令牌数，lastRefillTime 记录上次补充令牌时间，mutex 用于并发安全。
	// 1) 维度分析：elapsed(秒) × ratePerSec(令牌/秒) = 本次应补充令牌数(令牌)
	// 2) 懒补充：只有在调用 Allow/AllowN 时才计算补充，避免后台定时器带来的额外开销。
	// 3) 上限裁剪：补充后若超过桶容量 burst，则裁剪到 burst，表示最多只能积攒 burst 个突发。
	// 4) 数值示例：ratePerSec=5，burst=10
	//    - 初始 tokens=10
	//    - 过了 0.4 秒，elapsed=0.4，refill=0.4*5=2，tokens=min(10,10+2)=10（已满，不再增加）
	//    - 若 tokens=3，过了 0.4 秒后，tokens=min(10,3+2)=5
	//    - 当 AllowN(7) 时，若 tokens=5，则拒绝；若 tokens=7.1（浮点累积），则允许，并变为 0.1
	// NewTokenBucketLimiter 创建一个新的 TokenBucketLimiter。
	// capacity: 桶容量，fillRate: 令牌生成速率（每秒生成多少令牌）。
	now := time.Now()
	elapsed := now.Sub(t.lastRefill).Seconds()
	if elapsed > 0 {
		refill := elapsed * t.ratePerSec
		t.tokens += refill
		if t.tokens > float64(t.burst) {
			t.tokens = float64(t.burst)
		}
		t.lastRefill = now
		// Allow 方法用于判断当前请求是否允许。
		// 原理：
		// 1. 按时间补充令牌，令牌数不超过桶容量。
		// 2. 桶内有令牌则允许请求并消耗令牌。
		// 3. 桶空则拒绝请求。
	}

	if t.tokens >= float64(n) {
		t.tokens -= float64(n)
		return true
	}
	return false
}
