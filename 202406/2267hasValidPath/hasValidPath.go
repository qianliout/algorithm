package main

func main() {

}

func hasValidPath(grid [][]byte) bool {
	m, n := len(grid), len(grid[0])
	if (m+n)%2 == 0 || grid[0][0] == ')' || grid[m-1][n-1] == '(' { // 剪枝

		return false

	}

	var dfs func(i, j, path int) bool
	mem := make([][][]int, m)
	for i := range mem {
		mem[i] = make([][]int, n)
		for j := range mem[i] {
			// mem[i][j] = make([]int, m*n) // 如果这样写就会超内存
			mem[i][j] = make([]int, (m+n+1)/2)
			for k := range mem[i][j] {
				mem[i][j][k] = -1
			}
		}
	}

	dfs = func(i, j, path int) bool {
		if path > m-i+n-j-1 { // 剪枝：即使后面都是 ')' 也不能将 c 减为 0
			return false
		}
		// 代码实现时，由于找到合法路径就返回 true 了，不会继续执行 dfs，若 dfs(x,y,c) 最后返回的是 false，
		// 那后续访问同一个状态时（再次调用 dfs(x,y,c)），仍然会得到 false。因此没必要重复访问同一个状态，可以用一个 vis 数组标记，
		// 遇到访问过的状态可以直接返回 false
		// 要理解的是这里不只是对一个格子的 i,j 做了缓存，而是i,j 加上状态的缓存，所以如果走到同一个格式，状态不同，缓存也是不同的
		if mem[i][j][path] != -1 {
			// 为啥返问过就返回 false 呢，因为这
			return false
		}
		if i == m-1 && j == n-1 {
			// 最后一个一定是')'
			return path == 1
		}
		mem[i][j][path] = 1
		if grid[i][j] == '(' {
			path++
		} else if grid[i][j] == ')' {
			path--
		}

		if path < 0 {
			return false
		}
		// 这里一定要判断 i+1<m,这样减少递归
		a := (i+1 < m && dfs(i+1, j, path)) || (j+1 < n && dfs(i, j+1, path))

		return a
	}

	return dfs(0, 0, 0)
}

func in(m, n, i, j int) bool {
	if i < 0 || j < 0 {
		return false
	}
	if i >= m || j >= n {
		return false
	}

	return true
}
