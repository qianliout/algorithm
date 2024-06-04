package main

func main() {

}

func maximumScoreAfterOperations(edges [][]int, values []int) int64 {
	n := len(values)
	g := make([][]int, n)

	g[0] = append(g[0], -1) // 认为根节点的父节点是一个虚拟节点（小技巧）

	for _, ch := range edges {
		x, y := ch[0], ch[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}
	all := 0
	for _, ch := range values {
		all += ch
	}
	var dfs func(x int, fa int) int
	// 先假设全部选中，然后做 dfs
	// dfs的返回值是选和不选 x 节点时，所损失的点权
	// 也就是说：计算以 x 为根的子树是健康时，失去的最小分数
	dfs = func(x int, fa int) int {
		// 因为只有当父节点不选择时才会向下递归，如果递归到叶子节点了，所以这个叶子节点的所有向上父节点都没有选
		// 走到了叶子节点了
		// len(g[x])==1 表示此时 x 只和父节点相临了
		if len(g[x]) == 1 { // 如果上面把根节点加了一个虚拟父节点就不用判断 x!=0,如果不加就必须判断
			// if x != 0 && len(g[x]) == 1 {
			return values[x]
		}

		// 选x节点，那么损失值就是x 的点权，此时不需要向下递归
		// 因为选了根节点，所有的子节点都是健康
		los2 := values[x]

		// 不选x 的节点,那么在假设全选中的前提下会损失x的点权值就是0，
		// 此时就需要向下递归
		los1 := 0
		for _, y := range g[x] {
			if y != fa {
				los1 += dfs(y, x)
			}
		}
		return min(los2, los1)
	}

	return int64(all) - int64(dfs(0, -1))
}
