package main

func main() {

}

// 1 <= nums[i] <= 100 没有负数
func minDifference(nums []int, queries [][]int) []int {
	c := 100
	n := len(nums)
	sum := make([][]int, n+1)
	for i := range sum {
		sum[i] = make([]int, c+1) // 范围中包括100
	}
	for i, ch := range nums {
		sum[i+1] = append([]int{}, sum[i]...)
		sum[i+1][ch]++
	}
	ans := make([]int, len(queries))
	for i, ch := range queries {
		le, ri := ch[0], ch[1]
		last := 0 // 值是范围是1-100
		best := c
		for j := 1; j <= c; j++ {
			if sum[ri+1][j] != sum[le][j] {
				if last != 0 {
					best = min(best, j-last)
				}
				last = j
			}
		}
		if best != c {
			ans[i] = best
		} else {
			ans[i] = -1
		}
	}
	return ans
}
