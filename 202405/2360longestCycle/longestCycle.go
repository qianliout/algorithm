package main

func main() {

}

func longestCycle(edges []int) int {
	n := len(edges)
	ti := make([]int, n)
	// clock 初始化>=0就行
	clock, ans := 0, -1
	for x, ch := range ti {
		if ch > 0 {
			continue
		}
		start := clock
		for x >= 0 {
			if ti[x] > 0 {
				if ti[x] >= start {
					ans = max(ans, clock-ti[x])
				}
				break
			}
			ti[x] = clock
			x = edges[x]
			clock++
		}
	}
	return ans
}
