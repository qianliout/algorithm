package main

func main() {

}

func maxProfit(prices []int) int {
	n := len(prices)
	has := make([]int, n+10)
	not := make([]int, n+10)
	has[0] = -1 << 30

	for i := 0; i < n; i++ {
		has[i+1] = max(has[i], not[i]-prices[i])
		not[i+1] = max(not[i], has[i]+prices[i])
	}

	return not[n]
}
