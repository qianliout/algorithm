package main

func main() {

}

func minOperationsMaxProfit(customers []int, boardingCost int, runningCost int) int {
	ans := -1
	mx, i := 0, 0
	wait := 0
	t := 0
	cnt := 0 // 转了多少次轮
	for wait > 0 || i < len(customers) {
		if i < len(customers) {
			wait += customers[i]
		}
		up := min(4, wait)
		wait -= up
		t += up*boardingCost - runningCost
		if i < len(customers) {
			i++
		}
		cnt++
		if t > mx {
			mx = t
			ans = cnt
		}
	}
	return ans
}
