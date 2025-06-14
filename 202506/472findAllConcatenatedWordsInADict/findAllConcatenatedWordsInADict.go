package main

import (
	"fmt"
)

func main() {
	fmt.Println(findAllConcatenatedWordsInADict([]string{"cat", "cats", "catsdogcats", "dog", "dogcatsdog", "hippopotamuses", "rat", "ratcatdogcat"}))
}

func findAllConcatenatedWordsInADict(words []string) []string {
	// 先过滤空字符串并构建字典树
	trie := &Trie{Node: &Node{}}
	for _, word := range words {
		if len(word) > 0 {
			trie.Add(word)
		}
	}

	var result []string
	for _, word := range words {
		if len(word) == 0 {
			continue
		}
		// 检查是否能被拆分为至少两个单词
		if canFormByAtLeastTwoWords(trie, word) {
			result = append(result, word)
		}
	}
	return result
}

func canFormByAtLeastTwoWords(trie *Trie, word string) bool {
	cache := make(map[string]bool) // 缓存子串是否可分割
	return dfs(trie.Node, word, cache, false)
}

func dfs(node *Node, s string, cache map[string]bool, isFirstWord bool) bool {
	if len(s) == 0 {
		return isFirstWord // 只有已经找到至少一个单词时才返回true
	}

	// 检查缓存
	if val, exists := cache[s]; exists {
		return val
	}

	for i := 0; i < len(s); i++ {
		idx := s[i] - 'a'
		if node.Child[idx] == nil {
			break // 没有匹配的前缀，停止搜索
		}
		node = node.Child[idx]
		if node.End {
			// 如果是第一个单词，继续搜索剩余部分
			// 如果不是第一个单词，检查剩余部分是否能组成单词
			remaining := s[i+1:]
			if dfs(node, remaining, cache, true) {
				cache[s] = true
				return true
			}
		}
	}

	cache[s] = false
	return false
}

type Node struct {
	Child []*Node
	End   bool
}

type Trie struct {
	Node *Node
}

func (tr *Trie) Add(word string) {
	add(tr.Node, []byte(word), 0)
}

func (tr *Trie) Search(word string) bool {
	return search(tr.Node, []byte(word), 0)
}

func add(node *Node, word []byte, idx int) {
	if idx >= len(word) {
		node.End = true
		return
	}
	c := int(word[idx]) - int('a')
	if node.Child == nil {
		node.Child = make([]*Node, 26)
	}
	if node.Child[c] == nil {
		node.Child[c] = &Node{}
	}
	add(node.Child[c], word, idx+1)
}

func search(node *Node, word []byte, idx int) bool {
	if node == nil {
		return false
	}
	if idx >= len(word) {
		return node.End
	}
	c := int(word[idx]) - int('a')
	if node.Child == nil {
		return false
	}
	if node.Child[c] == nil {
		return false
	}
	return search(node.Child[c], word, idx+1)
}

/*
给你一个 不含重复 单词的字符串数组 words ，请你找出并返回 words 中的所有 连接词 。
连接词 定义为：一个完全由给定数组中的至少两个较短单词（不一定是不同的两个单词）组成的字符串
*/
