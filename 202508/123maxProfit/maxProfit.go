package main

func main() {

}

func maxProfit(prices []int) int {
	n := len(prices)
	inf := 1 << 30
	// dfs(i,true) 第i天结束时未持有股票
	// dfs(i,false) 第i天结束时持有股票
	var dfs func(i int, has bool, k int) int
	dfs = func(i int, has bool, k int) int {
		if i < 0 || k > 2 {
			return -inf
		}
		if has {
			a := dfs(i-1, true, k)
			b := dfs(i-1, false, k-1) - prices[i]
			return max(a, b)
		}
		a := dfs(i-1, false, k)
		b := dfs(i-1, true, k-1) - prices[i]
		return max(a, b)
	}
	ans := dfs(n-1, false, 2)
	return ans
}

// 你最多可以完成 两笔 交易。
