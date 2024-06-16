package main

import (
	"fmt"
)

func main() {
	fmt.Println(sortString("aaaabbbbcccc"))
}

func sortString(s string) string {
	cnt := make([]int, 26)
	for _, c := range s {
		cnt[int(c-'a')]++
	}
	ans := make([]byte, 0)
	for len(ans) < len(s) {
		v1 := make([]byte, 0)
		for i := 0; i < 26; i++ {
			if cnt[i] >= 1 {
				v1 = append(v1, byte(i+'a'))
				cnt[i]--
			}
		}
		v2 := make([]byte, 0)
		for i := 25; i >= 0; i-- {
			if cnt[i] >= 1 {
				v2 = append(v2, byte(i+'a'))
				cnt[i]--
			}
		}
		ans = append(ans, v1...)
		ans = append(ans, v2...)
	}
	return string(ans)
}
