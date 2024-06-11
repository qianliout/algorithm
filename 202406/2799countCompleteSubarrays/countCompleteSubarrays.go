package main

func main() {

}

func countCompleteSubarrays(nums []int) int {
	n, ans := len(nums), 0
	le, ri := 0, 0
	cnt := make(map[int]int)
	for _, ch := range nums {
		cnt[ch]++
	}
	exist := make(map[int]int)
	for le <= ri && ri < n {
		c := nums[ri]
		exist[c]++
		ri++
		for len(exist) >= len(cnt) {
			x := nums[le]
			exist[x]--
			if exist[x] == 0 {
				delete(exist, x)
			}
			le++
		}
		ans += le
	}
	return ans
}
