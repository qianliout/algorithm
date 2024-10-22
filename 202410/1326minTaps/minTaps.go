package main

// minTaps1 计算打开最少的水龙头数量以浇灌整个花园。
// n 表示花园的长度，ranges 表示每个水龙头的覆盖范围。
func minTaps1(n int, ranges []int) int {
	// rightMost数组用于记录每个位置上最远的覆盖范围。
	rightMost := make([]int, n+1)
	for i, ch := range ranges {
		// 计算并更新每个位置的最远覆盖范围。
		left := max(0, i-ch)
		rightMost[left] = max(rightMost[left], i+ch)
	}

	// cur表示当前覆盖的最远位置，nex表示下一步能覆盖的最远位置，ans记录打开的水龙头数量。
	cur, nex, ans := 0, 0, 0
	for i := 0; i < n; i++ {
		// 更新下一步能覆盖的最远位置。
		nex = max(nex, rightMost[i])
		// 当到达当前覆盖的最远位置时，检查是否需要打开新的水龙头。
		if cur == i {
			// 如果当前覆盖范围无法延伸，且没有新的水龙头可以打开，则无法完成浇灌。
			if cur == nex {
				return -1
			}
			// 打开新的水龙头，并更新当前覆盖的最远位置。
			ans++
			cur = nex
		}
	}
	// 返回打开的水龙头数量。
	return ans
}

// minTaps 计算打开最少的水龙头数量以浇灌整个花园。
// n 表示花园的长度，ranges 表示每个水龙头的覆盖范围。
func minTaps(n int, ranges []int) int {
	// rightMost数组用于记录每个位置上最远的覆盖范围。
	rightMost := make([]int, n+1)
	for i, ch := range ranges {
		// 计算并更新每个位置的最远覆盖范围。
		left := max(0, i-ch)
		rightMost[left] = max(rightMost[left], i+ch)
	}

	// # 已建造的桥的右端点，nex表示下一座桥的右端点的最大值
	cur, nex, ans := 0, 0, 0
	// 如果走到 n-1 时没有返回 -1，那么必然可以到达 n
	for i := 0; i < n; i++ {
		// 更新下一步能覆盖的最远位置。
		nex = max(nex, rightMost[i])
		// 到达已建造的桥的右端点
		if cur == i {
			// 无论怎么造桥，都无法从 i 到 i+1
			if cur == nex {
				return -1
			}
			// 造就一个新桥
			ans++
			cur = nex
		}
	}
	return ans
}
