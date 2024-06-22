package main

func main() {

}

func subarrayBitwiseORs(nums []int) int {
	ans := make(map[int]int)
	cur := make(map[int]int)
	cur[0] = 0
	for _, x := range nums {
		th := make(map[int]int)
		th[x]++
		for pre := range cur {
			th[pre|x]++
		}
		for k := range th {
			ans[k]++
		}
		cur = th

	}
	return len(ans)
}

func smallestSubarrays(nums []int) []int {
	n := len(nums)
	ans := make([]int, n)
	for i, x := range nums {
		ans[i] = 1
		for j := i - 1; j >= 0 && nums[j]|x != nums[j]; j-- {
			nums[j] = nums[j] | x
			ans[j] = i - j + 1
		}
	}
	return ans
}
