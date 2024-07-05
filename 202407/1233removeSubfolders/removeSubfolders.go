package main

import (
	"sort"
	"strings"
)

func main() {

}
func removeSubfolders(folder []string) []string {
	sort.Strings(folder)
	ans := make([]string, 0)

	for _, ch := range folder {
		if len(ans) == 0 {
			ans = append(ans, ch)
			continue
		}
		last := ans[len(ans)-1]
		// 一定得是 last+"/" 防止 /a/b /a/bc 这种
		if strings.HasPrefix(ch, last+"/") {
			continue
		}
		ans = append(ans, ch)
	}
	return ans
}
