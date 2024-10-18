package main

import (
	"fmt"
	"sort"
	"strings"
	"unicode"
)

func main() {
	fmt.Println(reorderLogFiles([]string{"dig1 8 1 5 1", "let1 art can", "dig2 3 6", "let2 own kit dig", "let3 art zero"}))
	// "let1 art can","let3 art zero","let2 own kit dig","dig1 8 1 5 1","dig2 3 6"
}

func reorderLogFiles(logs []string) []string {
	n := len(logs)
	nodes := make([]Node, n)

	for i, ch := range logs {
		split := strings.Split(ch, " ")
		no := Node{
			Fir: split[0],
			Sec: strings.Join(split[1:], " "),
		}
		no.FirN = Cha(no.Sec)
		// 字母日志 在内容不同时，忽略标识符后，按内容字母顺序排序；在内容相同时，按标识符排序。
		nodes[i] = no
	}

	sort.SliceStable(nodes, func(i, j int) bool {
		if nodes[i].FirN == 2 && nodes[j].FirN == 2 {
			// 一定要加这一句，才能是稳定排序
			return i < j
		}
		if nodes[i].FirN == 1 && nodes[j].FirN == 2 {
			return true
		} else if nodes[i].FirN == 2 && nodes[j].FirN == 1 {
			return false
		}

		if nodes[i].Sec != nodes[j].Sec {
			return nodes[i].Sec < nodes[j].Sec
		}

		return nodes[i].Fir < nodes[j].Fir
	})
	ans := make([]string, n)
	for i := range nodes {
		ans[i] = fmt.Sprintf("%s %s", nodes[i].Fir, nodes[i].Sec)
	}
	return ans
}

func Cha(sec string) int {
	for _, a := range sec {
		if a >= 'a' && a <= 'z' {
			return 1
		}
	}

	return 2
}

type Node struct {
	Fir  string
	FirN int
	Sec  string
}

func reorderLogFiles2(logs []string) []string {
	sort.SliceStable(logs, func(i, j int) bool {
		a, b := logs[i], logs[j]
		aDig := unicode.IsDigit(rune(a[len(a)-1]))
		bDig := unicode.IsDigit(rune(b[len(b)-1]))
		if aDig && bDig {
			return i < j
		} else if bDig {
			return true
		} else if aDig {
			return false
		}
		aSp := strings.SplitN(a, " ", 2)
		bSp := strings.SplitN(b, " ", 2)
		if aSp[1] == bSp[1] {
			return aSp[0] < bSp[0]
		}
		return aSp[1] < bSp[1]
	})
	return logs
}
