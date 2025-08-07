package main

func main() {

}

func findTargetSumWays2(nums []int, target int) int {
	all := 0
	for _, ch := range nums {
		all += ch
	}
	if (all+target)%2 != 0 {
		return 0
	}
	target = (target + all) / 2
	n := len(nums)
	var dfs func(i int, sum int) int
	dfs = func(i int, sum int) int {
		// 因为要全部选择，所以这里一定要判断i>=n
		if i >= n && sum == target {
			return 1
		}
		if i < 0 || i >= n {
			return 0
		}
		a := dfs(i+1, sum)
		b := dfs(i+1, sum+nums[i])
		return a + b
	}
	ans := dfs(0, 0)
	return ans
}

// [1,0],1

func findTargetSumWays1(nums []int, target int) int {
	all := 0
	for _, ch := range nums {
		all += ch
	}
	if (all+target)%2 != 0 {
		return 0
	}
	target = (target + all) / 2
	n := len(nums)
	var dfs func(i, sum int) int

	dfs = func(i, sum int) int {
		if i < 0 {
			if sum == target {
				return 1
			}
			return 0
		}
		a := dfs(i-1, sum)
		b := dfs(i-1, sum+nums[i])
		return a + b
	}
	ans := dfs(n-1, 0)
	return ans
}

// [1,0],1

func findTargetSumWays(nums []int, target int) int {
	all := 0
	for _, ch := range nums {
		all += ch
	}
	if target < 0 || (all+target)%2 != 0 {
		return 0
	}
	target = (target + all) / 2
	n := len(nums)
	f := make([][]int, n+5)
	for i := range f {
		f[i] = make([]int, target+5)
	}
	// 初值
	f[0][0] = 1
	for i := 0; i < n; i++ {
		ch := nums[i]
		for j := 0; j <= target; j++ {
			f[i+1][j] += f[i][j] // 不选

			if ch <= j {
				f[i+1][j] += f[i][j-ch]
			}
		}
	}
	return f[n][target]
}
