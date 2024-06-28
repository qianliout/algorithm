package main

func main() {

}

func gardenNoAdj(n int, paths [][]int) []int {
	g := make([][]int, n)
	for _, ch := range paths {
		x, y := ch[0]-1, ch[1]-1
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}
	ans := make([]int, n)
	for x := 0; x < n; x++ {
		used := make(map[int]bool)
		for _, y := range g[x] {
			used[ans[y]] = true
		}
		for i := 1; i < 5; i++ {
			if !used[i] {
				ans[x] = i
				break
			}
		}
	}
	return ans
}
