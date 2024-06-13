package main

func main() {

}

// 没有理解透彻
func totalSteps(nums []int) int {
	stark := make([]pair, 0)
	ans := 0
	for _, num := range nums {
		mt := 0
		// 严格的单调递减栈
		for len(stark) > 0 && stark[len(stark)-1].num <= num {
			mt = max(mt, stark[len(stark)-1].at)
			stark = stark[:len(stark)-1]
		}
		if len(stark) > 0 {
			mt++
			ans = max(ans, mt)
		} else {
			mt = 0
		}
		stark = append(stark, pair{num: num, at: mt})
	}
	return ans
}

type pair struct {
	num, at int
}
