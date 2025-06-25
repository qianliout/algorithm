package main

func main() {

}

func minOperations(nums []int) int {
	g, n, one := 0, len(nums), 0
	for _, ch := range nums {
		if ch == 1 {
			one++
		}
		g = gcd(g, ch)
	}
	if g > 1 {
		return -1
	}
	if one > 0 {
		return n - one
	}
	minCnt := n
	for i := 0; i < n; i++ {
		g := 0
		for j := i; j < n; j++ {
			g = gcd(g, nums[j])
			if g == 1 {
				minCnt = min(minCnt, j-i)
			}
		}
	}
	return minCnt + n - 1
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
