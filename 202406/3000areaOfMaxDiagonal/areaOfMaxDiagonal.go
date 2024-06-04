package main

func main() {

}

// 返回对角线最 长 的矩形的 面积 。如果存在多个对角线长度相同的矩形，返回面积最 大 的矩形的面积。
func areaOfMaxDiagonal(dimensions [][]int) int {
	mxs, mp := 0, 0
	for _, ch := range dimensions {
		s := ch[0] * ch[1]
		p := ch[0]*ch[1] + ch[1]*ch[1]
		if p > mp || (p == mp && s > mxs) {
			mp = p
			mxs = s
		}
	}
	return mxs
}
