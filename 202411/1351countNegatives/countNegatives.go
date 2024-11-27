package main

import (
	"fmt"
)

func main() {
	fmt.Println(countNegatives([][]int{{7, 1}}))
	fmt.Println(countNegatives([][]int{{-7, -11}}))
	fmt.Println(countNegatives2([][]int{{7, 1}}))
	fmt.Println(countNegatives2([][]int{{-7, -11}}))
	// fmt.Println(countNegatives([][]int{{7, 2}}))
	// fmt.Println(countNegatives([][]int{{3, 2}, {1, 0}}))
	// fmt.Println(countNegatives([][]int{{4, 3, 2, -1}, {3, 2, 1, -1}, {1, 1, -1, -2}, {-1, -1, -2, -3}}))
}

func countNegatives2(grid [][]int) int {
	n, m := len(grid), len(grid[0])
	cnt := 0
	for i := 0; i < n; i++ {
		le, ri := 0, m
		for le < ri {
			mid := le + (ri-le)/2
			if mid >= 0 && mid < m && grid[i][mid] < 0 {
				ri = mid
			} else {
				le = mid + 1
			}
		}
		fmt.Println("left:", le)
		cnt += max(0, m-le)
	}

	return cnt
}

func countNegatives(grid [][]int) int {
	n, m := len(grid), len(grid[0])
	cnt := 0
	for i := 0; i < n; i++ {
		le, ri := 0, m
		for le < ri {
			mid := le + (ri-le+1)/2
			if mid >= 0 && mid < m && grid[i][mid] >= 0 {
				le = mid
			} else {
				ri = mid - 1
			}
		}
		fmt.Println("left:", le)
		if le < m && le >= 0 {
			if grid[i][le] < 0 {
				cnt += max(0, m-le)
			} else {
				cnt += max(0, m-le-1)
			}
		}
	}
	return cnt
}

func nextGreatestLetter(letters []byte, target byte) byte {
	n := len(letters)
	le, ri := 0, n
	for le < ri {
		mid := le + (ri-le)/2
		if mid >= 0 && mid < n && letters[mid] > target {
			ri = mid
		} else {
			le = mid + 1
		}
	}
	if le < 0 || le >= n || letters[le] <= target {
		return letters[0]
	}
	return letters[le]
}
