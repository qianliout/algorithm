package main

import (
	"fmt"
	"math/bits"
)

func main() {
	fmt.Println(numberOfSpecialChars("cCceDC"))
}

// 这种写法不可以 :cCceDC
func numberOfSpecialChars1(word string) int {
	cnt1 := make([]int, 26)
	cnt2 := make([]int, 26)
	n := len(word)
	for i := 0; i < n; i++ {
		a := word[i]
		if a >= 'a' && a <= 'z' {
			idx := int(a - 'a')
			cnt1[idx] = i + 1
		} else {
			idx := int(a - 'A')
			cnt2[idx] = i + 1
		}
	}
	ans := 0
	for i := 0; i < 26; i++ {
		if cnt1[i] > 0 && cnt2[i] > 0 && cnt1[i] < cnt2[i] {
			ans++
		}
	}
	return ans
}
func numberOfSpecialChars2(word string) int {
	cnt1 := make([]int, 26)
	cnt2 := make([]int, 26)
	n := len(word)
	for i := 0; i < n; i++ {
		a := word[i]
		if a >= 'a' && a <= 'z' {
			idx := int(a - 'a')
			cnt1[idx] = i + 1
		} else {
			idx := int(a - 'A')
			if cnt2[idx] == 0 {
				cnt2[idx] = i + 1
			}
		}
	}
	ans := 0
	for i := 0; i < 26; i++ {
		if cnt1[i] > 0 && cnt2[i] > 0 && cnt1[i] < cnt2[i] {
			ans++
		}
	}
	return ans
}

func numberOfSpecialChars(word string) int {

	cnt1 := 0
	cnt2 := 0
	n := len(word)
	for i := 0; i < n; i++ {
		a := word[i]
		if a >= 'a' && a <= 'z' {
			idx := int(a - 'a')
			cnt1 = cnt1 | (1 << idx)
		} else {
			idx := int(a - 'A')
			cnt1 = cnt1 | (1 << idx)
		}
	}
	return bits.OnesCount(uint(cnt1 & cnt2))
}
