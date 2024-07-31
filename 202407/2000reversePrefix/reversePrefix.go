package main

import (
	"strings"
)

func main() {

}

func reversePrefix(word string, ch byte) string {
	ss := []byte(word)
	idx := strings.Index(word, string(ch))
	if idx < 0 {
		return word
	}
	le, ri := 0, idx

	for le < ri {
		ss[le], ss[ri] = ss[ri], ss[le]
		le++
		ri--
	}

	return string(ss)
}
