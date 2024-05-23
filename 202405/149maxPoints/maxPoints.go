package main

func main() {

}

func maxPoints(points [][]int) int {
	if len(points) <= 1 {
		return len(points)
	}
	ans := 0
	for i := 0; i < len(points); i++ {
		ax, ay := points[i][0], points[i][1]
		for j := i + 1; j < len(points); j++ {
			bx, by := points[j][0], points[j][1]
			// 看i和 j 这两个点上共有几个点
			cnt := 2
			for k := j + 1; k < len(points); k++ {
				cx, cy := points[k][0], points[k][1]
				if (bx-ax)*(cy-by) == (cx-bx)*(by-ay) {
					cnt++
				}
			}
			ans = max(ans, cnt)
		}
	}
	return ans
}
