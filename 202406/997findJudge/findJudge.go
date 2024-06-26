package main

func main() {

}

func findJudge(n int, trust [][]int) int {
	in := make([]int, n+1)
	out := make([]int, n+1)
	for _, ch := range trust {
		x, y := ch[0]-1, ch[1]-1
		in[y]++
		out[x]++
	}
	for i := 0; i < n; i++ {
		if out[i] == 0 && in[i] == n-1 {
			return i + 1
		}
	}
	return -1
}
