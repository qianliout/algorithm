package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(maxValue([][]int{{1, 2, 4}, {3, 4, 3}, {2, 3, 1}}, 2))

}

func maxValue(events [][]int, k int) int {
	sort.Slice(events, func(i, j int) bool { return events[i][1] < events[j][1] })
	dp := make([][]int, len(events)+1)
	for i := range dp {
		dp[i] = make([]int, k+1)
	}

	// 参加0次会议的价值是0，所以 dp[i][0] = 0
	for i := 0; i < len(events); i++ {
		e := events[i]
		for j := 1; j <= k; j++ {
			dp[i+1][j] = max(dp[i+1][j], dp[i][j]) // 不选 i
			// 那么你必须 完整 地参加完这个会议。会议结束日期是包含在会议内的，也就是说你不能同时参加一个开始日期与另一个结束日期相同的两个会议
			// 用二分查找， sort.Search 只能用于查左边界，但是本例是查找一个events[j],使events[j][1]<e[0],这样的值右边界
			// 所以得转化思路 查找一个 events[j][1]>=e[0] 这样的一个值的左边界，此时的 j，就是我们要找的值的下一个下标值
			p := sort.Search(i, func(j int) bool { return events[j][1] >= e[0] }) - 1
			// 这里为啥是 p+1呢，因为我们把 i 整体都+1了，就是为了防止-1的下标
			dp[i+1][j] = max(dp[i+1][j], dp[p+1][j-1]+e[2]) // 选
		}
	}
	return dp[len(events)][k]
}
