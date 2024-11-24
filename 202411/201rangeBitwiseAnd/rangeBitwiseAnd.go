package main

func main() {

}

func rangeBitwiseAnd(left int, right int) int {
	if left > right {
		left, right = right, left
	}
	move := 0
	for left != right {
		left = left >> 1
		right = right >> 1
		move++
	}
	return left << move
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	ans := make([]int, 0)
	for x != 0 {
		ans = append(ans, x%10)
		x = x / 10
	}
	l, r := 0, len(ans)-1
	for l < r {
		if ans[l] != ans[r] {
			return false
		}
		l++
		r--
	}
	return true
}
