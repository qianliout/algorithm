package main

func main() {

}

func canVisitAllRooms(rooms [][]int) bool {
	n := len(rooms)
	cnt := 0
	has := make([]bool, n)
	has[0] = true
	q := []int{0}
	for len(q) > 0 {
		fir := q[0]
		q = q[1:]
		cnt++
		for _, j := range rooms[fir] {
			if !has[j] {
				has[j] = true
				q = append(q, j)
			}
		}
	}
	return cnt == n
}
