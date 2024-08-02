package main

func main() {

}

func minCost(startPos []int, homePos []int, rowCosts []int, colCosts []int) int {
	x0, y0, x1, y1 := startPos[0], startPos[1], homePos[0], homePos[1]
	ans := 0
	// 下面循环中对每一个位置都计算了花费，所以先把起始位置的减去
	// 必须在这里先做减法，因为下面会改变x0,y0的值
	ans = ans - rowCosts[x0]
	ans = ans - colCosts[y0]
	if x0 > x1 {
		x0, x1 = x1, x0
	}
	if y0 > y1 {
		y0, y1 = y1, y0
	}
	for i := x0; i <= x1; i++ {
		ans += rowCosts[i]
	}
	for i := y0; i <= y1; i++ {
		ans += colCosts[i]
	}

	return ans
}
