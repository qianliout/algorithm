package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	fmt.Println(makeLargestSpecial("11011000"))
}

func makeLargestSpecial(s string) string {
	return dfs(s)
}

// 理解成有效的括号
func dfs(s string) string {
	candidates := sort.StringSlice{}

	if len(s) == 0 {
		return ""
	}
	cnt := 0
	last := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '1' {
			cnt++
		} else {
			cnt--
		}
		if cnt == 0 {
			// 如现有效的括号组了，那么此时有定是以下这种结构
			str := "1" + dfs(s[last+1:i]) + "0"
			candidates = append(candidates, str)
			last = i + 1
		}
	}
	sort.Sort(sort.Reverse(candidates))
	return strings.Join(candidates, "")
}
