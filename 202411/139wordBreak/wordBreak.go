package main

func main() {

}

func wordBreak(s string, wordDict []string) bool {
	n := len(s)
	word := make(map[string]bool)
	mx := 0
	for _, ch := range wordDict {
		word[ch] = true
		mx = max(mx, len(ch))
	}
	mem := make(map[int]bool)
	var dfs func(i int) bool

	dfs = func(i int) bool {
		if i >= n {
			return true
		}
		if v, ok := mem[i]; ok {
			return v
		}
		for j := i + 1; j <= min(n, j+mx); j++ {
			if word[s[i:j]] && dfs(j) {
				mem[i] = true
				return true
			}
		}
		mem[i] = false
		return false
	}
	return dfs(0)
}
