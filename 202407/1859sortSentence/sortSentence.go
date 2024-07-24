package main

import (
	"sort"
	"strings"
)

func main() {

}

func sortSentence(s string) string {
	// 序列的单词用单个空格连接起来，且开头和结尾没有任何空格
	split := strings.Split(s, " ")
	items := make([]item, 0)
	for _, ch := range split {
		items = append(items, conv(ch))
	}
	sort.Slice(items, func(i, j int) bool { return items[i].idx < items[j].idx })
	ans := make([]string, 0)
	for _, ch := range items {
		ans = append(ans, ch.str)
	}
	return strings.Join(ans, " ")
}

type item struct {
	str string
	idx int
}

func conv(s string) item {
	idx := 0
	i := len(s) - 1
	for ; i >= 0; i-- {
		if s[i] >= '0' && s[i] <= '9' {
			idx = idx*10 + int(s[i]) - int('0')
		} else {
			break
		}
	}

	return item{
		str: s[:i+1],
		idx: idx,
	}
}
