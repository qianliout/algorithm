package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(maximumMinutes([][]int{{0, 0, 0}, {2, 2, 0}, {1, 2, 0}}))
}

func maximumMinutes(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	fire := make([][]int, m)
	people := make([][]int, m)
	dirs := [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	fireq := make([]pair, 0)

	for i := range fire {
		fire[i] = make([]int, n)
		people[i] = make([]int, n)
		for j := range fire[i] {
			fire[i][j] = -1
			people[i][j] = -1
		}
	}

	for i := range grid {
		for j, ch := range grid[i] {
			if ch == 1 {
				fireq = append(fireq, pair{i, j, 0})
				fire[i][j] = 0
			}
		}
	}

	for len(fireq) > 0 {
		first := fireq[0]
		fireq = fireq[1:]
		for _, dir := range dirs {
			nx, ny := first.x+dir[0], first.y+dir[1]
			if nx < 0 || ny < 0 || nx >= m || ny >= n || grid[nx][ny] != 0 || fire[nx][ny] != -1 {
				continue
			}
			fire[nx][ny] = first.sec + 1
			fireq = append(fireq, pair{nx, ny, first.sec + 1})
		}
	}
	peopleq := []pair{{x: 0, y: 0, sec: 0}}
	people[0][0] = 0
	for len(peopleq) > 0 {
		first := peopleq[0]
		peopleq = peopleq[1:]
		for _, dir := range dirs {
			nx, ny := first.x+dir[0], first.y+dir[1]
			if nx < 0 || ny < 0 || nx >= m || ny >= n || grid[nx][ny] != 0 || people[nx][ny] != -1 {
				continue
			}
			people[nx][ny] = first.sec + 1
			peopleq = append(peopleq, pair{nx, ny, first.sec + 1})
		}
	}
	// 然后就分情况讨论了
	man2House, m1, m2 := people[m-1][n-1], people[m-1][n-2], people[m-2][n-1]
	fir2House, f1, f2 := fire[m-1][n-1], fire[m-1][n-2], fire[m-2][n-1]
	if man2House < 0 {
		return -1
	}
	if fir2House < 0 {
		return int(math.Pow10(9))
	}

	d := fir2House - man2House
	if d < 0 {
		return -1
	}
	if (m1 != -1 && m1+d < f1) || (m2 != -1 && m2+d < f2) {
		return d
	}
	return d - 1
}

type pair struct {
	x, y int
	sec  int
}
