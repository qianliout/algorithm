package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(racecar(3)) // 2
	fmt.Println(racecar(6)) // 5
	fmt.Println(racecar(5)) // 7
	fmt.Println(racecar(4)) // 5
}

func racecar(target int) int {
	if target < 0 {
		return 0
	}

	// dp[i]表示 target==i 时需要的最小步数
	dp := make([]int, target*10)
	for i := range dp {
		dp[i] = math.MaxInt / 10
	}
	dp[0] = 0
	// mx := target * 2
	for i := 1; i <= target; i++ {
		// 刚好走 n 步就到了 i
		// 向前走 forward 步, 这里的终点 2*i 得好好理解
		// 这里的终点，i*2，是一个难点，可以这样理解
		// 我们的速度是翻倍的，所以我们最多超过1个i，多了还是要回来
		// 2，结合下面 dp[forwardDis-i] 这里可以看出forwardDis-i>=0,超过之后也是没有意义的

		for forward := 1; 1<<forward-1 <= i*2; forward++ {
			forwardDis := 1<<forward - 1
			if forwardDis == i {
				dp[i] = min(dp[i], forward)
			} else if forwardDis > i {
				// 超过了，就要回头走，回头时需要一个 R 转向
				// 走过了，回头走
				dp[i] = min(dp[i], dp[forwardDis-i]+1+forward)
			} else {
				// 在没有到i就先回头走走
				// back 从0开始是一个容器出错的点，因为可以不往回走
				for back := 0; back < forward; back++ {
					backDis := 1<<back - 1
					dp[i] = min(dp[i], forward+1+back+1+dp[i-(forwardDis-backDis)])
				}
			}
		}
	}
	return dp[target]
}
