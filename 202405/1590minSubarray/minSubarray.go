package main

import (
	"fmt"
)

func main() {
	fmt.Println(minSubarray([]int{1, 2, 3}, 3))

}

func minSubarray(nums []int, p int) int {
	n := len(nums)
	preSum := make([]int, len(nums)+1)
	for i, ch := range nums {
		// 处理负数的技巧,这样不管值是多少，余数都会在[0:p)区间中
		preSum[i+1] = ((preSum[i]+ch)%p + p) % p
	}
	x := preSum[n] // 所有和的余数
	if x == 0 {
		return 0
	}
	last := make(map[int]int)
	ans := n
	for i, ch := range preSum {
		// 为啥这里要先加入呢，判断了之后再加入行不行？
		// 这里先加入就是应付 x=0,也就是不用删除的情况，如果在程序外面判断了 x==0的特例，就可以在后面加入
		last[ch] = i
		// 理解这个 a 的计算
		// 通过前缀和，问题转换成：在前缀和数组上找到两个数 right,left 满足right-left最小且preSum[right]和preSum[left]同余
		// 即是：(preSum[right]-preSum[left])%p=x，因为preSum 保存的就是前缀和的余数所以做如下变型
		// 变型：preSum[left] = preSum[right]-x
		// 加上应付负数的小技巧:preSum[left] = (preSum[right]-x +p)%p，设：a = preSum[left]
		// 那么我们遍历到 right 时，就知道的 preSum[right]的值，如果在这之前遍历中，有结果等于a,那么两次值的距离就是需要删除的
		a := (ch - x + p) % p
		if j, ok := last[a]; ok {
			ans = min(ans, i-j)
		}
	}
	if ans == n {
		return -1
	}
	return ans
}
