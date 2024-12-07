package main

import (
	"fmt"
)

func main() {
	fmt.Println(numTrees(3))
}

func numTrees2(n int) int {
	var dfs func(i, j int) int
	mem := make([][]int, n+1)
	for i := range mem {
		mem[i] = make([]int, n+1)
		for j := range mem[i] {
			mem[i][j] = -1
		}
	}

	dfs = func(i, j int) int {
		// 空树也是一个树
		if i >= j {
			return 1
		}
		if mem[i][j] != -1 {
			return mem[i][j]
		}
		ans := 0
		for k := i; k <= j; k++ {
			ans += dfs(i, k-1) * dfs(k+1, j)
		}
		mem[i][j] = ans
		return ans
	}
	ans := dfs(1, n)
	return ans
}

func numTrees1(n int) int {
	dp := make([]int, n+1)

	dp[0] = 1
	dp[1] = 1
	// 初值
	for i := 2; i <= n; i++ {
		for j := 1; j <= i; j++ {
			dp[i] = dp[i] + dp[i-j]*dp[j-1]
		}
	}
	return dp[n]
}

func numTrees(n int) int {
	// 创建一个二维切片来存储子问题的结果
	mem := make([][]int, n+1)
	for i := range mem {
		mem[i] = make([]int, n+1)
		for j := range mem[i] {
			mem[i][j] = 0
		}
	}

	// 初始化基本情况：当 i >= j 时，只有一种树的可能性，即空树或单节点树
	for i := 1; i <= n; i++ {
		for j := i; j <= n; j++ {
			if i == j {
				mem[i][j] = 1
			}
		}
	}

	// 填充 mem 数组
	for length := 2; length <= n; length++ { // 子树的长度从 2 开始
		for i := 1; i <= n-length+1; i++ {
			j := i + length - 1
			for k := i; k <= j; k++ {
				left := 1
				right := 1
				if k > i {
					left = mem[i][k-1]
				}
				if k < j {
					right = mem[k+1][j]
				}
				mem[i][j] += left * right
			}
		}
	}

	return mem[1][n]
}

func numTrees4(n int) int {
	var dfs func(i int) int
	mem := make([]int, n+1)
	for i := range mem {
		mem[i] = -1
	}

	dfs = func(i int) int {
		// 空树也是一个树
		if i <= 1 {
			return 1
		}
		if mem[i] != -1 {
			return mem[i]
		}
		ans := 0
		for k := 1; k <= i; k++ {
			ans += dfs(k-1) * dfs(i-k)
		}
		mem[i] = ans
		return ans
	}
	ans := dfs(n)
	return ans
}
