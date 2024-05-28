package main

import (
	"fmt"
)

func main() {
	// fmt.Println(maximumInvitations([]int{2, 2, 1, 2}))
	// fmt.Println(maximumInvitations([]int{1, 2, 0}))
	// fmt.Println(maximumInvitations([]int{3, 0, 1, 4, 1}))
	fmt.Println(maximumInvitations([]int{1, 0, 3, 2, 5, 6, 7, 4, 9, 8, 11, 10, 11, 12, 10}))
}

func maximumInvitations(favorite []int) int {
	n := len(favorite)
	rg := make([][]int, n)
	in := make([]int, n)
	for x, y := range favorite {
		rg[y] = append(rg[y], x)
		in[y]++
	}
	q := make([]int, 0)
	for x, ch := range in {
		if ch == 0 {
			q = append(q, x)
		}
	}
	for len(q) > 0 {
		x := q[0]
		q = q[1:]
		x = favorite[x]
		in[x]--
		if in[x] == 0 {
			q = append(q, x)
		}
	}
	ringSize, sumChainSize := 0, 0

	for st, ch := range in {
		if ch <= 0 {
			continue
		}
		ring := make([]int, 0)
		x := st
		for {
			ring = append(ring, x)
			in[x] = -1 // 把环的入度改了
			x = favorite[x]
			if x == st {
				break
			}
		}

		if len(ring) > 2 {
			ringSize = max(ringSize, len(ring))
		} else if len(ring) == 2 {
			sumChainSize += finMaxDep(rg, in, ring[0]) + finMaxDep(rg, in, ring[1]) + 2
		}
	}
	return max(ringSize, sumChainSize)
}

func finMaxDep(rg [][]int, in []int, x int) int {
	ans := 0
	for _, y := range rg[x] {
		if x == y {
			continue
		}
		if in[y] != 0 {
			continue
		}
		ans = max(ans, finMaxDep(rg, in, y)+1)
	}
	return ans
}
