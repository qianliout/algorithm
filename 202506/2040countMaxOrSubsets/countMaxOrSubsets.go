package main

func main() {

}

func countMaxOrSubsets(nums []int) int {
	n := len(nums)
	m := 1 << n
	mx := 0
	for _, c := range nums {
		mx = mx | c
	}
	ans := 0
	for i := 0; i < m; i++ {
		if getMax(i, nums) == mx {
			ans++
		}
	}
	return ans
}

func getMax(stat int, nums []int) int {
	mx := 0
	for i := range nums {
		if (stat>>i)&1 == 1 {
			mx = mx | nums[i]
		}
	}
	return mx
}
