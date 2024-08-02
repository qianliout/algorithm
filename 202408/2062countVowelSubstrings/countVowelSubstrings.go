package main

func main() {

}

func countVowelSubstrings(word string) int {
	// 数据量小，直接模拟
	ans, n := 0, len(word)
	for i := 0; i < n; i++ {
		for j := i + 5; j <= n; j++ {
			if check(word[i:j]) {
				ans++
			}
		}
	}
	return ans
}

func check(s string) bool {
	bi := map[byte]int{
		'a': 1,
		'e': 1,
		'i': 1,
		'o': 1,
		'u': 1,
	}
	cnt := make(map[byte]int)
	a := 0
	for _, ch := range s {
		if bi[byte(ch)] > 0 {
			a++
			cnt[byte(ch)]++
		}
	}
	return a == len(s) && len(cnt) == 5
}
