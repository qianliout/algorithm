package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	// fmt.Println(longestSubsequenceRepeatedK("letsleetcode", 2))
	fmt.Println(longestSubsequenceRepeatedK("babbbaabbabab", 2)) // babab
}

func longestSubsequenceRepeatedK(s string, k int) string {
	// 统计每个字符出现的次数
	num := make(map[byte]int)
	for _, ch := range s {
		num[byte(ch)]++
	}

	// 构建 hot 字符串
	var hot strings.Builder
	sortK := sortedKeys(num)

	for _, ch := range sortK {
		hot.WriteString(strings.Repeat(string(ch), num[ch]/k))
	}
	ss := []byte(hot.String())
	// 尝试所有可能的子序列
	for i := len(ss); i > 0; i-- {
		ans := ""
		p := permutations(ss, i)
		for _, ch := range p {
			if isSubsequenceRepeatedK(s, ch, k) {
				// 这样直接返回也是对的
				// return ch
				if ans == "" || ans < ch {
					ans = ch
				}
			}
		}
		// 为啥可以提前返回呢
		// 1：i 是从大到小，如果这一次找到了，那么其他的 i 值的长度比这一次少，一定不是正确的答案
		// 一定要提前返回，不然就会超时
		if ans != "" {
			return ans
		}

	}
	return ""
}

// sortedKeys 返回 map 的键按字典顺序排序后的切片
func sortedKeys(m map[byte]int) []byte {
	keys := make([]byte, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool {
		return keys[i] > keys[j]
	})
	return keys
}

// 从arr 中选出 m 个数的组合
func permutations(arr []byte, m int) []string {
	n := len(arr)
	ans := make([]string, 0)
	visit := make([]bool, n)
	var dfs func(path []byte)
	dfs = func(path []byte) {
		if len(path) >= m {
			ans = append(ans, string(path))
			return
		}
		for i := 0; i < n; i++ {
			if visit[i] {
				continue
			}
			if i > 0 && !visit[i-1] && arr[i] == arr[i-1] {
				continue
			}
			visit[i] = true
			path = append(path, arr[i])
			dfs(path)
			path = path[:len(path)-1]
			visit[i] = false
		}
	}
	dfs([]byte{})
	return ans
}

// isSubsequenceRepeatedK 检查 word 是否在 s 中重复出现 k 次
func isSubsequenceRepeatedK(s, word string, k int) bool {
	// count := strings.Count(s, word)
	// return count >= k
	ss := strings.SplitAfter(s, "")
	wordRepeated := strings.Repeat(word, k)
	j := 0
	for _, c := range ss {
		if j < len(wordRepeated) && c[0] == wordRepeated[j] {
			j++
		}
	}
	return j == len(wordRepeated)
}
