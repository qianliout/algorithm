package main

func main() {

}

func largeGroupPositions(s string) [][]int {
	ans := make([][]int, 0)
	start := 0
	end := 0
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			end++
		} else {
			if end-start >= 2 {
				ans = append(ans, []int{start, end})
			}
			start = i
			end = i
		}
	}
	if end-start >= 2 {
		ans = append(ans, []int{start, end})
	}
	return ans
}
