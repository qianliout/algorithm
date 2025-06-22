package main

import "fmt"

func main() {
	// 测试用例1：经典例子
	graph1 := [][]int{{1, 2}, {3}, {3}, {}}
	fmt.Printf("测试1: graph = %v\n", graph1)
	fmt.Println("图的结构:")
	printGraph(graph1)
	result1 := allPathsSourceTarget(graph1)
	fmt.Printf("所有路径: %v\n\n", result1)

	// 测试用例2：更复杂的例子
	graph2 := [][]int{{4, 3, 1}, {3, 2, 4}, {3}, {4}, {}}
	fmt.Printf("测试2: graph = %v\n", graph2)
	fmt.Println("图的结构:")
	printGraph(graph2)
	result2 := allPathsSourceTarget(graph2)
	fmt.Printf("所有路径: %v\n\n", result2)

	// 详细分析过程
	fmt.Println("=== 详细DFS过程分析 ===")
	allPathsSourceTargetDetailed(graph1)
}

// 打印图的结构
func printGraph(graph [][]int) {
	for i, neighbors := range graph {
		if len(neighbors) == 0 {
			fmt.Printf("节点%d -> 无出边\n", i)
		} else {
			fmt.Printf("节点%d -> %v\n", i, neighbors)
		}
	}
}

// ✅ 正确的解法：从起点0到终点n-1的DFS
func allPathsSourceTarget(graph [][]int) [][]int {
	/*
		🎯 问题分析：
		- 有向无环图(DAG)：保证不会有环，DFS不会无限循环
		- 起点：节点0
		- 终点：节点n-1
		- 目标：找出所有从0到n-1的路径

		🧠 解题思路：
		1. 使用DFS深度优先搜索
		2. 从节点0开始，沿着有向边探索
		3. 当到达节点n-1时，记录当前路径
		4. 使用回溯法探索所有可能的路径

		⚡ 关键技巧：
		- 路径记录：使用数组记录当前路径
		- 回溯：递归返回时移除当前节点
		- 终止条件：到达目标节点n-1
	*/

	n := len(graph)
	result := make([][]int, 0)

	// DFS函数：从当前节点开始搜索
	var dfs func(node int, path []int)
	dfs = func(node int, path []int) {
		// 将当前节点加入路径
		path = append(path, node)

		// 🎯 终止条件：到达目标节点n-1
		if node == n-1 {
			// 找到一条完整路径，复制并保存
			// 注意：必须复制，因为path会被后续修改
			pathCopy := make([]int, len(path))
			copy(pathCopy, path)
			result = append(result, pathCopy)
			return
		}

		// 🔄 递归探索：遍历当前节点的所有邻居
		for _, neighbor := range graph[node] {
			dfs(neighbor, path)
		}

		// 🔙 回溯：移除当前节点，为其他路径让路
		// 注意：这里不需要显式回溯，因为path是值传递
	}

	// 从节点0开始DFS
	dfs(0, []int{})

	return result
}

// 🔍 详细分析版本：展示DFS的完整过程
func allPathsSourceTargetDetailed(graph [][]int) {
	fmt.Printf("图结构分析: %v\n", graph)
	n := len(graph)
	result := make([][]int, 0)
	depth := 0

	var dfs func(node int, path []int)
	dfs = func(node int, path []int) {
		// 缩进显示递归深度
		indent := ""
		for i := 0; i < depth; i++ {
			indent += "  "
		}

		path = append(path, node)
		fmt.Printf("%s进入节点%d，当前路径: %v\n", indent, node, path)

		// 到达目标节点
		if node == n-1 {
			pathCopy := make([]int, len(path))
			copy(pathCopy, path)
			result = append(result, pathCopy)
			fmt.Printf("%s✅ 找到完整路径: %v\n", indent, pathCopy)
			return
		}

		// 探索邻居节点
		if len(graph[node]) == 0 {
			fmt.Printf("%s❌ 节点%d无出边，回溯\n", indent, node)
			return
		}

		fmt.Printf("%s🔍 节点%d的邻居: %v\n", indent, node, graph[node])
		for i, neighbor := range graph[node] {
			fmt.Printf("%s📍 探索邻居%d (%d/%d)\n", indent, neighbor, i+1, len(graph[node]))
			depth++
			dfs(neighbor, path)
			depth--
			fmt.Printf("%s🔙 从邻居%d返回\n", indent, neighbor)
		}
	}

	fmt.Println("\n=== DFS搜索过程 ===")
	dfs(0, []int{})
	fmt.Printf("\n最终结果: %v\n", result)
}

// 🚀 优化版本：使用引用传递和显式回溯
func allPathsSourceTargetOptimized(graph [][]int) [][]int {
	/*
		💡 优化思路：
		- 使用引用传递避免频繁的数组复制
		- 显式回溯，手动管理路径状态
		- 减少内存分配，提高性能
	*/

	n := len(graph)
	result := make([][]int, 0)
	path := make([]int, 0) // 使用共享的路径数组

	var dfs func(node int)
	dfs = func(node int) {
		// 将当前节点加入路径
		path = append(path, node)

		if node == n-1 {
			// 到达目标，复制当前路径
			pathCopy := make([]int, len(path))
			copy(pathCopy, path)
			result = append(result, pathCopy)
		} else {
			// 继续探索邻居
			for _, neighbor := range graph[node] {
				dfs(neighbor)
			}
		}

		// 🔙 显式回溯：移除当前节点
		path = path[:len(path)-1]
	}

	dfs(0)
	return result
}

func reverse(path []int) []int {
	ans := make([]int, 0)
	for i := len(path) - 1; i >= 0; i-- {
		ans = append(ans, path[i])
	}
	return ans
}

/*
给你一个有 n 个节点的 有向无环图（DAG），请你找出从节点 0 到节点 n-1 的所有路径并输出（不要求按特定顺序）
graph[i] 是一个从节点 i 可以访问的所有节点的列表（即从节点 i 到节点 graph[i][j]存在一条有向边）。
*/
