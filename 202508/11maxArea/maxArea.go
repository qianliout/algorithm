package main

func main() {

}

func maxArea(height []int) int {

	n := len(height)
	if n == 0 {
		return 0
	}

	ans := 0
	l, r := 0, n-1
	for l < r {
		mx := min(height[l], height[r])

		ans = max(ans, mx*(r-l))
		if height[l] > mx {
			r--
		} else {
			l++
		}
	}
	return ans
}
