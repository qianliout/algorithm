package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(isValid("UuE6"))
}

func isValid2(word string) bool {
	word2 := strings.ToLower(word)
	if len(word) < 3 {
		return false
	}
	yc, fc := 0, 0
	for _, ch := range word2 {
		if !((ch >= '0' && ch <= '9') || (ch >= 'a' && ch <= 'z')) {
			return false
		}
		if ch >= 'a' && ch <= 'z' {
			if !yuny(byte(ch)) {
				fc++
			} else {
				yc++
			}
		}
	}
	return yc > 0 && fc > 0
}

func isValid(word string) bool {
	// word := strings.ToLower(word)
	if len(word) < 3 {
		return false
	}
	y := map[byte]bool{
		'a': true,
		'e': true,
		'i': true,
		'o': true,
		'u': true,
		'A': true,
		'E': true,
		'I': true,
		'O': true,
		'U': true,
	}

	yc, fc, num := 0, 0, 0
	for _, ch := range word {
		if ch >= '0' && ch <= '9' {
			num++
			continue
		}
		if !((ch >= 'a' && ch <= 'z') || ch >= 'A' && ch <= 'Z') {
			return false
		}
		if y[byte(ch)] {
			yc++
		} else {
			fc++
		}
	}
	return yc > 0 && fc > 0
}

func yuny(ch byte) bool {
	ss := "aeiou"
	return strings.Contains(ss, string(ch))
}

/*
有效单词 需要满足以下几个条件：

    至少 包含 3 个字符。
    由数字 0-9 和英文大小写字母组成。（不必包含所有这类字符。）
    至少 包含一个 元音字母 。
    至少 包含一个 辅音字母 。

给你一个字符串 word 。如果 word 是一个有效单词，则返回 true ，否则返回 false 。

注意：

    'a'、'e'、'i'、'o'、'u' 及其大写形式都属于 元音字母 。
*/
