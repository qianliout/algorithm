package tokenbucket

import (
	"context"
	"fmt"
	"time"

	redis "github.com/redis/go-redis/v9"
)

// Manager：基于 Redis 的令牌桶限流器。
// 每个 key 的状态存入 Redis Hash：tokens、last_ms。使用 Lua 保证并发下的原子性。
type Manager struct {
	client     *redis.Client
	ratePerSec float64
	burst      int
	keyPrefix  string
	script     *redis.Script
}

// NewManager 创建 Redis 版令牌桶管理器。
// keyPrefix：键名前缀（命名空间），如 "rl:tb:"。
func NewManager(client *redis.Client, ratePerSec float64, burst int, keyPrefix string) *Manager {
	if keyPrefix == "" {
		keyPrefix = "rl:tb:"
	}
	s := redis.NewScript(tokenBucketLua)
	return &Manager{client: client, ratePerSec: ratePerSec, burst: burst, keyPrefix: keyPrefix, script: s}
}

// Allow 为指定 key 消耗 1 个令牌。
func (m *Manager) Allow(ctx context.Context, key string) (bool, error) { return m.AllowN(ctx, key, 1) }

// AllowN 为指定 key 一次性消耗 n 个令牌。
func (m *Manager) AllowN(ctx context.Context, key string, n int) (bool, error) {
	// 1. 参数校验，n<=0 直接允许。
	if n <= 0 {
		return true, nil
	}
	// 2. 获取当前毫秒时间。
	nowMs := time.Now().UnixMilli()
	// 3. 构造 Redis 的 key（带前缀，支持多租户/多用户隔离）。
	redisKey := m.keyPrefix + key
	// 4. 执行 Lua 脚本，原子补充令牌并尝试消耗 n 个令牌。
	// 参数依次为：key、速率、桶容量、消耗数量、当前时间。
	// 这里设置过期时间为 2 倍桶容量（秒），可根据业务调整
	ttl := int64(m.burst * 2)
	res, err := m.script.Run(ctx, m.client, []string{redisKey}, m.ratePerSec, m.burst, n, nowMs, ttl).Int()
	if err != nil {
		return false, err // Redis 或 Lua 执行异常
	}
	// 5. 脚本返回 1 表示允许，0 表示拒绝。
	if res == 1 {
		return true, nil
	}
	return false, nil
}

// tokenBucketLua：Redis Lua 脚本，原子地补充并尝试消耗 N 个令牌。
// 主要逻辑如下：
// 1. 读取当前桶的令牌数和最后更新时间。
// 2. 若首次访问，则令牌数设为桶容量，时间设为当前。
// 3. 否则按时间差补充令牌（速率 × 时间），但不超过桶容量。
// 4. 若令牌足够则消耗 n 个并允许，否则拒绝。
// 5. 最后写回令牌数和时间。
// 参数说明：
//
//	KEYS[1]：Redis Hash key，字段 tokens（浮点）、last_ms（整型）
//	ARGV[1]：ratePerSec（令牌生成速率，浮点）
//	ARGV[2]：burst（桶容量，整型）
//	ARGV[3]：n（本次消耗令牌数，整型）
//	ARGV[4]：now_ms（当前时间，毫秒）
//	ARGV[5]：ttl（key 过期时间，秒）
//
// 返回：允许=1，拒绝=0
const tokenBucketLua = `
-- 取参数
local key = KEYS[1]
local rate = tonumber(ARGV[1])
local burst = tonumber(ARGV[2])
local n = tonumber(ARGV[3])
local now = tonumber(ARGV[4])
local ttl = tonumber(ARGV[5])

-- 读取当前令牌数和最后更新时间
local tokens = tonumber(redis.call('HGET', key, 'tokens'))
local last = tonumber(redis.call('HGET', key, 'last_ms'))
if tokens == nil then
  -- 首次访问，桶满
  tokens = burst
  last = now
else
  if last == nil then last = now end
  -- 计算距离上次的时间差（秒）
  local elapsed = math.max(0, (now - last) / 1000.0)
  -- 按速率补充令牌
  local refill = elapsed * rate
  tokens = math.min(burst, tokens + refill) -- 不超过桶容量
  last = now -- 更新时间
end

-- 判断令牌是否足够
if tokens >= n then
  tokens = tokens - n -- 消耗令牌
  redis.call('HSET', key, 'tokens', tokens, 'last_ms', last)
  redis.call('EXPIRE', key, ttl) -- 设置 key 过期时间，自动清理无流量数据
  return 1 -- 允许
else
  redis.call('HSET', key, 'tokens', tokens, 'last_ms', last)
  redis.call('EXPIRE', key, ttl)
  return 0 -- 拒绝
end
`

// DebugString：返回 key 的可读状态（非原子，仅用于诊断）。
func (m *Manager) DebugString(ctx context.Context, key string) (string, error) {
	h, err := m.client.HGetAll(ctx, m.keyPrefix+key).Result()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", h), nil
}
