# DPDS - Dynamic Priority DAG Scheduler

一个基于 Go 实现的动态优先级 DAG 调度器，支持任务优先级调度、依赖管理、循环检测和优雅关闭。

## 功能特性

- **优先级调度**：严格按 P0 > P1 > P2 > P3 > P4 顺序调度（数值越小优先级越高）
- **DAG 依赖管理**：支持任务间依赖关系的动态添加和移除
- **循环依赖检测**：添加依赖时自动检测环
- **并发控制**：限制同时运行的任务数量
- **依赖感知调度**：只有前置依赖全部完成的任务才会被调度
- **任务状态追踪**：支持查询任务当前状态
- **优雅关闭**：支持等待运行中任务完成后关闭调度器

## 项目结构

```
dpds/
├── promt.md              # 需求文档
├── scheduler.go          # 调度器核心
├── task.go              # 任务定义与状态机
├── priority_queue.go     # 可更新优先级堆
├── dependency_graph.go   # 依赖关系图（DAG）
├── dpds_test.go         # 测试用例
└── README.md            # 本文档
```

## 核心组件

### TaskStatus 任务状态机

```
PENDING → READY → RUNNING → COMPLETED
```

| 状态 | 描述 |
|------|------|
| PENDING | 等待依赖完成 |
| READY | 依赖已满足，等待调度 |
| RUNNING | 执行中 |
| COMPLETED | 已完成 |

### PriorityQueue 优先级堆

- 基于二叉堆实现
- O(log n) 插入和删除
- 优先级取反存储，实现最大堆行为（数值越小优先级越高）

### DependencyGraph 依赖图

- 邻接表存储
- 支持动态添加/移除依赖
- 环检测使用 DFS

## 并发模型

- **Channel**：`taskChan` 用于提交任务
- **Atomic**：任务状态和计数器使用 `sync/atomic`
- **Mutex**：仅在优先级堆和依赖图修改时使用

## 调度流程详解（Scheduler）

这一节按“数据结构 → 任务提交流 → 调度循环 → 依赖解锁 → 动态依赖”把 `scheduler.go` 的执行路径串起来，方便对照源码理解。

### 关键数据结构

- `tasks map[string]*Task`：任务注册表（`taskID -> task`），用于：
  - 查询状态（`GetTaskStatus`）
  - 在依赖解锁时检查依赖任务是否已完成
- `depGraph *DependencyGraph`：依赖关系图（DAG）
  - `dependencies[task] = [dep1, dep2...]`：task 依赖谁
  - `dependents[dep] = [task1, task2...]`：谁依赖 dep
- `readyQueue *PriorityQueue`：就绪队列（优先级堆），只存可以立即运行的任务
  - 实现上把 `priority` 存成 `-task.Priority`，从而让 `Priority=0`（最高）先出队
- `maxConcurrent / runningCount`：并发控制
  - `schedule()` 在 `runningCount < maxConcurrent` 时持续出队并启动任务

### 生命周期总览

```
Submit() -> taskChan -> enqueueTask() -> schedule() -> startTask()
                                           |
                                           v
                                      handleTaskDone()
                                           |
                                           v
                           checkAndReady(dependent) -> schedule()
```

### 任务提交：Submit()

`Submit(task)` 做三件事：

1. 校验：task.ID 不能为空
2. 注册：写入 `tasks[task.ID] = task`
3. 建边：对 `task.Dependencies` 循环 `depGraph.AddDependency(task.ID, depID)`
4. 投递：把任务写入 `taskChan`，交给调度协程异步处理

为什么要 `taskChan`：
- 把“入队/触发调度”集中在一个调度循环里处理，让提交方不直接操作队列（更容易收敛并发点）。

### 调度主循环：Start() / schedulerLoop()

- `Start()` 启动一个 goroutine 执行 `schedulerLoop()`
- `schedulerLoop()` 是一个无限循环：
  - 收到 `taskChan` 的新任务 -> `enqueueTask(task)`
  - 收到 `stopChan` -> 退出循环（Shutdown）

### 入队判定：enqueueTask()

