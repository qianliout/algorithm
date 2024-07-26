package main

func main() {

}

// 未完成
func smallestChair(times [][]int, targetFriend int) int {
	n := len(times)
	pairs := make([]pair, n)
	for i, ch := range times {
		pairs[i] = pair{ID: i, arrival: ch[0], leaving: ch[1]}
	}
	return 0
}

type pair struct {
	ID      int
	arrival int
	leaving int
}
