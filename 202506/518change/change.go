package main

func main() {

}

func change(amount int, coins []int) int {
	n := len(coins)
	f := make([][]int, n+1)
	for i := range f {
		f[i] = make([]int, amount+1)
	}
	for i := 0; i < n; i++ {
		for t := 0; t <= amount; t++ {
			c := coins[i]
			f[i+1][t] = f[i][t]
			if t-c >= 0 {
				f[i+1][t] += f[i+1][t-c] + 1
			}
		}
	}
	return f[n][amount]
}
