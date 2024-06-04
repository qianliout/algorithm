package main

func main() {

}

func findChampion(n int, edges [][]int) int {
	g := make(map[int]bool)
	ans := -1
	for _, ch := range edges {
		g[ch[1]] = true
	}
	for k := 0; k < n; k++ {
		if !g[k] {
			if ans != -1 {
				return -1
			}
			ans = k
		}
	}
	return ans
}
