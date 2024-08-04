package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(int('a'), int('A'))

}

func capitalizeTitle(title string) string {
	split := strings.Split(title, " ")
	words := make([]string, 0)
	for _, ch := range split {
		ch = strings.ToLower(ch)
		if len(ch) > 2 {
			ss := []byte(ch)
			ss[0] = byte(ss[0] - 32)
			ch = string(ss)
		}
		words = append(words, ch)
	}
	return strings.Join(words, " ")
}
