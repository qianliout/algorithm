package main

func main() {}
func maxProfit(prices []int) int {
	n := len(prices)
	has := make([]int, n)
	notHas := make([]int, n)
	has[0] = -prices[0]
	notHas[0] = 0
	for i := 1; i < n; i++ {
		has[i] = max(has[i-1], notHas[i-1]-prices[i])
		notHas[i] = max(notHas[i-1], has[i-1]+prices[i])
	}
	return max(has[n-1], notHas[n-1])
}