`enqueueTask(task)` 的核心是决定“PENDING 还是 READY”：

- 如果 `depGraph.HasDependencies(task.ID)` 为 true：说明存在前置依赖
  - 任务状态设为 `PENDING`
- 否则：没有依赖
  - 任务状态设为 `READY`
  - 推入 `readyQueue`

无论走哪条分支，最后都会调用一次 `schedule()`，尝试尽快启动可运行的任务。

### 调度决策：schedule()

`schedule()` 的逻辑非常直接：只要还有并发槽位，就不断从 `readyQueue` 取任务启动。

- 循环条件：`runningCount < maxConcurrent`
- 从 `readyQueue.Pop()` 取一个 READY 任务
  - 若为空：说明当前没有可运行任务，直接返回
- 调用 `startTask(task)` 启动任务

这也是“优先级调度”的真正发生点：`readyQueue.Pop()` 总是返回当前最高优先级任务。

### 启动执行：startTask()

`startTask(task)` 做三件事：

1. 标记 `RUNNING`
2. `runningCount++`（atomic）并 `waitGroup.Add(1)`
3. 启 goroutine 执行 `TaskFunc`

goroutine 的 defer 里负责收尾：

- `handleTaskDone(taskID)`：更新状态、解锁依赖、触发下一轮调度
- `waitGroup.Done()`：配合 `Shutdown()` 等待所有运行中任务退出

### 任务完成：handleTaskDone()

任务完成后会发生依赖解锁：

1. `runningCount--`
2. 将自己标记为 `COMPLETED`
3. 查 `depGraph.GetDependents(taskID)`：取出所有“依赖我”的任务列表
4. 对每个 dependent 调 `checkAndReady(dependentID)`
5. 再调用一次 `schedule()`：如果有新 READY 的任务，就立即补满并发槽位

### 依赖检查与解锁：checkAndReady()

`checkAndReady(taskID)` 的目标是：把“依赖已满足”的 PENDING 任务推进 READY 队列。

- 如果任务不存在或任务状态不是 `PENDING`：直接返回
- 取依赖列表 `deps := depGraph.GetDependencies(taskID)`
- 遍历 deps：
  - 若任一依赖任务不存在，或依赖任务状态不是 `COMPLETED`：返回（依赖未满足）
- 所有依赖都完成：
  - 任务状态设为 `READY`
  - 推入 `readyQueue`

### 动态依赖：AddDependency / RemoveDependency

#### AddDependency(taskID, dependsOnID)

该 API 用来“运行前动态加依赖”，限制与处理要点：

- 限制：不能给 `RUNNING / COMPLETED` 的任务加依赖（否则语义不一致）
- 加边：`depGraph.AddDependency(taskID, dependsOnID)`（内部会做环检测）
- 若任务当前是 `READY`：
  - 需要把它从 `readyQueue` 移除
  - 状态改回 `PENDING`（因为新依赖可能没完成）
- 然后：
  - `checkAndReady(taskID)`（如果新依赖其实已完成，会立刻回到 READY）
  - `schedule()`（尝试启动）

#### RemoveDependency(taskID, dependsOnID)

- 移除边：`depGraph.RemoveDependency(...)`
- 再次判断是否可以运行：`checkAndReady(taskID)`
- `schedule()` 尝试启动

### 一个最小例子（帮助你建立直觉）

假设 `maxConcurrent=2`，任务依赖关系：

- A：无依赖
- B：依赖 A
- C：无依赖

那么调度一定满足：

1. B 先进入 `PENDING`（因为依赖 A 未完成）
2. A、C 进入 `READY`，被 `schedule()` 启动（最多同时 2 个）
3. A 完成后，`handleTaskDone(A)` 会唤醒 `B`：
   - `checkAndReady(B)` 发现 A 已 `COMPLETED`，将 B 推入 `readyQueue`
4. 随后 `schedule()` 在并发槽位空出来时启动 B

### 实现注意事项

