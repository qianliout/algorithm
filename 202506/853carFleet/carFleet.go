package main

import (
	"fmt"
	"sort"
)

func main() {
	// 测试用例1：经典例子
	target1 := 12
	position1 := []int{10, 8, 0, 5, 3}
	speed1 := []int{2, 4, 1, 1, 3}

	fmt.Printf("测试1:\n")
	fmt.Printf("目标位置: %d\n", target1)
	fmt.Printf("车辆位置: %v\n", position1)
	fmt.Printf("车辆速度: %v\n", speed1)
	result1 := carFleet(target1, position1, speed1)
	fmt.Printf("车队数量: %d\n\n", result1)

	// 详细分析过程
	fmt.Println("=== 详细分析过程 ===")
	carFleetDetailed(target1, position1, speed1)
}

func carFleet(target int, position []int, speed []int) int {
	/*
		核心思想：使用单调栈

		关键洞察：
		1. 从离终点最近的车开始分析（从右到左）
		2. 计算每辆车到达终点的时间
		3. 如果后面的车到达时间 >= 前面的车，说明会追上形成车队
		4. 使用单调栈维护独立车队的到达时间
	*/

	// 步骤1：将位置和速度组合，便于排序
	cars := make([]pair, 0)
	for i := range position {
		cars = append(cars, pair{pos: position[i], speed: speed[i]})
	}

	// 步骤2：按位置排序
	sort.Slice(cars, func(i, j int) bool { return cars[i].pos <= cars[j].pos })

	// 步骤3：计算每辆车到达终点的时间
	n := len(position)
	times := make([]float64, n)
	for i := range cars {
		times[i] = float64(target-cars[i].pos) / float64(cars[i].speed)
	}

	// 步骤4：使用单调栈维护车队
	stack := make([]float64, 0)

	for _, time := range times {
		/*
			单调栈的核心逻辑：
			- 如果当前车的到达时间 >= 栈顶车的时间，说明会追上
			- 追上后形成车队，以较慢的车（时间更长）为准
			- 因此要移除栈顶的较快车
		*/
		for len(stack) > 0 && time >= stack[len(stack)-1] {
			stack = stack[:len(stack)-1] // 移除被追上的车
		}
		stack = append(stack, time) // 当前车成为新的独立车队
	}

	return len(stack) // 栈的长度就是车队数量
}

// 详细分析函数 - 展示算法执行过程
func carFleetDetailed(target int, position []int, speed []int) {
	fmt.Printf("目标位置: %d\n", target)
	fmt.Printf("初始状态:\n")

	// 步骤1：组合数据
	cars := make([]pair, 0)
	for i := range position {
		cars = append(cars, pair{pos: position[i], speed: speed[i]})
		fmt.Printf("车%d: 位置=%d, 速度=%d\n", i, position[i], speed[i])
	}
	fmt.Println()

	// 步骤2：排序（从离终点近到远）
	sort.Slice(cars, func(i, j int) bool {
		return cars[i].pos > cars[j].pos
	})

	fmt.Println("按位置排序后（从近到远）:")
	for i, car := range cars {
		fmt.Printf("第%d辆: 位置=%d, 速度=%d\n", i+1, car.pos, car.speed)
	}
	fmt.Println()

	// 步骤3：计算到达时间
	fmt.Println("计算到达时间:")
	times := make([]float64, len(cars))
	for i, car := range cars {
		times[i] = float64(target-car.pos) / float64(car.speed)
		fmt.Printf("车(位置%d): 需要时间 = (%d-%d)/%d = %.2f\n",
			car.pos, target, car.pos, car.speed, times[i])
	}
	fmt.Println()

	// 步骤4：单调栈模拟
	fmt.Println("单调栈处理过程:")
	stack := make([]float64, 0)

	for i, time := range times {
		fmt.Printf("处理车%d (位置%d, 时间%.2f):\n", i+1, cars[i].pos, time)
		fmt.Printf("  当前栈: %v\n", formatStack(stack))

		// 检查是否会追上前面的车
		removed := 0
		for len(stack) > 0 && time >= stack[len(stack)-1] {
			removedTime := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			removed++
			fmt.Printf("  -> 会追上前车(时间%.2f)，移除\n", removedTime)
		}

		stack = append(stack, time)
		fmt.Printf("  -> 加入栈，当前栈: %v\n", formatStack(stack))
		fmt.Printf("  -> 当前车队数: %d\n\n", len(stack))
	}

	fmt.Printf("最终车队数: %d\n", len(stack))
}

// 格式化栈显示
func formatStack(stack []float64) []string {
	result := make([]string, len(stack))
	for i, v := range stack {
		result[i] = fmt.Sprintf("%.2f", v)
	}
	return result
}

type pair struct {
	pos   int // 位置
	speed int // 速度
}

// 在一条单行道上，有 n 辆车开往同一目的地。目的地是几英里以外的 target 。
// 给定两个整数数组 position 和 speed ，长度都是 n ，其中 position[i] 是第 i 辆车的位置， speed[i] 是第 i 辆车的速度(单位是英里/小时)。
// 一辆车永远不会超过前面的另一辆车，但它可以追上去，并以较慢车的速度在另一辆车旁边行驶。
// 车队 是指并排行驶的一辆或几辆汽车。车队的速度是车队中 最慢 的车的速度。
// 即便一辆车在 target 才赶上了一个车队，它们仍然会被视作是同一个车队。
// 返回到达目的地的车队数量 。
