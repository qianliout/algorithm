package main

import (
	"sort"
	"strings"
)

func main() {

}

func arrangeWords(text string) string {
	strs := strings.Split(text, " ")
	strs[0] = strings.ToLower(strs[0])
	sort.SliceStable(strs, func(i, j int) bool { return len(strs[i]) < len(strs[j]) })

	first := []byte(strs[0])
	up := strings.ToUpper(string(first[0]))
	first[0] = byte(up[0])
	strs[0] = string(first)

	s := strings.Join(strs, " ")
	return s
}
