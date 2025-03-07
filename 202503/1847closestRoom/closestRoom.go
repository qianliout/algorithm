package main

import (
	"math"
	"sort"

	"github.com/emirpasic/gods/v2/trees/redblacktree"
)

func main() {

}

func closestRoom(rooms [][]int, queries [][]int) []int {
	// 按放假面积从大到小排序
	sort.Slice(rooms, func(i, j int) bool {
		if rooms[i][1] != rooms[j][1] {
			return rooms[i][1] > rooms[j][1]
		}
		return rooms[i][0] < rooms[j][0]
	})
	queries2 := make([]pair, len(queries))
	for i, ch := range queries {
		queries2[i] = pair{id: i, preferID: ch[0], minSize: ch[1]}
	}
	// 按查询的面积从大到小排序
	sort.Slice(queries2, func(i, j int) bool {
		return queries2[i].minSize >= queries2[j].minSize
	})
	ans := make([]int, len(queries))
	for i := range ans {
		ans[i] = -1
	}
	tree := redblacktree.New[int, struct{}]()
	j := 0
	for _, qu := range queries2 {
		for j < len(rooms) && qu.minSize <= rooms[j][1] {
			tree.Put(rooms[j][0], struct{}{})
			j++
		}
		// 查左边
		diff := math.MaxInt64
		if node, ok := tree.Floor(qu.preferID); ok {
			diff = abs(node.Key - qu.preferID)
			ans[qu.id] = node.Key
		}
		// 查右边
		if node, ok := tree.Ceiling(qu.preferID); ok {
			// 因为如果差值相同，就返回小的 ID，所以这里是 < diff
			if abs(qu.preferID-node.Key) < diff {
				ans[qu.id] = node.Key
			}
		}
	}
	return ans
}

type pair struct {
	id       int
	preferID int
	minSize  int
}

func abs(a int) int {
	if a <= 0 {
		return -a
	}
	return a
}
