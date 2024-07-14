package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(reorderSpaces("  this   is  a sentence "))
	fmt.Println(reorderSpaces(" a "))
}

func reorderSpaces(text string) string {
	spaceCnt := 0
	ss := make([]string, 0)
	for _, ch := range text {
		if ch == ' ' {
			spaceCnt++
		}
	}
	split := strings.Split(text, " ")
	for _, ch := range split {
		if ch != "" {
			ss = append(ss, ch)
		}
	}
	if len(ss) == 0 {
		return text
	}
	if len(ss) == 1 {
		return ss[0] + strings.Repeat(" ", spaceCnt)
	}

	a := spaceCnt / (len(ss) - 1)
	sub := spaceCnt % (len(ss) - 1)
	spce := strings.Repeat(" ", a)
	res := strings.Join(ss, spce)
	res = res + strings.Repeat(" ", sub)
	return res
}
