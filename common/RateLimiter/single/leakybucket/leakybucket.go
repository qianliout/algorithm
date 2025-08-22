package leakybucket

import (
	"sync"
	"time"
)

// LeakyBucket：单机版漏桶限流器。
//
// 模型：以恒定速率漏水；每个请求视为加 1 单位水位。若水位将超过容量则拒绝，否则接受。
// 本实现按“懒计算”方式，根据时间差计算应漏出的水量。
type LeakyBucket struct {
	mu       sync.Mutex
	capacity float64 // 桶最大水位（容量）
	// Package leakybucket 实现了单机版的漏桶速率限制算法。
	// 漏桶算法将请求视为水流入桶，桶底以恒定速率漏水，超出桶容量的请求被丢弃。
	// 该算法能强制平滑流量，保护后端服务。
	leakPerSec   float64   // 漏出速率（单位/秒）
	level        float64   // 当前水位
	lastLeakTime time.Time // 上次计算漏出的时间
}

// New 创建漏桶限流器。
// leakPerSec：恒定处理速率；capacity：最大允许排队量。
func New(leakPerSec float64, capacity int) *LeakyBucket {
	if leakPerSec <= 0 {
		leakPerSec = 1
	}
	if capacity <= 0 {
		capacity = 1
	}
	return &LeakyBucket{
		capacity:   float64(capacity),
		leakPerSec: leakPerSec,
		level:      0,
		// LeakyBucketLimiter 结构体用于实现漏桶限流。
		// capacity 表示桶的最大容量，leakRate 表示漏水速率（单位：每秒漏多少请求），
		// mutex 用于并发安全，water 表示当前桶内的水量（即排队请求数），lastLeakTime 记录上次漏水时间。
		lastLeakTime: time.Now(),
	}
}

// Allow 入队 1 个请求，接受返回 true。
func (l *LeakyBucket) Allow() bool { return l.AllowN(1) }

// AllowN 原子地入队 n 个请求，接受返回 true。
func (l *LeakyBucket) AllowN(n int) bool {
	if n <= 0 {
		return true
	}

	l.mu.Lock()
	defer l.mu.Unlock()

	// 1. 计算距离上次漏水的时间（秒）。
	now := time.Now()
	elapsed := now.Sub(l.lastLeakTime).Seconds()
	// 2. 按漏水速率，计算这段时间应漏出的水量，并减少当前水位。
	if elapsed > 0 {
		l.level -= elapsed * l.leakPerSec // 水位减少 = 时间差 × 漏速
		if l.level < 0 {
			l.level = 0 // 水位不能为负
		}
		l.lastLeakTime = now // 更新漏水时间点
	}

	// 3. 判断本次请求（n个）是否能全部入桶。
	requested := float64(n)
	if l.level+requested <= l.capacity {
		l.level += requested // 累加水位，表示请求入队
		return true          // 接受请求
	}
	// 4. 超过桶容量则拒绝
	return false
}

// Manager：按 key 管理的漏桶集合（每用户/每接口独立限流）。
type Manager struct {
	mu         sync.Mutex
	buckets    map[string]*LeakyBucket
	leakPerSec float64
	capacity   int
}

// NewManager 创建按 key 的漏桶管理器。
func NewManager(leakPerSec float64, capacity int) *Manager {
	return &Manager{buckets: make(map[string]*LeakyBucket), leakPerSec: leakPerSec, capacity: capacity}
}

// Allow 为指定 key 入队 1 个请求。
func (m *Manager) Allow(key string) bool { return m.AllowN(key, 1) }

// AllowN 为指定 key 入队 n 个请求。
func (m *Manager) AllowN(key string, n int) bool {
	m.mu.Lock()
	b, ok := m.buckets[key]
	if !ok {
		b = New(m.leakPerSec, m.capacity)
		m.buckets[key] = b
	}
	m.mu.Unlock()
	return b.AllowN(n)
}
