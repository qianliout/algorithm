package main

func main() {

}

func countCharacters(words []string, chars string) int {
	ans := 0
	all := make(map[byte]int)
	for _, ch := range chars {
		all[byte(ch)]++
	}
	for _, word := range words {
		if check(word, all) {
			ans += len(word)
		}
	}
	return ans
}

func check(word string, all map[byte]int) bool {
	cnt := make(map[byte]int)
	for _, ch := range word {
		cnt[byte(ch)]++
	}
	for k, v := range cnt {
		if all[k] < v {
			return false
		}
	}
	return true
}
