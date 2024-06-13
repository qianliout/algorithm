package main

import (
	"container/heap"
	"fmt"

	. "outback/algorithm/common/commonHeap"
)

func main() {
	fmt.Println(earliestSecondToMarkIndices([]int{3, 2, 3}, []int{1, 3, 2, 2, 2, 2, 3}))
	fmt.Println(earliestSecondToMarkIndices([]int{1, 2, 3}, []int{1, 2, 3}))
	fmt.Println(earliestSecondToMarkIndices([]int{0}, []int{1}))
}

func earliestSecondToMarkIndices(nums []int, changeIndices []int) int {
	n, m, total := len(nums), len(changeIndices), 0
	for _, ch := range nums {
		total += ch
	}
	total += n

	first := make([]int, n)
	for i := range first {
		first[i] = -1
	}

	for i := m - 1; i >= 0; i-- {
		v := changeIndices[i]
		first[v-1] = i // 最开始出现的位置
	}

	var check func(n int) bool
	// 最多 mx 天能搞定所有的考试及复习不
	check = func(mx int) bool {
		// mx += n
		cnt := 0      // 还有的可用天数
		slow := total // 初始任务所有的课程都需要慢速复习(也就是题目中的减一操作)

		mh := make(MinHeap, 0)

		for t := mx - 1; t >= 0; t-- {
			i := changeIndices[t] - 1
			v := nums[i]
			if v <= 1 || first[i] != t {
				cnt++
				continue
			}
			if cnt == 0 {
				if len(mh) == 0 || v <= mh[0] {
					// 没有办反悔
					cnt++
					continue
				}
				pop := heap.Pop(&mh).(int)
				slow += pop + 1
				cnt += 2
			}
			slow -= v + 1
			cnt--
			heap.Push(&mh, v)
		}
		return cnt >= slow
	}

	// ans := n + sort.Search(m+1-n, check)
	// if ans > m {
	// 	return -1
	// }
	// return ans

	// 二分
	le, ri := n, m+1

	for le < ri {
		mid := le + (ri-le)/2
		if mid >= n && mid < m+1 && check(mid) {
			ri = mid
		} else {
			le = mid + 1
		}
	}
	if le > m {
		return -1
	}
	return le
}
