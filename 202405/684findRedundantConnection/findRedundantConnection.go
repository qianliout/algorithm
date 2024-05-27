package main

func main() {

}

func findRedundantConnection(edges [][]int) []int {
	n := len(edges)
	g := make([][]int, n+10)
	in := make([]int, n+10)
	for _, ch := range edges {
		x, y := ch[0], ch[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
		in[x]++
		in[y]++
	}
	queue := make([]int, 0)
	for i, v := range in {
		if v == 1 {
			queue = append(queue, i)
		}
	}
	for len(queue) > 0 {
		x := queue[0]
		queue = queue[1:]
		for _, nx := range g[x] {
			in[nx]--
			if in[nx] == 1 {
				queue = append(queue, nx)
			}
		}
	}
	for i := len(edges) - 1; i >= 0; i-- {
		ch := edges[i]
		if in[ch[0]] > 1 && in[ch[1]] > 1 {
			return ch
		}

	}
	return []int{}
}
