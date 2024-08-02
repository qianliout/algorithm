package main

import (
	"strings"
)

func main() {

}

func minimumBuckets(ham string) int {
	if ham == "H" || strings.HasPrefix(ham, "HH") || strings.HasSuffix(ham, "HH") ||
		strings.Contains(ham, "HHH") {
		return -1
	}
	next := strings.ReplaceAll(ham, "H.H", "")
	// 只有 H.H 这重情况下两个仓鼠共用一个盆，假如有4个这种情况，那么只就需要用4个盆，其他的仓鼠只能是一个仓鼠一个盆
	return (len(ham)-len(next))/3 + strings.Count(next, "H")
}
