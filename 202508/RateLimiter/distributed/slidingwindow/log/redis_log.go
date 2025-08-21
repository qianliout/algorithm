package log

import (
	"context"
	"time"

	redis "github.com/redis/go-redis/v9"
)

// Manager：基于 Redis 的滑动窗口（日志）限流器，使用 ZSET 存储时间戳。
// 每个 key 使用一个 ZSET，score 为毫秒时间戳，member 为唯一标识。
type Manager struct {
	client     *redis.Client
	limit      int
	windowSize time.Duration
	keyPrefix  string
	script     *redis.Script
}

// NewManager 创建 Redis 版滑动日志限流管理器。
func NewManager(client *redis.Client, limit int, windowSize time.Duration, keyPrefix string) *Manager {
	if keyPrefix == "" {
		keyPrefix = "rl:swl:"
	}
	return &Manager{client: client, limit: limit, windowSize: windowSize, keyPrefix: keyPrefix, script: redis.NewScript(luaSlidingLog)}
}

// Allow 为指定 key 接受 1 个请求。
func (m *Manager) Allow(ctx context.Context, key string) (bool, error) { return m.AllowN(ctx, key, 1) }

// AllowN 若窗口内仍有配额，则为指定 key 接受 n 个请求。
func (m *Manager) AllowN(ctx context.Context, key string, n int) (bool, error) {
	if n <= 0 {
		return true, nil
	}
	nowMs := time.Now().UnixMilli()
	res, err := m.script.Run(ctx, m.client, []string{m.keyPrefix + key}, m.limit, m.windowSize.Milliseconds(), n, nowMs).Int()
	if err != nil {
		return false, err
	}
	return res == 1, nil
}

// Lua 说明：
// KEYS[1]：ZSET
// ARGV：limit、window_ms、n、now_ms
// 返回：允许=1，否则=0
const luaSlidingLog = `
local key = KEYS[1]
local limit = tonumber(ARGV[1])
local window_ms = tonumber(ARGV[2])
local n = tonumber(ARGV[3])
local now = tonumber(ARGV[4])
local threshold = now - window_ms

-- 移除窗口外的旧成员
redis.call('ZREMRANGEBYSCORE', key, '-inf', threshold)

local cur = redis.call('ZCARD', key)
if cur + n <= limit then
  for i=1,n do
    local member = tostring(now) .. '-' .. tostring(i) .. '-' .. tostring(redis.call('INCR', key .. ':seq'))
    redis.call('ZADD', key, now, member)
  end
  return 1
else
  return 0
end
`
