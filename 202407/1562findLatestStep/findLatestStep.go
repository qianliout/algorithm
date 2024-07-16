package main

func main() {

}

func findLatestStep(arr []int, m int) int {
	cnt := 0
	res := -1
	n := len(arr)
	link := make([]int, n+2)
	for i, ch := range arr {
		l, r := ch, ch
		if link[ch-1] > 0 {
			l = link[ch-1]
		}
		if link[ch+1] > 0 {
			r = link[ch+1]
		}
		if ch-l == m {
			cnt--
		}
		if r-ch == m {
			cnt--
		}
		if r-l+1 == m {
			cnt++
		}
		if cnt > 0 {
			res = i + 1 // 下标从1开始
		}
		link[l] = r
		link[r] = l
	}

	return res
}

// 没有完全理解透彻
// https://leetcode.cn/problems/find-latest-group-of-size-m/solutions/386109/on-de-jie-jue-fang-fa-by-time-limit/
