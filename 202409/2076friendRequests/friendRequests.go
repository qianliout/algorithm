package main

import (
	"fmt"
	. "outback/algorithm/common/unionfind"
)

func main() {
	fmt.Println(friendRequests(3, [][]int{{0, 1}}, [][]int{{1, 2}, {0, 2}}))
	fmt.Println(friendRequests(5, [][]int{{0, 1}, {1, 2}, {2, 3}}, [][]int{{0, 4}, {1, 2}, {3, 1}, {3, 4}}))
}

func friendRequests(n int, restrictions [][]int, requests [][]int) []bool {
	uf := NewRankUnionFind(n)
	ans := make([]bool, len(requests))
	for i, ch := range requests {
		x, y := ch[0], ch[1]
		// 已经是朋友
		if uf.IsConnected(x, y) {
			ans[i] = true
			continue
		}
		ans[i] = true
		// 验证看看能不能成为朋友
		for _, re := range restrictions {
			a, b := re[0], re[1]
			if (uf.IsConnected(a, x) && uf.IsConnected(b, y)) ||
				(uf.IsConnected(a, y) && uf.IsConnected(b, x)) {
				ans[i] = false
			}
		}
		// 如果能成为朋友，就放在一个并查集里
		if ans[i] {
			uf.Union(x, y)
		}
	}
	return ans
}

/*
使用并查集，每个集合内都是直接朋友或者间接朋友。
对于：x=requests[i][0],y=requests[i][1]
    如果x所处的集合与y所处的集合相等，这样就意味着x和y已经是朋友了
    如果x所处的集合与y所处的集合不等，我们需要判断x所处的集合与y所处的集合能否合并，我们可以遍历restrictions，具体如下：
        对于：p=restrictions[j][0],q=restrictions[j][1] 有以下两种情况不合符条件
            如果x与p在同一个集合 且 q与y在同一个集合
            如果y与p在同一个集合 且 q与x在同一个集合
        上面两种情况意味着，x所属集合与y所属集合存在着不能成为朋友的用户，所以这两个集合不能够合并，这样x和y就不能够成为朋友了👫。
*/
