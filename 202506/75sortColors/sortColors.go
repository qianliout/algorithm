package main

func main() {

}

func sortColors(nums []int) {
	n := len(nums)
	red, blue := 0, n-1
	i := 0
	for i <= blue {
		if nums[i] == 2 {
			nums[i], nums[blue] = nums[blue], nums[i]
			blue--
		} else if nums[i] == 0 {
			nums[i], nums[red] = nums[red], nums[i]
			red++
			i++
		} else {
			i++
		}
	}
}

// 我们使用整数 0、 1 和 2 分别表示红色、白色和蓝色。
// 必须在不使用库内置的 sort 函数的情况下解决这个问题。
