package main

func main() {

}

func minOperations(boxes string) []int {
	n := len(boxes)
	nums := make([]int, n+1)
	for i, ch := range boxes {
		nums[i] = int(ch) - int('0')
	}

	ans := make([]int, n)
	for i := 0; i < n; i++ {
		ans[i] = cal(nums, i)
	}
	return ans
}

func cal(nums []int, i int) int {
	ans := 0
	for k := i - 1; k >= 0; k-- {
		ans += (i - k) * nums[k]
	}
	for k := i + 1; k < len(nums); k++ {
		ans += (k - i) * nums[k]
	}
	return ans
}
