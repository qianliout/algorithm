# Rate Limiter Implementations

本目录实现了 4 种经典限流算法的单机与 Redis 分布式版本：
- 令牌桶 Token Bucket
- 漏桶 Leaky Bucket
- 固定窗口 Fixed Window Counter
- 滑动窗口 Sliding Window（Log 与 Counter 两种）

目录结构：
- `single/`
  - `tokenbucket/`
  - `leakybucket/`
  - `fixedwindow/`
  - `slidingwindow/`
    - `log/`
    - `counter/`
- `distributed/`
  - `redisclient/` Redis 客户端辅助
  - `tokenbucket/`
  - `leakybucket/`
  - `fixedwindow/`
  - `slidingwindow/`
    - `log/` 使用 ZSET 精确滑动窗口
    - `counter/` 使用 ZSET+HASH 的分桶滑动窗口

注意：`slidingwindow/log` 包名为 `log`，与标准库 `log` 同名，建议 import 时使用别名。

## Quick Start（单机版）

```go
package main

import (
	"fmt"
	"time"
	singleTB "outback/algorithm/202508/RateLimiter/22/single/tokenbucket"
)

func main() {
	mgr := singleTB.NewManager(5 /* tokens/sec */, 10 /* burst */)
	key := "user:123"

	for i := 0; i < 12; i++ {
		ok := mgr.Allow(key)
		fmt.Println(i, ok)
		if i == 6 { time.Sleep(300 * time.Millisecond) }
	}
}
```

## Quick Start（Redis 分布式）

```go
package main

import (
	"context"
	"fmt"
	"time"

	redis "github.com/redis/go-redis/v9"
	dtb "outback/algorithm/202508/RateLimiter/22/distributed/tokenbucket"
)

func main() {
	ctx := context.Background()
	cli := redis.NewClient(&redis.Options{Addr: "127.0.0.1:6379"})
	defer cli.Close()

	mgr := dtb.NewManager(cli, 5 /* tokens/sec */, 10 /* burst */, "rl:tb:")
	key := "user:123"

	for i := 0; i < 12; i++ {
		ok, err := mgr.Allow(ctx, key)
		if err != nil { panic(err) }
		fmt.Println(i, ok)
		if i == 6 { time.Sleep(300 * time.Millisecond) }
	}
}
```

## 其它说明
- Redis 版本均使用 Lua 脚本保证并发原子性。
- 滑动窗口 Log 版本使用 ZSET，Counter 版本使用 ZSET+HASH 分桶；按窗口宽度清理过期数据。
- 单机版本均为并发安全（`sync.Mutex`），按调用时的时间差做懒计算。
- 如需对“每个用户/每个 API”分别限流，使用各自 `Manager`，传入不同 `key` 即可。

与 `ratelimit.md` 笔记对应的实现细节已在各文件内用中文注释说明，可直接阅读源码了解具体算法要点与权衡。


