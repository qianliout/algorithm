package main

func main() {

}

func arraySign(nums []int) int {
	a, b := 0, 0
	for _, ch := range nums {
		if ch == 0 {
			return 0
		}
		if ch > 0 {
			a++
		}
		if ch < 0 {
			b++
		}
	}
	if b%2 == 0 {
		return 1
	}
	return -1
}
