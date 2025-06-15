package main

import (
	"fmt"
)

func main() {
	// 测试用例1：基本测试
	words1 := []string{"cat", "cats", "catsdogcats", "dog", "dogcatsdog", "hippopotamuses", "rat", "ratcatdogcat"}
	fmt.Println("测试1:", findAllConcatenatedWordsInADict(words1))

	// 测试用例2：更复杂的测试
	words2 := []string{"cat", "dog", "catdog", "catcat", "dogdog", "catdogcat", "dogcatdog", "catcatdog", "dogdogcat"}
	fmt.Println("测试2:", findAllConcatenatedWordsInADict(words2))
}

func findAllConcatenatedWordsInADict(words []string) []string {
	// 构建 Trie 树
	trie := &Trie{Node: &Node{}}
	for _, word := range words {
		trie.Add(word)
	}

	result := make([]string, 0)

	// 对每个单词检查是否为连接词
	for _, word := range words {
		// 为每个单词创建独立的缓存
		cache := make(map[int]bool)
		if canFormByConcatenation(trie, word, word, cache) {
			result = append(result, word)
		}
	}

	return result
}

// DFS 检查单词是否可以由其他单词连接而成
// trie: Trie 树
// word: 要检查的单词
// originalWord: 原始单词（用于避免匹配自己）
// cache: 缓存，key为位置，value为从该位置开始能否构成连接词
func canFormByConcatenation(trie *Trie, word string, originalWord string, cache map[int]bool) bool {
	return dfs(trie, word, originalWord, 0, 0, cache)
}

// DFS 搜索函数（带缓存优化）
// trie: Trie 树
// word: 当前要匹配的单词
// originalWord: 原始单词
// pos: 当前位置
// count: 已匹配的单词数量
// cache: 缓存
func dfs(trie *Trie, word string, originalWord string, pos int, count int, cache map[int]bool) bool {
	// 如果已经遍历完整个单词
	if pos == len(word) {
		// 必须由至少两个单词组成
		return count >= 2
	}

	// 检查缓存（只有count > 0时才使用缓存，避免自匹配问题）
	if count > 0 {
		if result, exists := cache[pos]; exists {
			return result
		}
	}

	result := false

	// 从当前位置开始，在 Trie 中查找可能的单词
	node := trie.Node
	for i := pos; i < len(word); i++ {
		c := int(word[i]) - int('a')

		// 如果没有对应的子节点，停止搜索
		if node.Child == nil || node.Child[c] == nil {
			break
		}

		node = node.Child[c]

		// 如果找到一个完整的单词
		if node.End {
			currentWord := word[pos : i+1]
			// 避免单词匹配自己（只有在count=0且匹配整个原单词时才是匹配自己）
			if count == 0 && currentWord == originalWord {
				continue
			}

			// 递归检查剩余部分
			if dfs(trie, word, originalWord, i+1, count+1, cache) {
				result = true
				break
			}
		}
	}

	// 将结果存入缓存（只有count > 0时才缓存）
	if count > 0 {
		cache[pos] = result
	}
	return result
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
