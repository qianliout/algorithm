package main

func main() {

}

func maximumSegmentSum(nums []int, removeQueries []int) []int64 {
	n := len(nums)
	sum := make([]int64, n+1)
	fa := make([]int, n+1)
	for i := range fa {
		fa[i] = i
	}
	var find func(x int) int
	find = func(x int) int {
		if fa[x] != x {
			fa[x] = find(fa[x])
		}
		return fa[x]
	}
	ans := make([]int64, n)
	for i := n - 1; i > 0; i-- {
		x := removeQueries[i]
		// 把x合并到 x+1 上，这里没有理解
		// 为啥是合并到 x+1 而不合并 x-1呢
		// todo
		to := find(x + 1)
		fa[x] = to
		sum[to] += sum[x] + int64(nums[x])
		ans[i-1] = max(ans[i], sum[to])
	}
	return ans
}
