package main

func main() {

}

func equalPairs(grid [][]int) int {
	cnt := make(map[row]int)
	ans := 0
	for _, ch := range grid {
		a := row{}
		for j, v := range ch {
			a[j] = v
		}
		cnt[a]++
	}
	for i := 0; i < len(grid[0]); i++ {
		a := row{}
		for j := 0; j < len(grid); j++ {
			a[j] = grid[j][i]
		}
		ans += cnt[a]
	}
	return ans
}

type row [200]int

// 1 <= n <= 200
