package main

func main() {

}

func maxProfit(prices []int, fee int) int {
	n := len(prices)
	f1 := make([]int, n) // has
	f2 := make([]int, n) // not
	// 在卖出时交手续费
	f1[0] = -prices[0]
	for i := 1; i < n; i++ {
		f1[i] = max(f1[i-1], f2[i-1]-prices[i])
		f2[i] = max(f2[i-1], f1[i-1]+prices[i]-fee)
	}
	return f2[n-1]
}
