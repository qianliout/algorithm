package main

func main() {

}

// 直接算会超时
func countBadPairs(nums []int) int64 {
	n := len(nums)
	d := make([]int, n)
	for i, ch := range nums {
		d[i] = ch - i
	}

	ans := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if d[i] != d[j] {
				ans++
			}
		}
	}
	return int64(ans)
}
