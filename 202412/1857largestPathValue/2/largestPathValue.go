package main

func main() {

}

func largestPathValue(colors string, edges [][]int) int {
	n := len(colors)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, 26)
	}
	in := make([]int, n)
	out := make([][]int, n)
	for _, ch := range edges {
		x, y := ch[0], ch[1]
		out[x] = append(out[x], y)
		in[y]++
	}
	visited := make([]bool, n)
	q := make([]int, 0)
	for i := range in {
		if in[i] == 0 {
			q = append(q, i)
		}
	}
	find := 0
	for len(q) > 0 {
		find++
		x := q[0]
		q = q[1:]
		c := cIdx(colors[x])
		dp[x][c]++
		for _, y := range out[x] {
			if visited[y] {
				continue
			}
			for i := 0; i < 26; i++ {
				dp[y][i] = max(dp[y][i], dp[x][i])
			}
			in[y]--
			if in[y] == 0 {
				q = append(q, y)
			}
		}
	}
	if find != n {
		return -1
	}
	ans := 0

	for i := range dp {
		for _, ch := range dp[i] {
			ans = max(ans, ch)
		}
	}

	return ans
}

func cIdx(c byte) int {
	return int(c) - int('a')
}
