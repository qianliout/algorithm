package main

func main() {

}

// 图中的最长环
func longestCycle2(edges []int) int {
	/*
		核心思想：使用时间戳检测环

		关键概念：
		- ti[x]: 节点x的访问时间戳
		- start: 当前路径的起始时间戳
		- clock: 全局时间戳计数器

		为什么需要 ti[x] >= start 判断？
		这是区分"当前路径中的环"和"之前路径访问过的节点"的关键！
	*/

	n := len(edges)
	ans := -1
	ti := make([]int, n) // 时间戳数组，0表示未访问
	clock := 1           // 全局时间戳，从1开始（0表示未访问）

	for i := 0; i < n; i++ {
		if ti[i] > 0 {
			// 节点已经被访问过，跳过
			continue
		}

		start := clock // 记录当前路径的起始时间戳
		x := i

		// 沿着边走，直到遇到已访问节点或到达终点
		for x != -1 && ti[x] == 0 {
			ti[x] = clock // 标记访问时间
			clock++
			x = edges[x] // 移动到下一个节点
		}

		// 如果遇到已访问的节点
		if x != -1 && ti[x] >= start {
			/*
				关键判断：ti[x] >= start

				情况1：ti[x] >= start
				- 说明节点x是在当前路径中访问的
				- 这意味着找到了一个环！
				- 环的长度 = clock - ti[x]

				情况2：ti[x] < start
				- 说明节点x是在之前的路径中访问的
				- 当前路径只是连接到了之前的路径，没有形成新环
				- 不需要更新答案
			*/
			ans = max(ans, clock-ti[x])
		}
	}

	return ans
}

// 给你一个 n 个节点的 有向图 ，节点编号为 0 到 n - 1 ，
// 其中每个节点 至多 有一条出边。这个条件很重要，说明最多一个环，且是内向环。

/*
初始时间为 curTime=1。遍历图，每访问到一个新的节点 x，就记录首次访问时间 visTime[x]=curTime，然后将 curTime 加一。
假设我们从节点 i 开始。首先记录开始时间 startTime=curTime，然后继续走，如果走到死路，或者找到了一个之前访问过的点 x，则退出循环。
退出循环后，分类讨论：
如果 visTime[x]<startTime，说明 x 不是在本轮循环中访问的。例如上图从节点 0 开始，访问节点 0,3,2,4。然后接着从节点 1 开始，访问节点 3，
发现 visTime[3] 比访问节点 1 的时间还要早，那么包含节点 3 的环长我们之前已经计算过了，无需再次计算。
如果 visTime[x]≥startTime，说明 x 是在本轮循环中访问的，且被访问了两次。这只有一种可能，就是 x 在环上。根据前后两次访问 x 的时间差，就能算出环长，
即 curTime−visTime[x]。
注：本题保证每个连通块至多有一个环，所以可以根据时间差算出环长。如果没有这个保证，时间差算出的可能不是最长环。一般图的最长环是 NP-hard 问题。
*/

// 图中的最长环
func longestCycle(edges []int) int {
	ans, n, clock := -1, len(edges), 1
	tim := make([]int, n)
	for i := 0; i < n; i++ {
		if tim[i] > 0 {
			continue
		}
		x := i
		start := clock
		for x >= 0 {
			if tim[x] > 0 {
				if tim[x] >= start {
					// 说明x在当前路径中访问过
					ans = max(ans, clock-tim[x])
				}
				break // 找到已访问节点，退出循环
			}
			tim[x] = clock // 记录访问时间
			clock++
			x = edges[x] // 移动到下一个节点
		}
	}
	return ans
}
