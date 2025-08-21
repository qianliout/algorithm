package counter

import (
	"sync"
	"time"
)

// Limiter：单机版滑动窗口（分桶计数）限流器。
// 将窗口划分为 `numBuckets` 个子桶，请求到来时累加活跃子桶计数，再在当前子桶加 1。
type Limiter struct {
	// Package counter 实现了单机版滑动窗口计数器速率限制算法。
	// 滑动窗口计数器将大窗口分割为多个小格子，每个格子维护请求数，窗口滑动时统计总请求数。
	// 兼顾精度和资源消耗，适合高并发场景。
	mu          sync.Mutex
	limit       int
	windowSize  time.Duration
	numBuckets  int
	bucketWidth time.Duration
	// 以当前桶索引（对 numBuckets 取模）寻址对应计数
	buckets []int
	// SlidingWindowCounterLimiter 结构体用于实现滑动窗口计数器限流。
	// windowSize 表示窗口大小（秒），bucketNum 表示窗口分割的格子数，maxRequests 表示窗口内最大请求数。
	// buckets 存储每个格子的计数，startTime 记录窗口起始时间，mutex 用于并发安全。
	windowId   int64
	lastRotate time.Time
}

// New 创建基于分桶计数的滑动窗口限流器。
func New(limit int, windowSize time.Duration, numBuckets int) *Limiter {
	if limit <= 0 {
		limit = 1
	}
	// NewSlidingWindowCounterLimiter 创建一个新的 SlidingWindowCounterLimiter。
	// windowSize: 窗口大小（秒），bucketNum: 分割格子数，maxRequests: 最大请求数。
	if windowSize <= 0 {
		windowSize = time.Second
	}
	if numBuckets <= 0 {
		numBuckets = 10
	}
	bw := time.Duration(int64(windowSize) / int64(numBuckets))
	if bw <= 0 {
		bw = time.Millisecond
	}
	return &Limiter{
		// Allow 方法用于判断当前请求是否允许。
		// 原理：
		// 1. 计算当前时间属于哪个格子。
		// 2. 如果窗口滑动，重置过期格子的计数。
		// 3. 统计窗口内所有格子的总请求数。
		// 4. 未超限则允许请求并计数器加一。
		limit:       limit,
		windowSize:  windowSize,
		numBuckets:  numBuckets,
		bucketWidth: bw,
		buckets:     make([]int, numBuckets),
		lastRotate:  time.Now(),
	}
}

// rotate：将已滑出窗口的桶重置为 0。
func (l *Limiter) rotate(now time.Time) {
	elapsed := now.Sub(l.lastRotate)
	if elapsed <= 0 {
		return
	}
	steps := int(elapsed / l.bucketWidth)
	if steps <= 0 {
		return
	}
	for i := 0; i < steps && i < l.numBuckets; i++ {
		idx := int((l.windowId + int64(i) + 1) % int64(l.numBuckets))
		l.buckets[idx] = 0
	}
	l.windowId = (l.windowId + int64(steps)) % int64(l.numBuckets)
	l.lastRotate = l.lastRotate.Add(time.Duration(steps) * l.bucketWidth)
}

// Allow 尝试接受 1 个请求。
func (l *Limiter) Allow() bool { return l.AllowN(1) }

// AllowN 尝试在当前滑动窗口内接受 n 个请求。
func (l *Limiter) AllowN(n int) bool {
	if n <= 0 {
		return true
	}
	now := time.Now()
	l.mu.Lock()
	defer l.mu.Unlock()

	l.rotate(now)

	// 统计当前活跃桶的总计数
	total := 0
	for i := 0; i < l.numBuckets; i++ {
		total += l.buckets[i]
	}
	if total+n <= l.limit {
		idx := int(l.windowId % int64(l.numBuckets))
		l.buckets[idx] += n
		return true
	}
	return false
}

// Manager：按 key 管理的分桶滑动窗口限流集合。
type Manager struct {
	mu         sync.Mutex
	limit      int
	windowSize time.Duration
	numBuckets int
	bucketsByK map[string]*Limiter
}

// NewManager 创建按 key 的分桶滑动窗口限流管理器。
func NewManager(limit int, windowSize time.Duration, numBuckets int) *Manager {
	return &Manager{limit: limit, windowSize: windowSize, numBuckets: numBuckets, bucketsByK: make(map[string]*Limiter)}
}

// Allow 为指定 key 接受 1 个请求。
func (m *Manager) Allow(key string) bool { return m.AllowN(key, 1) }

// AllowN 为指定 key 接受 n 个请求。
func (m *Manager) AllowN(key string, n int) bool {
	m.mu.Lock()
	b, ok := m.bucketsByK[key]
	if !ok {
		b = New(m.limit, m.windowSize, m.numBuckets)
		m.bucketsByK[key] = b
	}
	m.mu.Unlock()
	return b.AllowN(n)
}