- `readyQueue` 与 `depGraph` 内部做了加锁，但 `tasks map` 在并发读写场景下需要额外保护（如果要在高并发生产环境使用，建议为 `tasks` 增加互斥或把对 `tasks` 的读写也收敛到同一协程/事件循环中）。

## 快速开始

```go
package main

import (
    "fmt"
    "time"
    "github.com/yourrepo/dpds"
)

func main() {
    // 创建调度器，最大并发数为 3
    s := dpds.NewScheduler(3)
    s.Start()

    // 提交任务
    task1 := dpds.NewTask("task-1", 0, 100, nil, func() error {
        fmt.Println("Task 1 running")
        time.Sleep(100 * time.Millisecond)
        return nil
    })

    task2 := dpds.NewTask("task-2", 1, 100, []string{"task-1"}, func() error {
        fmt.Println("Task 2 running after Task 1")
        time.Sleep(100 * time.Millisecond)
        return nil
    })

    s.Submit(task1)
    s.Submit(task2)

    // 等待任务完成
    time.Sleep(500 * time.Millisecond)

    // 查询状态
    status, _ := s.GetTaskStatus("task-1")
    fmt.Printf("Task 1 status: %s\n", status)

    // 关闭调度器
    s.Shutdown()
}
```

## API 文档

### NewScheduler(maxConcurrent int) *Scheduler

创建一个新的调度器实例。

- `maxConcurrent`：最大并发任务数

### func (s *Scheduler) Start()

启动调度器，开始处理任务。

### func (s *Scheduler) Submit(task *Task) error

提交新任务。

- `task`：要提交的任务
- 返回值：如果任务 ID 为空或存在循环依赖，返回错误

### func (s *Scheduler) AddDependency(taskID, dependsOnID string) error

动态添加任务依赖。

- `taskID`：任务 ID
- `dependsOnID`：被依赖的任务 ID
- 返回值：如果任务不存在、正在运行、已结束或添加后会产生循环依赖，返回错误

### func (s *Scheduler) RemoveDependency(taskID, dependsOnID string) error

移除任务依赖。

### func (s *Scheduler) GetTaskStatus(taskID string) (TaskStatus, error)

查询任务状态。

### func (s *Scheduler) Shutdown()

优雅关闭调度器，等待所有运行中的任务完成后返回。

### func (s *Scheduler) GetRunningCount() int32

获取当前运行中的任务数。

### func (s *Scheduler) GetReadyCount() int

获取就绪队列中的任务数。

## NewTask 函数签名

```go
func NewTask(id string, priority int, duration int, dependencies []string, fn func() error) *Task
```

- `id`：任务唯一标识
- `priority`：优先级（0-4，0 最高）
- `duration`：预估执行时长（毫秒）
- `dependencies`：依赖的任务 ID 列表
- `fn`：任务执行函数

## 优先级说明

| 优先级 | 数值 | 说明 |
|--------|------|------|
| P0 | 0 | 最高优先级 |
| P1 | 1 | |
| P2 | 2 | |
| P3 | 3 | |
| P4 | 4 | 最低优先级 |

任务按优先级从高到低调度，同优先级任务按提交顺序调度。

## 依赖管理

### 静态依赖

提交任务时通过 `dependencies` 参数指定：

```go
task := dpds.NewTask("B", 0, 100, []string{"A"}, fn)
// B 依赖 A，A 完成后 B 才能执行
```

### 动态依赖

任务提交后动态添加依赖：

```go
s.AddDependency("C", "A") // C 依赖 A
s.AddDependency("C", "B") // C 依赖 A 和 B
```

### 循环检测

添加依赖时自动检测环：

```go
// A 依赖 B，B 依赖 A -> 错误
s.AddDependency("A", "B")
s.AddDependency("B", "A") // 返回错误：adding dependency would create a cycle
```

### 动态依赖限制

- 只能给 PENDING 或 READY 状态的任务添加依赖
- 不能给 RUNNING 或 COMPLETED 状态的任务添加依赖

## 性能指标

- 单次调度决策：< 1ms
- O(log n) 优先级调整
- O(1) 依赖查询

## 运行测试

```bash
go test -v ./...
```

## 许可

MIT License
