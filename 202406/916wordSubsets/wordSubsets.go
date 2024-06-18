package main

func main() {

}

func wordSubsets(words1 []string, words2 []string) []string {
	w1 := make([][]int, len(words1))
	for i := range words1 {
		w1[i] = gen(words1[i])
	}

	w2 := make([][]int, len(words2))
	// 是words2的每一个字符串的母集，那就只关心最大值就行了
	ww2 := make([]int, 26)
	for i := range words2 {
		w2[i] = gen(words2[i])
		for j, ch := range w2[i] {
			ww2[j] = max(ww2[j], ch)
		}
	}

	ans := make([]string, 0)
	for i := 0; i < len(words1); i++ {
		if check(w1[i], ww2) {
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
