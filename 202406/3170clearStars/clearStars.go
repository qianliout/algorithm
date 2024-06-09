package main

import (
	"fmt"
)

func main() {
	fmt.Println(clearStars("d*o*"))
}

func clearStars2(s string) string {
	st := make([][]int, 26)
	del := make([]bool, len(s))
	for i, ch := range s {
		if ch != '*' {
			idx := int(ch) - 'a'
			st[idx] = append(st[idx], i)
			continue
		}
		del[i] = true
		// 这样写会有问题，p不会被更新
		for _, p := range st {
			if len(p) > 0 {
				del[p[len(p)-1]] = true
				p = p[:len(p)-1]
				break
			}
		}
	}
	ans := make([]byte, 0)
	for i, ch := range s {
		if !del[i] {
			ans = append(ans, byte(ch))
		}
	}
	return string(ans)
}

// 每次删除时要删除最右边的，才能保证字典序最小
func clearStars(s string) string {
	st := make([][]int, 26)
	del := make([]bool, len(s))
	for i, ch := range s {
		if ch != '*' {
			idx := int(ch) - 'a'
			st[idx] = append(st[idx], i)
			continue
		}
		del[i] = true
		for k := range st {
			p := st[k]
			if len(p) > 0 {
				del[p[len(p)-1]] = true
				p = p[:len(p)-1]
				st[k] = p
				break
			}
		}
	}
	ans := make([]byte, 0)
	for i, ch := range s {
		if !del[i] {
			ans = append(ans, byte(ch))
		}
	}
	return string(ans)
}
