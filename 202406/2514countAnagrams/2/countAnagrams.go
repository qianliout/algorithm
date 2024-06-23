package main

import (
	"math"
	"strings"
)

func main() {

}

// 会超时
func countAnagrams(s string) int {
	base := int(math.Pow10(9)) + 7
	split := strings.Split(s, " ")
	ans := 1
	for _, wo := range split {
		ans = (ans * count(wo)) % base
	}
	return ans % base
}

func count(word string) int {
	visit := make([]bool, len(word))
	ans := make(map[string]bool)
	dfs([]byte(word), []byte{}, visit, ans)
	return len(ans)
}

func dfs(data []byte, path []byte, visit []bool, ans map[string]bool) {
	if len(path) == len(data) {
		ans[string(path)] = true
		return
	}
	for i := 0; i < len(data); i++ {
		if visit[i] {
			continue
		}
		visit[i] = true
		path = append(path, data[i])
		dfs(data, path, visit, ans)
		path = path[:len(path)-1]
		visit[i] = false
	}
}
