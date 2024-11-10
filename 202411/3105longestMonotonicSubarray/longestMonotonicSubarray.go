package main

func main() {

}

func longestMonotonicSubarray1(nums []int) int {
	n := len(nums)
	ans := 1
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			// 升序
			if nums[j] >= nums[j-1] {
				break
			}
			ans = max(ans, j-i+1)
		}
		for j := i + 1; j < n; j++ {
			if nums[j] <= nums[j-1] {
				break
			}
			ans = max(ans, j-i+1)
		}
	}
	return ans
}

func longestMonotonicSubarray(nums []int) int {
	n, ans, i := len(nums), 1, 0
	for i < n-1 {
		if nums[i+1] == nums[i] {
			i++
			continue
		}
		i0 := i
		inc := nums[i+1] > nums[i] // 用一个变量判断了上升和下降
		i += 2
		for i < n && nums[i] != nums[i-1] && (nums[i] > nums[i-1]) == inc {
			i++
		}
		ans = max(ans, i-i0)
		i--
	}
	return ans
}
