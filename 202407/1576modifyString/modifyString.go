package main

import (
	"fmt"
)

func main() {
	fmt.Println(modifyString("?????"))
}

func modifyString(s string) string {
	ss := []byte(s)
	n := len(s)
	for i := 0; i < n; i++ {
		if ss[i] != '?' {
			continue
		}
		var a byte
		var b byte
		if i > 0 {
			a = ss[i-1]
		}
		if i < n-1 {
			b = ss[i+1]
		}
		ss[i] = get(a, b)
	}

	return string(ss)
}

func get(a, b byte) byte {
	for i := 'a'; i < 'z'; i++ {
		if byte(i) != a && byte(i) != b {
			return byte(i)
		}
	}
	return 'a'
}
