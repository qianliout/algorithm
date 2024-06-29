package main

import (
	"fmt"
)

func main() {
	fmt.Println(decodeMessage("the quick brown fox jumps over the lazy dog", "vkbs bs t suepuv"))
}

func decodeMessage(key string, message string) string {
	abc := make(map[byte]byte)
	used := 'a'
	vis := make(map[byte]bool)
	for _, ch := range key {
		if ch == ' ' {
			continue
		}
		if vis[byte(ch)] {
			continue
		}
		vis[byte(ch)] = true
		abc[byte(ch)] = byte(used)
		used++
	}
	ans := make([]byte, len(message))

	for i, ch := range message {
		if ch == ' ' {
			ans[i] = ' '
			continue
		}
		ans[i] = abc[byte(ch)]
	}

	return string(ans)
	// ans := make([]byte, 0)
	// for _, ch := range []byte(message) {
	// 	if va, ok := abc[ch]; ok {
	// 		ans = append(ans, va)
	// 	} else {
	// 		ans = append(ans, ' ')
	// 	}
	//
	// }
	// return string(ans)
}
