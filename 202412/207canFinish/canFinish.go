package main

func main() {

}

func canFinish(n int, pre [][]int) bool {
	in := make([]int, n)
	out := make([][]int, n)
	q := make([]int, 0)
	for _, ch := range pre {
		x, y := ch[0], ch[1]
		in[x]++
		out[y] = append(out[y], x)
	}
	for i, ch := range in {
		if ch == 0 {
			q = append(q, i)
		}
	}
	cnt := 0
	for len(q) > 0 {
		fir := q[0]
		cnt++
		q = q[1:]
		for _, j := range out[fir] {
			in[j]--
			if in[j] == 0 {
				q = append(q, j)
			}
		}
	}
	return cnt == n
}

// 其中 prerequisites[i] = [ai, bi] ，表示如果要学习课程 ai 则 必须 先学习课程  bi 。

func findOrder(n int, pre [][]int) []int {
	ans := make([]int, 0)
	in := make([]int, n)
	out := make([][]int, n)
	q := make([]int, 0)
	for _, ch := range pre {
		x, y := ch[0], ch[1]
		in[x]++
		out[y] = append(out[y], x)
	}
	for i, ch := range in {
		if ch == 0 {
			q = append(q, i)
		}
	}
	for len(q) > 0 {
		fir := q[0]
		ans = append(ans, fir)
		q = q[1:]
		for _, j := range out[fir] {
			in[j]--
			if in[j] == 0 {
				q = append(q, j)
			}
		}
	}
	if len(ans) == n {
		return ans
	}
	return []int{}
}
