package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	fmt.Println(minHeightShelves([][]int{{1, 1}, {2, 3}, {2, 3}, {1, 1}, {1, 1}, {1, 1}, {1, 2}}, 4))
	fmt.Println(minHeightShelves([][]int{{7, 3}, {8, 7}, {2, 7}, {2, 5}}, 10))
}

// 不可以用贪心，要用 dp
func minHeightShelves(books [][]int, shelfWidth int) int {
	n := len(books)
	var dfs func(i int) int
	mem := make([]int, n)
	for i := range mem {
		mem[i] = -1
	}
	dfs = func(i int) int {
		if i < 0 {
			return 0
		}
		if mem[i] >= 0 {
			return mem[i]
		}
		j := i
		w := 0
		h := 0
		res := math.MaxInt / 100
		for ; j >= 0; j-- {
			if books[j][0]+w > shelfWidth {
				break
			}
			w += books[j][0]
			h = max(h, books[j][1])
			res = min(res, dfs(j-1)+h)
		}

		mem[i] = res
		return res
	}

	return dfs(n - 1)
}

func defangIPaddr(address string) string {
	return strings.ReplaceAll(address, ".", "[.]")
}
