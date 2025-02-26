package main

func main() {

}

// 这样加缓存是错的，不知道是啥原因
func wordPuzzle2(grid [][]byte, target string) bool {
	m, n := len(grid), len(grid[0])
	tt := []byte(target)
	used := make([][]bool, m)
	for i := range used {
		used[i] = make([]bool, n)
	}
	mem := make([][][]int, m)

	var dfs func(i, j int, start int) bool
	dfs = func(i, j int, start int) bool {
		if start >= len(tt) {
			return true
		}
		if i < 0 || j < 0 || i >= m || j >= n {
			return false
		}
		if used[i][j] {
			return false
		}
		if grid[i][j] != tt[start] {
			return false
		}
		if mem[i][j][start] == 1 {
			return true
		} else if mem[i][j][start] == 2 {
			return false
		}

		used[i][j] = true
		a := dfs(i+1, j, start+1)
		b := dfs(i, j+1, start+1)
		c := dfs(i-1, j, start+1)
		d := dfs(i, j-1, start+1)
		used[i][j] = false
		if a || b || c || d {
			mem[i][j][start] = 1
			return true
		}
		mem[i][j][start] = 2
		return false
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			mem = cache(m, n, len(target))
			if dfs(i, j, 0) {
				return true
			}
		}
	}
	return false
}

func cache(m, n, k int) [][][]int {
	mem := make([][][]int, m)
	for i := range mem {
		mem[i] = make([][]int, n)
		for j := range mem[i] {
			mem[i][j] = make([]int, k+1)
		}
	}
	return mem
}

// 这样写会超时
func wordPuzzle1(grid [][]byte, target string) bool {
	m, n := len(grid), len(grid[0])
	tt := []byte(target)
	used := make([][]bool, m)
	for i := range used {
		used[i] = make([]bool, n)
	}
	var dfs func(i, j int, start int) bool
	dfs = func(i, j int, start int) bool {
		if start >= len(tt) {
			return true
		}
		if i < 0 || j < 0 || i >= m || j >= n {
			return false
		}
		if used[i][j] {
			return false
		}
		if grid[i][j] != tt[start] {
			return false
		}
		used[i][j] = true
		a := dfs(i+1, j, start+1)
		b := dfs(i, j+1, start+1)
		c := dfs(i-1, j, start+1)
		d := dfs(i, j-1, start+1)
		used[i][j] = false
		if a || b || c || d {
			return true
		}
		return false
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if dfs(i, j, 0) {
				return true
			}
		}
	}
	return false
}

func wordPuzzle3(grid [][]byte, target string) bool {
	m, n, k := len(grid), len(grid[0]), len(target)
	used := make([][]bool, m)
	tt := []byte(target)
	for i := range used {
		used[i] = make([]bool, n)
	}

	var dfs func(i, j int, start int) bool
	dfs = func(i, j int, start int) bool {
		if start >= k {
			return true
		}
		if i < 0 || j < 0 || i >= m || j >= n {
			return false
		}
		if used[i][j] {
			return false
		}
		if grid[i][j] != tt[start] {
			return false
		}
		used[i][j] = true
		// 这是导致超时的原因，不管下面的a,b,c,d 结果如何，都是要把这四个方向全部 dfs
		a := dfs(i+1, j, start+1)
		b := dfs(i, j+1, start+1)
		c := dfs(i-1, j, start+1)
		d := dfs(i, j-1, start+1)
		used[i][j] = false
		if a || b || c || d {
			return true
		}
		return false
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if dfs(i, j, 0) {
				return true
			}
		}
	}
	return false
}

func wordPuzzle(grid [][]byte, target string) bool {
	m, n, k := len(grid), len(grid[0]), len(target)
	used := make([][]bool, m)
	tt := []byte(target)
	for i := range used {
		used[i] = make([]bool, n)
	}

	var dfs func(i, j int, start int) bool
	dfs = func(i, j int, start int) bool {
		if start >= k {
			return true
		}
		if i < 0 || j < 0 || i >= m || j >= n {
			return false
		}
		if used[i][j] {
			return false
		}
		if grid[i][j] != tt[start] {
			return false
		}
		used[i][j] = true
		// 这是导致超时的原因，不管下面的a,b,c,d 结果如何，都是要把这四个方向全部 dfs
		// a := dfs(i+1, j, start+1)
		// b := dfs(i, j+1, start+1)
		// c := dfs(i-1, j, start+1)
		// d := dfs(i, j-1, start+1)
		// used[i][j] = false
		// if a || b || c || d {
		// 	return true
		// }
		a := dfs(i+1, j, start+1)
		if !a {
			a = dfs(i, j+1, start+1)
		}
		if !a {
			a = dfs(i-1, j, start+1)
		}
		if !a {
			a = dfs(i, j-1, start+1)
		}
		used[i][j] = false
		return a
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if dfs(i, j, 0) {
				return true
			}
		}
	}
	return false
}
