package main

func main() {

}

// queries[i] = [typei, indexi, vali] 。 0：全部行，1全部列
// 可以得到结果，会超时
func matrixSumQueries1(n int, queries [][]int) int64 {
	m := len(queries)
	ans := 0
	used := make([][]int, n)
	for i := range used {
		used[i] = make([]int, n)
	}
	for i := m - 1; i >= 0; i-- {
		t, idx, va := queries[i][0], queries[i][1], queries[i][2]
		res := 0
		if t == 0 {
			for j := 0; j < n; j++ {
				if used[idx][j] == 0 {
					res += va
					used[idx][j] = 1
				}
			}
		}
		if t == 1 {
			for j := 0; j < n; j++ {
				if used[j][idx] == 0 {
					res += va
					used[j][idx] = 1
				}
			}
		}
		ans += res
	}
	return int64(ans)
}

func matrixSumQueries(n int, queries [][]int) int64 {
	m := len(queries)
	ans := 0
	col, row := make(map[int]bool), make(map[int]bool)

	for i := m - 1; i >= 0; i-- {
		t, idx, va := queries[i][0], queries[i][1], queries[i][2]
		if t == 0 {
			if !col[idx] {
				// 看这一列中，还有多少行没有填
				ans += (n - len(row)) * va
			}
			col[idx] = true
		}
		if t == 1 {
			if !row[idx] {
				// 看这一行中，还有多少列没有填
				ans += (n - len(col)) * va
			}
			row[idx] = true
		}
	}
	return int64(ans)
}
