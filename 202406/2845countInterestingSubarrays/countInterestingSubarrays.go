package main

func main() {

}

func countInterestingSubarrays(nums []int, m int, k int) int64 {
	n := len(nums)
	pre := make([]int, n+1)
	for i, ch := range nums {
		if ch%m == k {
			pre[i+1] = 1
		}
		pre[i+1] += pre[i]
	}
	cnt := make(map[int]int)
	ans := 0
	for _, s := range pre {
		ans += cnt[(s%m-k+m)%m]
		cnt[s%m]++
	}
	return int64(ans)
}
