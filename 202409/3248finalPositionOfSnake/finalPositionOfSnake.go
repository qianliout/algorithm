package main

func main() {

}

func finalPositionOfSnake(n int, commands []string) int {
	// 生成的测评数据确保蛇不会移动到矩阵的边界外
	i, j := 0, 0
	for _, str := range commands {
		switch str[0] {
		case 'U':
			i--
		case 'D':
			i++
		case 'L':
			j--
		case 'R':
			j++
		}
	}
	return i*n + j
}
