package main

func main() {

}

func numberOfPoints(nums [][]int) int {
	// 最大值是100
	diff := [102]int{}
	ans := 0
	for _, p := range nums {
		diff[p[0]]++
		diff[p[1]+1]--
	}
	s := 0
	for _, d := range diff {
		s += d
		if s > 0 {
			ans++
		}
	}
	return ans
}
