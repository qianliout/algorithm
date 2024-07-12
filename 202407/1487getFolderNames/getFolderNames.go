package main

import (
	"fmt"
)

func main() {
	fmt.Println(getFolderNames([]string{"gta", "gta(1)", "gta", "gta", "avalon"}))
	fmt.Println(getFolderNames([]string{"kaido", "kaido(1)", "kaido", "kaido(1)", "kaido(2)"}))
}

func getFolderNames(names []string) []string {
	ans := make([]string, 0)
	exist := make(map[string]int)
	for _, ch := range names {
		if exist[ch] == 0 {
			exist[ch] = 1
			ans = append(ans, ch)
			continue
		}
		for i := exist[ch]; ; i++ {
			key := fmt.Sprintf("%s(%d)", ch, i)
			if exist[key] == 0 {
				ans = append(ans, key)
				exist[key] = 1
				// 这一步是重点，才不会导致超时
				exist[ch] = i + 1
				break
			}
		}
	}

	return ans
}
