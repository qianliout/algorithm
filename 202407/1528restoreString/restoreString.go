package main

import (
	"sort"
)

func main() {

}

func restoreString(s string, indices []int) string {
	n := len(s)
	res := make([]pair, n)
	for i, ch := range s {
		res[i] = pair{ch: byte(ch), idx: indices[i]}
	}
	sort.Slice(res, func(i, j int) bool { return res[i].idx < res[j].idx })
	res2 := make([]byte, 0)
	for i := range res {
		res2 = append(res2, res[i].ch)
	}
	return string(res2)
}

type pair struct {
	ch  byte
	idx int
}
