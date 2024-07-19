package main

func main() {

}

func check(nums []int) bool {
	ans := 0
	n := len(nums)
	for i := 0; i < n; i++ {
		if i == n-1 && nums[0] < nums[i] {
			ans++
		}
		if i < n-1 && nums[i+1] < nums[i] {
			ans++
		}
	}
	return ans <= 1
}
