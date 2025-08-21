package leakybucket

import (
	"context"
	"time"

	redis "github.com/redis/go-redis/v9"
)

// Manager：基于 Redis 的漏桶限流器。
// 每个 key 的状态存入 Hash：level（水位）、last_ms（上次计算时间）。
type Manager struct {
	client     *redis.Client
	leakPerSec float64
	capacity   int
	keyPrefix  string
	script     *redis.Script
}

// NewManager 创建 Redis 版漏桶管理器。
func NewManager(client *redis.Client, leakPerSec float64, capacity int, keyPrefix string) *Manager {
	if keyPrefix == "" {
		keyPrefix = "rl:lb:"
	}
	return &Manager{client: client, leakPerSec: leakPerSec, capacity: capacity, keyPrefix: keyPrefix, script: redis.NewScript(luaLeakyBucket)}
}

// Allow 为指定 key 入队 1 个请求。
func (m *Manager) Allow(ctx context.Context, key string) (bool, error) { return m.AllowN(ctx, key, 1) }

// AllowN 为指定 key 入队 n 个请求。
func (m *Manager) AllowN(ctx context.Context, key string, n int) (bool, error) {
	if n <= 0 {
		return true, nil
	}
	nowMs := time.Now().UnixMilli()
	res, err := m.script.Run(ctx, m.client, []string{m.keyPrefix + key}, m.leakPerSec, m.capacity, n, nowMs).Int()
	if err != nil {
		return false, err
	}
	return res == 1, nil
}

// Lua 说明：
// KEYS[1]：Hash，包含 level、last_ms
// ARGV：leakPerSec、capacity、n、now_ms
// 返回：接受=1，拒绝=0
const luaLeakyBucket = `
local key = KEYS[1]
local leak = tonumber(ARGV[1])
local capacity = tonumber(ARGV[2])
local n = tonumber(ARGV[3])
local now = tonumber(ARGV[4])

local level = tonumber(redis.call('HGET', key, 'level'))
local last = tonumber(redis.call('HGET', key, 'last_ms'))
if level == nil then level = 0 end
if last == nil then last = now end

local elapsed = math.max(0, (now - last) / 1000.0)
level = math.max(0, level - elapsed * leak)
last = now

if level + n <= capacity then
  level = level + n
  redis.call('HSET', key, 'level', level, 'last_ms', last)
  return 1
else
  redis.call('HSET', key, 'level', level, 'last_ms', last)
  return 0
end
`

// 解释上述脚本
// 1. 取出参数
// 2. 读取当前水位和上次计算时间
// 3. 计算时间差，按漏水速率更新水位
// 4. 判断是否允许入队
