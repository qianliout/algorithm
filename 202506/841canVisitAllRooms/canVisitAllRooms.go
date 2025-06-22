package main

func main() {

}

func canVisitAllRooms(rooms [][]int) bool {
	n := len(rooms)
	visit := make([]bool, n)
	visit[0] = true

	var dfs func(start int)
	find := 0
	dfs = func(start int) {
		find++
		for _, j := range rooms[start] {
			if !visit[j] {
				visit[j] = true
				dfs(j)
			}
		}
	}
	dfs(0)
	return find == n
}
