package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(longestDiverseString(1, 1, 7))
}

func longestDiverseString(a int, b int, c int) string {
	cnt := []pair{{ch: 'a', cnt: a}, {ch: 'b', cnt: b}, {ch: 'c', cnt: c}}
	ans := make([]byte, 0)
	for {
		sort.Slice(cnt, func(i, j int) bool { return cnt[i].cnt > cnt[j].cnt })
		hasNext := false
		for i, p := range cnt {
			if p.cnt == 0 {
				break
			}
			m := len(ans)
			if m >= 2 && ans[m-1] == p.ch && ans[m-2] == p.ch {
				continue
			}
			hasNext = true
			ans = append(ans, p.ch)
			cnt[i].cnt--
			// 一定要 break 进行下一轮重新选择
			break
		}
		if !hasNext {
			return string(ans)
		}
	}
}

type pair struct {
	ch  byte
	cnt int
}
