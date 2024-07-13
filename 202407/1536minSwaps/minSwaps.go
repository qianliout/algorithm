package main

func main() {

}

func minSwaps(grid [][]int) int {
	n := len(grid)
	row := make([]int, n)
	for i := 0; i < n; i++ {
		row[i] = cal1(grid[i]) // 每一行从后往前数有多少个连续0
	}
	ans := 0
	for i := 0; i < n; i++ {
		if row[i] >= n-1-i { // 说明这一行已经符合条件了
			continue
		}
		j := i + 1
		for ; j < n; j++ {
			if row[j] >= n-1-i {
				break
			}
		}
		// 找到最后都没有找到，说明不能完成
		if j >= n {
			return -1
		}
		// 从下到上交换
		for k := j; k > i; k-- {
			row[k], row[k-1] = row[k-1], row[k]
			ans++
		}
	}

	return ans
}

func cal1(row []int) int {
	cnt := 0
	n := len(row)
	for i := n - 1; i >= 0; i-- {
		if row[i] == 0 {
			cnt++
		} else {
			break
		}
	}
	return cnt
}
