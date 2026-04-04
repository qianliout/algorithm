package main

func main() {}

func trap_1(height []int) int {
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

func trap(height []int) int {
	n := len(height)
	st := make([]int, 0)
	ans := 0
	for i := 0; i < n; i++ {
		right := height[i]
		for len(st) > 0 && right > height[st[len(st)-1]] {
			bottom := height[st[len(st)-1]]
			st = st[:len(st)-1]
			if len(st) == 0 {
				break
			}
			left := height[st[len(st)-1]]
			width := i - st[len(st)-1] - 1
			h := min(left, right) - bottom
			ans += max(0, h*width)
		}
		st = append(st, i)
	}
	return ans
}
