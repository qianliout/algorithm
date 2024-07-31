package main

import (
	"math"
)

func main() {

}

func firstDayBeenInAllRooms(nextVisit []int) int {
	mod := int(math.Pow10(9)) + 7
	n := len(nextVisit)
	f := make([]int, n)
	for i := 1; i < n; i++ {
		f[i] = (f[i-1] + 1 + f[i-1] - f[nextVisit[i-1]] + 1 + mod) % mod
	}
	return f[n-1]
}

//  我们定义 f[i] 表示第一次访问第 i 号房间的日期编号，那么答案就是 f[n−1]。
// 我们考虑第一次到达第 i−1 号房间的日期编号，记为 f[i−1]，此时需要花一天的时间回退到第 nextVisit[i−1] 号房间，为什么是回退呢？因为题目限制了 0≤nextVisit[i]≤i。
// 回退之后，此时第 nextVisit[i−1] 号房间的访问为奇数次，而第 [nextVisit[i−1]+1,..i−1] 号房间均被访问偶数次，那么这时候我们从第
// nextVisit[i−1] 号房间再次走到第 i−1 号房间，就需要花费 f[i−1]−f[nextVisit[i−1]] 天的时间，然后再花费一天的时间到达第 i 号房间，因此 f[i]=f[i−1]+1+f[i−1]−f[nextVisit[i−1]]+1
