package main

func main() {

}

func reachableNodes(n int, edges [][]int, restricted []int) int {
	g := make([][]int, n)
	for _, ch := range edges {
		x, y := ch[0], ch[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}
	visited := make(map[int]bool)
	for _, ch := range restricted {
		visited[ch] = true
	}
	queue := []int{0}
	visited[0] = true
	cnt := 0
	for len(queue) > 0 {
		cnt += len(queue)
		lev := make([]int, 0)
		for _, no := range queue {
			for _, nx := range g[no] {
				if visited[nx] {
					continue
				}
				visited[nx] = true
				lev = append(lev, nx)
			}
		}
		queue = lev
	}
	return cnt
}
