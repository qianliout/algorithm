package main

func main() {

}

func minimumScore(nums []int, edges [][]int) int {
	n := len(nums)
	g := make([][]int, n)
	for _, ch := range edges {
		x, y := ch[0], ch[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}
	clock := 0
	in, out := make([]int, n), make([]int, n)
	xor := make([]int, n)
	var dfs func(x, fa int)
	dfs = func(x, fa int) {
		clock++
		in[x] = clock
		xor[x] = nums[x]
		for _, y := range g[x] {
			if y == fa {
				continue
			}
			dfs(y, x)
			xor[x] ^= xor[y]
		}
		out[x] = clock
	}
	dfs(0, -1)
	// 开始分类讨论
	ans := 1 << 31

	for i := 2; i < n; i++ {
		for j := 1; j < i; j++ {
			x, y, z := 0, 0, 0
			// 用这种方法来判断一个树是另一个树的子树
			if in[i] < in[j] && in[j] <= out[i] {
				x, y, z = xor[j], xor[i]^xor[j], xor[0]^xor[i]
			} else if in[j] < in[i] && in[i] <= out[j] {
				x, y, z = xor[i], xor[j]^xor[i], xor[0]^xor[j]
			} else {
				x, y, z = xor[i], xor[j], xor[0]^xor[i]^xor[j]
			}
			ans = min(ans, max(x, y, z)-min(x, y, z))
		}
	}
	return ans
}
