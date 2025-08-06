package main

func main() {

}

func trap2(height []int) int {
	n := len(height)
	if n == 0 {
		return 0
	}
	lm, rm := make([]int, n), make([]int, n)
	lm[0], rm[n-1] = height[0], height[n-1]
	for i := 1; i < n; i++ {
		lm[i] = max(lm[i-1], height[i])
	}
	for i := n - 2; i >= 0; i-- {
		rm[i] = max(rm[i+1], height[i])
	}
	ans := 0
	for i := 0; i < n; i++ {
		ans += max(0, min(lm[i], rm[i])-height[i])
	}

	return ans
}

func trap(height []int) int {
	n := len(height)
	if n == 0 {
		return 0
	}
	lm, rm := 0, 0
	l, r := 0, n-1

	ans := 0
	for l <= r {
		lm = max(lm, height[l])
		rm = max(rm, height[r])
		if lm <= rm {
			ans += max(0, lm-height[l])
			l++
		} else {
			ans += max(0, rm-height[r])
			r--
		}
	}

	return ans
}
