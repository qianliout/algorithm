package main

func main() {

}

// 直接暴力不能过
// 枚举中间点
func numTeams(rating []int) int {
	n := len(rating)
	cnt := 0
	for j := 1; j < n-1; j++ {
		leftLess, leftMore := 0, 0
		for i := j - 1; i >= 0; i-- {
			if rating[i] > rating[j] {
				leftMore++
			}
			if rating[i] < rating[j] {
				leftLess++
			}
		}
		rightLess, rightMore := 0, 0
		for k := j + 1; k < n; k++ {
			if rating[k] > rating[j] {
				rightMore++
			}
			if rating[k] < rating[j] {
				rightLess++
			}
		}
		cnt += leftLess*rightMore + rightLess*leftMore

	}
	return cnt
}
