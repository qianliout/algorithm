package main

func main() {

}

func findRedundantConnection(edges [][]int) []int {
	n := len(edges)
	in := make([]int, n)
	g := make([][]int, n)
	for _, ch := range edges {
		x, y := ch[0]-1, ch[1]-1
		in[x]++
		in[y]++
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}
	q := make([]int, 0)
	for k, v := range in {
		if v == 1 {
			q = append(q, k)
		}
	}
	for len(q) > 0 {
		fir := q[0]
		q = q[1:]
		for _, y := range g[fir] {
			in[y]--
			if in[y] == 1 {
				q = append(q, y)
			}
		}
	}
	for i := n - 1; i >= 0; i-- {
		x, y := edges[i][0]-1, edges[i][1]-1
		if in[x] > 0 && in[y] > 0 {
			return []int{x + 1, y + 1}
		}
	}
	return []int{}
}
