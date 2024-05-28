package main

import (
	"fmt"
)

func main() {
	fmt.Println(countVisitedNodes([]int{1, 2, 0, 0}))    // 3,3,3,4
	fmt.Println(countVisitedNodes([]int{1, 2, 3, 4, 0})) //
}

// 直接暴力会超时,加上缓存还是会超时
func countVisitedNodes1(edges []int) []int {
	n := len(edges)
	g := make([][]int, n)
	for i, ch := range edges {
		g[i] = append(g[i], ch)
	}

	ans := make([]int, n)
	mem := make(map[int]int, n)
	for i := range ans {
		ans[i] = bfs(g, i, n, mem)
	}
	return ans
}

func bfs(g [][]int, start, n int, mem map[int]int) int {
	if va, ok := mem[start]; ok {
		return va
	}
	ans := 0
	queue := make([]int, 0)
	queue = append(queue, start)
	visit := make([]bool, n)
	for len(queue) > 0 {
		ans++
		x := queue[0]
		visit[x] = true
		queue = queue[1:]
		nex := g[x] // 这里也说明，上面必须统计好出边
		for _, y := range nex {
			if visit[y] {
				continue
			}
			queue = append(queue, y)
		}
	}
	mem[start] = ans
	return ans
}

/*
对于本题来说：

	对于在基环上的点，其可以访问到的节点数，就是基环的大小。
	对于不在基环上的点 xxx，其可以访问到的节点数，是基环的大小，再加上点 xxx 的深度。这里的深度是指以基环上的点 root\textit{root}root 为根的树枝作为一棵树，点 xxx 在这棵树中的深度。这可以从 root\textit{root}root 出发，在反图上 DFS 得到。

注意题目给出的图可能不是连通的，可能有多棵内向基环树。
*/
func countVisitedNodes2(edges []int) []int {
	n := len(edges)
	g := make([][]int, n)
	in := make([]int, n)
	for x, y := range edges {
		g[y] = append(g[y], x) // 反图
		in[y]++
	}

	// 先把所有的枝找出来，入度是0
	queue := make([]int, 0)
	for x, ch := range in {
		if ch == 0 {
			queue = append(queue, x)
		}
	}

	for len(queue) > 0 {
		x := queue[0]
		queue = queue[1:]
		// 这里取y 的值容易出错,容易写成下面这种写法
		// for _,y:=range g[x]{}
		y := edges[x]
		in[y]--
		if in[y] == 0 {
			queue = append(queue, y)
		}
	}

	ans := make([]int, n)
	for i, d := range in {
		// 说明是树枝或已遍历过的点
		if d <= 0 {
			continue
		}
		// 只有环上的节点入度才大于0
		ring := make([]int, 0)
		x := i
		for {
			in[x] = -1 // 将基环上的点的入度标记为 -1，避免重复访问
			ring = append(ring, x)
			// 继续找下一个
			x = edges[x]
			// 说明回到了起点，这个环遍历完成
			if x == i {
				break
			}
		}
		// 找到一个环了，就把和这个环相连的树枝计算了，
		// 因为可能存在多个环，如果等全部找完再计算树枝的话会出错
		for _, ch := range ring {
			findDep(g, in, ch, len(ring), ans)
		}
	}

	return ans
}

func findDep(g [][]int, in []int, x int, dep int, ans []int) {
	ans[x] = dep
	for _, y := range g[x] {
		if in[y] == 0 { // 树枝上的点在拓扑排序后，入度均为 0
			findDep(g, in, y, dep+1, ans)
		}
	}
}

/*
对于本题来说：

	对于在基环上的点，其可以访问到的节点数，就是基环的大小。
	对于不在基环上的点 xxx，其可以访问到的节点数，是基环的大小，再加上点 xxx 的深度。这里的深度是指以基环上的点 root\textit{root}root 为根的树枝作为一棵树，点 xxx 在这棵树中的深度。这可以从 root\textit{root}root 出发，在反图上 DFS 得到。

注意题目给出的图可能不是连通的，可能有多棵内向基环树。
*/
func countVisitedNodes(edges []int) []int {
	n := len(edges)
	rg := make([][]int, n)
	g := make([][]int, n)
	in := make([]int, n)
	for x, y := range edges {
		rg[y] = append(rg[y], x) // 反图
		g[x] = append(g[x], y)   // 正图，按方向走
		in[y]++
	}

	// 先把所有的枝找出来，入度是0
	queue := make([]int, 0)
	for x, ch := range in {
		if ch == 0 {
			queue = append(queue, x)
		}
	}

	for len(queue) > 0 {
		x := queue[0]
		queue = queue[1:]
		// 这里取y 的值容易出错,因为上面 rg 存的是反图，所以不能用 rg,只能用 g，其实 g和 edges 的一致的，所以可以不用建正图
		for _, y := range g[x] {
			in[y]--
			if in[y] == 0 {
				queue = append(queue, y)
			}
		}
		// 下面这种写法写是对的
		// y := edges[x]
		// in[y]--
		// if in[y] == 0 {
		// 	queue = append(queue, y)
		// }
	}

	ans := make([]int, n)
	// visit := make([]bool, n)
	for i, d := range in {
		// 说明是树枝或已遍历过的点
		if d <= 0 {
			// 这种写法也是可以的，但是效率不高，因为同样一个环中每一个节点都要去计算一次
			// if d <= 0 || visit[i] {
			continue
		}
		// 这种写法也是可以的，但是效率不高，因为同样一个环中每一个节点都要去计算一次环大小，这是没有必要的
		// visit[i] = true
		// 只有环上的节点入度才大于0
		ring := make([]int, 0)
		x := i
		for {
			in[x] = -1 // 将基环上的点的入度标记为 -1，避免重复访问,这样的写法，一个环内就只会计算一次
			ring = append(ring, x)
			// 继续找下一个
			x = edges[x]
			// 说明回到了起点，这个环遍历完成
			if x == i {
				break
			}
		}
		// 找到一个环了，就把和这个环相连的树枝计算了，
		// 因为可能存在多个环，如果等全部找完再计算树枝的话会出错
		for _, ch := range ring {
			findDep(rg, in, ch, len(ring), ans)
		}
	}

	return ans
}
