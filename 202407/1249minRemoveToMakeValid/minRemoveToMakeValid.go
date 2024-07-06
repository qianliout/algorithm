package main

import "fmt"

func main() {
	fmt.Println(minRemoveToMakeValid("lee(t(c)o)de)"))
}

func minRemoveToMakeValid(s string) string {
	ans := make([]byte, 0)
	cnt := make([]int, 0)
	del := make([]int, 0)
	for i, ch := range s {
		switch ch {
		case ')':
			if len(cnt) == 0 {
				del = append(del, i)
			} else {
				cnt = cnt[:len(cnt)-1]
			}
		case '(':
			cnt = append(cnt, i)
		}
	}
	delM := make(map[int]bool)
	for _, ch := range del {
		delM[ch] = true
	}
	for _, ch := range cnt {
		delM[ch] = true
	}

	for i, ch := range s {
		if delM[i] {
			continue
		}
		ans = append(ans, byte(ch))
	}

	return string(ans)
}
