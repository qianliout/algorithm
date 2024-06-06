package main

import (
	"fmt"
)

func main() {
	fmt.Println(backspaceCompare("y#fo##f", "y#f#o##f"))
}

func backspaceCompare(s string, t string) bool {
	return strip([]byte(s)) == strip([]byte(t))
}

func strip(s []byte) string {
	ss := make([]byte, 0)
	for _, ch := range s {
		if ch == '#' {
			if len(ss) > 0 {
				ss = ss[:len(ss)-1]
			}
			continue
		}
		ss = append(ss, ch)
	}
	return string(ss)
}
