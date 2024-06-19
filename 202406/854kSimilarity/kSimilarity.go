package main

import (
	"fmt"
)

func main() {
	fmt.Println(kSimilarity("ab", "ba"))                                     // 1
	fmt.Println(kSimilarity("abc", "bca"))                                   // 2
	fmt.Println(kSimilarity("abac", "baca"))                                 // 2
	fmt.Println(kSimilarity("bccaba", "abacbc"))                             // 3
	fmt.Println(kSimilarity("abcdeabcdeabcdeabcde", "aaaabbbbccccddddeeee")) // 8
}

func kSimilarity(s1 string, s2 string) int {
	// mx := math.MaxInt / 2
	// res := dfs1([]byte(s1), []byte(s2), 0)
	// return res
	if s1 == s2 {
		return 0
	}
	return bfs([]byte(s1), []byte(s2))
}

func dfs1(s1, s2 []byte, start int) int {
	if start >= len(s1)-1 {
		return 0
	}
	res := 0
	for i := start; i < len(s1); i++ {
		if s1[i] == s2[i] {
			continue
		}
		for j := i + 1; j < len(s2); j++ {
			if s2[j] == s1[i] && s2[j] != s1[j] {
				s2[i], s2[j] = s2[j], s2[i]
				res = dfs1(s1, s2, i+1) + 1
				// s2[i], s2[j] = s2[j], s2[i]
			}
		}
	}
	return res
}

func bfs(s1 []byte, s2 []byte) int {
	queue := make([]pair, 0)

	queue = append(queue, pair{s2: s2, cost: 0})
	vis := make(map[string]bool)
	vis[string(s2)] = true
	ans := 0
	for len(queue) > 0 {
		lev := make([]pair, 0)
		for _, fir := range queue {
			if string(fir.s2) == string(s1) {
				return ans
			}

			for i := fir.start; i < len(s1); i++ {
				// for i := 0; i < len(s1); i++ {
				if s1[i] == s2[i] {
					continue
				}
				for j := i + 1; j < len(s2); j++ {
					if s2[j] == s1[i] && s2[j] != s1[j] {
						s2[i], s2[j] = s2[j], s2[i]
						if vis[string(s2)] {
							continue
						}
						vis[string(s2)] = true
						lev = append(lev, pair{
							s2:    append([]byte{}, s2...),
							cost:  fir.cost + 1,
							start: i,
						})
						s2[i], s2[j] = s2[j], s2[i]
					}
				}
			}
		}
		if len(lev) > 0 {
			ans++
		}
		lev = queue
	}
	return 0
}

func next(s1, s2 []byte) {

}

type pair struct {
	s2    []byte
	cost  int
	start int // 从那一步开始检测
}

func dfs(s1, s2 []byte, start int, cur int, mx *int) int {
	if start >= len(s2) {
		*mx = min(*mx, cur)
		return *mx
	}
	if cur >= *mx {
		return *mx
	}
	if start == len(s1)-1 {
		*mx = min(*mx, cur)
		return *mx
	}
	for i := start; i < len(s1); i++ {
		if s1[i] != s2[i] {
			for j := i + 1; j < len(s1); j++ {
				if s2[j] == s1[i] && s2[j] != s1[j] {
					s2[i], s2[j] = s2[j], s2[i]
					dfs(s1, s2, i+1, cur+1, mx)
					s2[i], s2[j] = s2[j], s2[i]
					// 这一步只是剪枝
					if s2[i] == s1[j] {
						break
					}
				}
			}
			return cur
		}
	}
	*mx = min(*mx, cur)
	return *mx
}
