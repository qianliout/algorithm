package main

func main() {

}

func loudAndRich(richer [][]int, quiet []int) []int {
	n := len(quiet)
	g := make([][]int, n)
	in := make([]int, n)
	for _, ch := range richer {
		// x 比y更有钱
		x, y := ch[0], ch[1]
		// g[x] 存的
		g[x] = append(g[x], y)
		in[y]++
	}
	ans := make([]int, n)
	for i := range ans {
		ans[i] = i
	}
	queue := make([]int, 0)
	for i, v := range in {
		if v == 0 {
			queue = append(queue, i)
		}
	}
	// 拓扑排序了
	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]
		// p 比q 有钱
		for _, q := range g[p] {
			if quiet[ans[p]] < quiet[ans[q]] {
				ans[q] = ans[p]
			}
			in[q]--
			if in[q] == 0 {
				queue = append(queue, q)
			}
		}
	}
	return ans
}
