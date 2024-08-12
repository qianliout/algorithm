package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(mostPopularCreator([]string{"alice", "alice", "alice"}, []string{"a", "b", "c"}, []int{1, 2, 2}))
}

func mostPopularCreator(creators []string, ids []string, views []int) [][]string {
	n := len(creators)
	cnt1 := make(map[string]int)
	cnt2 := make(map[string][]pair)
	for i := 0; i < n; i++ {
		cnt1[creators[i]] += views[i]
		cnt2[creators[i]] = append(cnt2[creators[i]], pair{creator: creators[i], id: ids[i], view: views[i]})
	}
	mx, cre := 0, make([]string, 0)
	for k, v := range cnt1 {
		if v == mx {
			cre = append(cre, k)
			continue
		}
		if v > mx {
			mx = v
			cre = []string{k}
		}
	}
	ans := make([][]string, 0)

	for _, c := range cre {
		ch := cnt2[c]
		sort.Slice(ch, func(i, j int) bool {
			if ch[i].view != ch[j].view {
				return ch[i].view > ch[j].view
			} else {
				return ch[i].id < ch[j].id
			}
		})
		ans = append(ans, []string{ch[0].creator, ch[0].id})
	}
	return ans
}

type pair struct {
	creator string
	id      string
	view    int
}
