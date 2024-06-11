package main

func main() {

}

func longestSubarray(nums []int, limit int) int {
	n := len(nums)
	mi := make([]int, 0) // 递增队列维护最小值，mi[0]是最小值
	mx := make([]int, 0) // 递减队列维护最大值，mx[0]是最大值
	le, ri, ans := 0, 0, 0
	for le <= ri && ri < n {
		c := nums[ri]
		ri++
		for len(mx) > 0 && mx[len(mx)-1] < c {
			mx = mx[:len(mx)-1]
		}
		mx = append(mx, c)
		for len(mi) > 0 && mi[len(mi)-1] > c {
			mi = mi[:len(mi)-1]
		}
		mi = append(mi, c)

		for len(mx) > 0 && len(mi) > 0 && (mx[0]-mi[0] > limit) {
			if mx[0] == nums[le] {
				mx = mx[1:]
			} else if mi[0] == nums[le] {
				mi = mi[1:]
			}
			le++
		}
		ans = max(ans, ri-le)
	}
	return ans
}
