package main

import (
	"fmt"
)

func main() {
	fmt.Println(robotWithString("bdda"))
}

func robotWithString(s string) string {
	st := make([]byte, 0)
	cnt := make([]int, 26)
	for _, ch := range s {
		idx := int(ch) - int('a')
		cnt[idx]++
	}
	mi := byte(0)
	ans := make([]byte, 0)
	for _, ch := range s {
		idx := int(ch) - int('a')
		cnt[idx]--
		for mi < 25 && cnt[mi] == 0 {
			mi++
		}
		st = append(st, byte(ch))
		for len(st) > 0 && st[len(st)-1]-'a' <= mi {
			ans = append(ans, st[len(st)-1])
			st = st[:len(st)-1]
		}
	}
	return string(ans)
}

// 贪心地思考，为了让字典序最小，在遍历 s 的过程中，如果栈顶字符 ≤ 后续字符（未入栈）的最小值，那么应该出栈并加到答案末尾，
// 否则应当继续遍历，取到比栈顶字符小的那个字符，这样才能保证字典序最小。
