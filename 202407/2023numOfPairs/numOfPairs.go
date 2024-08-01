package main

func main() {

}

func numOfPairs(nums []string, target string) int {
	// 两个字符串连接
	ans := 0
	n := len(nums)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j {
				continue
			}
			if nums[i]+nums[j] == target {
				ans++
			}
		}
	}
	return ans
}
