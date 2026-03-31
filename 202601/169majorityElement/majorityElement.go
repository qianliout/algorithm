package main

func main() {}

func majorityElement(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	ans := nums[0]
	cnt := 1
	n := len(nums)
	for i := 1; i < n; i++ {
		if nums[i] == ans {
			cnt++
		} else {
			cnt--
			if cnt == 0 {
				ans = nums[i]
				cnt = 1
			}
		}
	}
	return ans
}
