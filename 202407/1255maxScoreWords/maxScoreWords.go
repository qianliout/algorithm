package main

func main() {

}

func maxScoreWords(words []string, letters []byte, score []int) int {
	cnt := make(map[byte]int)
	for _, ch := range letters {
		cnt[byte(ch)]++
	}
	n := len(words)
	var dfs func(i int) int
	dfs = func(i int) int {
		if i >= n {
			return 0
		}
		do := 0
		wc, ok := check(words[i], cnt)
		if ok {
			for k, v := range wc {
				cnt[k] -= v
			}
			do = cal(words[i], score) + dfs(i+1)
			for k, v := range wc {
				cnt[k] += v
			}
		}
		not := dfs(i + 1)
		return max(do, not)
	}
	return dfs(0)
}

func check(word string, cnt2 map[byte]int) (map[byte]int, bool) {
	cnt := make(map[byte]int)
	for _, ch := range word {
		cnt[byte(ch)]++
	}
	for k, v := range cnt {
		if cnt2[k] < v {
			return cnt, false
		}
	}
	return cnt, true
}

func cal(word string, score []int) int {
	ans := 0
	for _, ch := range word {
		ans += score[int(ch)-'a']
	}

	return ans
}
