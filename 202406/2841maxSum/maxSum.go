package main

func main() {

}

func maxSum(nums []int, m int, k int) int64 {
	n := len(nums)
	sum := make([]int, n+1)
	for i, ch := range nums {
		sum[i+1] = sum[i] + ch
	}
	exist := make(map[int]int)
	le, ri, ans := 0, 0, 0

	for le <= ri && ri < n {
		exist[nums[ri]]++
		ri++
		for ri-le > k {
			exist[nums[le]]--
			if exist[nums[le]] == 0 {
				delete(exist, nums[le])
			}
			le++
		}
		if ri-le == k && len(exist) >= m {
			ans = max(ans, sum[ri]-sum[le])
		}
	}

	return int64(ans)
}
