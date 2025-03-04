package main

func main() {

}

func getResults(queries [][]int) []bool {
	return nil
}

type seg struct {
	left  int
	right int
	mx    int // 这个区间内空的区间的最大值
	sum   int
}

type segTree []seg

func (st segTree) update(o int, idx int, add int) {

}

func (st segTree) querySum(o int, L, R int) int {
	return o
}
