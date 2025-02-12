package main

func main() {

}

// 这样写是错的，因为题目中说是环形树组
func minSwaps1(nums []int) int {
	cnt := 0
	for _, ch := range nums {
		cnt += ch
	}
	le, ri, n := 0, 0, len(nums)
	mx, wi := 0, 0

	for le <= ri && ri < n {
		wi += nums[ri]
		ri++
		if ri-le == cnt {
			mx = max(mx, wi)
		}
		if ri-le >= cnt {
			wi -= nums[le]
			le++
		}
	}
	return cnt - mx
}
func minSwaps(nums []int) int {
	cnt := 0
	for _, ch := range nums {
		cnt += ch
	}
	le, ri, n := 0, 0, len(nums)
	mx, wi := 0, 0

	for le <= ri && ri < n*2 {
		wi += nums[ri%n]
		ri++
		if ri-le == cnt {
			mx = max(mx, wi)
		}
		if ri-le >= cnt {
			wi -= nums[le%n]
			le++
		}
	}
	return cnt - mx
}
