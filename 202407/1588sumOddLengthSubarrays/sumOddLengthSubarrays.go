package main

func main() {

}

func sumOddLengthSubarrays(arr []int) int {
	n := len(arr)
	sum := make([]int, n+1)
	for i, ch := range arr {
		sum[i+1] = sum[i] + ch
	}
	ans := 0
	left := 0
	for left <= n {
		for i := left + 1; i <= n; i += 2 {
			ans += sum[i] - sum[left]
		}
		left++
	}

	return ans
}
