package main

func main() {

}

func countMaxOrSubsets(nums []int) int {
	mx := 0
	for _, ch := range nums {
		mx = mx | ch
	}
	n := len(nums)
	m := 1 << n
	ans := 0
	for i := 0; i < m; i++ {
		if cal(i, nums, mx) {
			ans++
		}
	}
	return ans
}

func cal(stat int, nums []int, mx int) bool {
	ans := 0
	for i := 0; i < len(nums); i++ {
		if stat&(1<<i) != 0 {
			ans = ans | nums[i]
		}
	}
	return ans == mx
}
