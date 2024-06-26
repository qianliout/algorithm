package main

func main() {

}

/*
    (a, b) 表示白色车的位置。
    (c, d) 表示白色象的位置。
    (e, f) 表示黑皇后的位置。

假定你只能移动白色棋子，返回捕获黑皇后所需的最少移动次数。

请注意：

    车可以向垂直或水平方向移动任意数量的格子，但不能跳过其他棋子。
    象可以沿对角线方向移动任意数量的格子，但不能跳过其他棋子。
    如果车或象能移向皇后所在的格子，则认为它们可以捕获皇后。
    皇后不能移动。

*/
/*
	如果车能直接攻击到皇后，答案是 1。
	如果象能直接攻击到皇后，答案是 1。
	如果车被象挡住，那么移走象，车就可以攻击到皇后，答案是 2。小知识：这在国际象棋中称作「闪击」。
	如果象被车挡住，那么移走车，象就可以攻击到皇后，答案是 2。
	如果车不能直接攻击到皇后，那么车可以水平移动或者垂直移动，其中一种方式必定不会被象挡住，可以攻击到皇后，答案是 222。

对于车，如果和皇后在同一水平线或者同一竖直线，且中间没有象，那么就可以直接攻击到皇后。
*/
func minMovesToCaptureTheQueen(a int, b int, c int, d int, e int, f int) int {
	// che a,b
	// 象 c d
	// que e f
	// 车横着攻击
	if a == e && (c != e || !ok(b, d, f)) {
		return 1
	}
	// 车竖着攻击
	if b == f && (d != f || !ok(a, c, e)) {
		return 1
	}
	// 象45度攻击
	if c+d == e+f && (a+b != e+f || !ok(c, a, e)) {
		return 1
	}
	// 象135度攻击
	if c-d == e-f && (a-b != e-f || !ok(d, b, f)) {
		return 1
	}

	return 2
}

// 检查，m是否在 l 和 r 之间
func ok(l, m, r int) bool {
	return min(l, r) < m && m < max(l, r)
}
