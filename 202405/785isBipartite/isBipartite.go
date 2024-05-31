package main

func main() {

}

func isBipartite(graph [][]int) bool {
	n := len(graph)
	visit := make([]int, n) // 0 未染色，1蓝色 -1黄色
	queue := make([]int, 0)
	// 把所有点都加进去测试一下，都能二分，才算是能二分
	for i := 0; i < n; i++ {
		if visit[i] != 0 {
			continue
		}
		queue = append(queue, i)
		visit[i] = 1
		for len(queue) > 0 {
			fir := queue[0]
			queue = queue[1:]
			curColor := visit[fir]
			nextColor := -curColor

			next := graph[fir]

			for _, ne := range next {
				if visit[ne] == 0 {
					visit[ne] = nextColor
					queue = append(queue, ne)
				} else if visit[ne] != nextColor {
					return false
				}
			}
		}
	}
	return true
}
