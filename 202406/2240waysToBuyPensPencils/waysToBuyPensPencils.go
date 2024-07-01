package main

func main() {

}

func waysToBuyPensPencils(total int, cost1 int, cost2 int) int64 {
	ans := 0
	// 枚举买钢笔的机会
	for x := 0; x <= total/cost1; x++ {
		// 剩下的钱就是买铅笔
		y := (total - (x * cost1)) / cost2
		ans += y + 1 // 加 1  是加上买钢笔的方案
	}
	return int64(ans)
}
