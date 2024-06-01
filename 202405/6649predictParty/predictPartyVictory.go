package main

import (
	"fmt"
)

func main() {
	fmt.Println(predictPartyVictory("RD"))
	fmt.Println(predictPartyVictory("RDD"))
}

/*
禁止一名参议员的权利：参议员可以让另一位参议员在这一轮和随后的几轮中丧失 所有的权利 。
宣布胜利：如果参议员发现有权利投票的参议员都是 同一个阵营的 ，他可以宣布胜利并决定在游戏中的有关变化。
*/
func predictPartyVictory(senate string) string {
	r, d := make([]int, 0), make([]int, 0)
	for i, ch := range senate {
		if ch == 'D' {
			d = append(d, i)
		} else {
			r = append(r, i)
		}
	}
	for len(r) > 0 && len(d) > 0 {
		if r[0] < d[0] {
			r = append(r, r[0]+len(senate))
		} else {
			d = append(d, d[0]+len(senate))
		}
		r = r[1:]
		d = d[1:]
	}
	if len(r) > 0 {
		return "Radiant"
	}
	return "Dire"
}
