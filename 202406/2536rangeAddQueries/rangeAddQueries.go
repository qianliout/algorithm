package main

import (
	"fmt"
)

func main() {
	fmt.Println(rangeAddQueries(3, [][]int{{1, 1, 2, 2}, {0, 0, 1, 1}}))
}

// 注意不能转换为一维去处理，因为一行的数据的结尾，不应该影响另一行数据的开始
func rangeAddQueries(n int, queries [][]int) [][]int {
	// 第一行，第一列是其列，全为0
	// 因为差分数组会对下一个数做减一操作，所以这里是 n+2
	matrix := make([][]int, n+2)
	for i := range matrix {
		matrix[i] = make([]int, n+2)
	}

	for _, ch := range queries {
		// 为啥要加1呢，因为下面要用到前缀和的思想，第一行第一列是基行基列
		r1, c1, r2, c2 := ch[0]+1, ch[1]+1, ch[2]+1, ch[3]+1
		matrix[r1][c1]++
		matrix[r1][c2+1]--
		matrix[r2+1][c1]--
		matrix[r2+1][c2+1]++
	}
	// 利用差分数组的思想，前缀和的方法，复原数组
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			matrix[i][j] += matrix[i][j-1] + matrix[i-1][j] - matrix[i-1][j-1]
		}
	}
	// 只保留需要的部分
	matrix = matrix[1 : n+1]
	for i := range matrix {
		matrix[i] = matrix[i][1 : n+1]
	}

	return matrix
}

// 注意不能转换为一维去处理，因为一行的数据的结尾，不应该影响另一行数据的开始
// 分一行做一个一维差分
func rangeAddQueries2(n int, queries [][]int) [][]int {
	matrix := make([][]int, n+2)
	for i := range matrix {
		matrix[i] = make([]int, n+2)
	}

	for _, ch := range queries {
		// 为啥要加1呢，因为下面要用到前缀和的思想，第一行第一列是基行基列
		r1, c1, r2, c2 := ch[0]+1, ch[1]+1, ch[2]+1, ch[3]+1
		for i := r1; i <= r2; i++ {
			matrix[i][c1]++
			matrix[i][c2+1]--
		}
	}
	// 求前缀和
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			matrix[i][j] += matrix[i][j-1]
		}
	}

	// 只保留需要的部分
	matrix = matrix[1 : n+1]
	for i := range matrix {
		matrix[i] = matrix[i][1 : n+1]
	}

	return matrix
}

// 更简洁的写法
func rangeAddQueries3(n int, queries [][]int) [][]int {
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}

	for _, ch := range queries {
		// 为啥要加1呢，因为下面要用到前缀和的思想，第一行第一列是基行基列
		r1, c1, r2, c2 := ch[0], ch[1], ch[2], ch[3]
		for i := r1; i <= r2; i++ {
			matrix[i][c1]++
			if c2+1 < n {
				matrix[i][c2+1]--
			}
		}
	}
	// 求前缀和
	for i := 0; i < n; i++ {
		for j := 1; j < n; j++ {
			matrix[i][j] += matrix[i][j-1]
		}
	}

	return matrix
}
