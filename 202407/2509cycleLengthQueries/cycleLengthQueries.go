package main

func main() {

}

func cycleLengthQueries(n int, queries [][]int) []int {
	ans := make([]int, len(queries))
	for i, ch := range queries {
		a, b := ch[0], ch[1]
		cnt := 1
		for a != b {
			if a > b {
				a = a >> 1
			} else {
				b = b >> 1
			}
			cnt++
		}
		ans[i] = cnt
	}
	return ans
}
