package main

func main() {

}

func coloredCells(n int) int64 {
	a := 2*n - 1
	b := (1 + a) * n / 2
	return int64(2*b - a)
}
