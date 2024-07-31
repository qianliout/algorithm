package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(numberOfWeakCharacters([][]int{{2, 2}, {3, 3}}))
	fmt.Println(numberOfWeakCharacters([][]int{{7, 7}, {1, 2}, {9, 7}, {7, 3}, {3, 10}, {9, 8}, {8, 10}, {4, 3}, {1, 5}, {1, 5}}))
}

func numberOfWeakCharacters(properties [][]int) int {
	sort.Slice(properties, func(i, j int) bool {
		if properties[i][0] != properties[j][0] {
			return properties[i][0] > properties[j][0]
		}
		return properties[i][1] <= properties[j][1]
	})
	ans := 0
	n := len(properties)
	mx := properties[0][1] // 防御的最大值
	for i := 1; i < n; i++ {
		if properties[i][1] < mx {
			ans++
		}
		mx = max(mx, properties[i][1])
	}
	return ans
}
