package main

func main() {

}

func minAreaRect(points [][]int) int {
	col := make(map[int]int)
	row := make(map[int]int)
	mxC, mxR := 0, 0
	for _, ch := range points {
		x, y := ch[0], ch[1]
		col[x]++
		row[y]++
		if col[x] >= 2 && x > mxC {
			mxC = x
		}
		if row[y] >= 2 && y > mxR {
			mxR = y
		}
	}

	return mxR * mxC
}
