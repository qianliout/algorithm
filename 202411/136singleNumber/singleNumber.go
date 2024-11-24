package main

func main() {

}

func singleNumber(nums []int) int {
	ans := 0
	for _, ch := range nums {
		ans = ans ^ ch
	}
	return ans
}
