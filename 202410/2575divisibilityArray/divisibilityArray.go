package main

func main() {

}

func divisibilityArray(word string, m int) []int {
	n := len(word)
	ans := make([]int, n)
	num := 0
	for i, ch := range word {
		num = (num*10 + int(ch-'0')) % m
		if num == 0 {
			ans[i] = 1
		}
	}
	return ans
}
