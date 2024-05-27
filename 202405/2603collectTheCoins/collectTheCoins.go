package main

import (
	"fmt"
)

func main() {
	// fmt.Println(collectTheCoins([]int{1, 0, 0, 0, 0, 1}, [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 4}, {4, 5}}))
	fmt.Println(collectTheCoins([]int{1, 0, 1, 1, 1, 0, 1, 0, 0, 0, 0, 1}, [][]int{{0, 1}, {1, 2}, {2, 3}, {2, 4}, {2, 5}, {2, 6}, {4, 7}, {6, 8}, {5, 9}, {4, 10}, {6, 11}}))
}

// 给你一个 n 个节点的无向无根树
func collectTheCoins(coins []int, edges [][]int) int {
	n := len(coins)
	g := make([][]int, n)
	in := make([]int, n)
	for _, ch := range edges {
		x, y := ch[0], ch[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
		in[x]++
		in[y]++
	}
	// 把没有金币的页子节点去了
	leftNode := n // 表示剩余的节点数
	queue := make([]int, 0)
	for k, v := range in {
		// 无向无根树,说明每个节点至少有一个入度,不然这个节点就是一个孤树
		if v == 1 && coins[k] == 0 {
			queue = append(queue, k)
		}
	}
	for len(queue) > 0 {
		fir := queue[0]
		leftNode--
		queue = queue[1:]
		for _, nx := range g[fir] {
			in[nx]--
			// 删除下一叶子节点之后，nx 变成了新的叶子节点
			if in[nx] == 1 && coins[nx] == 0 {
				queue = append(queue, nx)
			}
		}
	}
	/*
		由于可以「收集距离当前节点距离为 2 以内的所有金币」，我们没有必要移动到叶子节点再收集，而是移动到叶子节点的父节点的父节点，就能收集到叶子节点上的金币。
		那么，去掉所有叶子，然后再去掉新产生的叶子，剩余节点就是必须要访问的节点。
	*/

	// 这种写法会报错，这种写法和下面的写法有啥不一致的呢
	// 这种写法会出错，原因是：我们需要t先用所有有金币的叶子节点先去更新其增节点的出度
	// 下面的这种写法，会导致在计算一个节点的父节点时，其没有被全部更新，会导致出错，
	// 但是我没有想到具体是那种场景会成这样
	// for k, v := range in {
	// 	if v == 1 && coins[k] == 1 {
	// 		queue = append(queue, k)
	// 		// 再向上找到父节点
	// 		for _, nx := range g[k] {
	// 			in[nx]--
	// 			if in[nx] == 1 {
	// 				queue = append(queue, nx)
	// 			}
	// 		}
	// 	}
	// }

	for k, v := range in {
		if v == 1 && coins[k] == 1 {
			queue = append(queue, k)
		}
	}

	for _, k := range queue {
		for _, nx := range g[k] {
			in[nx]--
			if in[nx] == 1 {
				queue = append(queue, nx)
			}
		}
	}

	leftNode = leftNode - len(queue)

	return max((leftNode-1)*2, 0)
}
