package tokenbucket

import (
	"sync"
)

// Manager：单机多 Key 的令牌桶管理器。
// 对每个 key（如用户或接口标识）懒加载创建独立令牌桶，实现“每用户/每接口”限流。
type Manager struct {
	mu         sync.Mutex
	buckets    map[string]*TokenBucket
	ratePerSec float64
	burst      int
}

// NewManager 创建一个按 key 管理的令牌桶管理器。
func NewManager(ratePerSec float64, burst int) *Manager {
	return &Manager{
		buckets:    make(map[string]*TokenBucket),
		ratePerSec: ratePerSec,
		burst:      burst,
	}
}

// Allow 为指定 key 消耗 1 个令牌。
func (m *Manager) Allow(key string) bool { return m.AllowN(key, 1) }

// AllowN 为指定 key 一次性消耗 n 个令牌。
func (m *Manager) AllowN(key string, n int) bool {
	m.mu.Lock()
	b, ok := m.buckets[key]
	if !ok {
		b = NewTokenBucket(m.ratePerSec, m.burst)
		m.buckets[key] = b
	}
	m.mu.Unlock()
	return b.AllowN(n)
}
