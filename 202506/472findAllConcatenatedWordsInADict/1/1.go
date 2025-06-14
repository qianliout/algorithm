package main

import (
	"sort"
)

func main() {

}

// 容易理解的解法
func findAllConcatenatedWordsInADict(words []string) []string {
	res := make([]string, 0)
	set := make(map[string]bool)

	if len(words) <= 1 {
		return res
	}

	sort.Slice(words, func(i, j int) bool { return len(words[i]) < len(words[j]) })
	set[words[0]] = true

	for i := 1; i < len(words); i++ {
		if getWord([]byte(words[i]), 0, set) {
			res = append(res, words[i])
		} else {
			set[words[i]] = true
		}
	}
	return res
}

func getWord(word []byte, start int, set map[string]bool) bool {
	length := len(word)
	str := ""
	for start < length {
		str = str + string([]byte{word[start]})
		if set[str] && (start == length-1 || getWord(word, start+1, set)) {
			return true
		}
		start++
	}
	return false
}
