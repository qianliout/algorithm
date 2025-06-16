package main

func main() {

}

func binaryGap(n int) int {
	ans := make([]int, 0)
	for n > 0 {
		ans = append(ans, n&1)
		n = n >> 1
	}
	pre := -1
	res := 0
	for i, c := range ans {
		if c == 1 {
			if pre != -1 {
				res = max(res, i-pre)
			}
			pre = i
		}
	}
	return res
}
