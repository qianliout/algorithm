package main

func main() {

}

func countPrimeSetBits(left int, right int) int {
	ans := 0

	ss := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97}
	mm := make(map[int]bool)
	for _, ch := range ss {
		mm[ch] = true
	}
	for i := left; i <= right; i++ {
		if mm[ha(i)] {
			ans++
		}
	}
	return ans
}

func ha(n int) int {
	ans := 0
	for n > 0 {
		if n&1 == 1 {
			ans++
		}
		n = n >> 1
	}
	return ans
}
