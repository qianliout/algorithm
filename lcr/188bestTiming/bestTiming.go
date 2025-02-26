package main

func main() {

}

func bestTiming(prices []int) int {
	ans := 0
	st := make([]int, 0)
	for _, ch := range prices {
		for len(st) > 0 && st[len(st)-1] >= ch {
			st = st[:len(st)-1]
		}
		if len(st) > 0 {
			ans = max(ans, ch-st[0])
		}
		st = append(st, ch)
	}
	return ans
}
