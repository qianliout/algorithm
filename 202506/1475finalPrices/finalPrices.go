package main

func main() {

}

func finalPrices(prices []int) []int {
	n := len(prices)
	dis := make([]int, n)

	st := make([]int, 0)
	for i, c := range prices {
		for len(st) > 0 && prices[st[len(st)-1]] >= c {
			last := st[len(st)-1]
			st = st[:len(st)-1]
			dis[last] = c
		}
		st = append(st, i)
	}
	ans := make([]int, n)
	for i, c := range prices {
		ans[i] = c - dis[i]
	}
	return ans
}
