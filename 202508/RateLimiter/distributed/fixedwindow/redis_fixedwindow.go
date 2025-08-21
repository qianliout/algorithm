package fixedwindow

import (
	"context"
	"time"

	redis "github.com/redis/go-redis/v9"
)

// Manager：基于 Redis 的固定窗口计数限流器。
// 每个 key 在 Hash 中维护 window_id 与 count。
type Manager struct {
	client     *redis.Client
	limit      int
	windowSize time.Duration
	keyPrefix  string
	script     *redis.Script
}

// NewManager 创建 Redis 版固定窗口管理器。
func NewManager(client *redis.Client, limit int, windowSize time.Duration, keyPrefix string) *Manager {
	if keyPrefix == "" {
		keyPrefix = "rl:fw:"
	}
	return &Manager{client: client, limit: limit, windowSize: windowSize, keyPrefix: keyPrefix, script: redis.NewScript(luaFixedWindow)}
}

// Allow 为指定 key 接受 1 个请求。
func (m *Manager) Allow(ctx context.Context, key string) (bool, error) { return m.AllowN(ctx, key, 1) }

// AllowN 为指定 key 接受 n 个请求。
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
// KEYS[1]：Hash，字段 window_id、count
// ARGV：limit、window_ms、n、now_ms
const luaFixedWindow = `
local key = KEYS[1]
local limit = tonumber(ARGV[1])
local window_ms = tonumber(ARGV[2])
local n = tonumber(ARGV[3])
local now = tonumber(redis.call('TIME')[1]) * 1000

local cur_id = math.floor(now / window_ms)
local stored_id = tonumber(redis.call('HGET', key, 'window_id'))
local count = tonumber(redis.call('HGET', key, 'count'))
if count == nil then count = 0 end

if stored_id == nil or stored_id ~= cur_id then
  stored_id = cur_id
  count = 0
	redis.call('HSET', key, 'window_id', cur_id, 'count', n)
	redis.call('PEXPIRE', key, window_ms * 2)
	if n <= limit then
		return 1
	else
		return 0
	end
else
	if count + n <= limit then
		redis.call('HINCRBY', key, 'count', n)
		redis.call('PEXPIRE', key, window_ms * 2)
		return 1
	else
		redis.call('PEXPIRE', key, window_ms * 2)
		return 0
	end
end

if count + n <= limit then
  count = count + n
  redis.call('HSET', key, 'window_id', stored_id, 'count', count)
  return 1
else
  redis.call('HSET', key, 'window_id', stored_id, 'count', count)
  return 0
end
`

// 详细说明：
// 本脚本实现了基于 Redis 的固定窗口限流器。每个 key 以 Hash 形式存储 window_id 和 count。
// 1. 每次请求会判断当前窗口是否已过期，若过期则重置计数。
// 2. 新增逻辑：每次窗口切换或访问时，都会为 Hash 设置过期时间（2倍窗口时长），避免 Redis key 长期堆积。
// 3. 这样可以保证限流数据自动清理，提升资源利用率。
// 4. 过期时间可根据实际业务需求调整。

// 详细解释上述脚本
// 1. 取出参数
// 2. 读取当前窗口 ID 和请求计数
// 3. 判断是否在同一时间窗口内
// 4. 如果在同一窗口内，判断是否超过限制
// 5. 更新状态并返回结果
