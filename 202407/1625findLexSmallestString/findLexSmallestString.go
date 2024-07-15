package main

import (
	"fmt"
)

func main() {
	fmt.Println(findLexSmallestString("5525", 9, 2))
}

// BFS 暴力
func findLexSmallestString(s string, a int, b int) string {
	ans := s
	queue := []string{s}
	visit := map[string]bool{s: true}
	for len(queue) > 0 {
		fist := queue[0]
		if fist < ans {
			ans = fist
		}
		queue = queue[1:]
		for _, nex := range cal(fist, a, b) {
			if !visit[nex] {
				queue = append(queue, nex)
				visit[nex] = true
			}
		}
	}
	return ans
}

// 累加：将  a 加到 s 中所有下标为奇数的元素上（下标从 0 开始）。数字一旦超过 9 就会变成 0，如此循环往复。例如，s = "3456" 且 a = 5，则执行此操作后 s 变成 "3951"。
// 轮转：将 s 向右轮转 b 位。例如，s = "3456" 且 b = 1，则执行此操作后 s 变成 "6345"。
func cal(s string, a, b int) []string {
	ans := make([]string, 0)
	ss := []byte(s)

	for i := 1; i < len(s); i += 2 {
		ch := byte((int(ss[i])-'0'+a)%10 + '0')
		ss[i] = ch
	}
	ans = append(ans, string(ss))
	ss = []byte(s)
	pre := ss[:b]
	ss = append(ss[b:], pre...)
	ans = append(ans, string(ss))
	return ans
}
