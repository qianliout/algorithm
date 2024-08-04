package main

func main() {

}

// 会超时
func getDistances(arr []int) []int64 {
	n := len(arr)
	ans := make([]int64, n)
	cnt := make(map[int][]int)

	for i, ch := range arr {
		cnt[ch] = append(cnt[ch], i)
	}
	for i, ch := range arr {
		ans[i] = cal(cnt[ch], i)
	}
	return ans
}

func cal(nums []int, a int) int64 {
	ans := 0
	for _, ch := range nums {
		if ch == a {
			continue
		}
		ans += abs(ch - a)
	}
	return int64(ans)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
