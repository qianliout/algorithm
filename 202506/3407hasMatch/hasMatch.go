package main

import (
	"strings"
)

func main() {

}

func hasMatch(s string, p string) bool {

	pp := strings.Split(p, "*")
	if len(pp) != 2 {
		return false
	}
	fir, sec := pp[0], pp[1]
	a := strings.Index(s, fir)
	b := strings.LastIndex(s, sec)
	if a == -1 || b == -1 {
		return false
	}
	if a+len(fir) > b {
		return false
	}
	return true
}
