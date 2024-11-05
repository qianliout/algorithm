package main

import (
	"fmt"
	"sort"
	"strings"
)

func longestSubsequenceRepeatedK(s string, k int) string {
	// 统计每个字符出现的次数
	num := make(map[rune]int)
	for _, ch := range s {
		num[ch]++
	}

	// 构建 hot 字符串
	var hot strings.Builder
	for _, ch := range sortedKeys(num) {
		hot.WriteString(strings.Repeat(string(ch), num[ch]/k))
	}

	// 尝试所有可能的子序列
	for i := len(hot.String()); i > 0; i-- {
		for _, item := range permutations([]rune(hot.String()), i) {
			word := string(item)
			if isSubsequenceRepeatedK(s, word, k) {
				return word
			}
		}
	}
	return ""
}

// sortedKeys 返回 map 的键按字典顺序排序后的切片
func sortedKeys(m map[rune]int) []rune {
	keys := make([]rune, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool {
		return keys[i] > keys[j]
	})
	return keys
}

// permutations 生成所有长度为 n 的排列
func permutations(arr []rune, n int) [][]rune {
	if n == 0 {
		return [][]rune{{}}
	}
	var result [][]rune
	permute(arr, 0, n, &result)
	return result
}

// permute 递归生成排列
func permute(arr []rune, start, n int, result *[][]rune) {
	if start == n {
		*result = append(*result, append([]rune(nil), arr[:n]...))
		return
	}
	used := make(map[rune]bool)
	for i := start; i < len(arr); i++ {
		if used[arr[i]] {
			continue
		}
		used[arr[i]] = true
		arr[start], arr[i] = arr[i], arr[start]
		permute(arr, start+1, n, result)
		arr[start], arr[i] = arr[i], arr[start]
	}
}

// isSubsequenceRepeatedK 检查 word 是否在 s 中重复出现 k 次
func isSubsequenceRepeatedK(s, word string, k int) bool {
	wordRepeated := strings.Repeat(word, k)
	j := 0
	for _, c := range s {
		if j < len(wordRepeated) && c == rune(wordRepeated[j]) {
			j++
		}
	}
	return j == len(wordRepeated)
}

func main() {
	s := "babbbaabbabab"
	k := 2
	fmt.Println(longestSubsequenceRepeatedK(s, k)) // 输出: "babab"
}
