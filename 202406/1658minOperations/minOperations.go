package main

func main() {

}

func minOperations(nums []int, x int) int {
	sum, n := 0, len(nums)
	for _, ch := range nums {
		sum += ch
	}
	if sum < x {
		return -1
	}
	x = sum - x
	add := 0
	le, ri, ans := 0, 0, -1
	for le <= ri && ri < n {
		c := nums[ri]
		add += c
		ri++
		for add > x {
			add -= nums[le]
			le++
		}

		if add == x {
			ans = max(ans, ri-le)
		}
	}
	if ans == -1 {
		return -1
	}
	return n - ans
}
