package main

import (
	"strings"
)

func main() {

}

func findOcurrences(text string, first string, second string) []string {
	split := strings.Split(text, " ")
	ans := make([]string, 0)
	for i := 2; i < len(split); i++ {
		if split[i-2] == first && split[i-1] == second {
			ans = append(ans, split[i])
		}
	}
	return ans
}
