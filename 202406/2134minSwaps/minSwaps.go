package main

func main() {

}

func minSwaps(nums []int) int {
	sum := 0
	for _, ch := range nums {
		sum += ch // 只有0和1两个数
	}
	n := len(nums)
	le, ri := 0, 0
	wind := 0
	ans := 0
	for le <= ri && ri < 2*n {
		idx := ri % n
		wind += nums[idx]
		ri++
		if ri-le > sum {
			idx2 := le % n
			wind -= nums[idx2]
			le++
		}
		if ri-le == sum {
			ans = max(ans, wind)
		}
	}

	return sum - ans
}
