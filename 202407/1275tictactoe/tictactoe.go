package main

import "fmt"

func main() {
	fmt.Println(tictactoe([][]int{{2, 2}, {0, 2}, {1, 0}, {0, 1}, {2, 0}, {0, 0}}))
}

func tictactoe(moves [][]int) string {
	moves = moves[:min(9, len(moves))]
	a := make([][]int, 0)
	for i := 0; i < len(moves); i += 2 {
		a = append(a, moves[i])
	}
	b := make([][]int, 0)
	for i := 1; i < len(moves); i += 2 {
		b = append(b, moves[i])
	}
	ai, bi := -1, -1
	for i := range a {
		if check(a[:i+1]) {
			ai = i + 1
			break
		}
	}
	for i := range b {
		if check(b[:i+1]) {
			bi = i + 1
			break
		}
	}
	if ai == -1 && bi == -1 {
		if len(moves) < 9 {
			return "Pending"
		} else {
			return "Draw"
		}
	}
	if ai != -1 {
		return "A"
	}
	return "B"

}

func check(arr [][]int) bool {
	g := make([][]int, 3)
	for i := 0; i < 3; i++ {
		g[i] = make([]int, 3)
	}

	for _, ch := range arr {
		x, y := ch[0], ch[1]
		g[x][y] = 1
	}

	for i := 0; i < 3; i++ {
		cnt := 0
		for j := 0; j < 3; j++ {
			cnt += g[i][j]
		}
		if cnt == 3 {
			return true
		}
	}
	for j := 0; j < 3; j++ {
		cnt := 0
		for i := 0; i < 3; i++ {
			cnt += g[i][j]
		}
		if cnt == 3 {
			return true
		}
	}
	if g[0][0]+g[1][1]+g[2][2] == 3 {
		return true
	}
	if g[0][2]+g[1][1]+g[2][0] == 3 {
		return true
	}
	return false
}
