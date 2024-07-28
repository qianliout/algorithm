package main

func main() {

}

func maximizeWin(pos []int, k int) int {
	n := len(pos)
	pre := make([]int, n+1)
	ans, left := 0, 0
	for right, p := range pos {
		for p-pos[left] > k {
			left++
		}
		ans = max(ans, right-left+1+pre[left])
		pre[right+1] = max(pre[right], right-left+1, pre[right+1])
	}
	return ans
}
