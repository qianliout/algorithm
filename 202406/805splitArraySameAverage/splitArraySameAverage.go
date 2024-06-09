package main

func main() {

}

func splitArraySameAverage(nums []int) bool {
	all, n := 0, len(nums)
	for _, ch := range nums {
		all += ch
	}
	if all%n != 0 || all%2 != 0 {
		return false
	}
	// 未完成
	return false
}
