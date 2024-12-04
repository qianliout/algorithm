package main

func main() {

}

func minCostClimbingStairs1(cost []int) int {
	f0, f1 := 0, 0
	n := len(cost)
	for i := 2; i <= n; i++ {
		f2 := min(f0+cost[i-2], f1+cost[i-1])
		f0, f1 = f1, f2
	}
	return f1
}
func minCostClimbingStairs(cost []int) int {
	n := len(cost)
	f := make([]int, n)
	for i := 2; i < n; i++ {
		f[i] = min(f[i-2]+cost[i-2], f[i-1]+cost[i-1])
	}
	return f[n-1]
}
