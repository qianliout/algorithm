package main

func main() {

}

func shortestPathBinaryMatrix(grid [][]int) int {
	n := len(grid)
	if grid[0][0] == 1 || grid[n-1][n-1] == 1 {
		return -1
	}
	q := []node{{0, 0, 1}}
	dirs := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}, {-1, -1}, {1, -1}, {-1, 1}, {1, 1}}
	for len(q) > 0 {
		lev := make([]node, 0)
		for _, fir := range q {
			// 以后写 bfs 都要在这里判断,这样才不会出错
			// 如果在下面判断，当开始节点和结束节点是一个时就出错了
			if fir.x == n-1 && fir.y == n-1 {
				return fir.dis
			}

			for _, dir := range dirs {
				nx, ny := fir.x+dir[0], fir.y+dir[1]
				if nx < 0 || nx >= n || ny < 0 || ny >= n {
					continue
				}
				if grid[nx][ny] == 1 {
					continue
				}
				// 可以在这里判断，但是容易出错，当开始节点和结束节点是一个时，就得不到正确的解
				// if nx == n-1 && ny == n-1 {
				// 	return fir.dis + 1
				// }

				grid[nx][ny] = 1
				lev = append(lev, node{nx, ny, fir.dis + 1})
			}
		}
		q = lev
	}
	return -1
}

type node struct {
	x, y int
	dis  int
}
