package main

func main() {

}

func findKDistantIndices(nums []int, key int, k int) []int {
	idx := make(map[int]int)
	for i, ch := range nums {
		if ch == key {
			idx[i] = ch
		}
	}
	ans := make([]int, 0)
	for i := range nums {
		if check(idx, k, i) {
			ans = append(ans, i)
		}
	}
	return ans
}

func check(idx map[int]int, k, i int) bool {
	for j := range idx {
		if abs(j-i) <= k {
			return true
		}
	}
	return false
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
