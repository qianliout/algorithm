package main

import (
	"strings"
)

func main() {

}
func truncateSentence(s string, k int) string {
	// 只有单个空格
	split := strings.Split(s, " ")
	n := min(k, len(split))
	return strings.Join(split[:n], " ")
}
