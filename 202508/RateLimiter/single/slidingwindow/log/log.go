package log

import (
	"container/list"
	"sync"
	"time"
)

// Limiter：单机版滑动窗口（日志）限流器。
// 维护请求的时间戳队列，移除窗口外的时间戳，窗口内数量不超过 limit 即允许。
type Limiter struct {
	mu         sync.Mutex
	limit      int
	windowSize time.Duration
	// 使用链表以便 O(1) 在两端插入/弹出；元素为到达时间 time.Time
	logs *list.List
}

// New 创建基于日志的滑动窗口限流器。
func New(limit int, windowSize time.Duration) *Limiter {
	// Package log 实现了单机版滑动窗口日志速率限制算法。
	// 滑动窗口日志记录每个请求的时间戳，窗口滑动时移除过期时间戳，统计窗口内请求数。
	// 精度高但内存消耗大，适合精确限流场景。
	if limit <= 0 {
		limit = 1
	}
	if windowSize <= 0 {
		windowSize = time.Second
	}
	return &Limiter{limit: limit, windowSize: windowSize, logs: list.New()}
	// SlidingWindowLogLimiter 结构体用于实现滑动窗口日志限流。
	// windowSize 表示窗口大小（秒），maxRequests 表示窗口内最大请求数。
	// timestamps 存储所有请求的时间戳，mutex 用于并发安全。
}

// Allow 尝试接受 1 个请求。
// Allow 方法用于判断当前请求是否允许。

func (l *Limiter) Allow() bool { return l.AllowN(1) }

// AllowN 尝试一次性接受 n 个请求（用于突发检查）。
func (l *Limiter) AllowN(n int) bool {
	if n <= 0 {
		return true
	}

	now := time.Now()
	threshold := now.Add(-l.windowSize)

	l.mu.Lock()

	defer l.mu.Unlock()
	// 原理：
	// 1. 移除窗口外的过期时间戳。
	// 2. 统计窗口内请求数，未超限则允许请求并记录时间戳。
	// 移除超出窗口的旧时间戳
	for l.logs.Len() > 0 {
		front := l.logs.Front()
		if front.Value.(time.Time).Before(threshold) {
			l.logs.Remove(front)
		} else {
			break
		}
	}

	if l.logs.Len()+n <= l.limit {
		for i := 0; i < n; i++ {
			l.logs.PushBack(now)
		}
		return true
	}
	return false
}

// Manager：按 key 管理的日志滑动窗口限流集合。
type Manager struct {
	mu         sync.Mutex
	limit      int
	windowSize time.Duration
	buckets    map[string]*Limiter
}

// NewManager 创建按 key 的日志滑动窗口限流管理器。
func NewManager(limit int, windowSize time.Duration) *Manager {
	return &Manager{limit: limit, windowSize: windowSize, buckets: make(map[string]*Limiter)}
}

// Allow 为指定 key 接受 1 个请求。
func (m *Manager) Allow(key string) bool { return m.AllowN(key, 1) }

// AllowN 为指定 key 接受 n 个请求。
func (m *Manager) AllowN(key string, n int) bool {
	m.mu.Lock()
	b, ok := m.buckets[key]
	if !ok {
		b = New(m.limit, m.windowSize)
		m.buckets[key] = b
	}
	m.mu.Unlock()
	return b.AllowN(n)
}
