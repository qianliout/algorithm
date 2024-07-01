package main

import (
	"fmt"
	"strings"
)

func main() {
	// fmt.Println(removeDigit("123", 3))
	// fmt.Println(removeDigit("551", '5'))
	fmt.Println(removeDigit("15454", '4'))
	fmt.Println('4', '5')
}

func removeDigit(number string, digit byte) string {
	cnt := strings.Count(number, string(digit))
	if cnt == 0 {
		return number
	}
	if cnt == 1 {
		return strings.ReplaceAll(number, string(digit), "")
	}
	rem := -1
	for i := 0; i < len(number); i++ {
		if number[i] == digit {
			if cnt == 1 {
				rem = i
				break
			}
			if i+1 < len(number) && int(number[i+1]) > int(digit) {
				rem = i
				break
			}
			cnt--
		}
	}
	ss := []byte(number)
	ans := append([]byte{}, ss[:rem]...)
	ans = append(ans, ss[rem+1:]...)
	return string(ans)
}
