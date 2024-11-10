package main

func main() {

}

func isValidSudoku(board [][]byte) bool {
	for i := 0; i < 9; i++ {
		cnt := make(map[byte]int)
		for j := 0; j < 9; j++ {
			b := board[i][j]
			if b == '.' {
				continue
			}
			if cnt[b] > 0 {
				return false
			}
			cnt[b]++
		}
	}
	for j := 0; j < 9; j++ {
		cnt := make(map[byte]int)
		for i := 0; i < 9; i++ {
			b := board[i][j]
			if b == '.' {
				continue
			}
			if cnt[b] > 0 {
				return false
			}
			cnt[b]++
		}
	}

	cnt := make([][]map[byte]int, 3)
	for i := range cnt {
		cnt[i] = make([]map[byte]int, 3)
		for j := range cnt[i] {
			cnt[i][j] = make(map[byte]int)
		}
	}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			b := board[i][j]
			if b == '.' {
				continue
			}
			if cnt[i/3][j/3][b] > 0 {
				return false
			}
			cnt[i/3][j/3][b]++

		}
	}
	return true
}
