package main

func main() {

}

func findOrder(numCourses int, prerequisites [][]int) []int {
	g := make([][]int, numCourses)
	in := make([]int, numCourses)
	for _, ch := range prerequisites {
		x, y := ch[0], ch[1]
		g[y] = append(g[y], x)
		in[x]++
	}
	q := make([]int, 0)
	for i, c := range in {
		if c == 0 {
			q = append(q, i)
		}
	}
	ans := make([]int, 0)
	cnt := 0
	for len(q) > 0 {
		fir := q[0]
		cnt++
		ans = append(ans, fir)
		q = q[1:]
		for _, x := range g[fir] {
			in[x]--
			if in[x] == 0 {
				q = append(q, x)
			}
		}
	}
	if cnt == numCourses {
		return ans
	}
	return []int{}
}
