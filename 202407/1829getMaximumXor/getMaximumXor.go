package main

func main() {

}

func getMaximumXor(nums []int, maximumBit int) []int {
	sumOr := 0
	for _, ch := range nums {
		sumOr = sumOr ^ ch
	}
	mask := 1<<maximumBit - 1 // 是一个各个位都是1的数

	n := len(nums)
	ans := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		ans[n-1-i] = sumOr ^ mask
		sumOr = sumOr ^ nums[i]
	}
	return ans
}
