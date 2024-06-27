package main

import (
	"sort"
)

func main() {

}

/*
先假设所有人都派到 A 市 去需要花费用sumA 值
现需要把这里面一半的人派到 B 去，那应该排那一部分人去 B 呢，
当然是排和A 差的少的去，比如现在两个人分别是 （10，20），（10，30），那当然是排第一个人到 B 去了
同理现在有两个人(30,40) (20，50)，那还是应该排第一个到 B 去

*/

func twoCitySchedCost(costs [][]int) int {
	n := len(costs)
	sub := make([]int, n)
	sumA := 0
	for i, ch := range costs {
		sub[i] = ch[1] - ch[0]
		sumA += ch[0]
	}
	sort.Ints(sub)
	for i := 0; i < n/2; i++ {
		sumA -= sub[i]
	}
	return sumA
}
