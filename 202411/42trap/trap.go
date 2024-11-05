package main

func main() {

}

func trap(height []int) int {
	n := len(height)
	lm, rm, l, r := height[0], height[n-1], 0, n-1
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
