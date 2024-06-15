package main

func main() {

}

func numSubmatrixSumTarget(matrix [][]int, target int) int {
	n, m := len(matrix), len(matrix[0])
	pre := make([][]int, n+1)
	for i := range pre {
		pre[i] = make([]int, m+1)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			pre[i+1][j+1] = pre[i+1][j] + pre[i][j+1] - pre[i][j] + matrix[i][j]
		}
	}
	ans := 0
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			for k := i - 1; k >= 0; k-- {
				for p := j - 1; p >= 0; p-- {
					if pre[i][j]-pre[k][j]-pre[i][p]+pre[k][p] == target {
						ans++
					}
				}
			}
		}
	}

	return ans
}

func maximumPopulation(logs [][]int) int {
	n := 2050 - 1950 + 1
	d := make([]int, n)

	for _, ch := range logs {
		x, y := ch[0]-1950, ch[1]-1950
		d[x]++
		d[y]--
	}
	mx, mxY, sum := 0, 0, 0
	for i, num := range d {
		sum += num
		if sum > mx {
			mx = sum
			mxY = i + 1950
		}
	}
	return mxY
}

func countTestedDevices(batteryPercentages []int) int {

}
