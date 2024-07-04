package main

import (
	"fmt"
)

func main() {
	fmt.Println(removeDuplicates("deeedbbcccbdaa", 3))
}

func removeDuplicates(s string, k int) string {
	st := make([]pair, 0)
	for _, ch := range s {
		if len(st) == 0 || st[len(st)-1].by != byte(ch) {
			st = append(st, pair{by: byte(ch), cnt: 1})
		} else {
			st[len(st)-1].cnt++
			if st[len(st)-1].cnt == k {
				st = st[:len(st)-1]
			}
		}
	}
	ans := make([]byte, 0)
	for i := range st {
		for j := 0; j < st[i].cnt; j++ {
			ans = append(ans, st[i].by)
		}
	}
	return string(ans)
}

type pair struct {
	by  byte
	cnt int
}
