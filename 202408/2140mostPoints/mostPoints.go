package main

func main() {

}

func mostPoints(questions [][]int) int64 {
	n := len(questions)
	f := make([]int64, n+1)
	for i := 0; i < n; i++ {
		points, brainpower := questions[i][0], questions[i][1]
		f[i+1] = max(f[i+1], f[i]) // 不做这一题
		j := min(n, i+brainpower+1)
		f[j] = max(f[j], f[i]+int64(points)) // 做这一题
	}
	return f[n]
}
