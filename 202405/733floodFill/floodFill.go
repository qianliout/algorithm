package main

func main() {

}

func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	if len(image) == 0 || len(image[0]) == 0 {
		return image
	}
	visit := make([][]bool, len(image))
	for i := range visit {
		visit[i] = make([]bool, len(image[i]))
	}

	dfs(image, sr, sc, color, image[sr][sc], visit)

	return image
}

var dirs = [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}

func dfs(images [][]int, sr, sc, color int, pre int, visit [][]bool) {
	if sr < 0 || sr >= len(images) || sc < 0 || sc >= len(images[sr]) {
		return
	}
	if images[sr][sc] != pre {
		return
	}
	images[sr][sc] = color
	visit[sr][sc] = true
	for _, dir := range dirs {
		nx, ny := sr+dir[0], sc+dir[1]
		if nx < 0 || nx >= len(images) || ny < 0 || ny >= len(images[nx]) {
			return
		}
		if visit[nx][ny] {
			continue
		}
		dfs(images, nx, ny, color, pre, visit)
	}
}
