package main

import (
	"fmt"
)

func main() {

	str := fmt.Sprintf("%02d:%02d", 120, 0)
	fmt.Println(str)
	fmt.Println(maximumTime("2?:?0"))
}

func maximumTime(s string) string {
	for h := 23; h >= 0; h-- {
		for m := 59; m >= 0; m-- {
			str := fmt.Sprintf("%02d:%02d", h, m)
			find := true
			for i := 0; i < len(str); i++ {
				if s[i] != str[i] && byte(s[i]) != '?' {
					find = false
				}
			}
			if find {
				return str
			}
		}
	}
	return ""
}
