package ratelimiter

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/go-redis/redis/v8"
)

// ===================== 单机限流实现 =====================

// 1. 令牌桶算法（单机版）
// 用于控制请求速率，允许突发流量。
type TokenBucket struct {
	capacity   int        // 桶容量，允许的最大令牌数
	tokens     int        // 当前桶内令牌数
	rate       int        // 令牌生成速率（每秒生成多少令牌）
	lastRefill time.Time  // 上次补充令牌的时间
	mu         sync.Mutex // 互斥锁，保证并发安全
}

// 新建一个令牌桶限流器
// capacity: 桶容量
// rate: 令牌生成速率
func NewTokenBucket(capacity, rate int) *TokenBucket {
	return &TokenBucket{
		capacity:   capacity,
		tokens:     capacity,
		rate:       rate,
		lastRefill: time.Now(),
	}
}

// 判断是否允许通过（令牌桶算法，单机版）
// 返回 true 表示允许，false 表示拒绝
func (tb *TokenBucket) Allow() bool {
	tb.mu.Lock()
	defer tb.mu.Unlock()
	now := time.Now()
	// 计算距离上次补充令牌的时间，补充令牌
	elapsed := now.Sub(tb.lastRefill).Seconds()
	refill := int(elapsed * float64(tb.rate))
	if refill > 0 {
		tb.tokens = min(tb.capacity, tb.tokens+refill)
		tb.lastRefill = now
	}
	// 有令牌则通过
	if tb.tokens > 0 {
		tb.tokens--
		return true
	}
	return false
}

// 2. 漏桶
type LeakyBucket struct {
	capacity     int
	remaining    int
	leakRate     float64 // 每秒漏出多少
	lastLeakTime time.Time
	mu           sync.Mutex
}

func NewLeakyBucket(capacity int, leakRate float64) *LeakyBucket {
	return &LeakyBucket{
		capacity:     capacity,
		remaining:    0,
		leakRate:     leakRate,
		lastLeakTime: time.Now(),
	}
}

func (lb *LeakyBucket) Allow() bool {
	lb.mu.Lock()
	defer lb.mu.Unlock()
	now := time.Now()
	elapsed := now.Sub(lb.lastLeakTime).Seconds()
	leaked := int(elapsed * lb.leakRate)
	if leaked > 0 {
		lb.remaining = max(0, lb.remaining-leaked)
		lb.lastLeakTime = now
	}
	if lb.remaining < lb.capacity {
		lb.remaining++
		return true
	}
	return false
}

// 3. 固定窗口计数器
type FixedWindow struct {
	windowSize  time.Duration
	limit       int
	count       int
	windowStart time.Time
	mu          sync.Mutex
}

func NewFixedWindow(windowSize time.Duration, limit int) *FixedWindow {
	return &FixedWindow{
		windowSize:  windowSize,
		limit:       limit,
		windowStart: time.Now(),
	}
}

func (fw *FixedWindow) Allow() bool {
	fw.mu.Lock()
	defer fw.mu.Unlock()
	now := time.Now()
	if now.Sub(fw.windowStart) >= fw.windowSize {
		fw.windowStart = now
		fw.count = 0
	}
	if fw.count < fw.limit {
		fw.count++
		return true
	}
	return false
}

// 4. 滑动窗口日志
type SlidingWindowLog struct {
	windowSize time.Duration
	limit      int
	timestamps []time.Time
	mu         sync.Mutex
}

func NewSlidingWindowLog(windowSize time.Duration, limit int) *SlidingWindowLog {
	return &SlidingWindowLog{
		windowSize: windowSize,
		limit:      limit,
	}
}

func (sw *SlidingWindowLog) Allow() bool {
	sw.mu.Lock()
	defer sw.mu.Unlock()
	now := time.Now()
	cutoff := now.Add(-sw.windowSize)
	// 移除窗口外的时间戳
	idx := 0
	for _, t := range sw.timestamps {
		if t.After(cutoff) {
			break
		}
		idx++
	}
	sw.timestamps = sw.timestamps[idx:]
	if len(sw.timestamps) < sw.limit {
		sw.timestamps = append(sw.timestamps, now)
		return true
	}
	return false
}

