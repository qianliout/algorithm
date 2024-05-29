package main

func main() {

}

func maxCount(m int, n int, ops [][]int) int {
	for _, ch := range ops {
		m = min(m, ch[0])
		n = min(n, ch[1])
	}
	return m * n
}
