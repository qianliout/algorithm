package main

import (
	"fmt"
)

func main() {
	fmt.Println(reversePairs([]int{9, 7, 5, 4, 6}))
	fmt.Println(reversePairs([]int{7, 5, 9, 4}))
	fmt.Println(reversePairs2([]int{7, 5, 9, 4}))
}

// 暴力解法,时间复杂度高
func reversePairs2(record []int) int {
	cnt := 0
	n := len(record)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if record[i] > record[j] {
				cnt++
			}
		}
	}
	return cnt
}

// 错的 [7,5,9,4]
func reversePairs3(record []int) int {
	st := make([]int, 0)
	cnt := 0
	for _, ch := range record {
		for len(st) > 0 && st[len(st)-1] <= ch {
			st = st[:len(st)-1]
		}
		cnt += len(st)
		st = append(st, ch)
	}
	return cnt
}

// 归并排序的思想
func reversePairs(record []int) int {
	var merge func(i, j int) int
	n := len(record)
	merge = func(left, right int) int {
		if left >= right {
			return 0
		}

		tem := make([]int, 0)
		mid := left + (right-left)/2
		a := merge(left, mid)
		b := merge(mid+1, right)
		ans := a + b
		i, j := left, mid+1
		for i <= mid && j <= right {
			if record[i] <= record[j] {
				ans += j - (mid + 1)
				tem = append(tem, record[i])
				i++
			} else {
				tem = append(tem, record[j])
				j++
			}
		}
		for ; i <= mid; i++ {
			ans += j - (mid + 1)
			tem = append(tem, record[i])
		}
		for ; j <= right; j++ {
			tem = append(tem, record[j])
		}
		for k := left; k <= right; k++ {
			record[k] = tem[k-left]
		}
		return ans
	}

	ans := merge(0, n-1)
	return ans
}

// https://leetcode.cn/problems/shu-zu-zhong-de-ni-xu-dui-lcof/solutions/216984/shu-zu-zhong-de-ni-xu-dui-by-leetcode-solution/?envType=problem-list-v2&envId=8LSpuXqD
