package main

func main() {

}

func trap(height []int) int {
	l, r, n := 0, len(height)-1, len(height)
	lm, rm := height[0], height[n-1]
	ans := 0
	for l < r {
		lm = max(lm, height[l])
		rm = max(rm, height[r])
		if lm < rm {
			ans += lm - height[l]
			l++
		} else {
			ans += rm - height[r]
			r--
		}
	}
	return ans
}
