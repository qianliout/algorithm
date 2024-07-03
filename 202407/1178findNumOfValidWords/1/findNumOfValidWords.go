package main

import (
	"fmt"
)

func main() {
	fmt.Println(findNumOfValidWords([]string{"aaaa", "asas", "able", "ability", "actt", "actor", "access"}, []string{"aboveyz", "abrodyz", "abslute", "absoryz", "actresz", "gaswxyz"}))
}

func findNumOfValidWords(words []string, puzzles []string) []int {
	ans := make([]int, len(puzzles))
	wordCnt := make(map[int]int)
	for _, wo := range words {
		wordCnt[gen1(wo)]++
	}
	for i, p := range puzzles {
		pop := gen2(p)
		res := 0
		for _, ch := range pop {
			res += wordCnt[ch]
		}
		ans[i] = res
	}
	return ans
}

// 根据一个单词生成二进制
func gen1(word string) int {
	ans := 0
	for _, ch := range word {
		ans |= 1 << (int(ch) - int('a'))
	}
	return ans
}

// 固定住第一位，变换其他位，所得到的可能的结果
func gen2(puz string) []int {
	ans := make([]int, 0)
	m := len(puz)
	first := int(puz[0]) - int('a')
	for i := 0; i < (1 << (m - 1)); i++ {
		u := 1 << first // 保留第一位
		for j := 1; j < m; j++ {
			if ((i >> (j - 1)) & 1) != 0 {
				u += 1 << (puz[j] - 'a')
			}
		}
		ans = append(ans, u)
	}

	return ans
}
