package main

import (
	"strings"
)

func main() {

}

func isValid(s string) bool {
	for s != "" {
		n := strings.ReplaceAll(s, "abc", "")
		if n == s && n != "" {
			return false
		}
		s = n
	}

	return true
}
