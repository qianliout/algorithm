package main

import (
	"fmt"
)

func main() {

	fmt.Println(smallestNumber("IIIDIDDD"))
}

func smallestNumber(pattern string) string {
	n := len(pattern)
	ans := make([]byte, n+1)
	for i := 0; i < n+1; i++ {
		ans[i] = byte(i + '1')
	}
	i := 0
	for i < n {
		if pattern[i] == 'I' {
			i += 1
			continue
		}
		i0 := i
		i++
		for i < n && pattern[i] == 'D' {
			i += 1
		}
		reverse(ans[i0 : i+1])
		// x := reverse(ans[i0 : i+1])

		// for j := i0; j <= i; j++ {
		// 	ans[j] = x[j-i0]
		// }
	}
	return string(ans)
}

func reverse(ans []byte) []byte {
	// ans := append([]byte{}, a...)
	l, r := 0, len(ans)-1
	for l < r {
		ans[l], ans[r] = ans[r], ans[l]
		l++
		r--
	}
	return ans
}
