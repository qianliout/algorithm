package main

func main() {

}

func maxSum(nums []int, m int, k int) int64 {
	le, ri, n := 0, 0, len(nums)
	ans := 0
	cnt := make(map[int]int)
	win := 0
	sum := 0
	for le <= ri && ri < n {
		if cnt[nums[ri]] == 0 {
			win++
		}
		cnt[nums[ri]]++
		sum += nums[ri]
		ri++

		if ri-le >= k && win >= m {
			ans = max(ans, sum)
		}
		if ri-le >= k {
			cnt[nums[le]]--
			sum -= nums[le]
			if cnt[nums[le]] == 0 {
				win--
			}
			le++
		}
	}
	return int64(ans)
}

// 会超时
func check(aa map[int]int, m int) bool {
	for _, v := range aa {
		if v == 0 {
			continue
		}
		m--
	}
	return m <= 0
}
