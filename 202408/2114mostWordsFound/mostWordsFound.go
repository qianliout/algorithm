package main

import (
	"strings"
)

func main() {

}

func mostWordsFound(sentences []string) int {
	ans := 0
	for _, ch := range sentences {
		split := strings.Split(ch, " ")
		ans = max(ans, len(split))
	}
	return ans
}
