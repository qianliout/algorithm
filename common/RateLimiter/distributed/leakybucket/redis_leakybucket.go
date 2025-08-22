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
  redis.call('PEXPIRE', key, 3600 * 1000) -- 设置过期时间为1小时 (3600秒 * 1000毫秒/秒)
  return 1
else
  redis.call('HSET', key, 'level', level, 'last_ms', last) -- 即使请求被拒绝，也更新last_ms和level（level未增加）
  redis.call('PEXPIRE', key, 3600 * 1000) -- 刷新过期时间
  return 0
end
`

// Lua脚本详细解析：
//
// 1. 参数获取与初始化
//    - key = KEYS[1]：获取Redis中的键名，这是一个Hash结构，用于存储漏桶的状态
//    - leak = tonumber(ARGV[1])：获取漏桶的漏出速率，即每秒可以处理的请求数
//    - capacity = tonumber(ARGV[2])：获取漏桶的容量，即桶中最多可以容纳的请求数
//    - n = tonumber(ARGV[3])：获取本次请求的数量
//    - now = tonumber(ARGV[4])：获取当前时间戳，单位为毫秒，由客户端传入以保证时间一致性
//
// 2. 获取存储的状态
//    - level = tonumber(redis.call('HGET', key, 'level'))：获取漏桶当前的水位（已累积的请求数）
//    - last = tonumber(redis.call('HGET', key, 'last_ms'))：获取上次更新水位的时间戳
//    - if level == nil then level = 0 end：如果水位不存在，则初始化为0
//    - if last == nil then last = now end：如果上次更新时间不存在，则初始化为当前时间
//
// 3. 计算漏出量并更新水位
//    - elapsed = math.max(0, (now - last) / 1000.0)：计算自上次更新以来经过的时间（秒），防止负值
//    - level = math.max(0, level - elapsed * leak)：根据经过的时间和漏出速率，计算漏出的请求数，并更新水位
//      确保水位不会低于0
//    - last = now：更新上次更新时间为当前时间
//
// 4. 判断请求是否允许
//    - if level + n <= capacity then：判断本次请求加入后是否会超过漏桶容量
//      - level = level + n：如果未超过，则将本次请求数量加入水位
//      - redis.call('HSET', key, 'level', level, 'last_ms', last)：更新Redis中漏桶的水位和上次更新时间
//      - redis.call('PEXPIRE', key, 3600 * 1000)：设置key的过期时间为1小时，防止key永久存在
//      - return 1：表示请求被允许
//    - else：如果本次请求加入后会超过漏桶容量
//      - redis.call('HSET', key, 'level', level, 'last_ms', last)：即使请求被拒绝，也更新上次更新时间，
//        并保持当前水位（因为请求未被接受，水位未增加）
//      - redis.call('PEXPIRE', key, 3600 * 1000)：刷新key的过期时间
//      - return 0：表示请求被拒绝
//
// 返回值说明：
// - 返回1表示请求被允许
// - 返回0表示请求被拒绝（漏桶已满）
//
// 性能与资源考量：
// 1. 使用Hash结构存储漏桶状态，减少Redis key的数量。
// 2. 每次操作都会更新key的过期时间，确保不活跃的限流器数据会被自动清理，避免内存泄漏。
// 3. Lua脚本保证了操作的原子性，避免了并发条件下的数据不一致问题。
// 4. 漏桶算法通过平滑请求速率来防止突发流量，提供更稳定的服务。
