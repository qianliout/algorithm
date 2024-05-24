package main

func main() {

}

func findShortestSubArray(nums []int) int {
	exit := make(map[int]int)
	for _, ch := range nums {
		exit[ch]++
	}
	mf := 0
	a := make([]int, 0)
	for _, v := range exit {
		mf = max(mf, v)
	}
	for k, v := range exit {
		if v == mf {
			a = append(a, k)
		}
	}
	ans := len(nums)
	for _, ch := range a {
		ans = min(ans, find(nums, ch, exit[ch]))
	}
	return ans
}

func find(nums []int, va int, k int) int {
	start := -1
	for i, ch := range nums {
		if ch == va {
			if start < 0 {
				start = i
			}
			k--
			if k == 0 {
				return i - start + 1
			}
		}
	}
	return 0
}
