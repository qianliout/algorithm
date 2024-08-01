package main

func main() {

}

func platesBetweenCandles(s string, queries [][]int) []int {
	n, m := len(s), len(queries)
	sum := make([]int, n+1)
	pivot := make([]int, 0)
	for i, ch := range s {
		if ch == '|' {
			pivot = append(pivot, i)
		}
		sum[i+1] = sum[i]
		if ch == '*' {
			sum[i+1] += 1
		}
	}
	ans := make([]int, m)

	for i, ch := range queries {
		start, end := ch[0], ch[1]
		// 查左边的蜡烛
		le, ri := start, end+1
		for le < ri {
			mid := le + (ri-le)/2
			if pivot[mid] >= start {

			}

		}
		ans[i] = 0
	}
	return ans
}
