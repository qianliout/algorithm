package fixedwindow

import (
	"sync"
	"time"
)

// Limiter：单机版固定窗口计数限流器。
// 在长度为 windowSize 的窗口内，最多允许 limit 次。窗口切换时计数器重置。
type Limiter struct {
	mu         sync.Mutex
	limit      int
	windowSize time.Duration
	windowId   int64 // 当前窗口标识（基于时间片计算）
	count      int
}

// New 创建固定窗口限流器。
func New(limit int, windowSize time.Duration) *Limiter {
	if limit <= 0 {
		limit = 1
	}
	if windowSize <= 0 {
		windowSize = time.Second
	}
	return &Limiter{limit: limit, windowSize: windowSize}
}

// Allow 判断当前窗口能否接受 1 个请求。
func (l *Limiter) Allow() bool { return l.AllowN(1) }

// AllowN 判断当前窗口能否接受 n 个请求。
func (l *Limiter) AllowN(n int) bool {
	if n <= 0 {
		return true
	}

	now := time.Now()
	windowId := now.UnixNano() / l.windowSize.Nanoseconds()

	l.mu.Lock()
	defer l.mu.Unlock()

	if l.windowId != windowId {
		l.windowId = windowId
		l.count = 0
	}

	if l.count+n <= l.limit {
		l.count += n
		return true
	}
	return false
}

// Manager：按 key 管理的固定窗口限流集合。
type Manager struct {
	mu         sync.Mutex
	limit      int
	windowSize time.Duration
	buckets    map[string]*Limiter
}

// NewManager 创建按 key 的固定窗口限流管理器。
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
