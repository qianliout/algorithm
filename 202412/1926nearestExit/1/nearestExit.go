package main

func main() {

}

func nearestExit(maze [][]byte, entrance []int) int {
	m, n := len(maze), len(maze[0])
	ans := 0
	ent := node{entrance[0], entrance[1]}
	q := []node{ent}
	dirs := [][]int{[]int{0, 1}, []int{0, -1}, []int{1, 0}, []int{-1, 0}}
	used := make([][]bool, m)
	for i := range used {
		used[i] = make([]bool, n)
	}
	used[ent.x][ent.y] = true

	for len(q) > 0 {
		lev := make([]node, 0)

		for _, no := range q {
			for _, dir := range dirs {
				x, y := no.x+dir[0], no.y+dir[1]

				nod1 := node{x, y}

				if in(maze, nod1) {
					if used[x][y] {
						continue
					}

					if exit(maze, nod1) && ent != nod1 {
						return ans + 1
					}
					lev = append(lev, nod1)
					used[x][y] = true
				}
			}
		}
		if len(lev) > 0 {
			ans++
		}
		q = lev

	}
	return -1
}

type node struct {
	x, y int
}

func exit(maze [][]byte, no node) bool {
	if !in(maze, no) {
		return false
	}
	m, n := len(maze), len(maze[0])
	if no.x == 0 || no.x == m-1 || no.y == 0 || no.y == n-1 {
		return true
	}
	return false
}

func in(maze [][]byte, no node) bool {
	m, n := len(maze), len(maze[0])
	x, y := no.x, no.y
	if x < 0 || x >= m || y < 0 || y >= n {
		return false
	}
	if maze[x][y] == '+' {
		return false
	}
	return true
}
