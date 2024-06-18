package main

import (
	"fmt"
)

func main() {
	fmt.Println(addMinimum("aaaaba"))
	fmt.Println(addMinimum2("aaaaba"))
}

func addMinimum(word string) int {
	n := len(word)
	ans := 0
	for i := 1; i < n; i++ {
		fmt.Println(word[i], word[i-1], int(word[i]-word[i-1]), int(word[i])-int(word[i-1]))
		ans += (int(word[i]) - int(word[i-1]) - 1 + 3) % 3
		// 这样写会出错，出错原因是,当word[i]='a',word[i-1]='b'时，直接相差会是255，因为word[i]当做 uint8了
		// 是非常容易出错的点
		// ans += (int(word[i]-word[i-1]) - 1 + 3) % 3
	}
	// 考虑首尾
	ans += int(word[0]-'a') + int('c'-word[n-1])
	return ans
}

func addMinimum2(s string) int {

	// ans := int(s[0]) - int(s[len(s)-1]) + 2
	ans := 0
	for i := 1; i < len(s); i++ {

		ans += (int(s[i]) - int(s[i-1]) + 2) % 3

	}

	return ans

}
