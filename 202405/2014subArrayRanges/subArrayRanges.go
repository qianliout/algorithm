package main

func main() {

}

func subArrayRanges(nums []int) int64 {
	ans := 0
	for i, ch := range nums {
		mi, ma := ch, ch
		for j := i + 1; j < len(nums); j++ {
			mi = min(mi, nums[j])
			ma = max(ma, nums[j])
			ans += ma - mi
		}
	}
	return int64(ans)
}
