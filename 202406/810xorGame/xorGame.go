package main

func main() {

}
func xorGame(nums []int) bool {
	if len(nums)%2 == 0 {
		return true
	}
	x := 0
	for _, ch := range nums {
		x = x ^ ch
	}
	return x == 0
}
