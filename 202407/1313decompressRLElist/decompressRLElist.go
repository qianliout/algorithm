package main

func main() {

}

func decompressRLElist(nums []int) []int {
	ans := make([]int, 0)
	n := len(nums)
	for i := 0; i+1 < n; i = i + 2 {
		nu := nums[i+1]
		fre := nums[i]
		for j := 0; j < fre; j++ {
			ans = append(ans, nu)
		}
	}
	return ans
}
