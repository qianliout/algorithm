package counter

import (
	"context"
	"time"

	redis "github.com/redis/go-redis/v9"
)

// Manager implements a Redis-backed Sliding Window Counter with bucketization.
// Data structures per key:
// - ZSET: scores = bucketStartMs, members = bucketStartMs (as string) for ordering and pruning
// - HASH: field = bucketStartMs string, value = count for that bucket
type Manager struct {
	client      *redis.Client
	limit       int
	windowSize  time.Duration
	bucketWidth time.Duration
	keyPrefix   string
	script      *redis.Script
}

// NewManager constructs a counter-based sliding window using buckets.
// bucketWidth controls the granularity; e.g., window=60s, bucketWidth=10s -> 6 buckets.
func NewManager(client *redis.Client, limit int, windowSize time.Duration, bucketWidth time.Duration, keyPrefix string) *Manager {
	if keyPrefix == "" {
		keyPrefix = "rl:swc:"
	}
	return &Manager{client: client, limit: limit, windowSize: windowSize, bucketWidth: bucketWidth, keyPrefix: keyPrefix, script: redis.NewScript(luaSlidingCounter)}
}

// Allow accepts one event.
func (m *Manager) Allow(ctx context.Context, key string) (bool, error) { return m.AllowN(ctx, key, 1) }

// AllowN accepts n events if capacity allows.
func (m *Manager) AllowN(ctx context.Context, key string, n int) (bool, error) {
	if n <= 0 {
		return true, nil
	}
	nowMs := time.Now().UnixMilli()
	res, err := m.script.Run(
		ctx,
		m.client,
		[]string{m.keyPrefix + key + ":z", m.keyPrefix + key + ":h"},
		m.limit,
		m.windowSize.Milliseconds(),
		m.bucketWidth.Milliseconds(),
		n,
		nowMs,
	).Int()
	if err != nil {
		return false, err
	}
	return res == 1, nil
}

// KEYS[1] = zset, KEYS[2] = hash
// ARGV: limit, window_ms, bucket_ms, n, now_ms
const luaSlidingCounter = `
local zkey = KEYS[1]
local hkey = KEYS[2]
local limit = tonumber(ARGV[1])
local window_ms = tonumber(ARGV[2])
local bucket_ms = tonumber(ARGV[3])
local n = tonumber(ARGV[4])
local now = tonumber(ARGV[5])

local threshold = now - window_ms

-- prune old buckets
local oldBuckets = redis.call('ZRANGEBYSCORE', zkey, '-inf', threshold)
if #oldBuckets > 0 then
  for i=1,#oldBuckets do
    redis.call('HDEL', hkey, oldBuckets[i])
  end
  redis.call('ZREMRANGEBYSCORE', zkey, '-inf', threshold)
end

-- sum active buckets
local activeBuckets = redis.call('ZRANGEBYSCORE', zkey, threshold, '+inf')
local total = 0
for i=1,#activeBuckets do
  local c = tonumber(redis.call('HGET', hkey, activeBuckets[i])) or 0
  total = total + c
end

if total + n <= limit then
  local curBucket = tostring(math.floor(now / bucket_ms) * bucket_ms)
  if redis.call('ZSCORE', zkey, curBucket) == false then
    redis.call('ZADD', zkey, tonumber(curBucket), curBucket)
  end
  redis.call('HINCRBY', hkey, curBucket, n)
  return 1
else
  return 0
end
`
