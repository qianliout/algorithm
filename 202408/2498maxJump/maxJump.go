package main

func main() {

}

func maxJump(stones []int) int {
	ans := stones[1] - stones[0]
	n := len(stones)
	for i := 2; i < n; i++ {
		ans = max(ans, stones[i]-stones[i-2])
	}
	return ans
}
