package main

func main() {

}

func edgeScore(edges []int) int {
	n := len(edges)
	g := make([]int, n)
	for i := 0; i < n; i++ {
		g[edges[i]] += i
	}
	s, a := g[0], 0
	for i := 0; i < n; i++ {
		if g[i] > s {
			s = g[i]
			a = i
		}
	}
	return a
}
