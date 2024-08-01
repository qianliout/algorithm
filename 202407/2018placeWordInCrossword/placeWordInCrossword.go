package main

import (
	"fmt"
)

func main() {
	fmt.Println(placeWordInCrossword([][]byte{[]byte("# #"), []byte("  #"), []byte("# c")}, "ca"))
	fmt.Println(placeWordInCrossword([][]byte{[]byte(" #a"), []byte(" #c"), []byte(" #a")}, "ac"))
	fmt.Println(placeWordInCrossword([][]byte{[]byte("  "), []byte("  ")}, "a"))
}

/*
该单词不占据任何 '#' 对应的格子。
每个字母对应的格子要么是 ' ' （空格）要么与 board 中已有字母 匹配 。
如果单词是 水平 放置的，那么该单词左边和右边 相邻 格子不能为 ' ' 或小写英文字母。
如果单词是 竖直 放置的，那么该单词上边和下边 相邻 格子不能为 ' ' 或小写英文字母。
*/

func placeWordInCrossword(board [][]byte, word string) bool {
	m, n, k := len(board), len(board[0]), len(word)

	// 横着匹配
	for _, row := range board {
		for j := 0; j < n; j++ { // 这里一定不能用 for range 的用法，因下面改了 j 的值
			if row[j] == '#' {
				continue
			}
			j0, ok1, ok2 := j, true, true
			for ; j < n && row[j] != '#'; j++ {
				// 相邻不能有空格，所以长度只要超过了 k就是不行的
				if j-j0 >= k || (row[j] != ' ' && row[j] != word[j-j0]) {
					ok1 = false // 正序不能放
				}
				// 倒着匹配时的下标计算：k-1 -(j-j0)
				if j-j0 >= k || (row[j] != ' ' && row[j] != word[k-1-j+j0]) {
					ok2 = false // 倒序不能放
				}
			}
			// 题目要求是刚好放的下
			if (ok1 || ok2) && j-j0 == k {
				return true
			}
		}
	}

	//  竖着匹配
	for j := 0; j < n; j++ {
		for i := 0; i < m; i++ {
			if board[i][j] == '#' {
				continue
			}
			i0, ok1, ok2 := i, true, true
			for ; i < m && board[i][j] != '#'; i++ {
				if i-i0 >= k || (board[i][j] != ' ' && board[i][j] != word[i-i0]) {
					ok1 = false // 正序不能放
				}
				if i-i0 >= k || (board[i][j] != ' ' && board[i][j] != word[k-1-i+i0]) {
					ok2 = false // 倒序不能放
				}
			}
			if (ok1 || ok2) && i-i0 == k {
				return true
			}
		}
	}

	return false
}
