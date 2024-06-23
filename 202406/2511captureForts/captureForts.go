package main

func main() {

}

// 题目很不好理解，其实就是找一个第于1的下标，向左或向右走找一个最近近的-1,且是间全是0，最大化这个距离
// 从1到-1，中间全是0，最大化这个距离
func captureForts(forts []int) int {
	ans := 0
	for i, ch := range forts {
		if ch != 1 {
			continue
		}
		// 向左
		j := i - 1
		for j >= 0 && forts[j] == 0 {
			j--
		}
		if j >= 0 && forts[j] == -1 {
			ans = max(ans, i-j-1)
		}
		// 向右
		j = i + 1
		for j < len(forts) && forts[j] == 0 {
			j++
		}
		if j < len(forts) && forts[j] == -1 {
			ans = max(ans, j-i-1)
		}
	}
	return ans
}
