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