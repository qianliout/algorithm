package main

func main() {

}

func numOfSubarrays(arr []int, k int, threshold int) int {
	ans := 0
	n := len(arr)
	sum := make([]int, n+1)
	for i, ch := range arr {
		sum[i+1] = sum[i] + ch
	}

	for ri := 0; ri+k <= n; ri++ {
		sub := sum[ri+k] - sum[ri]
		if sub/k >= threshold {
			ans++
		}
	}

	return ans
}
