package main

func main() {

}

func intervalIntersection(firstList [][]int, secondList [][]int) [][]int {
	le, ri := 0, 0
	m, n := len(firstList), len(secondList)
	// 已经排好需的了
	ans := make([][]int, 0)

	for le < m && ri < n {
		if firstList[le][1] < secondList[ri][0] {
			le++
		} else if firstList[le][0] > secondList[ri][1] {
			ri++
		} else {
			a := max(firstList[le][0], secondList[ri][0])
			b := min(firstList[le][1], secondList[ri][1])
			ans = append(ans, []int{a, b})
			if firstList[le][1] > secondList[ri][1] {
				ri++
			} else {
				le++
			}
		}
	}
	return ans
}
