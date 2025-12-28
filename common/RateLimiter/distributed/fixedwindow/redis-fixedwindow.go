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
local now = tonumber(ARGV[4])  -- 使用传入的时间参数

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
`

// Lua脚本详细解析：
//
// 1. 参数获取与初始化
//    - key = KEYS[1]：获取Redis中的键名，这是一个Hash结构
//    - limit = tonumber(ARGV[1])：获取限流阈值，表示在一个窗口期内允许的最大请求数
//    - window_ms = tonumber(ARGV[2])：获取窗口大小，单位为毫秒
//    - n = tonumber(ARGV[3])：获取本次请求的数量
//    - now = tonumber(ARGV[4])：获取当前时间戳，单位为毫秒，由客户端传入以保证时间一致性
//
// 2. 窗口ID计算
//    - cur_id = math.floor(now / window_ms)：计算当前时间所属的窗口ID
//      窗口ID是通过将当前时间戳除以窗口大小并向下取整得到的，同一窗口期内的所有请求具有相同的窗口ID
//
// 3. 获取存储的状态
//    - stored_id = tonumber(redis.call('HGET', key, 'window_id'))：获取存储的窗口ID
//    - count = tonumber(redis.call('HGET', key, 'count'))：获取当前窗口已处理的请求计数
//    - if count == nil then count = 0 end：如果计数不存在，则初始化为0
//
// 4. 窗口切换逻辑
//    - if stored_id == nil or stored_id ~= cur_id then：判断是否需要切换到新窗口
//      当存储的窗口ID不存在或与当前窗口ID不同时，表示需要切换到新窗口
//      - stored_id = cur_id：更新窗口ID
//      - count = 0：重置计数器
//      - redis.call('HSET', key, 'window_id', cur_id, 'count', n)：设置新窗口的ID和计数
//      - redis.call('PEXPIRE', key, window_ms * 2)：设置过期时间为窗口大小的2倍
//        这确保了即使在没有新请求的情况下，旧数据也会被自动清理，避免内存泄漏
//      - 判断本次请求是否超过限制并返回结果：
//        - if n <= limit then return 1 else return 0 end
//
// 5. 同窗口请求处理
//    - else：如果仍在当前窗口内
//      - if count + n <= limit then：判断添加新请求后是否超过限制
//        - redis.call('HINCRBY', key, 'count', n)：增加计数
//        - redis.call('PEXPIRE', key, window_ms * 2)：刷新过期时间
//        - return 1：表示请求被允许
//      - else：如果超过限制
//        - redis.call('PEXPIRE', key, window_ms * 2)：仍然刷新过期时间，确保限流状态被保留
//        - return 0：表示请求被拒绝
//
// 返回值说明：
// - 返回1表示请求被允许
// - 返回0表示请求被拒绝（超过限流阈值）
//
// 性能与资源考量：
// 1. 使用Hash结构而非String，可以在一个key中存储多个字段，减少key的数量
// 2. 通过设置过期时间，确保不活跃的限流器数据会被自动清理，避免Redis内存持续增长
// 3. 使用Lua脚本保证了操作的原子性，避免了竞态条件
// 4. 过期时间设置为窗口大小的2倍，确保在窗口切换时仍能正确获取上一个窗口的数据
