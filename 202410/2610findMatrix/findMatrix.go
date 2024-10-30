package main

func main() {

}

func findMatrix(nums []int) [][]int {
	exit := make(map[int]int)
	mx := 0
	for _, ch := range nums {
		exit[ch]++
		mx = max(mx, exit[ch])
	}
	ans := make([][]int, mx)
	for k, v := range exit {
		for i := 0; i < v; i++ {
			ans[i] = append(ans[i], k)
		}
	}
	return ans
}
