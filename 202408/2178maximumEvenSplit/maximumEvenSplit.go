package main

func main() {

}

func maximumEvenSplit(finalSum int64) []int64 {
	if finalSum%2 != 0 {
		return []int64{}
	}
	var a int64 = 2
	ans := make([]int64, 0)
	for a <= finalSum {
		ans = append(ans, int64(a))
		finalSum -= a
		a += 2
	}
	ans[len(ans)-1] += finalSum

	return ans
}
