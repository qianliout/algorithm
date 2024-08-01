package main

func main() {

}

func winnerOfGame(colors string) bool {
	cnt1 := cal(colors, 'A')
	cnt2 := cal(colors, 'B')
	return cnt1 > cnt2
}

func cal(colors string, by byte) int {
	ans := 0
	i, n := 0, len(colors)
	for i < n {
		if colors[i] != by {
			i++
			continue
		}
		j := i
		for j < n && colors[j] == by {
			j++
		}
		if j-i >= 3 {
			// 不能取 前后，只能取中间
			ans += j - i - 2
		}
		i = j
	}
	return ans
}
