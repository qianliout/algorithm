package main

import (
	"fmt"
)

func main() {
	fmt.Println(gcdOfStrings("ABCABC", "ABC"))
	fmt.Println(gcdOfStrings("ABABAB", "ABAB"))
}

func gcdOfStrings(str1 string, str2 string) string {
	if str1+str2 != str2+str1 {
		return ""
	}
	c := he(len(str1), len(str2))
	return str1[:c]
}

func he(a, b int) int {
	for b != 0 {
		b, a = a%b, b
	}
	return a
}

func he2(a, b int) int {
	if b == 0 {
		return a
	}
	return he2(b, a%b)
}
