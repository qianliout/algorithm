package main

func main() {

}

// 直接模拟会超时
func wordSubsets(words1 []string, words2 []string) []string {
	w1 := make([][]int, len(words1))
	for i := range words1 {
		w1[i] = gen(words1[i])
	}

	w2 := make([][]int, len(words2))
	for i := range words2 {
		w2[i] = gen(words2[i])
	}
	ans := make([]string, 0)
	for i := 0; i < len(words1); i++ {
		find := true
		for j := range words2 {
			if !check(w1[i], w2[j]) {
				find = false
				break
			}
		}
		if find {
			ans = append(ans, words1[i])
		}
	}
	return ans
}

func check(word1, word2 []int) bool {
	for i := 0; i < 26; i++ {
		if word1[i] < word2[i] {
			return false
		}
	}
	return true
}

func gen(word string) []int {
	ans := make([]int, 26)
	for _, ch := range word {
		ans[int(ch-'a')]++
	}
	return ans
}
