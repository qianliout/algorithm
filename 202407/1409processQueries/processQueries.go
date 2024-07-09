package main

import (
	"fmt"
)

func main() {
	fmt.Println(processQueries([]int{3, 1, 2, 1}, 5))
}

// 数据量不大，模拟就行
func processQueries1(queries []int, m int) []int {

	findIndexArr := make([]int, m+1)
	for i := 0; i < m; i++ {
		findIndexArr[i+1] = i
	}
	idx := 0
	ans := make([]int, len(queries))
	for j, ch := range queries {
		cur := findIndexArr[ch]
		ans[idx] = cur
		idx++
		for i := 0; i < len(findIndexArr); i++ {
			if findIndexArr[i] < cur {
				findIndexArr[i]++
			}
		}
		findIndexArr[queries[j]] = 0
	}

	return ans
}

func processQueries(queries []int, m int) []int {
	pairs := make([]pair, m)
	for i := 1; i <= m; i++ {
		pairs[i-1] = pair{value: i, index: i - 1}
	}
	ans := make([]int, 0)
	for _, ch := range queries {
		for j := m - 1; j >= 0; j-- {
			if pairs[j].value == ch {
				ans = append(ans, pairs[j].index)
				break
			}
		}
		idx := ans[len(ans)-1]
		for j := m - 1; j >= 0; j-- {
			if pairs[j].index < idx {
				pairs[j].index++
			} else if pairs[j].index == idx {
				pairs[j].index = 0
			}
		}
	}

	return ans
}

type pair struct {
	value int
	index int
}
