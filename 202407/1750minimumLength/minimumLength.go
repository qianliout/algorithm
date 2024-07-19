package main

import (
	"fmt"
)

func main() {
	fmt.Println(minimumLength("cabaabac"))
	fmt.Println(minimumLength("ca"))
	fmt.Println(minimumLength("bbbbbbbbbbbbbbbbbbbbbbbbbbbabbbbbbbbbbbbbbbccbcbcbccbbabbb"))
	fmt.Println(minimumLength("bbbbbbbbbbbbbbbbbbb"))
}

func minimumLength1(s string) int {
	i, j := 0, len(s)-1
	n := len(s)
	for i < j {
		if s[i] != s[j] {
			break
		}
		co := s[i]
		k, m := i, j
		for ; k <= j && k >= 0 && k+1 < n; k++ {
			if s[k+1] != co {
				break
			}
		}
		for ; m >= i && m-1 >= 1 && m <= n-1; m-- {
			if s[m-1] != co {
				break
			}
		}
		i, j = k+1, m-1
	}
	if i > j {
		return 0
	}
	return j - i + 1
}

// 更简洁的做法
func minimumLength(s string) int {
	i, j := 0, len(s)-1
	for i < j && s[i] == s[j] {
		co := s[i]
		for i <= j && s[i] == co {
			i++
		}
		for i <= j && s[j] == co {
			j--
		}
	}
	return j - i + 1
}
