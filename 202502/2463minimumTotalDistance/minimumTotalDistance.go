package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	// fmt.Println(minimumTotalDistance([]int{0, 4, 6}, [][]int{{2, 2}, {6, 2}}))
	// fmt.Println(minimumTotalDistance2([]int{0, 4, 6}, [][]int{{2, 2}, {6, 2}}))
	fmt.Println(minimumTotalDistance2([]int{789300819, -600989788, 529140594, -592135328, -840831288, 209726656, -671200998}, [][]int{{-865262624, 6}, {-717666169, 0}, {725929046, 2}, {449443632, 3}, {-912630111, 0}, {270903707, 3}, {-769206598, 2}, {-299780916, 4}, {-159433745, 5}, {-467185764, 3}, {849991650, 7}, {-292158515, 6}, {940410553, 6}, {258278787, 0}, {83034539, 2}, {54441577, 3}, {-235385712, 2}, {75791769, 3}}))
	fmt.Println(minimumTotalDistance([]int{789300819, -600989788, 529140594, -592135328, -840831288, 209726656, -671200998}, [][]int{{-865262624, 6}, {-717666169, 0}, {725929046, 2}, {449443632, 3}, {-912630111, 0}, {270903707, 3}, {-769206598, 2}, {-299780916, 4}, {-159433745, 5}, {-467185764, 3}, {849991650, 7}, {-292158515, 6}, {940410553, 6}, {258278787, 0}, {83034539, 2}, {54441577, 3}, {-235385712, 2}, {75791769, 3}}))
}

func minimumTotalDistance(robot []int, factory [][]int) int64 {
	sort.Ints(robot)
	sort.Slice(factory, func(i, j int) bool { return factory[i][0] < factory[j][0] })

	m := len(robot)
	f := make([]int, m+1)
	for i := range f {
		f[i] = math.MaxInt / 2
	}
	f[0] = 0
	for _, fa := range factory {
		for j := m; j > 0; j-- {
			cost := 0
			for k := 1; k <= min(j, fa[1]); k++ {
				cost += abs(robot[j-k] - fa[0])
				f[j] = min(f[j], f[j-k]+cost)
			}
		}
	}
	return int64(f[m])
}

func minimumTotalDistance2(robot []int, factory [][]int) int64 {
	// 对工厂按照位置排序
	sort.Slice(factory, func(i, j int) bool { return factory[i][0] < factory[j][0] })
	// 对机器人位置排序
	sort.Ints(robot)
	m, n := len(robot), len(factory)

	// 初始化二维 dp 数组，存储中间结果
	// 前i 个工厂修复前 j 个机器人
	inf := math.MaxInt64 / 100
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
		for j := range dp[i] {
			dp[i][j] = inf
		}
	}
	// 这里的初值容易出错，不能和一维一样，只写 dp[0][0] = 0
	for i := 0; i <= n; i++ {
		dp[i][0] = 0
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			dp[i][j] = dp[i-1][j] //  不修复
			// 修
			cost := 0
			limit := factory[i-1][1]
			x := factory[i-1][0]
			for k := 1; k <= limit; k++ {
				if j-k < 0 {
					break
				}
				cost += abs(robot[j-k] - x)
				dp[i][j] = min(dp[i][j], dp[i-1][j-k]+cost)
			}
		}
	}
	return int64(dp[n][m])
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
