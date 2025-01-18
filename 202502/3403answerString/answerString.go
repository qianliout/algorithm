package main

import (
	"fmt"
)

func main() {
	fmt.Println(answerString("dbca", 2))
	fmt.Println(answerString("usjxvmwnyethcnkwpdoqvbieqsgaycsdegmqwrkawlzhsxbawuetxvcykpdkxkbxvrwljkgqxotqpzyctxqudkgunlxlnoobndmgvhkyhbzetgxldvnsicfzjlrzxzbsgxtupgzfnczzywvlytlosuoxhypczdwqthhlczjepxbgimquwzlfnaktcyzmevspbsmgfgpddwnmxqcjypdbrpzqeijgmrshrpgonbzqxwudnbabyuymhxwsxhntnvpwxbfuqnnvjcshzlcxfavhefoqjzqnnthwpyeutifdiskqxmixarttwaqxhafmdbkgjmhhniytvkipuqkouihvdcopnbvfieeyelnszeqalaascouqjrraepuestggtuxxgjzreicyqutslvqcqcfpbyzetvpcseapnjrjpetjgdfvddgdrmqjrsfoexgakukqdaidzyqwjnrhkddprzppwurwejxlxguspompkmdmvybyltsmcxwd", 250))
}

// 这种方法没有办法 cache
func answerString2(word string, numFriends int) string {
	var dfs func(w []byte, n int) string
	dfs = func(w []byte, n int) string {
		m := len(w)
		if m < n {
			return ""
		}
		if n == 1 {
			return string(w)
		}

		ans := ""
		for k := m - 1; k >= 0; k-- {
			last := string(w[k:])
			pre := dfs(w[:k], n-1)
			if pre == "" {
				continue
			}
			a := max(pre, last)
			ans = max(ans, a)
		}
		return ans
	}
	w := []byte(word)
	ans := dfs(w, numFriends)
	return ans
}

// 这样写还是会超时
func answerString3(word string, numFriends int) string {
	mem := make([]map[string]string, numFriends+1)
	for i := range mem {
		mem[i] = make(map[string]string)
	}

	var dfs func(w string, n int) string

	dfs = func(w string, n int) string {
		m := len(w)
		if m < n {
			return ""
		}
		if n == 1 {
			return w
		}
		if mem[n][w] != "" {
			return mem[n][w]
		}

		ans := ""
		for k := m - 1; k >= 0; k-- {
			last := w[k:]
			pre := dfs(w[:k], n-1)
			if pre == "" {
				continue
			}
			a := max(pre, last)
			ans = max(ans, a)
		}
		mem[n][w] = ans
		return ans
	}
	ans := dfs(word, numFriends)
	return ans
}

// 这样写还是会超时
func answerString4(word string, numFriends int) string {
	mem := make([][]string, numFriends+1)
	mem2 := make([][]int, numFriends+1)
	for i := range mem {
		mem[i] = make([]string, len(word)+1)
		mem2[i] = make([]int, len(word)+1)
	}
	w := []byte(word)
	var dfs func(end int, n int) string

	dfs = func(end int, n int) string {
		if end < n {
			return ""
		}
		if n == 1 {
			return word[:end]
		}
		if mem2[n][end] > 0 {
			return mem[n][end]
		}

		ans := ""
		for k := end - 1; k >= 0; k-- {
			last := string(w[k:end])
			pre := dfs(k, n-1)
			if pre == "" {
				continue
			}
			a := max(pre, last)
			ans = max(ans, a)
		}
		mem[n][end] = ans
		mem2[n][end] = 1
		return ans
	}
	ans := dfs(len(w), numFriends)
	return ans
}

func answerString(word string, k int) string {
	n := len(word)
	if n < k {
		return ""
	}
	if k == 1 {
		return word
	}
	ans := ""
	// 枚举左端点，如果 pre 都相同，那么长度更长的字典序最大，比如 abcd比abc 更大
	// 又因为至少有k个子串,所以单个子串的长度最多 n-k+1
	for i := 0; i < n; i++ {
		ans = max(ans, word[i:min(n, i+n-k+1)])
	}
	return ans
}