// 工具函数
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// ===================== 分布式（Redis）实现 =====================

var ctx = context.Background()

// Redis令牌桶
func RedisTokenBucketAllow(rdb *redis.Client, key string, capacity, rate int) (bool, error) {
	now := time.Now().Unix()
	script := `
	local bucket = redis.call('HMGET', KEYS[1], 'tokens', 'lastRefill')
	local tokens = tonumber(bucket[1]) or tonumber(ARGV[1])
	local lastRefill = tonumber(bucket[2]) or 0
	local rate = tonumber(ARGV[2])
	local capacity = tonumber(ARGV[1])
	local now = tonumber(ARGV[3])
	local refill = math.floor((now - lastRefill) * rate)
	tokens = math.min(capacity, tokens + refill)
	if tokens > 0 then
		tokens = tokens - 1
		redis.call('HMSET', KEYS[1], 'tokens', tokens, 'lastRefill', now)
		redis.call('EXPIRE', KEYS[1], 3600)
		return 1
	else
		redis.call('HMSET', KEYS[1], 'tokens', tokens, 'lastRefill', now)
		redis.call('EXPIRE', KEYS[1], 3600)
		return 0
	end
	`
	res, err := rdb.Eval(ctx, script, []string{key}, capacity, rate, now).Int()
	return res == 1, err
}

// Redis漏桶
func RedisLeakyBucketAllow(rdb *redis.Client, key string, capacity int, leakRate float64) (bool, error) {
	now := time.Now().Unix()
	script := `
	local bucket = redis.call('HMGET', KEYS[1], 'remaining', 'lastLeakTime')
	local remaining = tonumber(bucket[1]) or 0
	local lastLeakTime = tonumber(bucket[2]) or 0
	local leakRate = tonumber(ARGV[1])
	local capacity = tonumber(ARGV[2])
	local now = tonumber(ARGV[3])
	local leaked = math.floor((now - lastLeakTime) * leakRate)
	remaining = math.max(0, remaining - leaked)
	if remaining < capacity then
		remaining = remaining + 1
		redis.call('HMSET', KEYS[1], 'remaining', remaining, 'lastLeakTime', now)
		redis.call('EXPIRE', KEYS[1], 3600)
		return 1
	else
		redis.call('HMSET', KEYS[1], 'remaining', remaining, 'lastLeakTime', now)
		redis.call('EXPIRE', KEYS[1], 3600)
		return 0
	end
	`
	res, err := rdb.Eval(ctx, script, []string{key}, leakRate, capacity, now).Int()
	return res == 1, err
}

// Redis固定窗口计数器
func RedisFixedWindowAllow(rdb *redis.Client, key string, windowSizeSec, limit int) (bool, error) {
	now := time.Now().Unix()
	window := now / int64(windowSizeSec)
	windowKey := key + ":" + fmt.Sprintf("%d", window)
	count, err := rdb.Incr(ctx, windowKey).Result()
	if err != nil {
		return false, err
	}
	if count == 1 {
		rdb.Expire(ctx, windowKey, time.Duration(windowSizeSec)*time.Second)
	}
	return count <= int64(limit), nil
}

// Redis滑动窗口日志（ZSET实现）
func RedisSlidingWindowLogAllow(rdb *redis.Client, key string, windowSizeSec, limit int) (bool, error) {
	now := time.Now().Unix()
	minScore := now - int64(windowSizeSec)
	pipe := rdb.TxPipeline()
	pipe.ZRemRangeByScore(ctx, key, "-inf", fmt.Sprintf("%d", minScore))
	pipe.ZAdd(ctx, key, &redis.Z{Score: float64(now), Member: now})
	pipe.Expire(ctx, key, time.Duration(windowSizeSec)*time.Second)
	pipe.ZCard(ctx, key)
	cmds, err := pipe.Exec(ctx)
	if err != nil {
		return false, err
	}
	card := cmds[len(cmds)-1].(*redis.IntCmd).Val()
	return card <= int64(limit), nil
}
