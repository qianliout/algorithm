package main

func main() {}

func trap(height []int) int {
	n := len(height)
	lm, rm, l, r := 0, 0, 0, n-1
	ans := 0
	for l < r {
		lm = max(lm, height[l])
		rm = max(rm, height[r])
		if lm <= rm {
			ans += max(0, min(lm, rm)-height[l])
			l++
		} else {
			ans += max(0, min(lm, rm)-height[r])
			r--
		}
	}
	return ans
}
